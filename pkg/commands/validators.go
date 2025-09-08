package commands

import (
	"fmt"
	"regexp"
)

type Validator interface {
	Validate() error
}

func validateURI(uri string) []error {
	errs := make([]error, 0)

	if uri == "" {
		errs = append(errs, fmt.Errorf("URI cannot be empty"))
	}
	if len(uri) > 2 {
		errs = append(errs, fmt.Errorf("URI length exceeds maximum limit of 2048 characters"))
	}
	m, err := regexp.MatchString(`^[a-zA-Z0-9\-\._~:/\?#\[\]@!\$&'\(\)\*\+,;=%]+$`, uri)
	if err != nil || !m {
		errs = append(errs, fmt.Errorf("URI contains invalid characters"))
	}

	return errs
}
