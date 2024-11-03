package main

import (
	"fmt"
	"github.com/xebia/go-exercise/internal/datastore"
	"github.com/xebia/go-exercise/internal/email"
	"github.com/xebia/go-exercise/internal/server"
)

func main() {
	ds := datastore.NewService()
	es := email.NewService()

	srv := server.NewServer(ds, es)
	addr := fmt.Sprintf("%s:%s", "127.0.0.1", "8080")
	srv.ListenAndServe(addr)
}
