package helper

import (
	"net/http"

	"appengine"
)

const UsedNamespace string = "mobileCMS"

func SetContext(r *http.Request) appengine.Context{
	context := appengine.NewContext(r)
	context,_ = appengine.Namespace(context, UsedNamespace)
	return context
}

func HasError(w http.ResponseWriter, err error, errorMessage string) bool {
	if err != nil {
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return true
	}
	return false
}