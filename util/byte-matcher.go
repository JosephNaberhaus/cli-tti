package util

import (
	"errors"
)

type ByteMatcher []byte

// Matches the given bytes ignoring extra bytes. Returns an error if one of the bytes doesn't match
func (b *ByteMatcher) Match(newBytes []byte) error {
	for _, newByte := range newBytes {
		if b.Complete() {
			break
		}

		if (*b)[0] == newByte {
			*b = (*b)[1:]
		} else {
			return errors.New("bytes do not match")
		}
	}

	return nil
}

// Returns true if all of the bytes have been matched
func (b *ByteMatcher) Complete() bool {
	return len(*b) == 0
}