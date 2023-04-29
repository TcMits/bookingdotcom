package api

import (
	"github.com/TcMits/bookingdotcom/usecase"
	"github.com/labstack/echo/v5"
)

type divisionAPIBuilder struct {}

func newDivisionAPIBuilder() *divisionAPIBuilder {
	return &divisionAPIBuilder{}
}

func (b *divisionAPIBuilder) Build(e *echo.Echo) {
  e.GET("/divisions", b.divisions)
}

//	@Summary		List vietnam divisions
//	@Description	get divisions
//	@Tags			divisions
//	@Accept			json
//	@Produce		json
//	@Param			q	query		usecase.DivisionAPIUseCaseFindConfig	false	"query"
//	@Success		200	{object}	usecase.DivisionAPIUseCaseFindResult
//	@Failure		400	{object}	echo.HTTPError
//	@Failure		404	{object}	echo.HTTPError
//	@Failure		500	{object}	echo.HTTPError
//	@Router			/divisions [get]
func (b *divisionAPIBuilder) divisions(c echo.Context) error {
  uc := usecase.NewDivisionAPIUseCase()

  config := &usecase.DivisionAPIUseCaseFindConfig{}
  if err := c.Bind(config); err != nil {
    return err
  }

  result, err := uc.Find(config)
  if err != nil {
    return echo.NewHTTPError(400, err.Error())
  }

  return c.JSON(200, result)
}
