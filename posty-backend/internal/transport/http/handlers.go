package http

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Проверяем размер загружаемого файла
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return
	} // Ограничиваем размер до 10 МБ

	// Получаем файл из формы
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error retrieving file from form-data: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Генерируем уникальное имя файла
	imageID := time.Now().UnixNano()
	filename := fmt.Sprintf("%d_%s", imageID, handler.Filename)
	filepath := path.Join("uploads", filename)

	// Создаем директорию uploads, если ее нет
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		log.Printf("Error creating uploads directory: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Создаем файл на сервере
	f, err := os.Create(filepath)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Копируем загруженный файл в созданный файл на сервере
	_, err = io.Copy(f, file)
	if err != nil {
		log.Printf("Error saving file: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Возвращаем ответ клиенту с идентификатором изображения
	response := map[string]interface{}{
		"filename": fmt.Sprintf("%d_%s", imageID, handler.Filename),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Извлекаем идентификатор изображения из URL
	imagePath := strings.TrimPrefix(r.URL.Path, "/v1/images/")
	filepath := path.Join("uploads", imagePath)

	log.Printf("Попытка доступа к файлу: %s", filepath)

	// Проверяем, существует ли файл
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		log.Printf("Файл не найден: %s", filepath)
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	log.Printf("Отправляем файл клиенту: %s", filepath)
	http.ServeFile(w, r, filepath)
}
