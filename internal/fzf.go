package internal

import (
	"errors"
	"fmt"
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
	fzf.Stderr = os.Stderr
	stdinPhony, err := fzf.StdinPipe()
	if err != nil {
		return err
	}
	for _, s := range b.ListBookMarks(tagFilter) {
		fmt.Fprintln(stdinPhony, s.Name)
	}
	fzfStdOut, err := fzf.Output()
	if err != nil {
		return err
	}
	bm := b.GetBookMark(string(fzfStdOut))
	if bm == nil {
		// The line ending doesn't work... sometimes
		bm = b.GetBookMark(string(fzfStdOut[:len(fzfStdOut)-1]))
		if bm == nil {
			return errors.New("Bookmark not in bookworm & newline trick didn't work")
		}
	}
	err = OpenURL(bm.Link)
	printIfVerbose(bm.Link)
	return err
}
