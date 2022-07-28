package initialize
import (
	"fmt"
	"github.com/fatih/color"

	//"github.com/fatih/color"
	"go.uber.org/zap"
	"sao-manage/global"
	"sao-manage/utils"
)
// InitLogger 初始化Logger
func InitLogger() {
	// 实例化zap 配置
	cfg := zap.NewDevelopmentConfig()
	// 注意global.Settings.LogsAddress是在settings-dev.yaml配置过的
	// 配置日志的输出地址
	cfg.OutputPaths = []string{
		fmt.Sprintf("%slog_%s.log", global.Settings.LogsAddress, utils.GetNowFormatTodayTime()), //
		"stdout",
	}
	color.Cyan(global.Settings.LogsAddress)
	// 创建logger实例
	logg, _ := cfg.Build()
	zap.ReplaceGlobals(logg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可

	global.Lg = logg  // 注册到全局变量中
}