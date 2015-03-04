package apass

import (
	"strings"
	"time"

	mrand "math/rand"

	"github.com/zerklabs/fifoq/string"
)

const (
	// If a value less than or equal to zero is given to Generate, the DefaultSegmentCount will be
	// used in its place
	DefaultSegmentCount = 4
	alphabet            = "0123456789abcdefghijklmnopqrstuvwxyz"
)

var ()

// Generate will return a multi-part character password,
// similar to those used by Google application passwords
func Generate(segments int) string {
	var password string

	if segments <= 0 {
		segments = DefaultSegmentCount
	}

	mrand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < segments; i++ {
		q := fifoqstring.New(8)

		for j := 0; j < 8; j++ {
			for {
				letter := getLetter()
				if !q.Contains(letter) {
					q.Push(letter)
					break
				}
			}
		}

		// combine the fifoq into a string
		password = password + strings.Join(q.View(), "")

		// separate the segments with a space
		if i < segments-1 {
			password = password + " "
		}
	}

	return password
}

// returns a random character from the alphabet
func getLetter() string {
	pos := mrand.Intn(len(alphabet))

	return string(alphabet[pos])
}
