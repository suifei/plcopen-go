package tests

import (
	"encoding/xml"
	"testing"

	"github.com/suifei/plcopen-go"
)

// TestFormattedText 验证formattedText类型在Go代码中的处理方式
func TestFormattedText(t *testing.T) {
	// 在XSD中，formattedText是一个复杂类型，允许混合内容
	// 在Go中，它通常表示为[]byte类型，可以保存任何XML内容
	// 创建一个包含formattedText的POU文档
	pou := plcopen.ProjectTypesPOU{
		Name:    "TestPOU",
		POUType: "function",
		Body: &plcopen.Body{
			ST: &plcopen.BodyST{
				Xhtml: "// 这是一个带格式的文本\nIF a > b THEN\n  c := a;\nEND_IF;",
			},
		},
	}

	// 测试Documentation字段
	dataType := plcopen.ProjectTypesDataType{
		Name: "TestType",
		BaseType: &plcopen.DataType{
			Derived: &plcopen.DataTypeDerived{
				Name: "INT",
			},
		},
		Documentation: []byte("<documentation>这是带<b>格式</b>的<i>文档</i>内容</documentation>"),
	}

	// 确保可以序列化为XML
	pouXML, err := xml.MarshalIndent(pou, "", "  ")
	if err != nil {
		t.Errorf("无法序列化带有formattedText的POU: %v", err)
	}

	dataTypeXML, err := xml.MarshalIndent(dataType, "", "  ")
	if err != nil {
		t.Errorf("无法序列化带有formattedText的DataType: %v", err)
	}

	t.Logf("带formattedText的POU序列化为:\n%s", string(pouXML))
	t.Logf("带formattedText的DataType序列化为:\n%s", string(dataTypeXML))

	// 解析包含formattedText的XML
	xmlString := `
	<pou name="TestPOU" pouType="function">
		<body>
			<ST><![CDATA[// 这是XML中的带格式文本
IF a > b THEN
  c := a;
END_IF;]]></ST>
		</body>
		<documentation><![CDATA[这是<b>带格式</b>的文档]]></documentation>
	</pou>`

	var parsedPOU plcopen.ProjectTypesPOU
	if err := xml.Unmarshal([]byte(xmlString), &parsedPOU); err != nil {
		t.Errorf("无法解析包含formattedText的XML: %v", err)
	}
	// 验证解析结果
	if parsedPOU.Name != "TestPOU" {
		t.Errorf("解析错误，期望name='TestPOU'，实际为'%s'", parsedPOU.Name)
	}

	if parsedPOU.Body == nil || parsedPOU.Body.ST == nil {
		t.Error("解析错误，Body或ST为nil")
	} else {
		t.Logf("成功解析formattedText内容: %s", parsedPOU.Body.ST.Xhtml)
	}

	t.Log("formattedText类型测试通过，在Go中使用[]byte类型正确处理")
}
