/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.0
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: common.i

package testgoupm

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


extern void _wrap_Swig_free_testgoupm_5354be522b10cd1b(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_testgoupm_5354be522b10cd1b(swig_intgo arg1);
extern uintptr_t _wrap_new_TESTIT_testgoupm_5354be522b10cd1b(swig_intgo arg1);
extern swig_intgo _wrap_TESTIT_blibbity_testgoupm_5354be522b10cd1b(uintptr_t arg1);
extern void _wrap_delete_TESTIT_testgoupm_5354be522b10cd1b(uintptr_t arg1);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_testgoupm_5354be522b10cd1b(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_testgoupm_5354be522b10cd1b(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrTESTIT uintptr

func (p SwigcptrTESTIT) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrTESTIT) SwigIsTESTIT() {
}

func NewTESTIT(arg1 int) (_swig_ret TESTIT) {
	var swig_r TESTIT
	_swig_i_0 := arg1
	swig_r = (TESTIT)(SwigcptrTESTIT(C._wrap_new_TESTIT_testgoupm_5354be522b10cd1b(C.swig_intgo(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrTESTIT) Blibbity() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_TESTIT_blibbity_testgoupm_5354be522b10cd1b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func DeleteTESTIT(arg1 TESTIT) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_TESTIT_testgoupm_5354be522b10cd1b(C.uintptr_t(_swig_i_0))
}

type TESTIT interface {
	Swigcptr() uintptr
	SwigIsTESTIT()
	Blibbity() (_swig_ret int)
}


