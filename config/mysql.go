package config

type Mysql struct {
	Path         string `json:"path" yaml:"path"`                             // 服务器地址:端口
	Config       string `json:"config" yaml:"config"`                       // 高级配置
	Db       	string `json:"dbname" yaml:"db"`                     // 数据库名
	Username     string `json:"username" yaml:"username"`                 // 数据库用户名
	Password     string `json:"password" yaml:"password"`                 // 数据库密码
	MaxIdleConns int    `json:"maxIdleConns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `json:"maxOpenConns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `json:"logMode" yaml:"log-mode"`                  // 是否开启Gorm全局日志
	LogZap       bool   `json:"logZap" yaml:"log-zap"`                     // 是否通过zap写入日志文件
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Db + "?" + m.Config
}