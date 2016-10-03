package main

import (
	"github.com/daved/grpcbasic0/idl"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// UserService ...
type UserService struct {
	Users *idl.Users
}

// NewUserService ...
func NewUserService() *UserService {
	return &UserService{
		Users: &idl.Users{
			Users: []*idl.User{
				{Id: 1, Name: "Alice", Age: 21, Fortune: fortunes.get(1)},
				{Id: 2, Name: "Bob", Age: 30, Fortune: fortunes.get(2)},
				{Id: 3, Name: "Charlie", Age: 25, Fortune: fortunes.get(3)},
				{Id: 4, Name: "Dana", Age: 18, Fortune: fortunes.get(4)},
			},
		},
	}
}

// RecordUser ...
func (s *UserService) RecordUser(ctx context.Context, req *idl.UserRecordReq) (*idl.User, error) {
	if req.Age == 0 {
		return nil, grpc.Errorf(codes.FailedPrecondition, "user record must include age (int64)")
	}

	tid := s.newestUserID()
	usr := &idl.User{
		Id:      tid,
		Name:    req.Name,
		Age:     req.Age,
		Fortune: fortunes.get(tid),
	}

	s.Users.Users = append(s.Users.Users, usr)

	return usr, nil
}

// GetUser ...
func (s *UserService) GetUser(ctx context.Context, req *idl.UserGetReq) (*idl.User, error) {
	for k, v := range s.Users.Users {
		if v.Id == req.Id {
			return s.Users.Users[k], nil
		}
	}

	return nil, grpc.Errorf(codes.NotFound, "cannot find user 'id: %d'", req.Id)
}

// ListUsers ...
func (s *UserService) ListUsers(ctx context.Context, req *idl.UsersListReq) (*idl.Users, error) {
	usrs := s.Users.Users
	if req.Down {
		usrs = reverse(usrs)
	}

	if req.Count == 0 {
		req.Count = int64(len(usrs))
	}

	first, last := req.Start, req.Start+req.Count
	if last > int64(len(usrs)) {
		last = int64(len(usrs))
	}

	return &idl.Users{Users: usrs[first:last]}, nil
}

func (s *UserService) newestUserID() int64 {
	return s.Users.Users[len(s.Users.Users)-1].Id + 1
}

func reverse(s []*idl.User) []*idl.User {
	r := make([]*idl.User, len(s))
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
