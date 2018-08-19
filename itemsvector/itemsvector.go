package itemsvector

type ItemsVector struct {
	items []int64
}

func NewItemsVector(data []int64, size int) *ItemsVector {
	return &ItemsVector{
		items: make([]int64, size),
	}
}
