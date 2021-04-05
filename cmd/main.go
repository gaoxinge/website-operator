package main

import (
	"log"

	"github.com/gaoxinge/website-operator/pkg"
)

func main() {
	websiteControl, err := pkg.NewWebsiteController("http://127.0.0.1:8001")
	if err != nil {
		log.Printf("run with error %v\n", err)
		return
	}

	websiteControl.Run()
}
