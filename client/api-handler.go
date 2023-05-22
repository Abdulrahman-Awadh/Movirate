package main

import (
	"context"
	"fmt"
	pb "github.com/Abdulrahman-Awadh/Movirate/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func getLatestMovies() []*pb.Movie {
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
	return movies.Movie
}

func getFavMovies() []*pb.Movie {
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error", err)
	}
	client := pb.NewMovieApiClient(conn)
	movies, err := client.GetFavMovies(context.Background(), &pb.RequestMovies{})
	if err != nil {
		log.Fatalf("Error", err)
	}
	return movies.Movie
}

func addMovieToFav(movie *pb.Movie) {
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error in connection: ", err)
	}

	client := pb.NewMovieApiClient(conn)

	response, err := client.AddMovieToFav(context.Background(), &pb.RequestAddMovieToFav{
		Movie: movie,
	})

	if err != nil {
		log.Fatal("Error in adding the movie: ", err)
	}
	fmt.Println(response.Message)
	return
}

func deleteMovieFromFav(movie *pb.Movie) {
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error", err)
	}
	client := pb.NewMovieApiClient(conn)
	response, err := client.DeleteMovieFromFav(context.Background(), &pb.RequestDeleteMovieFromFav{
		Id: movie.Id,
	})
	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Println(response.Message)
	return
}
func searchForMovie(keyword string) []*pb.Movie {
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error", err)
	}

	client := pb.NewMovieApiClient(conn)

	response, err := client.SearchForMovie(context.Background(), &pb.RequestSearchForMovie{
		Keyword: keyword,
	})

	if err != nil {
		log.Fatal("Error: ", err)
	}
	return response.Movie
}
