package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationEventHandler struct {
	ole.IUnknown
}

type IUIAutomationEventHandlerVtbl struct {
	ole.IUnknownVtbl
	HandleAutomationEvent uintptr
}

func (v *IUIAutomationEventHandler) VTable() *IUIAutomationEventHandlerVtbl {
	return (*IUIAutomationEventHandlerVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationEventHandler) HandleAutomationEvent() error {
	return ole.NewError(ole.E_NOTIMPL)
}
