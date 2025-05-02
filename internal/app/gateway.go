package app

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"

	pb "github.com/webbsalad/pet/internal/pb/github.com/webbsalad/pet"
)

var (
	gwPort = flag.Int("gw_port", 8080, "HTTP port")
)

func newCORSHandler(h http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}).Handler(h)
}

func gatewayOption() fx.Option {
	flag.Parse()

	return fx.Invoke(func(lc fx.Lifecycle) {
		mux := runtime.NewServeMux(
			runtime.WithForwardResponseOption(func(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
				if smd, ok := runtime.ServerMetadataFromContext(ctx); ok {
					if vals := smd.HeaderMD.Get("http-code"); len(vals) > 0 {
						if code, err := strconv.Atoi(vals[0]); err == nil {
							w.WriteHeader(code)
						} else {
							return err
						}
					}
				}
				return nil
			}),
		)

		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		endpoint := fmt.Sprintf("localhost:%d", *grpcPort)

		if err := pb.RegisterPetServiceHandlerFromEndpoint(
			context.Background(),
			mux,
			endpoint,
			opts,
		); err != nil {
			log.Fatalf("failed register PVZ gateway: %v", err)
		}

		handler := newCORSHandler(mux)

		srv := &http.Server{
			Addr:    fmt.Sprintf(":%d", *gwPort),
			Handler: handler,
		}

		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
						log.Fatalf("failed http server starting: %v", err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Println("stopping HTTP gateway")
				return srv.Shutdown(ctx)
			},
		})
	})
}
