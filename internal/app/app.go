package app

import (
	"fmt"
	"github.com/izaakdale/distcache/internal/consensus"
	"github.com/izaakdale/distcache/internal/discovery"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/izaakdale/distcache/api/v1"
	"github.com/izaakdale/distcache/internal/store"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	spec specification
)

type specification struct {
	GRPCHost  string `envconfig:"GRPC_HOST"`
	GRCPPort  int    `envconfig:"GRPC_PORT"`
	RedisAddr string `envconfig:"REDIS_ADDR"`
	Password  string `envconfig:"PASSWORD"`
	DB        int    `envconfig:"DB"`
	RecordTTL int    `envconfig:"RECORD_TTL"`

	discoveryCfg discovery.Specification
}

type App struct {
	ln   net.Listener
	gsrv *grpc.Server
}

func New() (*App, error) {
	if err := envconfig.Process("", &spec); err != nil {
		return nil, err
	}
	if err := envconfig.Process("", &spec.discoveryCfg); err != nil {
		return nil, err
	}

	cli, err := store.New(
		store.WithConfig(store.Config{
			RedisAddr: spec.RedisAddr,
			Password:  spec.Password,
			DB:        spec.DB,
			RecordTTL: spec.RecordTTL,
		}),
	)
	if err != nil {
		panic(err)
	}

	////TODO fill in new dist
	//dist, err := consensus.NewDistributedCache("test", consensus.Config{
	//	Txer: cli,
	//	Raft: &consensus.Raft{
	//		Config:      raft.Config{},
	//		BindAddr:    "",
	//		StreamLayer: nil,
	//		Bootstrap:   false,
	//	},
	//})
	//if err != nil {
	//	panic(err)
	//}
	////TODO
	//_ = dist

	gAddr := fmt.Sprintf("%s:%d", spec.GRPCHost, spec.GRCPPort)
	ln, err := net.Listen("tcp", gAddr)
	if err != nil {
		return nil, err
	}

	srv := Server{
		Txer: cli,
	}
	// for mTLS, leaving for now since I want to test via ingress
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

	v1.RegisterCacheServer(gsrv, &srv)

	return &App{
		ln:   ln,
		gsrv: gsrv,
	}, nil
}

func (a *App) Run() error {
	membership, err := discovery.NewMembership(&consensus.DistributedCache{}, spec.discoveryCfg, spec.GRCPPort)
	if err != nil {
		return err
	}

	// create error channel for the grpc server to relay information back to app
	errCh := make(chan error)
	go func(ch chan error) {
		ch <- a.gsrv.Serve(a.ln)
	}(errCh)

	defer func(ch chan error) {
		ch <- membership.Node.Leave()
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
		case ev := <-membership.EventChan:
			membership.HandleSerfEvent(ev, membership.Node)
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
