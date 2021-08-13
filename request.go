package base

import (
	"net/http"

	"github.com/gavv/httpexpect"
)

// PageRes 基础分页请求参数
var PageRes = map[string]interface{}{"page": 1, "pageSize": 10}

// PageKeys  基础分页请求返回参数
var PageKeys = ResponseKeys{
	{Type: "int", Key: "pageSize", Value: 10},
	{Type: "int", Key: "page", Value: 1},
	{Type: "int", Key: "total", Value: 0},
}

// PostList 分页请求-post
// auth  授权后的 *httpexpect.Expect
// url 请求地址
// res 请求参数
// pageKeys 相应数据
func PostList(auth *httpexpect.Expect, url string, res map[string]interface{}, pageKeys ResponseKeys) {
	obj := auth.POST(url).WithJSON(res).Expect().Status(http.StatusOK).JSON().Object()
	pageKeys.Test(obj)
}

// GetList 分页请求-get
// auth  授权后的 *httpexpect.Expect
// url 请求地址
// keys 响应数据
func GetList(auth *httpexpect.Expect, url string, keys ResponseKeys) {
	obj := auth.GET(url).Expect().Status(http.StatusOK).JSON().Object()
	keys.Test(obj)
}

// Create 创建数据
// auth  授权后的 *httpexpect.Expect
// url 请求地址
// create 请求数据
// keys 响应数据
func Create(auth *httpexpect.Expect, url string, create map[string]interface{}, keys ResponseKeys) *httpexpect.Object {
	return auth.POST(url).WithJSON(create).Expect().Status(http.StatusOK).JSON().Object()
}

// Update 更新数据
// auth  授权后的 *httpexpect.Expect
// url 请求地址
// update 请求数据
// keys 响应数据
func Update(auth *httpexpect.Expect, url string, update map[string]interface{}, keys ResponseKeys) {
	obj := auth.PUT(url).WithJSON(update).Expect().Status(http.StatusOK).JSON().Object()
	keys.Test(obj)
}

// Get 获取数据
// auth  授权后的 *httpexpect.Expect
// url 请求地址
// keys 响应数据
func Get(auth *httpexpect.Expect, url string, keys ResponseKeys) {
	obj := auth.GET(url).Expect().Status(http.StatusOK).JSON().Object()
	keys.Test(obj)
}

// Post 提交数据
// auth  授权后的 *httpexpect.Expect
// url 请求地址
// data 请求数据
// keys 响应数据
func Post(auth *httpexpect.Expect, url string, data map[string]interface{}, keys ResponseKeys) {
	obj := auth.POST(url).WithJSON(data).Expect().Status(http.StatusOK).JSON().Object()
	keys.Test(obj)
}

// Delete 删除数据
// auth  授权后的 *httpexpect.Expect
// url 请求地址
// keys 响应数据
func Delete(auth *httpexpect.Expect, url string, keys ResponseKeys) {
	obj := auth.DELETE(url).Expect().Status(http.StatusOK).JSON().Object()
	keys.Test(obj)
}

// DeleteMutil 批量删除数据
// auth  授权后的 *httpexpect.Expect
// url 请求地址
// data 请求数据
// keys 响应数据
func DeleteMutil(auth *httpexpect.Expect, url string, data map[string]interface{}, keys ResponseKeys) {
	obj := auth.DELETE(url).WithJSON(data).Expect().Status(http.StatusOK).JSON().Object()
	keys.Test(obj)
}
