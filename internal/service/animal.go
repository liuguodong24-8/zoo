package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
	"time"

	pb "helloworld/api/pledge/v1"
)

type AnimalService struct {
	pb.UnimplementedAnimalServer
	ple *biz.AnimalCase
	log *log.Helper
}

func NewAnimalService(ac *biz.AnimalCase, logger log.Logger) *AnimalService {
	return &AnimalService{
		ple: ac,
		log: log.NewHelper(logger),
	}
}

func (s *AnimalService) CreateAnimal(ctx context.Context, req *pb.CreateAnimalRequest) (*pb.CreateAnimalReply, error) {
	_, err := s.ple.CreateAnimal(ctx, &biz.AnimalRequest{
		TokenId:     req.TokenId,
		Address:     req.Address,
		Level:       req.Level,
		Category:    req.Category,
		Volume:      req.Volume,
		Rarity:      req.Rarity,
		Sex:         req.Sex,
		Age:         req.Age,
		StartTime:   time.Now(),
		StakeAmount: req.StakeAmount,
		ImgUrl:      "http://1.png", //从合约拿
		Status:      1,
	})

	if err != nil {
		return &pb.CreateAnimalReply{
			Code: 410,
			Msg:  err.Error(),
		}, err
	}
	return &pb.CreateAnimalReply{
		Code: 0,
		Msg:  "",
	}, nil
}
func (s *AnimalService) UpdateAnimal(ctx context.Context, req *pb.UpdateAnimalRequest) (*pb.UpdateAnimalReply, error) {
	return &pb.UpdateAnimalReply{}, nil
}
func (s *AnimalService) DeleteAnimal(ctx context.Context, req *pb.DeleteAnimalRequest) (*pb.DeleteAnimalReply, error) {
	return &pb.DeleteAnimalReply{}, nil
}
func (s *AnimalService) GetAnimal(ctx context.Context, req *pb.GetAnimalRequest) (*pb.GetAnimalReply, error) {
	a, err := s.ple.GetAnimal(ctx, req.TokenId)
	if err != nil {
		return &pb.GetAnimalReply{
			Code: 0,
			Msg:  "",
			Data: nil,
		}, nil
	}
	return &pb.GetAnimalReply{
		Code: 0,
		Msg:  "",
		Data: &pb.AnimalInfo{
			Id:          a.ID,
			TokenId:     a.TokenId,
			Status:      a.Status,
			Address:     a.Address,
			ImgUrl:      a.ImgUrl,
			Age:         a.Age,
			StakeAmount: a.StakeAmount,
			StartTime:   a.StartTime.Unix(),
			Volume:      a.Volume,
			Category:    a.Category,
			Level:       a.Level,
		},
	}, nil
	return &pb.GetAnimalReply{}, nil
}
func (s *AnimalService) ListAnimal(ctx context.Context, req *pb.ListAnimalRequest) (*pb.ListAnimalReply, error) {
	stus, err := s.ple.ListAnimal(ctx, "", req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	reply := &pb.ListAnimalReply{
		Data: make([]*pb.AnimalInfo, 0),
	}
	for _, x := range stus {
		item := &pb.AnimalInfo{
			Id:          x.ID,
			Address:     x.Address,
			StakeAmount: x.StakeAmount,
			Level:       x.Level,
			Category:    x.Category,
			Capacity:    x.Capacity,
			Volume:      x.Volume,
			Rarity:      x.Rarity,
			Sex:         x.Sex,
			Age:         x.Age,
			TokenId:     x.TokenId,
			ImgUrl:      x.ImgUrl,
			Status:      x.Status,
			StartTime:   x.StartTime.Unix(),
		}
		reply.Data = append(reply.Data, item)
	}
	return reply, nil
}

func (s *AnimalService) KillAnimal(ctx context.Context, req *pb.KillAnimalRequest) (*pb.KillAnimalReply, error) {
	_, err := s.ple.Kill(ctx, req.TokenId)
	if err != nil {
		return &pb.KillAnimalReply{
			Code: 411,
			Msg:  err.Error(),
		}, err
	}
	return &pb.KillAnimalReply{
		Code: 0,
		Msg:  "",
	}, nil
}

func (s *AnimalService) FeedingAnimal(ctx context.Context, req *pb.FeedingAnimalRequest) (*pb.FeedingAnimalReply, error) {
	_, err := s.ple.Feeding(ctx, req.FromTokenId, req.ToTokenId)
	if err != nil {
		return &pb.FeedingAnimalReply{
			Code: 411,
			Msg:  err.Error(),
		}, err
	}
	return &pb.FeedingAnimalReply{
		Code: 0,
		Msg:  "",
	}, nil
}

func (s *AnimalService) ComposeAnimal(ctx context.Context, req *pb.ComposeAnimalRequest) (*pb.ComposeAnimalReply, error) {
	_, err := s.ple.Compose(ctx, req.TokenId1, req.TokenId2, &biz.Animal{
		TokenId:     1,
		Address:     "xxxx",
		Level:       1,
		Category:    1,
		Volume:      1,
		Rarity:      1,
		Sex:         1,
		Age:         1,
		StartTime:   time.Now(),
		StakeAmount: 1,
		ImgUrl:      "http://1.png", //从合约拿
		Status:      1,
	})
	if err != nil {
		return &pb.ComposeAnimalReply{
			Code: 411,
			Msg:  err.Error(),
		}, err
	}
	return &pb.ComposeAnimalReply{
		Code: 0,
		Msg:  "",
	}, nil
}
