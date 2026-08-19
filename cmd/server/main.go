package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Holdrulff/autocomplete/graph"
	"github.com/Holdrulff/autocomplete/internal/autocomplete"
	"github.com/Holdrulff/autocomplete/internal/httpapi"
)

const (
	defaultSnapshotPath  = "data/stack-overflow-tags.json"
	defaultServerAddress = ":8080"
)

func buildHandler(snapshotPath string) (http.Handler, error) {
	catalog, err := autocomplete.LoadCatalog(snapshotPath)
	if err != nil {
		return nil, fmt.Errorf("load catalog: %w", err)
	}

	trie := autocomplete.NewTrieFromCatalog(catalog)

	service, err := autocomplete.NewService(trie)
	if err != nil {
		return nil, fmt.Errorf("create autocomplete service: %w", err)
	}

	resolver, err := graph.NewResolver(service)
	if err != nil {
		return nil, fmt.Errorf("create graphql resolver: %w", err)
	}

	graphqlHandler, err := httpapi.NewHandler(resolver)
	if err != nil {
		return nil, fmt.Errorf("create graphql handler: %w", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/query", graphqlHandler)
	mux.HandleFunc("GET /health", healthHandler)

	return mux, nil
}

func healthHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(http.StatusOK)

	_, _ = fmt.Fprintln(writer, "ok")
}

func main() {
	handler, err := buildHandler(defaultSnapshotPath)
	if err != nil {
		log.Fatalf("initialize server: %v", err)
	}

	server := &http.Server{
		Addr:              defaultServerAddress,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf(
		"graphql server listening at http://localhost%s/query",
		defaultServerAddress,
	)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("serve graphql API: %v", err)
	}
}
