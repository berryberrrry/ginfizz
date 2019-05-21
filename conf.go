/*
 * @Author: berryberry
 * @LastAuthor: Do not edit
 * @since: 2019-05-10 11:38:56
 * @lastTime: 2019-05-21 22:09:05
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
	DB       DBConfig
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

type LimitConfig struct {
	Enable     bool
	MaxAllowed int
}

type LogRotatorConfig struct {
	Filename   string // 日志文件路径
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
				DB: DBConfig{
					Enable:   true,
					DBType:   DBTypeMysql,
					Username: "root",
					Password: "root",
					Host:     "localhost",
					Port:     3306,
					DBName:   "tama",
					Charset:  "utf8",
				},
				Limit: LimitConfig{
					Enable:     true,
					MaxAllowed: 3,
				},
			},
			Monitor: MonitorConfig{
				Enable:   true,
				HttpPort: 10010,
			},
		}
	}
}
