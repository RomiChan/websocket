// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// !appengine

package websocket

import (
	"fmt"
	"math/bits"
	"testing"
)

const wordSize = 8

func maskBytesByByte(key uint32, b []byte) uint32 {
	for i := range b {
		b[i] ^= byte(key)
		key = bits.RotateLeft32(key, -8)
	}
	return key
}

func notzero(b []byte) int {
	for i := range b {
		if b[i] != 0 {
			return i
		}
	}
	return -1
}

func TestMaskBytes(t *testing.T) {
	origKey := newMaskKey()
	for size := 1; size <= 1024; size++ {
		for align := 0; align < wordSize; align++ {
			key := origKey
			for pos := 0; pos < 1; pos++ {
				b := make([]byte, size+align)[align:]
				r1 := maskBytes(key, b)
				r2 := maskBytesByByte(key, b)
				if i := notzero(b); i >= 0 {
					t.Errorf("size:%d, align:%d, pos:%d, offset:%d", size, align, pos, i)
				}
				if r1 != r2 {
					t.Errorf("size:%d orig: %d key got:%d, expected:%d", size, key, r1, r2)
				}
				key = bits.RotateLeft32(key, -8)
			}
		}
	}
}

func BenchmarkMaskBytes(b *testing.B) {
	for _, size := range []int{16, 32, 512, 1024, 4096} {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			for _, align := range []int{0, wordSize / 2} {
				b.Run(fmt.Sprintf("align-%d", align), func(b *testing.B) {
					for _, fn := range []struct {
						name string
						fn   func(key uint32, b []byte) uint32
					}{
						{"go", maskBytesGo},
						{"asm", maskBytes},
					} {
						b.Run(fn.name, func(b *testing.B) {
							key := newMaskKey()
							data := make([]byte, size+align)[align:]
							for i := 0; i < b.N; i++ {
								fn.fn(key, data)
							}
							b.SetBytes(int64(len(data)))
						})
					}
				})
			}
		})
	}
}
