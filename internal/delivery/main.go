package delivery

import (
	"github.com/haikalvidya/football-league/internal/usecase"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	League        *leagueDelivery
	ContainLetter *containLetterDelivery
}

type deliveryType struct {
	Usecase *usecase.Usecase
}

func NewDelivery(e *echo.Echo, usecase *usecase.Usecase) *Delivery {
	deliveryType := &deliveryType{
		Usecase: usecase,
	}
	delivery := &Delivery{
		League:        (*leagueDelivery)(deliveryType),
		ContainLetter: (*containLetterDelivery)(deliveryType),
	}

	Route(e, delivery)

	return delivery
}

func Route(e *echo.Echo, delivery *Delivery) {
	// league
	league := e.Group("/football")
	{
		league.POST("/recordgame", delivery.League.RecordGame)
		league.GET("/leaguestanding", delivery.League.LeagueStanding)
		league.GET("/rank", delivery.League.ClubStanding)
	}

	// contain letter
	e.POST("/is-contain-letters", delivery.ContainLetter.ContainLetter)
}
