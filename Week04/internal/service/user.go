package service

import (
	"../dao"
	"google.golang.org/grpc/status"
)

type Service struct {
	dao dao.Dao
}

func NewService(d dao.Dao) *Service {
	return &Service{dao: d}
}

func (s *Service) UserById(id int) (*model.User, error) {
	user, err := s.dao.UserById(id)
	if err != nil {
		if errors.Is(err, dao.ErrRecordNotFound) {
			return nil, status.Errorf(err, "User Not Found")
		}
		return nil, status.Errorf(err, "Internal Error:%v", err)
	}
	return user, nil
}
