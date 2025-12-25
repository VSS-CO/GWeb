package routes

import (
	"net/http"

	"co.vss.gweb/core"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	page := core.Div{
		Children: []core.Node{
			core.Header{
				Children: []core.Node{
					core.H1{Text: "Welcome to GWeb"},
					core.P{Text: "This page is rendered entirely in Go!"},
				},
			},
			core.Main{
				Children: []core.Node{
					core.Section{
						Children: []core.Node{
							core.Article{
								Children: []core.Node{
									core.P{Text: "Hello World!"},
								},
							},
						},
					},
				},
			},
			core.Footer{
				Children: []core.Node{
					core.P{Text: "Made with ❤️ using GWeb"},
				},
			},
		},
	}

	// Render the page and send it to the browser
	core.ServePage(w, r, page)
}
