package app

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/raft"
	config "github.com/izaakdale/distcache/internal/auth"
	"github.com/izaakdale/distcache/internal/consensus"
	"github.com/izaakdale/distcache/internal/discovery"
	cmux2 "github.com/soheilhy/cmux"
	"io"
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
	mux  cmux2.CMux
	ln   net.Listener
	gsrv *grpc.Server
	dist *consensus.DistributedCache
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

	gAddr := fmt.Sprintf("%s:%d", spec.GRPCHost, spec.GRCPPort)
	ln, err := net.Listen("tcp", gAddr)
	if err != nil {
		return nil, err
	}

	gsrv := grpc.NewServer() //grpc.Creds(creds))
	reflection.Register(gsrv)

	srv := Server{
		Txer: cli,
	}
	v1.RegisterCacheServer(gsrv, &srv)

	mux := cmux2.New(ln)
	raftLn := mux.Match(func(r io.Reader) bool {
		b := make([]byte, 1)
		if _, err := r.Read(b); err != nil {
			return false
		}
		return bytes.Compare(b, []byte{byte(consensus.RaftRPC)}) == 0
	})

	srvcfg, err := config.SetupTLSConfig(config.TLSConfig{
		CertFile:      config.ServerCertFile,
		KeyFile:       config.ServerKeyFile,
		CAFile:        config.CAFile,
		Server:        true,
		ServerAddress: raftLn.Addr().String(),
	})
	if err != nil {
		return nil, err
	}

	clicfg, err := config.SetupTLSConfig(config.TLSConfig{
		CertFile:      config.ClientCertFile,
		KeyFile:       config.ClientKeyFile,
		CAFile:        config.CAFile,
		Server:        false,
		ServerAddress: raftLn.Addr().String(),
	})
	if err != nil {
		return nil, err
	}

	dist, err := consensus.NewDistributedCache("test", consensus.Config{
		Txer: cli,
		Raft: &consensus.Raft{
			Config: raft.Config{
				LocalID: "TODO", //TODO
			},
			BindAddr:    gAddr,
			StreamLayer: consensus.NewStreamLayer(raftLn, srvcfg, clicfg),
			Bootstrap:   false,
		},
	})
	if err != nil {
		panic(err)
	}

	return &App{
		ln:   ln,
		gsrv: gsrv,
		dist: dist,
		mux:  mux,
	}, nil
}

func (a *App) Run() error {
	membership, err := discovery.NewMembership(a.dist, spec.discoveryCfg, spec.GRCPPort)
	if err != nil {
		return err
	}

	// create error channel for the grpc server to relay information back to app
	errCh := make(chan error)
	grpcLn := a.mux.Match(cmux2.Any())
	go func(ch chan error) {
		ch <- a.gsrv.Serve(grpcLn)
	}(errCh)
	go func(ch chan error) {
		ch <- a.mux.Serve()
	}(errCh)
	defer membership.Node.Leave()

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
