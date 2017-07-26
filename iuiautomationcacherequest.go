package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationCacheRequest struct {
	ole.IUnknown
}

type IUIAutomationCacheRequestVtbl struct {
	ole.IUnknownVtbl
	AddProperty               uintptr
	AddPattern                uintptr
	Clone                     uintptr
	Get_TreeScope             uintptr
	Put_TreeScope             uintptr
	Get_TreeFilter            uintptr
	Put_TreeFilter            uintptr
	Get_AutomationElementMode uintptr
	Put_AutomationElementMode uintptr
}

func (v *IUIAutomationCacheRequest) VTable() *IUIAutomationCacheRequestVtbl {
	return (*IUIAutomationCacheRequestVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationCacheRequest) AddProperty() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationCacheRequest) AddPattern() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationCacheRequest) Clone() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationCacheRequest) Get_TreeScope() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationCacheRequest) Put_TreeScope() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationCacheRequest) Get_TreeFilter() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationCacheRequest) Put_TreeFilter() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationCacheRequest) Get_AutomationElementMode() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationCacheRequest) Put_AutomationElementMode() error {
	return ole.NewError(ole.E_NOTIMPL)
}
