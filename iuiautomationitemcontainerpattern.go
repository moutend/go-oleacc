package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationItemContainerPattern struct {
	ole.IUnknown
}

type IUIAutomationItemContainerPatternVtbl struct {
	ole.IUnknownVtbl
	FindItemByProperty uintptr
}

func (v *IUIAutomationItemContainerPattern) VTable() *IUIAutomationItemContainerPatternVtbl {
	return (*IUIAutomationItemContainerPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationItemContainerPattern) FindItemByProperty() error {
	return ole.NewError(ole.E_NOTIMPL)
}
