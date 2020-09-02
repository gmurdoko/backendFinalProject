package config

import "finalproject/utils/getEnv"

//AuthSwitch app
func UseLogActivity() bool {
	isTrue := getEnv.ViperGetEnv("LOGACTIVITY", "YES")
	if isTrue == "YES" {
		return true
	}
	return false
}

func AuthSwitch() bool {
	isTrue := getEnv.ViperGetEnv("AUTH", "NO")
	if isTrue == "YES" {
		return true
	}
	return false
}
