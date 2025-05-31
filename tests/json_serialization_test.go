package tests

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/suifei/plcopen-go"
)

// TestJSONSerialization tests JSON marshaling and unmarshaling
func TestJSONSerialization(t *testing.T) {
	// Create a test project
	project := &plcopen.Project{
		FileHeader: &plcopen.ProjectFileHeader{
			CompanyName:        "JSON Test Company",
			ProductName:        "JSON Test Product",
			ProductVersion:     "1.0.0",
			ContentDescription: "PLCopen Go JSON serialization test",
			CreationDateTime:   time.Date(2025, 5, 31, 12, 0, 0, 0, time.UTC),
		},
		ContentHeader: &plcopen.ProjectContentHeader{
			Name:         "JSONTestProject",
			Version:      "1.0",
			Organization: "PLCopen-Go",
			Author:       "Test Author",
			Language:     "en",
		},
		Types: &plcopen.ProjectTypes{
			DataTypes: []plcopen.ProjectTypesDataType{
				{
					Name: "TestBool",
					BaseType: &plcopen.DataType{
						BOOL: &struct{}{},
					},
				},
			},
			POUs: []plcopen.ProjectTypesPOU{
				{
					Name:    "TestFunction",
					POUType: plcopen.POUTypeFunction,
					Interface: &plcopen.ProjectTypesPOUInterface{
						ReturnType: &plcopen.DataType{
							BOOL: &struct{}{},
						},
					},
					Body: &plcopen.Body{
						ST: &plcopen.BodyST{
							XMLNSXhtml: "http://www.w3.org/1999/xhtml",
							Xhtml:      "<xhtml:p>TestFunction := TRUE;</xhtml:p>",
						},
					},
				},
			},
		},
	}

	// Test JSON marshaling
	jsonData, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		t.Fatalf("JSON marshaling failed: %v", err)
	}

	t.Logf("JSON output length: %d bytes", len(jsonData))

	// Verify JSON contains expected fields
	jsonStr := string(jsonData)
	expectedFields := []string{
		`"companyName": "JSON Test Company"`,
		`"productName": "JSON Test Product"`,
		`"name": "JSONTestProject"`,
		`"organization": "PLCopen-Go"`,
		`"pOUType": "function"`,
	}
	for _, field := range expectedFields {
		if !jsonContainsString(jsonStr, field) {
			t.Errorf("JSON output missing expected field: %s", field)
		}
	}

	// Test JSON unmarshaling
	var unmarshaledProject plcopen.Project
	err = json.Unmarshal(jsonData, &unmarshaledProject)
	if err != nil {
		t.Fatalf("JSON unmarshaling failed: %v", err)
	}

	// Verify unmarshaled data
	if unmarshaledProject.FileHeader.CompanyName != project.FileHeader.CompanyName {
		t.Errorf("CompanyName mismatch: got %s, want %s",
			unmarshaledProject.FileHeader.CompanyName, project.FileHeader.CompanyName)
	}

	if unmarshaledProject.ContentHeader.Name != project.ContentHeader.Name {
		t.Errorf("Project name mismatch: got %s, want %s",
			unmarshaledProject.ContentHeader.Name, project.ContentHeader.Name)
	}

	if len(unmarshaledProject.Types.POUs) != len(project.Types.POUs) {
		t.Errorf("POU count mismatch: got %d, want %d",
			len(unmarshaledProject.Types.POUs), len(project.Types.POUs))
	} else if len(unmarshaledProject.Types.POUs) > 0 {
		if unmarshaledProject.Types.POUs[0].POUType != project.Types.POUs[0].POUType {
			t.Errorf("POU type mismatch: got %s, want %s",
				unmarshaledProject.Types.POUs[0].POUType, project.Types.POUs[0].POUType)
		}
	}

	t.Log("JSON serialization test passed successfully!")
}

// TestJSONXMLCompatibility tests that the same data can be serialized to both JSON and XML
func TestJSONXMLCompatibility(t *testing.T) {
	// Create a comprehensive test project
	project := &plcopen.Project{
		FileHeader: &plcopen.ProjectFileHeader{
			CompanyName:        "Compatibility Test Co",
			ProductName:        "Multi-Format Test",
			ProductVersion:     "2.0.0",
			CreationDateTime:   time.Date(2025, 5, 31, 15, 30, 0, 0, time.UTC),
			ContentDescription: "Testing JSON/XML compatibility",
		},
		ContentHeader: &plcopen.ProjectContentHeader{
			Name:         "CompatibilityTest",
			Version:      "2.0",
			Organization: "Test Org",
			Author:       "Compatibility Tester",
			Language:     "en-US",
		},
		Types: &plcopen.ProjectTypes{
			DataTypes: []plcopen.ProjectTypesDataType{
				{
					Name: "MyInt",
					BaseType: &plcopen.DataType{
						INT: &struct{}{},
					},
				},
				{
					Name: "MyEnum",
					BaseType: &plcopen.DataType{
						Enum: &plcopen.DataTypeEnum{
							Values: &plcopen.DataTypeEnumValues{
								Values: []plcopen.DataTypeEnumValuesValue{
									{Name: "VALUE1"},
									{Name: "VALUE2"},
									{Name: "VALUE3"},
								},
							},
							BaseType: &plcopen.DataType{
								INT: &struct{}{},
							},
						},
					},
				},
			},
			POUs: []plcopen.ProjectTypesPOU{
				{
					Name:    "TestProgram",
					POUType: plcopen.POUTypeProgram,
					Interface: &plcopen.ProjectTypesPOUInterface{
						LocalVars: &plcopen.VarList{
							Variables: []plcopen.VarListVariable{
								{
									Name: "counter",
									Type: &plcopen.DataType{
										INT: &struct{}{},
									},
								},
								{
									Name: "flag",
									Type: &plcopen.DataType{
										BOOL: &struct{}{},
									},
								},
							},
						},
					},
					Body: &plcopen.Body{
						ST: &plcopen.BodyST{
							XMLNSXhtml: "http://www.w3.org/1999/xhtml",
							Xhtml:      "<xhtml:p>counter := counter + 1;<xhtml:br/>flag := counter > 10;</xhtml:p>",
						},
					},
				},
			},
		},
		Instances: &plcopen.ProjectInstances{
			Configurations: []plcopen.ProjectInstancesConfiguration{
				{
					Name: "MainConfig",
					Resources: []plcopen.ProjectInstancesConfigurationResource{
						{
							Name: "CPU",
							POUInstances: []plcopen.POUInstance{
								{
									Name:     "MainProgram",
									TypeName: "TestProgram",
								},
							},
						},
					},
				},
			},
		},
	}

	// Test XML serialization
	xmlData, err := xml.MarshalIndent(project, "", "  ")
	if err != nil {
		t.Fatalf("XML marshaling failed: %v", err)
	}

	// Test JSON serialization
	jsonData, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		t.Fatalf("JSON marshaling failed: %v", err)
	}

	t.Logf("XML output length: %d bytes", len(xmlData))
	t.Logf("JSON output length: %d bytes", len(jsonData))

	// Test XML unmarshaling
	var xmlProject plcopen.Project
	err = xml.Unmarshal(xmlData, &xmlProject)
	if err != nil {
		t.Fatalf("XML unmarshaling failed: %v", err)
	}

	// Test JSON unmarshaling
	var jsonProject plcopen.Project
	err = json.Unmarshal(jsonData, &jsonProject)
	if err != nil {
		t.Fatalf("JSON unmarshaling failed: %v", err)
	}

	// Compare key fields from both unmarshaled projects
	if xmlProject.FileHeader.CompanyName != jsonProject.FileHeader.CompanyName {
		t.Errorf("CompanyName mismatch between XML and JSON: XML=%s, JSON=%s",
			xmlProject.FileHeader.CompanyName, jsonProject.FileHeader.CompanyName)
	}

	if xmlProject.ContentHeader.Name != jsonProject.ContentHeader.Name {
		t.Errorf("Project name mismatch between XML and JSON: XML=%s, JSON=%s",
			xmlProject.ContentHeader.Name, jsonProject.ContentHeader.Name)
	}

	if len(xmlProject.Types.DataTypes) != len(jsonProject.Types.DataTypes) {
		t.Errorf("DataType count mismatch: XML=%d, JSON=%d",
			len(xmlProject.Types.DataTypes), len(jsonProject.Types.DataTypes))
	}

	if len(xmlProject.Types.POUs) != len(jsonProject.Types.POUs) {
		t.Errorf("POU count mismatch: XML=%d, JSON=%d",
			len(xmlProject.Types.POUs), len(jsonProject.Types.POUs))
	}

	if len(xmlProject.Instances.Configurations) != len(jsonProject.Instances.Configurations) {
		t.Errorf("Configuration count mismatch: XML=%d, JSON=%d",
			len(xmlProject.Instances.Configurations), len(jsonProject.Instances.Configurations))
	}

	t.Log("JSON/XML compatibility test passed successfully!")
}

// TestJSONOmitEmpty tests that omitempty works correctly for optional fields
func TestJSONOmitEmpty(t *testing.T) {
	// Create project with minimal required fields only
	project := &plcopen.Project{
		FileHeader: &plcopen.ProjectFileHeader{
			CompanyName:      "Minimal Company",
			ProductName:      "Minimal Product",
			ProductVersion:   "1.0",
			CreationDateTime: time.Date(2025, 5, 31, 10, 0, 0, 0, time.UTC),
			// Note: ContentDescription is omitted (should not appear in JSON)
		},
		ContentHeader: &plcopen.ProjectContentHeader{
			Name: "MinimalProject",
			// Note: Version, Organization, etc. are omitted
		},
		// Note: Types and Instances are omitted
	}

	jsonData, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		t.Fatalf("JSON marshaling failed: %v", err)
	}

	jsonStr := string(jsonData)
	t.Logf("Minimal JSON output:\n%s", jsonStr)

	// Verify omitted optional fields are not present
	omittedFields := []string{
		`"contentDescription"`,
		`"version"`,
		`"organization"`,
		`"types"`,
		`"instances"`,
	}
	for _, field := range omittedFields {
		if jsonContainsString(jsonStr, field) {
			t.Errorf("JSON output should not contain omitted field: %s", field)
		}
	}

	// Verify required fields are present
	requiredFields := []string{
		`"companyName": "Minimal Company"`,
		`"productName": "Minimal Product"`,
		`"name": "MinimalProject"`,
	}

	for _, field := range requiredFields {
		if !jsonContainsString(jsonStr, field) {
			t.Errorf("JSON output missing required field: %s", field)
		}
	}

	t.Log("JSON omitempty test passed successfully!")
}

// Helper function to check if a string contains a substring
func jsonContainsString(s, substr string) bool {
	return len(s) >= len(substr) && jsonFindString(s, substr) >= 0
}

// Simple string search function
func jsonFindString(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
