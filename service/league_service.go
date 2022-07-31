package service

import (
	"github.com/CatBloom/MahjongMasterApi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LeagueService interface {
	SearchLeague(value string) ([]League, error)
	GetLeagueList(uid string) ([]League, error)
	GetLeague(id string) (League, error)
	CreateLeague(c *gin.Context) (League, error)
	UpdateLeague(id string, c *gin.Context) (League, error)
	DeleteLeague(id string) error
}
type leagueService struct {
	db *gorm.DB
}

func NewLeagueService(db *gorm.DB) LeagueService {
	return &leagueService{db}
}

type League = models.League

func (s leagueService) SearchLeague(value string) ([]League, error) {
	var l []League
	if err := s.db.Debug().Where("name Like ?", value+"%").Order("created_at desc").Limit(5).Find(&l).Error; err != nil {
		return nil, err
	}

	return l, nil
}

func (s leagueService) GetLeagueList(uid string) ([]League, error) {
	var l []League
	if err := s.db.Debug().Joins("left join admins_leagues as al on leagues.id = al.league_id").Where("uid = ?", uid).Order("created_at desc").Find(&l).Error; err != nil {
		return nil, err
	}

	return l, nil
}

func (s leagueService) GetLeague(id string) (League, error) {
	var l League

	if err := s.db.Debug().Preload("UIDS").Preload("Rules").Table("leagues").Find(&l, "leagues.id = ?", id).Error; err != nil {
		return l, err
	}

	return l, nil
}

func (s leagueService) CreateLeague(c *gin.Context) (League, error) {
	var l League

	//leagueの登録はトランザクション処理をする
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := c.BindJSON(&l); err != nil {
			return err
		}

		if err := tx.Debug().Create(&l).Error; err != nil {
			return err
		}

		a2l := &models.AdminsLeagues{
			LeagueID: l.ID,
			UID:      l.UID,
		}

		if err := tx.Create(&a2l).Error; err != nil {
			return err
		}

		return nil
	})
	//トランザクション処理のエラー判定
	if err != nil {
		return l, err
	}

	return l, nil
}

// if err := errorByErrorsNew(); err != nil {
// 	fmt.Println(err)
// 	return err
// }

// func errorByErrorsNew() error {
// 	return errors.New("error")
// }

func (s leagueService) UpdateLeague(id string, c *gin.Context) (League, error) {
	var l League

	if err := s.db.Where("id = ?", id).First(&l).Error; err != nil {
		return l, err
	}

	if err := c.BindJSON(&l); err != nil {
		return l, err
	}

	s.db.Save(&l)

	return l, nil
}

func (s leagueService) DeleteLeague(id string) error {
	var l League

	if err := s.db.Where("id = ?", id).Delete(&l).Error; err != nil {
		return err
	}

	return nil
}
