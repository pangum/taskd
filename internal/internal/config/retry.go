package config

import (
	"time"

	"github.com/pangum/pangu"
)

// Retry 重试配置
type Retry struct {
	// 间隔
	Interval time.Duration `default:"5s" json:"interval,omitempty" yaml:"interval" xml:"interval" toml:"interval"`
	// 最大次数
	Times uint8 `default:"10" json:"times,omitempty" yaml:"times" xml:"times" toml:"times"`
	// 每次重试拉取个数
	Count int `default:"10" json:"count,omitempty" yaml:"count" xml:"count" toml:"count"`
	// 最长执行时间
	Maximum time.Duration `default:"24h" json:"maximum,omitempty" yaml:"maximum" xml:"maximum" toml:"maximum"`
}

func retry(config *pangu.Config) (retry *Retry, err error) {
	retry = new(Retry)
	err = config.Build().Get(config)

	return
}
