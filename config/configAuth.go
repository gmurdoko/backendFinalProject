package config

import "finalproject/utils/getEnv"

//AuthSwitch app
func UseLogActivity() bool {
	isTrue := getEnv.ViperGetEnv("Auth", "YES")
	if isTrue == "YES" {
		return true
	}
	return false
}
