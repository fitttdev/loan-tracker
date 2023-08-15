package core

import (
	"fmt"
	"net/http"
	"io"
)
const RANDOM_QUOTE_API = "https://api.quotable.io/random"

func FetchRandomQuote() {
	resp, err := http.Get(RANDOM_QUOTE_API)

	if err != nil {
		fmt.Println("Error fetching quote:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Random Quote:")
	fmt.Println(string(body))
}
