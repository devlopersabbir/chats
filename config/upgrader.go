// Package config is a package for upgrader configuration
package config

import (
	"log"
	"net/http"
	"slices"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		var origin = r.Header.Get("Origin")

		if slices.Contains(SocketConfig.Origin, origin) {
			log.Println("Origin allowed", origin)
			return true
		} else {
			log.Fatal("Origin not allowed")
			return false
		}
	},
}
