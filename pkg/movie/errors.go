package movie

// Movie errors.
const (
	ErrMovieNotFound   = Error("movie not found")
	ErrMovieExists     = Error("movie already exists")
	ErrMovieIDRequired = Error("movie id required")
)

// Error represents a Movie error.
type Error string

// Error returns the error message.
func (e Error) Error() string {
	return string(e)
}