package util

import (
	"reflect"
)

func SplitArray(arry interface{}, num int64) [][]interface{} {
	kind := reflect.ValueOf(arry)
	var arr []interface{}
	for i := 0; i < kind.Len(); i++ {
		arr = append(arr, kind.Index(i))
	}
	max := int64(len(arr))
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]interface{}{arr}
	}
	//获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]interface{}, 0)
	//声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= quantity; i++ {
		end = i*num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i*num
	}
	return segments
}