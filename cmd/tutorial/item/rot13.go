package item

import (
	"bufio"
	"fmt"
	"go-learn/cmd/iohelper"
	"io"
	"os"
	"strings"
)

func StartRot13() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Print a string to get rot13 encoding")
	fmt.Println()

	for {
		command := iohelper.ReadString(reader);

		switch {

		case strings.Compare(command, "q") == 0:
			fmt.Println("Bye!")
			return
			break

		default:
			s := strings.NewReader(command)
			r := rot13Reader{s}
			io.Copy(os.Stdout, &r)
			fmt.Println()
			break
		}
	}
}

func rot13(r byte) byte {
	var base byte
	if r <= 'z' && r >= 'a' {
		base = 'a'

	} else if r <= 'Z' && r >= 'A' {
		base = 'A'

	} else {
		return r
	}

	b := r - base
	b = (b + 13) % 26
	return base + b
}

func (rr rot13Reader) Read(b []byte) (int, error) {
	l, err := rr.r.Read(b)
	for i := 0 ; i < l ; i++ {
		b[i] = rot13(b[i])
	}
	return l, err
}

type rot13Reader struct {
	r io.Reader
}
