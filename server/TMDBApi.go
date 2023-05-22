package main

import (
	"encoding/json"
	pb "github.com/Abdulrahman-Awadh/Movirate/proto"
	"io/ioutil"
	"log"
	"net/http"
)

// structure for json :)
type MoviesList struct {
	Dates struct {
		Maximum string `json:"maximum"`
		Minimum string `json:"minimum"`
	} `json:"dates"`
	Page    int `json:"page"`
	Results []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIds         []int   `json:"genre_ids"`
		ID               int     `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		Popularity       float64 `json:"popularity"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		VoteAverage      float64 `json:"vote_average"`
		VoteCount        int     `json:"vote_count"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

func getMoviesFromAPI() []*pb.Movie {
	url := "https://api.themoviedb.org/3/movie/now_playing?language=en-US&page=1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI3Njg0NTRmMzBmNDBkYjJkNzFjZGNmOGFhYmJmODYyYiIsInN1YiI6IjY0Njc1ZTllMDA2YjAxMDE0N2U5NDNkYiIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.l2lKzWqj0OJ1yAXHAkw0LoE5IQJvMneYBJOPuDvzSds")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))

	var decodedPerson MoviesList
	err := json.Unmarshal(body, &decodedPerson)
	if err != nil {
		log.Fatal(err)
	}

	var movies []*pb.Movie
	//fmt.Println("Decoded Person:", decodedPerson.Results[1].Title)
	for _, value := range decodedPerson.Results {
		movies = append(movies, &pb.Movie{
			Id:          int32(value.ID),
			Name:        value.Title,
			Poster:      value.PosterPath,
			Description: value.Overview,
			Genre:       "value.GenreIds",
		})
	}
	return movies
}

func searchForMovie() {

}
