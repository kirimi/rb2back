package migrations

import "embed"

//go:embed *.sql
var EmbedFs embed.FS
