package config

type Captcha struct {
	Height   int     `json:"height" yaml:"height"`
	Width    int     `json:"width" yaml:"width"`
	Length   int     `json:"length" yaml:"length"`
	MaxSkew  float64 `json:"max_skew" yaml:"max_skew" mapstructure:"max_skew"`
	DotCount int     `json:"doc_count" yaml:"doc_count" mapstructure:"doc_count"`
}
