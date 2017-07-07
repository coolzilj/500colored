package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
)

var colors = map[string]string{
	"blue":      "http://500colored.com/blue.html",
	"blue2":     "http://500colored.com/blue2.html",
	"purple":    "http://500colored.com/purple.html", // nil
	"pink":      "http://500colored.com/pink.html",
	"red":       "http://500colored.com/red.html",
	"orange":    "http://500colored.com/orange.html",
	"yellow":    "http://500colored.com/yellow.html",
	"yellow2":   "http://500colored.com/yellow2.html",
	"green":     "http://500colored.com/green.html",
	"green2":    "http://500colored.com/green2.html",
	"darkGreen": "http://500colored.com/dark_green.html",
	"brown":     "http://500colored.com/brown.html",
	"brown2":    "http://500colored.com/brown2.html",
	"mono":      "http://500colored.com/mono.html", // nil
}

var assetsFolder = "500colored-sync"

func downloadFile(filepath string, url string) {

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		// Create the file
		out, err := os.Create(filepath)
		if err != nil {
			log.Println(err)
		}
		defer out.Close()

		// Get the data
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()

		// Writer the body to file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			log.Println(err)
		}
	}

	if _, err := os.Stat(filepath); os.IsExist(err) {
		fmt.Printf("%s existed.", filepath)
	}

}

func main() {
	for color, URL := range colors {
		fmt.Printf("====\nColor: %s\n====\n", color)

		if _, err := os.Stat(assetsFolder + "/" + color); os.IsNotExist(err) {
			fmt.Printf("Creating folder %s.\n", color)
			os.MkdirAll(assetsFolder+"/"+color, 0777)
		}

		doc, err := goquery.NewDocument(URL)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".section").Each(func(i int, s *goquery.Selection) {
			flickrURL, _ := s.Find("a").Attr("href")
			fileName := s.Find(".title h3").Text()
			filepath := filepath.Join(assetsFolder, color, fileName+".jpg")

			fmt.Printf("Downloading Image %d: %s - %s\n", i+1, fileName, flickrURL)
			if flickrURL != "" {
				downloadFile(filepath, flickrURL)
			}
		})
	}
}
