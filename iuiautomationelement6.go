package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationElement6 struct {
	IUIAutomationElement5
}

type IUIAutomationElement6Vtbl struct {
	IUIAutomationElement5Vtbl
	Get_CurrentFullDescription uintptr
	Get_CachedFullDescription  uintptr
}

func (v *IUIAutomationElement6) VTable() *IUIAutomationElement6Vtbl {
	return (*IUIAutomationElement6Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationElement6) Get_CurrentFullDescription() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement6) Get_CachedFullDescription() error {
	return ole.NewError(ole.E_NOTIMPL)
}
