package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationChangesEventHandler struct {
	ole.IUnknown
}

type IUIAutomationChangesEventHandlerVtbl struct {
	ole.IUnknownVtbl
	HandleChangesEvent uintptr
}

func (v *IUIAutomationChangesEventHandler) VTable() *IUIAutomationChangesEventHandlerVtbl {
	return (*IUIAutomationChangesEventHandlerVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationChangesEventHandler) HandleChangesEvent() error {
	return ole.NewError(ole.E_NOTIMPL)
}
