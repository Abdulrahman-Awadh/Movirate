package main

import (
	"github.com/Abdulrahman-Awadh/Movirate/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	databaseInit()
	//server setup
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("can not create listner: %s", err)
	}
	grpcServer := grpc.NewServer()
	movie.RegisterMovieApiServer(grpcServer, &movieServer{})

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("can not Serve : %s", err)
	}

}
