package main

import "awesomeKonstru/backend/models"

func main() {
	models.ConnectDB()
	models.MigrateDB()
}
