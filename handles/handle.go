package handles

import (
	"database/sql"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/jjonline/go-mod-library/logger"
	"gorm.io/gorm"
	"net/http"
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
	res := H(SuccessCode, "success", data)
	ctx.JSON(http.StatusOK, res)
}

// F gin失败响应--接管错误处理
func F(ctx *gin.Context, err error) {
	eErr := handle(err)
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

// handle 错误处理
func handle(err error) E {
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
		// mysql错误
		return DbError.Wrap(e)
	case redis.Error:
		// redis错误
		return RedisError.Wrap(err)
	case CE:
		return e.Wrap(nil)
	case E:
		return e
	}
	return UnknownError.Wrap(err)
}

// region 检查连接断开导致异常方法

// CauseByLostConnection 字符串匹配方式检查是否为断开连接导致出错
func CauseByLostConnection(err error) bool {
	if "" == err.Error() {
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
