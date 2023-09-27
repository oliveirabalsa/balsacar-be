package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	advertisements, err := h.advertisementService.GetAllAdvertisements()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(advertisements) < 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	c.JSON(http.StatusOK, advertisements)
}

func (h *AdvertisementHandler) UpdateAdvertisementHandler(c *gin.Context) {
	var advertisement entity.Advertisement
	if err := c.ShouldBindJSON(&advertisement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAdvertisement, err := h.advertisementService.UpdateAdvertisement(&advertisement)
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

        jsonData := make([]map[string]interface{}, 0)

				for _, sheetName := range xlsxFile.GetSheetMap() {
					rows, err := xlsxFile.GetRows(sheetName)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
						return
					}

					for _, row := range rows {
						rowData := make(map[string]interface{})
						for _, cellValue := range row {
							rowData[cellValue] = cellValue
						}
						jsonData = append(jsonData, rowData)
					}
				}


	c.JSON(http.StatusOK, jsonData)

}
