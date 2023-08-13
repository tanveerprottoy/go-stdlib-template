package rbac

import (
	"encoding/json"
	"log"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/file"
)

var (
	// configs
	Data map[string]RBACModel
)

func init() {
	pwd, _ := file.GetPWD()
	log.Println(pwd)
	b, _ := file.ReadFile(pwd + "/config/rbac.json")
	_ = json.Unmarshal(b, &Data)
}

func GetJsonValue(key string) any {
	v, ok := Data[key]
	// If the key exists
	if ok {
		return v
	}
	return nil
}

type RBACModel struct {
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}
