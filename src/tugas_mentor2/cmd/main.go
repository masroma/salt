package main

import (
	"fmt"
)

type author struct {
	name       string
	created_at string
	updated_at string
}
type listFood struct {
	id      int
	title   string
	content string
	author
	created_at string
	updated_at string
}

func main() {
	var AlllistFood = []listFood{
		{
			id:      1,
			title:   "makan ayam",
			content: "ini adalah content makan ayam",
			author: author{
				name:       "roma",
				created_at: "2020-08-23 11:51:07",
				updated_at: "2020-08-28 20:50:12",
			},
			created_at: "2020-08-23 11:51:07",
			updated_at: "2020-08-28 20:50:12",
		},
		{
			id:      2,
			title:   "makan ikan",
			content: "ini adalah content makan ikan",
			author: author{
				name:       "roma",
				created_at: "2020-08-23 11:51:07",
				updated_at: "2020-08-28 20:50:12",
			},
			created_at: "2020-08-23 11:51:07",
			updated_at: "2020-08-28 20:50:12",
		},
	}

	var kondisi string

	for _, list := range AlllistFood {

		if list.id%2 == 0 {
			kondisi = "genap"
		} else {
			kondisi = "ganjil"
		}

		fmt.Println("The title of article is", list.title, "", kondisi)
		fmt.Println("The author of article is", list.author.name)
		fmt.Println("The article was created at", list.created_at)
		fmt.Println("==================================")
	}

}
