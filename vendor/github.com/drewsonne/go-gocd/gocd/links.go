package gocd

import (
	"encoding/json"
	"net/url"
	"strings"
)

type linkField map[string]map[string]string
type linkHref struct {
	H string `json:"href"`
}

// HALLinks describes a collection of HALLinks
type HALLinks struct {
	links []*HALLink
}

// Add a link
func (al *HALLinks) Add(link *HALLink) {
	al.links = append(al.links, link)
}

// Get a HALLink by name
func (al HALLinks) Get(name string) *HALLink {
	link, _ := al.GetOk(name)
	return link
}

// GetOk a HALLink by name, and if it doesn't exist, return false
func (al HALLinks) GetOk(name string) (*HALLink, bool) {
	for _, link := range al.links {
		if strings.ToLower(link.Name) == strings.ToLower(name) {
			return link, true
		}
	}
	return nil, false
}

// Keys returns a string list of link names
func (al HALLinks) Keys() []string {
	keys := make([]string, len(al.links))
	for i, l := range al.links {
		keys[i] = l.Name
	}
	return keys
}

// MarshallJSON allows the encoding of links into JSON
func (al HALLinks) MarshallJSON() ([]byte, error) {
	ls := map[string]*linkHref{}
	for _, link := range al.links {
		ls[link.Name] = &linkHref{H: link.URL.String()}
	}
	return json.Marshal(ls)
}

// UnmarshalJSON allows the decoding of links from JSON
func (al *HALLinks) UnmarshalJSON(j []byte) (e error) {
	var d linkField
	var u *url.URL
	if e := json.Unmarshal(j, &d); e != nil {
		return e
	}

	for linkName, value := range d {
		if u, e = url.Parse(value["href"]); e != nil {
			return e
		}
		al.Add(&HALLink{
			Name: linkName,
			URL:  u,
		})
	}
	return nil
}
