package app

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/serf/serf"
	msg "github.com/izaakdale/distcache/api/v1"
	"github.com/izaakdale/distcache/internal/store"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	spec              specification
	clusterMembership *serf.Serf
)

type specification struct {
	GRPCHost  string `envconfig:"GRPC_HOST"`
	GRCPPort  int    `envconfig:"GRPC_PORT"`
	RedisAddr string `envconfig:"REDIS_ADDR"`
	Password  string `envconfig:"PASSWORD"`
	DB        int    `envconfig:"DB"`
	RecordTTL int    `envconfig:"RECORD_TTL"`

	clutserCfg clusterSpec
}

type App struct {
	ln   net.Listener
	gsrv *grpc.Server
}

func New() (*App, error) {
	if err := envconfig.Process("", &spec); err != nil {
		return nil, err
	}
	if err := envconfig.Process("", &spec.clutserCfg); err != nil {
		return nil, err
	}

	if err := store.Init(
		store.WithConfig(
			store.Config{
				RedisAddr: spec.RedisAddr,
				Password:  spec.Password,
				DB:        spec.DB,
				RecordTTL: spec.RecordTTL,
			},
		),
	); err != nil {
		return nil, err
	}

	gAddr := fmt.Sprintf("%s:%d", spec.GRPCHost, spec.GRCPPort)
	ln, err := net.Listen("tcp", gAddr)
	if err != nil {
		return nil, err
	}

	srv := Server{}
	// cfg, err := config.SetupTLSConfig(config.TLSConfig{
	// 	CertFile:      config.ServerCertFile,
	// 	KeyFile:       config.ServerKeyFile,
	// 	CAFile:        config.CAFile,
	// 	Server:        true,
	// 	ServerAddress: ln.Addr().String(),
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// creds := credentials.NewTLS(cfg)

	gsrv := grpc.NewServer() //grpc.Creds(creds))
	reflection.Register(gsrv)

	msg.RegisterCacheServer(gsrv, &srv)

	return &App{
		ln:   ln,
		gsrv: gsrv,
	}, nil
}

func (a *App) Run() error {
	defer pool.Close()

	cluster, evCh, err := setupCluster(
		spec.clutserCfg.BindAddr,      // BIND defines where the agent listens for incomming connections
		spec.clutserCfg.BindPort,      // in k8s this would be the ip and port of the pod/container
		spec.clutserCfg.AdvertiseAddr, // ADVERTISE defines where the agent is reachable
		spec.clutserCfg.AdvertisePort, // in k8s this correlates to the cluster ip service
		spec.clutserCfg.Name,          // NAME must be unique, which is not possible for replicas with env vars. Uniqueness handled in setup
	)
	if err != nil {
		return err
	}
	clusterMembership = cluster
	defer cluster.Leave()

	// create error channel for the grpc server to relay information back to app
	errCh := make(chan error)
	go func(ch chan error) {
		ch <- a.gsrv.Serve(a.ln)
	}(errCh)

	// signal channel for the os/k8s
	shCh := make(chan os.Signal, 2)
	signal.Notify(shCh, os.Interrupt, syscall.SIGTERM)

Loop:
	for {
		select {
		// shutdown
		case <-shCh:
			break Loop
		// serf event channel
		case ev := <-evCh:
			handleSerfEvent(ev, cluster)
		// grpc error
		case err := <-errCh:
			return err
		}
	}
	return nil
}

func Must(a *App, err error) *App {
	if err != nil {
		log.Fatalf("error initialising: %e", err)
	}
	return a
}

func MustRun() {
	app := Must(New())
	if err := app.Run(); err != nil {
		log.Fatalf("error running: %e", err)
	}
}
