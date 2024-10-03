package dto

// SongsRequest представляет параметры запроса для фильтрации песен.
type SongsRequest struct {
	Group     string `schema:"group"`                                                                  // Группа или исполнитель
	Song      string `schema:"song"`                                                                   // Название песни
	FromDate  string `schema:"fromDate" validate:"omitempty,datetime=2006.01.02" example:"2023-01-01"` // Дата начала периода
	UntilDate string `schema:"untDate" validate:"omitempty,datetime=2006.01.02" example:"2023-12-31"`  // Дата конца периода
	Page      int    `schema:"page,default:1" validate:"number,gte=1" example:"1"`                     // Номер страницы для пагинации
}

// LyricsRequest представляет запрос для получения текста песни.
type LyricsRequest struct {
	ID   int `validate:"required,number,gte=1" example:"123"`              // Идентификатор песни
	Page int `schema:"page,default:1" validate:"number,gte=1" example:"1"` // Номер страницы для пагинации текста
}

// DeleteSongRequest представляет запрос для удаления песни.
type DeleteSongRequest struct {
	ID int `validate:"required,number,gte=1" example:"123"` // Идентификатор удаляемой песни
}

// UpdateSongRequest представляет запрос для обновления данных песни.
type UpdateSongRequest struct {
	ID          int     `json:"id" validate:"number,gte=1,required" swaggerignore:"true" example:"123"`       // Идентификатор песни
	Group       *string `json:"group" validate:"omitempty" example:"My Band"`                                 // Группа или исполнитель
	Title       *string `json:"title" validate:"omitempty" example:"My Song"`                                 // Название песни
	ReleaseDate *string `json:"releaseDate" validate:"omitempty,datetime=2006.01.02" example:"2023-01-01"`    // Дата релиза
	Text        *string `json:"text,omitempty" validate:"omitempty" example:"This is the lyrics..."`          // Текст песни
	Link        *string `json:"link,omitempty" validate:"omitempty,url" example:"https://example.com/mysong"` // Ссылка на песню
}

// AddSongRequest представляет запрос для добавления новой песни.
type AddSongRequest struct {
	Group string `json:"group" validate:"required" example:"My Band"` // Группа или исполнитель
	Song  string `json:"song"  validate:"required" example:"My Song"` // Название песни
}
