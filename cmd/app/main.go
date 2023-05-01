package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"

	"github.com/yudgxe/jetstyle-rest-api/internal/app"
	"github.com/yudgxe/jetstyle-rest-api/pkg/database"
	"github.com/yudgxe/jetstyle-rest-api/pkg/flagexp"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

// @title Jetstyle API
// @version 1.0
// @description API Server for Jetstyle test tasks

// @host localhost:8080
// @BasePath /

// @securitydefinitions.basic BasicAuth

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must pass at least one argument")
		os.Exit(0)
	}

	flag.Parse()

	config := app.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalln(err)
	}

	flags := make([]*flag.FlagSet, 0, 2)

	serveFlag := flag.NewFlagSet("serve", flag.ExitOnError)
	serveFlag.StringVar(&config.BindAddr, "bind-addr", config.BindAddr, "server address")

	flagexp.Subset(serveFlag, "db", ":", func(sub *flag.FlagSet) {
		config.DB = database.Export(sub, *config.DB)
	})

	var login, password string
	createFlag := flag.NewFlagSet("create", flag.ExitOnError)

	createFlag.StringVar(&login, "l", "", "user login")
	createFlag.StringVar(&password, "p", "", "user password")

	flagexp.Subset(createFlag, "db", ":", func(sub *flag.FlagSet) {
		config.DB = database.Export(sub, *config.DB)
	})

	flags = append(flags, serveFlag)
	flags = append(flags, createFlag)

	subcommand := strings.ToLower(os.Args[1])
	for _, v := range flags {
		if v.Name() == subcommand {
			if err := v.Parse(os.Args[2:]); err != nil {
				log.Fatalln(err)
			}
		}
	}

	switch subcommand {
	case "serve":
		app.Start(config)
	case "create":
		if err := app.CreateUser(config, login, password); err != nil {
			log.Println(err)
		}
	}
}
