package plcopen

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
)

// NamespaceV201 is the PLCopen TC6 XML 2.01 project namespace. It is
// distinct from both TC6 V1.0B and IEC 61131-10.
const NamespaceV201 = "http://www.plcopen.org/xml/tc6_0201"

// ProjectV201 is an immutable UTF-8 V2.01 project document. It keeps the
// complete XML, including vendor addData, XHTML and namespace prefixes.
// The legacy Project type and its XML/JSON behavior are unchanged.
type ProjectV201 struct {
	data []byte
	info ProjectV201Info
}

// ProjectV201Info indexes the main project sections. It is not an editable
// replacement for the original document or a complete generated XSD model.
type ProjectV201Info struct {
	Name           string
	CompanyName    string
	ProductName    string
	ProductVersion string
	CreationDate   string
	DataTypes      []string
	POUs           []POUV201Info
	Configurations []string
}

// POUV201Info identifies a POU and its top-level bodies in document order.
// Nested action/transition bodies remain available in the original XML.
type POUV201Info struct {
	Name      string
	Type      string
	Languages []string
}

var projectV201Name = xml.Name{Space: NamespaceV201, Local: "project"}

type v201TextElement struct {
	XMLName xml.Name
}

type v201Body struct {
	Children []v201TextElement `xml:",any"`
}

type v201POU struct {
	Name   string     `xml:"name,attr"`
	Type   string     `xml:"pouType,attr"`
	Bodies []v201Body `xml:"http://www.plcopen.org/xml/tc6_0201 body"`
}

type v201Named struct {
	Name string `xml:"name,attr"`
}

type v201Index struct {
	FileHeader struct {
		CompanyName    string `xml:"companyName,attr"`
		ProductName    string `xml:"productName,attr"`
		ProductVersion string `xml:"productVersion,attr"`
		CreationDate   string `xml:"creationDateTime,attr"`
	} `xml:"http://www.plcopen.org/xml/tc6_0201 fileHeader"`
	ContentHeader v201Named `xml:"http://www.plcopen.org/xml/tc6_0201 contentHeader"`
	Types         struct {
		DataTypes struct {
			Entries []v201Named `xml:"http://www.plcopen.org/xml/tc6_0201 dataType"`
		} `xml:"http://www.plcopen.org/xml/tc6_0201 dataTypes"`
		POUs struct {
			Entries []v201POU `xml:"http://www.plcopen.org/xml/tc6_0201 pou"`
		} `xml:"http://www.plcopen.org/xml/tc6_0201 pous"`
	} `xml:"http://www.plcopen.org/xml/tc6_0201 types"`
	Instances struct {
		Configurations struct {
			Entries []v201Named `xml:"http://www.plcopen.org/xml/tc6_0201 configuration"`
		} `xml:"http://www.plcopen.org/xml/tc6_0201 configurations"`
	} `xml:"http://www.plcopen.org/xml/tc6_0201 instances"`
}

// ParseProjectV201 reads exactly one well-formed V2.01 project. It checks
// the root namespace and rejects directives; it does not validate the XSD.
// It preserves the original bytes and does not compile or execute PLC code.
func ParseProjectV201(data []byte) (*ProjectV201, error) {
	dec := xml.NewDecoder(bytes.NewReader(bytes.TrimPrefix(data, []byte{0xef, 0xbb, 0xbf})))
	depth, seen := 0, false
	for {
		token, err := dec.Token()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("PLCopen 2.01 XML: %w", err)
		}
		switch token := token.(type) {
		case xml.StartElement:
			if depth == 0 {
				if seen {
					return nil, errors.New("PLCopen 2.01 XML has multiple roots")
				}
				if token.Name != projectV201Name {
					return nil, fmt.Errorf("expected PLCopen 2.01 project {%s}project, got {%s}%s",
						NamespaceV201, token.Name.Space, token.Name.Local)
				}
				seen = true
			}
			depth++
		case xml.EndElement:
			depth--
		case xml.CharData:
			if depth == 0 && len(bytes.Trim(token, " \t\r\n")) != 0 {
				return nil, errors.New("PLCopen 2.01 XML has text outside the project")
			}
		case xml.Directive:
			return nil, errors.New("PLCopen 2.01 XML directives are not supported")
		}
	}
	if !seen || depth != 0 {
		return nil, errors.New("PLCopen 2.01 XML has no complete project root")
	}

	var index v201Index
	if err := xml.Unmarshal(bytes.TrimPrefix(data, []byte{0xef, 0xbb, 0xbf}), &index); err != nil {
		return nil, fmt.Errorf("PLCopen 2.01 project index: %w", err)
	}
	info := ProjectV201Info{
		Name: index.ContentHeader.Name, CompanyName: index.FileHeader.CompanyName,
		ProductName: index.FileHeader.ProductName, ProductVersion: index.FileHeader.ProductVersion,
		CreationDate: index.FileHeader.CreationDate,
	}
	for _, item := range index.Types.DataTypes.Entries {
		info.DataTypes = append(info.DataTypes, item.Name)
	}
	for _, item := range index.Types.POUs.Entries {
		pou := POUV201Info{Name: item.Name, Type: item.Type}
		for _, body := range item.Bodies {
			for _, child := range body.Children {
				if child.XMLName.Space == NamespaceV201 {
					switch child.XMLName.Local {
					case "IL", "ST", "FBD", "LD", "SFC":
						pou.Languages = append(pou.Languages, child.XMLName.Local)
					}
				}
			}
		}
		info.POUs = append(info.POUs, pou)
	}
	for _, item := range index.Instances.Configurations.Entries {
		info.Configurations = append(info.Configurations, item.Name)
	}
	return &ProjectV201{data: bytes.Clone(data), info: info}, nil
}

// Info returns a detached index. Editing it does not change the document.
func (p *ProjectV201) Info() ProjectV201Info {
	out := p.info
	out.DataTypes = append([]string(nil), out.DataTypes...)
	out.Configurations = append([]string(nil), out.Configurations...)
	out.POUs = append([]POUV201Info(nil), out.POUs...)
	for i := range out.POUs {
		out.POUs[i].Languages = append([]string(nil), out.POUs[i].Languages...)
	}
	return out
}

// XML returns a copy of the exact original XML.
func (p *ProjectV201) XML() []byte { return bytes.Clone(p.data) }

// WriteTo writes the exact original XML.
func (p *ProjectV201) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(p.data)
	if err == nil && n != len(p.data) {
		err = io.ErrShortWrite
	}
	return int64(n), err
}
