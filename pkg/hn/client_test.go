package hn_test

import (
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"workshop-starter/pkg/hn"
)

func TestHTTPClient(t *testing.T) {
	t.Run("Maxitem", func(t *testing.T) {
		t.Run("500", func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(500)
			}))
			defer ts.Close()

			item, err := hn.NewHTTPClientFor(ts.URL).Maxitem()
			assert.Error(t, err)
			assert.Empty(t, item)
		})

		t.Run("success", func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, _ = io.WriteString(w, "13")
			}))
			defer ts.Close()

			item, err := hn.NewHTTPClientFor(ts.URL).Maxitem()
			assert.NoError(t, err)
			assert.Equal(t, 13, item)
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("500", func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(500)
			}))
			defer ts.Close()

			item, err := hn.NewHTTPClientFor(ts.URL).Get(123)
			assert.Error(t, err)
			assert.Empty(t, item)
		})

		t.Run("success for story", func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				buf, err := ioutil.ReadFile("testdata/8863.json")
				if assert.NoError(t, err) {
					_, _ = io.WriteString(w, string(buf))
				}
			}))
			defer ts.Close()

			item, err := hn.NewHTTPClientFor(ts.URL).Get(8863)
			assert.NoError(t, err)
			assert.Equal(t, "dhouston", item.Author)
			assert.Equal(t, "My YC app: Dropbox - Throw away your USB drive", item.Title)
			assert.Equal(t, 104, item.Score)
			if assert.NotEmpty(t, item.Kids) {
				assert.Equal(t, 9224, item.Kids[0])
				assert.Equal(t, 8876, item.Kids[len(item.Kids)-1])
			}
		})
	})

}

func TestMaxitem_Integration(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	item, err := hn.NewHTTPClient().Maxitem()
	assert.NoError(t, err)
	assert.NotEmpty(t, item)
}

//func TestGet_Integration(t *testing.T) {
//	if testing.Short() {
//		t.SkipNow()
//	}
//	item, err := hn.NewHTTPClient().Get(8863)
//	assert.NoError(t, err)
//	assert.NotEmpty(t, item)
//	assert.Equal(t, "dhouston", item.Author)
//	assert.Equal(t, "My YC app: Dropbox - Throw away your USB drive", item.Title)
//	assert.Equal(t, 104, item.Score)
//	if assert.NotEmpty(t, item.Kids) {
//		assert.Equal(t, 9224, item.Kids[0])
//		assert.Equal(t, 8876, item.Kids[len(item.Kids) - 1])
//	}
//}
