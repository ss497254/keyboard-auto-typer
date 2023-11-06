package public

import (
	"embed"
)

//go:embed all:*
var WebsiteSource embed.FS
