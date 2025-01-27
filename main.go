package main

import "log"

func main() {
	server := InitializedServer()

	log.Printf("Server Listening on %s", server.Addr)
	err := server.ListenAndServe()
	log.Fatal(err)
}
