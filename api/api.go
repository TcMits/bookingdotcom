package api

import (
	"github.com/TcMits/bookingdotcom"
	"github.com/labstack/echo/v5"
)

type APIBuilder struct {
	bookingdotcom bookingdotcom.BookingDotCom
}

func NewAPIBuilder(bookingdotcom bookingdotcom.BookingDotCom) *APIBuilder {
	return &APIBuilder{
		bookingdotcom: bookingdotcom,
	}
}

//	@title			Swagger bookingdotcom API
//	@version		1.0
//	@description	This a course project at HCMUS.
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:8080
//	@BasePath		/
func (b *APIBuilder) Build(e *echo.Echo) {
	newStayAPIBuilder(b.bookingdotcom).Build(e)
  newDivisionAPIBuilder().Build(e)
  newDocsAPIBuilder().Build(e)
}
