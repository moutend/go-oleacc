package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationProxyFactoryMapping struct {
	ole.IUnknown
}

type IUIAutomationProxyFactoryMappingVtbl struct {
	ole.IUnknownVtbl
	Get_Count     uintptr
	GetTable      uintptr
	GetEntry      uintptr
	SetTable      uintptr
	InsertEntries uintptr
	InsertEntry   uintptr
	RemoveEntry   uintptr
}

func (v *IUIAutomationProxyFactoryMapping) VTable() *IUIAutomationProxyFactoryMappingVtbl {
	return (*IUIAutomationProxyFactoryMappingVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationProxyFactoryMapping) Get_Count() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryMapping) GetTable() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryMapping) GetEntry() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryMapping) SetTable() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryMapping) InsertEntries() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryMapping) InsertEntry() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactoryMapping) RemoveEntry() error {
	return ole.NewError(ole.E_NOTIMPL)
}
