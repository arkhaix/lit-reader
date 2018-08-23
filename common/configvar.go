package common

import (
	"os"
)

// ConfigVar returns a string which is set, in order of priority, to:
//	flag value > environment value > defaultValue
func ConfigVar(defaultValue string, envVarName string, flagVar *string) string {
	res := defaultValue
	if envValue, ok := os.LookupEnv(envVarName); ok {
		res = envValue
	}
	if flagVar != nil {
		res = *flagVar
	}
	return res
}
