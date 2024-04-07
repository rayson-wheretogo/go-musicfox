// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package core

type MseReadyState int32

const SignatureMseReadyState string = "enum(Windows.Media.Core.MseReadyState;i4)"

const (
	MseReadyStateClosed MseReadyState = 0
	MseReadyStateOpen   MseReadyState = 1
	MseReadyStateEnded  MseReadyState = 2
)
