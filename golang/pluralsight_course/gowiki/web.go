package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

/*
The function handler is of the type http.HandlerFunc. It takes an http.ResponseWriter and an
http.Request as its arguments.

An http.ResponseWriter assembles the HTTP server's response.

An http.Request is a data structure that represents the client HTTP request.
*/
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World %s", r.URL.Path[1:])
}

type Page struct {
	Title string
	Body  []byte
}

/*
This is a method named save that takes as its receiver p, a pointer to Page. It takes no
parameters, and returns a value of type error.

The save method returns an error value because that is the return type of WriteFile.
*/
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

/*
Callers to loadPage can now check the second parameter, if it is nil then it has successfully
loaded a Page. If not, it will be an error that can be handled by the caller.
*/
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	/*
		The function template.ParseFiles will read the contents of edit.html and return a *template.Template
	*/
	t, err := template.ParseFiles(tmpl + ".html")

	if err != nil {
		/*
			http.Error function sends a specified HTTP response code and error message.
		*/
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*
		The method t.Execute executes the template, writing the generated HTML to the http.ResponseWriter.
		The .Title and .Body dotted identifiers refer to p.Title and p.Body
	*/
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
 */
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)

	if err != nil {
		// adds an HTTP status code of http.StatusFound (302) and a Location header to the HTTP Response.
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

/*
 */
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)

	//fmt.Fprintf(w, "<h1> Editing %s</h1>"+
	//"<form action=\"/save/%s\" method=\"POST\">"+
	//"<textarea name=\"body\">%s</textarea><br/>"+
	//"<input type=\"submit\" value=\"Save\">"+
	//"</form>",
	//p.Title, p.Title, p.Body)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

/*
The main function begins with a call to http.HandleFunc, which tells the http package to handle all
requests to the web root("/") with handler.

It then calls http.ListenAndServe, specifying that it should listen on port 8080 on any interface.
The function will block until the program is terminated.

ListenAndServe always return an error, since it only returns when an unexpected error occurs.
In order to log that error we wrap the function call with log.Fatal.
*/
func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
