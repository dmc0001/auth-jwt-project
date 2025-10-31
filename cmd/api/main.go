package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dmc0001/auth-jwt-project/internal/database"
	"github.com/dmc0001/auth-jwt-project/internal/env"
	"github.com/dmc0001/auth-jwt-project/internal/store"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		env.GetString("DB_USER", "admin"),
		env.GetString("DB_PASS", "password"),
		env.GetString("DB_HOST", "localhost"),
		env.GetString("DB_PORT", "3306"),
		env.GetString("DB_NAME", "name_db"),
	)

	port := env.GetString("PORT", ":80")

	db, err := database.InitDb(dsn)
	if err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}
	defer db.Close()

	cfg := &Config{
		addr:      port,
		userModel: store.NewUserModel(db),
	}

	app := &Application{
		config: cfg,
	}

	mux := http.NewServeMux()
	srv := app.Route(mux)

	log.Printf("✅ Server running on port %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("❌ Server stopped:", err)
	}
}
