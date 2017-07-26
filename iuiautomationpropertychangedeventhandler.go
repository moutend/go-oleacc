package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationPropertyChangedEventHandler struct {
	ole.IUnknown
}

type IUIAutomationPropertyChangedEventHandlerVtbl struct {
	ole.IUnknownVtbl
	HandlePropertyChangedEvent uintptr
}

func (v *IUIAutomationPropertyChangedEventHandler) VTable() *IUIAutomationPropertyChangedEventHandlerVtbl {
	return (*IUIAutomationPropertyChangedEventHandlerVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationPropertyChangedEventHandler) HandlePropertyChangedEvent() error {
	return ole.NewError(ole.E_NOTIMPL)
}
