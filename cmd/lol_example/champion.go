package main

type champion struct {
	name string
}

func NewChampion(name string) *champion {
	return &champion{
		name: name,
	}
}
