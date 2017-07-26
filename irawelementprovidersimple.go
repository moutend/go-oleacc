package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IRawElementProviderSimple struct {
	ole.IUnknown
}

type IRawElementProviderSimpleVtbl struct {
	ole.IUnknownVtbl
	Get_ProviderOptions        uintptr
	GetPatternProvider         uintptr
	GetPropertyValue           uintptr
	Get_HostRawElementProvider uintptr
}

func (v *IRawElementProviderSimple) VTable() *IRawElementProviderSimpleVtbl {
	return (*IRawElementProviderSimpleVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IRawElementProviderSimple) Get_ProviderOptions() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IRawElementProviderSimple) GetPatternProvider() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IRawElementProviderSimple) GetPropertyValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IRawElementProviderSimple) Get_HostRawElementProvider() error {
	return ole.NewError(ole.E_NOTIMPL)
}
