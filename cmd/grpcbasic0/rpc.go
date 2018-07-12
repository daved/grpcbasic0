package main

import (
	"context"

	"github.com/daved/grpcbasic0/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserService ...
type UserService struct {
	UsersResp *pb.UsersResp
}

// NewUserService ...
func NewUserService() *UserService {
	return &UserService{
		UsersResp: &pb.UsersResp{
			Users: []*pb.UserResp{
				{Id: 1, Name: "Alice", Age: 21, Fortune: fortunes.get(1)},
				{Id: 2, Name: "Bob", Age: 30, Fortune: fortunes.get(2)},
				{Id: 3, Name: "Charlie", Age: 25, Fortune: fortunes.get(3)},
				{Id: 4, Name: "Dana", Age: 18, Fortune: fortunes.get(4)},
			},
		},
	}
}

// NewUser ...
func (s *UserService) NewUser(ctx context.Context, req *pb.NewUserReq) (*pb.UserResp, error) {
	if req.Age == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "user record must include age (int64)")
	}

	tid := s.newestUserID()
	usr := &pb.UserResp{
		Id:      tid,
		Name:    req.Name,
		Age:     req.Age,
		Fortune: fortunes.get(tid),
	}

	s.UsersResp.Users = append(s.UsersResp.Users, usr)

	return usr, nil
}

// GetUser ...
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.UserResp, error) {
	for k, v := range s.UsersResp.Users {
		if v.Id == req.Id {
			return s.UsersResp.Users[k], nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "cannot find user 'id: %d'", req.Id)
}

// GetUsers ...
func (s *UserService) GetUsers(ctx context.Context, req *pb.GetUsersReq) (*pb.UsersResp, error) {
	usrs := s.UsersResp.Users
	if req.Desc {
		usrs = reverse(usrs)
	}

	if req.Start < 0 {
		req.Start = 0
	}
	if req.Count <= 0 {
		req.Count = int64(len(usrs))
	}

	first, last := req.Start, req.Start+req.Count
	first = filterMax(first, int64(len(usrs)))
	last = filterMax(last, int64(len(usrs)))

	return &pb.UsersResp{Users: usrs[first:last]}, nil
}

func (s *UserService) newestUserID() int64 {
	return s.UsersResp.Users[len(s.UsersResp.Users)-1].Id + 1
}

func filterMax(ind, length int64) int64 {
	if ind > length {
		return length
	}

	return ind
}

func reverse(s []*pb.UserResp) []*pb.UserResp {
	r := make([]*pb.UserResp, len(s))
	copy(r, s)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return r
}

type fortuneTeller []string

func (t *fortuneTeller) get(i int64) string {
	if i == 0 {
		i = int64(len(*t))
	}

	i--

	return (*t)[i%int64(len(*t))]
}

var (
	fortunes = fortuneTeller{
		"Great day!",
		"Nice day.",
		"OK day.",
		"Sad day.",
		"Terrible day!",
		"Interesting day.",
		"Calm day.",
	}
)
