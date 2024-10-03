package endpoints

import (
	"encoding/json"
	"errors"
	"github.com/CracherX/Song-service/internal/song/client"
	"github.com/CracherX/Song-service/internal/song/dto"
	mw "github.com/CracherX/Song-service/internal/song/middleware"
	"github.com/CracherX/Song-service/internal/song/services"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// SongsEndpoint структура точек входа API для работы с библиотекой песен.
type SongsEndpoint struct {
	Service *services.SongsService
	Client  *client.ApiClient
	Logger  *zap.Logger
}

// NewSongsEndpoint конструктор для SongsEndpoint.
func NewSongsEndpoint(srv *services.SongsService, clnt *client.ApiClient, log *zap.Logger) *SongsEndpoint {
	return &SongsEndpoint{
		Service: srv,
		Client:  clnt,
		Logger:  log,
	}
}

// @Summary Получить библиотеку песен
// @Description Получить список песен с пагинацией на основе предоставленных параметров запроса
// @Tags songs
// @Accept json
// @Produce json
// @Param group query string false "Название группы"
// @Param song query string false "Название песни"
// @Param fromDate query string false "Дата выпуска с (формат: 2006-01-02)"
// @Param untilDate query string false "Дата выпуска до (формат: 2006-01-02)"
// @Param page query int false "Номер страницы" default(1)
// @Success 200 {object} dto.PaginatedSongsResponse
// @Failure 400 {object} dto.e
// @Failure 500 {object} dto.e
// @Router / [GET]
func (ep *SongsEndpoint) Library(w http.ResponseWriter, r *http.Request) {
	ep.Logger.Debug("Начало выполнение запроса", zap.String("Функция", "Library"))

	var req dto.SongsRequest

	err := schema.NewDecoder().Decode(&req, r.URL.Query())
	if err != nil {
		dto.Error(w, http.StatusBadRequest, "Bad Request", "Неверные параметры запроса")
		ep.Logger.Debug("Bad Request", zap.String("Функция", "Library"))
		return
	}

	val := mw.GetValidator(r.Context())
	err = val.Struct(&req)
	if err != nil {
		dto.Error(w, http.StatusBadRequest, "Bad Request", "Неверный формат данных указанный в параметрах запроса")
		ep.Logger.Debug("Bad Request", zap.String("Функция", "Library"))
		return
	}

	songs, trc, err := ep.Service.GetLibrary(&req)
	if err != nil {
		dto.Error(w, http.StatusInternalServerError, "Ошибка на стороне сервера")
		ep.Logger.Error("Ошибка выполнения запроса", zap.String("Ошибка", err.Error()))
		return
	}

	// Преобразование данных из модели в DTO
	songResponses := make([]dto.SongResponse, len(songs))
	for i, song := range songs {
		songResponses[i] = dto.SongResponse{
			ID:          song.ID,
			Group:       song.Group,
			Song:        song.Song,
			ReleaseDate: song.ReleaseDate,
			Text:        song.Text,
			Link:        song.Link,
		}
	}

	// Рассчитываем общее количество страниц
	pageSize := 10
	totalPages := int((trc + int64(pageSize) - 1) / int64(pageSize))

	// Формируем ответ
	response := dto.PaginatedSongsResponse{
		Songs:      songResponses,
		Total:      trc,
		Page:       req.Page,
		TotalPages: totalPages,
		PageSize:   pageSize,
	}

	// Устанавливаем заголовок ответа
	w.Header().Set("Content-Type", "application/json")

	// Кодируем результат в JSON и отправляем ответ
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		dto.Error(w, http.StatusInternalServerError, "Ошибка на стороне сервера")
		ep.Logger.Error("Ошибка выполнения запроса.", zap.String("Ошибка", err.Error()))
		return
	}
	ep.Logger.Debug("Выполнение запроса завершено", zap.String("Функция", "Library"))
}

// @Summary Получить текст песни
// @Description Получить текст конкретной песни по ее ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param page query int false "Номер страницы" default(1)
// @Success 200 {object} dto.LyricsResponse
// @Failure 400 {object} dto.e
// @Failure 404 {object} dto.e
// @Failure 500 {object} dto.e
// @Router /lyrics/{id} [get]
func (ep *SongsEndpoint) Text(w http.ResponseWriter, r *http.Request) {
	ep.Logger.Debug("Начало выполнения запроса", zap.String("Функция", "Text"))
	var req dto.LyricsRequest

	err := schema.NewDecoder().Decode(&req, r.URL.Query())
	if err != nil {
		dto.Error(w, http.StatusBadRequest, "Bad Request", "Неверные параметры запроса")
		return
	}

	id := mux.Vars(r)
	req.ID, _ = strconv.Atoi(id["id"])

	val := mw.GetValidator(r.Context())
	err = val.Struct(&req)
	if err != nil {
		dto.Error(w, http.StatusBadRequest, "Bad Request", "Неверный формат данных указанный в параметрах запроса")
		ep.Logger.Debug("Bad Request", zap.String("Функция", "Text"))
		return
	}

	lyrics, totalPages, err := ep.Service.GetLyrics(&req)
	if err != nil {
		dto.Error(w, http.StatusInternalServerError, "Ошибка на стороне сервера")
		ep.Logger.Error("Ошибка выполнения запроса", zap.String("Ошибка", err.Error()))
		return
	}

	response := dto.LyricsResponse{
		Text:       lyrics,
		Page:       req.Page,
		TotalPages: totalPages,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		dto.Error(w, http.StatusInternalServerError, "Ошибка на стороне сервера")
		ep.Logger.Error("Ошибка кодирования ответа", zap.String("Ошибка", err.Error()))
	}
	ep.Logger.Debug("Выполнение запроса завершено", zap.String("Функция", "Text"))
}

// @Summary Удалить песню
// @Description Удалить песню из библиотеки по ее ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Success 204 {object} dto.s
// @Failure 400 {object} dto.e
// @Failure 404 {object} dto.e
// @Failure 500 {object} dto.e
// @Router /{id} [delete]
func (ep *SongsEndpoint) Delete(w http.ResponseWriter, r *http.Request) {
	ep.Logger.Debug("Начало выполнения запроса", zap.String("Функция", "DeleteSong"))
	var req dto.DeleteSongRequest

	vars := mux.Vars(r)
	req.ID, _ = strconv.Atoi(vars["id"])

	val := mw.GetValidator(r.Context())
	err := val.Struct(&req)
	if err != nil {
		dto.Error(w, http.StatusBadRequest, "Bad Request", "Указано невалидное значение ID")
		ep.Logger.Debug("Bad Request")
		return
	}
	err = ep.Service.DeleteSong(&req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			dto.Error(w, http.StatusNotFound, "Песня с таким ID не найдена")
			ep.Logger.Debug("Not Found", zap.Int("songID", req.ID))
		} else {
			dto.Error(w, http.StatusInternalServerError, "Ошибка на стороне сервера")
			ep.Logger.Error("Ошибка удаления песни", zap.Error(err))
		}
		return
	}
	dto.Success(w, http.StatusNoContent, "Успешное удаление")
	ep.Logger.Debug("Выполнение запроса завершено", zap.String("Функция", "DeleteSong"))
}

// @Summary Обновить песню
// @Description Обновить данные песни по ее ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param song body dto.UpdateSongRequest true "Обновленные данные песни"
// @Success 200 {object} dto.s
// @Failure 400 {object} dto.e
// @Failure 404 {object} dto.e
// @Failure 500 {object} dto.e
// @Router /{id} [patch]
func (ep *SongsEndpoint) UpdateSong(w http.ResponseWriter, r *http.Request) {
	ep.Logger.Debug("Начало выполнения запроса", zap.String("Функция", "UpdateSong"))
	var req dto.UpdateSongRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		dto.Error(w, http.StatusBadRequest, "Bad Request", "Неверные параметры запроса")
		ep.Logger.Debug("Bad Request", zap.String("Функция", "UpdateSong"))
		return
	}
	vars := mux.Vars(r)
	req.ID, _ = strconv.Atoi(vars["id"])

	val := mw.GetValidator(r.Context())
	err = val.Struct(&req)
	if err != nil {
		dto.Error(w, http.StatusBadRequest, "Bad Request", "Неверные параметры запроса")
		ep.Logger.Debug("Bad Request", zap.String("Функция", "UpdateSong"))
		return
	}

	// Вызываем сервис для обновления песни
	err = ep.Service.UpdateSong(&req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			dto.Error(w, http.StatusNotFound, "Not Found", "Песня не найдена")
			ep.Logger.Debug("Not Found", zap.String("Функция", "UpdateSong"))
		} else {
			dto.Error(w, http.StatusInternalServerError, "Ошибка на стороне сервера")
			ep.Logger.Error("Ошибка изменения данных песни", zap.String("Ошибка", err.Error()))
		}
		return
	}

	// Отправляем успешный ответ
	dto.Success(w, http.StatusOK, "Песня успешно обновлена")
	ep.Logger.Debug("Выполнение запроса завершено", zap.String("Функция", "UpdateSong"))
}

// @Summary Добавить новую песню
// @Description Добавляет новую песню в библиотеку
// @Tags songs
// @Accept json
// @Produce json
// @Param song body dto.AddSongRequest true "Данные новой песни"
// @Success 201 {object} dto.s
// @Failure 400 {object} dto.e
// @Failure 500 {object} dto.e
// @Router /add [post]
func (ep *SongsEndpoint) AddSong(w http.ResponseWriter, r *http.Request) {
	ep.Logger.Debug("Начало выполнения запроса", zap.String("Функция", "AddSong"))
	var req dto.AddSongRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		dto.Error(w, http.StatusBadRequest, "Bad Request", "Неверные параметры запроса")
		ep.Logger.Debug("Bad Request", zap.String("Функция", "AddSong"))
		return
	}

	val := mw.GetValidator(r.Context())
	err = val.Struct(&req)
	if err != nil {
		dto.Error(w, http.StatusBadRequest, "Bad Request", "Неверные параметры запроса")
		ep.Logger.Debug("Bad Request", zap.String("Функция", "AddSong"))
		return
	}

	params := map[string]string{
		"group": req.Group,
		"song":  req.Song,
	}

	cres, err := ep.Client.Get("/info", params)
	if err != nil {
		dto.Error(w, http.StatusBadGateway, "Bad Gateway", "Проблема в работе с внешним сервисом")
		ep.Logger.Error("Bad Gateway", zap.String("Ошибка", err.Error()))
		return
	}
	var cdto dto.SongResponse
	err = json.NewDecoder(cres.Body).Decode(&cdto)
	if err != nil {
		dto.Error(w, http.StatusBadGateway, "Bad Gateway", "Проблема в работе с внешним сервисом")
		ep.Logger.Error("Bad Gateway", zap.String("Ошибка", err.Error()))
		return
	}
	err = val.Struct(&cdto)
	if err != nil {
		dto.Error(w, http.StatusBadGateway, "Bad Gateway", "Проблема в работе с внешним сервисом")
		ep.Logger.Error("Bad Gateway", zap.String("Ошибка", err.Error()))
		return
	}
	err = ep.Service.AddSong(&cdto)
	if err != nil {
		dto.Error(w, http.StatusInternalServerError, "Ошибка на стороне сервера")
		ep.Logger.Error("Ошибка сохранения новой песни в базу данных", zap.String("Ошибка", err.Error()))
		return
	}
	dto.Success(w, http.StatusCreated, "Новая песня успешно добавлена")
	ep.Logger.Debug("Выполнение запроса завершено", zap.String("Функция", "AddSong"))
}
