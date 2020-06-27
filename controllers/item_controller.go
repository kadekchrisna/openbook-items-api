package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kadekchrisna/openbook-items-api/domain/items"
	"github.com/kadekchrisna/openbook-items-api/domain/queries"
	services "github.com/kadekchrisna/openbook-items-api/services/items"
	"github.com/kadekchrisna/openbook-items-api/utils/http_utils"
	"github.com/kadekchrisna/openbook-oauth-go/oauth"
	"github.com/kadekchrisna/openbook-utils-go/logger"
	errors "github.com/kadekchrisna/openbook-utils-go/rest_errors"
)

var (
	ItemController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Search(http.ResponseWriter, *http.Request)
}

type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		logger.Info("AuthenticateRequest")
		http_utils.ResponseError(w, nil)
		return
	}
	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		err := errors.NewUnAuthorizedError()
		http_utils.ResponseError(w, err)
		return
	}

	requestBody, errReadBody := ioutil.ReadAll(r.Body)
	if errReadBody != nil {
		respErr := errors.NewBadRequestError("invalid request body")
		http_utils.ResponseError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := errors.NewBadRequestError("invalid item json body")
		http_utils.ResponseError(w, respErr)
		return
	}
	itemRequest.Seller = sellerId

	result, errCreate := services.ItemService.Create(itemRequest)
	if errCreate != nil {
		http_utils.ResponseError(w, errCreate)
		return
	}
	fmt.Println(result)
	http_utils.ResponseSuccess(w, http.StatusCreated, result)
	return
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		logger.Info("AuthenticateRequest")
		http_utils.ResponseError(w, nil)
		return
	}
	if oauth.GetCallerId(r) == 0 {
		err := errors.NewUnAuthorizedError()
		http_utils.ResponseError(w, err)
		return
	}
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	fmt.Println(itemId)

	item, err := services.ItemService.Get(itemId)
	if err != nil {
		http_utils.ResponseError(w, err)
		return
	}
	http_utils.ResponseSuccess(w, http.StatusOK, item)
	return

}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		readErr := errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, readErr)
		return
	}
	defer r.Body.Close()
	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		readErr := errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, readErr)
		return
	}
	items, errSearch := services.ItemService.Search(query)
	if errSearch != nil {
		http_utils.ResponseError(w, errSearch)
		return
	}
	http_utils.ResponseSuccess(w, http.StatusOK, items)
	return
}
