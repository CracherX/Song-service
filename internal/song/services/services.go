package services

import (
	"github.com/CracherX/Song-service/internal/song/dto"
	"github.com/CracherX/Song-service/internal/song/storage/models"
	"gorm.io/gorm"
	"strings"
	"time"
)

// SongsService сервис для работы с библиотекой песен
type SongsService struct {
	DB *gorm.DB
}

// NewSongsService конструктор для SongsService.
func NewSongsService(db *gorm.DB) *SongsService {
	return &SongsService{
		DB: db,
	}
}

// GetLibrary получает список песен из библиотеки
func (ss *SongsService) GetLibrary(dto *dto.SongsRequest) ([]models.Song, int64, error) {

	var songs []models.Song
	var totalRecords int64

	query := ss.DB.Model(&models.Song{})

	if dto.Group != "" {
		query = query.Where(`"group" LIKE ?`, "%"+dto.Group+"%")
	}

	if dto.Song != "" {
		query = query.Where("title LIKE ?", "%"+dto.Song+"%")
	}

	if dto.FromDate != "" && dto.UntilDate != "" {
		fromDate, err := time.Parse("2006-01-02", dto.FromDate)
		if err != nil {
			return nil, 0, err
		}

		untilDate, err := time.Parse("2006-01-02", dto.UntilDate)
		if err != nil {
			return nil, 0, err
		}

		query = query.Where("release_date BETWEEN ? AND ?", fromDate, untilDate)
	}

	pageSize := 10
	offset := (dto.Page - 1) * pageSize

	query.Count(&totalRecords)

	err := query.Limit(pageSize).Offset(offset).Find(&songs).Error
	if err != nil {
		return nil, 0, err
	}

	return songs, totalRecords, nil
}

// GetLyrics получает текст песни по ID
func (ss *SongsService) GetLyrics(dto *dto.LyricsRequest) ([]string, int, error) {
	var song models.Song
	if err := ss.DB.First(&song, dto.ID).Error; err != nil {
		return nil, 0, err
	}

	verses := strings.Split(song.Text, "\\n")

	pageSize := 4
	totalVerses := len(verses)
	totalPages := (totalVerses + pageSize - 1) / pageSize

	startIndex := (dto.Page - 1) * pageSize
	endIndex := startIndex + pageSize

	if endIndex > totalVerses {
		endIndex = totalVerses
	}

	return verses[startIndex:endIndex], totalPages, nil
}

// DeleteSong удаляет песню по ID
func (ss *SongsService) DeleteSong(dto *dto.DeleteSongRequest) error {
	var song models.Song
	if err := ss.DB.First(&song, dto.ID).Error; err != nil {
		return err
	}

	if err := ss.DB.Delete(&song).Error; err != nil {
		return err
	}

	return nil
}

// UpdateSong обновляет данные песни по ID
func (ss *SongsService) UpdateSong(dto *dto.UpdateSongRequest) error {
	var song models.Song

	if err := ss.DB.First(&song, dto.ID).Error; err != nil {
		return err
	}

	if dto.Group != nil {
		song.Group = *dto.Group
	}
	if dto.Title != nil {
		song.Song = *dto.Title
	}
	if dto.ReleaseDate != nil {
		song.ReleaseDate = *dto.ReleaseDate
	}
	if dto.Text != nil {
		song.Text = *dto.Text
	}
	if dto.Link != nil {
		song.Link = *dto.Link
	}

	return ss.DB.Save(&song).Error
}

// AddSong добавляет новую песню в библиотеку по ID
func (ss *SongsService) AddSong(dto *dto.SongResponse) error {
	song := models.Song{
		Group:       dto.Group,
		Song:        dto.Song,
		ReleaseDate: dto.ReleaseDate,
		Text:        dto.Text,
		Link:        dto.Link,
	}

	if err := ss.DB.Create(&song).Error; err != nil {
		return err
	}
	return nil
}
