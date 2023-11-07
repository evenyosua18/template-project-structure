package grpc

import (
	"fmt"
	"github.com/evenyosua18/template-project-structure/app/constant"
	"github.com/evenyosua18/template-project-structure/app/infrastructure/container"
	"github.com/evenyosua18/template-project-structure/app/infrastructure/proto/pb"
	"github.com/evenyosua18/template-project-structure/app/infrastructure/server/grpc/service/user"
	"github.com/joho/godotenv"
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

	//setup environment variable (for local)
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(".env"); err != nil {
			panic(err)
		}
	}

	//config grpc server
	var options []grpc.ServerOption
	options = append(options, grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     1 * time.Hour,
		MaxConnectionAge:      5 * time.Minute,
		MaxConnectionAgeGrace: 5 * time.Second,
	}))

	//create grpc server
	grpcServer := grpc.NewServer(options...)

	//register grpc server
	Apply(grpcServer)
	reflection.Register(grpcServer)

	//run grpc server
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(`%s:%s`, os.Getenv(constant.GrpcHost), os.Getenv(constant.GrpcPort)))

		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to start grpc server: %v", err)
		}
	}()

	log.Println(fmt.Sprintf("grpc server is running at %s:%s", os.Getenv(constant.GrpcHost), os.Getenv(constant.GrpcPort)))

	//get signal when server interrupted
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	log.Fatalf("process killed with signal: %s", sig.String())
}

// Apply register all service here
func Apply(server *grpc.Server) {
	pb.RegisterUserServiceServer(server, user.NewUserService(container.InitializeUserInteraction()))
}
