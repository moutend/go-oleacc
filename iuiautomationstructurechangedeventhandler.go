package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationStructureChangedEventHandler struct {
	ole.IUnknown
}

type IUIAutomationStructureChangedEventHandlerVtbl struct {
	ole.IUnknownVtbl
	HandleStructureChangedEvent uintptr
}

func (v *IUIAutomationStructureChangedEventHandler) VTable() *IUIAutomationStructureChangedEventHandlerVtbl {
	return (*IUIAutomationStructureChangedEventHandlerVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationStructureChangedEventHandler) HandleStructureChangedEvent() error {
	return ole.NewError(ole.E_NOTIMPL)
}
