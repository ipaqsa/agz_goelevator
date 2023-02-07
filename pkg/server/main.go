package server

import (
	"go-altitude/pkg/logger"
	"go-altitude/pkg/service"
	"net/http"
)

var errorLogger = logger.NewLogger("ERROR")

func Run(port string) error {
	service.CreateService()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.HandleFunc("/download", downloadHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}
	return nil
}
