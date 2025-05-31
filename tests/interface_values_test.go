// Package models contains tests for PLCopen XML interface and value types
package tests

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/suifei/plcopen-go"
)

// TestInterfaceAndValues tests all variable interfaces and value types in the PLCopen schema
func TestInterfaceAndValues(t *testing.T) {
	// Create a project with comprehensive variable interfaces and values
	project := &plcopen.Project{
		FileHeader: &plcopen.ProjectFileHeader{
			CompanyName:        "Interface Test Company",
			ProductName:        "Interface Test Product",
			ProductVersion:     "1.0",
			ContentDescription: "Testing all variable interfaces and values",
			CreationDateTime:   time.Now(),
		},
		ContentHeader: &plcopen.ProjectContentHeader{
			Name:         "InterfaceTest",
			Version:      "1.0",
			Organization: "Test Org",
			Author:       "Test Author",
			Language:     "en",
			CoordinateInfo: &plcopen.ProjectContentHeaderCoordinateInfo{
				PageSize: &plcopen.ProjectContentHeaderCoordinateInfoPageSize{X: 210, Y: 297},
				FBD:      &plcopen.ProjectContentHeaderCoordinateInfoFBD{Scaling: &plcopen.ProjectContentHeaderCoordinateInfoFBDScaling{X: 1.0, Y: 1.0}},
				LD:       &plcopen.ProjectContentHeaderCoordinateInfoLD{Scaling: &plcopen.ProjectContentHeaderCoordinateInfoLDScaling{X: 1.0, Y: 1.0}},
				SFC:      &plcopen.ProjectContentHeaderCoordinateInfoSFC{Scaling: &plcopen.ProjectContentHeaderCoordinateInfoSFCScaling{X: 1.0, Y: 1.0}},
			},
		},
		Types: &plcopen.ProjectTypes{
			// Define data types for testing
			DataTypes: []plcopen.ProjectTypesDataType{
				{
					Name: "TestStruct",
					BaseType: &plcopen.DataType{
						Struct: &plcopen.VarListPlain{
							Variables: []plcopen.VarListPlainVariable{
								{Name: "field1", Type: &plcopen.DataType{INT: &struct{}{}}},
								{Name: "field2", Type: &plcopen.DataType{REAL: &struct{}{}}},
								{Name: "field3", Type: &plcopen.DataType{BOOL: &struct{}{}}},
							},
						},
					},
				},
				{
					Name: "TestArray",
					BaseType: &plcopen.DataType{
						Array: &plcopen.DataTypeArray{
							Dimensions: []plcopen.RangeSigned{{Lower: 1, Upper: 10}},
							BaseType:   &plcopen.DataType{INT: &struct{}{}},
						},
					},
				},
			},
			// Define POUs with different interface types
			POUs: []plcopen.ProjectTypesPOU{
				// Function with all interface types
				{
					Name:    "TestFunction",
					POUType: plcopen.POUTypeFunction,
					Interface: &plcopen.ProjectTypesPOUInterface{
						// Return type for function
						ReturnType: &plcopen.DataType{INT: &struct{}{}},

						// Input variables
						InputVars: &plcopen.ProjectTypesPOUInterfaceInputVars{
							Variables: []plcopen.VarListVariable{
								{
									Name:         "inputVar1",
									Type:         &plcopen.DataType{BOOL: &struct{}{}},
									InitialValue: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "TRUE"}},
								},
								{
									Name:         "inputVar2",
									Type:         &plcopen.DataType{INT: &struct{}{}},
									InitialValue: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "10"}},
								},
							},
							Documentation: []byte("<documentation>Input variables</documentation>"),
						},

						// Output variables
						OutputVars: &plcopen.ProjectTypesPOUInterfaceOutputVars{
							Variables: []plcopen.VarListVariable{
								{
									Name:         "outputVar1",
									Type:         &plcopen.DataType{BOOL: &struct{}{}},
									InitialValue: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "FALSE"}},
								},
								{
									Name: "outputVar2",
									Type: &plcopen.DataType{INT: &struct{}{}},
								},
							},
						},

						// InOut variables
						InOutVars: &plcopen.ProjectTypesPOUInterfaceInOutVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "inoutVar",
									Type: &plcopen.DataType{REAL: &struct{}{}},
								},
							},
						},

						// Local variables
						LocalVars: &plcopen.ProjectTypesPOUInterfaceLocalVars{
							Variables: []plcopen.VarListVariable{
								{
									Name:         "localVar1",
									Type:         &plcopen.DataType{INT: &struct{}{}},
									InitialValue: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "0"}},
								},
								{
									Name: "localArray",
									Type: &plcopen.DataType{Derived: &plcopen.DataTypeDerived{Name: "TestArray"}},
									InitialValue: &plcopen.Value{
										ArrayValue: &plcopen.ValueArrayValue{
											Values: []plcopen.ValueArrayValueValue{
												{RepeatCount: func() *uint64 { n := uint64(10); return &n }(), Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "0"}}},
											},
										},
									},
								},
								{
									Name: "localStruct",
									Type: &plcopen.DataType{Derived: &plcopen.DataTypeDerived{Name: "TestStruct"}},
									InitialValue: &plcopen.Value{
										StructValue: &plcopen.ValueStructValue{
											Values: []plcopen.ValueStructValueValue{
												{Member: "field1", Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "1"}}},
												{Member: "field2", Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "2.5"}}},
												{Member: "field3", Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "TRUE"}}},
											},
										},
									},
								},
							},
						},

						// Temp variables
						TempVars: &plcopen.ProjectTypesPOUInterfaceTempVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "tempVar",
									Type: &plcopen.DataType{INT: &struct{}{}},
								},
							},
						},

						// External variables
						ExternalVars: &plcopen.ProjectTypesPOUInterfaceExternalVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "externalVar",
									Type: &plcopen.DataType{INT: &struct{}{}},
								},
							},
						},

						// Global variables
						GlobalVars: &plcopen.ProjectTypesPOUInterfaceGlobalVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "globalVar",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
							},
						},

						// Documentation
						Documentation: []byte("<documentation>Function interface documentation</documentation>"),
					},
					Body: &plcopen.Body{
						ST: &plcopen.BodyST{
							XMLNSXhtml: "http://www.w3.org/1999/xhtml",
							Xhtml:      "<xhtml:p>// Function body</xhtml:p>",
						},
					},
				},

				// Function block with variable address attributes
				{
					Name:    "TestFunctionBlock",
					POUType: plcopen.POUTypeFunctionBlock,
					Interface: &plcopen.ProjectTypesPOUInterface{
						InputVars: &plcopen.ProjectTypesPOUInterfaceInputVars{
							Variables: []plcopen.VarListVariable{
								{
									Name:    "fbInput1",
									Type:    &plcopen.DataType{BOOL: &struct{}{}},
									Address: "%IX0.0",
								},
							},
						},
						OutputVars: &plcopen.ProjectTypesPOUInterfaceOutputVars{
							Variables: []plcopen.VarListVariable{
								{
									Name:    "fbOutput1",
									Type:    &plcopen.DataType{BOOL: &struct{}{}},
									Address: "%QX0.0",
								},
							},
						},
						LocalVars: &plcopen.ProjectTypesPOUInterfaceLocalVars{
							Variables: []plcopen.VarListVariable{
								{
									Name:    "fbLocal1",
									Type:    &plcopen.DataType{WORD: &struct{}{}},
									Address: "%MW10",
								},
							},
						},
					},
					Body: &plcopen.Body{
						ST: &plcopen.BodyST{
							XMLNSXhtml: "http://www.w3.org/1999/xhtml",
							Xhtml:      "<xhtml:p>// Function block body</xhtml:p>",
						},
					},
				},

				// Program with array and struct initialization
				{
					Name:    "TestProgram",
					POUType: plcopen.POUTypeProgram,
					Interface: &plcopen.ProjectTypesPOUInterface{
						LocalVars: &plcopen.ProjectTypesPOUInterfaceLocalVars{
							Variables: []plcopen.VarListVariable{
								// Nested array initialization
								{
									Name: "nestedArray",
									Type: &plcopen.DataType{
										Array: &plcopen.DataTypeArray{
											Dimensions: []plcopen.RangeSigned{
												{Lower: 1, Upper: 2},
												{Lower: 1, Upper: 3},
											},
											BaseType: &plcopen.DataType{INT: &struct{}{}},
										},
									},
									InitialValue: &plcopen.Value{
										ArrayValue: &plcopen.ValueArrayValue{
											Values: []plcopen.ValueArrayValueValue{
												{
													Value: &plcopen.Value{
														ArrayValue: &plcopen.ValueArrayValue{
															Values: []plcopen.ValueArrayValueValue{
																{Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "1"}}},
																{Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "2"}}},
																{Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "3"}}},
															},
														},
													},
												},
												{
													Value: &plcopen.Value{
														ArrayValue: &plcopen.ValueArrayValue{
															Values: []plcopen.ValueArrayValueValue{
																{Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "4"}}},
																{Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "5"}}},
																{Value: &plcopen.Value{SimpleValue: &plcopen.ValueSimpleValue{Value: "6"}}},
															},
														},
													},
												},
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
							Xhtml:      "<xhtml:p>// Program body</xhtml:p>",
						},
					},
				},
			},
		},
		Instances: &plcopen.ProjectInstances{
			Configurations: []plcopen.ProjectInstancesConfiguration{
				{
					Name: "VarConfig",
					GlobalVars: &plcopen.VarList{
						Variables: []plcopen.VarListVariable{
							{
								Name:    "globalCounter",
								Type:    &plcopen.DataType{INT: &struct{}{}},
								Address: "%MD100",
								InitialValue: &plcopen.Value{
									SimpleValue: &plcopen.ValueSimpleValue{Value: "0"},
								},
							},
						},
					},
					Resources: []plcopen.ProjectInstancesConfigurationResource{
						{
							Name: "VarResource",
							GlobalVars: &plcopen.VarList{
								Variables: []plcopen.VarListVariable{
									{
										Name:    "resourceVar",
										Type:    &plcopen.DataType{REAL: &struct{}{}},
										Address: "%MD200",
										InitialValue: &plcopen.Value{
											SimpleValue: &plcopen.ValueSimpleValue{Value: "0.0"},
										},
									},
								},
							},
							Tasks: []plcopen.ProjectInstancesConfigurationResourceTask{
								{
									Name:     "MainTask",
									Priority: 1,
									Interval: func() *string { s := "T#100ms"; return &s }(),
									Single:   func() *string { s := "FALSE"; return &s }(),
									POUInstances: []plcopen.POUInstance{
										{
											Name:     "MainProgram",
											TypeName: "TestProgram",
										},
									},
								},
							},
							POUInstances: []plcopen.POUInstance{
								{
									Name:     "FunctionBlockInstance",
									TypeName: "TestFunctionBlock",
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
		t.Fatalf("Failed to marshal interface and values project: %v", err)
	}

	// Unmarshal to verify well-formedness
	var parsed plcopen.Project
	err = xml.Unmarshal(xmlData, &parsed)
	if err != nil {
		t.Errorf("Interface and values XML is not well-formed: %v", err)
	} else {
		t.Log("Interface and values XML is well-formed")
	}

	// Check that we have all variable interface types
	interfaceTypes := map[string]bool{
		"InputVars":     false,
		"OutputVars":    false,
		"InOutVars":     false,
		"LocalVars":     false,
		"TempVars":      false,
		"ExternalVars":  false,
		"GlobalVars":    false,
		"ReturnType":    false,
		"Documentation": false,
	}

	// Find interface types in the parsed POUs
	for _, pou := range parsed.Types.POUs {
		if pou.Interface != nil {
			if pou.Interface.InputVars != nil {
				interfaceTypes["InputVars"] = true
			}
			if pou.Interface.OutputVars != nil {
				interfaceTypes["OutputVars"] = true
			}
			if pou.Interface.InOutVars != nil {
				interfaceTypes["InOutVars"] = true
			}
			if pou.Interface.LocalVars != nil {
				interfaceTypes["LocalVars"] = true
			}
			if pou.Interface.TempVars != nil {
				interfaceTypes["TempVars"] = true
			}
			if pou.Interface.ExternalVars != nil {
				interfaceTypes["ExternalVars"] = true
			}
			if pou.Interface.GlobalVars != nil {
				interfaceTypes["GlobalVars"] = true
			}
			if pou.Interface.ReturnType != nil {
				interfaceTypes["ReturnType"] = true
			}
			if pou.Interface.Documentation != nil {
				interfaceTypes["Documentation"] = true
			}
		}
	}

	// Verify all interface types were found
	for interfaceType, found := range interfaceTypes {
		t.Logf("Interface type %s found: %v", interfaceType, found)
		if !found {
			t.Errorf("Interface type %s not found in parsed project", interfaceType)
		}
	}

	// Check value types
	valueTypes := map[string]bool{
		"SimpleValue": false,
		"ArrayValue":  false,
		"StructValue": false,
		"RepeatCount": false,
	}

	// Find value types in the parsed POUs
	for _, pou := range parsed.Types.POUs {
		if pou.Interface != nil {
			// Check all variable lists for value types
			varLists := []struct {
				name string
				list *plcopen.VarList
			}{
				{"InputVars", pou.Interface.InputVars},
				{"OutputVars", pou.Interface.OutputVars},
				{"InOutVars", pou.Interface.InOutVars},
				{"LocalVars", pou.Interface.LocalVars},
				{"TempVars", pou.Interface.TempVars},
				{"ExternalVars", pou.Interface.ExternalVars},
				{"GlobalVars", pou.Interface.GlobalVars},
			}

			for _, vl := range varLists {
				if vl.list != nil {
					for _, v := range vl.list.Variables {
						if v.InitialValue != nil {
							if v.InitialValue.SimpleValue != nil {
								valueTypes["SimpleValue"] = true
							}
							if v.InitialValue.ArrayValue != nil {
								valueTypes["ArrayValue"] = true

								// Check for repeat count
								for _, arrValue := range v.InitialValue.ArrayValue.Values {
									if arrValue.RepeatCount != nil {
										valueTypes["RepeatCount"] = true
									}
								}
							}
							if v.InitialValue.StructValue != nil {
								valueTypes["StructValue"] = true
							}
						}
					}
				}
			}
		}
	}

	// Verify all value types were found
	for valueType, found := range valueTypes {
		t.Logf("Value type %s found: %v", valueType, found)
		if !found {
			t.Errorf("Value type %s not found in parsed project", valueType)
		}
	}

	// Check POU types
	pouTypes := map[plcopen.POUType]bool{
		plcopen.POUTypeFunction:      false,
		plcopen.POUTypeFunctionBlock: false,
		plcopen.POUTypeProgram:       false,
	}

	for _, pou := range parsed.Types.POUs {
		pouTypes[pou.POUType] = true
	}

	// Verify all POU types were found
	for pouType, found := range pouTypes {
		t.Logf("POU type %s found: %v", pouType, found)
		if !found {
			t.Errorf("POU type %s not found in parsed project", pouType)
		}
	}

	// Write to file for manual inspection if needed
	interfaceXMLFile := filepath.Join(os.TempDir(), "interface_plcopen.xml")
	if err := os.WriteFile(interfaceXMLFile, xmlData, 0644); err != nil {
		t.Logf("Could not write interface XML file: %v", err)
	} else {
		t.Logf("Interface XML written to: %s", interfaceXMLFile)
	}
}
