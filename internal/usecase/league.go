package usecase

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/haikalvidya/football-league/internal/delivery/payload"
	"github.com/haikalvidya/football-league/internal/models"
	"gorm.io/gorm"
)

type ILeagueUsecase interface {
	RecordGame(req string) error
	LeagueStanding() ([]*payload.LeagueStandingResponse, error)
	ClubStanding(nama string) ([]*payload.ClubStandingResponse, error)
}

type leagueUsecase usecaseType

func (u *leagueUsecase) RecordGame(data string) error {
	var listReq []payload.RecordGameRequest
	var req payload.RecordGameRequest

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		key = strings.Replace(key, "- ", "", -1)
		switch key {
		case "clubhomename":
			req.ClubHomeName = value
		case "clubawayname":
			req.ClubAwayName = value
		case "score":
			req.Score = value + " :" + parts[2]
			listReq = append(listReq, req)
			req = payload.RecordGameRequest{}
		}
	}

	for _, req := range listReq {
		var homeModel, awayModel *models.LeagueModel
		// get club home and club away name from database
		homeModel, err := u.Repo.League.SelectByClubName(req.ClubHomeName)
		if err != nil && err != gorm.ErrRecordNotFound {
			return fmt.Errorf("error to record game")
		}
		if err == gorm.ErrRecordNotFound {
			homeModel = &models.LeagueModel{
				Clubname: req.ClubHomeName,
			}
		}
		awayModel, err = u.Repo.League.SelectByClubName(req.ClubAwayName)
		if err != nil && err != gorm.ErrRecordNotFound {
			return fmt.Errorf("error to record game")
		}
		if err == gorm.ErrRecordNotFound {
			awayModel = &models.LeagueModel{
				Clubname: req.ClubAwayName,
			}
		}

		score := make(map[string]int)
		// parsing score "1 : 1" from req to map
		regex := regexp.MustCompile(`(\d+) : (\d+)`)
		scores := regex.FindStringSubmatch(req.Score)
		if len(scores) != 3 {
			return fmt.Errorf("invalid score format")
		}
		score[req.ClubHomeName], err = strconv.Atoi(scores[1])
		if err != nil {
			return errors.New("invalid score format")
		}
		score[req.ClubAwayName], err = strconv.Atoi(scores[2])
		if err != nil {
			return errors.New("invalid score format")
		}

		// if score is equal, both club get 1 point
		// the winner club get 3 point
		// the loser club get 0 point
		homeModel.Played++
		awayModel.Played++
		if score[req.ClubHomeName] == score[req.ClubAwayName] {
			homeModel.Points += 1
			awayModel.Points += 1
		} else if score[req.ClubHomeName] > score[req.ClubAwayName] {
			homeModel.Points += 3
		} else {
			awayModel.Points += 3
		}

		u.Repo.League.Update(homeModel)
		u.Repo.League.Update(awayModel)
	}
	return nil
}

func (u *leagueUsecase) LeagueStanding() ([]*payload.LeagueStandingResponse, error) {
	leagueModels, err := u.Repo.League.SelectOrderByPoints()
	if err != nil {
		return nil, fmt.Errorf("error to get league standing")
	}
	// convert to payload response
	res := make([]*payload.LeagueStandingResponse, len(leagueModels))
	for i, leagueModel := range leagueModels {
		res[i] = &payload.LeagueStandingResponse{
			ClubName: leagueModel.Clubname,
			Points:   leagueModel.Points,
		}
	}
	return res, nil
}

func (u *leagueUsecase) ClubStanding(name string) ([]*payload.ClubStandingResponse, error) {
	leagueModels, err := u.Repo.League.SelectOrderByPoints()
	if err != nil {
		return nil, fmt.Errorf("error to get club standing")
	}

	// convert to payload response
	res := make([]*payload.ClubStandingResponse, len(leagueModels))
	for i, leagueModel := range leagueModels {
		res[i] = &payload.ClubStandingResponse{
			ClubName: leagueModel.Clubname,
			Standing: i + 1,
		}
	}

	// if club name is not empty, filter the club name like the name
	if name != "" {
		filtered := make([]*payload.ClubStandingResponse, 0)
		for _, clubStanding := range res {
			if strings.Contains(strings.ToLower(clubStanding.ClubName), strings.ToLower(name)) {
				filtered = append(filtered, clubStanding)
			}
		}
		res = filtered
	}

	return res, nil
}
