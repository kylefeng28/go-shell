package shell

import (
	. "github.com/franela/goblin"
	"testing"
)

func Test(t *testing.T) {
	g := Goblin(t)

	g.Describe("Shell", func() {
		g.It("should be created with no errors", func() {
			shell, err := NewShell("/bin/bash")
			g.Assert(err).Equal(nil)
			defer shell.Close()
		})

		g.It("should echo hello", func() {
			shell, _ := NewShell("/bin/bash")
			defer shell.Close()
			out, err := shell.Run("echo hello")
			g.Assert(err).Equal(nil)
			g.Assert(out).Equal("hello\n")
		})
	})
}
