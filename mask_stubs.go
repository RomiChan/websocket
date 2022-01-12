// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

//go:build appengine || !amd64
// +build appengine !amd64

package websocket

func maskBytes(key [4]byte, pos int, b []byte) int {
	return maskBytesGo(key, pos, b)
}
