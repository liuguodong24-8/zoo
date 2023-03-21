package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
	"helloworld/internal/util"
)

type pledgeRepo struct {
	data *Data
	log  *log.Helper
}

func NewPledgeRepo(data *Data, logger log.Logger) biz.PledgeRepo {
	return &pledgeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *pledgeRepo) Save(ctx context.Context, ple *biz.Pledge) (*biz.Pledge, error) {
	zoo, _ := util.NewZoo(common.HexToAddress(repo.data.c.Eth.ZooAddr), repo.data.ethClient)
	stake, err := zoo.GetUserPledge(&bind.CallOpts{}, common.HexToAddress(ple.Address))
	if err != nil {
		return &biz.Pledge{}, err
	}
	fmt.Println(stake.TokenId.Int64())
	if stake.TokenId.Int64() == 0 {
		return &biz.Pledge{}, errors.New("you have not pledge")
	}
	//fmt.Println(stake)
	ple.TokenID = int32(stake.TokenId.Int64())
	tx := repo.data.gormDB.Create(ple)
	if tx.Error != nil {
		return &biz.Pledge{}, tx.Error
	}
	return ple, nil
}

func (repo *pledgeRepo) UpdatePledgeStatus(ctx context.Context, ple *biz.Pledge) (bool, error) {
	repo.log.WithContext(ctx).Info("gormDB: UpdatePledgeStatus, address: ", ple.Address)
	ple.Status = 0
	err := repo.data.gormDB.Where("address = ?", ple.Address).Updates(ple)
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (repo *pledgeRepo) Get(ctx context.Context, stu *biz.Pledge) (*biz.Pledge, error) {
	return stu, nil
}

func (repo *pledgeRepo) GetPledge(ctx context.Context, address string) (*biz.Pledge, error) {
	var ple biz.Pledge
	repo.data.gormDB.Where("address = ?", address).First(&ple) // gorm
	repo.log.WithContext(ctx).Info("gormDB: GetPledge, address: ", address)
	return &biz.Pledge{
		ID:        ple.ID,
		TokenID:   ple.TokenID,
		Address:   ple.Address,
		Price:     ple.Price,
		Times:     ple.Times,
		CreatedAt: ple.CreatedAt,
		UpdateAt:  ple.UpdateAt,
	}, nil
}

func (repo *pledgeRepo) GetPledges(ctx context.Context, name string, page int32, pageSize int32) ([]*biz.Pledge, error) {
	var ples []*biz.Pledge
	db := repo.data.gormDB
	if name != "" {
		db = db.Scopes(util.ColumnEqualScope("name", name))
	}
	if err := db.Scopes(util.PaginationScope((page-1)*pageSize, pageSize)).Find(&ples).Error; err != nil {
		return []*biz.Pledge{}, nil
	}
	return ples, nil
}
