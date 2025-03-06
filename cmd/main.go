package main

import (
	"log"
	"net/http"

	"resume-website/config"
	"resume-website/internal/api"
	"resume-website/internal/handler"
)

func main() {
	
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	userData, err := api.GetAllUserData(cfg)
	if err != nil {
		log.Fatalf("Ошибка получения данных пользователя: %v", err)
	}

	if err := handler.RenderTemplate(userData); err != nil {
		log.Fatalf("Ошибка рендеринга шаблона: %v", err)
	}

	http.HandleFunc("/", handler.ResumeHandler)
	http.HandleFunc("/update", handler.UpdateHandler(cfg))

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
