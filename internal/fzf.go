package internal

import (
	"os"
	"os/exec"
)

// Uses Fzf to open a BookMark
// Inspo https://github.com/zk-org/zk/blob/dev/internal/adapter/fzf/fzf.go
func (b BookWorm) FzfOpen(tagFilter string) error {
	fzfExec, err := exec.LookPath("fzf")
	if err != nil {
		return err
	}
	fzf := exec.Command(fzfExec)
	rPipe, wPipe, err := os.Pipe()
	if err != nil {
		return err
	}
	defer wPipe.Close()
	defer rPipe.Close()

	for _, s := range b.ListBookMarks(tagFilter) {
		_, err = wPipe.WriteString(s.Name)
		if err != nil {
			return err
		}
	}
	// This works with stdin, but not a pipe... why?
	fzf.Stdin = rPipe //os.Stdin
	fzfStdOut, err := fzf.Output()
	if err != nil {
		return err
	}
	return OpenURL(string(fzfStdOut))
}
