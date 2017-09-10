// Code generated by "gocd-response-links-generator -type=PluginsResponseLinks,PluginLinks"; DO NOT EDIT.
package gocd

import "encoding/json"

func (l PluginsResponseLinks) MarshalJSON() ([]byte, error) {
	ls := struct {
		Self *linkHref `json:"self,omitempty"`
		Doc  *linkHref `json:"doc,omitempty"`
	}{}
	if l.Self != nil {
		ls.Self = &linkHref{H: l.Self.String()}
	}
	if l.Doc != nil {
		ls.Doc = &linkHref{H: l.Doc.String()}
	}
	j, e := json.Marshal(ls)
	if e != nil {
		return nil, e
	}
	return j, nil
}
func (l *PluginsResponseLinks) UnmarshalJSON(j []byte) error {
	var d linkField
	if e := json.Unmarshal(j, &d); e != nil {
		return e
	}
	if err := unmarshallLinkField(d, "self", &l.Self); err != nil {
		return err
	}
	if err := unmarshallLinkField(d, "doc", &l.Doc); err != nil {
		return err
	}
	return nil
}
func (l PluginLinks) MarshalJSON() ([]byte, error) {
	ls := struct {
		Self *linkHref `json:"self,omitempty"`
		Doc  *linkHref `json:"doc,omitempty"`
		Find *linkHref `json:"find,omitempty"`
	}{}
	if l.Self != nil {
		ls.Self = &linkHref{H: l.Self.String()}
	}
	if l.Doc != nil {
		ls.Doc = &linkHref{H: l.Doc.String()}
	}
	if l.Find != nil {
		ls.Find = &linkHref{H: l.Find.String()}
	}
	j, e := json.Marshal(ls)
	if e != nil {
		return nil, e
	}
	return j, nil
}
func (l *PluginLinks) UnmarshalJSON(j []byte) error {
	var d linkField
	if e := json.Unmarshal(j, &d); e != nil {
		return e
	}
	if err := unmarshallLinkField(d, "self", &l.Self); err != nil {
		return err
	}
	if err := unmarshallLinkField(d, "doc", &l.Doc); err != nil {
		return err
	}
	if err := unmarshallLinkField(d, "find", &l.Find); err != nil {
		return err
	}
	return nil
}