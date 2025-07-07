package main

import (
	"fmt"
	"log"
	"os"

	"github.com/12153/hummingbird/cli/cmd/pkg"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: hummingbird <command> [args]")
		os.Exit(1)
	}

	switch args[1] {
	case "init":
		if len(args) < 3 {
			fmt.Println("Usage: hummingbird init <project-name>")
			os.Exit(1)
		}
		if err := pkg.Scaffold(args[2]); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("✅ Project '%s' initialized.\n", args[2])
	case "add":
		if len(args) < 4 {
			fmt.Println("Usage: hummingbird add route <name>")
			os.Exit(1)
		}
		if args[2] == "route" {
			if err := pkg.AddRoute(args[3], "."); err != nil {
				log.Fatal(err)
			}
		}
	case "dev":
		if len(args) < 3 {
			fmt.Println("Usage: hummingbird dev <project-dir>")
			os.Exit(1)
		}
		if err := pkg.Dev(args[2]); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("Unknown command:", args[1])
	}
}
