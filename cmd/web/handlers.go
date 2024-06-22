package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Функция-обработчик главной страницы
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Проверка, совпадает ли текущий путь URL-запроса с шаблоном "/"
	// Если не совпадает, ф-ция http.NotFound возвращает клиенту ошибку 404
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// Чтение шаблона и возвращение ошибки в случае проблемы
	tpf, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Execute для записи содержимого шаблона в тело HTTP ответа
	err = tpf.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

// Функция для создания заметки
func (app *application) createSnip(w http.ResponseWriter, r *http.Request) {
	// Проверка запроса, если это не POST-запрос, то у клиента возвращает ошибку 405
	// "return" в конце в случае ошибки завершает работу функции
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "GET/PUT/DELETE-Методы запрещены", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Создание заметки..."))
}

// Функция для отображения заметки
func (app *application) showSnip(w http.ResponseWriter, r *http.Request) {
	// Извлечение значения параметра id из URL
	// конвертация id в число
	// если id меньше или равно нулю, то ошибка 404
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Чтение заметки с ID %d...", id)
}
