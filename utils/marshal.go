package utils

import (
	"encoding/json"
	"encoding/xml"
	"os"
)

// JSON 核心函数
// 序列化 any → JSON（无缩进）
func ToJSON(v any) ([]byte, error) {
	return json.Marshal(v)
}

// 序列化 any → JSON（带缩进）
func ToJSONIndent(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

// 反序列化 JSON → 结构体（泛型）
func FromJSON[T any](data []byte) (T, error) {
	var result T
	return result, json.Unmarshal(data, &result)
}

// 序列化 map → JSON
func MapToJSON(m map[string]any) ([]byte, error) {
	return json.Marshal(m)
}

// 反序列化 JSON → map（简化版）
func JSONToMap(data []byte) (map[string]any, error) {
	return FromJSON[map[string]any](data)
}

// JSON Must 系列
// 序列化 any → JSON（出错 panic）
func MustToJSON(v any) []byte {
	data, err := ToJSON(v)
	if err != nil {
		panic(err)
	}
	return data
}

// 序列化 any → JSON（带缩进，出错 panic）
func MustToJSONIndent(v any) []byte {
	data, err := ToJSONIndent(v)
	if err != nil {
		panic(err)
	}
	return data
}

// 反序列化 JSON → 结构体（出错 panic）
func MustFromJSON[T any](data []byte) T {
	result, err := FromJSON[T](data)
	if err != nil {
		panic(err)
	}
	return result
}

// 序列化 map → JSON（出错 panic）
func MustMapToJSON(m map[string]any) []byte {
	data, err := MapToJSON(m)
	if err != nil {
		panic(err)
	}
	return data
}

// 反序列化 JSON → map（出错 panic）
func MustJSONToMap(data []byte) map[string]any {
	m, err := JSONToMap(data)
	if err != nil {
		panic(err)
	}
	return m
}

// JSON 文件操作
// 写入 JSON 文件
func WriteJSONFile(path string, v any) error {
	data, err := ToJSON(v)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
func WriteJSONFileIndent(path string, v any) error {
	data, err := ToJSONIndent(v)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// 读取 JSON 文件 → 结构体
func ReadJSONFile[T any](path string) (T, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		var zero T
		return zero, err
	}
	return FromJSON[T](data)
}

// Must 版本文件操作
func MustWriteJSONFile(path string, v any) {
	if err := WriteJSONFile(path, v); err != nil {
		panic(err)
	}
}
func MustWriteJSONFileIndent(path string, v any) {
	if err := WriteJSONFileIndent(path, v); err != nil {
		panic(err)
	}
}

// 读取 JSON 文件 → 结构体（出错 panic）
func MustReadJSONFile[T any](path string) T {
	result, err := ReadJSONFile[T](path)
	if err != nil {
		panic(err)
	}
	return result
}

// XML 补充函数
// 反序列化 XML → 结构体（泛型）
func FromXML[T any](data []byte) (T, error) {
	var result T
	return result, xml.Unmarshal(data, &result)
}

// Must 版本 XML 函数
func MustToXML(v any) []byte {
	data, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return data
}

// 反序列化 XML → 结构体（出错 panic）
func MustFromXML[T any](data []byte) T {
	result, err := FromXML[T](data)
	if err != nil {
		panic(err)
	}
	return result
}

// 结构体 → 写入 XML 文件
func WriteXMLFile(path string, v any) error {
	data, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// 读取 XML 文件 → 结构体
func ReadXMLFile[T any](path string) (T, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		var zero T
		return zero, err
	}
	var result T
	return result, xml.Unmarshal(data, &result)
}

// Must 版本 XML 文件操作
func MustWriteXMLFile(path string, v any) {
	if err := WriteXMLFile(path, v); err != nil {
		panic(err)
	}
}

// 读取 XML 文件 → 结构体（出错 panic）
func MustReadXMLFile[T any](path string) T {
	result, err := ReadXMLFile[T](path)
	if err != nil {
		panic(err)
	}
	return result
}

// 实用工具函数
// 合并多个 JSON 数据（覆盖式合并）
func MergeJSON(dataList ...[]byte) ([]byte, error) {
	merged := make(map[string]any)
	for _, data := range dataList {
		m := make(map[string]any)
		if err := json.Unmarshal(data, &m); err != nil {
			return nil, err
		}
		for k, v := range m {
			merged[k] = v
		}
	}
	return json.Marshal(merged)
}

// 快速验证 JSON 有效性
func IsValidJSON(data []byte) bool {
	return json.Valid(data)
}

// Must 版本合并
func MustMergeJSON(dataList ...[]byte) []byte {
	result, err := MergeJSON(dataList...)
	if err != nil {
		panic(err)
	}
	return result
}
