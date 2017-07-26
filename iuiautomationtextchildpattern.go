package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationTextChildPattern struct {
	ole.IUnknown
}

type IUIAutomationTextChildPatternVtbl struct {
	ole.IUnknownVtbl
	Get_TextContainer uintptr
	Get_TextRange     uintptr
}

func (v *IUIAutomationTextChildPattern) VTable() *IUIAutomationTextChildPatternVtbl {
	return (*IUIAutomationTextChildPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationTextChildPattern) Get_TextContainer() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextChildPattern) Get_TextRange() error {
	return ole.NewError(ole.E_NOTIMPL)
}
