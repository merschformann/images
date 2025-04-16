package main

type Item struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

var DB = []Item{
	{Message: "Hello, World!"},
}
