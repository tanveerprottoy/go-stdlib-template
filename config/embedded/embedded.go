package embedded

import _ "embed"

//go:embed configs.json
var Contents []byte
