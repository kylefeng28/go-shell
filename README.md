# go-shell

A small Go library to run shell commands.

Example:

```
package main

import (
	"fmt"
	. "github.com/kylefeng28/go-shell"
)

func main() {
	shell, err := NewShell("/bin/bash")
	if err != nil {
		panic("could not create shell")
	}
	defer shell.Close()

	out, err := shell.Run("echo hello world")
	if err != nil {
		panic("error execuing command")
	}

	fmt.Println(out)
}
```

## Future features
- [ ] stdin and stdout using channels
