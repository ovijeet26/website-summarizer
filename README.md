# 🧠 Website Summarizer

This is a Go application that fetches the contents of any public webpage, extracts its readable content, and sends it to a local LLM API (e.g. [Ollama](https://ollama.com/)) for summarization in Markdown format.

It uses a custom prompt to instruct the model to summarize only relevant content, ignoring navigation, ads, or other noise.

## ✨ Features

- Scrapes website title and body text
- Removes irrelevant elements like scripts, styles, inputs, and images
- Generates context-aware prompts for summarization
- Sends content to locally hosted LLMs (e.g., LLaMA 3) using OpenAI-compatible API format
- Outputs summaries in Markdown format

## 📦 Tech Stack

- **Go** (Golang)
- **goquery** (HTML parsing similar to BeautifulSoup)
- **net/http** (API communication)
- Optional: **glamour** (terminal markdown rendering)

## 📁 Project Structure

website-summarizer/
├── main.go
├── go.mod
└── src/
└── core/
├── website.go # Web scraping and prompt generation
├── llm.go # LLM API interaction
└── model.go # Shared message/request/response structs


## 🚀 Getting Started

### Prerequisites

- Go 1.20+
- [Ollama](https://ollama.com/) running locally (or another OpenAI-compatible LLM endpoint)
- Installed model (e.g. `ollama run llama3`)

### Installation

git clone https://github.com/yourusername/website-summarizer.git
cd website-summarizer
go mod tidy
go run main.go


## 🧪 Example Usage

### Basic Usage

summary, err := Summarize("https://example.com")
if err != nil {
log.Fatal(err)
}
fmt.Println(summary)


### Sample Output

Example.com
Example.com is a placeholder website often used in documentation. It provides a basic HTML structure for demonstration purposes and does not contain real content.

There are no news or announcements on the page.


## ⚙️ Configuration

Modify the LLM settings in `llm.go`:

const (
OpenAIEndpoint = "http://localhost:11434/v1/chat/completions"
ModelName = "llama3.2"
APIKey = "ollama"
)


---

_Project maintained by [Ovijeet](https://github.com/ovijeet26) | [Report Issue](https://github.com/ovijeet26/website-summarizer/issues)_
