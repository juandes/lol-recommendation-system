package champion

type Champion struct {
	name string
}

// test
func NewChampion(name string) *Champion {
	return &Champion{
		name: name,
	}
}
