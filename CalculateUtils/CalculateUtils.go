package CalculateUtils

import (
	"fmt"
	"strconv"
)

type CalculateUtils struct {
}

// PercentageCode
/**
 *  求两数的百分比(带符号)
 * @param v1 除数
 * @param v2 被除数
 * @param scale 要保留多少位小数
 * @param code 拼接符号，例如：%
 * @return
 */
func (CalculateUtils) PercentageCode(v1, v2, scale int, code string) string {

	itoa := strconv.Itoa(scale)

	result := float64(v1) / float64(v2) * 100

	// 保留两位小数
	formattedResult := fmt.Sprintf("%."+itoa+"f", result)

	return formattedResult + code
}

// Percentage
/**
 *  求两数的百分比
 * @param v1 除数
 * @param v2 被除数
 * @param scale 要保留多少位小数
 * @return
 */
func (CalculateUtils) Percentage(v1, v2, scale int) float64 {

	itoa := strconv.Itoa(scale)

	result := float64(v1) / float64(v2) * 100

	// 保留两位小数
	formattedResult := fmt.Sprintf("%."+itoa+"f", result)

	value, err := strconv.ParseFloat(formattedResult, 64)
	if err != nil {
		fmt.Println("Percentage:类型转换错误")
	}

	return value
}
