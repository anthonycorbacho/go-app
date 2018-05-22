package movie

// MovieID represents a movie identifier.
type MovieID string

// Movie represents the top level of a movie.
type Movie struct {
	ID          MovieID `json:"id"`
	Description string  `json:"description"`
	Year        int64   `json:"year"`
}

// Store represents a store for managing movies.
type Store interface {
	GetAll() []Movie
	Get(MovieID) (*Movie, error)
	Create(*Movie) error
	Update(*Movie) error
	Delete(MovieID) error
}
