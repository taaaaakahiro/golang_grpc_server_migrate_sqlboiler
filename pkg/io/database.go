package io

import (
	"database/sql"
	"fmt"
	"golang_grpc_proto/pkg/config"
	"log"

	_ "github.com/lib/pq"
)

func NewDataBase(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"user=%s dbname=%s password=%s  host=%s port=5432 sslmode=disable",
		cfg.DB.User, cfg.DB.Database, cfg.DB.Password, cfg.DB.Host)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
