package main

import (
	"fmt"
	"os"
	"os/exec"
)

var path = "app/database/migrations"
var database = "postgres://postgres:@localhost:5432/eco_db?sslmode=disable&search_path=public"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/migrate.go [migrate:up|migrate:down|migrate:create]")
		return
	}

	command := os.Args[1]

	switch command {
	case "migrate:up":
		runMigrate("up")
	case "migrate:down":
		runMigrate("down")
	case "migrate:create":
		if len(os.Args) < 3 {
			fmt.Println("Set migration name file -> ")
			return
		}
		name := os.Args[2]
		runCreate(name)
	default:
		fmt.Println("It is not migrate Comand line :", command)
	}
}

func runMigrate(direction string) {
	args := []string{
		"-path", path,
		"-database", database,
	}

	args = append(args, direction)

	if len(os.Args) > 2 {
		args = append(args, os.Args[2:]...)
	}

	cmd := exec.Command("migrate", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	cmd.Run()
}

func runCreate(name string) {
	cmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", path, "-seq", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
