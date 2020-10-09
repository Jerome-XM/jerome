package util

/**
响应成功返回函数
 */
func SuccResp(message string,data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	result["success"] = true
	result["message"] = message
	if data != nil {
		result["data"] = data
	}
	return result
}
/**
响应失败返回函数
 */
func FalseResp(message string) map[string]interface{} {
	result := make(map[string]interface{})
	result["success"] = false
	result["message"] = message
	return result
}

