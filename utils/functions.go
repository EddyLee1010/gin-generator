package utils

import (
	"github.com/jinzhu/inflection"
	"strings"
)

// SnakeToLowerCamel snake命名转小驼峰命名
func SnakeToLowerCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return inflection.Singular(strings.Join(parts, ""))
}

// SnakeToUpperCamel snake命名转大驼峰命名
func SnakeToUpperCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		if len(parts[i]) > 0 {
			// 每个单词首字母大写
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

// TableNameToStructName 使用jinzhu库，兼容gorm gen生成规则
// 把复数变成单数
func TableNameToStructName(tableName string) string {
	singular := inflection.Singular(tableName)
	return SnakeToUpperCamel(singular)
}
