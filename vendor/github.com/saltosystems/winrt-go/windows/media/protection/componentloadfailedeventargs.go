// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package protection

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const SignatureComponentLoadFailedEventArgs string = "rc(Windows.Media.Protection.ComponentLoadFailedEventArgs;{95972e93-7746-417e-8495-f031bbc5862c})"

type ComponentLoadFailedEventArgs struct {
	ole.IUnknown
}

func (impl *ComponentLoadFailedEventArgs) GetInformation() (*RevocationAndRenewalInformation, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiComponentLoadFailedEventArgs))
	defer itf.Release()
	v := (*iComponentLoadFailedEventArgs)(unsafe.Pointer(itf))
	return v.GetInformation()
}

func (impl *ComponentLoadFailedEventArgs) GetCompletion() (*MediaProtectionServiceCompletion, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiComponentLoadFailedEventArgs))
	defer itf.Release()
	v := (*iComponentLoadFailedEventArgs)(unsafe.Pointer(itf))
	return v.GetCompletion()
}

const GUIDiComponentLoadFailedEventArgs string = "95972e93-7746-417e-8495-f031bbc5862c"
const SignatureiComponentLoadFailedEventArgs string = "{95972e93-7746-417e-8495-f031bbc5862c}"

type iComponentLoadFailedEventArgs struct {
	ole.IInspectable
}

type iComponentLoadFailedEventArgsVtbl struct {
	ole.IInspectableVtbl

	GetInformation uintptr
	GetCompletion  uintptr
}

func (v *iComponentLoadFailedEventArgs) VTable() *iComponentLoadFailedEventArgsVtbl {
	return (*iComponentLoadFailedEventArgsVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iComponentLoadFailedEventArgs) GetInformation() (*RevocationAndRenewalInformation, error) {
	var out *RevocationAndRenewalInformation
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetInformation,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out RevocationAndRenewalInformation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iComponentLoadFailedEventArgs) GetCompletion() (*MediaProtectionServiceCompletion, error) {
	var out *MediaProtectionServiceCompletion
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetCompletion,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out MediaProtectionServiceCompletion
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
