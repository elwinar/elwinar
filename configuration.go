package main

// Configuration holds the configuration for the website: which port to listen on,
// which database to use, etc.
type Configuration struct {
	Database string
	Debug    bool
	Password string
	Port     int
	Public   string
	Secret   string
}
