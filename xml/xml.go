package main

import (
	"encoding/xml"
	"fmt"
)

// Similarly to the JSON examples, field tags contain directives for the encoder and decoder.
// The XMLName field name dictates the name of the XML element representing this struct;
// id,attr means that the Id field is an XML attribute rather than a nested element.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	// Using MarshalIndent to produce a more human-readable output.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))
	// To add a generic XML header to the output, append it explicitly.
	fmt.Println(xml.Header + string(out))

	var p Plant
	// Use Unmarshal to parse a stream of bytes with XML into a data structure.
	// If the XML is malformed or cannot be mapped onto Plant, a descriptive error will be returned.
	if err := xml.Unmarshal(out, &p); err != err {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}
	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
