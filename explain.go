package main

import (
	"bufio"
	"bytes"
	"io"
)

func Launch(r io.Reader) (int, error) {
	defer recover()

	reader := bufio.NewReader(r)

	for lineNo := 1; ; lineNo++ {
		buf := make([]byte, 0, 32)
		for {
			bs, err := reader.ReadBytes(byte('\n'))
			if err != nil {
				if err == io.EOF {
					return 0, nil
				}
				return lineNo, err
			}

			n := len(bs)
			if n == 1 {
				break
			}

			bs[n-1] = ' '
			if bs[n-2] != '\\' {
				buf = append(buf, bs...)
				break
			}

			buf = append(buf, bs[:n-2]...)
			buf = append(buf, ' ')
			lineNo++
		}

		buf = bytes.TrimSpace(buf)

		if len(buf) == 0 {
			continue
		}

		if buf[0] == '#' {
			continue
		}

		cmd, err := CmdParser(string(buf))

		if err != nil {
			return lineNo, err
		}

		if err := cmd.Exec(); err != nil {
			return lineNo, err
		}
	}
}
