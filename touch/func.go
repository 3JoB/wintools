package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/3JoB/ulib/fsutil"
	"github.com/urfave/cli/v2"
)

func Touch(c *cli.Context) error {
	if c.Args().Len() == 0 {
		return errors.New("Please specify a file to touch")
	}
	for _, r := range c.Args().Slice() {
		if !fsutil.IsExist(r) {
			s, err := fsutil.Create(r)
			if err != nil {
				return err
			}
			s.Close()
		} else {
			if fsutil.IsDir(r) {
				return fmt.Errorf("%s is a directory", r)
			}
		}
		info, _ := os.Stat(r)
		mTime := info.ModTime()
		if c.Bool("m") {
			mTime = time.Now()
		}
		os.Chtimes(r, mTime, mTime)
	}
	return nil
}
