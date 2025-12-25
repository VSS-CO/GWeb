package core

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// ---------------- ServePage ----------------

// ---------------- Router ----------------
func Router(distDir string) {
	// Static file server
	fs := http.FileServer(http.Dir(distDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Example page using your new renderer
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

	// Start HMR watcher
	go WatchFilesForChanges(distDir)

	// Start server
	fmt.Println("GWeb dev server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// ---------------- Hot Module Reload (HMR) ----------------
func WatchFilesForChanges(distDir string) {
	lastMod := make(map[string]time.Time)

	for {
		filepath.Walk(distDir, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}
			modTime := info.ModTime()
			if t, ok := lastMod[path]; !ok || t.Before(modTime) {
				fmt.Printf("File changed: %s\n", path)
				lastMod[path] = modTime
			}
			return nil
		})
		time.Sleep(500 * time.Millisecond) // poll every 0.5s
	}
}
