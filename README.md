# PLCopen-Go

一个用于处理 IEC 61131-3 PLCopen XML 格式的 Go 库，支持从 XSD 模式生成的完整结构体定义。

[![Go Version](https://img.shields.io/badge/go-1.22.3+-blue.svg)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/suifei/plcopen-go)](https://goreportcard.com/report/github.com/suifei/plcopen-go)
[![Changelog](https://img.shields.io/badge/changelog-available-green.svg)](CHANGELOG.md)

## 功能特性

- 🔧 完整的 PLCopen XML TC6 V1.0B 标准支持
- 📝 从 XSD 模式自动生成的 Go 结构体
- ✅ XML 模式验证支持
- 🛠️ 丰富的工具函数和实用程序
- 📊 全面的测试覆盖率
- 🌐 支持多种编程语言（ST、FBD、LD 等）
- 🔄 支持 XML 和 JSON 双重序列化格式

## 安装

```bash
go get github.com/suifei/plcopen-go
```

## 快速开始

### 创建 PLCopen 项目

```go
package main

import (
    "encoding/xml"
    "fmt"
    "time"
    
    "github.com/suifei/plcopen-go"
)

func main() {
    // 创建项目
    project := &plcopen.Project{
        FileHeader: &plcopen.ProjectFileHeader{
            CompanyName:        "Your Company",
            ProductName:        "Your Product",
            ProductVersion:     "1.0",
            ContentDescription: "PLCopen XML project",
            CreationDateTime:   time.Now(),
        },
        ContentHeader: &plcopen.ProjectContentHeader{
            Name:         "MyProject",
            Version:      "1.0",
            Organization: "Your Organization",
            Author:       "Your Name",
            Language:     "en",
        },
    }
    
    // 序列化为 XML
    xmlData, err := xml.MarshalIndent(project, "", "  ")
    if err != nil {
        panic(err)
    }
    
    fmt.Println(string(xmlData))
}
```

### 解析 PLCopen XML 文件

```go
package main

import (
    "encoding/xml"
    
    "github.com/suifei/plcopen-go"
    "github.com/suifei/plcopen-go/utils"
)

func main() {
    // 读取 XML 文件
    xmlContent, err := utils.ReadFile("project.xml")
    if err != nil {
        panic(err)
    }
    
    // 解析 XML
    var project plcopen.Project
    err = xml.Unmarshal([]byte(xmlContent), &project)
    if err != nil {
        panic(err)
    }
    
    // 访问项目信息
    fmt.Printf("项目名称: %s\n", project.ContentHeader.Name)
    fmt.Printf("公司名称: %s\n", project.FileHeader.CompanyName)
}
```

### PLCopen TC6 V2.01（版式独立）

现有 `plcopen.Project` / XML / JSON 接口仍对应 V1.0B。对于根命名空间为
`http://www.plcopen.org/xml/tc6_0201` 的 UTF-8 文件，使用独立的无损文档读取接口：

```go
data, err := os.ReadFile("plc.xml")
if err != nil {
    return err
}
p, err := plcopen.ParseProjectV201(data)
if err != nil {
    return err
}
fmt.Println(p.Info().Name, p.Info().POUs)
return os.WriteFile("plc-copy.xml", p.XML(), 0644)
```

`Info` 提供项目名称、文件头摘要、数据类型名、POU 名称/类型/正文语言和配置名称。
`XML` / `WriteTo` 原样保留整个文件，包括 XHTML、厂商扩展及命名空间前缀。
修改 `Info` 的副本不会修改项目。该接口只检查 XML 和版本命名空间，不是完整 V2.01
结构体、编辑器、XSD 验证器、V1→V2 转换器或 PLC 可执行程序编译器。

#### 官方 XSD 与 OpenPLC 兼容案例

- [PLCopen 官方 V2.01 XSD ZIP](https://www.plcopen.org/download_file/force/d397ab98-6a21-47d5-9fb1-37da9e94ef34/342/) 内含 `tc6_xml_v201.xsd`。提取后用 `xmllint --nonet --noout --schema tc6_xml_v201.xsd plc-copy.xml` 做独立校验。
- 复现样例取自 [OpenPLC Editor 所用的 editor 子模块](https://github.com/thiagoralves/editor/tree/d18889e683840eb6e721c0d5538be9119aff64fd/examples)，不是对旧版本格式选择或所有 OpenPLC 工具链的保证。
  - `examples/Blink/plc.xml`：LD。
  - `examples/Multi_Language/plc.xml`：ST、FBD、SFC、IL、LD。
  - `examples/Traffic_Light_FBD/plc.xml`：FBD。
- 仓库只提交自编测试 XML。运行 `go test ./...`；设置 `PLCOPEN_V201_XSD` 为官方 XSD 的绝对路径可启用校验用例。
  如果也要测试上述 OpenPLC 源文件，将它们分别保存为 `openplc-blink.xml`、`openplc-multi-language.xml`、
  `openplc-traffic-fbd.xml`，并把 `PLCOPEN_OPENPLC_FIXTURES` 指向该目录。
  测试检查 POU/语言、逐字节回写和回写文件的官方 XSD 校验，不编译或运行 PLC 程序。

### JSON 序列化支持

```go
package main

import (
    "encoding/json"
    "fmt"
    "time"
    
    "github.com/suifei/plcopen-go"
)

func main() {
    // 创建项目
    project := &plcopen.Project{
        FileHeader: &plcopen.ProjectFileHeader{
            CompanyName:        "Your Company",
            ProductName:        "Your Product",
            ProductVersion:     "1.0",
            ContentDescription: "PLCopen JSON project",
            CreationDateTime:   time.Now(),
        },
        ContentHeader: &plcopen.ProjectContentHeader{
            Name:         "MyProject",
            Version:      "1.0",
            Organization: "Your Organization",
            Author:       "Your Name",
            Language:     "en",
        },
    }
    
    // 序列化为 JSON
    jsonData, err := json.MarshalIndent(project, "", "  ")
    if err != nil {
        panic(err)
    }
    
    fmt.Println(string(jsonData))
    
    // 从 JSON 反序列化
    var newProject plcopen.Project
    err = json.Unmarshal(jsonData, &newProject)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("反序列化后的项目名称: %s\n", newProject.ContentHeader.Name)
}
```

## 项目结构

```
plcopen-go/
├── tc6_xml_v10_b.go        # 主要的结构体定义
├── utils/                   # 工具函数
│   ├── file_utils.go       # 文件操作工具
│   ├── marshal.go          # 序列化工具
│   └── validate_types_coverage.go
├── tests/                   # 测试文件
│   ├── tc6_xml_v10_b_test.go
│   ├── xml_validator_test.go
│   └── ...
├── docs/                    # 文档和 XSD 文件
│   ├── TC6_XML_V10_B.xsd
│   └── TC6_XML_V101.pdf
└── go.mod
```

## 支持的 PLCopen 元素

### 项目结构
- [`Project`](tc6_xml_v10_b.go) - 根项目元素
- [`ProjectFileHeader`](tc6_xml_v10_b.go) - 文件头信息
- [`ProjectContentHeader`](tc6_xml_v10_b.go) - 内容头信息
- [`ProjectTypes`](tc6_xml_v10_b.go) - 类型定义
- [`ProjectInstances`](tc6_xml_v10_b.go) - 实例配置

### 编程语言支持
- ST (Structured Text) - 结构化文本
- FBD (Function Block Diagram) - 功能块图
- LD (Ladder Diagram) - 梯形图
- IL (Instruction List) - 指令表
- SFC (Sequential Function Chart) - 顺序功能图

### 数据类型
- 基本数据类型 (BOOL, INT, REAL 等)
- 用户定义的数据类型
- 数组和结构体
- 枚举类型

## 工具函数

### 文件操作

```go
import "github.com/suifei/plcopen-go/utils"

// 读取文件
content, err := utils.ReadFile("project.xml")

// 写入文件
err = utils.WriteFile("output.xml", xmlContent)

// 检查文件是否存在
exists := utils.FileExists("project.xml")

// 读取文件行
lines, err := utils.ReadFileLines("config.txt")
```

## 测试

运行所有测试：

```bash
go test ./...
```

运行特定测试：

```bash
go test ./tests -v
```

运行 XML 验证测试：

```bash
go test ./tests -run TestXMLSchemaValidation -v
```

## 验证

该库包含全面的 XML 模式验证功能，确保生成的 XML 符合 PLCopen 标准：

- XML 格式验证
- 命名空间验证
- 结构完整性验证
- 类型覆盖率测试

详见 [`xml_validator_test.go`](tests/xml_validator_test.go) 中的验证测试。

## 示例

查看 [`tests`](tests/) 目录中的各种示例：

- [`tc6_xml_v10_b_test.go`](tests/tc6_xml_v10_b_test.go) - 基本项目创建和序列化
- [`xml_validator_test.go`](tests/xml_validator_test.go) - XML 验证示例
- [`data_types_test.go`](tests/data_types_test.go) - 数据类型使用示例

## API 文档

完整的 API 文档请访问：[GoDoc](https://godoc.org/github.com/suifei/plcopen-go)

## 贡献

欢迎贡献代码！请遵循以下步骤：

1. Fork 本项目
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件。

## 相关资源

- [PLCopen 官方网站](https://www.plcopen.org/)
- [IEC 61131-3 标准](https://en.wikipedia.org/wiki/IEC_61131-3)
- [PLCopen XML 规范](https://www.plcopen.org/technical-activities/xml-exchange-format)

## 版本历史

查看项目的详细更新历史和版本说明，请参考：[CHANGELOG.md](CHANGELOG.md)

## 致谢

感谢 PLCopen 组织提供的 XML 交换格式标准，使得工业自动化项目之间的互操作性成为可能。

---

如果您觉得这个项目有用，请给它一个 ⭐！