package restapi

import (
	"distributed_anomaly_detection_system/internal/databasehandlers"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type DADSReqHandler struct {
	dbConn databasehandlers.DBLayer
}

func NewDADSReqHandler() *DADSReqHandler {
	return new(DADSReqHandler)
}

func (reqhandler *DADSReqHandler) connect(o, conn string) error {
	dblayer, err := databasehandlers.ConnectDatabase(o, conn)
	if err != nil {
		return err
	}
	reqhandler.dbConn = dblayer
	return nil
}

func (reqhandler *DADSReqHandler) handleRequests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// /hydracrew/2
		ids := r.RequestURI[len("/hydracrew"):]
		id, err := strconv.Atoi(ids)
		if err != nil {
			// check the id
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "id %s provides is not of valid number. \n", ids)
			return
		} 
		
		// find the member with that id, cm is the crew member
		cm, err := reqhandler.dbConn.FindMember(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error %s occured while searching for id %d \n ", err.Error(), id)
			return
		}
		json.NewEncoder(w).Encode(&cm)
	
	case "POST":
		cm := new(databasehandlers.CrewMember)
		// fill the cm with bytes json body
		err :=  json.NewDecoder(r.Body).Decode(cm)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error %s occurred", err)
			return
		}
		
		err = reqhandler.dbConn.AddMember(cm)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error %s occurred while adding a member to dataset", err)
			return
		}
		
		fmt.Fprintf(w, "Successfully inserted id %d \n", cm.ID)
		}	

}