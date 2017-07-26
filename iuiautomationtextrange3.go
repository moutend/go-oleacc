package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationTextRange3 struct {
	IUIAutomationTextRange2
}

type IUIAutomationTextRange3Vtbl struct {
	IUIAutomationTextRange2Vtbl
	GetEnclosingElementBuildCache uintptr
	GetChildrenBuildCache         uintptr
	GetAttributeValues            uintptr
}

func (v *IUIAutomationTextRange3) VTable() *IUIAutomationTextRange3Vtbl {
	return (*IUIAutomationTextRange3Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationTextRange3) GetEnclosingElementBuildCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextRange3) GetChildrenBuildCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTextRange3) GetAttributeValues() error {
	return ole.NewError(ole.E_NOTIMPL)
}
