package dadsportal

import (
	"distributed_anomaly_detection_system/configer"
	"distributed_anomaly_detection_system/internal/databasehandlers/dadsdblayer"
	"distributed_anomaly_detection_system/web/restapi"
	"html/template"
	"log"
	"net/http"
)

var DADSWebTemplate *template.Template

func Run() error {
	var err error
	DADSWebTemplate, err = template.ParseFiles("./cover/Crew/crew.html", "./cover/about/about.html")

	if err != nil {
		return err
	}

	conf := struct {
		Filespath string `json:"filespath"`
	}{}
	// define the type, object, configfile
	err = configer.GetConfiguration(configer.JSON, &conf, "../portalconfig.json")
	if err != nil {
		return err
	}

	restapi.InitializeAPIHandlers()
	log.Println(conf.Filespath)
	fs := http.FileServer(http.Dir(conf.Filespath))
	// create a static website
	http.Handle("/", fs)
	http.HandleFunc("/Crew/", crewhandler)
	http.HandleFunc("/about/", abouthandler)
	return http.ListenAndServe(":8061", nil)

}

func crewhandler(w *http.ResponseWriter, r *http.Request) {
	// get the database object
	dblayer, err := dadsdblayer.ConnectDatabase("mysql", "root:admin@/DADS")
	if err != nil {
		return
	}
	all, err := dblayer.AllMembers()
	if err != nil {
		return
	}
	// merge the all members with crew to w, data will show when loading page
	// crew page has been cached before
	err = DADSWebTemplate.ExecuteTemplate(w, "crew.html", all)
	if err != nil {
		log.Println(err)
	}
}

func abouthandler(w *http.ResponseWriter, r *http.Request) {
	about := struct {
		Msg string `json:"message"`
	}{}

	err := configer.GetConfiguration(configer.JSON, &about, "../about.json")
	if err != nil {
		return
	}

	err = DADSWebTemplate.ExecuteTemplate(w, "about.html", about)
	if err != nil {
		log.Println(err)
	}

}