package restapi

import (
	"DADS/configer/Configurator"
	"fmt"
	"net/http"
)

type DBlayerconfig struct {
	DB string `json:"database"`
	Conn string `json:"connectionstring"`
}


func InitializeAPIHandlers() error {
	conf := new(DBlayerconfig)
	// first layer to read configuration: json configuration -- REST API server
	err := Configurator.GetConfiguration(Configurator.JSON, conf, "../../apiconfig.json")

	if err != nil {
		fmt.Println("Error decoding JSON", err)
		return err
	}
	// second layer to build the request handler: API server -- request handler -- DB Layer --- DB
	h := NewDADSReqHandler()
	// connect to db
	err = h.connect(conf.DB, conf.Conn)

	if err != nil {
		fmt.Println("Error connecting to db ", err)
		return err
	}
	// define the path for request handler
	http.HandleFunc("/dads", h.handleRequests)
	return nil
}

func RUNAPI() error {
	if err := InitializeAPIHandlers(); err != nil {
		return err
	}

	http.ListenAndServe(":8061", nil)
	return nil
}