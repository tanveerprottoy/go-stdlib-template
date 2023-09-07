package embedded

import _ "embed"

//go:embed dev.json
var Contents []byte
