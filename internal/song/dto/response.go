package dto

import (
	"encoding/json"
	"net/http"
)

// SongResponse представляет ответ с информацией о песне.
type SongResponse struct {
	ID          uint   `json:"id" validate:"required,number"`                       // Идентификатор песни
	Group       string `json:"group" validate:"required"`                           // Название группы или исполнителя
	Song        string `json:"song" validate:"required"`                            // Название песни
	ReleaseDate string `json:"releaseDate" validate:"required,datetime=2006.01.02"` // Дата релиза песни
	Text        string `json:"text" validate:"required"`                            // Текст песни
	Link        string `json:"link" validate:"required"`                            // Ссылка на песню
}

// PaginatedSongsResponse представляет ответ с пагинацией списка песен.
type PaginatedSongsResponse struct {
	Songs      []SongResponse `json:"songs"`      // Список песен
	Total      int64          `json:"total"`      // Общее количество записей
	Page       int            `json:"page"`       // Номер текущей страницы
	TotalPages int            `json:"totalPages"` // Общее количество страниц
	PageSize   int            `json:"pageSize"`   // Размер страницы
}

// LyricsResponse представляет ответ с текстом песни и информацией о пагинации.
type LyricsResponse struct {
	Text       []string `json:"text"`       // Текст песни, разбитый на страницы
	Page       int      `json:"page"`       // Номер текущей страницы
	TotalPages int      `json:"totalPages"` // Общее количество страниц с текстом
}

// e представляет структуру для передачи информации об ошибке.
type e struct {
	Status  int    `json:"status"`            // HTTP статус ошибки
	Error   string `json:"error"`             // Краткое описание ошибки
	Message string `json:"message"`           // Сообщение об ошибке
	Details string `json:"details,omitempty"` // Дополнительные детали ошибки (опционально)
}

// s представляет структуру для успешного ответа.
type s struct {
	Status  int    `json:"status"`  // HTTP статус успешного выполнения
	Message string `json:"message"` // Сообщение об успешном выполнении
}

// Success возвращает сообщение об успешном выполнении операции клиенту в формате JSON.
func Success(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := s{
		Status:  status,
		Message: msg,
	}
	json.NewEncoder(w).Encode(response)
}

// Error возвращает сообщение об ошибке клиенту в формате JSON.
func Error(w http.ResponseWriter, status int, errMsg string, details ...string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errorResponse := e{
		Status:  status,
		Error:   http.StatusText(status),
		Message: errMsg,
	}
	if len(details) > 0 {
		errorResponse.Details = details[0]
	}
	json.NewEncoder(w).Encode(errorResponse)
}
