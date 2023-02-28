package arrangeItemsByProperty

type Interface interface {
	Arrange(items []Item, propName string) []Item
}

type Item struct {
	name             string
	expiryTimeInDays int
}
