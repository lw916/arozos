package authlogger

import (
	"encoding/json"
	"net/http"
	"regexp"
)

//Handle of listing of the logger index (months)
func (l *Logger) HandleIndexListing(w http.ResponseWriter, r *http.Request) {
	indexes := l.ListSummary()
	js, err := json.Marshal(indexes)
	if err != nil {
		sendErrorResponse(w, err.Error())
		return
	}

	sendJSONResponse(w, string(js))
}

//Handle of the listing of a given index (month)
func (l *Logger) HandleTableListing(w http.ResponseWriter, r *http.Request) {
	//Get the record name request for listing
	month, err := mv(r, "record", true)
	if err != nil {
		sendErrorResponse(w, err.Error())
		return
	}

	records, err := l.ListRecords(month)
	if err != nil {
		sendErrorResponse(w, err.Error())
		return
	}

	//Filter the records before sending it to web UI
	results := []LoginRecord{}
	for _, record := range records {
		//Replace the username with a regex filtered one
		reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
		filteredUsername := reg.ReplaceAllString(record.TargetUsername, "░")
		record.TargetUsername = filteredUsername
		results = append(results, record)
	}

	js, _ := json.Marshal(results)
	sendJSONResponse(w, string(js))
}
