// Copyright 2022 The RomiChan WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// func mask(b *byte, len int, key uint64)
// Requires: AVX, AVX2
TEXT ·maskAsm(SB), NOSPLIT, $0-32
	MOVQ b_base+0(FP), AX
	MOVQ b_len+8(FP), CX
	MOVQ key+16(FP), DI
	MOVQ AX, DX
	ADDQ CX, DX
	CMPB ·useAVX2(SB), $1
	JE   avx2
	JMP  sse

// todo(wdvxdr): optimize unaligned case
avx2:
	VPBROADCASTQ key+16(FP), Y0
	MOVQ         AX, CX

avx2_loop:
	ADDQ    $0x80, CX
	CMPQ    CX, DX
	JAE     sse
	VPXOR   (AX), Y0, Y1
	VPXOR   32(AX), Y0, Y2
	VPXOR   64(AX), Y0, Y3
	VPXOR   96(AX), Y0, Y4
	VMOVDQU Y1, (AX)
	VMOVDQU Y2, 32(AX)
	VMOVDQU Y3, 64(AX)
	VMOVDQU Y4, 96(AX)
	MOVQ    CX, AX
	JMP     avx2_loop

sse:
	MOVQ       DI, X0
	PUNPCKLQDQ X0, X0
	MOVQ       AX, CX

sse_loop:
	ADDQ  $0x40, CX
	CMPQ  CX, DX
	JAE   x86
	MOVOU 0*16(AX), X1
	MOVOU 1*16(AX), X2
	MOVOU 2*16(AX), X3
	MOVOU 3*16(AX), X4
	PXOR  X0, X1
	PXOR  X0, X2
	PXOR  X0, X3
	PXOR  X0, X4
	MOVOU X1, 0*16(AX)
	MOVOU X2, 1*16(AX)
	MOVOU X3, 2*16(AX)
	MOVOU X4, 3*16(AX)
	MOVQ  CX, AX
	JMP   sse_loop

x86:
	MOVQ AX, CX

x86_loop:
	ADDQ $0x20, CX
	CMPQ CX, DX
	JAE  slow_loop
	XORQ DI, (AX)
	XORQ DI, 8(AX)
	XORQ DI, 16(AX)
	XORQ DI, 24(AX)
	MOVQ CX, AX
	JMP  x86_loop

slow_loop:
	CMPQ AX, DX
	JAE  done
	XORQ DI, (AX)
	ADDQ $0x08, AX
	JMP  slow_loop

done:
	RET
