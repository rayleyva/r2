package main

import ()

// Auto State Machine
type ASM struct {
	state int
}

func NewASM() *ASM {
	return &ASM{}
}

func (asm *ASM) init() {
	asm.state = 0
}

func (asm *ASM) cleanup() {
	//
}

func (asm *ASM) accept(r rune) (rc rune, toAppend bool, end bool, err error) {
	if r == rune(' ') {
		end = true
		return
	}
	rc = r
	toAppend = true
	end = false
	return
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
	return string(rc), len(s), nil
}
