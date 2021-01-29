package main

import (
	"errors"
	"fmt"
)

type MovieRepositoryInterface interface {
	GetAll() (string, error)
}

type getMovies struct {
	movieRepository MovieRepositoryInterface
}

func NewGetMovies(movieRepository MovieRepositoryInterface) *getMovies {
	return &getMovies{movieRepository: movieRepository}
}

func (rep *getMovies) Handle() (string, error) {
	fmt.Println("entro al caso de uso")

	movies, err := rep.movieRepository.GetAll()
	if err != nil {
		return "", errors.New("Error repositorio")
	}
	return movies, nil

}
