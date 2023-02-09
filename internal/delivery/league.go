package delivery

import (
	"io"
	"net/http"

	"github.com/haikalvidya/football-league/pkg/common"
	"github.com/labstack/echo/v4"
)

type leagueDelivery deliveryType

func (d *leagueDelivery) RecordGame(c echo.Context) error {
	res := common.Response{}

	data, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	req := string(data)

	err = d.Usecase.League.RecordGame(req)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Message = "Success Record Game"
	res.Status = true

	return c.JSON(http.StatusOK, res)
}

func (d *leagueDelivery) LeagueStanding(c echo.Context) error {
	res := common.Response{}

	leagueStanding, err := d.Usecase.League.LeagueStanding()
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Message = "Success Get League Standing"
	res.Data = leagueStanding
	res.Status = true

	return c.JSON(http.StatusOK, res)
}

func (d *leagueDelivery) ClubStanding(c echo.Context) error {
	res := common.Response{}
	// get parameter clubname if exist
	clubName := c.QueryParam("clubname")

	clubStanding, err := d.Usecase.League.ClubStanding(clubName)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Message = "Success Get Club Standing"
	res.Data = clubStanding
	res.Status = true

	return c.JSON(http.StatusOK, res)
}
