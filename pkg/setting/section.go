package setting

import "time"

//配置转换

type ServerSettings struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettings struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileEXT      string
	LogMaxSize      int
	MaxBackups      int
	MaxAge          int
	Compress        bool
	LogLevel        string
}

type DBSettings struct {
	DBType       string
	Username     string
	Password     string
	DBName       string
	Host         string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
