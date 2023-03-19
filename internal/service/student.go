package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"

	pb "helloworld/api/helloworld1/v1"
)

type StudentService struct {
	pb.UnimplementedStudentServer

	student *biz.StudentUserCase
	log     *log.Helper
}

func NewStudentService(stu *biz.StudentUserCase, logger log.Logger) *StudentService {
	return &StudentService{
		student: stu,
		log:     log.NewHelper(logger),
	}
}

func (s *StudentService) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentReply, error) {
	return &pb.CreateStudentReply{}, nil
}
func (s *StudentService) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentReply, error) {
	return &pb.UpdateStudentReply{}, nil
}
func (s *StudentService) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentReply, error) {
	return &pb.DeleteStudentReply{}, nil
}
func (s *StudentService) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentReply, error) {
	stu, err := s.student.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetStudentReply{
		Id:     stu.ID,
		Status: stu.Status,
		Name:   stu.Name,
	}, nil
}
func (s *StudentService) ListStudent(ctx context.Context, req *pb.ListStudentRequest) (*pb.ListStudentReply, error) {
	stus, err := s.student.List(ctx, req.Name, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	reply := &pb.ListStudentReply{
		Students: make([]*pb.ListStudentReply_Stu, 0),
	}
	for _, x := range stus {
		item := &pb.ListStudentReply_Stu{
			Id:        x.ID,
			Name:      x.Name,
			Info:      x.Info,
			CreatedAt: x.CreatedAt.String(),
		}
		reply.Students = append(reply.Students, item)
	}
	return reply, nil
}
func (s *StudentService) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{}, nil
}
