package couch

import (
	"github.com/anthonycorbacho/go-app/pkg/movie"
	"github.com/leesper/couchdb-golang"
)

// Ensure MovieService implement movie.service.
var _ movie.Service = &MovieService{}

const (
	movieDB   = "/featureflags"
	movieType = "movie_type"
	all       = `
      {
        "selector": {
          "type": "movieType"
        }
      }`
)

type document struct {
	Description string `json:"description"`
	Year        int64  `json:"description"`
	Type        string `json:"type"`
	couchdb.Document
}

// MovieService represents a service managing movies in couchDB.
type MovieService struct {
	couch *couchdb.Database
}

// New creates a instance of movie service
func New(u string) *MovieService {
	database, err := couchdb.NewDatabase(u + movieDB)
	if err != nil {
		// handle error
	}

	return &MovieService{
		couch: database,
	}
}

// GetAll returns all movies.
func (ms *MovieService) GetAll() []movie.Movie {
	docsQuery, err := ms.couch.QueryJSON(all)
	if err != nil {
		return []movie.Movie{}
	}

	size := len(docsQuery)
	movies := make([]movie.Movie, size)
	for i := 0; i < size; i++ {
		d := docsQuery[i]

		var movieDoc document
		err := couchdb.FromJSONCompatibleMap(&movieDoc, d)
		if err != nil {
			movies[i] = movie.Movie{}
			continue
		}

		var movie movie.Movie
		toMovie(&movie, &movieDoc)
		movies[i] = movie
	}
	return movies
}

// Get return a movie by its ID.
func (ms *MovieService) Get(ID movie.MovieID) (*movie.Movie, error) {
	doc, err := ms.couch.Get(string(ID), nil)
	if err != nil {
		return nil, err
	}

	var d document
	err = couchdb.FromJSONCompatibleMap(&d, doc)
	if err != nil {
		return nil, err
	}

	var m movie.Movie
	toMovie(&m, &d)

	return &m, nil
}

// Create creates a new Movie.
// This will return movie structure with new ID.
func (ms *MovieService) Create(m *movie.Movie) error {
	if m == nil {
		return movie.ErrMovieNotFound
	}

	if m.ID == "" {
		return movie.ErrMovieIDRequired
	}

	var d document
	fromMovie(&d, m)

	doc, err := couchdb.ToJSONCompatibleMap(d)
	if err != nil {
		return err
	}

	_, _, err = ms.couch.Save(doc, nil)
	if err != nil {
		return err
	}
	return nil
}

// Update updates a movie.
func (ms *MovieService) Update(m *movie.Movie) error {
	if m.ID == "" {
		return movie.ErrMovieIDRequired
	}
	var d document
	fromMovie(&d, m)
	d.Rev = ms.getRevision(d.ID)

	doc, err := couchdb.ToJSONCompatibleMap(d)
	if err != nil {
		return err
	}

	_, _, err = ms.couch.Save(doc, nil)
	if err != nil {
		return err
	}
	return nil
}

func (ms *MovieService) getRevision(ID string) string {
	doc, err := ms.couch.Get(ID, nil)
	if err != nil {
		return ""
	}

	var d document
	err = couchdb.FromJSONCompatibleMap(&d, doc)
	if err != nil {
		return ""
	}
	return d.Rev
}

// Delete deletes a movie by its ID.
func (ms *MovieService) Delete(ID movie.MovieID) error {

	if ID == "" {
		return movie.ErrMovieIDRequired
	}

	err := ms.couch.Delete(string(ID))
	if err != nil {
		return err
	}
	return nil
}

// toFlag converts a document to a movie.
func toMovie(m *movie.Movie, d *document) {
	m.ID = movie.MovieID(d.ID)
	m.Description = d.Description
	m.Year = d.Year
}

// fromMovie converts a movie to a document.
func fromMovie(d *document, m *movie.Movie) {
	d.ID = string(m.ID)
	d.Description = m.Description
	d.Year = m.Year
	d.Type = movieType
}
