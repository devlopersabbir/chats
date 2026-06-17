// Package config is a package for origin configuration
package config

type configStruct struct {
	Origin []string
}

var SocketConfig = configStruct{
	Origin: []string{
		"http://localhost:8080",
		"https://localhost:8080",
		"http://127.0.0.1:8080",
		"http://127.0.0.1:5500",
	},
}
