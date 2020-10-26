package encodepage

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"urlShortener/storages"
)

func TestHandler_Encode(t *testing.T) {
	tt := []struct {
		name   string
		url    string
		status int
		err    string
	}{
		{name: "successful request", url: "www.google.com", status: http.StatusOK},
		{name: "missing url argument", url: "", status: http.StatusBadRequest, err: "missing url"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			jsonData := map[string]string{"url": tc.url}
			data, _ := json.Marshal(jsonData)
			req, _ := http.NewRequest("POST", "http://localhost:8080/encode", bytes.NewBuffer(data))
			rec := httptest.NewRecorder()

			logger := log.New(os.Stdout, "", 0)
			storage := storages.NewFileSystem("C:\\webserver") //"C:\\webserver"
			h := NewHandler(logger, storage)
			h.Encode(rec, req)

			res := rec.Result()
			defer res.Body.Close()
			if res.StatusCode != tc.status {
				t.Errorf("expected status %v; got %v", tc.status, res.StatusCode)
			}
			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read from body; got %v", res.Body)
			}
			var encRes EncodeResponse
			err = json.Unmarshal(b, &encRes)
			if err != nil {
				t.Fatalf("not possible to unmarshal response; got %v", b)
			}
			if encRes.RedirectUrl == "" {
				t.Error("redirect url empty")
			}
		})
	}
}
