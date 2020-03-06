package repositories

type DummyMovie struct {
}

func NewDummyMovie() *DummyMovie {
	return &DummyMovie{}
}

func (rep *DummyMovie) GetAll() (string, error) {
	return `[
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
	{
		"id": 2,
		"vote_count": 2059,
		"vote_average": 8.1,
		"title": "1917",
		"original_title": "1917",
		"original_language": "en",
		"adult": 0,
		"poster_path": "/iZf0KyrE25z1sage4SYFLCCrMi9.jpg",
		"overview": "Nos encontramos en el año 1917. La Guerra Mundial amenazaba con cambiar, para siempre, el orden mundial. Ante la amenaza que se cernía, Estados Unidos decidió entrar en el conflicto con el objetivo de desequilibrar la balanza que caracterizaba a la contienda.",
		"release_date": "2019-12-10",
		"popularity": 99.999
	},
	{
		"id": 3,
		"vote_count": 85,
		"vote_average": 7.4,
		"title": "Aves de Presa (y la Fantabulosa Emancipación de Harley Quinn)",
		"original_title": "Birds of Prey (and the Fantabulous Emancipation of One Harley Quinn)",
		"original_language": "en",
		"adult": 0,
		"poster_path": "/sai5HgiBJOwqFxjECJyHONapzLk.jpg",
		"overview": "Después de separarse de Joker, Harley Quinn y otras tres heroínas (Canario Negro, Cazadora y Renée Montoya) unen sus fuerzas para salvar a una niña (Cassandra Cain) del malvado rey del crimen Máscara Negra.",
		"release_date": "2020-02-05",
		"popularity": 99.999
	}
]`, nil
}
