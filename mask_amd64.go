// Copyright 2022 The RomiChan WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !appengine || amd64
// +build !appengine amd64

package websocket

import (
	"encoding/binary"

	"golang.org/x/sys/cpu"
)

func maskBytes(key [4]byte, pos int, b []byte) int {
	if len(b) < 128 {
		return maskBytesGo(key, pos, b)
	}

	var k [4]byte
	for i := range k {
		k[i] = key[(pos+i)&3]
	}
	key32 := binary.LittleEndian.Uint32(k[:])
	key64 := uint64(key32) | uint64(key32)<<32

	off := len(b) / 8 * 8
	maskBlockAvx2(b[:off], key64)
	b = b[off:]
	// Mask one byte at a time for remaining bytes.
	for i := range b {
		b[i] ^= key[pos&3]
		pos++
	}
	return pos & 3
}

var useAVX2 = cpu.X86.HasAVX2
var useAVX512 = cpu.X86.HasAVX512F

//go:noescape
func maskBlockAvx2(b []byte, key uint64)
