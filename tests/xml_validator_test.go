package tests

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/suifei/plcopen-go"
	"github.com/suifei/plcopen-go/utils"
)

// TestXMLSchemaValidation validates generated XML against PLCopen XSD schema
func TestXMLSchemaValidation(t *testing.T) {
	// First, generate XML from our test cases
	project := createTestProject()

	t.Log(string(utils.MustToJSONIndent(project)))

	xmlData, err := xml.MarshalIndent(project, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal test project: %v", err)
	}
	// Write XML to temporary file
	xmlFile := filepath.Join(os.TempDir(), "test_plcopen.xml")
	err = os.WriteFile(xmlFile, xmlData, 0644)
	if err != nil {
		t.Fatalf("Failed to write XML file: %v", err)
	}
	// defer os.Remove(xmlFile) // Temporarily disabled for debugging

	t.Logf("Generated XML file path: %s", xmlFile)

	// Get current directory
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	// Path to XSD file
	xsdFile := filepath.Join(currentDir, "TC6_XML_V10_B.xsd")

	// Check if XSD file exists
	if _, err := os.Stat(xsdFile); os.IsNotExist(err) {
		t.Skipf("XSD file not found: %s", xsdFile)
	}

	fmt.Printf("Validating XML against XSD...\n")
	fmt.Printf("XML file: %s\n", xmlFile)
	fmt.Printf("XSD file: %s\n", xsdFile)

	// Try to validate using xmllint if available
	if validateWithXmllint(t, xmlFile, xsdFile) {
		return
	}

	// If xmllint is not available, do basic XML well-formedness check
	t.Log("xmllint not available, performing basic XML well-formedness check")

	// Parse the XML to ensure it's well-formed
	var parsed plcopen.Project
	err = xml.Unmarshal(xmlData, &parsed)
	if err != nil {
		t.Errorf("Generated XML is not well-formed: %v", err)
	} else {
		t.Log("XML is well-formed")
	}

	// Check namespace
	if !checkXMLNamespace(string(xmlData)) {
		t.Error("XML does not contain proper PLCopen namespace")
	} else {
		t.Log("XML contains proper PLCopen namespace")
	}

	// Validate XML structure matches our expectations
	validateProjectStructure(t, &parsed)
}

// createTestProject creates a comprehensive test project for validation
func createTestProject() *plcopen.Project {
	return &plcopen.Project{
		FileHeader: &plcopen.ProjectFileHeader{
			CompanyName:        "Test Company",
			ProductName:        "Test Product",
			ProductVersion:     "1.0",
			ContentDescription: "PLCopen XML validation test",
		}, ContentHeader: &plcopen.ProjectContentHeader{
			Name:         "ValidationTestProject",
			Version:      "1.0",
			Organization: "Test Organization",
			Author:       "Test Author",
			Language:     "en",
			CoordinateInfo: &plcopen.ProjectContentHeaderCoordinateInfo{
				PageSize: &plcopen.ProjectContentHeaderCoordinateInfoPageSize{
					X: 210.0,
					Y: 297.0,
				},
				FBD: &plcopen.ProjectContentHeaderCoordinateInfoFBD{
					Scaling: &plcopen.ProjectContentHeaderCoordinateInfoFBDScaling{
						X: 1.0,
						Y: 1.0,
					},
				},
				LD: &plcopen.ProjectContentHeaderCoordinateInfoLD{
					Scaling: &plcopen.ProjectContentHeaderCoordinateInfoLDScaling{
						X: 1.0,
						Y: 1.0,
					},
				},
				SFC: &plcopen.ProjectContentHeaderCoordinateInfoSFC{
					Scaling: &plcopen.ProjectContentHeaderCoordinateInfoSFCScaling{
						X: 1.0,
						Y: 1.0,
					},
				},
			},
		},
		Types: &plcopen.ProjectTypes{
			DataTypes: []plcopen.ProjectTypesDataType{
				{
					Name: "TestEnum",
					BaseType: &plcopen.DataType{
						Enum: &plcopen.DataTypeEnum{
							Values: &plcopen.DataTypeEnumValues{
								Values: []plcopen.DataTypeEnumValuesValue{
									{Name: "RED"},
									{Name: "GREEN"},
									{Name: "BLUE"},
								},
							},
							BaseType: &plcopen.DataType{INT: &struct{}{}},
						},
					},
				},
			},
			POUs: []plcopen.ProjectTypesPOU{
				{
					Name:    "TestProgram",
					POUType: plcopen.POUTypeProgram,
					Interface: &plcopen.ProjectTypesPOUInterface{
						LocalVars: &plcopen.ProjectTypesPOUInterfaceLocalVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "counter",
									Type: &plcopen.DataType{INT: &struct{}{}},
									InitialValue: &plcopen.Value{
										SimpleValue: &plcopen.ValueSimpleValue{Value: "0"},
									},
								},
								{
									Name: "status",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
							},
						},
					}, Body: &plcopen.Body{
						ST: &plcopen.BodyST{
							XMLNSXhtml: "http://www.w3.org/1999/xhtml",
							Xhtml:      "<xhtml:p>counter := counter + 1;<xhtml:br/>status := counter &gt; 10;</xhtml:p>",
						},
					},
				},
				{
					Name:    "TestFunction",
					POUType: plcopen.POUTypeFunction,
					Interface: &plcopen.ProjectTypesPOUInterface{
						ReturnType: &plcopen.DataType{BOOL: &struct{}{}},
						InputVars: &plcopen.ProjectTypesPOUInterfaceInputVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "input1",
									Type: &plcopen.DataType{INT: &struct{}{}},
								},
							},
						},
					}, Body: &plcopen.Body{
						ST: &plcopen.BodyST{
							XMLNSXhtml: "http://www.w3.org/1999/xhtml",
							Xhtml:      "<xhtml:p>TestFunction := input1 &gt; 0;</xhtml:p>",
						},
					},
				},
			},
		},
		Instances: &plcopen.ProjectInstances{
			Configurations: []plcopen.ProjectInstancesConfiguration{
				{
					Name: "DefaultConfig", Resources: []plcopen.ProjectInstancesConfigurationResource{
						{
							Name: "DefaultResource",
							POUInstances: []plcopen.POUInstance{
								{
									Name:     "MainProgramInstance",
									TypeName: "TestProgram",
								},
							},
						},
					},
				},
			},
		},
	}
}

// createComprehensiveProject creates a project with comprehensive PLCopen XML coverage
func createComprehensiveProject() *plcopen.Project {
	return &plcopen.Project{
		FileHeader: &plcopen.ProjectFileHeader{
			CompanyName:        "Comprehensive Test Company",
			CompanyURL:         "https://test-company.com",
			ProductName:        "Test PLCopen Suite",
			ProductVersion:     "2.0",
			ProductRelease:     "Beta",
			ContentDescription: "Comprehensive PLCopen XML test project with all node types",
		},
		ContentHeader: &plcopen.ProjectContentHeader{
			Name:         "ComprehensiveTestProject",
			Version:      "2.0",
			Organization: "Test Engineering",
			Author:       "Test Engineer",
			Language:     "en-US",
			Comment:      "This project tests all PLCopen XML node types",
			CoordinateInfo: &plcopen.ProjectContentHeaderCoordinateInfo{
				PageSize: &plcopen.ProjectContentHeaderCoordinateInfoPageSize{
					X: 210.0,
					Y: 297.0,
				},
				FBD: &plcopen.ProjectContentHeaderCoordinateInfoFBD{
					Scaling: &plcopen.ProjectContentHeaderCoordinateInfoFBDScaling{
						X: 1.0,
						Y: 1.0,
					},
				},
				LD: &plcopen.ProjectContentHeaderCoordinateInfoLD{
					Scaling: &plcopen.ProjectContentHeaderCoordinateInfoLDScaling{
						X: 1.0,
						Y: 1.0,
					},
				},
				SFC: &plcopen.ProjectContentHeaderCoordinateInfoSFC{
					Scaling: &plcopen.ProjectContentHeaderCoordinateInfoSFCScaling{
						X: 1.0,
						Y: 1.0,
					},
				},
			},
		}, Types: &plcopen.ProjectTypes{
			DataTypes: []plcopen.ProjectTypesDataType{
				// Basic data types that are required by XSD
				{
					Name:     "BOOL",
					BaseType: &plcopen.DataType{BOOL: &struct{}{}},
				},
				{
					Name:     "INT",
					BaseType: &plcopen.DataType{INT: &struct{}{}},
				},
				{
					Name:     "REAL",
					BaseType: &plcopen.DataType{REAL: &struct{}{}},
				},
				{
					Name:     "STRING",
					BaseType: &plcopen.DataType{String: &plcopen.DataTypeString{}},
				},
				{
					Name: "MyEnum",
					BaseType: &plcopen.DataType{
						Enum: &plcopen.DataTypeEnum{
							Values: &plcopen.DataTypeEnumValues{
								Values: []plcopen.DataTypeEnumValuesValue{
									{Name: "RED"},
									{Name: "GREEN"},
									{Name: "BLUE"},
								},
							},
							BaseType: &plcopen.DataType{INT: &struct{}{}},
						},
					},
				},
				{
					Name: "MyArray",
					BaseType: &plcopen.DataType{
						Array: &plcopen.DataTypeArray{
							Dimensions: []plcopen.RangeSigned{{Lower: 1, Upper: 10}},
							BaseType:   &plcopen.DataType{REAL: &struct{}{}},
						},
					},
				},
				{
					Name: "MyStruct",
					BaseType: &plcopen.DataType{
						Struct: &plcopen.VarListPlain{
							Variables: []plcopen.VarListPlainVariable{
								{
									Name: "field1",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
								{
									Name: "field2",
									Type: &plcopen.DataType{INT: &struct{}{}},
								},
							},
						},
					},
				},
			},
			POUs: []plcopen.ProjectTypesPOU{
				{
					Name:    "MainProgram",
					POUType: plcopen.POUTypeProgram,
					Interface: &plcopen.ProjectTypesPOUInterface{
						LocalVars: &plcopen.ProjectTypesPOUInterfaceLocalVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "counter",
									Type: &plcopen.DataType{INT: &struct{}{}},
									InitialValue: &plcopen.Value{
										SimpleValue: &plcopen.ValueSimpleValue{Value: "0"},
									},
								},
							},
						},
						InputVars: &plcopen.ProjectTypesPOUInterfaceInputVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "inputValue",
									Type: &plcopen.DataType{REAL: &struct{}{}},
								},
							},
						},
						OutputVars: &plcopen.ProjectTypesPOUInterfaceOutputVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "outputStatus",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
							},
						},
					}, Body: &plcopen.Body{
						ST: &plcopen.BodyST{
							XMLNSXhtml: "http://www.w3.org/1999/xhtml",
							Xhtml:      "<xhtml:p>counter := counter + 1;<xhtml:br/>outputStatus := counter &gt; 10;</xhtml:p>",
						},
					},
				},
				{
					Name:    "TestFunction",
					POUType: plcopen.POUTypeFunction,
					Interface: &plcopen.ProjectTypesPOUInterface{
						ReturnType: &plcopen.DataType{BOOL: &struct{}{}},
						InputVars: &plcopen.ProjectTypesPOUInterfaceInputVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "param1",
									Type: &plcopen.DataType{INT: &struct{}{}},
								},
							},
						},
					}, Body: &plcopen.Body{
						ST: &plcopen.BodyST{
							XMLNSXhtml: "http://www.w3.org/1999/xhtml",
							Xhtml:      "<xhtml:p>TestFunction := param1 &gt; 0;</xhtml:p>",
						},
					},
				},
				{
					Name:    "MyFunction",
					POUType: plcopen.POUTypeFunction,
					Interface: &plcopen.ProjectTypesPOUInterface{
						ReturnType: &plcopen.DataType{REAL: &struct{}{}}, InputVars: &plcopen.ProjectTypesPOUInterfaceInputVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "input1",
									Type: &plcopen.DataType{REAL: &struct{}{}},
								},
								{
									Name: "input2",
									Type: &plcopen.DataType{REAL: &struct{}{}},
								},
							},
						},
					},
					Body: &plcopen.Body{
						FBD: &plcopen.BodyFBD{
							Blocks: []plcopen.BodyFBDBlock{{
								Position: &plcopen.Position{X: 100, Y: 100},
								LocalID:  1,
								TypeName: "ADD",
								InOutVariables: []plcopen.BodyFBDBlockVariable2{
									{
										FormalParameter: "ENO",
									},
								},
								OutputVariables: []plcopen.BodyFBDBlockVariable1{
									{
										FormalParameter: "OUT",
									},
								},
							}},
						},
					},
				},
				{
					Name:    "TestFunctionBlock",
					POUType: plcopen.POUTypeFunctionBlock,
					Interface: &plcopen.ProjectTypesPOUInterface{
						InputVars: &plcopen.ProjectTypesPOUInterfaceInputVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "enable",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
							},
						}, OutputVars: &plcopen.ProjectTypesPOUInterfaceOutputVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "done",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
							},
						},
					}, Body: &plcopen.Body{
						ST: &plcopen.BodyST{
							XMLNSXhtml: "http://www.w3.org/1999/xhtml",
							Xhtml:      "<xhtml:p>done := enable;</xhtml:p>",
						},
					},
				},
				{
					Name:    "MyFunctionBlock",
					POUType: plcopen.POUTypeFunctionBlock,
					Interface: &plcopen.ProjectTypesPOUInterface{
						InputVars: &plcopen.ProjectTypesPOUInterfaceInputVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "start",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
								{
									Name: "reset",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
							},
						},
						OutputVars: &plcopen.ProjectTypesPOUInterfaceOutputVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "busy",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
								{
									Name: "finished",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
							},
						},
					}, Body: &plcopen.Body{
						LD: &plcopen.BodyLD{
							Contacts: []plcopen.BodyLDContact{
								{
									Position: &plcopen.Position{X: 50, Y: 100},
									Variable: "start",
									LocalID:  1,
								},
							},
							Coils: []plcopen.BodyLDCoil{
								{
									Position: &plcopen.Position{X: 200, Y: 100},
									Variable: "busy",
									LocalID:  2,
								},
							},
							LeftPowerRails: []plcopen.BodyLDLeftPowerRail{
								{
									Position: &plcopen.Position{X: 10, Y: 50},
									LocalID:  3,
								},
							},
							RightPowerRails: []plcopen.BodyLDRightPowerRail{
								{
									Position: &plcopen.Position{X: 250, Y: 50},
									LocalID:  4,
								},
							},
						},
					},
				},
			},
		}, Instances: &plcopen.ProjectInstances{
			Configurations: []plcopen.ProjectInstancesConfiguration{
				{
					Name: "DefaultConfig",
					GlobalVars: &plcopen.VarList{
						Variables: []plcopen.VarListVariable{
							{
								Name: "globalCounter",
								Type: &plcopen.DataType{INT: &struct{}{}},
								InitialValue: &plcopen.Value{
									SimpleValue: &plcopen.ValueSimpleValue{Value: "0"},
								},
							},
							{
								Name: "globalFlag",
								Type: &plcopen.DataType{BOOL: &struct{}{}},
								InitialValue: &plcopen.Value{
									SimpleValue: &plcopen.ValueSimpleValue{Value: "FALSE"},
								},
							},
						},
					},
					Resources: []plcopen.ProjectInstancesConfigurationResource{{
						Name: "DefaultResource",
						GlobalVars: &plcopen.VarList{
							Variables: []plcopen.VarListVariable{
								{
									Name: "resourceVar",
									Type: &plcopen.DataType{REAL: &struct{}{}},
									InitialValue: &plcopen.Value{
										SimpleValue: &plcopen.ValueSimpleValue{Value: "0.0"},
									},
								},
							},
						},
						POUInstances: []plcopen.POUInstance{
							{
								Name:     "MainInstance",
								TypeName: "MainProgram",
							},
						}, Tasks: []plcopen.ProjectInstancesConfigurationResourceTask{
							{
								Name:     "MainTask",
								Interval: func() *string { s := "00:00:00.1"; return &s }(),
								Priority: 1,
								POUInstances: []plcopen.POUInstance{
									{
										Name:     "MainInstance",
										TypeName: "MainProgram",
									},
								},
							},
						},
					},
					},
				},
			},
		},
	}
}

// validateWithXmllint attempts to validate XML using xmllint command-line tool
func validateWithXmllint(t *testing.T, xmlFile, xsdFile string) bool {
	// Try xmllint validation (available on most Unix systems and some Windows installations)
	cmd := exec.Command("xmllint", "--schema", xsdFile, "--noout", xmlFile)
	output, err := cmd.CombinedOutput()

	if err != nil {
		// xmllint not available or validation failed
		if exitError, ok := err.(*exec.ExitError); ok {
			if exitError.ExitCode() == 127 {
				// xmllint command not found
				t.Log("xmllint command not found, skipping XSD validation")
				return false
			} else {
				// Validation failed
				t.Errorf("XML validation failed with xmllint: %s\nOutput: %s", err, string(output))
				return true
			}
		}
		t.Logf("xmllint execution error: %v", err)
		return false
	}

	t.Log("XML successfully validated against XSD schema using xmllint")
	if len(output) > 0 {
		t.Logf("xmllint output: %s", string(output))
	}
	return true
}

// checkXMLNamespace verifies the XML contains the proper PLCopen namespace
func checkXMLNamespace(xmlContent string) bool {
	expectedNamespace := `xmlns="http://www.plcopen.org/xml/tc6.xsd"`
	return containsString(xmlContent, expectedNamespace)
}

// containsString checks if a string contains a substring
func containsString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// validateProjectStructure performs structural validation of the parsed project
func validateProjectStructure(t *testing.T, project *plcopen.Project) {
	if project.FileHeader == nil {
		t.Error("Project FileHeader is missing")
		return
	}

	if project.FileHeader.CompanyName == "" {
		t.Error("FileHeader CompanyName is empty")
	}

	if project.FileHeader.ProductName == "" {
		t.Error("FileHeader ProductName is empty")
	}

	if project.ContentHeader == nil {
		t.Error("Project ContentHeader is missing")
		return
	}

	if project.ContentHeader.Name == "" {
		t.Error("ContentHeader Name is empty")
	}

	if project.Types == nil {
		t.Error("Project Types is missing")
		return
	}

	if len(project.Types.POUs) == 0 {
		t.Error("No POUs defined in project")
	}

	// Validate each POU
	for i, pou := range project.Types.POUs {
		if pou.Name == "" {
			t.Errorf("POU[%d] has empty name", i)
		}

		if pou.POUType == "" {
			t.Errorf("POU[%d] '%s' has empty POUType", i, pou.Name)
		}

		if pou.Interface == nil {
			t.Errorf("POU[%d] '%s' has no Interface", i, pou.Name)
		}

		if pou.Body == nil {
			t.Errorf("POU[%d] '%s' has no Body", i, pou.Name)
		}
	}

	t.Log("Project structure validation passed")
}

// TestComplexXMLSchemaValidation tests more complex XML structures
func TestComplexXMLSchemaValidation(t *testing.T) {
	// Create a more complex project with various data types and programming languages
	project := &plcopen.Project{
		FileHeader: &plcopen.ProjectFileHeader{
			CompanyName:        "Complex Test Company",
			ProductName:        "Complex Test Suite",
			ProductVersion:     "2.0",
			ContentDescription: "Complex PLCopen XML validation test",
			CompanyURL:         "http://example.com",
			ProductRelease:     "Alpha",
			CreationDateTime:   time.Now(),
		},
		ContentHeader: &plcopen.ProjectContentHeader{
			Name:                 "ComplexValidationProject",
			Version:              "2.0",
			Organization:         "Complex Test Org",
			Author:               "Complex Test Author",
			Language:             "en-US",
			ModificationDateTime: func() *time.Time { now := time.Now(); return &now }(),
			Comment:              "This is a complex test project",
			CoordinateInfo: &plcopen.ProjectContentHeaderCoordinateInfo{
				PageSize: &plcopen.ProjectContentHeaderCoordinateInfoPageSize{
					X: 210.0,
					Y: 297.0,
				},
				FBD: &plcopen.ProjectContentHeaderCoordinateInfoFBD{
					Scaling: &plcopen.ProjectContentHeaderCoordinateInfoFBDScaling{
						X: 1.0,
						Y: 1.0,
					},
				},
				LD: &plcopen.ProjectContentHeaderCoordinateInfoLD{
					Scaling: &plcopen.ProjectContentHeaderCoordinateInfoLDScaling{
						X: 1.0,
						Y: 1.0,
					},
				},
				SFC: &plcopen.ProjectContentHeaderCoordinateInfoSFC{
					Scaling: &plcopen.ProjectContentHeaderCoordinateInfoSFCScaling{
						X: 1.0,
						Y: 1.0,
					},
				},
			},
		},
		Types: &plcopen.ProjectTypes{
			DataTypes: []plcopen.ProjectTypesDataType{
				// Basic data types
				{
					Name:          "BOOL_TYPE",
					BaseType:      &plcopen.DataType{BOOL: &struct{}{}},
					Documentation: []byte("<documentation>Boolean type</documentation>"),
				},
				{
					Name:     "BYTE_TYPE",
					BaseType: &plcopen.DataType{BYTE: &struct{}{}},
				},
				{
					Name:     "DATE_TYPE",
					BaseType: &plcopen.DataType{DATE: &struct{}{}},
				},
				{
					Name:     "DINT_TYPE",
					BaseType: &plcopen.DataType{DINT: &struct{}{}},
				},
				{
					Name:     "DT_TYPE",
					BaseType: &plcopen.DataType{DT: &struct{}{}},
				},
				{
					Name:     "DWORD_TYPE",
					BaseType: &plcopen.DataType{DWORD: &struct{}{}},
				},
				{
					Name:     "INT_TYPE",
					BaseType: &plcopen.DataType{INT: &struct{}{}},
				},
				{
					Name:     "LINT_TYPE",
					BaseType: &plcopen.DataType{LINT: &struct{}{}},
				},
				{
					Name:     "LREAL_TYPE",
					BaseType: &plcopen.DataType{LREAL: &struct{}{}},
				},
				{
					Name:     "LWORD_TYPE",
					BaseType: &plcopen.DataType{LWORD: &struct{}{}},
				},
				{
					Name:     "REAL_TYPE",
					BaseType: &plcopen.DataType{REAL: &struct{}{}},
				},
				{
					Name:     "SINT_TYPE",
					BaseType: &plcopen.DataType{SINT: &struct{}{}},
				},
				{
					Name:     "TIME_TYPE",
					BaseType: &plcopen.DataType{TIME: &struct{}{}},
				},
				{
					Name:     "TOD_TYPE",
					BaseType: &plcopen.DataType{TOD: &struct{}{}},
				},
				{
					Name:     "UDINT_TYPE",
					BaseType: &plcopen.DataType{UDINT: &struct{}{}},
				},
				{
					Name:     "UINT_TYPE",
					BaseType: &plcopen.DataType{UINT: &struct{}{}},
				},
				{
					Name:     "ULINT_TYPE",
					BaseType: &plcopen.DataType{ULINT: &struct{}{}},
				},
				{
					Name:     "USINT_TYPE",
					BaseType: &plcopen.DataType{USINT: &struct{}{}},
				},
				{
					Name:     "WORD_TYPE",
					BaseType: &plcopen.DataType{WORD: &struct{}{}},
				},
				// String types with length
				{
					Name: "STRING_TYPE",
					BaseType: &plcopen.DataType{
						String: &plcopen.DataTypeString{
							Length: func() *uint64 { n := uint64(80); return &n }(),
						},
					},
				},
				{
					Name: "WSTRING_TYPE",
					BaseType: &plcopen.DataType{
						WString: &plcopen.DataTypeWString{
							Length: func() *uint64 { n := uint64(40); return &n }(),
						},
					},
				},
				// Pointer type
				{
					Name: "POINTER_TYPE",
					BaseType: &plcopen.DataType{
						Pointer: &plcopen.DataTypePointer{
							BaseType: &plcopen.DataType{INT: &struct{}{}},
						},
					},
				},
				// Subrange types
				{
					Name: "SUBRANGE_SIGNED",
					BaseType: &plcopen.DataType{
						SubrangeSigned: &plcopen.DataTypeSubrangeSigned{
							Range:    &plcopen.RangeSigned{Lower: -10, Upper: 10},
							BaseType: &plcopen.DataType{INT: &struct{}{}},
						},
					},
				},
				{
					Name: "SUBRANGE_UNSIGNED",
					BaseType: &plcopen.DataType{
						SubrangeUnsigned: &plcopen.DataTypeSubrangeUnsigned{
							Range:    &plcopen.RangeUnsigned{Lower: 0, Upper: 100},
							BaseType: &plcopen.DataType{UINT: &struct{}{}},
						},
					},
				},
				// Derived type
				{
					Name: "DERIVED_TYPE",
					BaseType: &plcopen.DataType{
						Derived: &plcopen.DataTypeDerived{
							Name: "INT_TYPE",
						},
					},
				},
				{
					Name: "MyArray",
					BaseType: &plcopen.DataType{
						Array: &plcopen.DataTypeArray{
							Dimensions: []plcopen.RangeSigned{{Lower: 1, Upper: 10}},
							BaseType:   &plcopen.DataType{REAL: &struct{}{}},
						},
					},
				},
				{
					Name: "MyStruct",
					BaseType: &plcopen.DataType{
						Struct: &plcopen.VarListPlain{
							Variables: []plcopen.VarListPlainVariable{
								{
									Name: "x",
									Type: &plcopen.DataType{REAL: &struct{}{}},
								},
								{
									Name:          "y",
									Type:          &plcopen.DataType{REAL: &struct{}{}},
									Documentation: []byte("<documentation>Y coordinate</documentation>"),
								},
							},
							Documentation: []byte("<documentation>Structure documentation</documentation>"),
						},
					},
					Documentation: []byte("<documentation>MyStruct type documentation</documentation>"),
				},
			},
			POUs: []plcopen.ProjectTypesPOU{
				{
					Name:    "ComplexProgram",
					POUType: plcopen.POUTypeProgram,
					Interface: &plcopen.ProjectTypesPOUInterface{
						ReturnType: nil, // Program doesn't have a return type
						LocalVars: &plcopen.ProjectTypesPOUInterfaceLocalVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "data",
									Type: &plcopen.DataType{
										Array: &plcopen.DataTypeArray{
											Dimensions: []plcopen.RangeSigned{{Lower: 0, Upper: 9}},
											BaseType:   &plcopen.DataType{INT: &struct{}{}},
										},
									},
									Address: "%MW100",
									InitialValue: &plcopen.Value{
										SimpleValue: &plcopen.ValueSimpleValue{
											Value: "0",
										},
									},
									Documentation: []byte("<documentation>Data array</documentation>"),
								},
							},
						},
						InputVars: &plcopen.ProjectTypesPOUInterfaceInputVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "enable",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
							},
						},
						OutputVars: &plcopen.ProjectTypesPOUInterfaceOutputVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "done",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
							},
						},
						InOutVars: &plcopen.ProjectTypesPOUInterfaceInOutVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "inOutParam",
									Type: &plcopen.DataType{INT: &struct{}{}},
								},
							},
						},
						ExternalVars: &plcopen.ProjectTypesPOUInterfaceExternalVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "externalVar",
									Type: &plcopen.DataType{REAL: &struct{}{}},
								},
							},
						},
						GlobalVars: &plcopen.ProjectTypesPOUInterfaceGlobalVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "globalVar",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
							},
						},
						TempVars: &plcopen.ProjectTypesPOUInterfaceTempVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "tempVar",
									Type: &plcopen.DataType{INT: &struct{}{}},
								},
							},
						},
					},
					Body: &plcopen.Body{
						FBD: &plcopen.BodyFBD{
							InVariables: []plcopen.BodyFBDInVariable{
								{
									LocalID:    1,
									Expression: "enable",
									Position:   &plcopen.Position{X: 50, Y: 100},
								},
							},
							OutVariables: []plcopen.BodyFBDOutVariable{
								{
									LocalID:    2,
									Expression: "done",
									Position:   &plcopen.Position{X: 200, Y: 100},
								},
							},
							Blocks: []plcopen.BodyFBDBlock{
								{
									LocalID:  3,
									TypeName: "AND",
									Position: &plcopen.Position{X: 100, Y: 100},
									InputVariables: []plcopen.BodyFBDBlockVariable{
										{
											FormalParameter: "IN1",
										},
									},
									OutputVariables: []plcopen.BodyFBDBlockVariable1{
										{
											FormalParameter: "OUT",
										},
									},
								},
							},
							ActionBlocks: []plcopen.BodyFBDActionBlock{
								{
									LocalID:  4,
									Position: &plcopen.Position{X: 150, Y: 150},
									Actions: []plcopen.BodyFBDActionBlockAction{
										{
											Qualifier: func() *plcopen.BodyFBDActionBlockActionQualifier {
												q := plcopen.BodyFBDActionBlockActionQualifier("N")
												return &q
											}(),
										},
									},
								},
							},
							Comments: []plcopen.BodyFBDComment{
								{
									LocalID:  5,
									Position: &plcopen.Position{X: 75, Y: 75},
									Content:  "This is a comment",
								},
							},
							Connectors: []plcopen.BodyFBDConnector{
								{
									LocalID:  6,
									Position: &plcopen.Position{X: 125, Y: 125},
									Name:     "CN1",
								},
							},
							Continuations: []plcopen.BodyFBDContinuation{
								{
									LocalID:  7,
									Position: &plcopen.Position{X: 175, Y: 175},
									Name:     "CN1",
								},
							},
							InOutVariables: []plcopen.BodyFBDInOutVariable{
								{
									LocalID:    8,
									Expression: "inOutParam",
									Position:   &plcopen.Position{X: 100, Y: 200},
								},
							},
							Jumps: []plcopen.BodyFBDJump{
								{
									LocalID:  9,
									Position: &plcopen.Position{X: 225, Y: 225},
									Label:    "Label1",
								},
							},
							Labels: []plcopen.BodyFBDLabel{
								{
									LocalID:  10,
									Position: &plcopen.Position{X: 250, Y: 250},
									Label:    "Label1",
								},
							},
							Returns: []plcopen.BodyFBDReturn{
								{
									LocalID:  11,
									Position: &plcopen.Position{X: 275, Y: 275},
								},
							},
						},
					},
					Actions: []plcopen.ProjectTypesPOUAction{
						{
							Name: "ActionName",
							Body: &plcopen.Body{
								ST: &plcopen.BodyST{
									XMLNSXhtml: "http://www.w3.org/1999/xhtml",
									Xhtml:      "<xhtml:p>done := TRUE;</xhtml:p>",
								},
							},
							Documentation: []byte("<documentation>Action documentation</documentation>"),
						},
					},
					Transitions: []plcopen.ProjectTypesPOUTransition{
						{
							Name: "TransitionName",
							Body: &plcopen.Body{
								ST: &plcopen.BodyST{
									XMLNSXhtml: "http://www.w3.org/1999/xhtml",
									Xhtml:      "<xhtml:p>enable</xhtml:p>",
								},
							},
							Documentation: []byte("<documentation>Transition documentation</documentation>"),
						},
					},
					Documentation: []byte("<documentation>Program documentation</documentation>"),
				},
				// Add a POU with SFC body
				{
					Name:    "SFCProgram",
					POUType: plcopen.POUTypeProgram,
					Interface: &plcopen.ProjectTypesPOUInterface{
						LocalVars: &plcopen.ProjectTypesPOUInterfaceLocalVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "counter",
									Type: &plcopen.DataType{INT: &struct{}{}},
								},
							},
						},
					},
					Body: &plcopen.Body{
						SFC: &plcopen.BodySFC{
							Steps: []plcopen.BodySFCStep{
								{
									Name:        "Step1",
									LocalID:     1,
									Position:    &plcopen.Position{X: 100, Y: 100},
									InitialStep: func() *bool { b := true; return &b }(),
								},
								{
									Name:     "Step2",
									LocalID:  2,
									Position: &plcopen.Position{X: 100, Y: 200},
								},
							},
							Transitions: []plcopen.BodySFCTransition{
								{
									LocalID:  3,
									Position: &plcopen.Position{X: 100, Y: 150},
									Condition: &plcopen.BodySFCTransitionCondition{
										Inline: &plcopen.BodySFCTransitionConditionInline{
											Name: "Cond1",
											Body: &plcopen.Body{
												ST: &plcopen.BodyST{
													XMLNSXhtml: "http://www.w3.org/1999/xhtml",
													Xhtml:      "<xhtml:p>counter > 0</xhtml:p>",
												},
											},
										},
									},
								},
							},
						},
					},
				},
				// Add a POU with IL body
				{
					Name:    "ILProgram",
					POUType: plcopen.POUTypeProgram,
					Interface: &plcopen.ProjectTypesPOUInterface{
						LocalVars: &plcopen.ProjectTypesPOUInterfaceLocalVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "result",
									Type: &plcopen.DataType{INT: &struct{}{}},
								},
							},
						},
					},
					Body: &plcopen.Body{
						IL: &plcopen.BodyIL{
							XMLNSXhtml: "http://www.w3.org/1999/xhtml",
							Xhtml:      "<xhtml:p>LD 10<xhtml:br/>ADD 5<xhtml:br/>ST result</xhtml:p>",
						},
					},
				},
			},
		},
		Instances: &plcopen.ProjectInstances{
			Configurations: []plcopen.ProjectInstancesConfiguration{
				{
					Name: "TestConfig",
					GlobalVars: &plcopen.VarList{
						Variables: []plcopen.VarListVariable{
							{
								Name: "configVar",
								Type: &plcopen.DataType{INT: &struct{}{}},
								InitialValue: &plcopen.Value{
									SimpleValue: &plcopen.ValueSimpleValue{Value: "0"},
								},
							},
						},
						Documentation: []byte("<documentation>Global variable documentation</documentation>"),
					},
					Resources: []plcopen.ProjectInstancesConfigurationResource{
						{
							Name: "TestResource",
							GlobalVars: &plcopen.VarList{
								Variables: []plcopen.VarListVariable{
									{
										Name: "resourceVar",
										Type: &plcopen.DataType{BOOL: &struct{}{}},
									},
								},
							},
							Tasks: []plcopen.ProjectInstancesConfigurationResourceTask{
								{
									Name:     "TestTask",
									Priority: 10,
									Interval: func() *string { s := "T#100ms"; return &s }(),
									Single:   func() *string { s := "TRUE"; return &s }(),
									POUInstances: []plcopen.POUInstance{
										{
											Name:     "TestInstance",
											TypeName: "ComplexProgram",
										},
									},
								},
							},
							POUInstances: []plcopen.POUInstance{
								{
									Name:          "MainInstance",
									TypeName:      "SFCProgram",
									Documentation: []byte("<documentation>Instance documentation</documentation>"),
								},
							},
							Documentation: []byte("<documentation>Resource documentation</documentation>"),
						},
					},
					Documentation: []byte("<documentation>Configuration documentation</documentation>"),
				},
			},
		},
	}

	xmlData, err := xml.MarshalIndent(project, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal complex project: %v", err)
	}

	// Validate the XML is well-formed
	var parsed plcopen.Project
	err = xml.Unmarshal(xmlData, &parsed)
	if err != nil {
		t.Errorf("Complex XML is not well-formed: %v", err)
	} else {
		t.Log("Complex XML is well-formed")
	}

	// Write to file for manual inspection if needed
	complexXMLFile := filepath.Join(os.TempDir(), "complex_plcopen.xml")
	err = os.WriteFile(complexXMLFile, xmlData, 0644)
	if err != nil {
		t.Logf("Could not write complex XML file: %v", err)
	} else {
		t.Logf("Complex XML written to: %s", complexXMLFile)
	}

	fmt.Printf("Complex XML generated:\n%s\n", string(xmlData))
}

// TestXMLLintValidation tests XML validation using xmllint tool
func TestXMLLintValidation(t *testing.T) {
	// Check if xmllint is available
	_, err := exec.LookPath("xmllint")
	if err != nil {
		t.Skip("xmllint not available, skipping XSD validation test")
	}
	// Create a comprehensive project for validation
	project := createComprehensiveProject()

	// Marshal to XML
	xmlData, err := xml.MarshalIndent(project, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal project to XML: %v", err)
	}

	// Add XML declaration
	fullXML := []byte(xml.Header + string(xmlData))

	// Write XML to temporary file
	tmpFile, err := ioutil.TempFile("", "plcopen_test_*.xml")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(fullXML); err != nil {
		t.Fatalf("Failed to write XML to temporary file: %v", err)
	}
	tmpFile.Close()

	// Get XSD file path
	xsdPath := "TC6_XML_V10_B.xsd"
	if _, err := os.Stat(xsdPath); os.IsNotExist(err) {
		xsdPath = "TC6_XML_V10.xsd"
		if _, err := os.Stat(xsdPath); os.IsNotExist(err) {
			t.Skip("XSD schema file not found, skipping validation")
		}
	}

	// Validate XML against XSD using xmllint
	cmd := exec.Command("xmllint", "--schema", xsdPath, "--noout", tmpFile.Name())
	output, err := cmd.CombinedOutput()

	t.Logf("xmllint output: %s", string(output))

	if err != nil {
		t.Logf("Generated XML (first 2000 chars):\n%s", string(fullXML[:min(2000, len(fullXML))]))
		t.Fatalf("XML validation failed: %v\nOutput: %s", err, string(output))
	}

	// Additional check: validate that XML is well-formed
	cmd = exec.Command("xmllint", "--noout", tmpFile.Name())
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("XML well-formedness check failed: %v\nOutput: %s", err, string(output))
	}

	t.Logf("XML validation successful - document is valid against PLCopen XSD schema")
}

// TestGeneratedXMLStructure validates the structure of generated XML
func TestGeneratedXMLStructure(t *testing.T) {
	project := createComprehensiveProject()

	xmlData, err := xml.MarshalIndent(project, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal project to XML: %v", err)
	}

	xmlStr := string(xmlData)

	// Check for essential PLCopen XML elements
	requiredElements := []string{
		"<project",
		"xmlns=\"http://www.plcopen.org/xml/tc6.xsd\"",
		"<fileHeader",
		"<contentHeader",
		"<types>",
		"<dataTypes>",
		"<pous>",
		"<instances>",
		"<configurations>",
		"<globalVars>",
	}

	for _, element := range requiredElements {
		if !strings.Contains(xmlStr, element) {
			t.Errorf("Required XML element not found: %s", element)
		}
	}

	// Check for specific data types
	dataTypeElements := []string{
		"<dataType name=\"BOOL\">",
		"<dataType name=\"INT\">",
		"<dataType name=\"REAL\">",
		"<dataType name=\"STRING\">",
		"<dataType name=\"MyArray\">",
		"<dataType name=\"MyEnum\">",
		"<dataType name=\"MyStruct\">",
	}

	for _, element := range dataTypeElements {
		if !strings.Contains(xmlStr, element) {
			t.Errorf("Expected data type element not found: %s", element)
		}
	}

	// Check for POU types
	pouElements := []string{
		"<pou name=\"MainProgram\" pouType=\"program\">",
		"<pou name=\"MyFunction\" pouType=\"function\">",
		"<pou name=\"MyFunctionBlock\" pouType=\"functionBlock\">",
	}

	for _, element := range pouElements {
		if !strings.Contains(xmlStr, element) {
			t.Errorf("Expected POU element not found: %s", element)
		}
	} // Check for different body types
	bodyElements := []string{
		"<ST", // ST with attributes
		"<SFC>",
		"<FBD>",
		"<LD>",
		"<IL", // IL with attributes
	}

	bodyCount := 0
	foundBodyTypes := []string{}
	for _, element := range bodyElements {
		if strings.Contains(xmlStr, element) {
			bodyCount++
			foundBodyTypes = append(foundBodyTypes, element)
		}
	}

	t.Logf("Found body types: %v (count: %d)", foundBodyTypes, bodyCount)

	if bodyCount < 3 { // We should have at least 3 different body types
		t.Errorf("Expected multiple body types, found only %d. Found types: %v", bodyCount, foundBodyTypes)
	}

	t.Logf("XML structure validation passed - all required elements present")
}

// Helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
