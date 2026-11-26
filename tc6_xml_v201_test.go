package plcopen_test

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/suifei/plcopen-go"
)

func TestV201ProjectIndexAndExactRoundTrip(t *testing.T) {
	data, err := os.ReadFile("testdata/v201-project.xml")
	if err != nil {
		t.Fatal(err)
	}
	original := bytes.Clone(data)
	p, err := plcopen.ParseProjectV201(data)
	if err != nil {
		t.Fatal(err)
	}

	want := plcopen.ProjectV201Info{
		Name: "V201Example", CompanyName: "Example", ProductName: "Demo", ProductVersion: "1",
		CreationDate: "2026-11-26T00:00:00Z",
		DataTypes:    []string{"Count"},
		POUs: []plcopen.POUV201Info{
			{Name: "Counter", Type: "program", Languages: []string{"ST"}},
			{Name: "Process", Type: "program", Languages: []string{"FBD"}},
		},
		Configurations: []string{"Default"},
	}
	if got := p.Info(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Info() = %#v, want %#v", got, want)
	}
	data[0] = '!'
	p.XML()[0] = '!'
	summary := p.Info()
	summary.DataTypes[0] = "changed"
	summary.Configurations[0] = "changed"
	summary.POUs[0].Languages[0] = "changed"
	if !reflect.DeepEqual(p.Info(), want) {
		t.Fatal("Info aliases the project's index")
	}
	var b bytes.Buffer
	n, err := p.WriteTo(&b)
	if err != nil || n != int64(len(original)) || !bytes.Equal(b.Bytes(), original) {
		t.Fatalf("WriteTo = %d bytes, %v; input bytes not preserved", n, err)
	}
	if _, err := p.WriteTo(shortWriter{}); err != io.ErrShortWrite {
		t.Fatalf("short write returned %v", err)
	}
}

type shortWriter struct{}

func (shortWriter) Write(p []byte) (int, error) { return len(p) - 1, nil }

func TestV201ProjectNamespaceAndDocumentBoundaries(t *testing.T) {
	valid := `<project xmlns="http://www.plcopen.org/xml/tc6_0201"/>`
	if _, err := plcopen.ParseProjectV201(append([]byte{0xef, 0xbb, 0xbf}, []byte(valid)...)); err != nil {
		t.Fatalf("UTF-8 BOM: %v", err)
	}
	for _, input := range []string{
		`<project xmlns="http://www.plcopen.org/xml/tc6.xsd"/>`,
		`<project xmlns="http://www.plcopen.org/xml/tc6_0200"/>`,
		`<project/>`,
		`<project xmlns="http://www.plcopen.org/xml/tc6_0201"><types></project>`,
		valid + valid,
		valid + "junk",
		"<!DOCTYPE project>" + valid,
		"<!-- only a comment -->",
	} {
		if _, err := plcopen.ParseProjectV201([]byte(input)); err == nil {
			t.Errorf("accepted %q", input)
		}
	}
}

func TestV201OfficialSchemaAndOpenPLCExamples(t *testing.T) {
	xsd := os.Getenv("PLCOPEN_V201_XSD")
	if xsd == "" {
		t.Skip("set PLCOPEN_V201_XSD to test against the official PLCopen schema")
	}
	type sample struct {
		path  string
		pous  []string
		langs []string
	}
	files := []sample{
		{filepath.Join("testdata", "v201-project.xml"), []string{"Counter", "Process"}, []string{"ST", "FBD"}},
	}
	if base := os.Getenv("PLCOPEN_OPENPLC_FIXTURES"); base != "" {
		files = append(files,
			sample{filepath.Join(base, "openplc-traffic-fbd.xml"), []string{"Trafic_light"}, []string{"FBD"}},
			sample{filepath.Join(base, "openplc-blink.xml"), []string{"Blink"}, []string{"LD"}},
			sample{filepath.Join(base, "openplc-multi-language.xml"), []string{"AverageVal", "plc_prg", "CounterST", "CounterFBD", "CounterSFC", "CounterIL", "CounterLD"}, []string{"ST", "FBD", "ST", "FBD", "SFC", "IL", "LD"}},
		)
	}
	for _, file := range files {
		t.Run(filepath.Base(file.path), func(t *testing.T) {
			data, err := os.ReadFile(file.path)
			if err != nil {
				t.Fatal(err)
			}
			p, err := plcopen.ParseProjectV201(data)
			if err != nil {
				t.Fatal(err)
			}
			var names, langs []string
			for _, pou := range p.Info().POUs {
				names = append(names, pou.Name)
				langs = append(langs, pou.Languages...)
			}
			if !reflect.DeepEqual(names, file.pous) || !reflect.DeepEqual(langs, file.langs) {
				t.Fatalf("POUs/languages = %v/%v, want %v/%v", names, langs, file.pous, file.langs)
			}
			roundTrip := p.XML()
			if !bytes.Equal(data, roundTrip) {
				t.Fatal("round trip changed XML")
			}
			out := filepath.Join(t.TempDir(), "roundtrip.xml")
			if err := os.WriteFile(out, roundTrip, 0600); err != nil {
				t.Fatal(err)
			}
			if msg, err := exec.Command("xmllint", "--nonet", "--noout", "--schema", xsd, out).CombinedOutput(); err != nil {
				t.Fatalf("official XSD validation: %v\n%s", err, msg)
			}
		})
	}
}
