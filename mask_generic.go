// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

//go:build appengine || (!amd64 && !arm64)
// +build appengine !amd64,!arm64

package websocket

func maskBytes(key32 uint32, b []byte) uint32 {
	return maskBytesGo(key32, b)
}
