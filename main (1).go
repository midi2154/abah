package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("error")
		return
	}

	bannerfile := "standard.txt"
	if len(os.Args) == 3 {
		bannerfile = os.Args[2] + ".txt"
	}

	banner, err := LoadBanner(bannerfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(Render(os.Args[1], banner))
}
