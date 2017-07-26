package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationLegacyIAccessiblePattern struct {
	ole.IUnknown
}

type IUIAutomationLegacyIAccessiblePatternVtbl struct {
	ole.IUnknownVtbl
	Select                      uintptr
	SetValue                    uintptr
	Get_CurrentChildId          uintptr
	Get_CurrentName             uintptr
	Get_CurrentValue            uintptr
	Get_CurrentDescription      uintptr
	Get_CurrentRole             uintptr
	Get_CurrentState            uintptr
	Get_CurrentHelp             uintptr
	Get_CurrentKeyboardShortcut uintptr
	GetCurrentSelection         uintptr
	Get_CurrentDefaultAction    uintptr
	Get_CachedChildId           uintptr
	Get_CachedName              uintptr
	Get_CachedValue             uintptr
	Get_CachedDescription       uintptr
	Get_CachedRole              uintptr
	Get_CachedState             uintptr
	Get_CachedHelp              uintptr
	Get_CachedKeyboardShortcut  uintptr
	GetCachedSelection          uintptr
	Get_CachedDefaultAction     uintptr
	GetIAccessible              uintptr
}

func (v *IUIAutomationLegacyIAccessiblePattern) VTable() *IUIAutomationLegacyIAccessiblePatternVtbl {
	return (*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationLegacyIAccessiblePattern) Select() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) SetValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentChildId() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentDescription() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentRole() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentState() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentHelp() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentKeyboardShortcut() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) GetCurrentSelection() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CurrentDefaultAction() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedChildId() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedDescription() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedRole() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedState() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedHelp() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedKeyboardShortcut() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) GetCachedSelection() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) Get_CachedDefaultAction() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomationLegacyIAccessiblePattern) GetIAccessible() error {
	return ole.NewError(ole.E_NOTIMPL)
}
