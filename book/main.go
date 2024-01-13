package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Photojs struct {
	Photo []string `json:"myphoto"`
}

func main() {
	files, err := ioutil.ReadDir("pics/")
	if err != nil {
		log.Fatal(err)
	}
	photos := Photojs{}

	for _, file := range files {

		photos.Photo = append(photos.Photo, fmt.Sprintf("pics/%v", file.Name()))
	}
	out, _ := json.Marshal(photos.Photo)
	d1 := []byte(fmt.Sprintf("var myphoto=%v", string(out)))
	os.WriteFile("photo.js", d1, 0644)
	log.Println(string(d1))

}
