package handles

// 错误码定义
// 约定:
// 错误码使用建议放在 service (包裹其他错误) 返回
// 而 model层不用使用错误码
// controller层仅在必要得情况下使用错误码 如 参数校验错误
const (
	SuccessCode     = 0
	UnknownError CE = -1 // 未知錯誤
	// 基础服务错误
	DbError    CE = 100001 // 系統錯誤
	EsError    CE = 100002 // 系統錯誤
	AwsError   CE = 100003 // 系統錯誤
	RedisError CE = 100004 // 系統錯誤

	// 系统/服务错误
	InvalidParams       CE = 101001 // 請求參數錯誤
	PlatformParseFailed CE = 101002 // 請求不合法
	SystemBusy          CE = 101003 // 系統繁忙，請稍後重試
	DbRecordNotExist    CE = 101404 // 數據不存在

)
