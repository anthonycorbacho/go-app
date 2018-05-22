package pkg

// Configuration represents a general app configuration.
type Configuration struct {
	Port  string `default:":8001"`
	Couch couch
}

// Couch represents couch db configuration.
type couch struct {
	Url      string `default:"http://127.0.0.1:5984"`
	User     string `default:""`
	Password string `default:""`
}
