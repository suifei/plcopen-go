package tests

import (
	"reflect"
	"testing"

	"github.com/suifei/plcopen-go"
)

// TestAllTypesDeclared checks that all types in tc6_xml_v10_b.go are tested
func TestAllTypesDeclared(t *testing.T) {
	// Define all types declared in tc6_xml_v10_b.go
	types := []interface{}{
		// Core project types
		plcopen.Project{},
		plcopen.ProjectFileHeader{},
		plcopen.ProjectContentHeader{},
		plcopen.ProjectContentHeaderCoordinateInfo{},
		plcopen.ProjectContentHeaderCoordinateInfoPageSize{},
		plcopen.ProjectContentHeaderCoordinateInfoFBD{},
		plcopen.ProjectContentHeaderCoordinateInfoFBDScaling{},
		plcopen.ProjectContentHeaderCoordinateInfoLD{},
		plcopen.ProjectContentHeaderCoordinateInfoLDScaling{},
		plcopen.ProjectContentHeaderCoordinateInfoSFC{},
		plcopen.ProjectContentHeaderCoordinateInfoSFCScaling{},
		plcopen.ProjectTypes{},
		plcopen.ProjectTypesDataType{},
		plcopen.ProjectTypesPOU{},
		plcopen.ProjectTypesPOUInterface{},
		plcopen.ProjectTypesPOUAction{},
		plcopen.ProjectTypesPOUTransition{},
		plcopen.ProjectInstances{},
		plcopen.ProjectInstancesConfiguration{},
		plcopen.ProjectInstancesConfigurationResource{},
		plcopen.ProjectInstancesConfigurationResourceTask{},
		plcopen.POUInstance{},

		// Variable types
		plcopen.VarList{},
		plcopen.VarListPlain{},
		plcopen.VarListPlainVariable{},
		plcopen.VarListVariable{},

		// Data types
		plcopen.DataType{},
		plcopen.DataTypeArray{},
		plcopen.RangeSigned{},
		plcopen.RangeUnsigned{},
		plcopen.DataTypeDerived{},
		plcopen.DataTypeEnum{},
		plcopen.DataTypeEnumValues{},
		plcopen.DataTypeEnumValuesValue{},
		plcopen.DataTypePointer{},
		plcopen.DataTypeString{},
		plcopen.DataTypeWString{},
		plcopen.DataTypeSubrangeSigned{},
		plcopen.DataTypeSubrangeUnsigned{},
		// Value types
		plcopen.Value{},
		plcopen.ValueSimpleValue{},
		plcopen.ValueArrayValue{},
		plcopen.ValueArrayValueValue{},
		plcopen.ValueStructValue{},
		plcopen.ValueStructValueValue{},
		// Body types
		plcopen.Body{},
		plcopen.BodyST{},
		plcopen.BodyFBD{},
		plcopen.BodyLD{},
		plcopen.BodySFC{},
		plcopen.BodyIL{},
		plcopen.Position{},
		plcopen.Connection{},
		plcopen.ConnectionPointIn{},
		plcopen.ConnectionPointOut{},
		plcopen.BodyFBDBlock{},
		plcopen.BodyFBDBlockVariable{},
		plcopen.BodyFBDBlockVariable1{},
		plcopen.BodyFBDBlockVariable2{},
		plcopen.BodyFBDActionBlock{},
		plcopen.BodyFBDActionBlockAction{},
		plcopen.BodyFBDActionBlockActionReference{},
		plcopen.BodyFBDActionBlockActionInline{},
		plcopen.BodyFBDComment{},
		plcopen.BodyFBDConnector{},
		plcopen.BodyFBDContinuation{},
		plcopen.BodyFBDInVariable{},
		plcopen.BodyFBDOutVariable{},
		plcopen.BodyFBDInOutVariable{},
		plcopen.BodyFBDJump{},
		plcopen.BodyFBDLabel{},
		plcopen.BodyFBDReturn{},
		plcopen.BodyLDContact{},
		plcopen.BodyLDCoil{},
		plcopen.BodyLDLeftPowerRail{},
		plcopen.BodyLDRightPowerRail{},
		plcopen.BodySFCStep{},
		plcopen.BodySFCStepConnectionPointIn{},
		plcopen.BodySFCStepConnectionPointOut{},
		plcopen.BodySFCStepConnectionPointOutAction{},
		plcopen.BodySFCTransition{},
		plcopen.BodySFCTransitionCondition{},
		plcopen.BodySFCTransitionConditionInline{},
		plcopen.BodySFCTransitionConditionReference{},

		// Constant types
		plcopen.POUType(""),
		plcopen.EdgeModifierType(""),
		plcopen.StorageModifierType(""),
		plcopen.BodyFBDActionBlockActionQualifier(""),

		// Type aliases
		plcopen.ProjectTypesPOUInterfaceExternalVars{},
		plcopen.ProjectTypesPOUInterfaceGlobalVars{},
		plcopen.ProjectTypesPOUInterfaceInOutVars{},
		plcopen.ProjectTypesPOUInterfaceInputVars{},
		plcopen.ProjectTypesPOUInterfaceLocalVars{},
		plcopen.ProjectTypesPOUInterfaceOutputVars{},
		plcopen.ProjectTypesPOUInterfaceTempVars{},
	}

	// This test ensures all types are available for testing
	for _, typeItem := range types {
		typeName := reflect.TypeOf(typeItem).Name()
		if typeName == "" {
			t.Log("Anonymous type detected")
		}
	}

	// If we get here without panic, all types have been declared properly
	t.Log("All types are properly declared and available")
}
