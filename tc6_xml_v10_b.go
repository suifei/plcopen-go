// Package plcopen provides Go structs for IEC 61131-3 PLCopen XML format
// generated from XSD schema TC6_XML_V10_B.xsd
package plcopen

import (
	"encoding/xml"
	"time"
)

// Project represents the root element of a PLCopen XML document
type Project struct {
	XMLName       xml.Name              `xml:"http://www.plcopen.org/xml/tc6.xsd project" json:"-"`
	FileHeader    *ProjectFileHeader    `xml:"fileHeader" json:"fileHeader,omitempty"`
	ContentHeader *ProjectContentHeader `xml:"contentHeader" json:"contentHeader,omitempty"`
	Types         *ProjectTypes         `xml:"types" json:"types,omitempty"`
	Instances     *ProjectInstances     `xml:"instances" json:"instances,omitempty"`
}

// ProjectFileHeader contains file metadata
type ProjectFileHeader struct {
	CompanyName        string    `xml:"companyName,attr" json:"companyName"`
	CompanyURL         string    `xml:"companyURL,attr,omitempty" json:"companyURL,omitempty"`
	ProductName        string    `xml:"productName,attr" json:"productName"`
	ProductVersion     string    `xml:"productVersion,attr" json:"productVersion"`
	ProductRelease     string    `xml:"productRelease,attr,omitempty" json:"productRelease,omitempty"`
	CreationDateTime   time.Time `xml:"creationDateTime,attr" json:"creationDateTime"`
	ContentDescription string    `xml:"contentDescription,attr,omitempty" json:"contentDescription,omitempty"`
}

// ProjectContentHeader contains project content metadata
type ProjectContentHeader struct {
	Name                 string                              `xml:"name,attr" json:"name"`
	Version              string                              `xml:"version,attr,omitempty" json:"version,omitempty"`
	ModificationDateTime *time.Time                          `xml:"modificationDateTime,attr,omitempty" json:"modificationDateTime,omitempty"`
	Organization         string                              `xml:"organization,attr,omitempty" json:"organization,omitempty"`
	Author               string                              `xml:"author,attr,omitempty" json:"author,omitempty"`
	Language             string                              `xml:"language,attr,omitempty" json:"language,omitempty"`
	Comment              string                              `xml:"Comment,omitempty" json:"comment,omitempty"`
	CoordinateInfo       *ProjectContentHeaderCoordinateInfo `xml:"coordinateInfo" json:"coordinateInfo,omitempty"`
}

// ProjectContentHeaderCoordinateInfo contains coordinate information
type ProjectContentHeaderCoordinateInfo struct {
	PageSize *ProjectContentHeaderCoordinateInfoPageSize `xml:"pageSize,omitempty" json:"pageSize,omitempty"`
	FBD      *ProjectContentHeaderCoordinateInfoFBD      `xml:"fbd" json:"fBD,omitempty"`
	LD       *ProjectContentHeaderCoordinateInfoLD       `xml:"ld" json:"lD,omitempty"`
	SFC      *ProjectContentHeaderCoordinateInfoSFC      `xml:"sfc" json:"sFC,omitempty"`
}

// ProjectContentHeaderCoordinateInfoPageSize represents page size
type ProjectContentHeaderCoordinateInfoPageSize struct {
	X float64 `xml:"x,attr" json:"x"`
	Y float64 `xml:"y,attr" json:"y"`
}

// ProjectContentHeaderCoordinateInfoFBD represents FBD coordinate info
type ProjectContentHeaderCoordinateInfoFBD struct {
	Scaling *ProjectContentHeaderCoordinateInfoFBDScaling `xml:"scaling" json:"scaling,omitempty"`
}

// ProjectContentHeaderCoordinateInfoFBDScaling represents FBD scaling
type ProjectContentHeaderCoordinateInfoFBDScaling struct {
	X float64 `xml:"x,attr" json:"x"`
	Y float64 `xml:"y,attr" json:"y"`
}

// ProjectContentHeaderCoordinateInfoLD represents LD coordinate info
type ProjectContentHeaderCoordinateInfoLD struct {
	Scaling *ProjectContentHeaderCoordinateInfoLDScaling `xml:"scaling" json:"scaling,omitempty"`
}

// ProjectContentHeaderCoordinateInfoLDScaling represents LD scaling
type ProjectContentHeaderCoordinateInfoLDScaling struct {
	X float64 `xml:"x,attr" json:"x"`
	Y float64 `xml:"y,attr" json:"y"`
}

// ProjectContentHeaderCoordinateInfoSFC represents SFC coordinate info
type ProjectContentHeaderCoordinateInfoSFC struct {
	Scaling *ProjectContentHeaderCoordinateInfoSFCScaling `xml:"scaling" json:"scaling,omitempty"`
}

// ProjectContentHeaderCoordinateInfoSFCScaling represents SFC scaling
type ProjectContentHeaderCoordinateInfoSFCScaling struct {
	X float64 `xml:"x,attr" json:"x"`
	Y float64 `xml:"y,attr" json:"y"`
}

// ProjectTypes contains type definitions
type ProjectTypes struct {
	DataTypes []ProjectTypesDataType `xml:"dataTypes>dataType,omitempty" json:"dataTypes,omitempty"`
	POUs      []ProjectTypesPOU      `xml:"pous>pou,omitempty" json:"pOUs,omitempty"`
}

// ProjectTypesDataType represents a data type definition
type ProjectTypesDataType struct {
	Name          string    `xml:"name,attr" json:"name"`
	BaseType      *DataType `xml:"baseType" json:"baseType,omitempty"`
	InitialValue  *Value    `xml:"initialValue,omitempty" json:"initialValue,omitempty"`
	Documentation []byte    `xml:"documentation,omitempty" json:"documentation,omitempty"`
}

// ProjectTypesPOU represents a Program Organization Unit
type ProjectTypesPOU struct {
	Name          string                      `xml:"name,attr" json:"name"`
	POUType       POUType                     `xml:"pouType,attr" json:"pOUType"`
	Interface     *ProjectTypesPOUInterface   `xml:"interface,omitempty" json:"interface,omitempty"`
	Actions       []ProjectTypesPOUAction     `xml:"actions>action,omitempty" json:"actions,omitempty"`
	Transitions   []ProjectTypesPOUTransition `xml:"transitions>transition,omitempty" json:"transitions,omitempty"`
	Body          *Body                       `xml:"body,omitempty" json:"body,omitempty"`
	Documentation []byte                      `xml:"documentation,omitempty" json:"documentation,omitempty"`
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
	ReturnType    *DataType                             `xml:"returnType,omitempty" json:"returnType,omitempty"`
	LocalVars     *ProjectTypesPOUInterfaceLocalVars    `xml:"localVars,omitempty" json:"localVars,omitempty"`
	InputVars     *ProjectTypesPOUInterfaceInputVars    `xml:"inputVars,omitempty" json:"inputVars,omitempty"`
	InOutVars     *ProjectTypesPOUInterfaceInOutVars    `xml:"inOutVars,omitempty" json:"inOutVars,omitempty"`
	OutputVars    *ProjectTypesPOUInterfaceOutputVars   `xml:"outputVars,omitempty" json:"outputVars,omitempty"`
	ExternalVars  *ProjectTypesPOUInterfaceExternalVars `xml:"externalVars,omitempty" json:"externalVars,omitempty"`
	GlobalVars    *ProjectTypesPOUInterfaceGlobalVars   `xml:"globalVars,omitempty" json:"globalVars,omitempty"`
	TempVars      *ProjectTypesPOUInterfaceTempVars     `xml:"tempVars,omitempty" json:"tempVars,omitempty"`
	AccessVars    *VarList                              `xml:"accessVars,omitempty" json:"accessVars,omitempty"`
	Documentation []byte                                `xml:"documentation,omitempty" json:"documentation,omitempty"`
}

// ProjectTypesPOUAction represents an action within a POU
type ProjectTypesPOUAction struct {
	Name          string `xml:"name,attr" json:"name"`
	Body          *Body  `xml:"body" json:"body,omitempty"`
	Documentation []byte `xml:"documentation,omitempty" json:"documentation,omitempty"`
}

// ProjectTypesPOUTransition represents a transition within a POU
type ProjectTypesPOUTransition struct {
	Name          string `xml:"name,attr" json:"name"`
	Body          *Body  `xml:"body" json:"body,omitempty"`
	Documentation []byte `xml:"documentation,omitempty" json:"documentation,omitempty"`
}

// ProjectInstances contains configuration and resource instances
type ProjectInstances struct {
	Configurations []ProjectInstancesConfiguration `xml:"configurations>configuration,omitempty" json:"configurations,omitempty"`
}

// ProjectInstancesConfiguration represents a configuration
type ProjectInstancesConfiguration struct {
	Name          string                                  `xml:"name,attr" json:"name"`
	Resources     []ProjectInstancesConfigurationResource `xml:"resource,omitempty" json:"resources,omitempty"`
	GlobalVars    *VarList                                `xml:"globalVars,omitempty" json:"globalVars,omitempty"`
	Documentation []byte                                  `xml:"documentation,omitempty" json:"documentation,omitempty"`
}

// ProjectInstancesConfigurationResource represents a resource within a configuration
type ProjectInstancesConfigurationResource struct {
	Name          string                                      `xml:"name,attr" json:"name"`
	Tasks         []ProjectInstancesConfigurationResourceTask `xml:"task,omitempty" json:"tasks,omitempty"`
	GlobalVars    *VarList                                    `xml:"globalVars,omitempty" json:"globalVars,omitempty"`
	POUInstances  []POUInstance                               `xml:"pouInstance,omitempty" json:"pOUInstances,omitempty"`
	Documentation []byte                                      `xml:"documentation,omitempty" json:"documentation,omitempty"`
}

// ProjectInstancesConfigurationResourceTask represents a task within a resource
type ProjectInstancesConfigurationResourceTask struct {
	Name         string        `xml:"name,attr" json:"name"`
	Priority     uint64        `xml:"priority,attr" json:"priority"`
	Interval     *string       `xml:"interval,attr,omitempty" json:"interval,omitempty"`
	Single       *string       `xml:"single,attr,omitempty" json:"single,omitempty"`
	POUInstances []POUInstance `xml:"pouInstance,omitempty" json:"pOUInstances,omitempty"`
}

// POUInstance represents an instance of a POU
type POUInstance struct {
	Name          string `xml:"name,attr" json:"name"`
	TypeName      string `xml:"type,attr" json:"typeName"`
	Documentation []byte `xml:"documentation,omitempty" json:"documentation,omitempty"`
}

// VarList represents a list of variables
type VarList struct {
	Variables     []VarListVariable `xml:"variable,omitempty" json:"variables,omitempty"`
	Documentation []byte            `xml:"documentation,omitempty" json:"documentation,omitempty"`
}

// VarListPlain represents a plain variable list (extends VarList)
type VarListPlain struct {
	Variables     []VarListPlainVariable `xml:"variable,omitempty" json:"variables,omitempty"`
	Documentation []byte                 `xml:"documentation,omitempty" json:"documentation,omitempty"`
}

// VarListPlainVariable represents a plain variable
type VarListPlainVariable struct {
	Name          string    `xml:"name,attr" json:"name"`
	Address       string    `xml:"address,attr,omitempty" json:"address,omitempty"`
	Type          *DataType `xml:"type" json:"type,omitempty"`
	InitialValue  *Value    `xml:"initialValue,omitempty" json:"initialValue,omitempty"`
	Documentation []byte    `xml:"documentation,omitempty" json:"documentation,omitempty"`
}

// VarListVariable represents a variable with additional attributes
type VarListVariable struct {
	Name          string    `xml:"name,attr" json:"name"`
	Address       string    `xml:"address,attr,omitempty" json:"address,omitempty"`
	Type          *DataType `xml:"type" json:"type,omitempty"`
	InitialValue  *Value    `xml:"initialValue,omitempty" json:"initialValue,omitempty"`
	Documentation []byte    `xml:"documentation,omitempty" json:"documentation,omitempty"`
}

// DataType represents a data type with choice-style content
type DataType struct {
	// Basic data types (as empty elements)
	BOOL  *struct{} `xml:"BOOL,omitempty" json:"bOOL,omitempty"`
	BYTE  *struct{} `xml:"BYTE,omitempty" json:"bYTE,omitempty"`
	DATE  *struct{} `xml:"DATE,omitempty" json:"dATE,omitempty"`
	DINT  *struct{} `xml:"DINT,omitempty" json:"dINT,omitempty"`
	DT    *struct{} `xml:"DT,omitempty" json:"dT,omitempty"`
	DWORD *struct{} `xml:"DWORD,omitempty" json:"dWORD,omitempty"`
	INT   *struct{} `xml:"INT,omitempty" json:"iNT,omitempty"`
	LINT  *struct{} `xml:"LINT,omitempty" json:"lINT,omitempty"`
	LREAL *struct{} `xml:"LREAL,omitempty" json:"lREAL,omitempty"`
	LWORD *struct{} `xml:"LWORD,omitempty" json:"lWORD,omitempty"`
	REAL  *struct{} `xml:"REAL,omitempty" json:"rEAL,omitempty"`
	SINT  *struct{} `xml:"SINT,omitempty" json:"sINT,omitempty"`
	TIME  *struct{} `xml:"TIME,omitempty" json:"tIME,omitempty"`
	TOD   *struct{} `xml:"TOD,omitempty" json:"tOD,omitempty"`
	UDINT *struct{} `xml:"UDINT,omitempty" json:"uDINT,omitempty"`
	UINT  *struct{} `xml:"UINT,omitempty" json:"uINT,omitempty"`
	ULINT *struct{} `xml:"ULINT,omitempty" json:"uLINT,omitempty"`
	USINT *struct{} `xml:"USINT,omitempty" json:"uSINT,omitempty"`
	WORD  *struct{} `xml:"WORD,omitempty" json:"wORD,omitempty"`

	// Complex data types
	Array            *DataTypeArray            `xml:"array,omitempty" json:"array,omitempty"`
	Derived          *DataTypeDerived          `xml:"derived,omitempty" json:"derived,omitempty"`
	Enum             *DataTypeEnum             `xml:"enum,omitempty" json:"enum,omitempty"`
	Pointer          *DataTypePointer          `xml:"pointer,omitempty" json:"pointer,omitempty"`
	String           *DataTypeString           `xml:"string,omitempty" json:"string,omitempty"`
	Struct           *VarListPlain             `xml:"struct,omitempty" json:"struct,omitempty"`
	SubrangeSigned   *DataTypeSubrangeSigned   `xml:"subrangeSigned,omitempty" json:"subrangeSigned,omitempty"`
	SubrangeUnsigned *DataTypeSubrangeUnsigned `xml:"subrangeUnsigned,omitempty" json:"subrangeUnsigned,omitempty"`
	WString          *DataTypeWString          `xml:"wstring,omitempty" json:"wString,omitempty"`
}

// DataTypeArray represents an array data type
type DataTypeArray struct {
	Dimensions []RangeSigned `xml:"dimension" json:"dimensions,omitempty"`
	BaseType   *DataType     `xml:"baseType" json:"baseType,omitempty"`
}

// RangeSigned represents a signed range
type RangeSigned struct {
	Lower int64 `xml:"lower,attr" json:"lower"`
	Upper int64 `xml:"upper,attr" json:"upper"`
}

// RangeUnsigned represents an unsigned range
type RangeUnsigned struct {
	Lower uint64 `xml:"lower,attr" json:"lower"`
	Upper uint64 `xml:"upper,attr" json:"upper"`
}

// DataTypeDerived represents a derived data type
type DataTypeDerived struct {
	Name string `xml:"name,attr" json:"name"`
}

// DataTypeEnum represents an enumerated data type
type DataTypeEnum struct {
	Values   *DataTypeEnumValues `xml:"values" json:"values,omitempty"`
	BaseType *DataType           `xml:"baseType,omitempty" json:"baseType,omitempty"`
}

// DataTypeEnumValues represents enumerated values
type DataTypeEnumValues struct {
	Values []DataTypeEnumValuesValue `xml:"value" json:"values,omitempty"`
}

// DataTypeEnumValuesValue represents an enumerated value
type DataTypeEnumValuesValue struct {
	Name          string `xml:"name,attr" json:"name"`
	Documentation []byte `xml:"documentation,omitempty" json:"documentation,omitempty"`
}

// DataTypePointer represents a pointer data type
type DataTypePointer struct {
	BaseType *DataType `xml:"baseType" json:"baseType,omitempty"`
}

// DataTypeString represents a string data type
type DataTypeString struct {
	Length *uint64 `xml:"length,attr,omitempty" json:"length,omitempty"`
}

// DataTypeWString represents a wide string data type
type DataTypeWString struct {
	Length *uint64 `xml:"length,attr,omitempty" json:"length,omitempty"`
}

// DataTypeSubrangeSigned represents a signed subrange data type
type DataTypeSubrangeSigned struct {
	Range    *RangeSigned `xml:"range" json:"range,omitempty"`
	BaseType *DataType    `xml:"baseType" json:"baseType,omitempty"`
}

// DataTypeSubrangeUnsigned represents an unsigned subrange data type
type DataTypeSubrangeUnsigned struct {
	Range    *RangeUnsigned `xml:"range" json:"range,omitempty"`
	BaseType *DataType      `xml:"baseType" json:"baseType,omitempty"`
}

// Value represents a value with choice-style content
type Value struct {
	SimpleValue *ValueSimpleValue `xml:"simpleValue,omitempty" json:"simpleValue,omitempty"`
	ArrayValue  *ValueArrayValue  `xml:"arrayValue,omitempty" json:"arrayValue,omitempty"`
	StructValue *ValueStructValue `xml:"structValue,omitempty" json:"structValue,omitempty"`
}

// ValueArrayValue represents an array value
type ValueArrayValue struct {
	Values []ValueArrayValueValue `xml:"value" json:"values,omitempty"`
}

// ValueArrayValueValue represents a value within an array
type ValueArrayValueValue struct {
	RepeatCount *uint64 `xml:"repetition,attr,omitempty" json:"repeatCount,omitempty"`
	Value       *Value  `xml:",innerxml" json:"value,omitempty"`
}

// ValueSimpleValue represents a simple value
type ValueSimpleValue struct {
	Value string `xml:"value,attr" json:"value"`
}

// ValueStructValue represents a structured value
type ValueStructValue struct {
	Values []ValueStructValueValue `xml:"value" json:"values,omitempty"`
}

// ValueStructValueValue represents a value within a structure
type ValueStructValueValue struct {
	Member string `xml:"member,attr" json:"member"`
	Value  *Value `xml:",innerxml" json:"value,omitempty"`
}

// Body represents a body with choice-style content for different languages
type Body struct {
	FBD *BodyFBD `xml:"FBD,omitempty" json:"fBD,omitempty"`
	LD  *BodyLD  `xml:"LD,omitempty" json:"lD,omitempty"`
	SFC *BodySFC `xml:"SFC,omitempty" json:"sFC,omitempty"`
	IL  *BodyIL  `xml:"IL,omitempty" json:"iL,omitempty"`
	ST  *BodyST  `xml:"ST,omitempty" json:"sT,omitempty"`
}

// BodyFBD represents a Function Block Diagram body
type BodyFBD struct {
	Blocks         []BodyFBDBlock         `xml:"block,omitempty" json:"blocks,omitempty"`
	ActionBlocks   []BodyFBDActionBlock   `xml:"actionBlock,omitempty" json:"actionBlocks,omitempty"`
	Comments       []BodyFBDComment       `xml:"comment,omitempty" json:"comments,omitempty"`
	Connectors     []BodyFBDConnector     `xml:"connector,omitempty" json:"connectors,omitempty"`
	Continuations  []BodyFBDContinuation  `xml:"continuation,omitempty" json:"continuations,omitempty"`
	InVariables    []BodyFBDInVariable    `xml:"inVariable,omitempty" json:"inVariables,omitempty"`
	OutVariables   []BodyFBDOutVariable   `xml:"outVariable,omitempty" json:"outVariables,omitempty"`
	InOutVariables []BodyFBDInOutVariable `xml:"inOutVariable,omitempty" json:"inOutVariables,omitempty"`
	Jumps          []BodyFBDJump          `xml:"jump,omitempty" json:"jumps,omitempty"`
	Labels         []BodyFBDLabel         `xml:"label,omitempty" json:"labels,omitempty"`
	Returns        []BodyFBDReturn        `xml:"return,omitempty" json:"returns,omitempty"`
}

// BodyLD represents a Ladder Diagram body
type BodyLD struct {
	Contacts        []BodyLDContact        `xml:"contact,omitempty" json:"contacts,omitempty"`
	Coils           []BodyLDCoil           `xml:"coil,omitempty" json:"coils,omitempty"`
	LeftPowerRails  []BodyLDLeftPowerRail  `xml:"leftPowerRail,omitempty" json:"leftPowerRails,omitempty"`
	RightPowerRails []BodyLDRightPowerRail `xml:"rightPowerRail,omitempty" json:"rightPowerRails,omitempty"`
}

// BodySFC represents a Sequential Function Chart body
type BodySFC struct {
	Steps       []BodySFCStep       `xml:"step,omitempty" json:"steps,omitempty"`
	Transitions []BodySFCTransition `xml:"transition,omitempty" json:"transitions,omitempty"`
}

// BodyIL represents an Instruction List body
type BodyIL struct {
	XMLNSXhtml string `xml:"xmlns:xhtml,attr" json:"xMLNSXhtml"`
	Xhtml      string `xml:",innerxml" json:"xhtml"`
}

// BodyST represents a Structured Text body
type BodyST struct {
	XMLNSXhtml string `xml:"xmlns:xhtml,attr" json:"xMLNSXhtml"`
	Xhtml      string `xml:",innerxml" json:"xhtml"`
}

// Position represents a position coordinate
type Position struct {
	X float64 `xml:"x,attr" json:"x"`
	Y float64 `xml:"y,attr" json:"y"`
}

// Connection represents a connection
type Connection struct {
	RefLocalID      uint64  `xml:"refLocalId,attr" json:"refLocalID"`
	FormalParameter *string `xml:"formalParameter,attr,omitempty" json:"formalParameter,omitempty"`
}

// ConnectionPointIn represents an input connection point
type ConnectionPointIn struct {
	Connections []Connection `xml:"connection,omitempty" json:"connections,omitempty"`
}

// ConnectionPointOut represents an output connection point
type ConnectionPointOut struct {
	FormalParameter *string `xml:"formalParameter,attr,omitempty" json:"formalParameter,omitempty"`
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
	Position         *Position               `xml:"position,omitempty" json:"position,omitempty"`
	InputVariables   []BodyFBDBlockVariable  `xml:"inputVariables>variable,omitempty" json:"inputVariables,omitempty"`
	InOutVariables   []BodyFBDBlockVariable2 `xml:"inOutVariables>variable,omitempty" json:"inOutVariables,omitempty"`
	OutputVariables  []BodyFBDBlockVariable1 `xml:"outputVariables>variable,omitempty" json:"outputVariables,omitempty"`
	Documentation    []byte                  `xml:"documentation,omitempty" json:"documentation,omitempty"`
	LocalID          uint64                  `xml:"localId,attr" json:"localID"`
	Width            *float64                `xml:"width,attr,omitempty" json:"width,omitempty"`
	Height           *float64                `xml:"height,attr,omitempty" json:"height,omitempty"`
	TypeName         string                  `xml:"typeName,attr" json:"typeName"`
	InstanceName     *string                 `xml:"instanceName,attr,omitempty" json:"instanceName,omitempty"`
	ExecutionOrderID *uint64                 `xml:"executionOrderId,attr,omitempty" json:"executionOrderID,omitempty"`
}

// BodyFBDBlockVariable represents a variable in a block
type BodyFBDBlockVariable struct {
	FormalParameter   string             `xml:"formalParameter,attr" json:"formalParameter"`
	ConnectionPointIn *ConnectionPointIn `xml:"connectionPointIn,omitempty" json:"connectionPointIn,omitempty"`
}

// BodyFBDBlockVariable1 represents an output variable in a block
type BodyFBDBlockVariable1 struct {
	FormalParameter    string              `xml:"formalParameter,attr" json:"formalParameter"`
	ConnectionPointOut *ConnectionPointOut `xml:"connectionPointOut,omitempty" json:"connectionPointOut,omitempty"`
}

// BodyFBDBlockVariable2 represents an in-out variable in a block
type BodyFBDBlockVariable2 struct {
	FormalParameter    string              `xml:"formalParameter,attr" json:"formalParameter"`
	ConnectionPointIn  *ConnectionPointIn  `xml:"connectionPointIn,omitempty" json:"connectionPointIn,omitempty"`
	ConnectionPointOut *ConnectionPointOut `xml:"connectionPointOut,omitempty" json:"connectionPointOut,omitempty"`
}

// BodyFBDActionBlock represents an action block in FBD
type BodyFBDActionBlock struct {
	Position      *Position                  `xml:"position,omitempty" json:"position,omitempty"`
	Actions       []BodyFBDActionBlockAction `xml:"action,omitempty" json:"actions,omitempty"`
	Documentation []byte                     `xml:"documentation,omitempty" json:"documentation,omitempty"`
	LocalID       uint64                     `xml:"localId,attr" json:"localID"`
	Width         *float64                   `xml:"width,attr,omitempty" json:"width,omitempty"`
	Height        *float64                   `xml:"height,attr,omitempty" json:"height,omitempty"`
}

// BodyFBDActionBlockAction represents an action in an action block
type BodyFBDActionBlockAction struct {
	Reference *BodyFBDActionBlockActionReference `xml:"reference,omitempty" json:"reference,omitempty"`
	Inline    *BodyFBDActionBlockActionInline    `xml:"inline,omitempty" json:"inline,omitempty"`
	Qualifier *BodyFBDActionBlockActionQualifier `xml:"qualifier,attr,omitempty" json:"qualifier,omitempty"`
}

// BodyFBDActionBlockActionReference represents an action reference
type BodyFBDActionBlockActionReference struct {
	Name string `xml:"name,attr" json:"name"`
}

// BodyFBDActionBlockActionInline represents an inline action
type BodyFBDActionBlockActionInline struct {
	Body *Body  `xml:",innerxml" json:"body,omitempty"`
	Name string `xml:"name,attr" json:"name"`
}

// BodyFBDComment represents a comment in FBD
type BodyFBDComment struct {
	Position      *Position `xml:"position,omitempty" json:"position,omitempty"`
	Content       string    `xml:"content,omitempty" json:"content,omitempty"`
	Documentation []byte    `xml:"documentation,omitempty" json:"documentation,omitempty"`
	LocalID       uint64    `xml:"localId,attr" json:"localID"`
	Height        *float64  `xml:"height,attr,omitempty" json:"height,omitempty"`
	Width         *float64  `xml:"width,attr,omitempty" json:"width,omitempty"`
}

// BodyFBDConnector represents a connector in FBD
type BodyFBDConnector struct {
	Position           *Position           `xml:"position,omitempty" json:"position,omitempty"`
	ConnectionPointOut *ConnectionPointOut `xml:"connectionPointOut,omitempty" json:"connectionPointOut,omitempty"`
	Documentation      []byte              `xml:"documentation,omitempty" json:"documentation,omitempty"`
	Name               string              `xml:"name,attr" json:"name"`
	LocalID            uint64              `xml:"localId,attr" json:"localID"`
}

// BodyFBDContinuation represents a continuation in FBD
type BodyFBDContinuation struct {
	Position          *Position          `xml:"position,omitempty" json:"position,omitempty"`
	ConnectionPointIn *ConnectionPointIn `xml:"connectionPointIn,omitempty" json:"connectionPointIn,omitempty"`
	Documentation     []byte             `xml:"documentation,omitempty" json:"documentation,omitempty"`
	Name              string             `xml:"name,attr" json:"name"`
	LocalID           uint64             `xml:"localId,attr" json:"localID"`
}

// BodyFBDInVariable represents an input variable in FBD
type BodyFBDInVariable struct {
	Position           *Position           `xml:"position,omitempty" json:"position,omitempty"`
	ConnectionPointOut *ConnectionPointOut `xml:"connectionPointOut,omitempty" json:"connectionPointOut,omitempty"`
	Documentation      []byte              `xml:"documentation,omitempty" json:"documentation,omitempty"`
	Expression         string              `xml:"expression,attr" json:"expression"`
	LocalID            uint64              `xml:"localId,attr" json:"localID"`
	Height             *float64            `xml:"height,attr,omitempty" json:"height,omitempty"`
	Width              *float64            `xml:"width,attr,omitempty" json:"width,omitempty"`
	EdgeModifier       *EdgeModifierType   `xml:"edgeModifier,attr,omitempty" json:"edgeModifier,omitempty"`
	ExecutionOrderID   *uint64             `xml:"executionOrderId,attr,omitempty" json:"executionOrderID,omitempty"`
}

// BodyFBDOutVariable represents an output variable in FBD
type BodyFBDOutVariable struct {
	Position          *Position            `xml:"position,omitempty" json:"position,omitempty"`
	ConnectionPointIn *ConnectionPointIn   `xml:"connectionPointIn,omitempty" json:"connectionPointIn,omitempty"`
	Documentation     []byte               `xml:"documentation,omitempty" json:"documentation,omitempty"`
	Expression        string               `xml:"expression,attr" json:"expression"`
	LocalID           uint64               `xml:"localId,attr" json:"localID"`
	Height            *float64             `xml:"height,attr,omitempty" json:"height,omitempty"`
	Width             *float64             `xml:"width,attr,omitempty" json:"width,omitempty"`
	EdgeModifier      *EdgeModifierType    `xml:"edgeModifier,attr,omitempty" json:"edgeModifier,omitempty"`
	StorageModifier   *StorageModifierType `xml:"storageModifier,attr,omitempty" json:"storageModifier,omitempty"`
	ExecutionOrderID  *uint64              `xml:"executionOrderId,attr,omitempty" json:"executionOrderID,omitempty"`
}

// BodyFBDInOutVariable represents an in-out variable in FBD
type BodyFBDInOutVariable struct {
	Position           *Position            `xml:"position,omitempty" json:"position,omitempty"`
	ConnectionPointIn  *ConnectionPointIn   `xml:"connectionPointIn,omitempty" json:"connectionPointIn,omitempty"`
	ConnectionPointOut *ConnectionPointOut  `xml:"connectionPointOut,omitempty" json:"connectionPointOut,omitempty"`
	Documentation      []byte               `xml:"documentation,omitempty" json:"documentation,omitempty"`
	Expression         string               `xml:"expression,attr" json:"expression"`
	LocalID            uint64               `xml:"localId,attr" json:"localID"`
	Height             *float64             `xml:"height,attr,omitempty" json:"height,omitempty"`
	Width              *float64             `xml:"width,attr,omitempty" json:"width,omitempty"`
	EdgeModifier       *EdgeModifierType    `xml:"edgeModifier,attr,omitempty" json:"edgeModifier,omitempty"`
	StorageModifier    *StorageModifierType `xml:"storageModifier,attr,omitempty" json:"storageModifier,omitempty"`
	ExecutionOrderID   *uint64              `xml:"executionOrderId,attr,omitempty" json:"executionOrderID,omitempty"`
}

// BodyFBDJump represents a jump in FBD
type BodyFBDJump struct {
	Position      *Position `xml:"position,omitempty" json:"position,omitempty"`
	Documentation []byte    `xml:"documentation,omitempty" json:"documentation,omitempty"`
	Label         string    `xml:"label,attr" json:"label"`
	LocalID       uint64    `xml:"localId,attr" json:"localID"`
	Height        *float64  `xml:"height,attr,omitempty" json:"height,omitempty"`
	Width         *float64  `xml:"width,attr,omitempty" json:"width,omitempty"`
}

// BodyFBDLabel represents a label in FBD
type BodyFBDLabel struct {
	Position      *Position `xml:"position,omitempty" json:"position,omitempty"`
	Documentation []byte    `xml:"documentation,omitempty" json:"documentation,omitempty"`
	Label         string    `xml:"label,attr" json:"label"`
	LocalID       uint64    `xml:"localId,attr" json:"localID"`
	Height        *float64  `xml:"height,attr,omitempty" json:"height,omitempty"`
	Width         *float64  `xml:"width,attr,omitempty" json:"width,omitempty"`
}

// BodyFBDReturn represents a return in FBD
type BodyFBDReturn struct {
	Position      *Position `xml:"position,omitempty" json:"position,omitempty"`
	Documentation []byte    `xml:"documentation,omitempty" json:"documentation,omitempty"`
	LocalID       uint64    `xml:"localId,attr" json:"localID"`
	Height        *float64  `xml:"height,attr,omitempty" json:"height,omitempty"`
	Width         *float64  `xml:"width,attr,omitempty" json:"width,omitempty"`
}

// BodyLDContact represents a contact in LD
type BodyLDContact struct {
	Position           *Position           `xml:"position,omitempty" json:"position,omitempty"`
	ConnectionPointIn  *ConnectionPointIn  `xml:"connectionPointIn,omitempty" json:"connectionPointIn,omitempty"`
	ConnectionPointOut *ConnectionPointOut `xml:"connectionPointOut,omitempty" json:"connectionPointOut,omitempty"`
	Variable           string              `xml:"variable" json:"variable"`
	Documentation      []byte              `xml:"documentation,omitempty" json:"documentation,omitempty"`
	LocalID            uint64              `xml:"localId,attr" json:"localID"`
	Height             *float64            `xml:"height,attr,omitempty" json:"height,omitempty"`
	Width              *float64            `xml:"width,attr,omitempty" json:"width,omitempty"`
	EdgeModifier       *EdgeModifierType   `xml:"edgeModifier,attr,omitempty" json:"edgeModifier,omitempty"`
	Negated            *bool               `xml:"negated,attr,omitempty" json:"negated,omitempty"`
}

// BodyLDCoil represents a coil in LD
type BodyLDCoil struct {
	Position           *Position            `xml:"position,omitempty" json:"position,omitempty"`
	ConnectionPointIn  *ConnectionPointIn   `xml:"connectionPointIn,omitempty" json:"connectionPointIn,omitempty"`
	ConnectionPointOut *ConnectionPointOut  `xml:"connectionPointOut,omitempty" json:"connectionPointOut,omitempty"`
	Variable           string               `xml:"variable" json:"variable"`
	Documentation      []byte               `xml:"documentation,omitempty" json:"documentation,omitempty"`
	LocalID            uint64               `xml:"localId,attr" json:"localID"`
	Height             *float64             `xml:"height,attr,omitempty" json:"height,omitempty"`
	Width              *float64             `xml:"width,attr,omitempty" json:"width,omitempty"`
	EdgeModifier       *EdgeModifierType    `xml:"edgeModifier,attr,omitempty" json:"edgeModifier,omitempty"`
	StorageModifier    *StorageModifierType `xml:"storageModifier,attr,omitempty" json:"storageModifier,omitempty"`
	Negated            *bool                `xml:"negated,attr,omitempty" json:"negated,omitempty"`
}

// BodyLDLeftPowerRail represents a left power rail in LD
type BodyLDLeftPowerRail struct {
	Position           *Position           `xml:"position,omitempty" json:"position,omitempty"`
	ConnectionPointOut *ConnectionPointOut `xml:"connectionPointOut,omitempty" json:"connectionPointOut,omitempty"`
	Documentation      []byte              `xml:"documentation,omitempty" json:"documentation,omitempty"`
	LocalID            uint64              `xml:"localId,attr" json:"localID"`
	Height             *float64            `xml:"height,attr,omitempty" json:"height,omitempty"`
	Width              *float64            `xml:"width,attr,omitempty" json:"width,omitempty"`
}

// BodyLDRightPowerRail represents a right power rail in LD
type BodyLDRightPowerRail struct {
	Position          *Position          `xml:"position,omitempty" json:"position,omitempty"`
	ConnectionPointIn *ConnectionPointIn `xml:"connectionPointIn,omitempty" json:"connectionPointIn,omitempty"`
	Documentation     []byte             `xml:"documentation,omitempty" json:"documentation,omitempty"`
	LocalID           uint64             `xml:"localId,attr" json:"localID"`
	Height            *float64           `xml:"height,attr,omitempty" json:"height,omitempty"`
	Width             *float64           `xml:"width,attr,omitempty" json:"width,omitempty"`
}

// BodySFCStep represents a step in SFC
type BodySFCStep struct {
	Position                 *Position                            `xml:"position,omitempty" json:"position,omitempty"`
	ConnectionPointIn        *BodySFCStepConnectionPointIn        `xml:"connectionPointIn,omitempty" json:"connectionPointIn,omitempty"`
	ConnectionPointOut       *BodySFCStepConnectionPointOut       `xml:"connectionPointOut,omitempty" json:"connectionPointOut,omitempty"`
	ConnectionPointOutAction *BodySFCStepConnectionPointOutAction `xml:"connectionPointOutAction,omitempty" json:"connectionPointOutAction,omitempty"`
	Documentation            []byte                               `xml:"documentation,omitempty" json:"documentation,omitempty"`
	Name                     string                               `xml:"name,attr" json:"name"`
	LocalID                  uint64                               `xml:"localId,attr" json:"localID"`
	Height                   *float64                             `xml:"height,attr,omitempty" json:"height,omitempty"`
	Width                    *float64                             `xml:"width,attr,omitempty" json:"width,omitempty"`
	InitialStep              *bool                                `xml:"initialStep,attr,omitempty" json:"initialStep,omitempty"`
}

// BodySFCStepConnectionPointIn represents a step's input connection point
type BodySFCStepConnectionPointIn struct {
	Connections []Connection `xml:"connection,omitempty" json:"connections,omitempty"`
}

// BodySFCStepConnectionPointOut represents a step's output connection point
type BodySFCStepConnectionPointOut struct {
	FormalParameter *string `xml:"formalParameter,attr,omitempty" json:"formalParameter,omitempty"`
}

// BodySFCStepConnectionPointOutAction represents a step's action output connection point
type BodySFCStepConnectionPointOutAction struct {
	FormalParameter *string `xml:"formalParameter,attr,omitempty" json:"formalParameter,omitempty"`
}

// BodySFCTransition represents a transition in SFC
type BodySFCTransition struct {
	Position           *Position                   `xml:"position,omitempty" json:"position,omitempty"`
	ConnectionPointIn  *ConnectionPointIn          `xml:"connectionPointIn,omitempty" json:"connectionPointIn,omitempty"`
	ConnectionPointOut *ConnectionPointOut         `xml:"connectionPointOut,omitempty" json:"connectionPointOut,omitempty"`
	Condition          *BodySFCTransitionCondition `xml:"condition,omitempty" json:"condition,omitempty"`
	Documentation      []byte                      `xml:"documentation,omitempty" json:"documentation,omitempty"`
	LocalID            uint64                      `xml:"localId,attr" json:"localID"`
	Height             *float64                    `xml:"height,attr,omitempty" json:"height,omitempty"`
	Width              *float64                    `xml:"width,attr,omitempty" json:"width,omitempty"`
	Priority           *uint64                     `xml:"priority,attr,omitempty" json:"priority,omitempty"`
}

// BodySFCTransitionCondition represents a transition condition
type BodySFCTransitionCondition struct {
	Inline    *BodySFCTransitionConditionInline    `xml:"inline,omitempty" json:"inline,omitempty"`
	Reference *BodySFCTransitionConditionReference `xml:"reference,omitempty" json:"reference,omitempty"`
}

// BodySFCTransitionConditionInline represents an inline transition condition
type BodySFCTransitionConditionInline struct {
	Body *Body  `xml:",innerxml" json:"body,omitempty"`
	Name string `xml:"name,attr" json:"name"`
}

// BodySFCTransitionConditionReference represents a transition condition reference
type BodySFCTransitionConditionReference struct {
	Name string `xml:"name,attr" json:"name"`
}

// Variable list types for different scopes (type aliases to VarList)
type ProjectTypesPOUInterfaceExternalVars = VarList
type ProjectTypesPOUInterfaceGlobalVars = VarList
type ProjectTypesPOUInterfaceInOutVars = VarList
type ProjectTypesPOUInterfaceInputVars = VarList
type ProjectTypesPOUInterfaceLocalVars = VarList
type ProjectTypesPOUInterfaceOutputVars = VarList
type ProjectTypesPOUInterfaceTempVars = VarList
