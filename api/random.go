package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type WikiResponse struct {
	Query struct {
		Random []struct {
			Title string `json:"title"`
		} `json:"random"`
	} `json:"query"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// seed
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get("https://en.wikipedia.org/w/api.php?action=query&format=json&list=random&rnnamespace=0&rnlimit=1")
	if err != nil {
		http.Error(w, "failed to fetch random article", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var wikiResp WikiResponse
	if err := json.NewDecoder(resp.Body).Decode(&wikiResp); err != nil {
		http.Error(w, "failed to parse wikipedia response", http.StatusInternalServerError)
		return
	}

	if len(wikiResp.Query.Random) == 0 {
		http.Error(w, "no random article found", http.StatusInternalServerError)
		return
	}

	title := wikiResp.Query.Random[0].Title
	wikiURL := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", title)

	http.Redirect(w, r, wikiURL, http.StatusFound)
}
