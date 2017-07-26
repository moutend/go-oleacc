package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationProxyFactory struct {
	ole.IUnknown
}

type IUIAutomationProxyFactoryVtbl struct {
	ole.IUnknownVtbl
	CreateProvider     uintptr
	Get_ProxyFactoryId uintptr
}

func (v *IUIAutomationProxyFactory) VTable() *IUIAutomationProxyFactoryVtbl {
	return (*IUIAutomationProxyFactoryVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationProxyFactory) CreateProvider() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationProxyFactory) Get_ProxyFactoryId() error {
	return ole.NewError(ole.E_NOTIMPL)
}
