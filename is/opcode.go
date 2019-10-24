package is

import (
	"strings"
)

type Opcode byte

const (
	LDR Opcode = iota
	STR
	LDI
	MOV
	POP
	PSH
	BNZ
	CAL
	ADR
	ADD
	SHR
	AND
	ORR
	EOR
	OST
	IST
	opUNK
)

var Opcodes = []Opcode{
	LDR,
	STR,
	LDI,
	MOV,
	POP,
	PSH,
	BNZ,
	CAL,
	ADR,
	ADD,
	SHR,
	AND,
	ORR,
	EOR,
	OST,
	IST,
}

var opStrings = []string{
	"LDR",
	"STR",
	"LDI",
	"MOV",
	"POP",
	"PSH",
	"BNZ",
	"CAL",
	"ADR",
	"ADD",
	"SHR",
	"AND",
	"ORR",
	"EOR",
	"OST",
	"IST",
}

func (op Opcode) String() string {
	if op >= opUNK {
		return "UNK"
	}

	return opStrings[op]
}

func DecodeOpcode(s string) Opcode {
	s = strings.ToLower(s)
	for i, v := range opStrings {
		if v == s {
			return Opcode(i)
		}
	}

	return opUNK
}
