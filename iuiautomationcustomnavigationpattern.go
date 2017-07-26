package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationCustomNavigationPattern struct {
	ole.IUnknown
}

type IUIAutomationCustomNavigationPatternVtbl struct {
	ole.IUnknownVtbl
	Navigate uintptr
}

func (v *IUIAutomationCustomNavigationPattern) VTable() *IUIAutomationCustomNavigationPatternVtbl {
	return (*IUIAutomationCustomNavigationPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationCustomNavigationPattern) Navigate() error {
	return ole.NewError(ole.E_NOTIMPL)
}
