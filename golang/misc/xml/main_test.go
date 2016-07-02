package main

import (
	"bytes"
	"encoding/xml"
	"testing"
)

func TestOmitempty(t *testing.T) {
	p1 := Person{
		ID:  1,
		Age: 33,
	}
	p2 := Person{
		ID: 2,
		Name: &Name{
			First: "Ichiro",
			Last:  "Suzuki",
		},
	}
	ps := []Person{p1, p2}
	for _, p := range ps {
		buf, _ := xml.MarshalIndent(p, "", "  ")
		t.Logf("\n%s", string(buf))
	}
}

func TestSOAPEnvelopeMarshal(t *testing.T) {
	e := SOAPEnvelope{
		Header: &SOAPHeader{
			Content: MySOAPHeader{
				UserID:   "achiku",
				Password: "pass",
			},
		},
		Body: SOAPBody{
			Content: Person{
				ID: 1,
				Name: &Name{
					First: "Akira",
					Last:  "Chiku",
				},
				Age: 30,
			},
		},
	}
	buf, err := xml.MarshalIndent(e, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("\n%s", string(buf))
}

func TestUnmarshal(t *testing.T) {
	d1 := []byte(`
	<person>
	  <name>
		<first>Taro</first>
		<last>Suzuki</last>
	  </name>
	  <age>30</age>
	</person>
	`)
	d2 := []byte(`
	<person>
	  <age>30</age>
	</person>
	`)
	d3 := []byte(`
	<person>
	  <id>3</id>
	</person>
	`)
	docs := [][]byte{d1, d2, d3}
	for _, d := range docs {
		p := Person{}
		xml.Unmarshal(d, &p)
		t.Logf("%+v", p)
	}
}

func TestNoHeaderSOAPEnvelopeMarshal(t *testing.T) {
	e := SOAPEnvelope{
		Body: SOAPBody{
			Content: Person{
				ID: 1,
				Name: &Name{
					First: "Akira",
					Last:  "Chiku",
				},
				Age: 30,
			},
		},
	}
	buf, err := xml.MarshalIndent(e, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("\n%s", string(buf))
}

func TestNoHeaderSOAPEnvelopeUnmarshal(t *testing.T) {
	xmldoc := []byte(`
    <Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
      <Body xmlns="http://schemas.xmlsoap.org/soap/envelope/">
        <person>
          <id>1</id>
          <name>
            <first>Akira</first>
            <last>Chiku</last>
          </name>
          <age>30</age>
        </person>
      </Body>
    </Envelope>
	`)
	p := new(Person)
	res := SOAPEnvelope{}
	res.Body = SOAPBody{Content: p}
	if err := xml.Unmarshal(xmldoc, &res); err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)
	t.Logf("%+v", p)
}

func TestWithHeaderSOAPEnvelopeUnmarshal(t *testing.T) {
	xmldoc := []byte(`
    <Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
      <Body xmlns="http://schemas.xmlsoap.org/soap/envelope/">
        <person>
          <id>1</id>
          <name>
            <first>Akira</first>
            <last>Chiku</last>
          </name>
          <age>30</age>
        </person>
      </Body>
      <Header xmlns="http://schemas.xmlsoap.org/soap/envelope/">
        <MySOAPHeader xmlns="http://akirachiku.com/soap/">
          <UserId>achiku</UserId>
          <Password>pass</Password>
        </MySOAPHeader>
      </Header>
    </Envelope>
	`)
	p := new(Person)
	h := new(MySOAPHeader)
	res := SOAPEnvelope{}
	res.Body = SOAPBody{Content: p}
	res.Header = &SOAPHeader{Content: h}
	if err := xml.Unmarshal(xmldoc, &res); err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)
	t.Logf("%+v", res.Header.Content)
	t.Logf("%+v", p)
	t.Logf("%+v", p.Name)
	t.Logf("%s, %s", h.UserID, h.Password)
}

func TestNoEnvelopeSOAPEnvelopeUnmarshal(t *testing.T) {
	xmldoc := []byte(`
    <Body xmlns="http://schemas.xmlsoap.org/soap/envelope/">
      <person>
        <id>1</id>
        <name>
          <first>Akira</first>
          <last>Chiku</last>
        </name>
        <age>30</age>
      </person>
    </Body>
	`)
	p := new(Person)
	b := SOAPBody{Content: p}
	if err := xml.Unmarshal(xmldoc, &b); err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", b)
	t.Logf("%+v", p)
}

func TestTraverseXML(t *testing.T) {
	xmldoc := []byte(`
    <Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
      <Body xmlns="http://schemas.xmlsoap.org/soap/envelope/">
        <person>
          <id>1</id>
          <name>
            <first>Akira</first>
            <last>Chiku</last>
          </name>
          <age>30</age>
        </person>
      </Body>
      <Header xmlns="http://schemas.xmlsoap.org/soap/envelope/">
        <MySOAPHeader xmlns="http://akirachiku.com/soap/">
          <UserId>achiku</UserId>
          <Password>pass</Password>
        </MySOAPHeader>
      </Header>
    </Envelope>
	`)
	buf := bytes.NewBuffer(xmldoc)
	dec := xml.NewDecoder(buf)

	var n Node
	err := dec.Decode(&n)
	if err != nil {
		t.Fatal(err)
	}

	walk([]Node{n}, func(n Node) bool {
		if len(n.Nodes) > 0 {
			if n.XMLName.Local == "Body" {
				t.Logf("%s", n.Nodes[0].XMLName.Local)
				return false
			}
		}
		return true
	})
}

func TestGetSOAPActionType(t *testing.T) {
	personDoc := []byte(`
    <Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
      <Body xmlns="http://schemas.xmlsoap.org/soap/envelope/">
        <person>
          <id>1</id>
        </person>
      </Body>
    </Envelope>
	`)
	cardDoc := []byte(`
    <Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
      <Body xmlns="http://schemas.xmlsoap.org/soap/envelope/">
        <card>
          <id>1</id>
        </card>
      </Body>
    </Envelope>
	`)

	// this is not really valid SOAP request
	irregularDoc := []byte(`
    <Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
      <Body xmlns="http://schemas.xmlsoap.org/soap/envelope/">
        <barcode>
          <id>1</id>
        </barcode>
        <authcode>
          <id>1</id>
        </authcode>
      </Body>
    </Envelope>
	`)

	data := []struct {
		input    []byte
		expected string
	}{
		{input: personDoc, expected: "person"},
		{input: cardDoc, expected: "card"},
		{input: irregularDoc, expected: "barcode"},
	}

	for _, d := range data {
		name, err := getSOAPActionType(d.input)
		if err != nil {
			t.Error(err)
		}
		t.Log(name)
		if name != d.expected {
			t.Errorf("want %s got %s", d.expected, name)
		}
	}
}
