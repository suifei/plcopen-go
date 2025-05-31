# PLCopen-Go 更新日志

## [v1.1.1] - 2025-05-31

### 新增功能 ✨
- **JSON 序列化测试套件**: 添加了完整的 JSON 功能测试覆盖
  - 基础 JSON 序列化/反序列化测试
  - XML 与 JSON 格式兼容性测试
  - `omitempty` 功能验证测试

### 改进 🚀
- **文档优化**: 完善了 CHANGELOG.md 的内容和格式
- **README 增强**: 添加了 CHANGELOG 链接和版本历史部分
- **徽章更新**: 为 README 添加了 CHANGELOG 徽章

### 测试覆盖 🧪
- 新增 3 个专门的 JSON 序列化测试
- 验证 JSON 与 XML 格式的数据一致性
- 确保 `omitempty` 功能正常工作

### 技术细节 🔧
- 测试文件：`tests/json_serialization_test.go`
- 测试覆盖：序列化、反序列化、兼容性、字段处理
- 所有测试通过，确保功能稳定性

## [v1.1.0] - 2025-05-31

### 新增功能 ✨
- **JSON 序列化支持**: 为所有结构体添加了完整的 JSON 标签支持
  - 支持将 PLCopen 对象序列化为 JSON 格式
  - 支持从 JSON 格式反序列化为 PLCopen 对象
  - 智能处理 `omitempty` 选项，优化输出格式
  - 自动将驼峰命名转换为小写开头的驼峰命名（符合 JSON 规范）

### 改进 🚀
- **双重序列化格式**: 现在同时支持 XML 和 JSON 两种序列化格式
- **更好的互操作性**: JSON 支持使得数据交换更加灵活
- **API 兼容性**: 保持完全向后兼容，现有代码无需修改

### 技术细节 🔧
- 为 140+ 个结构体字段添加了 JSON 标签
- 特殊处理 `XMLName` 字段（使用 `json:"-"` 排除）
- 指针类型和切片类型自动添加 `omitempty` 选项
- 保持与现有 XML 标签的一致性

### 使用示例

#### XML 序列化（现有功能）
```go
xmlData, err := xml.MarshalIndent(project, "", "  ")
```

#### JSON 序列化（新功能）
```go
jsonData, err := json.MarshalIndent(project, "", "  ")
```

#### JSON 反序列化（新功能）
```go
var project plcopen.Project
err := json.Unmarshal(jsonData, &project)
```

### 兼容性说明
- 完全向后兼容 v1.0.x
- Go 版本要求保持不变（>=1.22.3）
- 所有现有 API 保持不变

### 文档更新
- 更新 README.md，添加 JSON 序列化示例
- 添加功能特性说明
- 提供完整的使用示例

---

## [v1.0.0] - 2025-05-30

### 初始版本 🎉
- 完整的 PLCopen XML TC6 V1.0B 标准支持
- 从 XSD 模式生成的 Go 结构体
- XML 模式验证支持
- 工具函数和实用程序
- 全面的测试覆盖率
- 支持多种编程语言（ST、FBD、LD、IL、SFC）

---

## 版本说明

本项目遵循 [语义化版本](https://semver.org/lang/zh-CN/) 规范：

- **MAJOR**: 不兼容的 API 变更
- **MINOR**: 向后兼容的功能新增
- **PATCH**: 向后兼容的问题修正

### 版本格式
`MAJOR.MINOR.PATCH`，例如 `1.1.0`

### 发布频率
- 主要版本：根据需要发布
- 次要版本：每月发布新功能
- 补丁版本：根据需要修复问题

---

*更多详细信息请参考 [GitHub Releases](https://github.com/suifei/plcopen-go/releases)*
