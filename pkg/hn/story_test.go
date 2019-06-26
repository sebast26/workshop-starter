package hn_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"workshop-starter/pkg/hn"
	"workshop-starter/pkg/hn/mock"
)

func TestStoryBuilder(t *testing.T) {
	t.Run("error getting root item", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mock.NewMockMyClient(ctrl)
		client.EXPECT().Get(123).Return(hn.Item{}, errors.New("boom"))

		story, err := hn.NewStoryBuild(client).Build(123)
		assert.Error(t, err)
		assert.Empty(t, story)
	})

	t.Run("story as root, no childs", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mock.NewMockMyClient(ctrl)
		client.EXPECT().Get(123).Return(hn.Item{
			Author: "test",
			Title:  "Title test",
			Score:  10,
			Kids:   []int{},
		}, nil)

		story, err := hn.NewStoryBuild(client).Build(123)
		assert.NoError(t, err)
		assert.NotEmpty(t, story)
		assert.Equal(t, "test", story.Author)
		assert.Equal(t, "Title test", story.Title)
		assert.Equal(t, 10, story.Score)
		assert.Empty(t, story.Comments)
	})

	t.Run("story as root, single child", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mock.NewMockMyClient(ctrl)
		client.EXPECT().Get(123).Return(hn.Item{
			Id:     123,
			Author: "test",
			Title:  "Title test",
			Score:  10,
			Kids:   []int{456},
		}, nil)
		client.EXPECT().Get(456).Return(hn.Item{
			Id:     456,
			Author: "author1",
			Parent: 123,
			Text:   "Some text",
			Type:   "comment",
		}, nil)

		story, err := hn.NewStoryBuild(client).Build(123)
		assert.NoError(t, err)
		assert.NotEmpty(t, story)
		assert.Equal(t, 123, story.Id)
		assert.Equal(t, "test", story.Author)
		assert.Equal(t, "Title test", story.Title)
		assert.Equal(t, 10, story.Score)
		assert.NotEmpty(t, story.Comments)
		assert.Equal(t, 456, story.Comments[0].Id)
		assert.Equal(t, "author1", story.Comments[0].Author)
		assert.Equal(t, 123, story.Comments[0].Parent)
		assert.Equal(t, "Some text", story.Comments[0].Text)
		assert.Empty(t, story.Comments[0].Comments)
		assert.Equal(t, 1, len(story.Comments))
	})

	t.Run("story as root, child of child", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mock.NewMockMyClient(ctrl)
		client.EXPECT().Get(123).Return(hn.Item{Id: 123, Kids: []int{456}}, nil)
		client.EXPECT().Get(456).Return(hn.Item{Id: 456, Kids: []int{789}}, nil)
		client.EXPECT().Get(789).Return(hn.Item{Id: 789, Kids: []int{}}, nil)

		story, err := hn.NewStoryBuild(client).Build(123)
		assert.NoError(t, err)
		assert.Equal(t, 123, story.Id)
		assert.Equal(t, 456, story.Comments[0].Id)
		assert.Empty(t, story.Comments[0].Comments[0].Comments)
		assert.Equal(t, 789, story.Comments[0].Comments[0].Id)
	})
}
