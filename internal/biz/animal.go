package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Animal struct {
	ID          int64
	Level       int64
	Category    int64
	Capacity    int64
	Volume      int64
	Rarity      int64
	Sex         int64
	Age         int64
	StartTime   time.Time
	KillTime    time.Time
	StakeAmount int64
	Status      int64
	ImgUrl      string
	TokenId     int64
	Address     string
}

type AnimalRequest struct {
	ID          int64
	Level       int64
	Category    int64
	Capacity    int64
	Volume      int64
	Rarity      int64
	Sex         int64
	Age         int64
	StartTime   time.Time
	KillTime    time.Time
	StakeAmount int64
	Status      int64
	ImgUrl      string
	TokenId     int64
	Address     string
	Hash        string
}

// AnimalRepo 定义animal的操作接口
type AnimalRepo interface {
	Save(ctx context.Context, Animal *AnimalRequest) (*AnimalRequest, error)
	Get(ctx context.Context, Animal *Animal) (*Animal, error)
	GetAnimal(context.Context, int64) (*Animal, error)
	GetAnimals(context.Context, string, int32, int32) ([]*Animal, error)
	Kill(context.Context, int64) (bool, error)
	Feeding(context.Context, int64, int64) (bool, error)
	Compose(context.Context, int64, int64, *Animal) (bool, error)
}

type AnimalCase struct {
	repo AnimalRepo
	log  *log.Helper
}

// NewAnimalCase 初始化StudentUserCase
func NewAnimalCase(repo AnimalRepo, logger log.Logger) *AnimalCase {
	return &AnimalCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AnimalCase) CreateAnimal(ctx context.Context, a *AnimalRequest) (*AnimalRequest, error) {
	uc.log.WithContext(ctx).Info("CreateAnimal: %v", a.ID)
	return uc.repo.Save(ctx, a)
}

func (uc *AnimalCase) GetAnimal(ctx context.Context, tokenId int64) (*Animal, error) {
	uc.log.WithContext(ctx).Infof("biz.Get: %s", tokenId)
	return uc.repo.GetAnimal(ctx, tokenId)
}

func (uc *AnimalCase) ListAnimal(ctx context.Context, name string, page int32, pageSize int32) ([]*Animal, error) {
	uc.log.WithContext(ctx).Info("List: %v", name)
	return uc.repo.GetAnimals(ctx, name, page, pageSize)
}

func (uc *AnimalCase) Kill(ctx context.Context, tokenId int64) (bool, error) {
	uc.log.WithContext(ctx).Info("Kill: %d", tokenId)
	return uc.repo.Kill(ctx, tokenId)
}

func (uc *AnimalCase) Feeding(ctx context.Context, fromTokenId int64, toTokenId int64) (bool, error) {
	uc.log.WithContext(ctx).Info("Feeding fromTokenId: %d", fromTokenId)
	uc.log.WithContext(ctx).Info("Feeding toTokenId: %d", toTokenId)
	return uc.repo.Feeding(ctx, fromTokenId, toTokenId)
}

func (uc *AnimalCase) Compose(ctx context.Context, tokenId1 int64, tokenId2 int64, a *Animal) (bool, error) {
	uc.log.WithContext(ctx).Info("Compose: %d", tokenId1)
	return uc.repo.Compose(ctx, tokenId1, tokenId2, a)
}
