package core

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type website struct {
	URL   string
	Title string
	Text  string
}

const SystemPrompt = `You are an assistant that analyzes the contents of a website
and provides a short summary, ignoring text that might be navigation related.
Respond in markdown.`

func NewWebsite(url string) (*website, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return scrape(url, resp)
}

func scrape(url string, resp *http.Response) (*website, error) {

	// Parse HTML with goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	// Get title
	title := doc.Find("title").First().Text()
	if strings.TrimSpace(title) == "" {
		title = "No title found"
	}

	// Remove irrelevant elements
	doc.Find("script, style, img, input").Remove()

	// Get body text
	bodySelection := doc.Find("body")
	text := strings.TrimSpace(bodySelection.Text())

	return &website{
		URL:   url,
		Title: title,
		Text:  text,
	}, nil
}

func (w *website) UserPrompt() string {
	userPrompt := fmt.Sprintf("You are looking at a website titled %s\n", w.Title)
	userPrompt += "The contents of this website is as follows; " +
		"please provide a short summary of this website in markdown. " +
		"If it includes news or announcements, then summarize these too.\n\n"
	userPrompt += w.Text
	return userPrompt
}

func (w *website) MessagesFor() []Message {
	return []Message{
		{
			Role:    "system",
			Content: SystemPrompt,
		},
		{
			Role:    "user",
			Content: w.UserPrompt(),
		},
	}
}
