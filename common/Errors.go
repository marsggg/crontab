package common

import "github.com/pkg/errors"

var (
	ERR_LOCK_ALREADY_REQUIRED = errors.New("锁已被占用")
	ERR_CRONEXPR_PARSE        = errors.New("CronExpr表达式解析失败")
)
