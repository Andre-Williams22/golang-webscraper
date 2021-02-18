package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/

type Fact struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func main() {
	// fmt.Println("Start")
	// // Instantiate default collector
	var counter int

	german_facts := make([]Fact, 0)

	collector := colly.NewCollector()
	// ("https://www.expatica.com/de/moving/about/germany-facts-108768/"),

	// query selector on site
	collector.OnHTML(".entry__content li strong", func(element *colly.HTMLElement) {

		// Increments ID by 1
		counter++

		// Makes ID into a String
		// fmt.Println(element.Attr("id"))
		// factId, err := strconv.Atoi(element.Attr("id"))

		// if err != nil {
		// 	log.Println("Failed to get ID")
		// }

		factDesc := element.Text

		fact := Fact{
			ID:          counter,
			Description: factDesc,
		}

		german_facts = append(german_facts, fact)

		// fmt.Println(element.Text)

	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())

	})

	collector.Visit("https://www.expatica.com/de/moving/about/germany-facts-108768/")
	writeJSON(german_facts)

	// enc := json.NewEncoder(os.Stdout)
	// enc.SetIndent("", " ")
	// enc.Encode(german_facts)
	// writeJSON(german_facts)

	// fmt.Println("End")

}

func writeJSON(data []Fact) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("unable to create json file")
		return
	}

	_ = ioutil.WriteFile("germanfacts.json", file, 0644)

}

//
// go get -u github.com/gocolly/colly
// go build -o main
// go run main.go
