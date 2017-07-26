package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationDockPattern struct {
	ole.IUnknown
}

type IUIAutomationDockPatternVtbl struct {
	ole.IUnknownVtbl
	SetDockPosition         uintptr
	Get_CurrentDockPosition uintptr
	Get_CachedDockPosition  uintptr
}

func (v *IUIAutomationDockPattern) VTable() *IUIAutomationDockPatternVtbl {
	return (*IUIAutomationDockPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationDockPattern) SetDockPosition() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationDockPattern) Get_CurrentDockPosition() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationDockPattern) Get_CachedDockPosition() error {
	return ole.NewError(ole.E_NOTIMPL)
}
