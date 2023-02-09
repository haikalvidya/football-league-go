package delivery

import (
	"net/http"

	"github.com/haikalvidya/football-league/internal/delivery/payload"
	"github.com/haikalvidya/football-league/pkg/common"
	"github.com/haikalvidya/football-league/pkg/utils"
	"github.com/labstack/echo/v4"
)

type containLetterDelivery deliveryType

func (d *containLetterDelivery) ContainLetter(c echo.Context) error {
	res := common.Response{}
	req := &payload.ContainLetterRequest{}

	c.Bind(req)

	if err := c.Validate(req); err != nil {
		res.Error = utils.GetErrorValidation(err)
		res.Status = false
		res.Message = "Failed Contain Letter"
		return c.JSON(http.StatusBadRequest, res)
	}

	result := d.Usecase.ContainLetter.ContainLetter(req)

	res.Message = "Success Contain Letter"
	res.Data = result
	res.Status = true

	return c.JSON(http.StatusOK, res)
}
