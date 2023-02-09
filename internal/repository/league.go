package repository

import "github.com/haikalvidya/football-league/internal/models"

type ILeagueRepository interface {
	SelectOrderByPoints() ([]*models.LeagueModel, error)
	SelectByClubName(clubName string) (*models.LeagueModel, error)
	Delete(league *models.LeagueModel) error
	Update(league *models.LeagueModel) error
}

type leagueRepository repositoryType

func (r *leagueRepository) SelectOrderByPoints() ([]*models.LeagueModel, error) {
	leagues := []*models.LeagueModel{}
	err := r.DB.Order("points desc").Find(&leagues).Error
	if err != nil {
		return nil, err
	}
	return leagues, nil
}

func (r *leagueRepository) SelectByClubName(clubName string) (*models.LeagueModel, error) {
	league := &models.LeagueModel{}
	err := r.DB.Where("clubname = ?", clubName).First(league).Error
	if err != nil {
		return nil, err
	}
	return league, nil
}

func (r *leagueRepository) Delete(league *models.LeagueModel) error {
	err := r.DB.Delete(league).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *leagueRepository) Update(league *models.LeagueModel) error {
	err := r.DB.Save(league).Error
	if err != nil {
		return err
	}
	return nil
}
