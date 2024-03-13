package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	orders := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggets": {preptime: 12},
	}
	i, ch := orders[order]
	if !ch {
		return 404
	}
	return i.preptime
}
