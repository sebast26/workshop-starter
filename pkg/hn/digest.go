package hn

import (
	"fmt"
	"strconv"
)

type Mail struct {
	Email        string
	Template     string
	Placeholders map[string]string
}

type MailSender interface {
	Send(mail Mail) error
}

func BuildDigest(email string, stories []Story) Mail {
	return Mail{Email: email, Template: "digest", Placeholders: buildPlaceholders(stories)}
}

func buildPlaceholders(stories []Story) map[string]string {
	placeholders := make(map[string]string, len(stories)+1)
	placeholders["entriesNumber"] = strconv.Itoa(len(stories))

	for _, story := range stories {
		count := countComments(story)
		placeholders[strconv.Itoa(story.Id)] = fmt.Sprintf("%s (%d)", story.Title, count)
	}

	return placeholders
}

func countComments(story Story) int {
	count := len(story.Comments)
	for _, c := range story.Comments {
		count += countCommentsOnComment(c)
	}
	return count
}

func countCommentsOnComment(comment Comment) int {
	count := len(comment.Comments)
	for _, c := range comment.Comments {
		count += countCommentsOnComment(c)
	}
	return count
}
