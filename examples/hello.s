start main
section .data
hello_str:
	s "hello world"

section .text
main:
	LDI hello_str	; r0 = hello_str
	MOV R1, R0	; r1 = r0
	LDI #11		; r0 = 11 = len(hello_str)
	OST R1, R0	; write(0, hello_str, len(hello_str))
