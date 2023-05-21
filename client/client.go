package main

import (
	"context"
	pb "github.com/Abdulrahman-Awadh/Movirate/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error", err)
	}

	client := pb.NewMovieApiClient(conn)

	// Note how we are calling the GetBookList method on the server
	// This is available to us through the auto-generated code
	movies, err := client.GetMovies(context.Background(), &pb.RequestMovies{})
	if err != nil {
		log.Fatalf("Error", err)
	}

	log.Printf("Movie list: %v", movies)

}
