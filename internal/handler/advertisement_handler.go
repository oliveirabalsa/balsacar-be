package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/oliveirabalsa/balsacar-be/internal/dto"
	"github.com/oliveirabalsa/balsacar-be/internal/entity"
	"github.com/oliveirabalsa/balsacar-be/internal/service"
	"github.com/xuri/excelize/v2"
)

type AdvertisementHandler struct {
	advertisementService service.AdvertisementService
}

func NewAdvertisementHandler(advertisementService service.AdvertisementService) *AdvertisementHandler {
	return &AdvertisementHandler{
		advertisementService: advertisementService,
	}
}

func (h *AdvertisementHandler) CreateAdvertisementHandler(c *gin.Context) {
	var advertisement entity.Advertisement
	if err := c.ShouldBindJSON(&advertisement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAdvertisement, err := h.advertisementService.CreateAdvertisement(&advertisement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdAdvertisement)
}

func (h *AdvertisementHandler) GetAdvertisementByIDHandler(c *gin.Context) {
	advertisementIDStr := c.Param("id")
	advertisementID, err := uuid.Parse(advertisementIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	advertisement, err := h.advertisementService.GetAdvertisementByID(advertisementID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, advertisement)
}

func (h *AdvertisementHandler) GetAllAdvertisementsHandler(c *gin.Context) {
	fmt.Println("GetAllAdvertisementsHandler")
	filters := dto.AdvertisementParamsDto{
		Page:     1,
		PageSize: 10,
	}

	fmt.Println("filters", filters)

	if err := c.ShouldBindQuery(&filters); err != nil {
		fmt.Println("err filters", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	advertisements, total, err := h.advertisementService.GetAllAdvertisements(&filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(advertisements) < 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	response := dto.Paginated{
		Data:       advertisements,
		Page:       filters.Page,
		PageSize:   filters.PageSize,
		Total:      total,
		TotalPages: (total + filters.PageSize - 1) / filters.PageSize,
	}

	c.JSON(http.StatusOK, response)
}

func (h *AdvertisementHandler) UpdateAdvertisementHandler(c *gin.Context) {
	advertisementIDStr := c.Param("id")
	advertisementID, err := uuid.Parse(advertisementIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var advertisement entity.Advertisement
	if err := c.ShouldBindJSON(&advertisement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAdvertisement, err := h.advertisementService.UpdateAdvertisement(advertisementID, &advertisement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedAdvertisement)
}

func (h *AdvertisementHandler) DeleteAdvertisementHandler(c *gin.Context) {
	advertisementIDStr := c.Param("id")
	advertisementID, err := uuid.Parse(advertisementIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.advertisementService.DeleteAdvertisement(advertisementID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *AdvertisementHandler) UploadSheetAdvertisementHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	xlsxFile, err := excelize.OpenReader(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sheetName := xlsxFile.GetSheetName(0)
	rows, err := xlsxFile.GetRows(sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	headerRow := rows[0]
	jsonData := []entity.Advertisement{}
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		var advertisement entity.Advertisement
		for j, key := range headerRow {
			value := row[j]
			advertisement.FromKeyValue(&advertisement, key, value)
		}
		h.advertisementService.CreateAdvertisement(&advertisement)
		jsonData = append(jsonData, advertisement)
	}
	c.JSON(http.StatusOK, jsonData)
}
