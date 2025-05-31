package tests

import (
	"encoding/xml"
	"strings"
	"testing"
	"time"

	"github.com/suifei/plcopen-go"
)

// TestPOUTypeConstants tests all POUType constants
func TestPOUTypeConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant plcopen.POUType
		expected string
	}{
		{"Function", plcopen.POUTypeFunction, "function"},
		{"FunctionBlock", plcopen.POUTypeFunctionBlock, "functionBlock"},
		{"Program", plcopen.POUTypeProgram, "program"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.constant) != tt.expected {
				t.Errorf("POUType constant %s = %v, want %v", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

// TestEdgeModifierTypeConstants tests all EdgeModifierType constants
func TestEdgeModifierTypeConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant plcopen.EdgeModifierType
		expected string
	}{
		{"None", plcopen.EdgeModifierTypeNone, "none"},
		{"Falling", plcopen.EdgeModifierTypeFalling, "falling"},
		{"Rising", plcopen.EdgeModifierTypeRising, "rising"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.constant) != tt.expected {
				t.Errorf("EdgeModifierType constant %s = %v, want %v", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

// TestStorageModifierTypeConstants tests all StorageModifierType constants
func TestStorageModifierTypeConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant plcopen.StorageModifierType
		expected string
	}{
		{"None", plcopen.StorageModifierTypeNone, "none"},
		{"Set", plcopen.StorageModifierTypeSet, "set"},
		{"Reset", plcopen.StorageModifierTypeReset, "reset"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.constant) != tt.expected {
				t.Errorf("StorageModifierType constant %s = %v, want %v", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

// TestBodyFBDActionBlockActionQualifierConstants tests all action qualifier constants
func TestBodyFBDActionBlockActionQualifierConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant plcopen.BodyFBDActionBlockActionQualifier
		expected string
	}{
		{"P1", plcopen.BodyFBDActionBlockActionQualifierP1, "P1"},
		{"N", plcopen.BodyFBDActionBlockActionQualifierN, "N"},
		{"P0", plcopen.BodyFBDActionBlockActionQualifierP0, "P0"},
		{"R", plcopen.BodyFBDActionBlockActionQualifierR, "R"},
		{"S", plcopen.BodyFBDActionBlockActionQualifierS, "S"},
		{"L", plcopen.BodyFBDActionBlockActionQualifierL, "L"},
		{"D", plcopen.BodyFBDActionBlockActionQualifierD, "D"},
		{"P", plcopen.BodyFBDActionBlockActionQualifierP, "P"},
		{"DS", plcopen.BodyFBDActionBlockActionQualifierDS, "DS"},
		{"DL", plcopen.BodyFBDActionBlockActionQualifierDL, "DL"},
		{"SD", plcopen.BodyFBDActionBlockActionQualifierSD, "SD"},
		{"SL", plcopen.BodyFBDActionBlockActionQualifierSL, "SL"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.constant) != tt.expected {
				t.Errorf("BodyFBDActionBlockActionQualifier constant %s = %v, want %v", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

// TestProjectXMLMarshaling tests XML marshaling and unmarshaling of Project
func TestProjectXMLMarshaling(t *testing.T) {
	creationTime := time.Now()
	modificationTime := time.Now().Add(time.Hour)

	original := &plcopen.Project{
		FileHeader: &plcopen.ProjectFileHeader{
			CompanyName:        "Test Company",
			CompanyURL:         "https://test.com",
			ProductName:        "Test Product",
			ProductVersion:     "1.0",
			ProductRelease:     "Release 1",
			CreationDateTime:   creationTime,
			ContentDescription: "Test Description",
		},
		ContentHeader: &plcopen.ProjectContentHeader{
			Name:                 "TestProject",
			Version:              "1.0",
			ModificationDateTime: &modificationTime,
			Organization:         "Test Org",
			Author:               "Test Author",
			Language:             "en",
			Comment:              "Test Comment",
			CoordinateInfo: &plcopen.ProjectContentHeaderCoordinateInfo{
				PageSize: &plcopen.ProjectContentHeaderCoordinateInfoPageSize{
					X: 100.0,
					Y: 200.0,
				},
				FBD: &plcopen.ProjectContentHeaderCoordinateInfoFBD{
					Scaling: &plcopen.ProjectContentHeaderCoordinateInfoFBDScaling{
						X: 1.0,
						Y: 1.0,
					},
				},
				LD: &plcopen.ProjectContentHeaderCoordinateInfoLD{
					Scaling: &plcopen.ProjectContentHeaderCoordinateInfoLDScaling{
						X: 2.0,
						Y: 2.0,
					},
				},
				SFC: &plcopen.ProjectContentHeaderCoordinateInfoSFC{
					Scaling: &plcopen.ProjectContentHeaderCoordinateInfoSFCScaling{
						X: 3.0,
						Y: 3.0,
					},
				},
			},
		},
		Types: &plcopen.ProjectTypes{
			DataTypes: []plcopen.ProjectTypesDataType{
				{
					Name: "TestDataType",
					BaseType: &plcopen.DataType{
						BOOL: &struct{}{},
					},
					InitialValue: &plcopen.Value{
						SimpleValue: &plcopen.ValueSimpleValue{
							Value: "TRUE",
						},
					},
					Documentation: []byte("Test documentation"),
				},
			},
			POUs: []plcopen.ProjectTypesPOU{
				{
					Name:    "TestPOU",
					POUType: plcopen.POUTypeFunction,
					Interface: &plcopen.ProjectTypesPOUInterface{
						ReturnType: &plcopen.DataType{
							INT: &struct{}{},
						},
						InputVars: &plcopen.VarList{
							Variables: []plcopen.VarListVariable{
								{
									Name: "input1",
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
							Xhtml:      "RETURN input1;",
						},
					},
				},
			},
		},
		Instances: &plcopen.ProjectInstances{
			Configurations: []plcopen.ProjectInstancesConfiguration{
				{
					Name: "Config1",
					Resources: []plcopen.ProjectInstancesConfigurationResource{
						{
							Name: "Resource1",
							Tasks: []plcopen.ProjectInstancesConfigurationResourceTask{
								{
									Name:     "Task1",
									Priority: 1,
								},
							},
						},
					},
				},
			},
		},
	}

	// Test marshaling
	data, err := xml.MarshalIndent(original, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal Project: %v", err)
	}

	// Test unmarshaling
	var unmarshaled plcopen.Project
	err = xml.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal Project: %v", err)
	}

	// Verify basic fields
	if unmarshaled.FileHeader.CompanyName != original.FileHeader.CompanyName {
		t.Errorf("CompanyName = %v, want %v", unmarshaled.FileHeader.CompanyName, original.FileHeader.CompanyName)
	}
	if unmarshaled.ContentHeader.Name != original.ContentHeader.Name {
		t.Errorf("Name = %v, want %v", unmarshaled.ContentHeader.Name, original.ContentHeader.Name)
	}
}

// TestDataTypeAllFields tests all DataType fields
func TestDataTypeAllFields(t *testing.T) {
	tests := []struct {
		name     string
		dataType *plcopen.DataType
	}{
		{"BOOL", &plcopen.DataType{BOOL: &struct{}{}}},
		{"BYTE", &plcopen.DataType{BYTE: &struct{}{}}},
		{"DATE", &plcopen.DataType{DATE: &struct{}{}}},
		{"DINT", &plcopen.DataType{DINT: &struct{}{}}},
		{"DT", &plcopen.DataType{DT: &struct{}{}}},
		{"DWORD", &plcopen.DataType{DWORD: &struct{}{}}},
		{"INT", &plcopen.DataType{INT: &struct{}{}}},
		{"LINT", &plcopen.DataType{LINT: &struct{}{}}},
		{"LREAL", &plcopen.DataType{LREAL: &struct{}{}}},
		{"LWORD", &plcopen.DataType{LWORD: &struct{}{}}},
		{"REAL", &plcopen.DataType{REAL: &struct{}{}}},
		{"SINT", &plcopen.DataType{SINT: &struct{}{}}},
		{"TIME", &plcopen.DataType{TIME: &struct{}{}}},
		{"TOD", &plcopen.DataType{TOD: &struct{}{}}},
		{"UDINT", &plcopen.DataType{UDINT: &struct{}{}}},
		{"UINT", &plcopen.DataType{UINT: &struct{}{}}},
		{"ULINT", &plcopen.DataType{ULINT: &struct{}{}}},
		{"USINT", &plcopen.DataType{USINT: &struct{}{}}},
		{"WORD", &plcopen.DataType{WORD: &struct{}{}}},
		{
			"Array",
			&plcopen.DataType{
				Array: &plcopen.DataTypeArray{
					Dimensions: []plcopen.RangeSigned{
						{Lower: 0, Upper: 10},
					},
					BaseType: &plcopen.DataType{INT: &struct{}{}},
				},
			},
		},
		{
			"Derived",
			&plcopen.DataType{
				Derived: &plcopen.DataTypeDerived{
					Name: "MyType",
				},
			},
		},
		{
			"Enum",
			&plcopen.DataType{
				Enum: &plcopen.DataTypeEnum{
					Values: &plcopen.DataTypeEnumValues{
						Values: []plcopen.DataTypeEnumValuesValue{
							{Name: "VALUE1"},
							{Name: "VALUE2"},
						},
					},
					BaseType: &plcopen.DataType{INT: &struct{}{}},
				},
			},
		},
		{
			"Pointer",
			&plcopen.DataType{
				Pointer: &plcopen.DataTypePointer{
					BaseType: &plcopen.DataType{INT: &struct{}{}},
				},
			},
		},
		{
			"String",
			&plcopen.DataType{
				String: &plcopen.DataTypeString{
					Length: uint64Ptr(255),
				},
			},
		},
		{
			"WString",
			&plcopen.DataType{
				WString: &plcopen.DataTypeWString{
					Length: uint64Ptr(255),
				},
			},
		},
		{
			"Struct",
			&plcopen.DataType{
				Struct: &plcopen.VarListPlain{
					Variables: []plcopen.VarListPlainVariable{
						{
							Name: "field1",
							Type: &plcopen.DataType{INT: &struct{}{}},
						},
					},
				},
			},
		},
		{
			"SubrangeSigned",
			&plcopen.DataType{
				SubrangeSigned: &plcopen.DataTypeSubrangeSigned{
					Range:    &plcopen.RangeSigned{Lower: -100, Upper: 100},
					BaseType: &plcopen.DataType{INT: &struct{}{}},
				},
			},
		},
		{
			"SubrangeUnsigned",
			&plcopen.DataType{
				SubrangeUnsigned: &plcopen.DataTypeSubrangeUnsigned{
					Range:    &plcopen.RangeUnsigned{Lower: 0, Upper: 100},
					BaseType: &plcopen.DataType{UINT: &struct{}{}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test marshaling
			data, err := xml.Marshal(tt.dataType)
			if err != nil {
				t.Fatalf("Failed to marshal DataType %s: %v", tt.name, err)
			}

			// Test unmarshaling
			var unmarshaled plcopen.DataType
			err = xml.Unmarshal(data, &unmarshaled)
			if err != nil {
				t.Fatalf("Failed to unmarshal DataType %s: %v", tt.name, err)
			}
		})
	}
}

// TestValueTypes tests all Value types
func TestValueTypes(t *testing.T) {
	tests := []struct {
		name  string
		value *plcopen.Value
	}{
		{
			"SimpleValue",
			&plcopen.Value{
				SimpleValue: &plcopen.ValueSimpleValue{
					Value: "42",
				},
			},
		},
		{
			"ArrayValue",
			&plcopen.Value{
				ArrayValue: &plcopen.ValueArrayValue{
					Values: []plcopen.ValueArrayValueValue{
						{
							RepeatCount: uint64Ptr(3),
							Value: &plcopen.Value{
								SimpleValue: &plcopen.ValueSimpleValue{
									Value: "1",
								},
							},
						},
					},
				},
			},
		},
		{
			"StructValue",
			&plcopen.Value{
				StructValue: &plcopen.ValueStructValue{
					Values: []plcopen.ValueStructValueValue{
						{
							Member: "field1",
							Value: &plcopen.Value{
								SimpleValue: &plcopen.ValueSimpleValue{
									Value: "42",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test marshaling
			data, err := xml.Marshal(tt.value)
			if err != nil {
				t.Fatalf("Failed to marshal Value %s: %v", tt.name, err)
			}

			// Test unmarshaling
			var unmarshaled plcopen.Value
			err = xml.Unmarshal(data, &unmarshaled)
			if err != nil {
				t.Fatalf("Failed to unmarshal Value %s: %v", tt.name, err)
			}
		})
	}
}

// TestBodyTypes tests all Body types
func TestBodyTypes(t *testing.T) {
	tests := []struct {
		name string
		body *plcopen.Body
	}{
		{
			"FBD",
			&plcopen.Body{
				FBD: &plcopen.BodyFBD{
					Blocks: []plcopen.BodyFBDBlock{
						{
							LocalID:  1,
							TypeName: "AND",
							Position: &plcopen.Position{X: 100, Y: 200},
							InputVariables: []plcopen.BodyFBDBlockVariable{
								{
									FormalParameter: "IN1",
									ConnectionPointIn: &plcopen.ConnectionPointIn{
										Connections: []plcopen.Connection{
											{RefLocalID: 2},
										},
									},
								},
							},
							OutputVariables: []plcopen.BodyFBDBlockVariable1{
								{
									FormalParameter:    "OUT",
									ConnectionPointOut: &plcopen.ConnectionPointOut{},
								},
							},
						},
					},
					InVariables: []plcopen.BodyFBDInVariable{
						{
							LocalID:            2,
							Expression:         "input1",
							Position:           &plcopen.Position{X: 50, Y: 200},
							ConnectionPointOut: &plcopen.ConnectionPointOut{},
						},
					},
					OutVariables: []plcopen.BodyFBDOutVariable{
						{
							LocalID:    3,
							Expression: "output1",
							Position:   &plcopen.Position{X: 200, Y: 200},
							ConnectionPointIn: &plcopen.ConnectionPointIn{
								Connections: []plcopen.Connection{
									{RefLocalID: 1},
								},
							},
						},
					},
				},
			},
		},
		{
			"LD",
			&plcopen.Body{
				LD: &plcopen.BodyLD{
					LeftPowerRails: []plcopen.BodyLDLeftPowerRail{
						{
							LocalID:            1,
							Position:           &plcopen.Position{X: 0, Y: 0},
							ConnectionPointOut: &plcopen.ConnectionPointOut{},
						},
					},
					Contacts: []plcopen.BodyLDContact{
						{
							LocalID:            2,
							Variable:           "input1",
							Position:           &plcopen.Position{X: 50, Y: 0},
							ConnectionPointIn:  &plcopen.ConnectionPointIn{},
							ConnectionPointOut: &plcopen.ConnectionPointOut{},
						},
					},
					Coils: []plcopen.BodyLDCoil{
						{
							LocalID:            3,
							Variable:           "output1",
							Position:           &plcopen.Position{X: 100, Y: 0},
							ConnectionPointIn:  &plcopen.ConnectionPointIn{},
							ConnectionPointOut: &plcopen.ConnectionPointOut{},
						},
					},
					RightPowerRails: []plcopen.BodyLDRightPowerRail{
						{
							LocalID:           4,
							Position:          &plcopen.Position{X: 150, Y: 0},
							ConnectionPointIn: &plcopen.ConnectionPointIn{},
						},
					},
				},
			},
		},
		{
			"SFC",
			&plcopen.Body{
				SFC: &plcopen.BodySFC{
					Steps: []plcopen.BodySFCStep{
						{
							LocalID:     1,
							Name:        "Step1",
							Position:    &plcopen.Position{X: 100, Y: 50},
							InitialStep: boolPtr(true),
							ConnectionPointIn: &plcopen.BodySFCStepConnectionPointIn{
								Connections: []plcopen.Connection{
									{RefLocalID: 2},
								},
							},
							ConnectionPointOut: &plcopen.BodySFCStepConnectionPointOut{},
						},
					},
					Transitions: []plcopen.BodySFCTransition{
						{
							LocalID:            2,
							Position:           &plcopen.Position{X: 100, Y: 100},
							ConnectionPointIn:  &plcopen.ConnectionPointIn{},
							ConnectionPointOut: &plcopen.ConnectionPointOut{},
							Condition: &plcopen.BodySFCTransitionCondition{
								Inline: &plcopen.BodySFCTransitionConditionInline{
									Name: "condition1",
									Body: &plcopen.Body{
										ST: &plcopen.BodyST{
											Xhtml: "TRUE",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			"IL",
			&plcopen.Body{
				IL: &plcopen.BodyIL{
					XMLNSXhtml: "http://www.w3.org/1999/xhtml",
					Xhtml:      "LD input1\nST output1",
				},
			},
		},
		{
			"ST",
			&plcopen.Body{
				ST: &plcopen.BodyST{
					XMLNSXhtml: "http://www.w3.org/1999/xhtml",
					Xhtml:      "output1 := input1;",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test marshaling
			data, err := xml.Marshal(tt.body)
			if err != nil {
				t.Fatalf("Failed to marshal Body %s: %v", tt.name, err)
			}

			// Test unmarshaling
			var unmarshaled plcopen.Body
			err = xml.Unmarshal(data, &unmarshaled)
			if err != nil {
				t.Fatalf("Failed to unmarshal Body %s: %v", tt.name, err)
			}
		})
	}
}

// TestComplexStructures tests complex nested structures
func TestComplexStructures(t *testing.T) {
	// Test ProjectTypesPOU with all optional fields
	pou := &plcopen.ProjectTypesPOU{
		Name:    "ComplexPOU",
		POUType: plcopen.POUTypeFunctionBlock,
		Interface: &plcopen.ProjectTypesPOUInterface{
			ReturnType: &plcopen.DataType{BOOL: &struct{}{}},
			LocalVars: &plcopen.VarList{
				Variables: []plcopen.VarListVariable{
					{
						Name:    "localVar",
						Address: "%MW100",
						Type:    &plcopen.DataType{INT: &struct{}{}},
						InitialValue: &plcopen.Value{
							SimpleValue: &plcopen.ValueSimpleValue{Value: "0"},
						},
						Documentation: []byte("Local variable"),
					},
				},
				Documentation: []byte("Local variables"),
			},
			InputVars: &plcopen.VarList{
				Variables: []plcopen.VarListVariable{
					{
						Name: "inputVar",
						Type: &plcopen.DataType{BOOL: &struct{}{}},
					},
				},
			},
			InOutVars: &plcopen.VarList{
				Variables: []plcopen.VarListVariable{
					{
						Name: "inOutVar",
						Type: &plcopen.DataType{REAL: &struct{}{}},
					},
				},
			},
			OutputVars: &plcopen.VarList{
				Variables: []plcopen.VarListVariable{
					{
						Name: "outputVar",
						Type: &plcopen.DataType{DINT: &struct{}{}},
					},
				},
			},
			ExternalVars:  &plcopen.VarList{},
			GlobalVars:    &plcopen.VarList{},
			TempVars:      &plcopen.VarList{},
			AccessVars:    &plcopen.VarList{},
			Documentation: []byte("Interface documentation"),
		},
		Actions: []plcopen.ProjectTypesPOUAction{
			{
				Name: "Action1",
				Body: &plcopen.Body{
					ST: &plcopen.BodyST{
						Xhtml: "// Action code",
					},
				},
				Documentation: []byte("Action documentation"),
			},
		},
		Transitions: []plcopen.ProjectTypesPOUTransition{
			{
				Name: "Transition1",
				Body: &plcopen.Body{
					ST: &plcopen.BodyST{
						Xhtml: "// Transition code",
					},
				},
				Documentation: []byte("Transition documentation"),
			},
		},
		Body: &plcopen.Body{
			FBD: &plcopen.BodyFBD{
				ActionBlocks: []plcopen.BodyFBDActionBlock{
					{
						LocalID:  1,
						Position: &plcopen.Position{X: 100, Y: 100},
						Actions: []plcopen.BodyFBDActionBlockAction{
							{Reference: &plcopen.BodyFBDActionBlockActionReference{
								Name: "Action1",
							},
								Qualifier: actionQualifierPtr(plcopen.BodyFBDActionBlockActionQualifierP),
							},
						},
					},
				},
				Comments: []plcopen.BodyFBDComment{
					{
						LocalID:  2,
						Position: &plcopen.Position{X: 200, Y: 200},
						Content:  "This is a comment",
					},
				},
				Connectors: []plcopen.BodyFBDConnector{
					{
						LocalID:            3,
						Name:               "Connector1",
						Position:           &plcopen.Position{X: 150, Y: 150},
						ConnectionPointOut: &plcopen.ConnectionPointOut{},
					},
				},
				Continuations: []plcopen.BodyFBDContinuation{
					{
						LocalID:           4,
						Name:              "Connector1",
						Position:          &plcopen.Position{X: 250, Y: 150},
						ConnectionPointIn: &plcopen.ConnectionPointIn{},
					},
				},
				InOutVariables: []plcopen.BodyFBDInOutVariable{
					{
						LocalID:            5,
						Expression:         "inOutVar",
						Position:           &plcopen.Position{X: 300, Y: 300},
						ConnectionPointIn:  &plcopen.ConnectionPointIn{},
						ConnectionPointOut: &plcopen.ConnectionPointOut{},
					},
				},
				Jumps: []plcopen.BodyFBDJump{
					{
						LocalID:  6,
						Label:    "LABEL1",
						Position: &plcopen.Position{X: 400, Y: 400},
					},
				},
				Labels: []plcopen.BodyFBDLabel{
					{
						LocalID:  7,
						Label:    "LABEL1",
						Position: &plcopen.Position{X: 500, Y: 500},
					},
				},
				Returns: []plcopen.BodyFBDReturn{
					{
						LocalID:  8,
						Position: &plcopen.Position{X: 600, Y: 600},
					},
				},
			},
		},
		Documentation: []byte("POU documentation"),
	}

	// Test marshaling
	data, err := xml.MarshalIndent(pou, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal complex POU: %v", err)
	}

	// Test unmarshaling
	var unmarshaled plcopen.ProjectTypesPOU
	err = xml.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal complex POU: %v", err)
	}

	// Verify some complex fields
	if unmarshaled.Name != pou.Name {
		t.Errorf("Name = %v, want %v", unmarshaled.Name, pou.Name)
	}
	if len(unmarshaled.Actions) != len(pou.Actions) {
		t.Errorf("Actions length = %v, want %v", len(unmarshaled.Actions), len(pou.Actions))
	}
	if len(unmarshaled.Body.FBD.ActionBlocks) != len(pou.Body.FBD.ActionBlocks) {
		t.Errorf("ActionBlocks length = %v, want %v", len(unmarshaled.Body.FBD.ActionBlocks), len(pou.Body.FBD.ActionBlocks))
	}
}

// TestAllFBDElements tests all FBD element types
func TestAllFBDElements(t *testing.T) {
	fbd := &plcopen.BodyFBD{
		Blocks: []plcopen.BodyFBDBlock{
			{
				LocalID:          1,
				TypeName:         "AND",
				InstanceName:     stringPtr("andBlock"),
				Position:         &plcopen.Position{X: 100, Y: 100},
				Width:            float64Ptr(50),
				Height:           float64Ptr(30),
				ExecutionOrderID: uint64Ptr(1),
				InputVariables: []plcopen.BodyFBDBlockVariable{
					{
						FormalParameter: "IN1",
						ConnectionPointIn: &plcopen.ConnectionPointIn{
							Connections: []plcopen.Connection{
								{
									RefLocalID:      2,
									FormalParameter: stringPtr("OUT"),
								},
							},
						},
					},
				},
				InOutVariables: []plcopen.BodyFBDBlockVariable2{
					{
						FormalParameter:    "INOUT1",
						ConnectionPointIn:  &plcopen.ConnectionPointIn{},
						ConnectionPointOut: &plcopen.ConnectionPointOut{},
					},
				},
				OutputVariables: []plcopen.BodyFBDBlockVariable1{
					{
						FormalParameter:    "OUT",
						ConnectionPointOut: &plcopen.ConnectionPointOut{},
					},
				},
				Documentation: []byte("Block documentation"),
			},
		},
		ActionBlocks: []plcopen.BodyFBDActionBlock{
			{
				LocalID:  10,
				Position: &plcopen.Position{X: 200, Y: 200},
				Width:    float64Ptr(80),
				Height:   float64Ptr(60),
				Actions: []plcopen.BodyFBDActionBlockAction{
					{
						Reference: &plcopen.BodyFBDActionBlockActionReference{
							Name: "TestAction",
						},
						Qualifier: actionQualifierPtr(plcopen.BodyFBDActionBlockActionQualifierP1),
					},
					{
						Inline: &plcopen.BodyFBDActionBlockActionInline{
							Name: "InlineAction",
							Body: &plcopen.Body{
								ST: &plcopen.BodyST{
									Xhtml: "// Inline action code",
								},
							},
						},
						Qualifier: actionQualifierPtr(plcopen.BodyFBDActionBlockActionQualifierN),
					},
				},
				Documentation: []byte("Action block documentation"),
			},
		},
	}

	// Test marshaling
	data, err := xml.MarshalIndent(fbd, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal FBD: %v", err)
	}

	// Test unmarshaling
	var unmarshaled plcopen.BodyFBD
	err = xml.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal FBD: %v", err)
	}

	// Verify structure
	if len(unmarshaled.Blocks) != 1 {
		t.Errorf("Blocks length = %v, want 1", len(unmarshaled.Blocks))
	}
	if len(unmarshaled.ActionBlocks) != 1 {
		t.Errorf("ActionBlocks length = %v, want 1", len(unmarshaled.ActionBlocks))
	}
}

// TestLDElements tests Ladder Diagram elements
func TestLDElements(t *testing.T) {
	ld := &plcopen.BodyLD{
		LeftPowerRails: []plcopen.BodyLDLeftPowerRail{
			{
				LocalID:            1,
				Position:           &plcopen.Position{X: 0, Y: 100},
				Width:              float64Ptr(10),
				Height:             float64Ptr(100),
				ConnectionPointOut: &plcopen.ConnectionPointOut{},
				Documentation:      []byte("Left power rail"),
			},
		},
		Contacts: []plcopen.BodyLDContact{
			{
				LocalID:            2,
				Variable:           "input1",
				Position:           &plcopen.Position{X: 50, Y: 100},
				Width:              float64Ptr(30),
				Height:             float64Ptr(20),
				ConnectionPointIn:  &plcopen.ConnectionPointIn{},
				ConnectionPointOut: &plcopen.ConnectionPointOut{},
				EdgeModifier:       edgeModifierPtr(plcopen.EdgeModifierTypeRising),
				Negated:            boolPtr(false),
				Documentation:      []byte("Contact documentation"),
			},
		},
		Coils: []plcopen.BodyLDCoil{
			{
				LocalID:            3,
				Variable:           "output1",
				Position:           &plcopen.Position{X: 100, Y: 100},
				Width:              float64Ptr(30),
				Height:             float64Ptr(20),
				ConnectionPointIn:  &plcopen.ConnectionPointIn{},
				ConnectionPointOut: &plcopen.ConnectionPointOut{}, EdgeModifier: edgeModifierPtr(plcopen.EdgeModifierTypeFalling),
				StorageModifier: storageModifierPtr(plcopen.StorageModifierTypeSet),
				Negated:         boolPtr(true),
				Documentation:   []byte("Coil documentation"),
			},
		},
		RightPowerRails: []plcopen.BodyLDRightPowerRail{
			{
				LocalID:           4,
				Position:          &plcopen.Position{X: 150, Y: 100},
				Width:             float64Ptr(10),
				Height:            float64Ptr(100),
				ConnectionPointIn: &plcopen.ConnectionPointIn{},
				Documentation:     []byte("Right power rail"),
			},
		},
	}

	// Test marshaling
	data, err := xml.MarshalIndent(ld, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal LD: %v", err)
	}

	// Test unmarshaling
	var unmarshaled plcopen.BodyLD
	err = xml.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal LD: %v", err)
	}

	// Verify structure
	if len(unmarshaled.LeftPowerRails) != 1 {
		t.Errorf("LeftPowerRails length = %v, want 1", len(unmarshaled.LeftPowerRails))
	}
	if len(unmarshaled.Contacts) != 1 {
		t.Errorf("Contacts length = %v, want 1", len(unmarshaled.Contacts))
	}
	if len(unmarshaled.Coils) != 1 {
		t.Errorf("Coils length = %v, want 1", len(unmarshaled.Coils))
	}
	if len(unmarshaled.RightPowerRails) != 1 {
		t.Errorf("RightPowerRails length = %v, want 1", len(unmarshaled.RightPowerRails))
	}
}

// TestSFCElements tests Sequential Function Chart elements
func TestSFCElements(t *testing.T) {
	sfc := &plcopen.BodySFC{
		Steps: []plcopen.BodySFCStep{
			{
				LocalID:     1,
				Name:        "START",
				Position:    &plcopen.Position{X: 100, Y: 50},
				Width:       float64Ptr(80),
				Height:      float64Ptr(40),
				InitialStep: boolPtr(true),
				ConnectionPointIn: &plcopen.BodySFCStepConnectionPointIn{
					Connections: []plcopen.Connection{
						{RefLocalID: 2},
					},
				},
				ConnectionPointOut: &plcopen.BodySFCStepConnectionPointOut{
					FormalParameter: stringPtr("OUT"),
				},
				ConnectionPointOutAction: &plcopen.BodySFCStepConnectionPointOutAction{
					FormalParameter: stringPtr("ACTION"),
				},
				Documentation: []byte("Initial step"),
			},
		},
		Transitions: []plcopen.BodySFCTransition{
			{
				LocalID:            2,
				Position:           &plcopen.Position{X: 100, Y: 150},
				Width:              float64Ptr(60),
				Height:             float64Ptr(20),
				Priority:           uint64Ptr(1),
				ConnectionPointIn:  &plcopen.ConnectionPointIn{},
				ConnectionPointOut: &plcopen.ConnectionPointOut{},
				Condition: &plcopen.BodySFCTransitionCondition{
					Reference: &plcopen.BodySFCTransitionConditionReference{
						Name: "condition1",
					},
				},
				Documentation: []byte("Transition documentation"),
			},
		},
	}

	// Test marshaling
	data, err := xml.MarshalIndent(sfc, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal SFC: %v", err)
	}

	// Test unmarshaling
	var unmarshaled plcopen.BodySFC
	err = xml.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal SFC: %v", err)
	}

	// Verify structure
	if len(unmarshaled.Steps) != 1 {
		t.Errorf("Steps length = %v, want 1", len(unmarshaled.Steps))
	}
	if len(unmarshaled.Transitions) != 1 {
		t.Errorf("Transitions length = %v, want 1", len(unmarshaled.Transitions))
	}
}

// TestProjectInstances tests project instance structures
func TestProjectInstances(t *testing.T) {
	interval := "T#100ms"
	single := "T#1s"

	instances := &plcopen.ProjectInstances{
		Configurations: []plcopen.ProjectInstancesConfiguration{
			{
				Name: "Configuration1",
				Resources: []plcopen.ProjectInstancesConfigurationResource{
					{
						Name: "Resource1",
						Tasks: []plcopen.ProjectInstancesConfigurationResourceTask{
							{
								Name:     "CyclicTask",
								Priority: 1,
								Interval: &interval,
								POUInstances: []plcopen.POUInstance{
									{
										Name:          "MainProgram",
										TypeName:      "MAIN",
										Documentation: []byte("Main program instance"),
									},
								},
							},
							{
								Name:     "SingleTask",
								Priority: 2,
								Single:   &single,
								POUInstances: []plcopen.POUInstance{
									{
										Name:     "InitProgram",
										TypeName: "INIT",
									},
								},
							},
						},
						GlobalVars: &plcopen.VarList{
							Variables: []plcopen.VarListVariable{
								{
									Name: "globalVar1",
									Type: &plcopen.DataType{INT: &struct{}{}},
								},
							},
						},
						POUInstances: []plcopen.POUInstance{
							{
								Name:     "FBInstance",
								TypeName: "MyFunctionBlock",
							},
						},
						Documentation: []byte("Resource documentation"),
					},
				},
				GlobalVars: &plcopen.VarList{
					Variables: []plcopen.VarListVariable{
						{
							Name: "configVar1",
							Type: &plcopen.DataType{BOOL: &struct{}{}},
						},
					},
				},
				Documentation: []byte("Configuration documentation"),
			},
		},
	}

	// Test marshaling
	data, err := xml.MarshalIndent(instances, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal ProjectInstances: %v", err)
	}

	// Test unmarshaling
	var unmarshaled plcopen.ProjectInstances
	err = xml.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal ProjectInstances: %v", err)
	}

	// Verify structure
	if len(unmarshaled.Configurations) != 1 {
		t.Errorf("Configurations length = %v, want 1", len(unmarshaled.Configurations))
	}

	config := unmarshaled.Configurations[0]
	if len(config.Resources) != 1 {
		t.Errorf("Resources length = %v, want 1", len(config.Resources))
	}

	resource := config.Resources[0]
	if len(resource.Tasks) != 2 {
		t.Errorf("Tasks length = %v, want 2", len(resource.Tasks))
	}
}

// TestTypeAliases tests all type aliases
func TestTypeAliases(t *testing.T) {
	// Test that type aliases work correctly
	var localVars plcopen.ProjectTypesPOUInterfaceLocalVars = plcopen.VarList{
		Variables: []plcopen.VarListVariable{
			{Name: "var1", Type: &plcopen.DataType{INT: &struct{}{}}},
		},
	}

	var inputVars plcopen.ProjectTypesPOUInterfaceInputVars = plcopen.VarList{
		Variables: []plcopen.VarListVariable{
			{Name: "var2", Type: &plcopen.DataType{BOOL: &struct{}{}}},
		},
	}

	var inOutVars plcopen.ProjectTypesPOUInterfaceInOutVars = plcopen.VarList{}
	var outputVars plcopen.ProjectTypesPOUInterfaceOutputVars = plcopen.VarList{}
	var externalVars plcopen.ProjectTypesPOUInterfaceExternalVars = plcopen.VarList{}
	var globalVars plcopen.ProjectTypesPOUInterfaceGlobalVars = plcopen.VarList{}
	var tempVars plcopen.ProjectTypesPOUInterfaceTempVars = plcopen.VarList{}

	// Test that they can be used interchangeably with VarList
	interface_ := &plcopen.ProjectTypesPOUInterface{
		LocalVars:    &localVars,
		InputVars:    &inputVars,
		InOutVars:    &inOutVars,
		OutputVars:   &outputVars,
		ExternalVars: &externalVars,
		GlobalVars:   &globalVars,
		TempVars:     &tempVars,
	}

	// Test marshaling
	data, err := xml.MarshalIndent(interface_, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal interface with type aliases: %v", err)
	}

	// Test unmarshaling
	var unmarshaled plcopen.ProjectTypesPOUInterface
	err = xml.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal interface with type aliases: %v", err)
	}

	// Verify that local vars were preserved
	if len(unmarshaled.LocalVars.Variables) != 1 {
		t.Errorf("LocalVars length = %v, want 1", len(unmarshaled.LocalVars.Variables))
	}
	if unmarshaled.LocalVars.Variables[0].Name != "var1" {
		t.Errorf("LocalVars variable name = %v, want var1", unmarshaled.LocalVars.Variables[0].Name)
	}
}

// TestRangeTypes tests signed and unsigned ranges
func TestRangeTypes(t *testing.T) {
	// Test RangeSigned
	signedRange := &plcopen.RangeSigned{
		Lower: -100,
		Upper: 100,
	}

	data, err := xml.Marshal(signedRange)
	if err != nil {
		t.Fatalf("Failed to marshal RangeSigned: %v", err)
	}

	var unmarshaledSigned plcopen.RangeSigned
	err = xml.Unmarshal(data, &unmarshaledSigned)
	if err != nil {
		t.Fatalf("Failed to unmarshal RangeSigned: %v", err)
	}

	if unmarshaledSigned.Lower != signedRange.Lower {
		t.Errorf("RangeSigned Lower = %v, want %v", unmarshaledSigned.Lower, signedRange.Lower)
	}
	if unmarshaledSigned.Upper != signedRange.Upper {
		t.Errorf("RangeSigned Upper = %v, want %v", unmarshaledSigned.Upper, signedRange.Upper)
	}

	// Test RangeUnsigned
	unsignedRange := &plcopen.RangeUnsigned{
		Lower: 0,
		Upper: 255,
	}

	data, err = xml.Marshal(unsignedRange)
	if err != nil {
		t.Fatalf("Failed to marshal RangeUnsigned: %v", err)
	}

	var unmarshaledUnsigned plcopen.RangeUnsigned
	err = xml.Unmarshal(data, &unmarshaledUnsigned)
	if err != nil {
		t.Fatalf("Failed to unmarshal RangeUnsigned: %v", err)
	}

	if unmarshaledUnsigned.Lower != unsignedRange.Lower {
		t.Errorf("RangeUnsigned Lower = %v, want %v", unmarshaledUnsigned.Lower, unsignedRange.Lower)
	}
	if unmarshaledUnsigned.Upper != unsignedRange.Upper {
		t.Errorf("RangeUnsigned Upper = %v, want %v", unmarshaledUnsigned.Upper, unsignedRange.Upper)
	}
}

// TestFullProjectXML tests marshaling and unmarshaling of a complete project
func TestFullProjectXML(t *testing.T) {
	// Create a complete project structure
	creationTime := time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC)
	modTime := time.Date(2023, 1, 1, 11, 0, 0, 0, time.UTC)

	project := &plcopen.Project{
		FileHeader: &plcopen.ProjectFileHeader{
			CompanyName:        "Test Corp",
			CompanyURL:         "https://test.com",
			ProductName:        "TestPLC",
			ProductVersion:     "2.0",
			ProductRelease:     "Beta",
			CreationDateTime:   creationTime,
			ContentDescription: "Test project for unit testing",
		},
		ContentHeader: &plcopen.ProjectContentHeader{
			Name:                 "CompleteTest",
			Version:              "1.5",
			ModificationDateTime: &modTime,
			Organization:         "Test Organization",
			Author:               "Test Author",
			Language:             "en-US",
			Comment:              "Complete test project",
			CoordinateInfo: &plcopen.ProjectContentHeaderCoordinateInfo{
				PageSize: &plcopen.ProjectContentHeaderCoordinateInfoPageSize{X: 210, Y: 297},
				FBD:      &plcopen.ProjectContentHeaderCoordinateInfoFBD{Scaling: &plcopen.ProjectContentHeaderCoordinateInfoFBDScaling{X: 1.0, Y: 1.0}},
				LD:       &plcopen.ProjectContentHeaderCoordinateInfoLD{Scaling: &plcopen.ProjectContentHeaderCoordinateInfoLDScaling{X: 1.0, Y: 1.0}},
				SFC:      &plcopen.ProjectContentHeaderCoordinateInfoSFC{Scaling: &plcopen.ProjectContentHeaderCoordinateInfoSFCScaling{X: 1.0, Y: 1.0}},
			},
		},
		Types: &plcopen.ProjectTypes{
			DataTypes: []plcopen.ProjectTypesDataType{
				{
					Name: "MyEnum",
					BaseType: &plcopen.DataType{
						Enum: &plcopen.DataTypeEnum{
							Values: &plcopen.DataTypeEnumValues{
								Values: []plcopen.DataTypeEnumValuesValue{
									{Name: "VALUE1", Documentation: []byte("First value")},
									{Name: "VALUE2", Documentation: []byte("Second value")},
								},
							},
							BaseType: &plcopen.DataType{INT: &struct{}{}},
						},
					},
					Documentation: []byte("Custom enumeration"),
				},
			},
			POUs: []plcopen.ProjectTypesPOU{
				{
					Name:    "TestFunction",
					POUType: plcopen.POUTypeFunction,
					Interface: &plcopen.ProjectTypesPOUInterface{
						ReturnType: &plcopen.DataType{BOOL: &struct{}{}},
						InputVars: &plcopen.VarList{
							Variables: []plcopen.VarListVariable{
								{
									Name: "Input1",
									Type: &plcopen.DataType{INT: &struct{}{}},
									InitialValue: &plcopen.Value{
										SimpleValue: &plcopen.ValueSimpleValue{Value: "0"},
									},
								},
							},
						},
						OutputVars: &plcopen.VarList{
							Variables: []plcopen.VarListVariable{
								{
									Name: "Output1",
									Type: &plcopen.DataType{BOOL: &struct{}{}},
								},
							},
						},
					},
					Body: &plcopen.Body{
						ST: &plcopen.BodyST{
							XMLNSXhtml: "http://www.w3.org/1999/xhtml",
							Xhtml:      "Output1 := Input1 > 0;",
						},
					},
					Documentation: []byte("Test function"),
				},
			},
		},
		Instances: &plcopen.ProjectInstances{
			Configurations: []plcopen.ProjectInstancesConfiguration{
				{
					Name: "Config",
					Resources: []plcopen.ProjectInstancesConfigurationResource{
						{
							Name: "CPU",
							POUInstances: []plcopen.POUInstance{
								{
									Name:     "Program1",
									TypeName: "TestFunction",
								},
							},
						},
					},
				},
			},
		},
	}

	// Test marshaling to XML
	xmlData, err := xml.MarshalIndent(project, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal complete project: %v", err)
	}

	// Verify XML contains expected elements
	xmlStr := string(xmlData)
	expectedElements := []string{
		"project",
		"fileHeader",
		"contentHeader",
		"types",
		"instances",
		"TestFunction",
		"MyEnum",
		"Config",
	}

	for _, elem := range expectedElements {
		if !strings.Contains(xmlStr, elem) {
			t.Errorf("XML does not contain expected element: %s", elem)
		}
	}

	// Test unmarshaling from XML
	var unmarshaledProject plcopen.Project
	err = xml.Unmarshal(xmlData, &unmarshaledProject)
	if err != nil {
		t.Fatalf("Failed to unmarshal complete project: %v", err)
	}

	// Verify critical fields
	if unmarshaledProject.FileHeader.CompanyName != project.FileHeader.CompanyName {
		t.Errorf("CompanyName mismatch")
	}
	if unmarshaledProject.ContentHeader.Name != project.ContentHeader.Name {
		t.Errorf("Project name mismatch")
	}
	if len(unmarshaledProject.Types.POUs) != len(project.Types.POUs) {
		t.Errorf("POUs count mismatch")
	}
	if len(unmarshaledProject.Types.DataTypes) != len(project.Types.DataTypes) {
		t.Errorf("DataTypes count mismatch")
	}
}

// Helper functions for pointer creation
func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func float64Ptr(f float64) *float64 {
	return &f
}

func uint64Ptr(u uint64) *uint64 {
	return &u
}

func edgeModifierPtr(e plcopen.EdgeModifierType) *plcopen.EdgeModifierType {
	return &e
}

func storageModifierPtr(s plcopen.StorageModifierType) *plcopen.StorageModifierType {
	return &s
}

func actionQualifierPtr(q plcopen.BodyFBDActionBlockActionQualifier) *plcopen.BodyFBDActionBlockActionQualifier {
	return &q
}
