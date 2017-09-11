package teamcity

import (
	"fmt"
	"regexp"
)

func ValidateID(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if !regexp.MustCompile(`^[0-9A-Za-z_]+$`).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q should contain only latin letters, digits and underscores: %q",
			k, value))
	}
	if len(value) > 80 {
		errors = append(errors, fmt.Errorf(
			"%q cannot be longer than 80 characters: %q", k, value))
	}
	if regexp.MustCompile(`^-`).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q cannot begin with a hyphen: %q", k, value))
	}
	if regexp.MustCompile(`^[0-9]`).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q cannot begin with a digit: %q", k, value))
	}
	return

}
