package main

import (
	"github.com/jmorave/budgeter/pkg/api"
	"github.com/jmorave/databse"
)

func main() {
	database_create()
	handleRequests()
}