// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

package websocket

import (
	"encoding/binary"
	"math/bits"
)

func maskBytesGo(key32 uint32, b []byte) uint32 {
	if len(b) >= 16 {
		key64 := uint64(key32) | uint64(key32)<<32
		// At some point in the future we can clean these unrolled loops up.
		// See https://github.com/golang/go/issues/31586#issuecomment-487436401

		// Then we xor until b is less than 128 bytes.
		for len(b) >= 128 {
			v := binary.LittleEndian.Uint64(b)
			binary.LittleEndian.PutUint64(b, v^key64)
			v = binary.LittleEndian.Uint64(b[8:16])
			binary.LittleEndian.PutUint64(b[8:16], v^key64)
			v = binary.LittleEndian.Uint64(b[16:24])
			binary.LittleEndian.PutUint64(b[16:24], v^key64)
			v = binary.LittleEndian.Uint64(b[24:32])
			binary.LittleEndian.PutUint64(b[24:32], v^key64)
			v = binary.LittleEndian.Uint64(b[32:40])
			binary.LittleEndian.PutUint64(b[32:40], v^key64)
			v = binary.LittleEndian.Uint64(b[40:48])
			binary.LittleEndian.PutUint64(b[40:48], v^key64)
			v = binary.LittleEndian.Uint64(b[48:56])
			binary.LittleEndian.PutUint64(b[48:56], v^key64)
			v = binary.LittleEndian.Uint64(b[56:64])
			binary.LittleEndian.PutUint64(b[56:64], v^key64)
			v = binary.LittleEndian.Uint64(b[64:72])
			binary.LittleEndian.PutUint64(b[64:72], v^key64)
			v = binary.LittleEndian.Uint64(b[72:80])
			binary.LittleEndian.PutUint64(b[72:80], v^key64)
			v = binary.LittleEndian.Uint64(b[80:88])
			binary.LittleEndian.PutUint64(b[80:88], v^key64)
			v = binary.LittleEndian.Uint64(b[88:96])
			binary.LittleEndian.PutUint64(b[88:96], v^key64)
			v = binary.LittleEndian.Uint64(b[96:104])
			binary.LittleEndian.PutUint64(b[96:104], v^key64)
			v = binary.LittleEndian.Uint64(b[104:112])
			binary.LittleEndian.PutUint64(b[104:112], v^key64)
			v = binary.LittleEndian.Uint64(b[112:120])
			binary.LittleEndian.PutUint64(b[112:120], v^key64)
			v = binary.LittleEndian.Uint64(b[120:128])
			binary.LittleEndian.PutUint64(b[120:128], v^key64)
			b = b[128:]
		}

		// Then we xor until b is less than 64 bytes.
		for len(b) >= 64 {
			v := binary.LittleEndian.Uint64(b)
			binary.LittleEndian.PutUint64(b, v^key64)
			v = binary.LittleEndian.Uint64(b[8:16])
			binary.LittleEndian.PutUint64(b[8:16], v^key64)
			v = binary.LittleEndian.Uint64(b[16:24])
			binary.LittleEndian.PutUint64(b[16:24], v^key64)
			v = binary.LittleEndian.Uint64(b[24:32])
			binary.LittleEndian.PutUint64(b[24:32], v^key64)
			v = binary.LittleEndian.Uint64(b[32:40])
			binary.LittleEndian.PutUint64(b[32:40], v^key64)
			v = binary.LittleEndian.Uint64(b[40:48])
			binary.LittleEndian.PutUint64(b[40:48], v^key64)
			v = binary.LittleEndian.Uint64(b[48:56])
			binary.LittleEndian.PutUint64(b[48:56], v^key64)
			v = binary.LittleEndian.Uint64(b[56:64])
			binary.LittleEndian.PutUint64(b[56:64], v^key64)
			b = b[64:]
		}

		// Then we xor until b is less than 32 bytes.
		for len(b) >= 32 {
			v := binary.LittleEndian.Uint64(b)
			binary.LittleEndian.PutUint64(b, v^key64)
			v = binary.LittleEndian.Uint64(b[8:16])
			binary.LittleEndian.PutUint64(b[8:16], v^key64)
			v = binary.LittleEndian.Uint64(b[16:24])
			binary.LittleEndian.PutUint64(b[16:24], v^key64)
			v = binary.LittleEndian.Uint64(b[24:32])
			binary.LittleEndian.PutUint64(b[24:32], v^key64)
			b = b[32:]
		}

		// Then we xor until b is less than 16 bytes.
		for len(b) >= 16 {
			v := binary.LittleEndian.Uint64(b)
			binary.LittleEndian.PutUint64(b, v^key64)
			v = binary.LittleEndian.Uint64(b[8:16])
			binary.LittleEndian.PutUint64(b[8:16], v^key64)
			b = b[16:]
		}

		// Then we xor until b is less than 8 bytes.
		for len(b) >= 8 {
			v := binary.LittleEndian.Uint64(b)
			binary.LittleEndian.PutUint64(b, v^key64)
			b = b[8:]
		}
	}

	for i := range b {
		b[i] ^= byte(key32)
		key32 = bits.RotateLeft32(key32, -8)
	}
	return key32
}
