package main

import (
	"context"
	"fmt"
	"github.com/EugeneTsydenov/chesshub-{{cookiecutter.domain}}-service/cmd/{{cookiecutter.domain}}/app"
	"github.com/EugeneTsydenov/chesshub-{{cookiecutter.domain}}-service/config"
	"os"
)

func main() {
	env := os.Getenv("APP_ENV")
	cfgPath := os.Getenv("CONFIG_PATH")
	cfg, err := config.Load(env, cfgPath)

	if err != nil {
		fmt.Printf("Failed to laoding config: %v\n", err)
	}

	ctx := context.Background()

	a := app.New(cfg)
	if err = a.InitDeps(ctx); err != nil {
		fmt.Printf("Failed to initialize dependencies: %v\n", err)
		os.Exit(1)
	}

	a.SetupGRPCServer()

	if err = a.Run(ctx); err != nil {
		fmt.Printf("Application error: %v\n", err)
		os.Exit(1)
	}
}
