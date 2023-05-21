package main

import (
	"context"
	"fmt"
	pb "github.com/Abdulrahman-Awadh/Movirate/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
	"strings"
)

func main() {
	//select endpoint menu => done
	//Get Latest Movies => done
	//select one => done
	//show details => done
	//add to fav
	//Search for a movie
	//add to fav
	//get my favorites
	//select one
	//show details
	//delete from favorite
	for {
		// Print the menu
		fmt.Println("=== Welcome To Movirate App ===")
		fmt.Println("=== Select an option ===")
		fmt.Println("1. Get The Latest Movies")
		fmt.Println("2. Search for a movie")
		fmt.Println("3. Get your favorites movies")
		fmt.Println("0. Exit")

		// Read user input
		var choice string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		// Process the user's choice
		switch strings.TrimSpace(choice) {
		case "1":
			fmt.Println("Option 1 selected")
			movies := getLatestMovies()
			movie := selectMovie(movies)
			if movie == nil {
				continue
			}
			movieMenu(movie)
			// Add code for Option 1
		case "2":
			fmt.Println("Option 2 selected")
			// Add code for Option 2
		case "3":
			fmt.Println("Option 3 selected")
			// Add code for Option 3
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Println() // Add an empty line for better readability
	}

}
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

func printMovies(movies []*pb.Movie) {
	log.Printf("Movie list:")
	for index, value := range movies {
		fmt.Printf("%d- %s\n", index+1, value.Name)
	}

}

func printMovieDetails(movie *pb.Movie) {
	fmt.Println("Movie Name:", movie.Name)
	fmt.Println("Movie Description:", movie.Description)
	fmt.Println("Movie Genre:", movie.Genre)
	fmt.Println("Movie Poster:", movie.Poster)
	fmt.Println()
}

func selectMovie(movies []*pb.Movie) *pb.Movie {
	printMovies(movies)
	for {
		// Print the menu
		fmt.Println("=== Select a Movie ===")
		fmt.Println("B. Back")

		// Read user input
		var choice string
		fmt.Print("Enter The Number: ")
		fmt.Scanln(&choice)
		fmt.Println()

		// Process the user's choice
		num, err := strconv.Atoi(choice)
		if (err == nil) && (num <= len(movies)) {
			return movies[num-1]
		} else if choice == "b" {
			return nil
		} else {
			fmt.Println("Invalid choice. Please try again.")
		}

	}

	fmt.Println() // Add an empty line for better readability
	return nil
}
func movieMenu(movie *pb.Movie) {
	printMovieDetails(movie)
	for {
		// Print the menu
		fmt.Println("F. Add to favorites movies")
		fmt.Println("B. Back")

		// Read user input
		var choice string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		// Process the user's choice
		switch strings.TrimSpace(choice) {
		case "f":
			//add to fav

		case "b":
			return
			// Add code for Option 2
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Println() // Add an empty line for better readability
	}

}

func getFavoriteMovies(movie pb.Movie) {

}

func addMovieToFavorite(movie pb.Movie) {

}
func deleteMovieFromFavorite(movie pb.Movie) {

}
