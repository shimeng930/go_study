package base

import "testing"

func TestString(t *testing.T) {
	t.Run("strLoop", func(t *testing.T) {
		strLoop("hello中国")
	})
}
