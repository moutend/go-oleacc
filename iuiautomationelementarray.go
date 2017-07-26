package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationElementArray struct {
	ole.IUnknown
}

type IUIAutomationElementArrayVtbl struct {
	ole.IUnknownVtbl
	Get_Length uintptr
	GetElement uintptr
}

func (v *IUIAutomationElementArray) VTable() *IUIAutomationElementArrayVtbl {
	return (*IUIAutomationElementArrayVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationElementArray) Get_Length() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElementArray) GetElement() error {
	return ole.NewError(ole.E_NOTIMPL)
}
