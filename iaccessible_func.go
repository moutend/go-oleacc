// +build !windows

package waa

import (
	"github.com/go-ole/go-ole"
)

func accNavigate(acc *IAccessible, navdir uint32, varStart *ole.VARIANT, varEndUpAt **ole.VARIANT) error {
	return ole.NewError(ole.E_NOTIMPL)
}
