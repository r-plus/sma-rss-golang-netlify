package main

import (
	"encoding/xml"
	"testing"
)

func TestGeneratedAtomFeedUnmarshalable(t *testing.T) {
	// MONOEYES, HIATUS
	ids := []string{"351", "237"}
	for _, v := range ids {
		smaAtomFeedTestUtil(v, t)
	}
}

func smaAtomFeedTestUtil(id string, t *testing.T) error {
	s, nwErr := makeSMAAtomFeed(id)
	if nwErr != nil {
		// ignore network error, use original feed data to daily test.
		return nil
	}
	var v interface{}

	err := xml.Unmarshal([]byte(s), &v)
	if err != nil {
		t.Logf("%s feed: %#v", id, err)
		t.Fail()
		return err
	}
	return nil
}
