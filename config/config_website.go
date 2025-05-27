package config

type Website struct {
	Logo        string `json:"logo" yaml:"logo"`
	Name        string `json:"name" yaml:"name"`
	Title       string `json:"title" yaml:"title"`
	Slogan      string `json:"slogan" yaml:"slogan"`
	SloganEn    string `json:"slogan_en" yaml:"slogan_en" mapstructure:"slogan_en"`
	Description string `json:"description" yaml:"description"`
	Version     string `json:"version" yaml:"version"`
	CreatedAt   string `json:"created_at" yaml:"created_at" mapstructure:"created_at"`
}
