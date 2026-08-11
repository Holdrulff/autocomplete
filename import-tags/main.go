package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type tag struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type tagsResponse struct {
	Items          []tag `json:"items"`
	HasMore        bool  `json:"has_more"`
	QuotaRemaining int   `json:"quota_remaining"`
	Backoff        int   `json:"backoff"`
}

type snapshot struct {
	Source      string    `json:"source"`
	Attribution string    `json:"attribution"`
	GeneratedAt time.Time `json:"generated_at"`
	Tags        []tag     `json:"tags"`
}

var httpClient = &http.Client{
	Timeout: 15 * time.Second,
}

func fetchPage(page int, apiKey string) (tagsResponse, error) {
	endpoint := "https://api.stackexchange.com/2.3/tags"

	params := url.Values{}
	params.Set("site", "stackoverflow")
	params.Set("order", "desc")
	params.Set("sort", "popular")
	params.Set("page", strconv.Itoa(page))
	params.Set("pagesize", "100")
	params.Set("key", apiKey)

	requestURL := endpoint + "?" + params.Encode()

	request, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return tagsResponse{}, fmt.Errorf("create request: %w", err)
	}

	response, err := httpClient.Do(request)

	if err != nil {
		return tagsResponse{}, fmt.Errorf("fetch API: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return tagsResponse{}, fmt.Errorf("unexpected status: %s", response.Status)
	}

	var data tagsResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return tagsResponse{}, fmt.Errorf("JSON interpreter: %w", err)
	}

	return data, nil
}

func normalizeAndDeduplicate(tags []tag) []tag {
	uniqueTags := make(map[string]tag, len(tags))

	for _, currentTag := range tags {
		normalizedName := strings.ToLower(strings.TrimSpace(currentTag.Name))

		if normalizedName == "" {
			continue
		}

		currentTag.Name = normalizedName

		existingTag, alreadyExists := uniqueTags[normalizedName]
		if !alreadyExists || currentTag.Count > existingTag.Count {
			uniqueTags[normalizedName] = currentTag
		}
	}

	normalizedTags := make([]tag, 0, len(uniqueTags))

	for _, currentTag := range uniqueTags {
		normalizedTags = append(normalizedTags, currentTag)
	}

	return normalizedTags
}

func writeSnapshot(path string, tags []tag) error {
	if err := os.MkdirAll("data", 0o755); err != nil {
		return fmt.Errorf("create data directory: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create snapshot file: %w", err)
	}
	defer file.Close()

	data := snapshot{
		Source:      "https://api.stackexchange.com/2.3/tags?site=stackoverflow",
		Attribution: "Stack Overflow data via Stack Exchange API; user contributions licensed under CC BY-SA.",
		GeneratedAt: time.Now().UTC(),
		Tags:        tags,
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("encode snapshot: %w", err)
	}

	return nil
}

func main() {
	const maxTags = 5000
	const snapshotPath = "data/stack-overflow-tags.json"

	apiKey := os.Getenv("STACK_EXCHANGE_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "STACK_EXCHANGE_API_KEY was not defined")
		os.Exit(1)
	}

	var allTags []tag

	for page := 1; len(allTags) < maxTags; page++ {
		data, err := fetchPage(page, apiKey)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}

		remaining := maxTags - len(allTags)
		if len(data.Items) > remaining {
			data.Items = data.Items[:remaining]
		}

		allTags = append(allTags, data.Items...)

		fmt.Printf(
			"page %d: %d tags; remaining quota: %d\n",
			page,
			len(data.Items),
			data.QuotaRemaining,
		)

		if !data.HasMore {
			break
		}

		if data.Backoff > 0 {
			fmt.Printf("API requestes %d seconds\n", data.Backoff)

			time.Sleep(time.Duration(data.Backoff) * time.Second)
		}
	}

	collectedCount := len(allTags)

	allTags = normalizeAndDeduplicate(allTags)

	fmt.Printf(
		"\nTags únicas: %d; duplicadas removidas: %d\n",
		len(allTags),
		collectedCount-len(allTags),
	)

	sort.Slice(allTags, func(i, j int) bool {
		left := allTags[i]
		right := allTags[j]

		if left.Count == right.Count {
			return left.Name < right.Name
		}

		return left.Count > right.Count
	})

	if err := writeSnapshot(snapshotPath, allTags); err != nil {
		fmt.Fprintln(os.Stderr, "erro ao salvar snapshot:", err)
		os.Exit(1)
	}

	fmt.Println("Snapshot salvo em:", snapshotPath)

	previewCount := 20

	if len(allTags) < previewCount {
		previewCount = len(allTags)
	}

	fmt.Println("\n20 most popular tags:")

	for _, currentTag := range allTags[:previewCount] {
		fmt.Printf("%-25s %d\n", currentTag.Name, currentTag.Count)
	}

	fmt.Println("total tags:", len(allTags))
}
