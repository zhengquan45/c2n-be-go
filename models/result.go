package models

import (
	"encoding/json"
	"fmt"
)

type Result[T any] struct {
	Code    int    `json:"code"`    // 返回Code
	Data    T      `json:"data"`    // 返回数据
	Message string `json:"message"` // 返回消息
}

// 构造函数: 成功返回 (不含数据)
func Ok[T any]() Result[T] {
	return Result[T]{
		Code:    200,
		Data:    zeroValue[T](),
		Message: "success",
	}
}

// 构造函数: 成功返回 (包含数据)
func OkWithData[T any](data T) Result[T] {
	return Result[T]{
		Code:    200,
		Data:    data,
		Message: "success",
	}
}

// 构造函数: 失败返回 (自定义消息)
func Fail[T any](message string) Result[T] {
	return Result[T]{
		Code:    500,
		Data:    zeroValue[T](),
		Message: message,
	}
}

// 构造函数: 失败返回 (自定义状态码和消息)
func FailWithCode[T any](code int, message string) Result[T] {
	return Result[T]{
		Code:    code,
		Data:    zeroValue[T](),
		Message: message,
	}
}

// Helper function to get zero value of type T
func zeroValue[T any]() T {
	var zero T
	return zero
}

func main() {
	// 示例1: 成功 (不含数据)
	res1 := Ok[string]()
	jsonRes1, _ := json.Marshal(res1)
	fmt.Println(string(jsonRes1)) // {"code":200,"data":"","message":"success"}

	// 示例2: 成功 (包含数据)
	res2 := OkWithData[string]("Hello World")
	jsonRes2, _ := json.Marshal(res2)
	fmt.Println(string(jsonRes2)) // {"code":200,"data":"Hello World","message":"success"}

	// 示例3: 失败 (自定义消息)
	res3 := Fail[string]("Something went wrong")
	jsonRes3, _ := json.Marshal(res3)
	fmt.Println(string(jsonRes3)) // {"code":500,"data":"","message":"Something went wrong"}

	// 示例4: 失败 (自定义状态码和消息)
	res4 := FailWithCode[string](404, "Not Found")
	jsonRes4, _ := json.Marshal(res4)
	fmt.Println(string(jsonRes4)) // {"code":404,"data":"","message":"Not Found"}
}
