package hn

type Story struct {
	Id       int
	Author   string
	Title    string
	Score    int
	Comments []Comment
}

type Comment struct {
	Id       int
	Author   string
	Parent   int
	Text     string
	Comments []Comment
}

type MyStoryBuilder struct {
	client MyClient
}

type StoryBuilder interface {
	Build(id int)
}

func NewStoryBuild(client MyClient) *MyStoryBuilder {
	return &MyStoryBuilder{client}
}

func (sb *MyStoryBuilder) Build(id int) (Story, error) {
	item, err := sb.client.Get(id)
	if err != nil {
		return Story{}, err
	}
	return Story{Id: item.Id,
		Author:   item.Author,
		Title:    item.Title,
		Score:    item.Score,
		Comments: sb.fetchComments(item.Kids)}, nil
}

func (sb *MyStoryBuilder) fetchComments(ids []int) []Comment {
	var comments []Comment
	for _, childId := range ids {
		item, err := sb.client.Get(childId)
		if err != nil {
			return []Comment{}
		}
		comments = append(comments, Comment{
			Id:       item.Id,
			Author:   item.Author,
			Text:     item.Text,
			Parent:   item.Parent,
			Comments: sb.fetchComments(item.Kids),
		})
	}
	return comments
}
