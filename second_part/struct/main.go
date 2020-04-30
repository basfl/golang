package main

import "fmt"

type songs struct {
	title, artist string
}
type playlist struct {
	genra string
	songs []songs
}

func main() {
	songs := []songs{
		{
			title:  "walking on moon",
			artist: "pink floyed",
		}, {
			title:  "rock on",
			artist: "bryan",
		},
	}
	rocks := playlist{genra: "rock", songs: songs}
	fmt.Printf("\n%v", rocks)
	for _, v := range rocks.songs {
		fmt.Printf("\n title=%s and artist is %s ", v.title, v.artist)
	}
	rocks.songs[0].title = "mine"
	fmt.Printf("\n%v", rocks)

}
