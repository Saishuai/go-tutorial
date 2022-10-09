package page

import (
	"errors"
	"example/gowiki/common"
	"html/template"
	"net/http"
	"os"
	"reflect"
	"unsafe"
)

/*
******************

	type HtmlTemplate struct {
		ViewTemplate *template.Template
		EditTemplate *template.Template
	}

var allTemplates = &HtmlTemplate{}

********************
*/
var allTemplates = template.Must(template.ParseFiles("./template/edit.html", "./template/view.html"))

// var allTemplates, _ = template.ParseFiles("edit.html", "view.html")

type Page struct {
	Title    string             `json:"title"`
	Body     []byte             `json:"body"`
	Template *template.Template `json:"template"`
	Meta     string             `json:"meta"`
}

// type PageImp struct {
// 	PageItem Page
// }

func (p *Page) Save() error {
	filename := "./page/source/" + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func (p *Page) Render(w http.ResponseWriter) {
	err := p.Template.ExecuteTemplate(w, p.Meta+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *Page) LoadPage(title string) error {
	filename := "./page/source/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	p.Title = title
	p.Body = body
	p.Template = allTemplates
	return nil
}

func (p *Page) SetValue(field string, value any) error {
	pr := reflect.ValueOf(p).Elem()
	if pr.Kind() != reflect.Struct {
		return errors.New("non-struct can not implete this function")
	}
	// valueR := reflect.TypeOf(value)
	fieldValue := pr.FieldByName(field)
	switch fieldValue.Kind() {
	case reflect.String:
		fieldValue.SetString(value.(string))
	case reflect.Uint8, reflect.Slice:
		fieldValue.SetBytes(value.([]byte))
	case reflect.Pointer:
		fieldValue.SetPointer(value.(unsafe.Pointer))
	}
	return nil
}

func (p *Page) GetValue(field string) (string, error) {
	pr := reflect.ValueOf(p).Elem()
	fieldValue := pr.FieldByName(field)
	if reflect.DeepEqual(fieldValue, reflect.Value{}) {
		return "", errors.New("non-exist field")
	}
	return common.Any(fieldValue), nil
}

func InitPage() PageI {
	return &Page{Template: allTemplates}
}

type PageI interface {
	LoadPage(title string) error
	Render(w http.ResponseWriter)
	Save() error
	SetValue(field string, value any) error
	GetValue(field string) (string, error)
}
