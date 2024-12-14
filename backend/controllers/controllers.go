package controllers

import (
	"log"
	"net/http"
	"news_app/models"
	"news_app/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type NewsController struct {
	service  service.NewsService
	validate *validator.Validate
}

func NewNewsController(service service.NewsService) *NewsController {
	return &NewsController{
		service:  service,
		validate: validator.New(),
	}
}

// @Summary Get all news
// @Description Get all news articles
// @Tags news
// @Accept  json
// @Produce  json
// @Success 200 {array} models.News
// @Router /api/news [get]
func (c *NewsController) GetNews(ctx *gin.Context) {
	news, err := c.service.GetAllNews(ctx)
	if err != nil {
		log.Printf("Ошибка получении новости: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"ошибка": "Не удалось получить новости"})
		return
	}
	ctx.JSON(http.StatusOK, news)
}

// @Summary Get a news article by ID
// @Description Get a single news article by its ID
// @Tags news
// @Accept  json
// @Produce  json
// @Param id path int true "News ID"
// @Success 200 {object} models.News
// @Failure 404 {object} gin.H{"error": "News not found"}
// @Router /api/news/{id} [get]
func (c *NewsController) GetNewsByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"ошибка": "Неправильный формат ID"})
		return
	}

	news, err := c.service.GetNewsByID(ctx, id)
	if err != nil {
		if err.Error() == "новость не найдена" {
			ctx.JSON(http.StatusNotFound, gin.H{"ошибка": "Новость не найдена"})
		} else {
			log.Printf("Ошибка получения новости по ID: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"ошибка": "Не удалось получить новости"})
		}
		return
	}

	ctx.JSON(http.StatusOK, news)
}

// @Summary Create a news article
// @Description Create a new news article
// @Tags news
// @Accept  json
// @Produce  json
// @Param news body models.News true "News Data"
// @Success 201 {object} models.News
// @Router /api/news/create [post]
func (c *NewsController) CreateNews(ctx *gin.Context) {
	var news models.News
	if err := ctx.ShouldBindJSON(&news); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ошибка": "Некорректный запрос"})
		return
	}

	if err := c.validate.Struct(news); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errors := make(map[string]string)
			for _, vErr := range validationErrors {
				errors[vErr.Field()] = vErr.ActualTag()
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"ошибка": "Проверка не удалась", "детали": errors})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"ошибка": "Некорректный запрос"})
		return
	}

	createdNews, err := c.service.CreateNews(ctx, news)
	if err != nil {
		log.Printf("Ошибка создания новости: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"ошибка": "Не удалось создать новость"})
		return
	}

	ctx.JSON(http.StatusCreated, createdNews)
}

// @Summary Update an existing news article
// @Description Update an existing news article by ID
// @Tags news
// @Accept  json
// @Produce  json
// @Param id path int true "News ID"
// @Param news body models.News true "Updated News Data"
// @Success 200 {object} models.News
// @Router /api/news/update/{id} [put]
func (c *NewsController) UpdateNews(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"ошибка": "Неправильный формат ID"})
		return
	}

	var news models.News
	if err := ctx.ShouldBindJSON(&news); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ошибка": "Некорректный запрос"})
		return
	}

	if err := c.validate.Struct(news); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ошибка": "Проверка не удалась", "детали": err.Error()})
		return
	}

	updatedNews, err := c.service.UpdateNews(ctx, id, news)
	if err != nil {
		log.Printf("Ошибка обновления новости: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"ошибка": "Не удалось обновить новость"})
		return
	}

	ctx.JSON(http.StatusOK, updatedNews)
}

// @Summary Delete a news article
// @Description Delete a news article by ID
// @Tags news
// @Param id path int true "News ID"
// @Success 200 {string} string "News deleted"
// @Router /api/news/delete/{id} [delete]
func (c *NewsController) DeleteNews(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"ошибка": "Неправильный формат ID"})
		return
	}

	err = c.service.DeleteNews(ctx, id)
	if err != nil {
		log.Printf("Ошибка удаления новости: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"ошибка": "Не удалось удалить новость"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"уведомление": "Новость удалена"})
}
