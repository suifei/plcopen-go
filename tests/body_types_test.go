// Package plcopen contains tests for PLCopen XML body types
package tests

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/suifei/plcopen-go"
)

// TestAllBodyTypes tests all body types defined in the PLCopen schema
// This includes ST, FBD, LD, SFC and IL body types
func TestAllBodyTypes(t *testing.T) {
	// Create a project with all body types
	project := &plcopen.Project{
		FileHeader: &plcopen.ProjectFileHeader{
			CompanyName:        "Body Types Test Company",
			ProductName:        "Body Types Test Product",
			ProductVersion:     "1.0",
			ContentDescription: "Testing all PLC body types",
			CreationDateTime:   time.Now(),
		},
		ContentHeader: &plcopen.ProjectContentHeader{
			Name:         "BodyTypesTest",
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
			POUs: []plcopen.ProjectTypesPOU{
				// ST body type
				{
					Name:    "ST_Program",
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
						ST: &plcopen.BodyST{
							XMLNSXhtml: "http://www.w3.org/1999/xhtml",
							Xhtml:      "<xhtml:p>counter := counter + 1;</xhtml:p>",
						},
					},
				},

				// FBD body type
				{
					Name:    "FBD_Program",
					POUType: plcopen.POUTypeProgram,
					Interface: &plcopen.ProjectTypesPOUInterface{
						InputVars: &plcopen.ProjectTypesPOUInterfaceInputVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "in1",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
								{
									Name: "in2",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
							},
						},
						OutputVars: &plcopen.ProjectTypesPOUInterfaceOutputVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "out1",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
							},
						},
					},
					Body: &plcopen.Body{
						FBD: &plcopen.BodyFBD{
							// Input variables
							InVariables: []plcopen.BodyFBDInVariable{
								{
									LocalID:    1,
									Position:   &plcopen.Position{X: 50, Y: 50},
									Expression: "in1",
								},
								{
									LocalID:    2,
									Position:   &plcopen.Position{X: 50, Y: 100},
									Expression: "in2",
								},
							},
							// Output variables
							OutVariables: []plcopen.BodyFBDOutVariable{
								{
									LocalID:    3,
									Position:   &plcopen.Position{X: 300, Y: 75},
									Expression: "out1",
								},
							},
							// Blocks
							Blocks: []plcopen.BodyFBDBlock{
								{
									LocalID:  4,
									Position: &plcopen.Position{X: 150, Y: 75},
									TypeName: "AND",
									InputVariables: []plcopen.BodyFBDBlockVariable{
										{
											FormalParameter: "IN1",
											ConnectionPointIn: &plcopen.ConnectionPointIn{
												Connections: []plcopen.Connection{
													{
														RefLocalID: 1,
													},
												},
											},
										},
										{
											FormalParameter: "IN2",
											ConnectionPointIn: &plcopen.ConnectionPointIn{
												Connections: []plcopen.Connection{
													{
														RefLocalID: 2,
													},
												},
											},
										},
									},
									OutputVariables: []plcopen.BodyFBDBlockVariable1{
										{
											FormalParameter: "OUT",
										},
									},
								},
							},
							// FBD Comments
							Comments: []plcopen.BodyFBDComment{
								{
									LocalID:  5,
									Position: &plcopen.Position{X: 150, Y: 20},
									Content:  "This is an AND gate",
								},
							},
							// Connectors
							Connectors: []plcopen.BodyFBDConnector{
								{
									LocalID:  6,
									Position: &plcopen.Position{X: 200, Y: 200},
									Name:     "CN1",
								},
							},
							// Continuations
							Continuations: []plcopen.BodyFBDContinuation{
								{
									LocalID:  7,
									Position: &plcopen.Position{X: 250, Y: 200},
									Name:     "CN1",
								},
							},
							// Labels and jumps
							Labels: []plcopen.BodyFBDLabel{
								{
									LocalID:  8,
									Position: &plcopen.Position{X: 300, Y: 200},
									Label:    "LBL1",
								},
							},
							Jumps: []plcopen.BodyFBDJump{
								{
									LocalID:  9,
									Position: &plcopen.Position{X: 350, Y: 200},
									Label:    "LBL1",
								},
							},
							// Return
							Returns: []plcopen.BodyFBDReturn{
								{
									LocalID:  10,
									Position: &plcopen.Position{X: 400, Y: 200},
								},
							},
							// ActionBlocks
							ActionBlocks: []plcopen.BodyFBDActionBlock{
								{
									LocalID:  11,
									Position: &plcopen.Position{X: 450, Y: 200},
									Actions: []plcopen.BodyFBDActionBlockAction{
										{
											Qualifier: func() *plcopen.BodyFBDActionBlockActionQualifier {
												q := plcopen.BodyFBDActionBlockActionQualifier("N")
												return &q
											}(),
											Reference: &plcopen.BodyFBDActionBlockActionReference{
												Name: "ActionName",
											},
										},
									},
								},
							},
						},
					},
					// POU Actions
					Actions: []plcopen.ProjectTypesPOUAction{
						{
							Name: "ActionName",
							Body: &plcopen.Body{
								ST: &plcopen.BodyST{
									XMLNSXhtml: "http://www.w3.org/1999/xhtml",
									Xhtml:      "<xhtml:p>// Action code</xhtml:p>",
								},
							},
						},
					},
				},

				// LD body type
				{
					Name:    "LD_Program",
					POUType: plcopen.POUTypeProgram,
					Interface: &plcopen.ProjectTypesPOUInterface{
						LocalVars: &plcopen.ProjectTypesPOUInterfaceLocalVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "input1",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
								{
									Name: "input2",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
								{
									Name: "output1",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
							},
						},
					},
					Body: &plcopen.Body{
						LD: &plcopen.BodyLD{
							// Left power rail
							LeftPowerRails: []plcopen.BodyLDLeftPowerRail{
								{
									LocalID:  1,
									Position: &plcopen.Position{X: 10, Y: 100},
								},
							},
							// Contacts
							Contacts: []plcopen.BodyLDContact{
								{
									LocalID:  2,
									Position: &plcopen.Position{X: 100, Y: 100},
									Variable: "input1",
									Negated:  func() *bool { b := false; return &b }(),
								},
								{
									LocalID:      3,
									Position:     &plcopen.Position{X: 100, Y: 200},
									Variable:     "input2",
									Negated:      func() *bool { b := true; return &b }(),
									EdgeModifier: func() *plcopen.EdgeModifierType { e := plcopen.EdgeModifierTypeRising; return &e }(),
								},
							},
							// Coils
							Coils: []plcopen.BodyLDCoil{
								{
									LocalID:         4,
									Position:        &plcopen.Position{X: 300, Y: 100},
									Variable:        "output1",
									StorageModifier: func() *plcopen.StorageModifierType { s := plcopen.StorageModifierTypeSet; return &s }(),
								},
							},
							// Right power rail
							RightPowerRails: []plcopen.BodyLDRightPowerRail{
								{
									LocalID:  5,
									Position: &plcopen.Position{X: 400, Y: 100},
								},
							},
						},
					},
				},

				// SFC body type
				{
					Name:    "SFC_Program",
					POUType: plcopen.POUTypeProgram,
					Interface: &plcopen.ProjectTypesPOUInterface{
						LocalVars: &plcopen.ProjectTypesPOUInterfaceLocalVars{
							Variables: []plcopen.VarListVariable{
								{
									Name: "counter",
									Type: &plcopen.DataType{INT: &struct{}{}},
								},
								{
									Name: "done",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
							},
						},
					},
					Body: &plcopen.Body{
						SFC: &plcopen.BodySFC{
							// Steps
							Steps: []plcopen.BodySFCStep{
								{
									Name:        "Init",
									LocalID:     1,
									Position:    &plcopen.Position{X: 100, Y: 100},
									InitialStep: func() *bool { b := true; return &b }(),
								},
								{
									Name:     "Step1",
									LocalID:  2,
									Position: &plcopen.Position{X: 100, Y: 200},
								},
								{
									Name:     "Step2",
									LocalID:  3,
									Position: &plcopen.Position{X: 100, Y: 300},
								},
							},
							// Transitions
							Transitions: []plcopen.BodySFCTransition{
								{
									LocalID:  4,
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
								{
									LocalID:  5,
									Position: &plcopen.Position{X: 100, Y: 250},
									Condition: &plcopen.BodySFCTransitionCondition{
										Reference: &plcopen.BodySFCTransitionConditionReference{
											Name: "TransCond",
										},
									},
									Priority: func() *uint64 { n := uint64(10); return &n }(),
								},
							},
						},
					},
					// Transitions defined for SFC
					Transitions: []plcopen.ProjectTypesPOUTransition{
						{
							Name: "TransCond",
							Body: &plcopen.Body{
								ST: &plcopen.BodyST{
									XMLNSXhtml: "http://www.w3.org/1999/xhtml",
									Xhtml:      "<xhtml:p>done</xhtml:p>",
								},
							},
						},
					},
				},

				// IL body type
				{
					Name:    "IL_Program",
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
					Name: "BodyTypesConfig",
					Resources: []plcopen.ProjectInstancesConfigurationResource{
						{
							Name: "BodyTypesResource",
							POUInstances: []plcopen.POUInstance{
								{
									Name:     "ST_Instance",
									TypeName: "ST_Program",
								},
								{
									Name:     "FBD_Instance",
									TypeName: "FBD_Program",
								},
								{
									Name:     "LD_Instance",
									TypeName: "LD_Program",
								},
								{
									Name:     "SFC_Instance",
									TypeName: "SFC_Program",
								},
								{
									Name:     "IL_Instance",
									TypeName: "IL_Program",
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
		t.Fatalf("Failed to marshal body types project: %v", err)
	}

	// Unmarshal to verify well-formedness
	var parsed plcopen.Project
	err = xml.Unmarshal(xmlData, &parsed)
	if err != nil {
		t.Errorf("Body types XML is not well-formed: %v", err)
	} else {
		t.Log("Body types XML is well-formed")
	}

	// Check that we have all body types in the parsed project
	bodyTypes := map[string]bool{
		"ST":  false,
		"FBD": false,
		"LD":  false,
		"SFC": false,
		"IL":  false,
	}

	// Find body types in the parsed POUs
	for _, pou := range parsed.Types.POUs {
		body := pou.Body
		if body != nil {
			if body.ST != nil {
				bodyTypes["ST"] = true
			}
			if body.FBD != nil {
				bodyTypes["FBD"] = true
			}
			if body.LD != nil {
				bodyTypes["LD"] = true
			}
			if body.SFC != nil {
				bodyTypes["SFC"] = true
			}
			if body.IL != nil {
				bodyTypes["IL"] = true
			}
		}
	}

	// Verify all body types were found
	for bodyType, found := range bodyTypes {
		t.Logf("Body type %s found: %v", bodyType, found)
		if !found {
			t.Errorf("Body type %s not found in parsed project", bodyType)
		}
	}

	// Write to file for manual inspection if needed
	bodyTypesXMLFile := filepath.Join(os.TempDir(), "bodytypes_plcopen.xml")
	if err := os.WriteFile(bodyTypesXMLFile, xmlData, 0644); err != nil {
		t.Logf("Could not write body types XML file: %v", err)
	} else {
		t.Logf("Body types XML written to: %s", bodyTypesXMLFile)
	}
}
