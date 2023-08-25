package rbac

import (
	"fmt"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/stringsext"
)

func buildAndFindForKey(slice []string, offset int, method, param string) any {
	// build the key
	l := len(slice)
	if l-offset >= 0 {
		var k string
		if param == "" {
			k = fmt.Sprintf("%s.%s", slice[len(slice)-offset], method)
		} else {
			k = fmt.Sprintf("%s.%s.%s", slice[len(slice)-offset], method, param)
		}
		fmt.Println(k)
		return GetJsonValue(k)
	}
	return nil
}

// GetRBAC fetches from the Role based access control
// json file if the data exists
func GetRBAC(path, method string) any {
	s := stringsext.Split(path, "/")
	fmt.Println(s)
	v := buildAndFindForKey(s, 1, method, "")
	if v == nil {
		// try with the value before the last one
		return buildAndFindForKey(s, 2, method, "p")
	}
	return nil
}
