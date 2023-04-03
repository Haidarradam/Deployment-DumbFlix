package filmdto

import "dumbflix/models"

type FilmResponse struct {
	ID            int                     `json:"id"`
	Title         string                  `json:"title"`
	Thumbnailfilm string                  `json:"thumbnailfilm"`
	Year          int                     `json:"year"`
	Category      models.CategoryResponse `json:"category"`
	CategoryID    int                     `json:"category_id"`
	LinkFilm      string                  `json:"linkfilm"`
	Description   string                  `json:"description"`
}
