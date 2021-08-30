package render

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/jjonline/go-lib-backend/logger"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

// H 格式化响应
func H(code int, msg string, data interface{}) gin.H {
	return gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}

// S gin成功响应
func S(ctx *gin.Context, data interface{}) {
	res := H(SuccessCode, "ok", data)
	ctx.JSON(http.StatusOK, res)
}

// F gin失败响应--接管错误处理
func F(ctx *gin.Context, err error) {
	eErr := translateError(err)
	// 记录错误日志
	LogErrWithGin(ctx, eErr, false)
	res := H(eErr.Code(), eErr.Format(), nil)
	ctx.JSON(http.StatusOK, res)
}

// LogErr 记录错误日志
func LogErr(err error, mark string, isAlert bool) {

}

// LogErrWithGin 记录错误日志
func LogErrWithGin(ctx *gin.Context, err error, isAlert bool) {
	logger.GinLogHttpFail(ctx, err)
}

// translateError 默认错误翻译
func translateError(err error) E {
	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, sql.ErrNoRows) {
		// gorm 数据不存在错误
		return DbRecordNotExist.Wrap(err)
	}

	if CauseByLostConnection(err) {
		// 各种原因丢失链接导致异常
		return LostConnectionError.Wrap(err)
	}

	switch e := err.(type) {
	case *mysql.MySQLError:
		return DbError.Wrap(e) // mysql错误
	case redis.Error:
		// redis错误
		return RedisError.Wrap(err)
	case *validator.InvalidValidationError:
		return ErrDefineWithMsg.Wrap(err, "内部错误：参数绑定条件语法错误") // 需要修改参数结构体的tag为binding条件
	case validator.ValidationErrors:
		return ErrDefineWithMsg.Wrap(err, "参数错误：参数值未满足限定条件")
	case *json.UnmarshalTypeError:
		return ErrDefineWithMsg.Wrap(err, "参数错误：参数值类型不符")
	case *strconv.NumError:
		if e.Func == "ParseBool" {
			return ErrDefineWithMsg.Wrap(err, "参数错误：限定传参布尔值")
		}
		return ErrDefineWithMsg.Wrap(err, "参数错误：数值范围超过限定类型界限")
	case CE:
		return e.Wrap(nil)
	case E:
		return e
	default:
		return UnknownError.Wrap(err)
	}
}

// region 检查连接断开导致异常方法

// CauseByLostConnection 字符串匹配方式检查是否为断开连接导致出错
func CauseByLostConnection(err error) bool {
	if err == nil || "" == err.Error() {
		return false
	}

	needles := []string{
		"server has gone away",
		"no connection to the server",
		"lost connection",
		"is dead or not enabled",
		"error while sending",
		"decryption failed or bad record mac",
		"server closed the connection unexpectedly",
		"ssl connection has been closed unexpectedly",
		"error writing data to the connection",
		"resource deadlock avoided",
		"transaction() on null",
		"child connection forced to terminate due to client_idle_limit",
		"query_wait_timeout",
		"reset by peer",
		"broken pipe",
		"connection refused",
	}

	msg := strings.ToLower(err.Error())
	for _, needle := range needles {
		if strings.Contains(msg, needle) {
			return true
		}
	}
	return false
}

// endregion
