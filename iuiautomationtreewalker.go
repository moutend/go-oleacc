package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationTreeWalker struct {
	ole.IUnknown
}

type IUIAutomationTreeWalkerVtbl struct {
	ole.IUnknownVtbl
	GetParentElement                    uintptr
	GetFirstChildElement                uintptr
	GetLastChildElement                 uintptr
	GetNextSiblingElement               uintptr
	GetPreviousSiblingElement           uintptr
	NormalizeElement                    uintptr
	GetParentElementBuildCache          uintptr
	GetFirstChildElementBuildCache      uintptr
	GetLastChildElementBuildCache       uintptr
	GetNextSiblingElementBuildCache     uintptr
	GetPreviousSiblingElementBuildCache uintptr
	NormalizeElementBuildCache          uintptr
	Get_Condition                       uintptr
}

func (v *IUIAutomationTreeWalker) VTable() *IUIAutomationTreeWalkerVtbl {
	return (*IUIAutomationTreeWalkerVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationTreeWalker) GetParentElement() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTreeWalker) GetFirstChildElement() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTreeWalker) GetLastChildElement() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTreeWalker) GetNextSiblingElement() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTreeWalker) GetPreviousSiblingElement() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTreeWalker) NormalizeElement() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTreeWalker) GetParentElementBuildCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTreeWalker) GetFirstChildElementBuildCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTreeWalker) GetLastChildElementBuildCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTreeWalker) GetNextSiblingElementBuildCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTreeWalker) GetPreviousSiblingElementBuildCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTreeWalker) NormalizeElementBuildCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationTreeWalker) Get_Condition() error {
	return ole.NewError(ole.E_NOTIMPL)
}
