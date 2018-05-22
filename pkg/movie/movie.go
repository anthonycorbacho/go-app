package movie

// MovieID represents a movie identifier.
type MovieID string

// Movie represents the top level of a movie.
type Movie struct {
	ID          MovieID `json:"id"`
	Description string  `json:"description"`
	Year        int64   `json:"year"`
}

// Service represents a service for managing movies.
type Service interface {
	GetAll() []Movie
	Get(MovieID) (*Movie, error)
	Create(*Movie) error
	Update(*Movie) error
	Delete(MovieID) error
}
