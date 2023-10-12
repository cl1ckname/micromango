package user

import (
	"crypto/sha1"
	"fmt"
)

func hashString(s string, salt string) string {
	pwd := sha1.New()
	pwd.Write([]byte(s))
	pwd.Write([]byte(salt))
	return fmt.Sprintf("%x", pwd.Sum(nil))
}
