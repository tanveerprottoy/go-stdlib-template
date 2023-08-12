package rbac

import (
	"encoding/json"
	"log"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/file"
)

var (
	// configs
	Data map[string]any
)

func init() {
	pwd, _ := file.GetPWD()
	log.Println(pwd)
	b, _ := file.ReadFile(pwd + "/config/rbac.json")
	_ = json.Unmarshal(b, &Data)
}

func GetJsonValue(key string) any {
	return Data[key]
}
