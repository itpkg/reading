package sitemap

import (
	"encoding/xml"
	"fmt"
	"io"
	"time"
)

type Handler func() []*Url
type TimeStr string
type PriorityStr string
type ChangeFreqStr string

const (
	Always  = ChangeFreqStr("always")
	Hourly  = ChangeFreqStr("hourly")
	Daily   = ChangeFreqStr("daily")
	Weekly  = ChangeFreqStr("weekly")
	Monthly = ChangeFreqStr("monthly")
	Yearly  = ChangeFreqStr("yearly")
	Never   = ChangeFreqStr("never")
)

func Priority(p float32) PriorityStr {
	return PriorityStr(fmt.Sprintf("%.1f", p))
}

func Time(t time.Time) TimeStr {
	return TimeStr(t.Format("2006-01-02"))
}

type UrlSet struct {
	XMLName xml.Name `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
	Urls    []*Url   `xml:"url"`
}

type Url struct {
	Loc        string        `xml:"loc"`
	LastMod    TimeStr       `xml:"lastmod"`
	ChangeFreq ChangeFreqStr `xml:"changefreq"`
	Priority   PriorityStr   `xml:"priority"`
}

func Xml(wrt io.Writer, handlers ...Handler) error {
	si := &UrlSet{
		Urls: make([]*Url, 0),
	}
	for _, fn := range handlers {
		urls := fn()
		si.Urls = append(si.Urls, urls...)
	}
	_, err := wrt.Write([]byte(xml.Header))
	if err != nil {
		return err
	}
	en := xml.NewEncoder(wrt)
	//en.Indent("", "  ")
	return en.Encode(si)
}
