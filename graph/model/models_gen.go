// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Game struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description *string  `json:"description"`
	ImageURL    *string  `json:"imageUrl"`
	Offers      []*Offer `json:"offers"`
}

type Notification struct {
	GameID         string   `json:"gameId"`
	PriceLowerThan *float64 `json:"priceLowerThan"`
}

type NotificationInput struct {
	GameID         string   `json:"gameId"`
	PriceLowerThan *float64 `json:"priceLowerThan"`
}

type Offer struct {
	Shop      Shop   `json:"shop"`
	IsDigital bool   `json:"isDigital"`
	IsUsed    bool   `json:"isUsed"`
	Link      string `json:"link"`
	Price     *Price `json:"price"`
}

type Price struct {
	Original    *float64 `json:"original"`
	Discrounted *float64 `json:"discrounted"`
	Real        float64  `json:"real"`
}

type Shop string

const (
	ShopNintendo Shop = "NINTENDO"
	ShopSavela   Shop = "SAVELA"
	ShopInverest Shop = "INVEREST"
)

var AllShop = []Shop{
	ShopNintendo,
	ShopSavela,
	ShopInverest,
}

func (e Shop) IsValid() bool {
	switch e {
	case ShopNintendo, ShopSavela, ShopInverest:
		return true
	}
	return false
}

func (e Shop) String() string {
	return string(e)
}

func (e *Shop) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Shop(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Shop", str)
	}
	return nil
}

func (e Shop) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
