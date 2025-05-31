// Package plcopen provides Go structs for IEC 61131-3 PLCopen XML format
// generated from XSD schema TC6_XML_V10_B.xsd
package plcopen

import (
	"encoding/xml"
	"time"
)

// Project represents the root element of a PLCopen XML document
type Project struct {
	XMLName       xml.Name              `xml:"http://www.plcopen.org/xml/tc6.xsd project"`
	FileHeader    *ProjectFileHeader    `xml:"fileHeader"`
	ContentHeader *ProjectContentHeader `xml:"contentHeader"`
	Types         *ProjectTypes         `xml:"types"`
	Instances     *ProjectInstances     `xml:"instances"`
}

// ProjectFileHeader contains file metadata
type ProjectFileHeader struct {
	CompanyName        string    `xml:"companyName,attr"`
	CompanyURL         string    `xml:"companyURL,attr,omitempty"`
	ProductName        string    `xml:"productName,attr"`
	ProductVersion     string    `xml:"productVersion,attr"`
	ProductRelease     string    `xml:"productRelease,attr,omitempty"`
	CreationDateTime   time.Time `xml:"creationDateTime,attr"`
	ContentDescription string    `xml:"contentDescription,attr,omitempty"`
}

// ProjectContentHeader contains project content metadata
type ProjectContentHeader struct {
	Name                 string                              `xml:"name,attr"`
	Version              string                              `xml:"version,attr,omitempty"`
	ModificationDateTime *time.Time                          `xml:"modificationDateTime,attr,omitempty"`
	Organization         string                              `xml:"organization,attr,omitempty"`
	Author               string                              `xml:"author,attr,omitempty"`
	Language             string                              `xml:"language,attr,omitempty"`
	Comment              string                              `xml:"Comment,omitempty"`
	CoordinateInfo       *ProjectContentHeaderCoordinateInfo `xml:"coordinateInfo"`
}

// ProjectContentHeaderCoordinateInfo contains coordinate information
type ProjectContentHeaderCoordinateInfo struct {
	PageSize *ProjectContentHeaderCoordinateInfoPageSize `xml:"pageSize,omitempty"`
	FBD      *ProjectContentHeaderCoordinateInfoFBD      `xml:"fbd"`
	LD       *ProjectContentHeaderCoordinateInfoLD       `xml:"ld"`
	SFC      *ProjectContentHeaderCoordinateInfoSFC      `xml:"sfc"`
}

// ProjectContentHeaderCoordinateInfoPageSize represents page size
type ProjectContentHeaderCoordinateInfoPageSize struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

// ProjectContentHeaderCoordinateInfoFBD represents FBD coordinate info
type ProjectContentHeaderCoordinateInfoFBD struct {
	Scaling *ProjectContentHeaderCoordinateInfoFBDScaling `xml:"scaling"`
}

// ProjectContentHeaderCoordinateInfoFBDScaling represents FBD scaling
type ProjectContentHeaderCoordinateInfoFBDScaling struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

// ProjectContentHeaderCoordinateInfoLD represents LD coordinate info
type ProjectContentHeaderCoordinateInfoLD struct {
	Scaling *ProjectContentHeaderCoordinateInfoLDScaling `xml:"scaling"`
}

// ProjectContentHeaderCoordinateInfoLDScaling represents LD scaling
type ProjectContentHeaderCoordinateInfoLDScaling struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

// ProjectContentHeaderCoordinateInfoSFC represents SFC coordinate info
type ProjectContentHeaderCoordinateInfoSFC struct {
	Scaling *ProjectContentHeaderCoordinateInfoSFCScaling `xml:"scaling"`
}

// ProjectContentHeaderCoordinateInfoSFCScaling represents SFC scaling
type ProjectContentHeaderCoordinateInfoSFCScaling struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

// ProjectTypes contains type definitions
type ProjectTypes struct {
	DataTypes []ProjectTypesDataType `xml:"dataTypes>dataType,omitempty"`
	POUs      []ProjectTypesPOU      `xml:"pous>pou,omitempty"`
}

// ProjectTypesDataType represents a data type definition
type ProjectTypesDataType struct {
	Name          string    `xml:"name,attr"`
	BaseType      *DataType `xml:"baseType"`
	InitialValue  *Value    `xml:"initialValue,omitempty"`
	Documentation []byte    `xml:"documentation,omitempty"`
}

// ProjectTypesPOU represents a Program Organization Unit
type ProjectTypesPOU struct {
	Name          string                      `xml:"name,attr"`
	POUType       POUType                     `xml:"pouType,attr"`
	Interface     *ProjectTypesPOUInterface   `xml:"interface,omitempty"`
	Actions       []ProjectTypesPOUAction     `xml:"actions>action,omitempty"`
	Transitions   []ProjectTypesPOUTransition `xml:"transitions>transition,omitempty"`
	Body          *Body                       `xml:"body,omitempty"`
	Documentation []byte                      `xml:"documentation,omitempty"`
}

// POUType represents the type of POU
type POUType string

const (
	POUTypeFunction      POUType = "function"
	POUTypeFunctionBlock POUType = "functionBlock"
	POUTypeProgram       POUType = "program"
)

// ProjectTypesPOUInterface represents the interface of a POU
type ProjectTypesPOUInterface struct {
	ReturnType    *DataType                             `xml:"returnType,omitempty"`
	LocalVars     *ProjectTypesPOUInterfaceLocalVars    `xml:"localVars,omitempty"`
	InputVars     *ProjectTypesPOUInterfaceInputVars    `xml:"inputVars,omitempty"`
	InOutVars     *ProjectTypesPOUInterfaceInOutVars    `xml:"inOutVars,omitempty"`
	OutputVars    *ProjectTypesPOUInterfaceOutputVars   `xml:"outputVars,omitempty"`
	ExternalVars  *ProjectTypesPOUInterfaceExternalVars `xml:"externalVars,omitempty"`
	GlobalVars    *ProjectTypesPOUInterfaceGlobalVars   `xml:"globalVars,omitempty"`
	TempVars      *ProjectTypesPOUInterfaceTempVars     `xml:"tempVars,omitempty"`
	AccessVars    *VarList                              `xml:"accessVars,omitempty"`
	Documentation []byte                                `xml:"documentation,omitempty"`
}

// ProjectTypesPOUAction represents an action within a POU
type ProjectTypesPOUAction struct {
	Name          string `xml:"name,attr"`
	Body          *Body  `xml:"body"`
	Documentation []byte `xml:"documentation,omitempty"`
}

// ProjectTypesPOUTransition represents a transition within a POU
type ProjectTypesPOUTransition struct {
	Name          string `xml:"name,attr"`
	Body          *Body  `xml:"body"`
	Documentation []byte `xml:"documentation,omitempty"`
}

// ProjectInstances contains configuration and resource instances
type ProjectInstances struct {
	Configurations []ProjectInstancesConfiguration `xml:"configurations>configuration,omitempty"`
}

// ProjectInstancesConfiguration represents a configuration
type ProjectInstancesConfiguration struct {
	Name          string                                  `xml:"name,attr"`
	Resources     []ProjectInstancesConfigurationResource `xml:"resource,omitempty"`
	GlobalVars    *VarList                                `xml:"globalVars,omitempty"`
	Documentation []byte                                  `xml:"documentation,omitempty"`
}

// ProjectInstancesConfigurationResource represents a resource within a configuration
type ProjectInstancesConfigurationResource struct {
	Name          string                                      `xml:"name,attr"`
	Tasks         []ProjectInstancesConfigurationResourceTask `xml:"task,omitempty"`
	GlobalVars    *VarList                                    `xml:"globalVars,omitempty"`
	POUInstances  []POUInstance                               `xml:"pouInstance,omitempty"`
	Documentation []byte                                      `xml:"documentation,omitempty"`
}

// ProjectInstancesConfigurationResourceTask represents a task within a resource
type ProjectInstancesConfigurationResourceTask struct {
	Name         string        `xml:"name,attr"`
	Priority     uint64        `xml:"priority,attr"`
	Interval     *string       `xml:"interval,attr,omitempty"`
	Single       *string       `xml:"single,attr,omitempty"`
	POUInstances []POUInstance `xml:"pouInstance,omitempty"`
}

// POUInstance represents an instance of a POU
type POUInstance struct {
	Name          string `xml:"name,attr"`
	TypeName      string `xml:"type,attr"`
	Documentation []byte `xml:"documentation,omitempty"`
}

// VarList represents a list of variables
type VarList struct {
	Variables     []VarListVariable `xml:"variable,omitempty"`
	Documentation []byte            `xml:"documentation,omitempty"`
}

// VarListPlain represents a plain variable list (extends VarList)
type VarListPlain struct {
	Variables     []VarListPlainVariable `xml:"variable,omitempty"`
	Documentation []byte                 `xml:"documentation,omitempty"`
}

// VarListPlainVariable represents a plain variable
type VarListPlainVariable struct {
	Name          string    `xml:"name,attr"`
	Address       string    `xml:"address,attr,omitempty"`
	Type          *DataType `xml:"type"`
	InitialValue  *Value    `xml:"initialValue,omitempty"`
	Documentation []byte    `xml:"documentation,omitempty"`
}

// VarListVariable represents a variable with additional attributes
type VarListVariable struct {
	Name          string    `xml:"name,attr"`
	Address       string    `xml:"address,attr,omitempty"`
	Type          *DataType `xml:"type"`
	InitialValue  *Value    `xml:"initialValue,omitempty"`
	Documentation []byte    `xml:"documentation,omitempty"`
}

// DataType represents a data type with choice-style content
type DataType struct {
	// Basic data types (as empty elements)
	BOOL  *struct{} `xml:"BOOL,omitempty"`
	BYTE  *struct{} `xml:"BYTE,omitempty"`
	DATE  *struct{} `xml:"DATE,omitempty"`
	DINT  *struct{} `xml:"DINT,omitempty"`
	DT    *struct{} `xml:"DT,omitempty"`
	DWORD *struct{} `xml:"DWORD,omitempty"`
	INT   *struct{} `xml:"INT,omitempty"`
	LINT  *struct{} `xml:"LINT,omitempty"`
	LREAL *struct{} `xml:"LREAL,omitempty"`
	LWORD *struct{} `xml:"LWORD,omitempty"`
	REAL  *struct{} `xml:"REAL,omitempty"`
	SINT  *struct{} `xml:"SINT,omitempty"`
	TIME  *struct{} `xml:"TIME,omitempty"`
	TOD   *struct{} `xml:"TOD,omitempty"`
	UDINT *struct{} `xml:"UDINT,omitempty"`
	UINT  *struct{} `xml:"UINT,omitempty"`
	ULINT *struct{} `xml:"ULINT,omitempty"`
	USINT *struct{} `xml:"USINT,omitempty"`
	WORD  *struct{} `xml:"WORD,omitempty"`

	// Complex data types
	Array            *DataTypeArray            `xml:"array,omitempty"`
	Derived          *DataTypeDerived          `xml:"derived,omitempty"`
	Enum             *DataTypeEnum             `xml:"enum,omitempty"`
	Pointer          *DataTypePointer          `xml:"pointer,omitempty"`
	String           *DataTypeString           `xml:"string,omitempty"`
	Struct           *VarListPlain             `xml:"struct,omitempty"`
	SubrangeSigned   *DataTypeSubrangeSigned   `xml:"subrangeSigned,omitempty"`
	SubrangeUnsigned *DataTypeSubrangeUnsigned `xml:"subrangeUnsigned,omitempty"`
	WString          *DataTypeWString          `xml:"wstring,omitempty"`
}

// DataTypeArray represents an array data type
type DataTypeArray struct {
	Dimensions []RangeSigned `xml:"dimension"`
	BaseType   *DataType     `xml:"baseType"`
}

// RangeSigned represents a signed range
type RangeSigned struct {
	Lower int64 `xml:"lower,attr"`
	Upper int64 `xml:"upper,attr"`
}

// RangeUnsigned represents an unsigned range
type RangeUnsigned struct {
	Lower uint64 `xml:"lower,attr"`
	Upper uint64 `xml:"upper,attr"`
}

// DataTypeDerived represents a derived data type
type DataTypeDerived struct {
	Name string `xml:"name,attr"`
}

// DataTypeEnum represents an enumerated data type
type DataTypeEnum struct {
	Values   *DataTypeEnumValues `xml:"values"`
	BaseType *DataType           `xml:"baseType,omitempty"`
}

// DataTypeEnumValues represents enumerated values
type DataTypeEnumValues struct {
	Values []DataTypeEnumValuesValue `xml:"value"`
}

// DataTypeEnumValuesValue represents an enumerated value
type DataTypeEnumValuesValue struct {
	Name          string `xml:"name,attr"`
	Documentation []byte `xml:"documentation,omitempty"`
}

// DataTypePointer represents a pointer data type
type DataTypePointer struct {
	BaseType *DataType `xml:"baseType"`
}

// DataTypeString represents a string data type
type DataTypeString struct {
	Length *uint64 `xml:"length,attr,omitempty"`
}

// DataTypeWString represents a wide string data type
type DataTypeWString struct {
	Length *uint64 `xml:"length,attr,omitempty"`
}

// DataTypeSubrangeSigned represents a signed subrange data type
type DataTypeSubrangeSigned struct {
	Range    *RangeSigned `xml:"range"`
	BaseType *DataType    `xml:"baseType"`
}

// DataTypeSubrangeUnsigned represents an unsigned subrange data type
type DataTypeSubrangeUnsigned struct {
	Range    *RangeUnsigned `xml:"range"`
	BaseType *DataType      `xml:"baseType"`
}

// Value represents a value with choice-style content
type Value struct {
	SimpleValue *ValueSimpleValue `xml:"simpleValue,omitempty"`
	ArrayValue  *ValueArrayValue  `xml:"arrayValue,omitempty"`
	StructValue *ValueStructValue `xml:"structValue,omitempty"`
}

// ValueArrayValue represents an array value
type ValueArrayValue struct {
	Values []ValueArrayValueValue `xml:"value"`
}

// ValueArrayValueValue represents a value within an array
type ValueArrayValueValue struct {
	RepeatCount *uint64 `xml:"repetition,attr,omitempty"`
	Value       *Value  `xml:",innerxml"`
}

// ValueSimpleValue represents a simple value
type ValueSimpleValue struct {
	Value string `xml:"value,attr"`
}

// ValueStructValue represents a structured value
type ValueStructValue struct {
	Values []ValueStructValueValue `xml:"value"`
}

// ValueStructValueValue represents a value within a structure
type ValueStructValueValue struct {
	Member string `xml:"member,attr"`
	Value  *Value `xml:",innerxml"`
}

// Body represents a body with choice-style content for different languages
type Body struct {
	FBD *BodyFBD `xml:"FBD,omitempty"`
	LD  *BodyLD  `xml:"LD,omitempty"`
	SFC *BodySFC `xml:"SFC,omitempty"`
	IL  *BodyIL  `xml:"IL,omitempty"`
	ST  *BodyST  `xml:"ST,omitempty"`
}

// BodyFBD represents a Function Block Diagram body
type BodyFBD struct {
	Blocks         []BodyFBDBlock         `xml:"block,omitempty"`
	ActionBlocks   []BodyFBDActionBlock   `xml:"actionBlock,omitempty"`
	Comments       []BodyFBDComment       `xml:"comment,omitempty"`
	Connectors     []BodyFBDConnector     `xml:"connector,omitempty"`
	Continuations  []BodyFBDContinuation  `xml:"continuation,omitempty"`
	InVariables    []BodyFBDInVariable    `xml:"inVariable,omitempty"`
	OutVariables   []BodyFBDOutVariable   `xml:"outVariable,omitempty"`
	InOutVariables []BodyFBDInOutVariable `xml:"inOutVariable,omitempty"`
	Jumps          []BodyFBDJump          `xml:"jump,omitempty"`
	Labels         []BodyFBDLabel         `xml:"label,omitempty"`
	Returns        []BodyFBDReturn        `xml:"return,omitempty"`
}

// BodyLD represents a Ladder Diagram body
type BodyLD struct {
	Contacts        []BodyLDContact        `xml:"contact,omitempty"`
	Coils           []BodyLDCoil           `xml:"coil,omitempty"`
	LeftPowerRails  []BodyLDLeftPowerRail  `xml:"leftPowerRail,omitempty"`
	RightPowerRails []BodyLDRightPowerRail `xml:"rightPowerRail,omitempty"`
}

// BodySFC represents a Sequential Function Chart body
type BodySFC struct {
	Steps       []BodySFCStep       `xml:"step,omitempty"`
	Transitions []BodySFCTransition `xml:"transition,omitempty"`
}

// BodyIL represents an Instruction List body
type BodyIL struct {
	XMLNSXhtml string `xml:"xmlns:xhtml,attr"`
	Xhtml      string `xml:",innerxml"`
}

// BodyST represents a Structured Text body
type BodyST struct {
	XMLNSXhtml string `xml:"xmlns:xhtml,attr"`
	Xhtml      string `xml:",innerxml"`
}

// Position represents a position coordinate
type Position struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

// Connection represents a connection
type Connection struct {
	RefLocalID      uint64  `xml:"refLocalId,attr"`
	FormalParameter *string `xml:"formalParameter,attr,omitempty"`
}

// ConnectionPointIn represents an input connection point
type ConnectionPointIn struct {
	Connections []Connection `xml:"connection,omitempty"`
}

// ConnectionPointOut represents an output connection point
type ConnectionPointOut struct {
	FormalParameter *string `xml:"formalParameter,attr,omitempty"`
}

// EdgeModifierType represents edge modifier types
type EdgeModifierType string

const (
	EdgeModifierTypeNone    EdgeModifierType = "none"
	EdgeModifierTypeFalling EdgeModifierType = "falling"
	EdgeModifierTypeRising  EdgeModifierType = "rising"
)

// StorageModifierType represents storage modifier types
type StorageModifierType string

const (
	StorageModifierTypeNone  StorageModifierType = "none"
	StorageModifierTypeSet   StorageModifierType = "set"
	StorageModifierTypeReset StorageModifierType = "reset"
)

// BodyFBDActionBlockActionQualifier represents action block action qualifiers
type BodyFBDActionBlockActionQualifier string

const (
	BodyFBDActionBlockActionQualifierP1 BodyFBDActionBlockActionQualifier = "P1"
	BodyFBDActionBlockActionQualifierN  BodyFBDActionBlockActionQualifier = "N"
	BodyFBDActionBlockActionQualifierP0 BodyFBDActionBlockActionQualifier = "P0"
	BodyFBDActionBlockActionQualifierR  BodyFBDActionBlockActionQualifier = "R"
	BodyFBDActionBlockActionQualifierS  BodyFBDActionBlockActionQualifier = "S"
	BodyFBDActionBlockActionQualifierL  BodyFBDActionBlockActionQualifier = "L"
	BodyFBDActionBlockActionQualifierD  BodyFBDActionBlockActionQualifier = "D"
	BodyFBDActionBlockActionQualifierP  BodyFBDActionBlockActionQualifier = "P"
	BodyFBDActionBlockActionQualifierDS BodyFBDActionBlockActionQualifier = "DS"
	BodyFBDActionBlockActionQualifierDL BodyFBDActionBlockActionQualifier = "DL"
	BodyFBDActionBlockActionQualifierSD BodyFBDActionBlockActionQualifier = "SD"
	BodyFBDActionBlockActionQualifierSL BodyFBDActionBlockActionQualifier = "SL"
)

// BodyFBDBlock represents a block in FBD
type BodyFBDBlock struct {
	Position         *Position               `xml:"position,omitempty"`
	InputVariables   []BodyFBDBlockVariable  `xml:"inputVariables>variable,omitempty"`
	InOutVariables   []BodyFBDBlockVariable2 `xml:"inOutVariables>variable,omitempty"`
	OutputVariables  []BodyFBDBlockVariable1 `xml:"outputVariables>variable,omitempty"`
	Documentation    []byte                  `xml:"documentation,omitempty"`
	LocalID          uint64                  `xml:"localId,attr"`
	Width            *float64                `xml:"width,attr,omitempty"`
	Height           *float64                `xml:"height,attr,omitempty"`
	TypeName         string                  `xml:"typeName,attr"`
	InstanceName     *string                 `xml:"instanceName,attr,omitempty"`
	ExecutionOrderID *uint64                 `xml:"executionOrderId,attr,omitempty"`
}

// BodyFBDBlockVariable represents a variable in a block
type BodyFBDBlockVariable struct {
	FormalParameter   string             `xml:"formalParameter,attr"`
	ConnectionPointIn *ConnectionPointIn `xml:"connectionPointIn,omitempty"`
}

// BodyFBDBlockVariable1 represents an output variable in a block
type BodyFBDBlockVariable1 struct {
	FormalParameter    string              `xml:"formalParameter,attr"`
	ConnectionPointOut *ConnectionPointOut `xml:"connectionPointOut,omitempty"`
}

// BodyFBDBlockVariable2 represents an in-out variable in a block
type BodyFBDBlockVariable2 struct {
	FormalParameter    string              `xml:"formalParameter,attr"`
	ConnectionPointIn  *ConnectionPointIn  `xml:"connectionPointIn,omitempty"`
	ConnectionPointOut *ConnectionPointOut `xml:"connectionPointOut,omitempty"`
}

// BodyFBDActionBlock represents an action block in FBD
type BodyFBDActionBlock struct {
	Position      *Position                  `xml:"position,omitempty"`
	Actions       []BodyFBDActionBlockAction `xml:"action,omitempty"`
	Documentation []byte                     `xml:"documentation,omitempty"`
	LocalID       uint64                     `xml:"localId,attr"`
	Width         *float64                   `xml:"width,attr,omitempty"`
	Height        *float64                   `xml:"height,attr,omitempty"`
}

// BodyFBDActionBlockAction represents an action in an action block
type BodyFBDActionBlockAction struct {
	Reference *BodyFBDActionBlockActionReference `xml:"reference,omitempty"`
	Inline    *BodyFBDActionBlockActionInline    `xml:"inline,omitempty"`
	Qualifier *BodyFBDActionBlockActionQualifier `xml:"qualifier,attr,omitempty"`
}

// BodyFBDActionBlockActionReference represents an action reference
type BodyFBDActionBlockActionReference struct {
	Name string `xml:"name,attr"`
}

// BodyFBDActionBlockActionInline represents an inline action
type BodyFBDActionBlockActionInline struct {
	Body *Body  `xml:",innerxml"`
	Name string `xml:"name,attr"`
}

// BodyFBDComment represents a comment in FBD
type BodyFBDComment struct {
	Position      *Position `xml:"position,omitempty"`
	Content       string    `xml:"content,omitempty"`
	Documentation []byte    `xml:"documentation,omitempty"`
	LocalID       uint64    `xml:"localId,attr"`
	Height        *float64  `xml:"height,attr,omitempty"`
	Width         *float64  `xml:"width,attr,omitempty"`
}

// BodyFBDConnector represents a connector in FBD
type BodyFBDConnector struct {
	Position           *Position           `xml:"position,omitempty"`
	ConnectionPointOut *ConnectionPointOut `xml:"connectionPointOut,omitempty"`
	Documentation      []byte              `xml:"documentation,omitempty"`
	Name               string              `xml:"name,attr"`
	LocalID            uint64              `xml:"localId,attr"`
}

// BodyFBDContinuation represents a continuation in FBD
type BodyFBDContinuation struct {
	Position          *Position          `xml:"position,omitempty"`
	ConnectionPointIn *ConnectionPointIn `xml:"connectionPointIn,omitempty"`
	Documentation     []byte             `xml:"documentation,omitempty"`
	Name              string             `xml:"name,attr"`
	LocalID           uint64             `xml:"localId,attr"`
}

// BodyFBDInVariable represents an input variable in FBD
type BodyFBDInVariable struct {
	Position           *Position           `xml:"position,omitempty"`
	ConnectionPointOut *ConnectionPointOut `xml:"connectionPointOut,omitempty"`
	Documentation      []byte              `xml:"documentation,omitempty"`
	Expression         string              `xml:"expression,attr"`
	LocalID            uint64              `xml:"localId,attr"`
	Height             *float64            `xml:"height,attr,omitempty"`
	Width              *float64            `xml:"width,attr,omitempty"`
	EdgeModifier       *EdgeModifierType   `xml:"edgeModifier,attr,omitempty"`
	ExecutionOrderID   *uint64             `xml:"executionOrderId,attr,omitempty"`
}

// BodyFBDOutVariable represents an output variable in FBD
type BodyFBDOutVariable struct {
	Position          *Position            `xml:"position,omitempty"`
	ConnectionPointIn *ConnectionPointIn   `xml:"connectionPointIn,omitempty"`
	Documentation     []byte               `xml:"documentation,omitempty"`
	Expression        string               `xml:"expression,attr"`
	LocalID           uint64               `xml:"localId,attr"`
	Height            *float64             `xml:"height,attr,omitempty"`
	Width             *float64             `xml:"width,attr,omitempty"`
	EdgeModifier      *EdgeModifierType    `xml:"edgeModifier,attr,omitempty"`
	StorageModifier   *StorageModifierType `xml:"storageModifier,attr,omitempty"`
	ExecutionOrderID  *uint64              `xml:"executionOrderId,attr,omitempty"`
}

// BodyFBDInOutVariable represents an in-out variable in FBD
type BodyFBDInOutVariable struct {
	Position           *Position            `xml:"position,omitempty"`
	ConnectionPointIn  *ConnectionPointIn   `xml:"connectionPointIn,omitempty"`
	ConnectionPointOut *ConnectionPointOut  `xml:"connectionPointOut,omitempty"`
	Documentation      []byte               `xml:"documentation,omitempty"`
	Expression         string               `xml:"expression,attr"`
	LocalID            uint64               `xml:"localId,attr"`
	Height             *float64             `xml:"height,attr,omitempty"`
	Width              *float64             `xml:"width,attr,omitempty"`
	EdgeModifier       *EdgeModifierType    `xml:"edgeModifier,attr,omitempty"`
	StorageModifier    *StorageModifierType `xml:"storageModifier,attr,omitempty"`
	ExecutionOrderID   *uint64              `xml:"executionOrderId,attr,omitempty"`
}

// BodyFBDJump represents a jump in FBD
type BodyFBDJump struct {
	Position      *Position `xml:"position,omitempty"`
	Documentation []byte    `xml:"documentation,omitempty"`
	Label         string    `xml:"label,attr"`
	LocalID       uint64    `xml:"localId,attr"`
	Height        *float64  `xml:"height,attr,omitempty"`
	Width         *float64  `xml:"width,attr,omitempty"`
}

// BodyFBDLabel represents a label in FBD
type BodyFBDLabel struct {
	Position      *Position `xml:"position,omitempty"`
	Documentation []byte    `xml:"documentation,omitempty"`
	Label         string    `xml:"label,attr"`
	LocalID       uint64    `xml:"localId,attr"`
	Height        *float64  `xml:"height,attr,omitempty"`
	Width         *float64  `xml:"width,attr,omitempty"`
}

// BodyFBDReturn represents a return in FBD
type BodyFBDReturn struct {
	Position      *Position `xml:"position,omitempty"`
	Documentation []byte    `xml:"documentation,omitempty"`
	LocalID       uint64    `xml:"localId,attr"`
	Height        *float64  `xml:"height,attr,omitempty"`
	Width         *float64  `xml:"width,attr,omitempty"`
}

// BodyLDContact represents a contact in LD
type BodyLDContact struct {
	Position           *Position           `xml:"position,omitempty"`
	ConnectionPointIn  *ConnectionPointIn  `xml:"connectionPointIn,omitempty"`
	ConnectionPointOut *ConnectionPointOut `xml:"connectionPointOut,omitempty"`
	Variable           string              `xml:"variable"`
	Documentation      []byte              `xml:"documentation,omitempty"`
	LocalID            uint64              `xml:"localId,attr"`
	Height             *float64            `xml:"height,attr,omitempty"`
	Width              *float64            `xml:"width,attr,omitempty"`
	EdgeModifier       *EdgeModifierType   `xml:"edgeModifier,attr,omitempty"`
	Negated            *bool               `xml:"negated,attr,omitempty"`
}

// BodyLDCoil represents a coil in LD
type BodyLDCoil struct {
	Position           *Position            `xml:"position,omitempty"`
	ConnectionPointIn  *ConnectionPointIn   `xml:"connectionPointIn,omitempty"`
	ConnectionPointOut *ConnectionPointOut  `xml:"connectionPointOut,omitempty"`
	Variable           string               `xml:"variable"`
	Documentation      []byte               `xml:"documentation,omitempty"`
	LocalID            uint64               `xml:"localId,attr"`
	Height             *float64             `xml:"height,attr,omitempty"`
	Width              *float64             `xml:"width,attr,omitempty"`
	EdgeModifier       *EdgeModifierType    `xml:"edgeModifier,attr,omitempty"`
	StorageModifier    *StorageModifierType `xml:"storageModifier,attr,omitempty"`
	Negated            *bool                `xml:"negated,attr,omitempty"`
}

// BodyLDLeftPowerRail represents a left power rail in LD
type BodyLDLeftPowerRail struct {
	Position           *Position           `xml:"position,omitempty"`
	ConnectionPointOut *ConnectionPointOut `xml:"connectionPointOut,omitempty"`
	Documentation      []byte              `xml:"documentation,omitempty"`
	LocalID            uint64              `xml:"localId,attr"`
	Height             *float64            `xml:"height,attr,omitempty"`
	Width              *float64            `xml:"width,attr,omitempty"`
}

// BodyLDRightPowerRail represents a right power rail in LD
type BodyLDRightPowerRail struct {
	Position          *Position          `xml:"position,omitempty"`
	ConnectionPointIn *ConnectionPointIn `xml:"connectionPointIn,omitempty"`
	Documentation     []byte             `xml:"documentation,omitempty"`
	LocalID           uint64             `xml:"localId,attr"`
	Height            *float64           `xml:"height,attr,omitempty"`
	Width             *float64           `xml:"width,attr,omitempty"`
}

// BodySFCStep represents a step in SFC
type BodySFCStep struct {
	Position                 *Position                            `xml:"position,omitempty"`
	ConnectionPointIn        *BodySFCStepConnectionPointIn        `xml:"connectionPointIn,omitempty"`
	ConnectionPointOut       *BodySFCStepConnectionPointOut       `xml:"connectionPointOut,omitempty"`
	ConnectionPointOutAction *BodySFCStepConnectionPointOutAction `xml:"connectionPointOutAction,omitempty"`
	Documentation            []byte                               `xml:"documentation,omitempty"`
	Name                     string                               `xml:"name,attr"`
	LocalID                  uint64                               `xml:"localId,attr"`
	Height                   *float64                             `xml:"height,attr,omitempty"`
	Width                    *float64                             `xml:"width,attr,omitempty"`
	InitialStep              *bool                                `xml:"initialStep,attr,omitempty"`
}

// BodySFCStepConnectionPointIn represents a step's input connection point
type BodySFCStepConnectionPointIn struct {
	Connections []Connection `xml:"connection,omitempty"`
}

// BodySFCStepConnectionPointOut represents a step's output connection point
type BodySFCStepConnectionPointOut struct {
	FormalParameter *string `xml:"formalParameter,attr,omitempty"`
}

// BodySFCStepConnectionPointOutAction represents a step's action output connection point
type BodySFCStepConnectionPointOutAction struct {
	FormalParameter *string `xml:"formalParameter,attr,omitempty"`
}

// BodySFCTransition represents a transition in SFC
type BodySFCTransition struct {
	Position           *Position                   `xml:"position,omitempty"`
	ConnectionPointIn  *ConnectionPointIn          `xml:"connectionPointIn,omitempty"`
	ConnectionPointOut *ConnectionPointOut         `xml:"connectionPointOut,omitempty"`
	Condition          *BodySFCTransitionCondition `xml:"condition,omitempty"`
	Documentation      []byte                      `xml:"documentation,omitempty"`
	LocalID            uint64                      `xml:"localId,attr"`
	Height             *float64                    `xml:"height,attr,omitempty"`
	Width              *float64                    `xml:"width,attr,omitempty"`
	Priority           *uint64                     `xml:"priority,attr,omitempty"`
}

// BodySFCTransitionCondition represents a transition condition
type BodySFCTransitionCondition struct {
	Inline    *BodySFCTransitionConditionInline    `xml:"inline,omitempty"`
	Reference *BodySFCTransitionConditionReference `xml:"reference,omitempty"`
}

// BodySFCTransitionConditionInline represents an inline transition condition
type BodySFCTransitionConditionInline struct {
	Body *Body  `xml:",innerxml"`
	Name string `xml:"name,attr"`
}

// BodySFCTransitionConditionReference represents a transition condition reference
type BodySFCTransitionConditionReference struct {
	Name string `xml:"name,attr"`
}

// Variable list types for different scopes (type aliases to VarList)
type ProjectTypesPOUInterfaceExternalVars = VarList
type ProjectTypesPOUInterfaceGlobalVars = VarList
type ProjectTypesPOUInterfaceInOutVars = VarList
type ProjectTypesPOUInterfaceInputVars = VarList
type ProjectTypesPOUInterfaceLocalVars = VarList
type ProjectTypesPOUInterfaceOutputVars = VarList
type ProjectTypesPOUInterfaceTempVars = VarList
