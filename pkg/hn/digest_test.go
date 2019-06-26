package hn_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"workshop-starter/pkg/hn"
)

func TestDigest(t *testing.T) {
	t.Run("empty stories", func(t *testing.T) {
		mail := hn.BuildDigest("m@olx.pl", []hn.Story{})
		assert.Equal(t, "m@olx.pl", mail.Email)
		assert.Equal(t, "digest", mail.Template)
		assert.Equal(t, 1, len(mail.Placeholders))
		assert.Equal(t, "0", mail.Placeholders["entriesNumber"])
	})

	t.Run("build digest with stories without comments", func(t *testing.T) {
		mail := hn.BuildDigest("m@olx.pl", []hn.Story{
			{Title: "Story 1", Id: 1},
			{Title: "Story 2", Id: 2},
			{Title: "Story 3", Id: 3},
		})
		assert.Equal(t, "m@olx.pl", mail.Email)
		assert.Equal(t, "digest", mail.Template)
		assert.Equal(t, 4, len(mail.Placeholders))
		assert.Equal(t, "3", mail.Placeholders["entriesNumber"])
		assert.Equal(t, "Story 1 (0)", mail.Placeholders["1"])
		assert.Equal(t, "Story 2 (0)", mail.Placeholders["2"])
		assert.Equal(t, "Story 3 (0)", mail.Placeholders["3"])
	})

	t.Run("build digest with stories with comments", func(t *testing.T) {
		mail := hn.BuildDigest("m@olx.pl", []hn.Story{
			{Title: "Story 1", Id: 1, Comments: []hn.Comment{
				{Id: 4},
				{Id: 5},
			}},
		})
		assert.Equal(t, "m@olx.pl", mail.Email)
		assert.Equal(t, "digest", mail.Template)
		assert.Equal(t, 2, len(mail.Placeholders))
		assert.Equal(t, "1", mail.Placeholders["entriesNumber"])
		assert.Equal(t, "Story 1 (2)", mail.Placeholders["1"])
	})

	t.Run("build digest with stories with nested comments", func(t *testing.T) {
		mail := hn.BuildDigest("m@olx.pl", []hn.Story{
			{Title: "Story 1", Id: 1, Comments: []hn.Comment{
				{Id: 4, Comments: []hn.Comment{
					{Id: 5},
				}},
			}},
		})
		assert.Equal(t, "m@olx.pl", mail.Email)
		assert.Equal(t, "digest", mail.Template)
		assert.Equal(t, 2, len(mail.Placeholders))
		assert.Equal(t, "1", mail.Placeholders["entriesNumber"])
		assert.Equal(t, "Story 1 (2)", mail.Placeholders["1"])
	})

	t.Run("build digest with stories with nested comments", func(t *testing.T) {
		mail := hn.BuildDigest("m@olx.pl", []hn.Story{
			{Title: "Story 1", Id: 1, Comments: []hn.Comment{
				{Id: 4, Comments: []hn.Comment{
					{Id: 5, Comments: []hn.Comment{
						{Id: 6},
					}},
				}},
			}},
		})
		assert.Equal(t, "m@olx.pl", mail.Email)
		assert.Equal(t, "digest", mail.Template)
		assert.Equal(t, 2, len(mail.Placeholders))
		assert.Equal(t, "1", mail.Placeholders["entriesNumber"])
		assert.Equal(t, "Story 1 (3)", mail.Placeholders["1"])
	})
}
