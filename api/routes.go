package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danesparza/centralconfig/datastores"
)

func ShowHelp(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "ShowHelp method")
}

//	Gets a specfic config item based on application and config item name
func GetConfig(rw http.ResponseWriter, req *http.Request) {
	//	req.Body is a ReadCloser -- we need to remember to close it:
	defer req.Body.Close()

	//	Decode the request:
	request := &datastores.ConfigItem{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusBadRequest)
		return
	}

	//	Get the current datastore:
	ds := datastores.GetConfigDatastore()

	//	Send the request to the datastore and get a response:
	response, err := ds.Get(request)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusInternalServerError)
		return
	}

	//	If we found an item, return it (otherwise, return an empty array):
	configItems := []datastores.ConfigItem{}
	if response.Name != "" {
		configItems = append(configItems, response)
		sendDataResponse(rw, "Config item found", configItems)
		return
	}

	sendDataResponse(rw, "No config item found with that application and name", configItems)
}

//	Set a specific config item
func SetConfig(rw http.ResponseWriter, req *http.Request) {
	//	req.Body is a ReadCloser -- we need to remember to close it:
	defer req.Body.Close()

	//	Decode the request:
	request := &datastores.ConfigItem{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusBadRequest)
		return
	}

	//	Get the current datastore:
	ds := datastores.GetConfigDatastore()

	//	Send the request to the datastore and get a response:
	err = ds.Set(request)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusInternalServerError)
	} else {
		configItems := append([]datastores.ConfigItem{}, *request)
		sendDataResponse(rw, "Config item updated", configItems)
	}
}

//	Removes a specific config item
func RemoveConfig(rw http.ResponseWriter, req *http.Request) {
	//	req.Body is a ReadCloser -- we need to remember to close it:
	defer req.Body.Close()

	//	Decode the request:
	request := &datastores.ConfigItem{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusBadRequest)
		return
	}

	//	Get the current datastore:
	ds := datastores.GetConfigDatastore()

	//	Send the request to the datastore and get a response:
	err = ds.Remove(request)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusInternalServerError)
	} else {
		configItems := append([]datastores.ConfigItem{}, *request)
		sendDataResponse(rw, "Config item removed", configItems)
	}
}

//	Gets all config information for a given application
func GetAllConfig(rw http.ResponseWriter, req *http.Request) {
	//	req.Body is a ReadCloser -- we need to remember to close it:
	defer req.Body.Close()

	//	Decode the request:
	request := &datastores.ConfigItem{}
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusBadRequest)
		return
	}

	//	Get the current datastore:
	ds := datastores.GetConfigDatastore()

	//	Send the request to the datastore and get a response:
	configItems, err := ds.GetAll(request.Application)
	if err != nil {
		sendErrorResponse(rw, err, http.StatusInternalServerError)
		return
	}

	//	If we found an item, return it (otherwise, return an empty array):
	if len(configItems) > 0 {
		sendDataResponse(rw, "Config items found", configItems)
		return
	}

	sendDataResponse(rw, "No config items found with that application", configItems)
}

//	Initializes a store
func InitStore(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "InitStore method")
}

//	Used to send back an error:
func sendErrorResponse(rw http.ResponseWriter, err error, code int) {
	//	Our return value
	response := datastores.ConfigResponse{
		Status:  code,
		Message: "Error: " + err.Error()}

	//	Serialize to JSON & return the response:
	rw.WriteHeader(code)
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(rw).Encode(response)
}

//	Used to send back a response with data
func sendDataResponse(rw http.ResponseWriter, message string, configItems []datastores.ConfigItem) {
	//	Our return value
	response := datastores.ConfigResponse{
		Status:  http.StatusOK,
		Message: message,
		Data:    configItems}

	//	Serialize to JSON & return the response:
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(rw).Encode(response)
}
