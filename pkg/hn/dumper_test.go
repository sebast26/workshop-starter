package hn_test

import (
	"bytes"
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
	"workshop-starter/pkg/hn"
	"workshop-starter/pkg/hn/mock"
)

type errWriter struct {
}

func (e errWriter) Write(p []byte) (int, error) {
	return 0, errors.New("boom")
}

func TestDumper(t *testing.T) {
	t.Run("write error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mock.NewMockMyClient(ctrl)
		client.EXPECT().Maxitem().Return(1, nil)
		client.EXPECT().Get(1).Return(hn.Item{Id: 1}, nil)

		err := hn.NewDump(client).Dump(&errWriter{})

		assert.Error(t, err)
	})

	t.Run("sucess", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mock.NewMockMyClient(ctrl)
		client.EXPECT().Maxitem().Return(2, nil)
		client.EXPECT().Get(2).Return(hn.Item{Title: "A"}, nil)
		client.EXPECT().Get(1).Return(hn.Item{Title: "B"}, nil)

		var b bytes.Buffer
		err := hn.NewDump(client).Dump(&b)
		assert.NoError(t, err)
		assert.Equal(t, "A\nB\n", b.String())
	})

	t.Run("sucess", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mock.NewMockMyClient(ctrl)
		client.EXPECT().Maxitem().Return(4, nil)
		client.EXPECT().Get(4).Return(hn.Item{Title: "D"}, nil)
		client.EXPECT().Get(3).Return(hn.Item{Title: "C"}, nil)

		var b bytes.Buffer
		err := hn.NewDump(client).Dump(&b)
		assert.NoError(t, err)
		assert.Equal(t, "D\nC\n", b.String())
	})

	t.Run("sucess", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mock.NewMockMyClient(ctrl)
		client.EXPECT().Maxitem().Return(1, nil)
		client.EXPECT().Get(1).Return(hn.Item{Title: "C"}, nil)

		var b bytes.Buffer
		err := hn.NewDump(client).Dump(&b)
		assert.NoError(t, err)
		assert.Equal(t, "C\n", b.String())
	})
}

func TestDumperIntegration(t *testing.T) {
	f, err := ioutil.TempFile(os.TempDir(), "hn-dump-")
	assert.NoError(t, err)
	defer f.Close()

	err = hn.NewDump(hn.NewHTTPClient()).Dump(f)
	assert.NoError(t, err)

	spew.Dump(f.Name())
}
