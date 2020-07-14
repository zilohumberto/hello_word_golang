package cache

import (
	"context"
	"errors"
	"github.com/go-redis/redis"
)

type Service interface {
	Set(string, string) error
	Get(string) (*string, error)
}

type service struct {
	ctx context.Context
	redis redis.Client
}

func NewService() Service {
	return &service{
		context.Background(),
		*redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})}
}

func (s *service) Set(key string, value string) error {
	err := s.redis.Set(s.ctx, key, value, 0).Err()
	return err
}

func (s *service) Get(key string) (*string, error){
	raw, err := s.redis.Get(s.ctx, key).Result()
	if err == redis.Nil{
		return nil, errors.New("key does not exits")
	}
	if err != nil{
		return nil, err
	}
	return &raw, nil
}
