package api

import (
	"net/http"
	"text/template"

	d "github.com/Username/Project-Name/web/data"
	n "github.com/Username/Project-Name/web/scraping"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))
var tpl2 = template.Must(template.ParseFiles("templates/processor.html"))
var ChannelYou string

func index(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
func process(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	ChannelYou = r.FormValue("channel")
	n.Scrape(ChannelYou)
	data := d.Data{
		Name:  n.NameRet,
		Desc:  n.DescRet,
		Date:  n.DateRet,
		Views: n.ViewsRet,
	}
	tpl2.Execute(w, data)
}
func StartServer() {
	fs := http.FileServer(http.Dir("assets/styles"))
	http.Handle("/styles/", http.StripPrefix("/styles", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/process", process)
	http.ListenAndServe(":8080", nil)
}
