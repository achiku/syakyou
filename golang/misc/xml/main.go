package main

// http://stackoverflow.com/questions/27246275/golang-hide-xml-parent-tag-if-empty
// http://tatsushid.github.io/blog/2014/09/go-structure-pointer-and-omitempty/

import (
	"encoding/xml"
	"fmt"
	"log"
)

// SOAPEnvelope envelope
type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    SOAPBody
	Header  *SOAPHeader `xml:",omitempty"`
}

// SOAPHeader header
type SOAPHeader struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Content interface{} `xml:",omitempty"`
}

// SOAPBody body
type SOAPBody struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

// UnmarshalXML unmarshal SOAPBody
func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}
	var (
		token    xml.Token
		err      error
		consumed bool
	)
Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}
		if token == nil {
			break
		}
		envelopeNameSpace := "http://schemas.xmlsoap.org/soap/envelope/"
		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError(
					"Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == envelopeNameSpace && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil
				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}
				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}
				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}
	return nil
}

// SOAPFault fault
type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	Code    string   `xml:"faultcode,omitempty"`
	String  string   `xml:"faultstring,omitempty"`
	Actor   string   `xml:"faultactor,omitempty"`
	Detail  string   `xml:"detail,omitempty"`
}

// MySOAPHeader soap header
type MySOAPHeader struct {
	XMLName  xml.Name `xml:"http://akirachiku.com/soap/ MySOAPHeader"`
	UserID   string   `xml:"UserId,omitempty"`
	Password string   `xml:"Password,omitempty"`
}

func (f *SOAPFault) Error() string {
	return f.String
}

// Name struct
type Name struct {
	XMLName xml.Name `xml:"name"`
	First   string   `xml:"first,omitempty"`
	Last    string   `xml:"last,omitempty"`
}

// Person struct
type Person struct {
	XMLName xml.Name `xml:"person"`
	ID      int      `xml:"id,omitempty"`
	Name    *Name
	Age     int `xml:"age,omitempty"`
}

func marshalPerson(p Person) {
	buf, _ := xml.MarshalIndent(p, "", "  ")
	log.Printf("%+v", p)
	log.Printf("\n%s", string(buf))
	return
}

func unmarshalPerson(xmldoc []byte) {
	p := Person{}
	xml.Unmarshal(xmldoc, &p)
	log.Println(p)
	return
}

func main() {
	fmt.Printf("Hello, world\n")
}
