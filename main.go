package main

import (
	"github.com/cardlet/api"
)

func main() {
	api.SetupConfig()

	api.InitDb()

	api.CreateRouter()
}