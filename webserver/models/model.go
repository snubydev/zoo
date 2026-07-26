package models

type PageItem struct {
	Icon  string
	Label string
	Link  string
	Page  string
}

type Data struct {
	Active string
	HeaderTitle string
	Pages []PageItem
	ListTitle string
	ListAnimals []string
}

