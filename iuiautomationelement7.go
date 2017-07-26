package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationElement7 struct {
	IUIAutomationElement6
}

type IUIAutomationElement7Vtbl struct {
	IUIAutomationElement6Vtbl
	FindFirstWithOptions           uintptr
	FindAllWithOptions             uintptr
	FindFirstWithOptionsBuildCache uintptr
	FindAllWithOptionsBuildCache   uintptr
	GetCurrentMetadataValue        uintptr
}

func (v *IUIAutomationElement7) VTable() *IUIAutomationElement7Vtbl {
	return (*IUIAutomationElement7Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationElement7) FindFirstWithOptions() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement7) FindAllWithOptions() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement7) FindFirstWithOptionsBuildCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement7) FindAllWithOptionsBuildCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationElement7) GetCurrentMetadataValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}
