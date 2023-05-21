package main

import (
	"context"
	pb "github.com/Abdulrahman-Awadh/Movirate/proto"
)

type movieServer struct {
	pb.UnimplementedMovieApiServer
}

// this to implement the interface of the endpoint
func (s *movieServer) GetMovies(ctx context.Context, req *pb.RequestMovies) (*pb.ResponseMovies, error) {
	return &pb.ResponseMovies{
		Movie: getMoviesFromAPI(),
	}, nil
}

func (s *movieServer) AddMovieToFav(context.Context, *pb.RequestMAddMovieToFav) (*pb.ResponseAddMovieToFav, error) {
	return nil, nil
}

func (s *movieServer) DeleteMovieFromFav(context.Context, *pb.RequestDeleteMovieFromFav) (*pb.ResponseDeleteMovieFromFav, error) {
	return nil, nil
}

func (s *movieServer) SearchForMovie(context.Context, *pb.RequestSearchForMovie) (*pb.ResponseSearchForMovie, error) {
	return nil, nil
}

func (s *movieServer) GetFavMovies(context.Context, *pb.RequestMovies) (*pb.ResponseMovies, error) {
	return nil, nil
}
