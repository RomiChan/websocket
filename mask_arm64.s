// Copyright 2022 The RomiChan WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// func maskBlockAsm(b []byte, key uint64)
// Requires: AVX, AVX2
TEXT ·maskAsm(SB), NOSPLIT, $0-32
	MOVD b_ptr+0(FP), R0
	MOVD b_len+8(FP), R1
	MOVD key+16(FP), R2
	VDUP R2, V0.D2
	CMP  $64, R1
	BLT  tail

loop_64:
	VLD1   (R0), [V1.B16, V2.B16, V3.B16, V4.B16]
	VEOR   V1.B16, V0.B16, V1.B16
	VEOR   V2.B16, V0.B16, V2.B16
	VEOR   V3.B16, V0.B16, V3.B16
	VEOR   V4.B16, V0.B16, V4.B16
	VST1.P [V1.B16, V2.B16, V3.B16, V4.B16], 64(R0)
	SUBS   $64, R1
	CMP    $64, R1
	BGE    loop_64

tail:
	// quick end
	CBZ    R1, end
	TBZ    $5, R1, less_than32
	VLD1   (R0), [V1.B16, V2.B16]
	VEOR   V1.B16, V0.B16, V1.B16
	VEOR   V2.B16, V0.B16, V2.B16
	VST1.P [V1.B16, V2.B16], 32(R0)

less_than32:
	TBZ   $4, R1, less_than16
	LDP   (R0), (R11, R12)
	EOR   R11, R2, R11
	EOR   R12, R2, R12
	STP.P (R11, R12), 16(R0)

less_than16:
	TBZ    $3, R1, end
	MOVD   (R0), R11
	EOR    R2, R11, R11
	MOVD   R11, (R0)

end:
	RET
