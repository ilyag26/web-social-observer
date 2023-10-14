package api

import (
	"net/http"
	"text/template"

	d "github.com/Username/Project-Name/web/data"
	n "github.com/Username/Project-Name/web/scraping"
)

var (
	videoURL, channelURL, accountURL string
)

var indexTpl = template.Must(template.ParseFiles("templates/index.gohtml"))
var videoCheckProcessTpl = template.Must(template.ParseFiles("templates/videoCheckProcess.gohtml"))
var videoCheckTpl = template.Must(template.ParseFiles("templates/videocheck.gohtml"))
var channelCheckProcessTpl = template.Must(template.ParseFiles("templates/channelCheckProcess.gohtml"))
var channelCheckTpl = template.Must(template.ParseFiles("templates/channelcheck.gohtml"))
var tiktokPage = template.Must(template.ParseFiles("templates/tiktok.gohtml"))
var youtubePage = template.Must(template.ParseFiles("templates/youtube.gohtml"))
var tiktokCheckTpl = template.Must(template.ParseFiles("templates/tiktokcheck.gohtml"))
var tiktokCheckProgressTpl = template.Must(template.ParseFiles("templates/tiktokcheckprogress.gohtml"))

func index(w http.ResponseWriter, r *http.Request) {
	indexTpl.Execute(w, nil)
}

func videoCheck(w http.ResponseWriter, r *http.Request) {
	videoCheckTpl.Execute(w, nil)
}

func channelCheck(w http.ResponseWriter, r *http.Request) {
	channelCheckTpl.Execute(w, nil)
}

func tiktokP(w http.ResponseWriter, r *http.Request) {
	tiktokPage.Execute(w, nil)
}

func youtubeP(w http.ResponseWriter, r *http.Request) {
	youtubePage.Execute(w, nil)
}

func tiktokCheck(w http.ResponseWriter, r *http.Request) {
	tiktokCheckTpl.Execute(w, nil)
}

func tiktokCheckProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	accountURL = r.FormValue("accountURL")
	n.ScrapeAccount(accountURL)
	accountData := d.AccountData{
		Name:  n.NameAcc,
		Desc:  n.DescAcc,
		Likes: n.LikesAcc,
		Subs:  n.SubsAcc,
	}
	tiktokCheckProgressTpl.Execute(w, accountData)
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
	http.HandleFunc("/tiktok", tiktokP)
	http.HandleFunc("/youtube", youtubeP)
	http.HandleFunc("/ttcheck", tiktokCheck)
	http.HandleFunc("/ttcheckprocess", tiktokCheckProcess)
	http.ListenAndServe(":8080", nil)
}
