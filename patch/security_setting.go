package patch

import (
	"fmt"
	"regexp"
)

// RemoveEnvironmentVariableContend permit to nto compare env variable
// Opensearch substitute env variable with real value
func RemoveEnvironmentVariableContend(actualByte []byte, expectedByte []byte) ([]byte, []byte, error) {

	rEnvVar := regexp.MustCompile(`("[^"]+"):("\$\{env\.[^"]+\}")`)

	matches := rEnvVar.FindAllStringSubmatch(string(expectedByte), -1)
	for _, match := range matches {
		rSubstitute := regexp.MustCompile(fmt.Sprintf(`%s:"[^"]*"`, match[1]))
		actualByte = rSubstitute.ReplaceAll(actualByte, []byte(match[0]))
	}

	return actualByte, expectedByte, nil
}
