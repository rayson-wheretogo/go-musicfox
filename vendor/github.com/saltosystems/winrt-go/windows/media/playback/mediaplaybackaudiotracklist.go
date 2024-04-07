// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package playback

import (
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
	"github.com/saltosystems/winrt-go/windows/media/core"
)

const SignatureMediaPlaybackAudioTrackList string = "rc(Windows.Media.Playback.MediaPlaybackAudioTrackList;{77206f1f-c34f-494f-8077-2bad9ff4ecf1})"

type MediaPlaybackAudioTrackList struct {
	ole.IUnknown
}

func (impl *MediaPlaybackAudioTrackList) AddSelectedIndexChanged(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(core.GUIDISingleSelectMediaTrackList))
	defer itf.Release()
	v := (*core.ISingleSelectMediaTrackList)(unsafe.Pointer(itf))
	return v.AddSelectedIndexChanged(handler)
}

func (impl *MediaPlaybackAudioTrackList) RemoveSelectedIndexChanged(token foundation.EventRegistrationToken) error {
	itf := impl.MustQueryInterface(ole.NewGUID(core.GUIDISingleSelectMediaTrackList))
	defer itf.Release()
	v := (*core.ISingleSelectMediaTrackList)(unsafe.Pointer(itf))
	return v.RemoveSelectedIndexChanged(token)
}

func (impl *MediaPlaybackAudioTrackList) SetSelectedIndex(value int32) error {
	itf := impl.MustQueryInterface(ole.NewGUID(core.GUIDISingleSelectMediaTrackList))
	defer itf.Release()
	v := (*core.ISingleSelectMediaTrackList)(unsafe.Pointer(itf))
	return v.SetSelectedIndex(value)
}

func (impl *MediaPlaybackAudioTrackList) GetSelectedIndex() (int32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(core.GUIDISingleSelectMediaTrackList))
	defer itf.Release()
	v := (*core.ISingleSelectMediaTrackList)(unsafe.Pointer(itf))
	return v.GetSelectedIndex()
}
