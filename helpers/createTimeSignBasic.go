package helpers

import (
	"fmt"
	"net/url"
	"sort"

	"github.com/Dasha-Kinsely/Countryside/configs"
)

func CreateTimeSignBasic(param url.Values) string {
	var key []string
	var str = ""
	for k := range param {
		if k != "sn" {
			key = append(key, k)
		}
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], param.Get(key[i]))
		} else {
			str = str + fmt.Sprintf("%v=%v", key[i], param.Get(key[i]))
		}
	}
	sign := MD5(MD5(str) + MD5(configs.APP_NAME+"6YJSuc50uJ18zj45"))
	return sign
}
