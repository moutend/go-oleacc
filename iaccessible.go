package waa

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAccessible struct {
	ole.IDispatch
}

type IAccessibleVtbl struct {
	ole.IDispatchVtbl
	Get_accParent           uintptr
	Get_accChildCount       uintptr
	Get_accChild            uintptr
	Get_accName             uintptr
	Get_accValue            uintptr
	Get_accDescription      uintptr
	Get_accRole             uintptr
	Get_accState            uintptr
	Get_accHelp             uintptr
	Get_accHelpTopic        uintptr
	Get_accKeyboardShortcut uintptr
	Get_accFocus            uintptr
	Get_accSelection        uintptr
	Get_accDefaultAction    uintptr
	AccSelect               uintptr
	AccLocation             uintptr
	AccNavigate             uintptr
	AccHitTest              uintptr
	AccDoDefaultAction      uintptr
	Put_accName             uintptr
	Put_accValue            uintptr
}

func (v *IAccessible) VTable() *IAccessibleVtbl {
	return (*IAccessibleVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAccessible) Get_accParent() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) Get_accChildCount() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) Get_accChild() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) Get_accName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) Get_accValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) Get_accDescription() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) Get_accRole() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) Get_accState() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) Get_accHelp() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) Get_accHelpTopic() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) Get_accKeyboardShortcut() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) Get_accFocus() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) Get_accSelection() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) Get_accDefaultAction() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) AccSelect() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) AccLocation() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) AccNavigate(navdir uint32, varStart *ole.VARIANT, varEndUpAt **ole.VARIANT) error {
	return accNavigate(v, navdir, varStart, varEndUpAt)
}

func (v *IAccessible) AccHitTest() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) AccDoDefaultAction() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) Put_accName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAccessible) Put_accValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}
