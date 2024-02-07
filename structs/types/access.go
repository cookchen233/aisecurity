package types

type Access struct {
	ID      string `toml:"id" json:"id"`
	Title   string `toml:"title" json:"title"`
	Enabled bool   `toml:"enabled" json:"enabled"`
}
