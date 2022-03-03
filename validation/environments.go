package validation

import "errors"

func Validate(source string, destination string) (err error) {
	if source == destination {
		return errors.New("error: source and destination can't match")
	}

	if destination == "production" {
		return errors.New("error: destination can't be production")
	}

	return nil
}
