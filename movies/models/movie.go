package models

type Movie struct {
	ID               int     `json:"id"`
	Title            string  `json:"title"`
	VoteCount        int     `json:"vote_count"`
	VoteAverage      float32 `json:"vote_average"`
	OriginalTitle    string  `json:"original_title"`
	OriginalLanguage string  `json:"original_languaje"`
	Adult            bool    `json:"adult"`
	PosterPath       string  `json:"poster_path"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	Popularity       float32 `json:"popularity"`
}
