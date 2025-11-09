package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	ntfyTopic := "your0dhaJobAlert"

	c := colly.NewCollector(
		colly.AllowedDomains("careers.zerodha.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML("#page-job-openings", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "There are no job openings currently.") {
			fmt.Println("No job openings available at the moment.")
		} else {
			fmt.Println("Job openings are available!")
			req, _ := http.NewRequest("POST", "https://ntfy.sh/"+ntfyTopic, strings.NewReader("Job openings are available at Zerodha! Check https://careers.zerodha.com/"))
			req.Header.Set("Actions", "view, Open Zerodha Careers, https://careers.zerodha.com/")
			req.Header.Set("X-Icon", "https://zerodha.com/static/images/favicon.png")
			http.DefaultClient.Do(req)
		}
	})

	c.Visit("https://careers.zerodha.com/")
}