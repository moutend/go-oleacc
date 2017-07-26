package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationFocusChangedEventHandler struct {
	ole.IUnknown
}

type IUIAutomationFocusChangedEventHandlerVtbl struct {
	ole.IUnknownVtbl
	HandleFocusChangedEvent uintptr
}

func (v *IUIAutomationFocusChangedEventHandler) VTable() *IUIAutomationFocusChangedEventHandlerVtbl {
	return (*IUIAutomationFocusChangedEventHandlerVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationFocusChangedEventHandler) HandleFocusChangedEvent() error {
	return ole.NewError(ole.E_NOTIMPL)
}
