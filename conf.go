/*
 * @Author: berryberry
 * @LastAuthor: Do not edit
 * @since: 2019-05-10 11:38:56
 * @lastTime: 2019-05-29 20:24:54
 */
package ginfizz

var (
	FizzConfig *Config
)

const (
	DBTypeMysql = "mysql"
	DBTypeMongo = "mongo"
)

type Config struct {
	App     AppConfig
	Monitor MonitorConfig
}

type AppConfig struct {
	RunMode  string
	HttpPort int
	Limit    LimitConfig
	Log      LogConfig
	Pprof    PprofConfig
	//DB       DBConfig  #v0版本暂时不支持数据库操作
}

type MonitorConfig struct {
	Enable   bool
	HttpPort int
}

type LogConfig struct {
	Enable      bool
	LogLevel    string
	LogsDirPath string
	LogRotator  LogRotatorConfig
}

type DBConfig struct {
	Enable   bool
	DBType   string
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
	Charset  string
}

type PprofConfig struct {
	Enable bool
}

type LimitConfig struct {
	Enable     bool
	MaxAllowed int
}

type LogRotatorConfig struct {
	Filename   string // 日志文件名
	MaxSize    int    // megabytes
	MaxBackups int    // 最多保留3个备份
	MaxAge     int    // days
	Compress   bool   // 是否压缩 disabled by default
}

func init() {
	if FizzConfig == nil {
		FizzConfig = &Config{
			App: AppConfig{
				RunMode:  "dev",
				HttpPort: 8080,
				Log: LogConfig{
					Enable:      true,
					LogLevel:    "info",
					LogsDirPath: "logs",
					LogRotator: LogRotatorConfig{
						Filename:   "tama.log",
						MaxSize:    10,
						MaxBackups: 30,
						MaxAge:     30,
						Compress:   false,
					},
				},
				Limit: LimitConfig{
					Enable:     true,
					MaxAllowed: 100,
				},
				Pprof: PprofConfig{
					Enable: true,
				},
				// DB: DBConfig{
				// 	Enable:   false,
				// 	DBType:   DBTypeMysql,
				// 	Username: "root",
				// 	Password: "root",
				// 	Host:     "localhost",
				// 	Port:     3306,
				// 	DBName:   "tama",
				// 	Charset:  "utf8",
				// },
			},
			Monitor: MonitorConfig{
				Enable:   true,
				HttpPort: 10010,
			},
		}
	}
}
