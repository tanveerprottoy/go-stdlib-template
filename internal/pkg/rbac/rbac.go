package rbac

import (
	"fmt"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/stringspkg"
)

// GetRBAC fetches from the Role based access control
// json file if the data exists
func GetRBAC(path, method string) {
	s := stringspkg.Split(path, "/")
	fmt.Println(s)
}