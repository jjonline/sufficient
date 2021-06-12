package cfg

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

// Viper 依赖viper实现的文件模式配置解析
// viper支持的格式见包内常量
type Viper struct{}

// Parse
// @param resource 文件路径（string） 或 配置资源（[]byte）
// @param cType    配置文件类型：请使用包内常量
// @param target   解析结果集引用（struct pointer）
func (v Viper) Parse(resource interface{}, cType string, target interface{}) error {
	if !IsCfgTypeSupport(cType) {
		return ConfigTypeNotSupport
	}

	var stream []byte
	var err error

	switch resource.(type) {
	case string:
		// file dir for check file exist
		filePath := resource.(string)
		if !IsFileExist(filePath) {
			return ConfigFileNotExist
		}

		stream, err = os.ReadFile(filePath)
		if err != nil {
			return ConfigFileParseFailed
		}
	case []byte:
		stream = resource.([]byte)
	default:
		return ConfigFileNotExist
	}

	// use viper parse
	vip := viper.New()
	vip.SetConfigType(cType)

	if err = vip.ReadConfig(strings.NewReader(string(stream))); err != nil {
		return ConfigFileParseFailed
	}

	// set config value
	return vip.Unmarshal(target)
}
