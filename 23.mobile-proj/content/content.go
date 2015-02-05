package content

import (
	"errors"
	"appengine"
	"appengine/datastore"
)

type Content struct {
	Name string
	JsonData string
	Version int
}

func NewContent (name string, jsonData string) (*Content, error) {
	if len(name) == 0 {
		return nil, errors.New("Empty Name")
	}
	var createdContent = new(Content)
	createdContent.Name = name
	createdContent.JsonData = jsonData
	createdContent.Version = 1
	return createdContent, nil
}

func (this *Content) Save(context appengine.Context) (*datastore.Key, error) {
	key, err := datastore.Put(context, datastore.NewIncompleteKey(context, "Content", nil), this)
    if err != nil {
        return nil, err
    }

	return key, nil
}