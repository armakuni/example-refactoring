package main

import (
	"encoding/json"
	"net/http"
	"reflect"
)

type ListResults struct {
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
	Paginated
}

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Error struct {
	Error string `json:"error"`
}

type ListFunc func() (interface{}, Paginated, error)

//
//
//
//
//
//
//
//
//

func HandleStuff(w http.ResponseWriter, req *http.Request) {
	filename := req.URL.Query().Get("filename")
	if filename == "" {
		filename = "names.txt"
	}
	statusCode, body := RenderList(func() (interface{}, Paginated, error) {
		return Search(filename)
	})
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(body)
}

func Search(filename string) ([]Item, Paginated, error) {
	items := make([]Item, 0)
	pagination, err := NewFileReader(filename).Paginate(func(index int, name string) {
		items = append(items, Item{index + 1, name})
	})
	return items, pagination, err
}

//
//
//
//
//
//
//
//
//
//
//
//

func RenderError(err error) (int, []byte) {
	payload := &Error{Error: err.Error()}
	encoded, _ := json.MarshalIndent(payload, "", " ")
	return 500, encoded
}

func RenderList(listFunc ListFunc) (int, []byte) {
	if data, pagination, err := listFunc(); err != nil {
		return RenderError(err)
	} else {
		results := ListResults{
			Paginated: pagination,
			Count:     SliceLength(data),
			Data:      data,
		}
		encoded, _ := json.MarshalIndent(&results, "", "  ")
		return 200, encoded
	}
}

func SliceLength(obj interface{}) int {
	v := reflect.ValueOf(obj)
	return v.Len()
}
