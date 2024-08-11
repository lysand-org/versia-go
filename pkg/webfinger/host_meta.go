package webfinger

import (
	"encoding/json"
	"net/url"
)

type HostMeta struct {
	JSON []byte
	XML  []byte
}

func NewHostMeta(baseURL *url.URL) HostMeta {
	template := &url.URL{Path: "/.well-known/webfinger?resource={uri}"}
	template = baseURL.ResolveReference(template)

	return HostMeta{
		JSON: generateJSONHostMeta(template),
		XML:  generateXMLHostMeta(template),
	}
}

func generateXMLHostMeta(template *url.URL) []byte {
	return []byte(`<?xml version="1.0"?>
	<XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0">
	<Link rel="lrdd" template="` + template.String() + `" />
	</XRD>`)
}

func generateJSONHostMeta(template *url.URL) []byte {
	b, err := json.Marshal(hostMetaStruct{
		Links: []hostMetaLink{{
			Rel:      "lrdd",
			Template: template.String(),
		}},
	})
	if err != nil {
		panic(err)
	}

	return b
}

type hostMetaStruct struct {
	Links []hostMetaLink `json:"links"`
}

type hostMetaLink struct {
	Rel      string `json:"rel"`
	Template string `json:"template"`
}
