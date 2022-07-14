package util

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func ParseTemplate(w http.ResponseWriter, tmpl string, data any) {
	short := fmt.Sprintf("%s.tmpl.html", tmpl)
	t, err := template.New(short).ParseFiles(
		"templates/tmpl/base.tmpl.html",
		"templates/tmpl/common.tmpl.html",

		fmt.Sprintf("templates/%s", short),
	)

	if err != nil {
		log.Println("could not read", tmpl, err)
		os.Exit(1)
	}

	t.Execute(w, data)
}
