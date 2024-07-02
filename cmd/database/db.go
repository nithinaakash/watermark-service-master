package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/aayushrangwala/watermark-service/internal/database"
	dbsvc "github.com/aayushrangwala/watermark-service/pkg/database"
	"github.com/aayushrangwala/watermark-service/pkg/database/endpoints"
	"github.com/aayushrangwala/watermark-service/pkg/database/transport"

	"github.com/go-kit/kit/log"
	// kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/oklog/oklog/pkg/group"
	// "google.golang.org/grpc"
)

const (
	defaultHTTPPort = "8081"
	defaultGRPCPort = "8082"
)

var (
	logger   log.Logger
	httpAddr = net.JoinHostPort("localhost", envString("HTTP_PORT", defaultHTTPPort))
	grpcAddr = net.JoinHostPort("localhost", envString("GRPC_PORT", defaultGRPCPort))
	db       *gorm.DB
)

func main() {
	var err error
	db, err = database.Init(database.DefaultHost, database.DefaultPort, database.DefaultDBUser, database.DefaultDatabase, database.DefaultPassword)
	if err != nil {
		logger.Log("FATAL", fmt.Sprintf("failed to load db with error: %s", err.Error()))
		os.Exit(1)
	}
	defer db.Close()

	var (
		service     = dbsvc.NewService(db)
		eps         = endpoints.NewEndpointSet(service)
		httpHandler = transport.NewHTTPHandler(eps)
		// grpcServer  = transport.NewGRPCServer(eps)f
	)

	var g group.Group
	{
		// The HTTP listener mounts the Go kit HTTP handler we created.
		httpListener, err := net.Listen("tcp", httpAddr)
		if err != nil {
			logger.Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "HTTP", "addr", httpAddr)
			return http.Serve(httpListener, httpHandler)
		}, func(error) {
			httpListener.Close()
		})
	}
	// {
	// 	// The gRPC listener mounts the Go kit gRPC server we created.
	// 	grpcListener, err := net.Listen("tcp", grpcAddr)
	// 	if err != nil {
	// 		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	// 		os.Exit(1)
	// 	}
	// 	g.Add(func() error {
	// 		logger.Log("transport", "gRPC", "addr", grpcAddr)
	// 		// we add the Go Kit gRPC Interceptor to our gRPC service as it is used by
	// 		// the here demonstrated zipkin tracing middleware.
	// 		baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	// 		pb.RegisterDatabaseServer(baseServer, grpcServer)
	// 		return baseServer.Serve(grpcListener)
	// 	}, func(error) {
	// 		grpcListener.Close()
	// 	})
	// }
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
