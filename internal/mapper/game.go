package mapper

import (
	"NintendoCenter/gateway/graph/model"
	"NintendoCenter/gateway/internal/protos"
)

var shopMap = map[protos.Shop]model.Shop{
	protos.Shop_NINTENDO: model.ShopNintendo,
}

type GameMapper struct {
}

func (m GameMapper) GameToResponse(game *protos.Game) *model.Game {
	bestOffer := m.findBestOffer(game.Offers)
	offers := make([]*model.Offer, 0, len(game.Offers))
	for _, o := range game.GetOffers() {
		offers = append(offers, m.offerToResponse(o))
	}

	return &model.Game{
		ID:          game.Id,
		Title:       game.Title,
		Description: &game.Description,
		ImageURL:    &game.ImageUrl,
		Offers:      offers,
		BestOffer:   m.offerToResponse(bestOffer),
	}
}

func (m GameMapper) GamesToResponse(games []*protos.Game) []*model.Game {
	result := make([]*model.Game, 0, len(games))
	for _, g := range games {
		result = append(result, m.GameToResponse(g))
	}

	return result
}

func (m GameMapper) offerToResponse(offer *protos.Offer) *model.Offer {
	return &model.Offer{
		Shop:      shopMap[offer.Shop],
		IsDigital: offer.IsDigital,
		IsUsed:    offer.IsUsed,
		Link:      offer.Link,
		Price:     m.priceToResponse(offer.Price),
	}
}

func (m GameMapper) priceToResponse(price *protos.Price) *model.Price {
	if price == nil {
		return nil
	}
	return &model.Price{
		Original:    floatPtr(float64(price.Original)),
		Discrounted: floatPtr(float64(price.Discounted)),
		Real:        float64(price.Real),
	}
}

func floatPtr(n float64) *float64 {
	return &n
}


func (m GameMapper) findBestOffer(offers []*protos.Offer) *protos.Offer {
	var bestOffer *protos.Offer
	for _, o := range offers {
		if bestOffer == nil || o.Price.Real < bestOffer.Price.Real {
			bestOffer = o
		}
	}
	return bestOffer
}