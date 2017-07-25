package oleacc

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomation struct {
	ole.IUnknown
}

type IUIAutomationVtbl struct {
	ole.IUnknownVtbl
	CompareElements                           uintptr
	CompareRuntimeIds                         uintptr
	GetRootElement                            uintptr
	ElementFromHandle                         uintptr
	ElementFromPoint                          uintptr
	GetFocusedElement                         uintptr
	GetRootElementBuildCache                  uintptr
	ElementFromHandleBuildCache               uintptr
	ElementFromPointBuildCache                uintptr
	GetFocusedElementBuildCache               uintptr
	CreateTreeWalker                          uintptr
	Get_ControlViewWalker                     uintptr
	Get_ContentViewWalker                     uintptr
	Get_RawViewWalker                         uintptr
	Get_RawViewCondition                      uintptr
	Get_ControlViewCondition                  uintptr
	Get_ContentViewCondition                  uintptr
	CreateCacheRequest                        uintptr
	CreateTrueCondition                       uintptr
	CreateFalseCondition                      uintptr
	CreatePropertyCondition                   uintptr
	CreatePropertyConditionEx                 uintptr
	CreateAndCondition                        uintptr
	CreateAndConditionFromArray               uintptr
	CreateAndConditionFromNativeArray         uintptr
	CreateOrCondition                         uintptr
	CreateOrConditionFromArray                uintptr
	CreateOrConditionFromNativeArray          uintptr
	CreateNotCondition                        uintptr
	AddAutomationEventHandler                 uintptr
	RemoveAutomationEventHandler              uintptr
	AddPropertyChangedEventHandlerNativeArray uintptr
	AddPropertyChangedEventHandler            uintptr
	RemovePropertyChangedEventHandler         uintptr
	AddStructureChangedEventHandler           uintptr
	RemoveStructureChangedEventHandler        uintptr
	AddFocusChangedEventHandler               uintptr
	RemoveFocusChangedEventHandler            uintptr
	IntNativeArrayToSafeArray                 uintptr
	IntSafeArrayToNativeArray                 uintptr
	RectToVariant                             uintptr
	VariantToRect                             uintptr
	SafeArrayToRectNativeArray                uintptr
	CreateProxyFactoryEntry                   uintptr
	Get_ProxyFactoryMapping                   uintptr
	GetPropertyProgrammaticName               uintptr
	GetPatternProgrammaticName                uintptr
	PollForPotentialSupportedPatterns         uintptr
	PollForPotentialSupportedProperties       uintptr
	CheckNotSupported                         uintptr
	Get_ReservedNotSupportedValue             uintptr
	Get_ReservedMixedAttributeValue           uintptr
	ElementFromIAccessible                    uintptr
	ElementFromIAccessibleBuildCache          uintptr
}

func (v *IUIAutomation) VTable() *IUIAutomationVtbl {
	return (*IUIAutomationVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomation) CompareElements() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) CompareRuntimeIds() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) GetRootElement() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) ElementFromHandle() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) ElementFromPoint() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) GetFocusedElement() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) GetRootElementBuildCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) ElementFromHandleBuildCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) ElementFromPointBuildCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) GetFocusedElementBuildCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) CreateTreeWalker() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) Get_ControlViewWalker() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) Get_ContentViewWalker() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) Get_RawViewWalker() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) Get_RawViewCondition() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) Get_ControlViewCondition() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) Get_ContentViewCondition() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) CreateCacheRequest() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) CreateTrueCondition() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) CreateFalseCondition() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) CreatePropertyCondition() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) CreatePropertyConditionEx() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) CreateAndCondition() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) CreateAndConditionFromArray() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) CreateAndConditionFromNativeArray() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) CreateOrCondition() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) CreateOrConditionFromArray() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) CreateOrConditionFromNativeArray() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) CreateNotCondition() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) AddAutomationEventHandler() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) RemoveAutomationEventHandler() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) AddPropertyChangedEventHandlerNativeArray() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) AddPropertyChangedEventHandler() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) RemovePropertyChangedEventHandler() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) AddStructureChangedEventHandler() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) RemoveStructureChangedEventHandler() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) AddFocusChangedEventHandler() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) RemoveFocusChangedEventHandler() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) IntNativeArrayToSafeArray() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) IntSafeArrayToNativeArray() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) RectToVariant() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) VariantToRect() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) SafeArrayToRectNativeArray() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) CreateProxyFactoryEntry() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) Get_ProxyFactoryMapping() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) GetPropertyProgrammaticName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) GetPatternProgrammaticName() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) PollForPotentialSupportedPatterns() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) PollForPotentialSupportedProperties() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) CheckNotSupported() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) Get_ReservedNotSupportedValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) Get_ReservedMixedAttributeValue() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) ElementFromIAccessible() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IUIAutomation) ElementFromIAccessibleBuildCache() error {
	return ole.NewError(ole.E_NOTIMPL)
}
