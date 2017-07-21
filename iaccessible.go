package waa

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAccessible struct {
	ole.IDispatch
}

type IAccessibleVtbl struct {
	ole.IDispatchBtbl
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
	return nil
}

func (v *IAccessible) Get_accChildCount() error {
	return nil
}

func (v *IAccessible) Get_accChild() error {
	return nil
}

func (v *IAccessible) Get_accName() error {
	return nil
}

func (v *IAccessible) Get_accValue() error {
	return nil
}

func (v *IAccessible) Get_accDescription() error {
	return nil
}

func (v *IAccessible) Get_accRole() error {
	return nil
}

func (v *IAccessible) Get_accState() error {
	return nil
}

func (v *IAccessible) Get_accHelp() error {
	return nil
}

func (v *IAccessible) Get_accHelpTopic() error {
	return nil
}

func (v *IAccessible) Get_accKeyboardShortcut() error {
	return nil
}

func (v *IAccessible) Get_accFocus() error {
	return nil
}

func (v *IAccessible) Get_accSelection() error {
	return nil
}

func (v *IAccessible) Get_accDefaultAction() error {
	return nil
}

func (v *IAccessible) AccSelect() error {
	return nil
}

func (v *IAccessible) AccLocation() error {
	return nil
}

func (v *IAccessible) AccNavigate() error {
	return nil
}

func (v *IAccessible) AccHitTest() error {
	return nil
}

func (v *IAccessible) AccDoDefaultAction() error {
	return nil
}

func (v *IAccessible) Put_accName() error {
	return nil
}

func (v *IAccessible) Put_accValue() error {
	return nil
}
