// Package plcopen contains tests for PLCopen XML data types
package tests

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/suifei/plcopen-go"
)

// TestAllDataTypes tests all data types defined in the PLCopen schema
func TestAllDataTypes(t *testing.T) {
	// Create a project that includes all possible data types and structures
	project := &plcopen.Project{
		FileHeader: &plcopen.ProjectFileHeader{
			CompanyName:        "DataTypes Test Company",
			ProductName:        "DataTypes Test Product",
			ProductVersion:     "1.0",
			CompanyURL:         "http://datatypes.example.com",
			ProductRelease:     "Beta",
			ContentDescription: "Comprehensive data types test",
			CreationDateTime:   time.Now(),
		},
		ContentHeader: &plcopen.ProjectContentHeader{
			Name:                 "DataTypesTest",
			Version:              "1.0",
			Organization:         "Test Org",
			Author:               "Test Author",
			Language:             "en",
			ModificationDateTime: func() *time.Time { now := time.Now(); return &now }(),
			Comment:              "This project tests all data types",
			CoordinateInfo: &plcopen.ProjectContentHeaderCoordinateInfo{
				PageSize: &plcopen.ProjectContentHeaderCoordinateInfoPageSize{X: 210, Y: 297},
				FBD:      &plcopen.ProjectContentHeaderCoordinateInfoFBD{Scaling: &plcopen.ProjectContentHeaderCoordinateInfoFBDScaling{X: 1.0, Y: 1.0}},
				LD:       &plcopen.ProjectContentHeaderCoordinateInfoLD{Scaling: &plcopen.ProjectContentHeaderCoordinateInfoLDScaling{X: 1.0, Y: 1.0}},
				SFC:      &plcopen.ProjectContentHeaderCoordinateInfoSFC{Scaling: &plcopen.ProjectContentHeaderCoordinateInfoSFCScaling{X: 1.0, Y: 1.0}},
			},
		},
		Types: &plcopen.ProjectTypes{
			DataTypes: []plcopen.ProjectTypesDataType{
				// Test all basic data types
				{Name: "T_BOOL", BaseType: &plcopen.DataType{BOOL: &struct{}{}}},
				{Name: "T_BYTE", BaseType: &plcopen.DataType{BYTE: &struct{}{}}},
				{Name: "T_DATE", BaseType: &plcopen.DataType{DATE: &struct{}{}}},
				{Name: "T_DINT", BaseType: &plcopen.DataType{DINT: &struct{}{}}},
				{Name: "T_DT", BaseType: &plcopen.DataType{DT: &struct{}{}}},
				{Name: "T_DWORD", BaseType: &plcopen.DataType{DWORD: &struct{}{}}},
				{Name: "T_INT", BaseType: &plcopen.DataType{INT: &struct{}{}}},
				{Name: "T_LINT", BaseType: &plcopen.DataType{LINT: &struct{}{}}},
				{Name: "T_LREAL", BaseType: &plcopen.DataType{LREAL: &struct{}{}}},
				{Name: "T_LWORD", BaseType: &plcopen.DataType{LWORD: &struct{}{}}},
				{Name: "T_REAL", BaseType: &plcopen.DataType{REAL: &struct{}{}}},
				{Name: "T_SINT", BaseType: &plcopen.DataType{SINT: &struct{}{}}},
				{Name: "T_TIME", BaseType: &plcopen.DataType{TIME: &struct{}{}}},
				{Name: "T_TOD", BaseType: &plcopen.DataType{TOD: &struct{}{}}},
				{Name: "T_UDINT", BaseType: &plcopen.DataType{UDINT: &struct{}{}}},
				{Name: "T_UINT", BaseType: &plcopen.DataType{UINT: &struct{}{}}},
				{Name: "T_ULINT", BaseType: &plcopen.DataType{ULINT: &struct{}{}}},
				{Name: "T_USINT", BaseType: &plcopen.DataType{USINT: &struct{}{}}},
				{Name: "T_WORD", BaseType: &plcopen.DataType{WORD: &struct{}{}}},

				// String types
				{
					Name:     "T_STRING80",
					BaseType: &plcopen.DataType{String: &plcopen.DataTypeString{Length: func() *uint64 { n := uint64(80); return &n }()}},
				},
				{
					Name:     "T_WSTRING40",
					BaseType: &plcopen.DataType{WString: &plcopen.DataTypeWString{Length: func() *uint64 { n := uint64(40); return &n }()}},
				},

				// Array types
				{
					Name: "T_ARRAY_INT",
					BaseType: &plcopen.DataType{
						Array: &plcopen.DataTypeArray{
							Dimensions: []plcopen.RangeSigned{{Lower: 1, Upper: 10}},
							BaseType:   &plcopen.DataType{INT: &struct{}{}},
						},
					},
				},
				{
					Name: "T_ARRAY_2D",
					BaseType: &plcopen.DataType{
						Array: &plcopen.DataTypeArray{
							Dimensions: []plcopen.RangeSigned{
								{Lower: 1, Upper: 5},
								{Lower: 1, Upper: 5},
							},
							BaseType: &plcopen.DataType{REAL: &struct{}{}},
						},
					},
				},

				// Struct type
				{
					Name: "T_POINT",
					BaseType: &plcopen.DataType{
						Struct: &plcopen.VarListPlain{
							Variables: []plcopen.VarListPlainVariable{
								{Name: "x", Type: &plcopen.DataType{REAL: &struct{}{}}},
								{Name: "y", Type: &plcopen.DataType{REAL: &struct{}{}}},
								{Name: "z", Type: &plcopen.DataType{REAL: &struct{}{}}},
							},
							Documentation: []byte("<documentation>3D point</documentation>"),
						},
					},
				},

				// Pointer type
				{
					Name: "T_POINTER_TO_INT",
					BaseType: &plcopen.DataType{
						Pointer: &plcopen.DataTypePointer{
							BaseType: &plcopen.DataType{INT: &struct{}{}},
						},
					},
				},

				// Enumerated type
				{
					Name: "T_COLOR",
					BaseType: &plcopen.DataType{
						Enum: &plcopen.DataTypeEnum{
							Values: &plcopen.DataTypeEnumValues{
								Values: []plcopen.DataTypeEnumValuesValue{
									{Name: "RED", Documentation: []byte("<documentation>Red color</documentation>")},
									{Name: "GREEN"},
									{Name: "BLUE"},
									{Name: "YELLOW"},
								},
							},
							BaseType: &plcopen.DataType{INT: &struct{}{}},
						},
					},
				},

				// Subrange types
				{
					Name: "T_SUBRANGE_SIGNED",
					BaseType: &plcopen.DataType{
						SubrangeSigned: &plcopen.DataTypeSubrangeSigned{
							Range:    &plcopen.RangeSigned{Lower: -100, Upper: 100},
							BaseType: &plcopen.DataType{INT: &struct{}{}},
						},
					},
				},
				{
					Name: "T_SUBRANGE_UNSIGNED",
					BaseType: &plcopen.DataType{
						SubrangeUnsigned: &plcopen.DataTypeSubrangeUnsigned{
							Range:    &plcopen.RangeUnsigned{Lower: 0, Upper: 200},
							BaseType: &plcopen.DataType{UINT: &struct{}{}},
						},
					},
				},

				// Derived type
				{
					Name: "T_DERIVED_TYPE",
					BaseType: &plcopen.DataType{
						Derived: &plcopen.DataTypeDerived{
							Name: "T_INT",
						},
					},
				},
			},
			POUs: []plcopen.ProjectTypesPOU{
				{
					Name:    "DataTypesProgram",
					POUType: plcopen.POUTypeProgram,
					Interface: &plcopen.ProjectTypesPOUInterface{
						LocalVars: &plcopen.ProjectTypesPOUInterfaceLocalVars{
							Variables: []plcopen.VarListVariable{
								{
									Name:         "boolVar",
									Type:         &plcopen.DataType{BOOL: &struct{}{}},
									InitialValue: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "TRUE"}},
								},
								{
									Name:         "intVar",
									Type:         &plcopen.DataType{INT: &struct{}{}},
									InitialValue: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "123"}},
								},
								{
									Name:         "realVar",
									Type:         &plcopen.DataType{REAL: &struct{}{}},
									InitialValue: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "123.456"}},
								},
								{
									Name: "arrayVar",
									Type: &plcopen.DataType{Derived: &plcopen.DataTypeDerived{Name: "T_ARRAY_INT"}},
									InitialValue: &plcopen.Value{
										ArrayValue: &plcopen.ValueArrayValue{
											Values: []plcopen.ValueArrayValueValue{
												{RepeatCount: func() *uint64 { n := uint64(5); return &n }(), Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "0"}}},
												{Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "1"}}},
												{Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "2"}}},
												{Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "3"}}},
												{Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "4"}}},
											},
										},
									},
								},
								{
									Name: "structVar",
									Type: &plcopen.DataType{Derived: &plcopen.DataTypeDerived{Name: "T_POINT"}},
									InitialValue: &plcopen.Value{
										StructValue: &plcopen.ValueStructValue{
											Values: []plcopen.ValueStructValueValue{
												{Member: "x", Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "1.0"}}},
												{Member: "y", Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "2.0"}}},
												{Member: "z", Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "3.0"}}},
											},
										},
									},
								},
							},
						},
					},
					Body: &plcopen.Body{
						ST: &plcopen.BodyST{
							XMLNSXhtml: "http://www.w3.org/1999/xhtml",
							Xhtml:      "<xhtml:p>// Data types test program</xhtml:p>",
						},
					},
				},
			},
		},
		Instances: &plcopen.ProjectInstances{
			Configurations: []plcopen.ProjectInstancesConfiguration{
				{
					Name: "TypesConfig",
					Resources: []plcopen.ProjectInstancesConfigurationResource{
						{
							Name: "TypesResource",
							POUInstances: []plcopen.POUInstance{
								{
									Name:          "DataTypesInstance",
									TypeName:      "DataTypesProgram",
									Documentation: []byte("<documentation>Types test instance</documentation>"),
								},
							},
						},
					},
				},
			},
		},
	}

	// Marshal to XML and verify
	xmlData, err := xml.MarshalIndent(project, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal data types project: %v", err)
	}

	// Unmarshal to verify well-formedness
	var parsed plcopen.Project
	err = xml.Unmarshal(xmlData, &parsed)
	if err != nil {
		t.Errorf("Data types XML is not well-formed: %v", err)
	} else {
		t.Log("Data types XML is well-formed")
	}

	// Verify all data types
	if len(parsed.Types.DataTypes) != len(project.Types.DataTypes) {
		t.Errorf("Parsed data types count mismatch. Expected: %d, Got: %d",
			len(project.Types.DataTypes), len(parsed.Types.DataTypes))
	}

	// Verify that the data types were correctly parsed
	dataTypeCount := map[string]int{
		"BOOL": 0, "BYTE": 0, "DATE": 0,
		"DINT": 0, "DT": 0, "DWORD": 0,
		"INT": 0, "LINT": 0, "LREAL": 0,
		"LWORD": 0, "REAL": 0, "SINT": 0,
		"TIME": 0, "TOD": 0, "UDINT": 0,
		"UINT": 0, "ULINT": 0, "USINT": 0,
		"WORD": 0, "String": 0, "WString": 0,
		"Array": 0, "Struct": 0, "Pointer": 0,
		"Enum": 0, "SubrangeSigned": 0, "SubrangeUnsigned": 0,
		"Derived": 0,
	}

	// Count the occurrences of each data type in the parsed project
	for _, dataType := range parsed.Types.DataTypes {
		baseType := dataType.BaseType
		if baseType.BOOL != nil {
			dataTypeCount["BOOL"]++
		}
		if baseType.BYTE != nil {
			dataTypeCount["BYTE"]++
		}
		if baseType.DATE != nil {
			dataTypeCount["DATE"]++
		}
		if baseType.DINT != nil {
			dataTypeCount["DINT"]++
		}
		if baseType.DT != nil {
			dataTypeCount["DT"]++
		}
		if baseType.DWORD != nil {
			dataTypeCount["DWORD"]++
		}
		if baseType.INT != nil {
			dataTypeCount["INT"]++
		}
		if baseType.LINT != nil {
			dataTypeCount["LINT"]++
		}
		if baseType.LREAL != nil {
			dataTypeCount["LREAL"]++
		}
		if baseType.LWORD != nil {
			dataTypeCount["LWORD"]++
		}
		if baseType.REAL != nil {
			dataTypeCount["REAL"]++
		}
		if baseType.SINT != nil {
			dataTypeCount["SINT"]++
		}
		if baseType.TIME != nil {
			dataTypeCount["TIME"]++
		}
		if baseType.TOD != nil {
			dataTypeCount["TOD"]++
		}
		if baseType.UDINT != nil {
			dataTypeCount["UDINT"]++
		}
		if baseType.UINT != nil {
			dataTypeCount["UINT"]++
		}
		if baseType.ULINT != nil {
			dataTypeCount["ULINT"]++
		}
		if baseType.USINT != nil {
			dataTypeCount["USINT"]++
		}
		if baseType.WORD != nil {
			dataTypeCount["WORD"]++
		}
		if baseType.String != nil {
			dataTypeCount["String"]++
		}
		if baseType.WString != nil {
			dataTypeCount["WString"]++
		}
		if baseType.Array != nil {
			dataTypeCount["Array"]++
		}
		if baseType.Struct != nil {
			dataTypeCount["Struct"]++
		}
		if baseType.Pointer != nil {
			dataTypeCount["Pointer"]++
		}
		if baseType.Enum != nil {
			dataTypeCount["Enum"]++
		}
		if baseType.SubrangeSigned != nil {
			dataTypeCount["SubrangeSigned"]++
		}
		if baseType.SubrangeUnsigned != nil {
			dataTypeCount["SubrangeUnsigned"]++
		}
		if baseType.Derived != nil {
			dataTypeCount["Derived"]++
		}
	}

	// Print the data type counts for debugging
	for typeName, count := range dataTypeCount {
		t.Logf("Data type %s count: %d", typeName, count)
		if count == 0 {
			t.Errorf("Data type %s not found in parsed project", typeName)
		}
	}

	// Check value types (SimpleValue, ArrayValue, StructValue)
	valueTypes := map[string]bool{
		"SimpleValue": false,
		"ArrayValue":  false,
		"StructValue": false,
	}

	for _, pou := range parsed.Types.POUs {
		for _, variable := range pou.Interface.LocalVars.Variables {
			if variable.InitialValue != nil {
				if variable.InitialValue.SimpleValue != nil {
					valueTypes["SimpleValue"] = true
				}
				if variable.InitialValue.ArrayValue != nil {
					valueTypes["ArrayValue"] = true
				}
				if variable.InitialValue.StructValue != nil {
					valueTypes["StructValue"] = true
				}
			}
		}
	}

	// Verify all value types were used
	for valueType, found := range valueTypes {
		if !found {
			t.Errorf("Value type %s not found in parsed project", valueType)
		}
	}

	// Write to file for manual inspection if needed
	dataTypesXMLFile := filepath.Join(os.TempDir(), "datatypes_plcopen.xml")
	if err := os.WriteFile(dataTypesXMLFile, xmlData, 0644); err != nil {
		t.Logf("Could not write data types XML file: %v", err)
	} else {
		t.Logf("Data types XML written to: %s", dataTypesXMLFile)
	}
}
