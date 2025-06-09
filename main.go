/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"todo/cmd"
	"todo/database"
)

func main() {
	database.InitDB()
	defer database.DB.Close()
	cmd.Execute()
}
