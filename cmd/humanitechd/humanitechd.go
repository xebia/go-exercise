package main

import (
	"fmt"
	"os"

	"github.com/xebia/go-exercise/internal/db"
	"github.com/xebia/go-exercise/internal/email"
	"github.com/xebia/go-exercise/internal/server"
)

func main() {
	emailAPIKey := os.Getenv("SENDGRID_API_KEY")

	ds := db.NewService()
	es := email.NewService(emailAPIKey)

	srv := server.NewServer(ds, es)

	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	srv.ListenAndServe(addr)
}
