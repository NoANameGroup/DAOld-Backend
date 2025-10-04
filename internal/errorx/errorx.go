// Copyright © 2025 univero. All rights reserved.
// Licensed under the GNU Affero General Public License v3 (AGPL-3.0).
// license that can be found in the LICENSE file.

package errorx

import (
	"errors"
	"fmt"

	"github.com/NoANameGroup/DAOld-Backend/pkg/log"
)

const unknowCode = 999

// Errorx 是HTTP服务的业务异常
// 若返回Errorx给前端, 则HTTP状态码应该是200, 且响应体为Errorx内容
// 最佳实践:
// - 业务处理链路的末端使用Errorx, PostProcess处理后给出用户友好的响应
// - 预定义一些Errorx作为常量
// - 除却末端的Errorx外, 其余的error照常处理
type Errorx struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func New(code int, msg string) *Errorx {
	return &Errorx{
		Code: code,
		Msg:  msg,
	}
}

// Error 实现了error接口, 返回错误字符串
func (e Errorx) Error() string {
	return fmt.Sprintf("code=%d, msg=%s", e.Code, e.Msg)
}

// EndE 的作用是记录错误日志, 并返回一个与err相同的Errorx
func EndE(err error) error {
	log.Error("error: ", err)
	var ex Errorx
	if errors.As(err, &ex) {
		return ex
	}
	return &Errorx{Code: unknowCode, Msg: err.Error()}
}

// EndM 记录错误日志, 并返回一个自定义消息的Errorx
func EndM(err error, msg string) error {
	log.Error("error: ", msg)
	return &Errorx{Code: unknowCode, Msg: msg}
}

// EndX 记录错误日志, 并返回一个自定义消息和code的Errorx
func EndX(err error, code int, msg string) error {
	log.Error("error: ", msg)
	return &Errorx{Code: code, Msg: msg}
}

// 定义常量错误
// 用户相关
var (
	ErrEmailExisted                = New(1001, "邮箱已被注册")
	ErrUsernameExisted             = New(1002, "用户名已被注册")
	ErrUsernameOrPasswordIncorrect = New(1003, "用户名或密码错误")
	ErrContextUserIDInvalid        = New(1004, "上下文中用户ID无效")
	ErrPasswordIncorrect           = New(1005, "密码错误")
	ErrOldAndNewPasswordSame       = New(1006, "新密码不能与原密码相同")
	ErrConfirmPasswordNotMatch     = New(1007, "确认密码与新密码不匹配")
	ErrConfirmationNotMatch        = New(1008, "确认信息不匹配")
	ErrBirthdayFormatInvalid       = New(1009, "生日格式无效")
	ErrUserPermissionsInsufficient = New(1010, "用户权限不足")
)
