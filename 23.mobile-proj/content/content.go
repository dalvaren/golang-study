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
	ProductKey *datastore.Key
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

func Load(context appengine.Context, encodedKey string) (*Content, error) {
	key, _ := datastore.DecodeKey(encodedKey)
	var loadedEntity = new(Content)
	err := datastore.Get(context, key, loadedEntity);
	if err != nil {
		return nil, err
	}
	return loadedEntity, nil
}