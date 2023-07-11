package static

import "embed"

//go:embed txt/*
var Txt embed.FS
