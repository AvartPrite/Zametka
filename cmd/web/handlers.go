package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AvartPrite/Zametka/pkg/models"
)

// Функция-обработчик главной страницы
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Проверка, совпадает ли текущий путь URL-запроса с шаблоном "/"
	// Если не совпадает, ф-ция http.NotFound возвращает клиенту ошибку 404
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})
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

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})
}

// Функция для создания заметки
func (app *application) createSnip(w http.ResponseWriter, r *http.Request) {
	// Проверка запроса, если это не POST-запрос, то у клиента возвращает ошибку 405
	// "return" в конце в случае ошибки завершает работу функции
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "Рецепт пельменей"
	content := "Берём и покупаем пельмени"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snip?id=%d", id), http.StatusSeeOther)
}
