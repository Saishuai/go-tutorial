package main

import (
	"errors"
	"example/gowiki/page"
	"log"
	"net/http"
	"regexp"
)

// set regexp rule
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getValidTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	titlePath := validPath.FindStringSubmatch(r.URL.Path)
	if titlePath == nil {
		http.NotFound(w, r)
		return "", errors.New("url path doesn't match regexp rule")
	}
	return titlePath[2], nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	if len(title) == 0 {
		return
	}
	var p = page.InitPage()
	err := p.LoadPage(title)
	if err != nil {
		// fmt.Fprintf(w, "Load page /%q error, redirect to edit page.", filename)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	err = p.SetValue("Meta", "view")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	// t, _ := template.ParseFiles("view.html")
	// t.Execute(w, p)
	p.Render(w)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	if len(title) == 0 {
		return
	}
	var p = page.InitPage()
	err := p.LoadPage(title)
	if err != nil {
		// p = &Page{Title: title, Template: allTemplates}
		p.SetValue("Title", title)
	}
	// p.Meta = "edit"
	err = p.SetValue("Meta", "edit")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Fprintf(w, "<h1>editing %s</h1>"+
	// 	"<form action=\"/save/%s\" method=\"POST\">"+
	// 	"<testarea name=\"body\">%s</testarea><br>"+
	// 	"<input type=\"submit\" value=\"Save\">"+
	// 	"</form>",
	// 	p.Title, p.Title, p.Body)
	// t, _ := template.ParseFiles("edit.html")
	// t.Execute(w, p)
	p.Render(w)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	if len(title) == 0 {
		return
	}
	newBody := r.FormValue("body")
	var p = page.InitPage()
	// p := &Page{Title: title, Body: []byte(newBody)}
	p.SetValue("Title", title)
	p.SetValue("Body", []byte(newBody))
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func Init() error {
	return nil
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// extract `title` from request
		// use the input param fn
		title, err := getValidTitle(w, r)
		if err != nil {
			fn(w, r, "")
		}
		fn(w, r, title)
	}
}

func main() {
	/*********************
	p1 := &Page{Title: "Test", Body: []byte("Hello! Body of page.")}
	err := p1.Save()
	if err != nil {
		log.Fatalf("Save page stream error.")
	}
	p2, err2 := loadPage("Test")
	if err2 != nil {
		log.Fatalf("Load page to stream error.")
	}
	fmt.Printf("Load page from file, title: %q, body: %q", p2.Title, string(p2.Body))
	*********************/

	err := Init()
	if err != nil {
		log.Fatalf("Load template error: %s", err.Error())
		return
	}
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
