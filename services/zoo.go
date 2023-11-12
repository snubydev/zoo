package services

import "strings"

type Zoo struct {
}

func NewZoo() *Zoo {
	return &Zoo{}
}

func (c *Zoo) Search(text string) []string {
	result := []string{}
	for _, v := range animals {
		if strings.Index(strings.ToLower(v), strings.ToLower(text)) > -1 {
			result = append(result, v)
		}
	}
	return result
}
