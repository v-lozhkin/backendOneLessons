package models

import "errors"

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageLink   string `json:"image_link,omitempty"`
}

func (i Item) Validate() error {
	if i.Name == "" {
		return errors.New("name can't be empty")
	}

	if i.Price < 0 {
		return errors.New("price can't be less than 0")
	}

	return nil
}

type ItemFilter struct {
	ID       int
	PriceMin int
	PriceMax int
}

func (i ItemFilter) Validate() error {
	if i.PriceMax < i.PriceMin {
		return errors.New("wrong price range")
	}

	return nil
}
