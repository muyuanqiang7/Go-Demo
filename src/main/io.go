package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
)

type Page struct {
	title string
	body  []byte
}

func (p *Page) save() error {
	filename := p.title + ".txt"
	return ioutil.WriteFile(filename, p.body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{title: title, body: body}, nil
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/view/"):]
	page, _ := loadPage(title)
	fmt.Fprintf(writer, "<h1>%s</h1><div>%s</div>", title, page.body)
}

func main() {
	page := &Page{title: "TestPage", body: []byte("this is simple test page")}
	saveErr := page.save()
	if saveErr != nil {
		fmt.Println(saveErr)
		return
	}
	readPage, err := loadPage("TestPage")
	if err != nil {

	} else {
		fmt.Println(string(readPage.body))
	}
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
