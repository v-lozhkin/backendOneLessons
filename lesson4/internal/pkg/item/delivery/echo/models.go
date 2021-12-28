package echo

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageLink   string `json:"image_link,omitempty"`
}

type ItemFilter struct {
	ID       int `param:"id"`
	PriceMin int `query:"price_min"`
	PriceMax int `query:"price_max"`
}
