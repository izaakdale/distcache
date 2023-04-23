package app

import (
	"fmt"
	"log"
	"net"

	msg "github.com/izaakdale/distcache/api/v1"
	"github.com/izaakdale/distcache/internal/store"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var spec specification

type specification struct {
	GRPCHost  string `envconfig:"GRPC_HOST"`
	GRCPPort  int    `envconfig:"GRPC_PORT"`
	RedisAddr string `envconfig:"REDIS_ADDR"`
	Password  string `envconfig:"PASSWORD"`
	DB        int    `envconfig:"DB"`
	RecordTTL int    `envconfig:"RECORD_TTL"`
}

type App struct {
	ln   net.Listener
	gsrv *grpc.Server
}

func (a *App) Run() error {
	log.Printf("running on: %s", a.ln.Addr().String())
	return a.gsrv.Serve(a.ln)
}

func New() (*App, error) {
	if err := envconfig.Process("", &spec); err != nil {
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

	srv := Server{}
	gsrv := grpc.NewServer()
	reflection.Register(gsrv)

	msg.RegisterCacheServer(gsrv, &srv)

	gAddr := fmt.Sprintf("%s:%d", spec.GRPCHost, spec.GRCPPort)
	ln, err := net.Listen("tcp", gAddr)
	if err != nil {
		return nil, err
	}

	return &App{
		ln:   ln,
		gsrv: gsrv,
	}, nil
}

func Must(a *App, err error) *App {
	if err != nil {
		log.Fatalf("error initialising: %e", err)
	}
	return a
}
