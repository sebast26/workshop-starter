package hn

import (
	"github.com/davecgh/go-spew/spew"
	"io"
)

type NewDumper struct {
	client MyClient
}

func NewDump(client MyClient) *NewDumper {
	return &NewDumper{client}
}

func (d NewDumper) Dump(writer io.Writer) error {

	maxItemId, _ := d.client.Maxitem()
	for i := maxItemId; (i > maxItemId-2) && (i > 0); i-- {
		spew.Dump(i)
		item, _ := d.client.Get(i)
		if item.Type == "comment" {
			_, err := writer.Write([]byte(item.Text + "\n"))
			if err != nil {
				return err
			}
		}
		_, err := writer.Write([]byte(item.Title + "\n"))
		if err != nil {
			return err
		}

	}

	return nil
}
