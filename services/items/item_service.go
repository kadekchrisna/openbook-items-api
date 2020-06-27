package services

import (
	"github.com/kadekchrisna/openbook-items-api/domain/items"
	"github.com/kadekchrisna/openbook-items-api/domain/queries"
	errors "github.com/kadekchrisna/openbook-utils-go/rest_errors"
)

var (
	ItemService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, errors.ResErr)
	Get(string) (*items.Item, errors.ResErr)
	Search(queries.EsQuery) ([]items.Item, errors.ResErr)
}

type itemsService struct {
}

func (s *itemsService) Create(item items.Item) (*items.Item, errors.ResErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, errors.ResErr) {
	item := items.Item{Id: id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, errors.ResErr) {
	dao := items.Item{}
	return dao.Search(query)
}
