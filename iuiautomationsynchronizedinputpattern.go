package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationSynchronizedInputPattern struct {
	ole.IUnknown
}

type IUIAutomationSynchronizedInputPatternVtbl struct {
	ole.IUnknownVtbl
	StartListening uintptr
}

func (v *IUIAutomationSynchronizedInputPattern) VTable() *IUIAutomationSynchronizedInputPatternVtbl {
	return (*IUIAutomationSynchronizedInputPatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationSynchronizedInputPattern) StartListening() error {
	return ole.NewError(ole.E_NOTIMPL)
}
