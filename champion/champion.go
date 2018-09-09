package champion

type Champion struct {
	name string
}

func NewChampion(name string) *Champion {
	return &Champion{
		name: name,
	}
}
