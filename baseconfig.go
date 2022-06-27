package commond

// rewrite this file for your app flags
type BaseConfig struct {
	Debug bool `json:"debug"`
}

func (b *BaseConfig) Flags() Flags {
	return Flags{
		Flag{Name: "debug", Default: b.Debug, Ptr: &b.Debug, Description: "verbose logs"},
	}
}

func DefaultBaseConfig() *BaseConfig {
	return &BaseConfig{
		Debug: true,
	}
}
