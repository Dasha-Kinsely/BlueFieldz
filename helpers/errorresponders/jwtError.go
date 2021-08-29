package errorresponders

import "errors"

func JwtError(errorType string) (error) {
	switch errorType{
		case "malformed":
			return errors.New("token is malformed...\n")
		case "expired":
			return errors.New("token expired...\n")
		case "generation":
			return errors.New("error occurred at generation step, might be cased by invalid signiture...")
		default:
			return errors.New("unknown error regarding jwt reading has occurred...\n")
	}
}