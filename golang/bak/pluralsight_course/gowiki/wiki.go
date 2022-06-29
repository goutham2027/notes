package main

import (
	"fmt"
	"io/ioutil"
)

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

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page")}
	p1.save()

	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
