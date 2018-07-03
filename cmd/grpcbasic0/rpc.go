package main

import (
	"github.com/daved/grpcbasic0/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// UserService ...
type UserService struct {
	Users *pb.Users
}

// NewUserService ...
func NewUserService() *UserService {
	return &UserService{
		Users: &pb.Users{
			Users: []*pb.User{
				{Id: 1, Name: "Alice", Age: 21, Fortune: fortunes.get(1)},
				{Id: 2, Name: "Bob", Age: 30, Fortune: fortunes.get(2)},
				{Id: 3, Name: "Charlie", Age: 25, Fortune: fortunes.get(3)},
				{Id: 4, Name: "Dana", Age: 18, Fortune: fortunes.get(4)},
			},
		},
	}
}

// RecordUser ...
func (s *UserService) RecordUser(ctx context.Context, req *pb.UserRecordReq) (*pb.User, error) {
	if req.Age == 0 {
		return nil, grpc.Errorf(codes.FailedPrecondition, "user record must include age (int64)")
	}

	tid := s.newestUserID()
	usr := &pb.User{
		Id:      tid,
		Name:    req.Name,
		Age:     req.Age,
		Fortune: fortunes.get(tid),
	}

	s.Users.Users = append(s.Users.Users, usr)

	return usr, nil
}

// GetUser ...
func (s *UserService) GetUser(ctx context.Context, req *pb.UserGetReq) (*pb.User, error) {
	for k, v := range s.Users.Users {
		if v.Id == req.Id {
			return s.Users.Users[k], nil
		}
	}

	return nil, grpc.Errorf(codes.NotFound, "cannot find user 'id: %d'", req.Id)
}

// GetUsers ...
func (s *UserService) GetUsers(ctx context.Context, req *pb.UsersGetReq) (*pb.Users, error) {
	usrs := s.Users.Users
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

	return &pb.Users{Users: usrs[first:last]}, nil
}

func (s *UserService) newestUserID() int64 {
	return s.Users.Users[len(s.Users.Users)-1].Id + 1
}

func filterMax(ind, length int64) int64 {
	if ind > length {
		return length
	}

	return ind
}

func reverse(s []*pb.User) []*pb.User {
	r := make([]*pb.User, len(s))
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
