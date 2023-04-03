package types

const (
	BreakerSearchMovieByKeyword = 0
	BreakerGetMovieDetails      = 1

	OmdbAPIURL = "http://www.omdbapi.com"

	APIKey = "faf7e5bb"
)

var (
	ExtAPIBreakerList = []int{
		BreakerSearchMovieByKeyword,
		BreakerGetMovieDetails,
	}
)
