package rendering

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Profiles []string
	Rolls    string
}

func ServeTemplate(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("public/html/index.html"))

	data := PageData{
		Profiles: []string{"Default"},
		Rolls:    "Roll Data",
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
