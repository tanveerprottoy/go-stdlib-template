package rbac

import (
	"fmt"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/stringspkg"
)

func buildAndFindForKey(slice []string, offset int, method string) any {
	// build the key
	k := fmt.Sprintf("%s.%s", slice[len(slice)-offset], method)
	return GetJsonValue(k)
}

// GetRBAC fetches from the Role based access control
// json file if the data exists
func GetRBAC(path, method string) any {
	s := stringspkg.Split(path, "/")
	fmt.Println(s)
	if len(s) > 0 {
		v := buildAndFindForKey(s, 1, method)
		if v == nil {
			// try with the value before the last one
			return buildAndFindForKey(s, 2, method)
		}
	}
	return nil
}
