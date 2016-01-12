package epub

import (
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
)

type Book struct {
	MimeType  string
	Container *Container
}

type Container struct {
	XMLName   xml.Name   `xml:"container"`
	RootFiles []RootFile `xml:"rootfiles>rootfile"`
}

type RootFile struct {
	FullPath  string `xml:"full-path,attr"`
	MediaType string `xml:"media-type,attr"`
	Opf       *Opf
}

type Opf struct {
	XMLName  xml.Name `xml:"package"`
	Metadata Metadata `xml:"metadata"`
	Manifest Manifest `xml:"manifest"`
	Spine    Spine    `xml:"spine"`
}

type Metadata struct {
	Title      string `xml:"title"`
	Creator    string `xml:"creator"`
	Language   string `xml:"language"`
	Identifier string `xml:"identifier"`
	Subject    string `xml:"subject"`
	Publisher  string `xml:"publisher"`
	Date       string `xml:"date"`
}

type Manifest struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Id        string `xml:"id,attr"`
	Href      string `xml:"href,attr"`
	MediaType string `xml:"media-type,attr"`
}

type Spine struct {
	Toc      string    `xml:"toc,attr"`
	ItemRefs []ItemRef `xml:"itemref"`
}

func (p *Spine) Index() int {
	for k, v := range p.ItemRefs {
		if v.Linear != "no" {
			return k
		}
	}
	return 0
}

func (p *Spine) Cover() int {
	for k, v := range p.ItemRefs {
		if v.Linear == "no" {
			return k
		}
	}
	return 0
}

type ItemRef struct {
	IdRef  string `xml:"idref,attr"`
	Linear string `xml:"linear,attr"`
}

func readRootFile(target string, file RootFile) (*Opf, error) {
	buf, err := ioutil.ReadFile(filepath.Join(target, file.FullPath))
	if err != nil {
		return nil, err
	}
	var opf Opf
	if err = xml.Unmarshal(buf, &opf); err != nil {
		return nil, err
	}
	return &opf, nil
}

func readContainer(target string) (*Container, error) {

	buf, err := ioutil.ReadFile(filepath.Join(target, "META-INF", "container.xml"))
	if err != nil {
		return nil, err
	}
	var ct Container
	if err := xml.Unmarshal(buf, &ct); err != nil {
		return nil, err
	}

	return &ct, nil
}
