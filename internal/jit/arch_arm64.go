//go:build arm64
// +build arm64

/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package jit

import (
	"unsafe"

	"github.com/twitchyliquid64/golang-asm/obj"
)

// Stub implementation for ARM64 - not implemented
func As(op string) obj.As {
	panic("jit: ARM64 implementation not available")
}

func ImmPtr(imm unsafe.Pointer) obj.Addr {
	panic("jit: ARM64 implementation not available")
}

func Imm(imm int64) obj.Addr {
	panic("jit: ARM64 implementation not available")
}

func Reg(reg string) obj.Addr {
	panic("jit: ARM64 implementation not available")
}

func Ptr(reg obj.Addr, offs int64) obj.Addr {
	panic("jit: ARM64 implementation not available")
}

func Sib(reg obj.Addr, idx obj.Addr, scale int16, offs int64) obj.Addr {
	panic("jit: ARM64 implementation not available")
}
