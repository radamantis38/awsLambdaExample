package repositories

import (
	"awsLambdaExample/movies/models"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlMovie struct {
}

func NewMysqlMovie() *MysqlMovie {
	return &MysqlMovie{}
}

func (rep *MysqlMovie) GetAll() (string, error) {
	fmt.Println("Go MySQL")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:strongpassword@tcp(127.0.0.1:3306)/movies")

	// if there is an error opening the connection, handle it
	if err != nil {
		return "", errors.New("Error de repositorio")
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT *  FROM movies limit 10")
	if err != nil {
		return "", errors.New("Error de repositorio")
	}

	var movies []models.Movie
	for results.Next() {
		var movie models.Movie
		// for each row, scan the result into our tag composite object
		err = results.Scan(&movie.ID, &movie.VoteCount, &movie.VoteAverage, &movie.Title, &movie.OriginalTitle, &movie.OriginalLanguage, &movie.Adult, &movie.PosterPath, &movie.Overview, &movie.ReleaseDate, &movie.Popularity)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		movies = append(movies, movie)
	}

	moviesJson, err := json.Marshal(movies)
	if err != nil {
		return "", errors.New("Error de repositorio")
	}

	return string(moviesJson), nil
}
