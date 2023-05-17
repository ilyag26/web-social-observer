package api

import (
	"net/http"
	"text/template"

	n "github.com/Username/Project-Name/web/scraping"
)

type Data struct {
	Name  string
	Desc  string
	Date  string
	Views string
}

var tpl = template.Must(template.ParseFiles("web/templates/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("web/templates/styles"))))
	data := Data{
		Name:  n.NameRet,
		Desc:  n.DescRet,
		Date:  n.DateRet,
		Views: n.ViewsRet,
	}
	tpl.Execute(w, data)
}

func StartServer() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
