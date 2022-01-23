// Copyright 2022 The RomiChan WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !appengine && (amd64 || arm64)
// +build !appengine
// +build amd64 arm64

package websocket

import (
	"golang.org/x/sys/cpu"
)

func maskBytes(key32 uint32, b []byte) uint32 {
	if len(b) > 0 {
		return maskAsm(&b[0], len(b), key32)
	}
	return key32
}

var useAVX2 = cpu.X86.HasAVX2

//go:noescape
func maskAsm(b *byte, len int, key uint32) uint32
