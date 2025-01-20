package main

import (
	"cyoa"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the cyoa web app")
	filename := flag.String("file", "stories.json", "json file with story")
	flag.Parse()
	fmt.Printf("Using the story in %s. \n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	arcs, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%+v\n", arcs)
	// fmt.Println(arcs["debate"].ArcSuggestions[0].ArcName)

	// tpl := template.Must(template.New("").Parse("Hello"))
	// h := cyoa.NewHandler(arcs, cyoa.WithTemplate(tpl))

	tpl := template.Must(template.New("").Parse(storyTpl))

	h := cyoa.NewHandler(arcs,
		cyoa.WithTemplate(tpl),
		cyoa.WithPathFunc(pathFn),
	)

	mux := http.NewServeMux()
	mux.Handle("/story/", h)

	fmt.Printf("Starting the server on port %d\n", *port)
	addr := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

var storyTpl = `
<!DOCTYPE html>
<html>
  <head>
    <title>Choose your own adventure</title>
  </head>
  <body>
    <h1>"{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <ul>
    {{range .ArcSuggestions}}
    <li><a href="/story/{{.ArcName}}"> {{.SuggestionText}}</a></li>
    {{end}}
  </ul>
    </body>
</html>`
