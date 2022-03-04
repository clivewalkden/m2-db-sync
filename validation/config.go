package validation

import "github.com/clivewalkden/m2-db-sync/common"
import "errors"

func ConfigValidation(config common.Config) (err error) {
	// old and new domains count match
	if len(config.Src.Domain) != len(config.Dest.Domain) {
		return errors.New("error: source and destination domains count doesn't match")
	}

	return nil
}
