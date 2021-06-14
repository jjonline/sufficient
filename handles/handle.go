package handles

import (
	"database/sql"
	"gorm.io/gorm"
	"net/http"

	"github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

// H 格式化响应
func H(code int, msg string, data interface{}) gin.H {
	return gin.H{
		"code":   code,
		"msg":    msg,
		"data":   data,
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

}

// handle 错误处理
func handle(err error) E {
	if err == gorm.ErrRecordNotFound || err == sql.ErrNoRows {
		return DbRecordNotExist.Wrap(err)
	}
	switch e := err.(type) {
	case *mysql.MySQLError: // mysql错误
		return DbError.Wrap(e)
	case CE:
		return e.Wrap(nil)
	case E:
		return e
	}
	return UnknownError.Wrap(err)
}
