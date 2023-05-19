package main

import (
	"context"
	"github.com/Abdulrahman-Awadh/Movirate/movie"
	pb "github.com/Abdulrahman-Awadh/Movirate/movie"
	"google.golang.org/grpc"
	"log"
	"net"
)

// no idea yet
type movieServer struct {
	pb.UnimplementedMovieApiServer
}

// this to implement the interface of the endpoint
func (s *movieServer) GetMovie(ctx context.Context, req *pb.RequestMovies) (*pb.ResponseMovie, error) {
	return &pb.ResponseMovie{
		Movie: &pb.Movie{
			Id:          1,
			Name:        "Intersteller",
			Poster:      "^_^",
			Description: "number 1 movie",
			Genre:       "perfect",
		},
	}, nil

}

func (s *movieServer) GetMovies(ctx context.Context, req *pb.RequestMovies) (*pb.ResponseMovies, error) {
	movies := getMoviesFromAPI()
	response := &pb.ResponseMovies{Movie: movies}

	return response, nil
}

func main() {
	connectToDB()
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
