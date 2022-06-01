package shipping

type ShippingService struct {
	Rate float64
}

func NewShippingService(rate float64) *ShippingService {
	return &ShippingService{Rate: rate}
}

func (s *ShippingService) CalculateShippingCost(weight float64) float64 {
	return weight * s.Rate
}
