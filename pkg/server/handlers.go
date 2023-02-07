package server

import (
	"encoding/json"
	"go-altitude/pkg/service"
	"html/template"
	"net/http"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if !exists("./data/data.xlsx") {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.Header().Set("Content-Disposition", "attachment; filename=data.xlsx")
		w.Header().Set("Content-Type", "application/octet-stream")
		http.ServeFile(w, r, "./data/data.xlsx")
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		var request requestT
		err := decoder.Decode(&request)
		if err != nil {
			errorLogger.Println(err.Error())
			answerReport(w, http.StatusBadRequest, err.Error())
			return
		}
		err = service.Service.GetAltitudes(request.A, request.B, request.LongitudeStep, request.LatitudeStep, request.Source)
		if err != nil {
			errorLogger.Println(err.Error())
			answerReport(w, http.StatusBadRequest, err.Error())
			return
		}
		answerReport(w, http.StatusOK, "OK")
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		files, err := template.ParseFiles("web/main.html")
		if err != nil {
			errorLogger.Println(err.Error())
			return
		}
		err = files.Execute(w, nil)
		if err != nil {
			errorLogger.Println(err.Error())
			return
		}
	}
}

func answerReport(w http.ResponseWriter, status int, comment string) {
	w.WriteHeader(status)
	data := &response{
		Comment: comment,
	}
	jdata, err := json.Marshal(data)
	if err != nil {
		errorLogger.Printf("json marshal error: %s\n", err.Error())
		return
	}
	_, err = w.Write(jdata)
	if err != nil {
		errorLogger.Printf("write answer error: %s\n", err.Error())
		return
	}
}
