package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationTextEditTextChangedEventHandler struct {
	ole.IUnknown
}

type IUIAutomationTextEditTextChangedEventHandlerVtbl struct {
	ole.IUnknownVtbl
	HandleTextEditTextChangedEvent uintptr
}

func (v *IUIAutomationTextEditTextChangedEventHandler) VTable() *IUIAutomationTextEditTextChangedEventHandlerVtbl {
	return (*IUIAutomationTextEditTextChangedEventHandlerVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationTextEditTextChangedEventHandler) HandleTextEditTextChangedEvent() error {
	return ole.NewError(ole.E_NOTIMPL)
}
