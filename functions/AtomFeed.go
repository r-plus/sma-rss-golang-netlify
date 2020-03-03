package main

import (
	"fmt"
	"html"
	"strings"
	"time"
)

// AtomFeed is feed generator
type AtomFeed struct {
	ArtistID string
	Info     SMAInfo
}

// MakeFeed generate atom rss feed xml string
func (a AtomFeed) MakeFeed() string {
	return a.atomMakeHeader() + a.atomMakeRootTag() + a.atomMakeEntries() + a.atomMakeEndTag()
}

func (a AtomFeed) atomMakeHeader() string {
	return `<?xml version='1.0' encoding='UTF-8'?>
<feed xmlns='http://www.w3.org/2005/Atom' xml:lang='ja'>`
}

func (a AtomFeed) atomMakeRootTag() string {
	item := a.Info.Items[0]
	return fmt.Sprintf(`<id>sma %s feed</id>
        <title>sma %s feed</title>
        <updated>%s</updated>`, a.ArtistID, a.ArtistID, a.entryPubDateString(item))
}

// escape html and SMA feed injected control char like U+0008 Backspace.
func escapeString(s string) string {
	ss := strings.ReplaceAll(s, "\b", "")
	return html.EscapeString(ss)
	// return html.EscapeString(s)
}

func (a AtomFeed) atomMakeEntries() string {
	entries := make([]string, len(a.Info.Items))
	for i, v := range a.Info.Items {
		date := a.entryPubDateString(v)
		title := escapeString(v.Title)
		html := escapeString(v.Text)
		entries[i] = fmt.Sprintf(`<entry>
          <id>%s</id>
          <title>%s</title>
          <updated>%s</updated>
          <summary type="html">%s</summary>
          <content type="html">%s</content>
      </entry>`, v.ID, title, date, html, html)
	}
	return strings.Join(entries, "\n")
}

func (a AtomFeed) atomMakeEndTag() string {
	return "</feed>"
}

func (a AtomFeed) entryPubDateString(item SMAItem) string {
	d, _ := time.Parse("2006/01/02", item.PubDate)
	return d.Format(time.RFC3339)
}
