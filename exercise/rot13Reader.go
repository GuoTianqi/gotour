package exercise

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rotR *rot13Reader) Read(buff []byte) (int, error) {
	// TODO: 未实现
	return 0, nil
}

func rot13ReaderMain() {
	s := strings.NewReader("Lbh penpxrq gurpbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
