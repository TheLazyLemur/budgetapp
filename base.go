package budgetapp

import "embed"

//go:embed db/migrations/*.sql
var EmbedMigrations embed.FS
