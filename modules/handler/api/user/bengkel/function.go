package bengkel

import (
	"math"
	"net/http"
	"strconv"

	cs "github.com/fazaalexander/montirku-be/helper/customstatus"
	bt "github.com/fazaalexander/montirku-be/modules/entity/bengkel"
	"github.com/labstack/echo/v4"
)

func (bh *BengkelHandler) GetAllBengkel() echo.HandlerFunc {
	return func(c echo.Context) error {
		var bengkels []*bt.Bengkel

		pageParam := c.QueryParam("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			page = 1
		}

		pageSize := 10
		offset := (page - 1) * pageSize

		bengkels, total, err := bh.bengkelUseCase.GetAllBengkel(bengkels, offset, pageSize)
		if err != nil {
			code, msg := cs.CustomStatus(err.Error())
			return c.JSON(code, echo.Map{
				"Status":  code,
				"Message": msg,
			})
		}

		if len(bengkels) == 0 {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": "Belum ada list bengkel",
				"Status":  http.StatusNotFound,
			})
		}

		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
		if page > totalPages {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": "Halaman Tidak Ditemukan",
				"Status":  http.StatusNotFound,
			})
		}

		var bengkelResponses []bt.BengkelResponse
		for _, bengkel := range bengkels {
			bengkelResponse := bt.BengkelResponse{
				Category:        bengkel.BengkelCategory.Category,
				Name:            bengkel.Name,
				PhoneNumber:     bengkel.PhoneNumber,
				StartingPrice:   bengkel.StartingPrice,
				BengkelPhotoUrl: bengkel.BengkelPhotoUrl,
				AvgRating:       bengkel.AvgRating,
				Status:          bengkel.Status,
				DistrictName:    bengkel.BengkelAddress.DistrictName,
				CityName:        bengkel.BengkelAddress.CityName,
			}

			bengkelResponses = append(bengkelResponses, bengkelResponse)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Bengkels":  bengkelResponses,
			"Page":      page,
			"TotalPage": totalPages,
			"Status":    http.StatusOK,
		})
	}
}

func (bh *BengkelHandler) GetBengkelById() echo.HandlerFunc {
	return func(c echo.Context) error {
		var bengkel *bt.Bengkel
		bengkelId := c.Param("id")

		bengkel, err := bh.bengkelUseCase.GetBengkelById(bengkelId, bengkel)
		if err != nil {
			code, msg := cs.CustomStatus(err.Error())
			return c.JSON(code, echo.Map{
				"Status":  code,
				"Message": msg,
			})
		}

		var bengkelServices []*bt.BengkelServices
		bengkelServices, err = bh.bengkelUseCase.GetAllBengkelServices(bengkelId, bengkelServices)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "record not found 2",
				"Status":  http.StatusInternalServerError,
			})
		}

		var operationalTimes []*bt.OperationalTime
		operationalTimes, err = bh.bengkelUseCase.GetAllOperationalTime(bengkelId, operationalTimes)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "record not found 3",
				"Status":  http.StatusInternalServerError,
			})
		}

		var bengkelServicesResponse []bt.BengkelServicesResponse
		for _, service := range bengkelServices {
			BengkelService := bt.BengkelServicesResponse{
				Name:  service.Name,
				Price: service.Price,
			}
			bengkelServicesResponse = append(bengkelServicesResponse, BengkelService)
		}

		var operationalTimesResponse []bt.OperationalTimeResponse
		for _, time := range operationalTimes {
			operationalTime := bt.OperationalTimeResponse{
				Day:       time.Day,
				OpenTime:  time.OpenTime,
				CloseTime: time.CloseTime,
			}
			operationalTimesResponse = append(operationalTimesResponse, operationalTime)
		}

		bengkelResponse := bt.BengkelDetailResponse{
			Category:        bengkel.BengkelCategory.Category,
			Name:            bengkel.Name,
			PhoneNumber:     bengkel.PhoneNumber,
			StartingPrice:   bengkel.StartingPrice,
			BengkelPhotoUrl: bengkel.BengkelPhotoUrl,
			AvgRating:       bengkel.AvgRating,
			Status:          bengkel.Status,
			Address: bt.BengkelAddressResponse{
				DistrictName: bengkel.BengkelAddress.DistrictName,
				CityName:     bengkel.BengkelAddress.CityName,
				Address:      bengkel.BengkelAddress.Address,
				Latitude:     bengkel.BengkelAddress.Latitude,
				Longitude:    bengkel.BengkelAddress.Longitude,
			},
			Services:        bengkelServicesResponse,
			OperationalTime: operationalTimesResponse,
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Bengkel": bengkelResponse,
			"Status":  http.StatusOK,
		})
	}
}

func (bh *BengkelHandler) FilterBengkel() echo.HandlerFunc {
	return func(c echo.Context) error {
		var bengkels []*bt.Bengkel

		pageParam := c.QueryParam("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			page = 1
		}

		pageSize := 10
		offset := (page - 1) * pageSize

		category := c.QueryParam("category")
		status := c.QueryParam("status")

		validParams := map[string]bool{"category": true, "status": true, "page": true}
		for param := range c.QueryParams() {
			if !validParams[param] {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"Message": "Masukkan parameter dengan benar",
					"Status":  http.StatusBadRequest,
				})
			}
		}

		bengkels, total, err := bh.bengkelUseCase.FilterBengkel(category, status, offset, pageSize)
		if err != nil {
			code, msg := cs.CustomStatus(err.Error())
			return c.JSON(code, echo.Map{
				"Status":  code,
				"Message": msg,
			})
		}

		if len(bengkels) == 0 {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": "Bengkel tidak ditemukan",
				"Status":  http.StatusNotFound,
			})
		} else {
			var bengkelResponses []bt.BengkelResponse
			for _, bengkel := range bengkels {
				bengkelResponse := bt.BengkelResponse{
					Category:        bengkel.BengkelCategory.Category,
					Name:            bengkel.Name,
					PhoneNumber:     bengkel.PhoneNumber,
					StartingPrice:   bengkel.StartingPrice,
					BengkelPhotoUrl: bengkel.BengkelPhotoUrl,
					AvgRating:       bengkel.AvgRating,
					Status:          bengkel.Status,
					DistrictName:    bengkel.BengkelAddress.DistrictName,
					CityName:        bengkel.BengkelAddress.CityName,
				}

				bengkelResponses = append(bengkelResponses, bengkelResponse)
			}
			return c.JSON(http.StatusOK, map[string]interface{}{
				"Bengkels":  bengkelResponses,
				"Page":      page,
				"TotalPage": int(math.Ceil(float64(total) / float64(pageSize))),
				"Status":    http.StatusOK,
			})
		}
	}
}
