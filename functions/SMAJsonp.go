package main

// SMAJsonp jsonp top object
type SMAJsonp struct {
	Info SMAInfo
}

// SMAInfo info
type SMAInfo struct {
	Items []SMAItem `json:"item"`
}

// SMAItem item
type SMAItem struct {
	ID      string
	Title   string
	PubDate string
	Text    string
}
