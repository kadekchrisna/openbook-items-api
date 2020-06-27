package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/kadekchrisna/openbook-items-api/client/elasticsearch"
	"github.com/kadekchrisna/openbook-items-api/domain/queries"
	"github.com/kadekchrisna/openbook-items-api/logger"
	rest_error "github.com/kadekchrisna/openbook-utils-go/rest_errors"
)

const (
	itemsIndex = "items"
	docType    = "_doc"
)

func (i *Item) Save() rest_error.ResErr {
	result, err := elasticsearch.Client.Index(itemsIndex, docType, i)
	if err != nil {
		logger.Error(err.Error(), err)
		return rest_error.NewInternalServerError("Error when inserting item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_error.ResErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(itemsIndex, docType, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_error.NewNotFoundError(fmt.Sprintf("item not found when getting item with id %s", i.Id))
		}
		return rest_error.NewInternalServerError(fmt.Sprintf("Error when getting item with id %s", i.Id), errors.New("database error"))
	}
	bytes, errSource := result.Source.MarshalJSON()
	if errSource != nil {
		return rest_error.NewInternalServerError(fmt.Sprintf("Error when marshalling item with id %s", i.Id), errors.New("database error"))
	}
	if err := json.Unmarshal(bytes, &i); err != nil {
		return rest_error.NewInternalServerError(fmt.Sprintf("Error when marshalling item with id %s", i.Id), errors.New("database error"))
	}
	i.Id = itemId
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, rest_error.ResErr) {
	result, err := elasticsearch.Client.Search(itemsIndex, query.Build())
	if err != nil {
		return nil, rest_error.NewInternalServerError(fmt.Sprint("Error when searching in document query."), errors.New("database error"))
	}
	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_error.NewInternalServerError(fmt.Sprint("Error when trying to parse the result."), errors.New("database error"))
		}
		items[index] = item
	}
	if len(items) == 0 {
		return nil, rest_error.NewNotFoundError(fmt.Sprint("no item found."))
	}
	return items, nil
}
