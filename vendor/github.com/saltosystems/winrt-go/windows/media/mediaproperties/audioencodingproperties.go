// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package mediaproperties

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const SignatureAudioEncodingProperties string = "rc(Windows.Media.MediaProperties.AudioEncodingProperties;{62bc7a16-005c-4b3b-8a0b-0a090e9687f3})"

type AudioEncodingProperties struct {
	ole.IUnknown
}

func NewAudioEncodingProperties() (*AudioEncodingProperties, error) {
	inspectable, err := ole.RoActivateInstance("Windows.Media.MediaProperties.AudioEncodingProperties")
	if err != nil {
		return nil, err
	}
	return (*AudioEncodingProperties)(unsafe.Pointer(inspectable)), nil
}

func (impl *AudioEncodingProperties) SetBitrate(value uint32) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiAudioEncodingProperties))
	defer itf.Release()
	v := (*iAudioEncodingProperties)(unsafe.Pointer(itf))
	return v.SetBitrate(value)
}

func (impl *AudioEncodingProperties) GetBitrate() (uint32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiAudioEncodingProperties))
	defer itf.Release()
	v := (*iAudioEncodingProperties)(unsafe.Pointer(itf))
	return v.GetBitrate()
}

func (impl *AudioEncodingProperties) SetChannelCount(value uint32) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiAudioEncodingProperties))
	defer itf.Release()
	v := (*iAudioEncodingProperties)(unsafe.Pointer(itf))
	return v.SetChannelCount(value)
}

func (impl *AudioEncodingProperties) GetChannelCount() (uint32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiAudioEncodingProperties))
	defer itf.Release()
	v := (*iAudioEncodingProperties)(unsafe.Pointer(itf))
	return v.GetChannelCount()
}

func (impl *AudioEncodingProperties) SetSampleRate(value uint32) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiAudioEncodingProperties))
	defer itf.Release()
	v := (*iAudioEncodingProperties)(unsafe.Pointer(itf))
	return v.SetSampleRate(value)
}

func (impl *AudioEncodingProperties) GetSampleRate() (uint32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiAudioEncodingProperties))
	defer itf.Release()
	v := (*iAudioEncodingProperties)(unsafe.Pointer(itf))
	return v.GetSampleRate()
}

func (impl *AudioEncodingProperties) SetBitsPerSample(value uint32) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiAudioEncodingProperties))
	defer itf.Release()
	v := (*iAudioEncodingProperties)(unsafe.Pointer(itf))
	return v.SetBitsPerSample(value)
}

func (impl *AudioEncodingProperties) GetBitsPerSample() (uint32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiAudioEncodingProperties))
	defer itf.Release()
	v := (*iAudioEncodingProperties)(unsafe.Pointer(itf))
	return v.GetBitsPerSample()
}

func (impl *AudioEncodingProperties) GetProperties() (*MediaPropertySet, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDIMediaEncodingProperties))
	defer itf.Release()
	v := (*IMediaEncodingProperties)(unsafe.Pointer(itf))
	return v.GetProperties()
}

func (impl *AudioEncodingProperties) GetType() (string, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDIMediaEncodingProperties))
	defer itf.Release()
	v := (*IMediaEncodingProperties)(unsafe.Pointer(itf))
	return v.GetType()
}

func (impl *AudioEncodingProperties) SetSubtype(value string) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDIMediaEncodingProperties))
	defer itf.Release()
	v := (*IMediaEncodingProperties)(unsafe.Pointer(itf))
	return v.SetSubtype(value)
}

func (impl *AudioEncodingProperties) GetSubtype() (string, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDIMediaEncodingProperties))
	defer itf.Release()
	v := (*IMediaEncodingProperties)(unsafe.Pointer(itf))
	return v.GetSubtype()
}

func (impl *AudioEncodingProperties) SetFormatUserData(valueSize uint32, value []uint8) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiAudioEncodingPropertiesWithFormatUserData))
	defer itf.Release()
	v := (*iAudioEncodingPropertiesWithFormatUserData)(unsafe.Pointer(itf))
	return v.SetFormatUserData(valueSize, value)
}

func (impl *AudioEncodingProperties) GetFormatUserData() (uint32, []uint8, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiAudioEncodingPropertiesWithFormatUserData))
	defer itf.Release()
	v := (*iAudioEncodingPropertiesWithFormatUserData)(unsafe.Pointer(itf))
	return v.GetFormatUserData()
}

func (impl *AudioEncodingProperties) GetIsSpatial() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiAudioEncodingProperties2))
	defer itf.Release()
	v := (*iAudioEncodingProperties2)(unsafe.Pointer(itf))
	return v.GetIsSpatial()
}

func (impl *AudioEncodingProperties) Copy() (*AudioEncodingProperties, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiAudioEncodingProperties3))
	defer itf.Release()
	v := (*iAudioEncodingProperties3)(unsafe.Pointer(itf))
	return v.Copy()
}

const GUIDiAudioEncodingProperties string = "62bc7a16-005c-4b3b-8a0b-0a090e9687f3"
const SignatureiAudioEncodingProperties string = "{62bc7a16-005c-4b3b-8a0b-0a090e9687f3}"

type iAudioEncodingProperties struct {
	ole.IInspectable
}

type iAudioEncodingPropertiesVtbl struct {
	ole.IInspectableVtbl

	SetBitrate       uintptr
	GetBitrate       uintptr
	SetChannelCount  uintptr
	GetChannelCount  uintptr
	SetSampleRate    uintptr
	GetSampleRate    uintptr
	SetBitsPerSample uintptr
	GetBitsPerSample uintptr
}

func (v *iAudioEncodingProperties) VTable() *iAudioEncodingPropertiesVtbl {
	return (*iAudioEncodingPropertiesVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iAudioEncodingProperties) SetBitrate(value uint32) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetBitrate,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(value),             // in uint32
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iAudioEncodingProperties) GetBitrate() (uint32, error) {
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetBitrate,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iAudioEncodingProperties) SetChannelCount(value uint32) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetChannelCount,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(value),             // in uint32
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iAudioEncodingProperties) GetChannelCount() (uint32, error) {
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetChannelCount,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iAudioEncodingProperties) SetSampleRate(value uint32) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetSampleRate,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(value),             // in uint32
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iAudioEncodingProperties) GetSampleRate() (uint32, error) {
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetSampleRate,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iAudioEncodingProperties) SetBitsPerSample(value uint32) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetBitsPerSample,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(value),             // in uint32
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iAudioEncodingProperties) GetBitsPerSample() (uint32, error) {
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetBitsPerSample,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

const GUIDiAudioEncodingPropertiesWithFormatUserData string = "98f10d79-13ea-49ff-be70-2673db69702c"
const SignatureiAudioEncodingPropertiesWithFormatUserData string = "{98f10d79-13ea-49ff-be70-2673db69702c}"

type iAudioEncodingPropertiesWithFormatUserData struct {
	ole.IInspectable
}

type iAudioEncodingPropertiesWithFormatUserDataVtbl struct {
	ole.IInspectableVtbl

	SetFormatUserData uintptr
	GetFormatUserData uintptr
}

func (v *iAudioEncodingPropertiesWithFormatUserData) VTable() *iAudioEncodingPropertiesWithFormatUserDataVtbl {
	return (*iAudioEncodingPropertiesWithFormatUserDataVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iAudioEncodingPropertiesWithFormatUserData) SetFormatUserData(valueSize uint32, value []uint8) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetFormatUserData,
		uintptr(unsafe.Pointer(v)),         // this
		uintptr(valueSize),                 // in uint32
		uintptr(unsafe.Pointer(&value[0])), // in uint8
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iAudioEncodingPropertiesWithFormatUserData) GetFormatUserData() (uint32, []uint8, error) {
	var valueSize uint32
	var value []uint8 = make([]uint8, valueSize)
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetFormatUserData,
		uintptr(unsafe.Pointer(v)),          // this
		uintptr(unsafe.Pointer(&valueSize)), // out uint32
		uintptr(unsafe.Pointer(&value[0])),  // out uint8
	)

	if hr != 0 {
		return 0, nil, ole.NewError(hr)
	}

	return valueSize, value, nil
}

const GUIDiAudioEncodingProperties2 string = "c45d54da-80bd-4c23-80d5-72d4a181e894"
const SignatureiAudioEncodingProperties2 string = "{c45d54da-80bd-4c23-80d5-72d4a181e894}"

type iAudioEncodingProperties2 struct {
	ole.IInspectable
}

type iAudioEncodingProperties2Vtbl struct {
	ole.IInspectableVtbl

	GetIsSpatial uintptr
}

func (v *iAudioEncodingProperties2) VTable() *iAudioEncodingProperties2Vtbl {
	return (*iAudioEncodingProperties2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iAudioEncodingProperties2) GetIsSpatial() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsSpatial,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

const GUIDiAudioEncodingProperties3 string = "87600341-748c-4f8d-b0fd-10caf08ff087"
const SignatureiAudioEncodingProperties3 string = "{87600341-748c-4f8d-b0fd-10caf08ff087}"

type iAudioEncodingProperties3 struct {
	ole.IInspectable
}

type iAudioEncodingProperties3Vtbl struct {
	ole.IInspectableVtbl

	Copy uintptr
}

func (v *iAudioEncodingProperties3) VTable() *iAudioEncodingProperties3Vtbl {
	return (*iAudioEncodingProperties3Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iAudioEncodingProperties3) Copy() (*AudioEncodingProperties, error) {
	var out *AudioEncodingProperties
	hr, _, _ := syscall.SyscallN(
		v.VTable().Copy,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out AudioEncodingProperties
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

const GUIDiAudioEncodingPropertiesStatics string = "0cad332c-ebe9-4527-b36d-e42a13cf38db"
const SignatureiAudioEncodingPropertiesStatics string = "{0cad332c-ebe9-4527-b36d-e42a13cf38db}"

type iAudioEncodingPropertiesStatics struct {
	ole.IInspectable
}

type iAudioEncodingPropertiesStaticsVtbl struct {
	ole.IInspectableVtbl

	AudioEncodingPropertiesCreateAac     uintptr
	AudioEncodingPropertiesCreateAacAdts uintptr
	AudioEncodingPropertiesCreateMp3     uintptr
	AudioEncodingPropertiesCreatePcm     uintptr
	AudioEncodingPropertiesCreateWma     uintptr
}

func (v *iAudioEncodingPropertiesStatics) VTable() *iAudioEncodingPropertiesStaticsVtbl {
	return (*iAudioEncodingPropertiesStaticsVtbl)(unsafe.Pointer(v.RawVTable))
}

func AudioEncodingPropertiesCreateAac(sampleRate uint32, channelCount uint32, bitrate uint32) (*AudioEncodingProperties, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Media.MediaProperties.AudioEncodingProperties", ole.NewGUID(GUIDiAudioEncodingPropertiesStatics))
	if err != nil {
		return nil, err
	}
	v := (*iAudioEncodingPropertiesStatics)(unsafe.Pointer(inspectable))

	var out *AudioEncodingProperties
	hr, _, _ := syscall.SyscallN(
		v.VTable().AudioEncodingPropertiesCreateAac,
		0,                             // this is a static func, so there's no this
		uintptr(sampleRate),           // in uint32
		uintptr(channelCount),         // in uint32
		uintptr(bitrate),              // in uint32
		uintptr(unsafe.Pointer(&out)), // out AudioEncodingProperties
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func AudioEncodingPropertiesCreateAacAdts(sampleRate uint32, channelCount uint32, bitrate uint32) (*AudioEncodingProperties, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Media.MediaProperties.AudioEncodingProperties", ole.NewGUID(GUIDiAudioEncodingPropertiesStatics))
	if err != nil {
		return nil, err
	}
	v := (*iAudioEncodingPropertiesStatics)(unsafe.Pointer(inspectable))

	var out *AudioEncodingProperties
	hr, _, _ := syscall.SyscallN(
		v.VTable().AudioEncodingPropertiesCreateAacAdts,
		0,                             // this is a static func, so there's no this
		uintptr(sampleRate),           // in uint32
		uintptr(channelCount),         // in uint32
		uintptr(bitrate),              // in uint32
		uintptr(unsafe.Pointer(&out)), // out AudioEncodingProperties
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func AudioEncodingPropertiesCreateMp3(sampleRate uint32, channelCount uint32, bitrate uint32) (*AudioEncodingProperties, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Media.MediaProperties.AudioEncodingProperties", ole.NewGUID(GUIDiAudioEncodingPropertiesStatics))
	if err != nil {
		return nil, err
	}
	v := (*iAudioEncodingPropertiesStatics)(unsafe.Pointer(inspectable))

	var out *AudioEncodingProperties
	hr, _, _ := syscall.SyscallN(
		v.VTable().AudioEncodingPropertiesCreateMp3,
		0,                             // this is a static func, so there's no this
		uintptr(sampleRate),           // in uint32
		uintptr(channelCount),         // in uint32
		uintptr(bitrate),              // in uint32
		uintptr(unsafe.Pointer(&out)), // out AudioEncodingProperties
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func AudioEncodingPropertiesCreatePcm(sampleRate uint32, channelCount uint32, bitsPerSample uint32) (*AudioEncodingProperties, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Media.MediaProperties.AudioEncodingProperties", ole.NewGUID(GUIDiAudioEncodingPropertiesStatics))
	if err != nil {
		return nil, err
	}
	v := (*iAudioEncodingPropertiesStatics)(unsafe.Pointer(inspectable))

	var out *AudioEncodingProperties
	hr, _, _ := syscall.SyscallN(
		v.VTable().AudioEncodingPropertiesCreatePcm,
		0,                             // this is a static func, so there's no this
		uintptr(sampleRate),           // in uint32
		uintptr(channelCount),         // in uint32
		uintptr(bitsPerSample),        // in uint32
		uintptr(unsafe.Pointer(&out)), // out AudioEncodingProperties
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func AudioEncodingPropertiesCreateWma(sampleRate uint32, channelCount uint32, bitrate uint32) (*AudioEncodingProperties, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Media.MediaProperties.AudioEncodingProperties", ole.NewGUID(GUIDiAudioEncodingPropertiesStatics))
	if err != nil {
		return nil, err
	}
	v := (*iAudioEncodingPropertiesStatics)(unsafe.Pointer(inspectable))

	var out *AudioEncodingProperties
	hr, _, _ := syscall.SyscallN(
		v.VTable().AudioEncodingPropertiesCreateWma,
		0,                             // this is a static func, so there's no this
		uintptr(sampleRate),           // in uint32
		uintptr(channelCount),         // in uint32
		uintptr(bitrate),              // in uint32
		uintptr(unsafe.Pointer(&out)), // out AudioEncodingProperties
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

const GUIDiAudioEncodingPropertiesStatics2 string = "7489316f-77a0-433d-8ed5-4040280e8665"
const SignatureiAudioEncodingPropertiesStatics2 string = "{7489316f-77a0-433d-8ed5-4040280e8665}"

type iAudioEncodingPropertiesStatics2 struct {
	ole.IInspectable
}

type iAudioEncodingPropertiesStatics2Vtbl struct {
	ole.IInspectableVtbl

	AudioEncodingPropertiesCreateAlac uintptr
	AudioEncodingPropertiesCreateFlac uintptr
}

func (v *iAudioEncodingPropertiesStatics2) VTable() *iAudioEncodingPropertiesStatics2Vtbl {
	return (*iAudioEncodingPropertiesStatics2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func AudioEncodingPropertiesCreateAlac(sampleRate uint32, channelCount uint32, bitsPerSample uint32) (*AudioEncodingProperties, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Media.MediaProperties.AudioEncodingProperties", ole.NewGUID(GUIDiAudioEncodingPropertiesStatics2))
	if err != nil {
		return nil, err
	}
	v := (*iAudioEncodingPropertiesStatics2)(unsafe.Pointer(inspectable))

	var out *AudioEncodingProperties
	hr, _, _ := syscall.SyscallN(
		v.VTable().AudioEncodingPropertiesCreateAlac,
		0,                             // this is a static func, so there's no this
		uintptr(sampleRate),           // in uint32
		uintptr(channelCount),         // in uint32
		uintptr(bitsPerSample),        // in uint32
		uintptr(unsafe.Pointer(&out)), // out AudioEncodingProperties
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func AudioEncodingPropertiesCreateFlac(sampleRate uint32, channelCount uint32, bitsPerSample uint32) (*AudioEncodingProperties, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Media.MediaProperties.AudioEncodingProperties", ole.NewGUID(GUIDiAudioEncodingPropertiesStatics2))
	if err != nil {
		return nil, err
	}
	v := (*iAudioEncodingPropertiesStatics2)(unsafe.Pointer(inspectable))

	var out *AudioEncodingProperties
	hr, _, _ := syscall.SyscallN(
		v.VTable().AudioEncodingPropertiesCreateFlac,
		0,                             // this is a static func, so there's no this
		uintptr(sampleRate),           // in uint32
		uintptr(channelCount),         // in uint32
		uintptr(bitsPerSample),        // in uint32
		uintptr(unsafe.Pointer(&out)), // out AudioEncodingProperties
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
