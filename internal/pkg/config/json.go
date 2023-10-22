package config

import (
	_ "embed"
	"encoding/json"
	"log"

	"github.com/tanveerprottoy/stdlib-go-template/config/embedded"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/file"
)

var (
	// configs
	Configs map[string]any
)

func init() {
	var b []byte
	if constant.ReadFromEmbedFile {
		b = readFromEmbeddedFile()
	} else {
		b = readFromFile()
	}
	_ = json.Unmarshal(b, &Configs)
	log.Print("configs: ", Configs)
}

func readFromEmbeddedFile() []byte {
	return embedded.Contents
}

func readFromFile() []byte {
	pwd, _ := file.GetPWD()
	b, _ := file.ReadFile(pwd + "/config.json")
	return b
}

func GetJsonValue(key string) any {
	return Configs[key]
}
