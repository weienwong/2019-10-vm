; this might be wrong
start main
section .data
param:
	b 10

section .text
fib_ret:		; this is address 3
	pop r0
	ldi r1, #1
	psh r1
	psh r0
	ret
fib:
	pop r1		; r1 is the param
	ldi r2, #3
	biz r1, r2	; if r1 == 0 { fib_ret() }
	ldi r3, #1
	add r1, r3	; r1++
	biz r1, r2	; if r1 == 0 { fib_ret() }

	ldi r2, #254
	add r1, r2	; r1 -= 2
	psh r1		; push r1

	ldi r2, #255
	add r1, r2	; r1 --
	psh r1		; push r1

	cal fib		; result is in the stack
	pop r1
	pop r2
	psh r1
	psh r2		; swap the return value with the other param
	cal fib		; call fib again

	pop r1
	pop r2
	add r1, r2	; add the two returned values
	pop r2		; pop the return address
	psh r1		; push the return value
	psh r2		; push the return address
	ret		; return to value at the top of the stack

main:
	psh param
	cal fib
	pop r1
	ldi #1
	cal r1, r0

