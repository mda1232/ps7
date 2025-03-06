package handler

import (
	"bytes"
	"html/template"
	"net/http"
	"resume-website/internal/api"
	"resume-website/internal/models"
)

var renderedPage []byte

func RenderTemplate(userData models.UserData) error {
	tmpl, err := template.ParseFiles("templates/resume.html")
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, userData); err != nil {
		return err
	}

	renderedPage = buf.Bytes()
	return nil
}

func ResumeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write(renderedPage)
}

func UpdateHandler(cfg *models.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userData, err := api.GetAllUserData(cfg)
		if err != nil {
			http.Error(w, "Ошибка обновления данных: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if err := RenderTemplate(userData); err != nil {
			http.Error(w, "Ошибка рендеринга шаблона: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Страница успешно обновлена!"))
	}
}
