package config

type System struct {
	Env string `json:"env" yaml:"env"`
	Addr string `json:"addr" yaml:"addr"`
	DbType        string `json:"dbType" yaml:"db-type"`                      // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string `json:"ossType" yaml:"oss-type"`                   // Oss类型
	UseMultipoint bool   `json:"useMultipoint" yaml:"use-multipoint"` // 多点登录拦截
	UcsKey   string   `json:"ucs_key" yaml:"ucsKey"` 												// 签名密钥
	UcsTimeStamp   int   `json:"ucs_time_stamp" yaml:"ucsTimeStamp"` 								// 签名过期时间
}