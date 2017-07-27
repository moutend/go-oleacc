package oleacc

import (
	"github.com/go-ole/go-ole"
)

func AccessibleObjectFromEvent(hWindow uintptr, objectId, childId uint32, acc **IAccessible, varChild **ole.VARIANT) error {
	return accessibleObjectFromEvent(hWindow, objectId, childId, acc, varChild)
}

func AccessibleObjectFromPoint(point Point, acc **IAccessible, varChild **ole.VARIANT) error {
	return accessibleObjectFromPoint(point, acc, varChild)
}

func OpenInputDesktop(flags uint32, isInherit bool, accessMask uint32) HDESK {
	return openInputDesktop(flags, isInherit, accessMask)
}

func CloseDesktop(hDesk HDESK) bool {
	return closeDesktop(hDesk)
}
