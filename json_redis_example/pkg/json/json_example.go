package json

import (
	"encoding/json"
	"github.com/zilohumberto/hello_word_golang/json_redis_example/pkg/cache"
)

type Message struct{
	Action string `json:"action"`
	Body string `json:"body"`
	Producer string `json:"producer"`
	Destination string `json:"destination"`
}


type Service interface {
	GetMessage() (*Message, error)
	SetMessage(*Message) error
}

type service struct{
	key string
	cache cache.Service
}

func NewService(key string) Service{
	return &service{key, cache.NewService()}
}

func (s *service) GetMessage() (*Message, error){
	rawJson, err := s.cache.Get(s.key)
	if err != nil{
		return nil, err
	}
	message := Message{}
	err = json.Unmarshal([]byte(*rawJson), &message)
	return &message, err
}

func (s *service) SetMessage(message *Message) error{
	rawJson, err := json.Marshal(message)
	if err != nil{
		return err
	}
	err = s.cache.Set(s.key, string(rawJson))
	return err
}
