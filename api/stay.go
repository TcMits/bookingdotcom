package api

import (
	"github.com/TcMits/bookingdotcom"
	"github.com/TcMits/bookingdotcom/usecase"
	"github.com/labstack/echo/v5"
)

type stayAPIBuilder struct {
	bookingdotcom bookingdotcom.BookingDotCom
}

func newStayAPIBuilder(bookingdotcom bookingdotcom.BookingDotCom) *stayAPIBuilder {
	return &stayAPIBuilder{bookingdotcom: bookingdotcom}
}

func (b *stayAPIBuilder) Build(e *echo.Echo) {
	e.GET("/stays", b.stays)
	e.GET("/stays/:id", b.stay)
	e.PATCH("/stays/:id/reserve", b.reserve)
}

//	@Summary		List stays
//	@Description	get stays example provinceCode=79&districtCode=770&wardCode=27127&guests=2&guests=0&checkTimes=2024-01-31T23:34:25.191Z&checkTimes=2024-09-21T06:10:36.275Z
//	@Tags			stays
//	@Accept			json
//	@Produce		json
//	@Param			q	query		usecase.StayAPIUseCaseFindStaysQuery	false	"query"
//	@Success		200	{object}	usecase.StayAPIUseCaseFindStaysResult
//	@Failure		400	{object}	echo.HTTPError
//	@Failure		404	{object}	echo.HTTPError
//	@Failure		500	{object}	echo.HTTPError
//	@Router			/stays [get]
func (b *stayAPIBuilder) stays(c echo.Context) error {
	uc := usecase.NewStayAPIUseCase(b.bookingdotcom)

	config := &usecase.StayAPIUseCaseFindStaysConfig{}
	if err := c.Bind(config); err != nil {
		return err
	}

	result, err := uc.FindStays(c.Request().Context(), config)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, result)
}

//	@Summary		Show a stay
//	@Description	get stay
//	@Tags			stays
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string									true	"id"
//	@Param			q	query		usecase.StayAPIUseCaseFindStaysQuery	false	"query"
//	@Success		200	{object}	usecase.StayAPIUseCaseReserveRoomResult
//	@Failure		400	{object}	echo.HTTPError
//	@Failure		404	{object}	echo.HTTPError
//	@Failure		500	{object}	echo.HTTPError
//	@Router			/stays/{id} [get]
func (b *stayAPIBuilder) stay(c echo.Context) error {
	uc := usecase.NewStayAPIUseCase(b.bookingdotcom)

	config := &usecase.StayAPIUseCaseFindStaysConfig{}
	if err := c.Bind(config); err != nil {
		return err
	}

	result, err := uc.FindStays(c.Request().Context(), config)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	if result.Count == 0 {
		return echo.NewHTTPError(404, "Not Found")
	}

	return c.JSON(200, result.Items[0])
}

//	@Summary		Reserve a room in a stay
//	@Description	reserve room
//	@Tags			stays
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string									true	"id"
//	@Param			payload	body		usecase.StayAPIUseCaseReserveRoomBody	true	"payload"
//	@Success		200		{object}	usecase.StayAPIUseCaseReserveRoomResult
//	@Failure		400		{object}	echo.HTTPError
//	@Failure		404		{object}	echo.HTTPError
//	@Failure		500		{object}	echo.HTTPError
//	@Router			/stays/{id}/reserve [patch]
func (b *stayAPIBuilder) reserve(c echo.Context) error {
	uc := usecase.NewStayAPIUseCase(b.bookingdotcom)

	config := &usecase.StayAPIUseCaseReserveRoomConfig{}
	if err := c.Bind(config); err != nil {
		return err
	}

	result, err := uc.ReserveRoom(c.Request().Context(), config)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, result)
}
