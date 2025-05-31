package utils

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

// XSD类型信息
type TypeInfo struct {
	Name       string
	Path       string
	IsNamed    bool
	Elements   []ElementInfo
	Attributes []AttributeInfo
}

// XSD元素信息
type ElementInfo struct {
	Name     string
	Type     string
	Optional bool
}

// XSD属性信息
type AttributeInfo struct {
	Name     string
	Type     string
	Optional bool
}

// XSD解析器
type XSDParser struct {
	NamedTypes     map[string]TypeInfo
	AnonymousTypes []TypeInfo
	Count          int
	CurrentPath    []string
}

// 解析XSD文件
func (p *XSDParser) ParseXSD(filename string) error {
	// 初始化数据结构
	p.NamedTypes = make(map[string]TypeInfo)
	p.AnonymousTypes = []TypeInfo{}
	p.Count = 0
	p.CurrentPath = []string{}

	// 读取XSD文件
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// 创建XML解析器
	decoder := xml.NewDecoder(strings.NewReader(string(data)))

	// 处理XML tokens
	err = p.processXML(decoder)
	if err != nil {
		return err
	}

	return nil
}

// 处理XML内容
func (p *XSDParser) processXML(decoder *xml.Decoder) error {
	var currentType *TypeInfo
	// var currentElement *ElementInfo  // 移除未使用的变量

	for {
		token, err := decoder.Token()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}
		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			// 更新路径
			p.CurrentPath = append(p.CurrentPath, se.Name.Local)

			if se.Name.Local == "complexType" {
				p.Count++

				// 创建新的类型信息
				typeInfo := TypeInfo{
					Path:       strings.Join(p.CurrentPath, "/"),
					Elements:   []ElementInfo{},
					Attributes: []AttributeInfo{},
				}

				// 尝试获取name属性
				var name string
				for _, attr := range se.Attr {
					if attr.Name.Local == "name" {
						name = attr.Value
						break
					}
				}

				// 如果有名字，添加到命名类型映射中
				if name != "" {
					typeInfo.Name = name
					typeInfo.IsNamed = true
					p.NamedTypes[name] = typeInfo
					fmt.Printf("找到命名复杂类型: %s，路径：%s\n", name, typeInfo.Path)
				} else {
					typeInfo.IsNamed = false
					parentPath := ""
					if len(p.CurrentPath) >= 2 {
						parentPath = p.CurrentPath[len(p.CurrentPath)-2]
					}
					typeInfo.Name = fmt.Sprintf("anonymous-%s-%d", parentPath, len(p.AnonymousTypes))
					p.AnonymousTypes = append(p.AnonymousTypes, typeInfo)
					//fmt.Printf("找到匿名复杂类型，路径：%s\n", typeInfo.Path)
				}

				currentType = &typeInfo
			} else if se.Name.Local == "element" {
				// 处理元素
				elementInfo := ElementInfo{}
				for _, attr := range se.Attr {
					if attr.Name.Local == "name" {
						elementInfo.Name = attr.Value
					} else if attr.Name.Local == "type" {
						elementInfo.Type = attr.Value
					} else if attr.Name.Local == "minOccurs" {
						if attr.Value == "0" {
							elementInfo.Optional = true
						}
					}
				}

				if elementInfo.Name != "" && elementInfo.Type != "" && currentType != nil {
					// 如果是在complexType内部，添加到当前类型的元素列表
					if strings.Contains(p.CurrentPath[len(p.CurrentPath)-2], "complexType") {
						currentType.Elements = append(currentType.Elements, elementInfo)
						//fmt.Printf("添加元素 '%s' 类型 '%s' 到 '%s'\n", elementInfo.Name, elementInfo.Type, currentType.Name)
					}
				} // currentElement = &elementInfo  // 移除未使用的变量赋值
			} else if se.Name.Local == "attribute" {
				// 处理属性
				attributeInfo := AttributeInfo{}
				for _, attr := range se.Attr {
					if attr.Name.Local == "name" {
						attributeInfo.Name = attr.Value
					} else if attr.Name.Local == "type" {
						attributeInfo.Type = attr.Value
					} else if attr.Name.Local == "use" && attr.Value == "optional" {
						attributeInfo.Optional = true
					}
				}

				if attributeInfo.Name != "" && currentType != nil {
					currentType.Attributes = append(currentType.Attributes, attributeInfo)
					//fmt.Printf("添加属性 '%s' 类型 '%s' 到 '%s'\n", attributeInfo.Name, attributeInfo.Type, currentType.Name)
				}
			}

		case xml.EndElement:
			// 到达结束标签时，从路径中移除最后一个元素
			if len(p.CurrentPath) > 0 {
				p.CurrentPath = p.CurrentPath[:len(p.CurrentPath)-1]
			}

			// 如果结束的是complexType，清除当前类型
			if se.Name.Local == "complexType" {
				currentType = nil // } else if se.Name.Local == "element" {
				// 	currentElement = nil
			}
		}
	}

	return nil
}

// Go类型信息
type GoTypeInfo struct {
	Name   string
	Fields []string
}

// 从Go源文件中提取类型信息
func ExtractGoTypes(filename string) (map[string]GoTypeInfo, error) {
	// 读取Go源文件
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("读取Go文件失败: %v", err)
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	types := make(map[string]GoTypeInfo)

	// 简单解析，查找type定义
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		// 检测类型定义
		if strings.HasPrefix(line, "type ") {
			parts := strings.Fields(line)
			if len(parts) >= 3 {
				typeName := parts[1]

				// 处理字符串常量类型，如 POUType string
				if parts[2] == "string" {
					types[typeName] = GoTypeInfo{
						Name:   typeName,
						Fields: []string{"string_constant"},
					}
					continue
				}

				// 检查是否是结构体
				if parts[2] == "struct" && strings.HasSuffix(line, "{") {
					fields := []string{}

					// 解析字段
					j := i + 1
					for j < len(lines) {
						fieldLine := strings.TrimSpace(lines[j])

						// 结构体结束
						if fieldLine == "}" {
							break
						}

						// 添加字段
						if fieldLine != "" && !strings.HasPrefix(fieldLine, "//") {
							fieldParts := strings.Fields(fieldLine)
							if len(fieldParts) >= 2 {
								fieldName := fieldParts[0]
								fields = append(fields, fieldName)
							}
						}

						j++
					}

					types[typeName] = GoTypeInfo{
						Name:   typeName,
						Fields: fields,
					}

					// 更新索引
					i = j
				}
			}
		}
	}

	return types, nil
}

// 验证XSD和Go类型匹配
func ValidateTypesCoverage(xsdFileName, goFileName string) error {
	// 解析XSD文件
	parser := &XSDParser{}
	err := parser.ParseXSD(xsdFileName)
	if err != nil {
		return fmt.Errorf("解析XSD失败: %v", err)
	}

	// 解析Go文件
	goTypes, err := ExtractGoTypes(goFileName)
	if err != nil {
		return fmt.Errorf("解析Go文件失败: %v", err)
	}

	// 输出解析结果
	fmt.Printf("\n\n解析结果:\n")
	fmt.Printf("在XSD中找到 %d 个总复杂类型\n", parser.Count)
	fmt.Printf("在XSD中找到 %d 个命名复杂类型\n", len(parser.NamedTypes))
	fmt.Printf("在XSD中找到 %d 个匿名复杂类型\n", len(parser.AnonymousTypes))
	fmt.Printf("在Go文件中找到 %d 个类型\n", len(goTypes))

	// 寻找对应关系
	fmt.Println("\n命名复杂类型及其在Go代码中的对应情况:")
	for _, typeInfo := range parser.NamedTypes {
		// 尝试找到匹配的Go类型
		found := false
		for goTypeName := range goTypes {
			// 使用简单的名称匹配策略
			if strings.EqualFold(typeInfo.Name, goTypeName) ||
				strings.EqualFold("ppx:"+typeInfo.Name, goTypeName) {
				found = true
				fmt.Printf("✓ XSD类型 '%s' 与Go类型 '%s' 匹配\n", typeInfo.Name, goTypeName)
				break
			}
		}

		if !found {
			fmt.Printf("✗ XSD类型 '%s' 在Go代码中未找到对应类型\n", typeInfo.Name)
		}
	}
	// 检查Go类型是否都有对应的XSD类型
	fmt.Println("\nGo类型及其在XSD中的对应情况:")
	for goTypeName := range goTypes {
		found := false
		for _, typeInfo := range parser.NamedTypes {
			if strings.EqualFold(typeInfo.Name, goTypeName) ||
				strings.EqualFold("ppx:"+typeInfo.Name, goTypeName) {
				found = true
				break
			}
		}

		if !found {
			// 检查可能对应匿名复杂类型的情况
			fmt.Printf("? Go类型 '%s' 可能对应XSD中的匿名复杂类型\n", goTypeName)
		}
	}

	return nil
}

// // 主函数
// func main() {
// 	// 获取当前工作目录
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		fmt.Printf("获取当前目录失败: %v\n", err)
// 		return
// 	}

// 	// 构建XSD和Go文件路径
// 	xsdPath := filepath.Join(dir, "TC6_XML_V10_B.xsd")
// 	goPath := filepath.Join(dir, "tc6_xml_v10_b.go")

// 	// 验证类型覆盖
// 	err = ValidateTypesCoverage(xsdPath, goPath)
// 	if err != nil {
// 		fmt.Printf("验证失败: %v\n", err)
// 		return
// 	}
// }
