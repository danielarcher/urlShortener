package storages

import (
	"io/ioutil"
	"urlShortener/rand"
)

type FileSystem struct {
	Path string
}
func (f FileSystem) Save(url string) (string, error) {
	code := rand.String(8)
	err := ioutil.WriteFile(f.Path+"/"+code, []byte(url), 0644)

	return code, err
}

func (f FileSystem) Load(code string) (string, error) {
	urlBytes, err := ioutil.ReadFile(f.Path +"/"+code)

	return string(urlBytes), err
}