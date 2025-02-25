// +build linux

package main

/*
#cgo CFLAGS: -g -Wall -I./include
#cgo LDFLAGS: -lmsc -lrt -ldl -lpthread -lstdc++

#include "iat.h"
*/
import "C"

import (
	"fmt"
	"sync"
)

// APPID 讯飞 SDK ID
var APPID = "5ec24fa4"
var mutex sync.Mutex

func wavToText(path string) string {
	var result string
	// 登录参数，appid与msc库绑定,请勿随意改动
	loginParams := fmt.Sprintf("appid = %s, work_dir = .", APPID)

	/*
	* sub:				请求业务类型
	* domain:			领域
	* language:			语言
	* accent:			方言
	* sample_rate:		音频采样率
	* result_type:		识别结果格式
	* result_encoding:	结果编码格式
	*
	 */
	// const char* session_begin_params ="sub = iat, domain = iat, language = zh_cn, accent = mandarin, sample_rate = 16000, result_type = plain, result_encoding = utf8";
	sessionBeginParams := fmt.Sprintf("sub = iat, domain = iat, language = zh_cn, accent = mandarin, sample_rate = 16000, result_type = plain, result_encoding = utf8")

	/* 用户登录 */
	//第一个参数是用户名，第二个参数是密码，均传NULL即可，第三个参数是登录参数
	ret := C.MSPLogin(nil, nil, C.CString(loginParams))
	if C.MSP_SUCCESS != ret {
		fmt.Printf("MSPLogin failed , Error code %d.\n", ret)
		goto exit //登录失败，退出登录
	}
	result = C.GoString(C.run_iat(C.CString(path), C.CString(sessionBeginParams)))

exit:
	C.MSPLogout() //退出登录
	return result
}
