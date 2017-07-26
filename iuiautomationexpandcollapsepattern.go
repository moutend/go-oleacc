package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationExpandCollapsePattern struct {
	ole.IUnknown
}

type IUIAutomationExpandCollapsePatternVtbl struct {
	ole.IUnknownVtbl
	Get_CurrentExpandCollapseState uintptr
	Get_CachedExpandCollapseState  uintptr
}

func (v *IUIAutomationExpandCollapsePattern) VTable() *IUIAutomationExpandCollapsePatternVtbl {
	return (*IUIAutomationExpandCollapsePatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationExpandCollapsePattern) Get_CurrentExpandCollapseState() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationExpandCollapsePattern) Get_CachedExpandCollapseState() error {
	return ole.NewError(ole.E_NOTIMPL)
}
