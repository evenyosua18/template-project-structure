package grpc

import (
	"fmt"
	"github.com/evenyosua18/template-project-structure/app/infrastructure/proto/pb"
	exampleSvc "github.com/evenyosua18/template-project-structure/app/infrastructure/server/grpc/service/example"
	"github.com/evenyosua18/template-project-structure/app/usecase"
	exampleUc "github.com/evenyosua18/template-project-structure/app/usecase/example"
	"github.com/evenyosua18/template-project-structure/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer() {
	/*start service*/

	//get config
	cfg := config.GetConfig().Server

	//config grpc server
	var options []grpc.ServerOption
	//maxIdle, err := time.ParseDuration(cfg.Grpc.MaxIdle + "h")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//maxAge, err := time.ParseDuration(cfg.Grpc.MaxAge + "m")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//maxAgeGrace, err := time.ParseDuration(cfg.Grpc.MaxAge + "s")
	//
	//if err != nil {
	//	panic(err)
	//}
	options = append(options, grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     1 * time.Hour,
		MaxConnectionAge:      5 * time.Minute,
		MaxConnectionAgeGrace: 5 * time.Second,
	}))

	//create grpc server
	grpcServer := grpc.NewServer(options...)

	//setup dependency injection for use case (interaction)
	ctn := usecase.NewContainer()

	//register grpc server
	Apply(grpcServer, ctn)
	reflection.Register(grpcServer)

	//run grpc server
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(`%s:%d`, cfg.Grpc.Host, cfg.Grpc.Port))

		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to start grpc server: %v", err)
		}
	}()

	log.Println(fmt.Sprintf("grpc server is running at %s:%d", cfg.Grpc.Host, cfg.Grpc.Port))

	//get signal when server interrupted
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	log.Fatalf("process killed with signal: %s", sig.String())
}

// Apply register all service here
func Apply(server *grpc.Server, ctn *usecase.Container) {
	pb.RegisterExampleServiceServer(server, exampleSvc.NewServiceExample(ctn.Resolve("exampleCTN").(*exampleUc.InteractionExample)))
}
