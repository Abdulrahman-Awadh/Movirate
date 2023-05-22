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

func (s *movieServer) AddMovieToFav(ctx context.Context, req *pb.RequestAddMovieToFav) (*pb.ResponseAddMovieToFav, error) {
	movie := req.Movie
	if movie != nil {
		AddToFav(movie)
		return &pb.ResponseAddMovieToFav{
			IsDone:  true,
			Message: "Done",
		}, nil
	}

	return nil, nil
}

func (s *movieServer) DeleteMovieFromFav(ctx context.Context, req *pb.RequestDeleteMovieFromFav) (*pb.ResponseDeleteMovieFromFav, error) {
	id := req.Id
	DeleteFromFav(int(id))

	return &pb.ResponseDeleteMovieFromFav{
		IsDone:  true,
		Message: "Done",
	}, nil
}

func (s *movieServer) SearchForMovie(context.Context, *pb.RequestSearchForMovie) (*pb.ResponseSearchForMovie, error) {
	return nil, nil
}

func (s *movieServer) GetFavMovies(context.Context, *pb.RequestMovies) (*pb.ResponseMovies, error) {
	movies := GetFavMovies()

	return &pb.ResponseMovies{
		Movie: movies,
	}, nil
}
