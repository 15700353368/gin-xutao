package config

type Server struct {
	Mysql Mysql `json:"mysql" yaml:"mysql"`
	System System `json:"system" yaml:"system"`
	Zap Zap `json:"zap" yaml:"zap"`
	Redis   Redis   `json:"redis" yaml:"redis"`
}