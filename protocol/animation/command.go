// Code generated by cdpgen. DO NOT EDIT.

package animation

import (
	"github.com/icc-fathom/cdp/protocol/runtime"
)

// GetCurrentTimeArgs represents the arguments for GetCurrentTime in the Animation domain.
type GetCurrentTimeArgs struct {
	ID string `json:"id"` // Id of animation.
}

// NewGetCurrentTimeArgs initializes GetCurrentTimeArgs with the required arguments.
func NewGetCurrentTimeArgs(id string) *GetCurrentTimeArgs {
	args := new(GetCurrentTimeArgs)
	args.ID = id
	return args
}

// GetCurrentTimeReply represents the return values for GetCurrentTime in the Animation domain.
type GetCurrentTimeReply struct {
	CurrentTime float64 `json:"currentTime"` // Current time of the page.
}

// GetPlaybackRateReply represents the return values for GetPlaybackRate in the Animation domain.
type GetPlaybackRateReply struct {
	PlaybackRate float64 `json:"playbackRate"` // Playback rate for animations on page.
}

// ReleaseAnimationsArgs represents the arguments for ReleaseAnimations in the Animation domain.
type ReleaseAnimationsArgs struct {
	Animations []string `json:"animations"` // List of animation ids to seek.
}

// NewReleaseAnimationsArgs initializes ReleaseAnimationsArgs with the required arguments.
func NewReleaseAnimationsArgs(animations []string) *ReleaseAnimationsArgs {
	args := new(ReleaseAnimationsArgs)
	args.Animations = animations
	return args
}

// ResolveAnimationArgs represents the arguments for ResolveAnimation in the Animation domain.
type ResolveAnimationArgs struct {
	AnimationID string `json:"animationId"` // Animation id.
}

// NewResolveAnimationArgs initializes ResolveAnimationArgs with the required arguments.
func NewResolveAnimationArgs(animationID string) *ResolveAnimationArgs {
	args := new(ResolveAnimationArgs)
	args.AnimationID = animationID
	return args
}

// ResolveAnimationReply represents the return values for ResolveAnimation in the Animation domain.
type ResolveAnimationReply struct {
	RemoteObject runtime.RemoteObject `json:"remoteObject"` // Corresponding remote object.
}

// SeekAnimationsArgs represents the arguments for SeekAnimations in the Animation domain.
type SeekAnimationsArgs struct {
	Animations  []string `json:"animations"`  // List of animation ids to seek.
	CurrentTime float64  `json:"currentTime"` // Set the current time of each animation.
}

// NewSeekAnimationsArgs initializes SeekAnimationsArgs with the required arguments.
func NewSeekAnimationsArgs(animations []string, currentTime float64) *SeekAnimationsArgs {
	args := new(SeekAnimationsArgs)
	args.Animations = animations
	args.CurrentTime = currentTime
	return args
}

// SetPausedArgs represents the arguments for SetPaused in the Animation domain.
type SetPausedArgs struct {
	Animations []string `json:"animations"` // Animations to set the pause state of.
	Paused     bool     `json:"paused"`     // Paused state to set to.
}

// NewSetPausedArgs initializes SetPausedArgs with the required arguments.
func NewSetPausedArgs(animations []string, paused bool) *SetPausedArgs {
	args := new(SetPausedArgs)
	args.Animations = animations
	args.Paused = paused
	return args
}

// SetPlaybackRateArgs represents the arguments for SetPlaybackRate in the Animation domain.
type SetPlaybackRateArgs struct {
	PlaybackRate float64 `json:"playbackRate"` // Playback rate for animations on page
}

// NewSetPlaybackRateArgs initializes SetPlaybackRateArgs with the required arguments.
func NewSetPlaybackRateArgs(playbackRate float64) *SetPlaybackRateArgs {
	args := new(SetPlaybackRateArgs)
	args.PlaybackRate = playbackRate
	return args
}

// SetTimingArgs represents the arguments for SetTiming in the Animation domain.
type SetTimingArgs struct {
	AnimationID string  `json:"animationId"` // Animation id.
	Duration    float64 `json:"duration"`    // Duration of the animation.
	Delay       float64 `json:"delay"`       // Delay of the animation.
}

// NewSetTimingArgs initializes SetTimingArgs with the required arguments.
func NewSetTimingArgs(animationID string, duration float64, delay float64) *SetTimingArgs {
	args := new(SetTimingArgs)
	args.AnimationID = animationID
	args.Duration = duration
	args.Delay = delay
	return args
}
