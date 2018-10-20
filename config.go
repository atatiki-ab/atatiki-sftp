package main

// Config file for the service
type Config struct {
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
}
