package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomation2 struct {
	IUIAutomation
}

type IUIAutomation2Vtbl struct {
	IUIAutomationVtbl
	Get_AutoSetFocus       uintptr
	Put_AutoSetFocus       uintptr
	Get_ConnectionTimeout  uintptr
	Put_ConnectionTimeout  uintptr
	Get_TransactionTimeout uintptr
	Put_TransactionTimeout uintptr
}

func (v *IUIAutomation2) VTable() *IUIAutomation2Vtbl {
	return (*IUIAutomation2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomation2) Get_AutoSetFocus() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation2) Put_AutoSetFocus() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation2) Get_ConnectionTimeout() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation2) Put_ConnectionTimeout() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation2) Get_TransactionTimeout() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation2) Put_TransactionTimeout() error {
	return ole.NewError(ole.E_NOTIMPL)
}
