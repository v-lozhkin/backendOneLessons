package delivery

import (
	"backendOneLessons/lesson4/internal/pkg/image"
	"backendOneLessons/lesson4/internal/pkg/item"
	"backendOneLessons/lesson4/internal/pkg/models"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type delivery struct {
	items  item.ItemUsecase
	images image.Storage
}

func New(items item.ItemUsecase, images image.Storage) item.RESTDelivery {
	return delivery{
		items:  items,
		images: images,
	}
}

func (d delivery) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var response interface{}
	var err error

	switch r.Method {
	case http.MethodGet:
		response, err = d.handleGET(ctx, r)
	case http.MethodPost:
		response, err = d.handlePOST(ctx, r)
	case http.MethodPut:
		response, err = d.handlePUT(ctx, r)
	case http.MethodDelete:
		response, err = d.handleDELETE(ctx, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if err != nil {
		//TODO: handle errors
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&response); err != nil {
		log.Println("error happened during serve")
	}

}

func (d delivery) handleGET(ctx context.Context, r *http.Request) (interface{}, error) {
	filter := models.ItemFilter{}

	parts := strings.Split(r.URL.Path, "/")[1:]
	if len(parts) == 2 && parts[1] != "" {
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		filter.ID = id
	}

	max := r.FormValue("price_max")
	min := r.FormValue("price_min")

	if max != "" {
		val, err := strconv.Atoi(max)
		if err != nil {
			return nil, err
		}

		filter.PriceMax = val
	}

	if min != "" {
		val, err := strconv.Atoi(min)
		if err != nil {
			return nil, err
		}

		filter.PriceMin = val
	}

	return d.items.List(ctx, filter)
}

func (d delivery) handlePOST(ctx context.Context, r *http.Request) (interface{}, error) {
	parts := strings.Split(r.URL.Path, "/")[1:]

	if len(parts) == 3 && parts[2] == "upload" {
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		return d.handleUPLOAD(ctx, r, id)
	}

	var newItem = &models.Item{}
	if err := json.NewDecoder(r.Body).Decode(newItem); err != nil {
		return nil, err
	}

	if err := d.items.Create(ctx, newItem); err != nil {
		return nil, err
	}

	return *newItem, nil
}

func (d delivery) handlePUT(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, errors.New("not implemented")
}

func (d delivery) handleDELETE(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, errors.New("not implemented")
}

func (d delivery) handleUPLOAD(ctx context.Context, r *http.Request, id int) (interface{}, error) {
	existingItems, err := d.items.List(ctx, models.ItemFilter{ID: id})
	if err != nil {
		return nil, err
	}

	if len(existingItems) == 0 {
		return nil, errors.New("item doesn't exists")
	}

	file, headers, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	path, err := d.images.Save(ctx, headers.Filename, data)
	if err != nil {
		return nil, err
	}

	existingItems[0].ImageLink = path

	if err = d.items.Update(ctx, existingItems[0]); err != nil {
		return nil, err
	}

	return existingItems[0], nil
}
