package main

import (
	"context"

	"github.com/carlmjohnson/requests"
)

func main() {
	ctx := context.Background()
	url := "https://discord.com/api/webhooks/{guildid}/{token}"

	// get(ctx, url)

	type embed struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Color       int    `json:"color"`
	}

	type payload struct {
		Username string  `json:"username"`
		Content  string  `json:"content"`
		Embeds   []embed `json:"embeds"`
	}

	var res string
	req := payload{
		// Content:  "Hello, world!",
		Username: "test",
		Embeds: []embed{
			{
				Title:       "Title",
				Description: "Description",
				Color:       1127128,
			},
		},
	}

	err := requests.
		URL(url).
		BodyJSON(&req).
		ToString(&res).
		Fetch(ctx)

	if err != nil {
		panic(err)
	}

}

func get(ctx context.Context, url string) {
	var s string
	err := requests.
		URL(url).
		ToString(&s).
		Fetch(ctx)

	if err != nil {
		panic(err)
	}

	println(s)
}
