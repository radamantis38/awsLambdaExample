package main

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"
)

//mock repository
type MovieRepositoryMock struct {
	mock.Mock
}

func (m *MovieRepositoryMock) GetAll() (string, error) {
	args := m.Called()
	return args.Get(0).(string), args.Error(1)
}

// fin mock

func Test_getMovies_Handle(t *testing.T) {
	type fields struct {
		movieRepository *MovieRepositoryMock
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
		mocks   func(f fields)
	}{
		{
			name:   "Exitoso",
			fields: fields{movieRepository: &MovieRepositoryMock{}},
			want: `[
	{
		"id": 1,
		"vote_count": 2219,
		"vote_average": 6,
		"title": "Ad astra",
		"original_title": "Ad Astra",
		"original_language": "en",
		"adult": 0,
		"poster_path": "/iR8TvOxTQNiTXFHs1S43RE2j0kg.jpg",
		"overview": "El futuro cercano, un tiempo en el que tanto la esperanza como la adversidad impulsan a la humanidad a mirar hacia las estrellas y más allá. Mientras un misterioso fenómeno amenaza con destruir la vida en el planeta Tierra, el astronauta Roy McBride emprende una misión a través de la inmensidad del espacio y sus muchos peligros para descubrir la verdad sobre una expedición perdida que décadas antes afrontó audazmente el vacío y el silencio en busca de lo desconocido.",
		"release_date": "2019-09-17",
		"popularity": 99.999
	},
]`,
			wantErr: false,
			mocks: func(f fields) {
				f.movieRepository.On("GetAll").Return(
					`[
	{
		"id": 1,
		"vote_count": 2219,
		"vote_average": 6,
		"title": "Ad astra",
		"original_title": "Ad Astra",
		"original_language": "en",
		"adult": 0,
		"poster_path": "/iR8TvOxTQNiTXFHs1S43RE2j0kg.jpg",
		"overview": "El futuro cercano, un tiempo en el que tanto la esperanza como la adversidad impulsan a la humanidad a mirar hacia las estrellas y más allá. Mientras un misterioso fenómeno amenaza con destruir la vida en el planeta Tierra, el astronauta Roy McBride emprende una misión a través de la inmensidad del espacio y sus muchos peligros para descubrir la verdad sobre una expedición perdida que décadas antes afrontó audazmente el vacío y el silencio en busca de lo desconocido.",
		"release_date": "2019-09-17",
		"popularity": 99.999
	},
]`, nil).Once()
			},
		},
		{
			name:    "Fail",
			fields:  fields{movieRepository: &MovieRepositoryMock{}},
			want:    "",
			wantErr: true,
			mocks: func(f fields) {
				f.movieRepository.On("GetAll").Return("", errors.New("Error")).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks(tt.fields)
			rep := NewGetMovies(tt.fields.movieRepository)
			got, err := rep.Handle()
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
			tt.fields.movieRepository.AssertExpectations(t)
		})
	}
}
