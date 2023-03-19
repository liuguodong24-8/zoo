package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
	"helloworld/internal/util"
	"time"
)

type AnimalRepo struct {
	data *Data
	log  *log.Helper
}

func NewAnimalRepo(data *Data, logger log.Logger) biz.AnimalRepo {
	return &AnimalRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *AnimalRepo) Save(ctx context.Context, a *biz.Animal) (*biz.Animal, error) {
	err := repo.data.gormDB.Create(a)
	if err != nil {
		return &biz.Animal{}, err.Error
	}
	return a, nil
}

func (repo *AnimalRepo) Get(ctx context.Context, stu *biz.Animal) (*biz.Animal, error) {
	return stu, nil
}

func (repo *AnimalRepo) GetAnimal(ctx context.Context, tokenId int64) (*biz.Animal, error) {
	var a biz.Animal
	repo.data.gormDB.Where("token_id = ?", tokenId).First(&a) // gorm
	repo.log.WithContext(ctx).Info("gormDB: GetAnimal, token_id: ", tokenId)
	return &biz.Animal{
		ID:          a.ID,
		TokenId:     a.TokenId,
		Status:      a.Status,
		Address:     a.Address,
		ImgUrl:      a.ImgUrl,
		Age:         a.Age,
		StakeAmount: a.StakeAmount,
		StartTime:   a.StartTime,
		Volume:      a.Volume,
		Category:    a.Category,
		Level:       a.Level,
	}, nil
}

func (repo *AnimalRepo) GetAnimals(ctx context.Context, name string, page int32, pageSize int32) ([]*biz.Animal, error) {
	var ples []*biz.Animal
	db := repo.data.gormDB
	/*if name != "" {
		db = db.Scopes(util.ColumnEqualScope("name", name))
	}*/
	db = db.Scopes(util.ColumnEqualScope("status", 1))
	if err := db.Scopes(util.PaginationScope((page-1)*pageSize, pageSize)).Find(&ples).Error; err != nil {
		return []*biz.Animal{}, nil
	}
	return ples, nil
}

// Kill 杀死
func (repo *AnimalRepo) Kill(ctx context.Context, tokenId int64) (bool, error) {
	repo.log.WithContext(ctx).Info("gormDB: KillAnimal, token_id: ", tokenId)
	err := repo.data.gormDB.Select("status", "kill_time").Where("token_id", tokenId).Updates(biz.Animal{
		Status:   0,
		KillTime: time.Now(),
	})
	if err != nil {
		return false, err.Error
	}
	return true, nil
}

// Feeding 投喂 Feeding
func (repo *AnimalRepo) Feeding(ctx context.Context, fromTokenId int64, toTokenId int64) (bool, error) {
	var a, b biz.Animal
	repo.data.gormDB.Where("token_id = ?", fromTokenId).First(&a) // gorm
	repo.data.gormDB.Where("token_id = ?", toTokenId).First(&b)   // gorm
	if a.Status == 0 || b.Status == 0 {
		return false, errors.New("status not equal 0")
	}
	repo.log.WithContext(ctx).Info("gormDB: FeedingAnimal, token_id: ", toTokenId)
	err := repo.data.gormDB.Select("stake_amount", "level").Where("token_id", toTokenId).Updates(biz.Animal{
		StakeAmount: a.StakeAmount + b.StakeAmount,
		Level:       b.Level + 1,
	})
	if err.Error != nil {
		return false, err.Error
	}
	repo.log.WithContext(ctx).Info("gormDB: KillAnimal, token_id: ", fromTokenId)
	err1 := repo.data.gormDB.Select("status", "kill_time").Where("token_id", fromTokenId).Updates(biz.Animal{
		Status:   0,
		KillTime: time.Now(),
	})
	if err1.Error != nil {
		return false, err.Error
	}
	return true, nil
}

// Compose 合成 Compose
func (repo *AnimalRepo) Compose(ctx context.Context, tokenId1 int64, tokenId2 int64, ani *biz.Animal) (bool, error) {
	var a, b biz.Animal
	repo.data.gormDB.Where("token_id = ?", tokenId1).First(&a) // gorm
	repo.data.gormDB.Where("token_id = ?", tokenId2).First(&b) // gorm
	if a.Status == 0 || b.Status == 0 {
		return false, errors.New("status not equal 0")
	}
	if a.Sex == b.Sex {
		return false, errors.New("not sample sex")
	}
	repo.log.WithContext(ctx).Info("gormDB: KillAnimal, token_id: ", tokenId1)
	err := repo.data.gormDB.Select("status", "kill_time").Where("token_id", tokenId1).Updates(biz.Animal{
		Status:   0,
		KillTime: time.Now(),
	})

	if err.Error != nil {
		return false, err.Error
	}
	repo.log.WithContext(ctx).Info("gormDB: KillAnimal, token_id: ", tokenId2)
	err = repo.data.gormDB.Select("status", "kill_time").Where("token_id", tokenId2).Updates(biz.Animal{
		Status:   0,
		KillTime: time.Now(),
	})
	if err.Error != nil {
		return false, err.Error
	}
	err = repo.data.gormDB.Create(ani)
	if err.Error != nil {
		return false, err.Error
	}
	return true, nil
}
