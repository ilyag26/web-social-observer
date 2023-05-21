package api

import (
	"net/http"
	"text/template"

	d "github.com/Username/Project-Name/web/data"
	n "github.com/Username/Project-Name/web/scraping"
)

var indexTpl = template.Must(template.ParseFiles("templates/index.gohtml"))
var videoCheckProcessTpl = template.Must(template.ParseFiles("templates/videoCheckProcess.gohtml"))
var videoCheckTpl = template.Must(template.ParseFiles("templates/videocheck.gohtml"))
var channelCheckProcessTpl = template.Must(template.ParseFiles("templates/channelCheckProcess.gohtml"))
var channelCheckTpl = template.Must(template.ParseFiles("templates/channelcheck.gohtml"))
var videoURL string
var channelURL string

func index(w http.ResponseWriter, r *http.Request) {
	indexTpl.Execute(w, nil)
}
func videoCheck(w http.ResponseWriter, r *http.Request) {
	videoCheckTpl.Execute(w, nil)
}
func channelCheck(w http.ResponseWriter, r *http.Request) {
	channelCheckTpl.Execute(w, nil)
}
func videoCheckProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	videoURL = r.FormValue("videoURL")
	n.ScrapeVideo(videoURL)
	videoData := d.VideoData{
		Name:  n.NameVideoRet,
		Desc:  n.DescVideoRet,
		Date:  n.DateVideoRet,
		Views: n.ViewsVideoRet,
	}
	videoCheckProcessTpl.Execute(w, videoData)
}
func channelCheckProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	channelURL = r.FormValue("channelURL")
	n.ScrapeChannel(channelURL)
	dataChannel := d.ChannelData{
		Name:  n.NameChannelRet,
		Desc:  n.DescChannelRet,
		Views: n.ViewsChannelRet,
		Date:  n.DateChannelRet,
		Subs:  n.SubsChannelRet,
	}
	channelCheckProcessTpl.Execute(w, dataChannel)
}
func StartServer() {
	fs := http.FileServer(http.Dir("assets/styles"))
	http.Handle("/styles/", http.StripPrefix("/styles", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/channelcheck", channelCheck)
	http.HandleFunc("/videocheck", videoCheck)
	http.HandleFunc("/videocheckprocess", videoCheckProcess)
	http.HandleFunc("/channelcheckprocess", channelCheckProcess)
	http.ListenAndServe(":8080", nil)
}
