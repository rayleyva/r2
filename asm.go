package main

import (
	"errors"
	"strconv"
)

/*
String format
1: (\w|\d)+(\w|\d|^\s)*\s
2: ".*"
*/

/*
* 1:
                            \s
                   +---------------> E
                   |  +---------+
       \w,\d       |  |  \w\d   |
    S ----------->  A1 <--------+
    |              |  ^
    |           '\'|  |
    |'\'           |  |
    +---> CONV1 <--+  |rune
            |         |
            +---------+

* 2:
                            "
                   +---------------> E
                   |  +---------+
           "       |  |  \w\d\s |
    S ----------->  A2 <--------+
                   |  ^
                '\'|  |
                   |  |
          CONV2 <--+  |rune
            |         |
            +---------+
*/
const (
	S = iota

	A1
	CONV1

	A2
	CONV2

	E
)

// Auto State Machine
type ASM struct {
	state int
}

func NewASM() *ASM {
	return &ASM{
		state: S,
	}
}

func (asm *ASM) init() {
	asm.state = S
}

func (asm *ASM) cleanup() {
	asm.state = S
}

func conv(r rune) (rune, error) {
	switch r {
	case rune('"'):
		return rune('"'), nil
	case rune('\''):
		return rune('\''), nil
	case rune('\\'):
		return rune('\\'), nil
	case rune('b'):
		return rune('\b'), nil
	case rune('f'):
		return rune('\f'), nil
	case rune('n'):
		return rune('\n'), nil
	case rune('r'):
		return rune('\r'), nil
	case rune('t'):
		return rune('\t'), nil
	default:
		return rune(' '), errors.New("Unrecognized char of: \\" + string(r))
	}
}

func (asm *ASM) accept(r rune) (rc rune, toAppend bool, end bool, err error) {
	switch asm.state {
	case S:
		if r == rune(' ') {
			return rune(' '), false, false, nil
		}
		if r == rune('"') {
			asm.state = A2
			return rune(' '), false, false, nil
		}
		if r == rune('\\') {
			asm.state = CONV1
			return rune(' '), false, false, nil
		}
		asm.state = A1
		return r, true, false, nil

	case A1:
		if r == rune('\\') {
			asm.state = CONV1
			return rune(' '), false, false, nil
		}
		if r == rune(' ') {
			asm.state = E
			return rune(' '), false, true, nil
		}
		return r, true, false, nil

	case CONV1:
		rc, err := conv(r)
		if err == nil {
			asm.state = A1
			return rc, true, false, err
		}
		return rune(' '), false, false, err

	case A2:
		if r == rune('"') {
			asm.state = E
			return rune(' '), false, true, nil
		}
		if r == rune('\\') {
			asm.state = CONV2
			return rune(' '), false, false, nil
		}
		return r, true, false, nil

	case CONV2:
		rc, err := conv(r)
		if err == nil {
			asm.state = A2
			return rc, true, false, nil
		}
		return rune(' '), false, false, err

	default:
		panic("FATAL: untouch code in asm")
	}
}

func (asm *ASM) GetString(s []rune) (string, int, error) {
	asm.init()
	defer asm.cleanup()

	rc := make([]rune, 0, 32)

	for i := 0; i < len(s); i++ {
		r, toAppend, end, err := asm.accept(s[i])
		if err != nil {
			return "", -1, err
		}
		if end {
			return string(rc), i + 1, nil
		}
		if toAppend {
			rc = append(rc, r)
		}
	}

	if asm.state != E && asm.state != A1 {
		return "", 0, errors.New("args error, state: " + strconv.Itoa(asm.state))
	}

	return string(rc), len(s), nil
}
