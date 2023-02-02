package arrangeGoodsByProperty

type GoodsByPropertyArranger interface {
	ArrangeGoodsByProperty(items []Item, propName string) []Item
}

type Item struct {
	name             string
	expiryTimeInDays int
}
