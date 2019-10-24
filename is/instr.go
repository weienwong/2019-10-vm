package is

type Instr byte

func (instr Instr) Op() Opcode {
	// TODO
	return 0
}

func (instr Instr) R1() Reg {
	// TODO
	return 0
}

func (instr Instr) R2() Reg {
	// TODO
	return 0
}

func (instr Instr) IReg() IReg {
	return IReg(instr.R2())
}

func (instr Instr) Imm() byte {
	// TODO
	return 0
}

/*
This could be helpful for debugging. Given a binary file it would return
the assmebly code used to create it.
func (instr Instr) String() string {

}
*/

func (instr Instr) Encode() byte {
	// TODO
	return 0
}

func DecodeInstr(b byte) Instr {
	// TODO
	return 0
}

type Reg byte

const (
	R0 Reg = iota
	R1
	R2
	R3
)

type IReg byte

const (
	PC IReg = iota
	FP
	SP
)
