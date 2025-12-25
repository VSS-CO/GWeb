package core

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

// HMR clients
var hmrClients = make(map[chan string]struct{})

// StartServer starts the dev server with HMR and static files
func StartServer() {
	fmt.Println("Starting GWeb dev server on http://localhost:8080")

	// Start file watcher for HMR
	go WatchFilesPolling()

	// Setup routes
	SetupRoutes()

	// Serve static files
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// HMR endpoint
	http.HandleFunc("/__hmr", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming not supported", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")

		ch := make(chan string)
		hmrClients[ch] = struct{}{}
		defer func() { delete(hmrClients, ch) }()

		for msg := range ch {
			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// SetupRoutes replaces the old RegisterRoutes
func SetupRoutes() {
	// Root page example using new renderer
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := Div{
			Children: []Node{
				Header{
					Children: []Node{
						H1{Text: "Welcome to GWeb"},
						P{Text: "This page is rendered entirely in Go!"},
					},
				},
				Main{
					Children: []Node{
						Section{
							Children: []Node{
								Article{
									Children: []Node{
										P{Text: fmt.Sprintf("Requested path: %s", r.URL.Path)},
									},
								},
							},
						},
					},
				},
				Footer{
					Children: []Node{
						P{Text: "Made with ❤️ using GWeb"},
					},
				},
			},
		}

		ServePage(w, r, page)
	})
}

// ServePage renders a Node and writes it to the ResponseWriter
func ServePage(w http.ResponseWriter, r *http.Request, page Node) {
	html := Render(page)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}

// WatchFilesPolling watches "routes" and "core" folders for changes and triggers HMR
func WatchFilesPolling() {
	lastModTimes := make(map[string]time.Time)

	for {
		changed := false

		folders := []string{"routes", "core"}
		for _, folder := range folders {
			filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
				if err != nil || info.IsDir() {
					return nil
				}
				modTime := info.ModTime()
				if t, ok := lastModTimes[path]; !ok || modTime.After(t) {
					lastModTimes[path] = modTime
					changed = true
				}
				return nil
			})
		}

		if changed {
			log.Println("[GWeb HMR] Detected file changes")
			for ch := range hmrClients {
				go func(c chan string) { c <- "reload" }(ch)
			}
		}

		time.Sleep(time.Second)
	}
}
