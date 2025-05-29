//go:build arm64
// +build arm64

/*
 * Copyright 2022 ByteDance Inc.
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

package abi

import (
	"reflect"

	"github.com/douglasboyle/sonic/loader/internal/rt"
)

const (
	PtrSize  = 8 // pointer size on ARM64
	PtrAlign = 8 // pointer alignment on ARM64
)

type floatKind uint8

const (
	notFloatKind floatKind = iota
	floatKind32
	floatKind64
)

type Parameter struct {
	InRegister bool
	IsPointer  bool
	IsFloat    floatKind
	Reg        interface{} // placeholder for register type
	Mem        uint32
	Type       reflect.Type
}

func (self Parameter) String() string {
	if self.InRegister {
		return "[register parameter - ARM64 not implemented]"
	} else {
		return "[stack parameter - ARM64 not implemented]"
	}
}

type FunctionLayout struct {
	FP   uint32
	Args []Parameter
	Rets []Parameter
}

func (self FunctionLayout) ArgSize() uint32 {
	size := uintptr(0)
	for _, arg := range self.Args {
		size += arg.Type.Size()
	}
	return uint32(size)
}

type Frame struct {
	desc   *FunctionLayout
	locals []bool
	ccall  bool
}

// NewFunctionLayout creates a new function layout (ARM64 stub)
func NewFunctionLayout(ft reflect.Type) FunctionLayout {
	// For ARM64, we'll create a basic layout with all parameters on stack
	// This is a stub implementation
	layout := FunctionLayout{
		FP:   0,
		Args: []Parameter{},
		Rets: []Parameter{},
	}

	if ft.Kind() != reflect.Func {
		panic("NewFunctionLayout: expected function type")
	}

	// For now, just create empty parameter lists
	// A proper implementation would analyze the function signature
	return layout
}

// NewFrame creates a new frame (ARM64 stub)
func NewFrame(desc *FunctionLayout, locals []bool, ccall bool) Frame {
	fr := Frame{}
	fr.desc = desc
	fr.locals = locals
	fr.ccall = ccall
	return fr
}

func (self *Frame) Size() uint32 {
	// ARM64 stub - return minimal frame size
	return uint32(PtrSize)
}

func (self *Frame) ArgPtrs() *rt.StackMap {
	// ARM64 stub - return empty stackmap
	var m rt.StackMapBuilder
	return m.Build()
}

func (self *Frame) LocalPtrs() *rt.StackMap {
	// ARM64 stub - return empty stackmap
	var m rt.StackMapBuilder
	return m.Build()
}

func (self *Frame) StackCheckTextSize() uint32 {
	// ARM64 stub - return 0 for now
	return 0
}

func (self *Frame) GrowStackTextSize() uint32 {
	// ARM64 stub - return 0 for now
	return 0
}

// ReservedRegs returns an empty slice for ARM64 (not implemented)
func ReservedRegs(ccall bool) []interface{} {
	return []interface{}{}
}

// Stub implementations for methods that would be called
func (self *Frame) emitStackCheck(p interface{}, stack interface{}, maxStack uintptr) {
	panic("ARM64 implementation not available")
}

func (self *Frame) emitExchangeArgs(p interface{}) {
	panic("ARM64 implementation not available")
}

func (self *Frame) emitExchangeRets(p interface{}) {
	panic("ARM64 implementation not available")
}

func (self *Frame) emitRestoreRegs(p interface{}) {
	panic("ARM64 implementation not available")
}

// CallC stub for ARM64
func CallC(addr uintptr, fr Frame, maxStack uintptr) []byte {
	panic("CallC not implemented for ARM64")
}
