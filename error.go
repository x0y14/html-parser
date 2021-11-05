package html_parser

import "fmt"

func UnexpectedRuneError(r rune) error {
	return fmt.Errorf("unexpected rune error : %v", string(r))
}
