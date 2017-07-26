package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationProxyFactoryEntry struct {
	ole.IUnknown
}

type IUIAutomationProxyFactoryEntryVtbl struct {
	ole.IUnknownVtbl
	Get_ProxyFactory               uintptr
	Get_ClassName                  uintptr
	Get_ImageName                  uintptr
	Get_AllowSubstringMatch        uintptr
	Get_CanCheckBaseClass          uintptr
	Get_NeedsAdviseEvents          uintptr
	Put_ClassName                  uintptr
	Put_ImageName                  uintptr
	Put_AllowSubstringMatch        uintptr
	Put_CanCheckBaseClass          uintptr
	Put_NeedsAdviseEvents          uintptr
	SetWinEventsForAutomationEvent uintptr
	GetWinEventsForAutomationEvent uintptr
}

func (v *IUIAutomationProxyFactoryEntry) VTable() *IUIAutomationProxyFactoryEntryVtbl {
	return (*IUIAutomationProxyFactoryEntryVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationProxyFactoryEntry) Get_ProxyFactory() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryEntry) Get_ClassName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryEntry) Get_ImageName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryEntry) Get_AllowSubstringMatch() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryEntry) Get_CanCheckBaseClass() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryEntry) Get_NeedsAdviseEvents() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryEntry) Put_ClassName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryEntry) Put_ImageName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryEntry) Put_AllowSubstringMatch() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryEntry) Put_CanCheckBaseClass() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryEntry) Put_NeedsAdviseEvents() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryEntry) SetWinEventsForAutomationEvent() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryEntry) GetWinEventsForAutomationEvent() error {
	return ole.NewError(ole.E_NOTIMPL)
}
