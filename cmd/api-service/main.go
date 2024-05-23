package main

import (
	"cityger-be/internal/config"
	"cityger-be/internal/logger"
	"cityger-be/internal/storage/postgresql"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustRead(os.Args)

	//! FIXME: To be removed
	if cfg.Env == "dev" {
		fmt.Println("Config:", *cfg)
		fmt.Println("Server:", (*cfg).HTTPServer.Address)
	}

	log := logger.Init(cfg.Env)
	//! FIXME: To be removed
	log.Info("logger initialized successfully", slog.String("env", cfg.Env))

	storage, err := postgresql.New(cfg.Storage)
	if err != nil {
		log.Error("storage init failed", slog.Attr{Key: "error", Value: slog.StringValue(err.Error())})
		os.Exit(1)
	}
	defer storage.Close()

	//! FIXME: To be removed
	log.Info("storage initialized successfully", slog.String("env", cfg.Env))

	//! Check DB data - ROLE, to be removed
	roles, err := storage.GetAllRoles()
	if err != nil {
		log.Error("roles data receive failed", slog.Attr{Key: "error", Value: slog.StringValue(err.Error())})
	} else {
		fmt.Println("Roles data:\n", roles)
	}
	id, err := storage.AddNewRole("new_role2", "Nova Role to play")
	if err != nil {
		log.Error("role add failed", slog.Attr{Key: "error", Value: slog.StringValue(err.Error())})
	} else {
		fmt.Println("New ID:", id)
		roles, err := storage.GetAllRoles()
		if err != nil {
			log.Error("roles data receive failed", slog.Attr{Key: "error", Value: slog.StringValue(err.Error())})
		} else {
			fmt.Println("Roles data:\n", roles)
		}
	}

	// TODO: Init Router

	// TODO: Run Service

	// TODO: Auth FE-BE (???)
}
