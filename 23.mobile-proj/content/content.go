package content

import (
	"errors"
)

type Content struct {
	ProductId string
	Name string
	JsonData string
	Version int
}

func NewContent (productId string, name string, jsonData string) (*Content, error) {
	if len(name) == 0 {
		return nil, errors.New("Empty Name")
	}
	var createdContent = new(Content)
	createdContent.ProductId = productId
	createdContent.Name = name
	createdContent.JsonData = jsonData
	createdContent.Version = 1
	return createdContent, nil
}