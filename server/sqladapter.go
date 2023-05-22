package main

import (
	"database/sql"
	"fmt"
	pb "github.com/Abdulrahman-Awadh/Movirate/proto"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

const (
	host     = "localhost"
	port     = 8001
	user     = "postgres"
	password = "example"
	dbname   = "movirate"
)

func databaseInit() {
	CreateDatabase()
	CreateTable()
}
func CreateDatabase() {

	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s sslmode=disable",
		host, port, user, password)
	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE DATABASE movirate")
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		fmt.Println("Database creation error:", err)
		return
	}
	//defer db.Close()

}

func CreateTable() {
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Database created successfully.")
	createTable := `
		CREATE TABLE IF NOT EXISTS movie (
		    id bigint PRIMARY KEY,
		    name text,
			description text,
			poster_path text
		);
	`

	_, err = db.Exec(createTable)
	if err != nil {
		fmt.Println("Table creation error:", err)
		return
	}
	fmt.Println("Table created successfully.")
	insertQuery := `
		INSERT INTO movie (id, name, description, poster_path)
		VALUES (123,'test','test','test');
	`
	_, err = db.Exec(insertQuery)
	if err != nil {
		fmt.Println("Table creation error:", err)
		return
	}

}

func AddToFav(movie pb.Movie) {
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Define the insert query with placeholders
	insertQuery := `
		INSERT INTO movie (id, name, description, poster_path)
		VALUES ($1,$2,$3,$4);
	`

	// Prepare the statement
	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		fmt.Println("Statement preparation error:", err)
		return
	}
	defer stmt.Close()

	// Execute the prepared statement with variables

	err = stmt.QueryRow(movie.Id, movie.Name, movie.Description, movie.Poster).Scan()
	if err != nil {
		fmt.Println("Record insertion error:", err)
		return
	}

	fmt.Println("Record inserted successfully.")

}

func DeleteFromFav(id int) {
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Define the insert query with placeholders
	insertQuery := `
		DELETE FROM movie
		WHERE id = $1;
`

	// Prepare the statement
	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		fmt.Println("Statement preparation error:", err)
		return
	}
	defer stmt.Close()

	// Execute the prepared statement with variables

	err = stmt.QueryRow(id).Scan()
	if err != nil {
		fmt.Println("Record insertion error:", err)
		return
	}

	fmt.Println("Record inserted successfully.")

}

func GetFavMovies() []pb.Movie {
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		panic(err)
	}
	rows, err := db.Query(`SELECT * FROM movie;`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var movies []pb.Movie
	for rows.Next() {
		var id int
		var name string
		var description string
		var poster_path string

		err = rows.Scan(&id, &name, &description, &poster_path)
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, pb.Movie{
			Id:          int32(id),
			Name:        name,
			Poster:      poster_path,
			Description: description,
		})
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return movies
}
