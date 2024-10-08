// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Получить список песен с пагинацией на основе предоставленных параметров запроса",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Получить библиотеку песен",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Название группы",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Название песни",
                        "name": "song",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Дата выпуска с (формат: 2006-01-02)",
                        "name": "fromDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Дата выпуска до (формат: 2006-01-02)",
                        "name": "untilDate",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Номер страницы",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.PaginatedSongsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.e"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.e"
                        }
                    }
                }
            }
        },
        "/add": {
            "post": {
                "description": "Добавляет новую песню в библиотеку",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Добавить новую песню",
                "parameters": [
                    {
                        "description": "Данные новой песни",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AddSongRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.s"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.e"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.e"
                        }
                    }
                }
            }
        },
        "/lyrics/{id}": {
            "get": {
                "description": "Получить текст конкретной песни по ее ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Получить текст песни",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID песни",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Номер страницы",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.LyricsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.e"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.e"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.e"
                        }
                    }
                }
            }
        },
        "/{id}": {
            "delete": {
                "description": "Удалить песню из библиотеки по ее ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Удалить песню",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID песни",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/dto.s"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.e"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.e"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.e"
                        }
                    }
                }
            },
            "patch": {
                "description": "Обновить данные песни по ее ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Обновить песню",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID песни",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Обновленные данные песни",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateSongRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.s"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.e"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.e"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.e"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AddSongRequest": {
            "type": "object",
            "required": [
                "group",
                "song"
            ],
            "properties": {
                "group": {
                    "description": "Группа или исполнитель",
                    "type": "string",
                    "example": "My Band"
                },
                "song": {
                    "description": "Название песни",
                    "type": "string",
                    "example": "My Song"
                }
            }
        },
        "dto.LyricsResponse": {
            "type": "object",
            "properties": {
                "page": {
                    "description": "Номер текущей страницы",
                    "type": "integer"
                },
                "text": {
                    "description": "Текст песни, разбитый на страницы",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "totalPages": {
                    "description": "Общее количество страниц с текстом",
                    "type": "integer"
                }
            }
        },
        "dto.PaginatedSongsResponse": {
            "type": "object",
            "properties": {
                "page": {
                    "description": "Номер текущей страницы",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "Размер страницы",
                    "type": "integer"
                },
                "songs": {
                    "description": "Список песен",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.SongResponse"
                    }
                },
                "total": {
                    "description": "Общее количество записей",
                    "type": "integer"
                },
                "totalPages": {
                    "description": "Общее количество страниц",
                    "type": "integer"
                }
            }
        },
        "dto.SongResponse": {
            "type": "object",
            "required": [
                "group",
                "id",
                "link",
                "releaseDate",
                "song",
                "text"
            ],
            "properties": {
                "group": {
                    "description": "Название группы или исполнителя",
                    "type": "string"
                },
                "id": {
                    "description": "Идентификатор песни",
                    "type": "integer"
                },
                "link": {
                    "description": "Ссылка на песню",
                    "type": "string"
                },
                "releaseDate": {
                    "description": "Дата релиза песни",
                    "type": "string"
                },
                "song": {
                    "description": "Название песни",
                    "type": "string"
                },
                "text": {
                    "description": "Текст песни",
                    "type": "string"
                }
            }
        },
        "dto.UpdateSongRequest": {
            "type": "object",
            "properties": {
                "group": {
                    "description": "Группа или исполнитель",
                    "type": "string",
                    "example": "My Band"
                },
                "link": {
                    "description": "Ссылка на песню",
                    "type": "string",
                    "example": "https://example.com/mysong"
                },
                "releaseDate": {
                    "description": "Дата релиза",
                    "type": "string",
                    "example": "2023-01-01"
                },
                "text": {
                    "description": "Текст песни",
                    "type": "string",
                    "example": "This is the lyrics..."
                },
                "title": {
                    "description": "Название песни",
                    "type": "string",
                    "example": "My Song"
                }
            }
        },
        "dto.e": {
            "type": "object",
            "properties": {
                "details": {
                    "description": "Дополнительные детали ошибки (опционально)",
                    "type": "string"
                },
                "error": {
                    "description": "Краткое описание ошибки",
                    "type": "string"
                },
                "message": {
                    "description": "Сообщение об ошибке",
                    "type": "string"
                },
                "status": {
                    "description": "HTTP статус ошибки",
                    "type": "integer"
                }
            }
        },
        "dto.s": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "Сообщение об успешном выполнении",
                    "type": "string"
                },
                "status": {
                    "description": "HTTP статус успешного выполнения",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/songs",
	Schemes:          []string{},
	Title:            "Реализация онлайн библиотеки песен",
	Description:      "Выполнялось как тестовое задание для Effective Mobile",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
