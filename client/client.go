package main

import (
	"fmt"
	pb "github.com/Abdulrahman-Awadh/Movirate/proto"
	"strconv"
	"strings"
)

func main() {

	//Search for a movie
	//add to fav
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
			movies := getLatestMovies()
			if movies == nil {
				continue
			}
			movie := selectMovie(movies)
			if movie == nil {
				continue
			}
			if movie == nil {
				continue
			}
			movieMenu(movie)
		case "2":
			var keyword string
			fmt.Print("Enter Movie Name: ")
			fmt.Scanln(&keyword)

			movies := searchForMovie(keyword)
			if movies == nil {
				fmt.Print("No Movie Found ")
				continue
			}
			movie := selectMovie(movies)
			if movie == nil {
				continue
			}
			if movie == nil {
				continue
			}
			movieMenu(movie)
		case "3":
			movies := getFavMovies()
			if movies == nil {
				continue
			}
			movie := selectMovie(movies)
			if movie == nil {
				continue
			}
			if movie == nil {
				continue
			}
			favMovieMenu(movie)
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Println()
	}

}

func printMovies(movies []*pb.Movie) {
	fmt.Println("Movie list:")
	for index, value := range movies {
		fmt.Printf("%d- %s\n", index+1, value.Name)
	}

}

func printMovieDetails(movie *pb.Movie) {
	fmt.Println("Movie Name:", movie.Name)
	fmt.Println("Movie Description:", movie.Description)
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
		fmt.Print("Enter your choice: ")
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

	fmt.Println()
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
			addMovieToFav(movie)
			return

		case "b":
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Println()
	}

}

func favMovieMenu(movie *pb.Movie) {
	printMovieDetails(movie)
	for {
		// Print the menu
		fmt.Println("D. Delete from favorites movies")
		fmt.Println("B. Back")

		// Read user input
		var choice string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		// Process the user's choice
		switch strings.TrimSpace(choice) {
		case "d":
			//add to fav
			deleteMovieFromFav(movie)
			return
		case "b":
			return
			// Add code for Option 2
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Println() // Add an empty line for better readability
	}

}
