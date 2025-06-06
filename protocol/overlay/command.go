// Code generated by cdpgen. DO NOT EDIT.

package overlay

import (
	"encoding/json"

	"github.com/icc-fathom/cdp/protocol/dom"
	"github.com/icc-fathom/cdp/protocol/page"
	"github.com/icc-fathom/cdp/protocol/runtime"
)

// GetHighlightObjectForTestArgs represents the arguments for GetHighlightObjectForTest in the Overlay domain.
type GetHighlightObjectForTestArgs struct {
	NodeID                dom.NodeID  `json:"nodeId"`                          // Id of the node to get highlight object for.
	IncludeDistance       *bool       `json:"includeDistance,omitempty"`       // Whether to include distance info.
	IncludeStyle          *bool       `json:"includeStyle,omitempty"`          // Whether to include style info.
	ColorFormat           ColorFormat `json:"colorFormat,omitempty"`           // The color format to get config with (default: hex).
	ShowAccessibilityInfo *bool       `json:"showAccessibilityInfo,omitempty"` // Whether to show accessibility info (default: true).
}

// NewGetHighlightObjectForTestArgs initializes GetHighlightObjectForTestArgs with the required arguments.
func NewGetHighlightObjectForTestArgs(nodeID dom.NodeID) *GetHighlightObjectForTestArgs {
	args := new(GetHighlightObjectForTestArgs)
	args.NodeID = nodeID
	return args
}

// SetIncludeDistance sets the IncludeDistance optional argument.
// Whether to include distance info.
func (a *GetHighlightObjectForTestArgs) SetIncludeDistance(includeDistance bool) *GetHighlightObjectForTestArgs {
	a.IncludeDistance = &includeDistance
	return a
}

// SetIncludeStyle sets the IncludeStyle optional argument. Whether to
// include style info.
func (a *GetHighlightObjectForTestArgs) SetIncludeStyle(includeStyle bool) *GetHighlightObjectForTestArgs {
	a.IncludeStyle = &includeStyle
	return a
}

// SetColorFormat sets the ColorFormat optional argument. The color
// format to get config with (default: hex).
func (a *GetHighlightObjectForTestArgs) SetColorFormat(colorFormat ColorFormat) *GetHighlightObjectForTestArgs {
	a.ColorFormat = colorFormat
	return a
}

// SetShowAccessibilityInfo sets the ShowAccessibilityInfo optional argument.
// Whether to show accessibility info (default: true).
func (a *GetHighlightObjectForTestArgs) SetShowAccessibilityInfo(showAccessibilityInfo bool) *GetHighlightObjectForTestArgs {
	a.ShowAccessibilityInfo = &showAccessibilityInfo
	return a
}

// GetHighlightObjectForTestReply represents the return values for GetHighlightObjectForTest in the Overlay domain.
type GetHighlightObjectForTestReply struct {
	Highlight json.RawMessage `json:"highlight"` // Highlight data for the node.
}

// GetGridHighlightObjectsForTestArgs represents the arguments for GetGridHighlightObjectsForTest in the Overlay domain.
type GetGridHighlightObjectsForTestArgs struct {
	NodeIDs []dom.NodeID `json:"nodeIds"` // Ids of the node to get highlight object for.
}

// NewGetGridHighlightObjectsForTestArgs initializes GetGridHighlightObjectsForTestArgs with the required arguments.
func NewGetGridHighlightObjectsForTestArgs(nodeIDs []dom.NodeID) *GetGridHighlightObjectsForTestArgs {
	args := new(GetGridHighlightObjectsForTestArgs)
	args.NodeIDs = nodeIDs
	return args
}

// GetGridHighlightObjectsForTestReply represents the return values for GetGridHighlightObjectsForTest in the Overlay domain.
type GetGridHighlightObjectsForTestReply struct {
	Highlights json.RawMessage `json:"highlights"` // Grid Highlight data for the node ids provided.
}

// GetSourceOrderHighlightObjectForTestArgs represents the arguments for GetSourceOrderHighlightObjectForTest in the Overlay domain.
type GetSourceOrderHighlightObjectForTestArgs struct {
	NodeID dom.NodeID `json:"nodeId"` // Id of the node to highlight.
}

// NewGetSourceOrderHighlightObjectForTestArgs initializes GetSourceOrderHighlightObjectForTestArgs with the required arguments.
func NewGetSourceOrderHighlightObjectForTestArgs(nodeID dom.NodeID) *GetSourceOrderHighlightObjectForTestArgs {
	args := new(GetSourceOrderHighlightObjectForTestArgs)
	args.NodeID = nodeID
	return args
}

// GetSourceOrderHighlightObjectForTestReply represents the return values for GetSourceOrderHighlightObjectForTest in the Overlay domain.
type GetSourceOrderHighlightObjectForTestReply struct {
	Highlight json.RawMessage `json:"highlight"` // Source order highlight data for the node id provided.
}

// HighlightFrameArgs represents the arguments for HighlightFrame in the Overlay domain.
type HighlightFrameArgs struct {
	FrameID             page.FrameID `json:"frameId"`                       // Identifier of the frame to highlight.
	ContentColor        *dom.RGBA    `json:"contentColor,omitempty"`        // The content box highlight fill color (default: transparent).
	ContentOutlineColor *dom.RGBA    `json:"contentOutlineColor,omitempty"` // The content box highlight outline color (default: transparent).
}

// NewHighlightFrameArgs initializes HighlightFrameArgs with the required arguments.
func NewHighlightFrameArgs(frameID page.FrameID) *HighlightFrameArgs {
	args := new(HighlightFrameArgs)
	args.FrameID = frameID
	return args
}

// SetContentColor sets the ContentColor optional argument. The
// content box highlight fill color (default: transparent).
func (a *HighlightFrameArgs) SetContentColor(contentColor dom.RGBA) *HighlightFrameArgs {
	a.ContentColor = &contentColor
	return a
}

// SetContentOutlineColor sets the ContentOutlineColor optional argument.
// The content box highlight outline color (default: transparent).
func (a *HighlightFrameArgs) SetContentOutlineColor(contentOutlineColor dom.RGBA) *HighlightFrameArgs {
	a.ContentOutlineColor = &contentOutlineColor
	return a
}

// HighlightNodeArgs represents the arguments for HighlightNode in the Overlay domain.
type HighlightNodeArgs struct {
	HighlightConfig HighlightConfig         `json:"highlightConfig"`         // A descriptor for the highlight appearance.
	NodeID          *dom.NodeID             `json:"nodeId,omitempty"`        // Identifier of the node to highlight.
	BackendNodeID   *dom.BackendNodeID      `json:"backendNodeId,omitempty"` // Identifier of the backend node to highlight.
	ObjectID        *runtime.RemoteObjectID `json:"objectId,omitempty"`      // JavaScript object id of the node to be highlighted.
	Selector        *string                 `json:"selector,omitempty"`      // Selectors to highlight relevant nodes.
}

// NewHighlightNodeArgs initializes HighlightNodeArgs with the required arguments.
func NewHighlightNodeArgs(highlightConfig HighlightConfig) *HighlightNodeArgs {
	args := new(HighlightNodeArgs)
	args.HighlightConfig = highlightConfig
	return args
}

// SetNodeID sets the NodeID optional argument. Identifier of the node
// to highlight.
func (a *HighlightNodeArgs) SetNodeID(nodeID dom.NodeID) *HighlightNodeArgs {
	a.NodeID = &nodeID
	return a
}

// SetBackendNodeID sets the BackendNodeID optional argument.
// Identifier of the backend node to highlight.
func (a *HighlightNodeArgs) SetBackendNodeID(backendNodeID dom.BackendNodeID) *HighlightNodeArgs {
	a.BackendNodeID = &backendNodeID
	return a
}

// SetObjectID sets the ObjectID optional argument. JavaScript object
// id of the node to be highlighted.
func (a *HighlightNodeArgs) SetObjectID(objectID runtime.RemoteObjectID) *HighlightNodeArgs {
	a.ObjectID = &objectID
	return a
}

// SetSelector sets the Selector optional argument. Selectors to
// highlight relevant nodes.
func (a *HighlightNodeArgs) SetSelector(selector string) *HighlightNodeArgs {
	a.Selector = &selector
	return a
}

// HighlightQuadArgs represents the arguments for HighlightQuad in the Overlay domain.
type HighlightQuadArgs struct {
	Quad         dom.Quad  `json:"quad"`                   // Quad to highlight
	Color        *dom.RGBA `json:"color,omitempty"`        // The highlight fill color (default: transparent).
	OutlineColor *dom.RGBA `json:"outlineColor,omitempty"` // The highlight outline color (default: transparent).
}

// NewHighlightQuadArgs initializes HighlightQuadArgs with the required arguments.
func NewHighlightQuadArgs(quad dom.Quad) *HighlightQuadArgs {
	args := new(HighlightQuadArgs)
	args.Quad = quad
	return args
}

// SetColor sets the Color optional argument. The highlight fill color
// (default: transparent).
func (a *HighlightQuadArgs) SetColor(color dom.RGBA) *HighlightQuadArgs {
	a.Color = &color
	return a
}

// SetOutlineColor sets the OutlineColor optional argument. The
// highlight outline color (default: transparent).
func (a *HighlightQuadArgs) SetOutlineColor(outlineColor dom.RGBA) *HighlightQuadArgs {
	a.OutlineColor = &outlineColor
	return a
}

// HighlightRectArgs represents the arguments for HighlightRect in the Overlay domain.
type HighlightRectArgs struct {
	X            int       `json:"x"`                      // X coordinate
	Y            int       `json:"y"`                      // Y coordinate
	Width        int       `json:"width"`                  // Rectangle width
	Height       int       `json:"height"`                 // Rectangle height
	Color        *dom.RGBA `json:"color,omitempty"`        // The highlight fill color (default: transparent).
	OutlineColor *dom.RGBA `json:"outlineColor,omitempty"` // The highlight outline color (default: transparent).
}

// NewHighlightRectArgs initializes HighlightRectArgs with the required arguments.
func NewHighlightRectArgs(x int, y int, width int, height int) *HighlightRectArgs {
	args := new(HighlightRectArgs)
	args.X = x
	args.Y = y
	args.Width = width
	args.Height = height
	return args
}

// SetColor sets the Color optional argument. The highlight fill color
// (default: transparent).
func (a *HighlightRectArgs) SetColor(color dom.RGBA) *HighlightRectArgs {
	a.Color = &color
	return a
}

// SetOutlineColor sets the OutlineColor optional argument. The
// highlight outline color (default: transparent).
func (a *HighlightRectArgs) SetOutlineColor(outlineColor dom.RGBA) *HighlightRectArgs {
	a.OutlineColor = &outlineColor
	return a
}

// HighlightSourceOrderArgs represents the arguments for HighlightSourceOrder in the Overlay domain.
type HighlightSourceOrderArgs struct {
	SourceOrderConfig SourceOrderConfig       `json:"sourceOrderConfig"`       // A descriptor for the appearance of the overlay drawing.
	NodeID            *dom.NodeID             `json:"nodeId,omitempty"`        // Identifier of the node to highlight.
	BackendNodeID     *dom.BackendNodeID      `json:"backendNodeId,omitempty"` // Identifier of the backend node to highlight.
	ObjectID          *runtime.RemoteObjectID `json:"objectId,omitempty"`      // JavaScript object id of the node to be highlighted.
}

// NewHighlightSourceOrderArgs initializes HighlightSourceOrderArgs with the required arguments.
func NewHighlightSourceOrderArgs(sourceOrderConfig SourceOrderConfig) *HighlightSourceOrderArgs {
	args := new(HighlightSourceOrderArgs)
	args.SourceOrderConfig = sourceOrderConfig
	return args
}

// SetNodeID sets the NodeID optional argument. Identifier of the node
// to highlight.
func (a *HighlightSourceOrderArgs) SetNodeID(nodeID dom.NodeID) *HighlightSourceOrderArgs {
	a.NodeID = &nodeID
	return a
}

// SetBackendNodeID sets the BackendNodeID optional argument.
// Identifier of the backend node to highlight.
func (a *HighlightSourceOrderArgs) SetBackendNodeID(backendNodeID dom.BackendNodeID) *HighlightSourceOrderArgs {
	a.BackendNodeID = &backendNodeID
	return a
}

// SetObjectID sets the ObjectID optional argument. JavaScript object
// id of the node to be highlighted.
func (a *HighlightSourceOrderArgs) SetObjectID(objectID runtime.RemoteObjectID) *HighlightSourceOrderArgs {
	a.ObjectID = &objectID
	return a
}

// SetInspectModeArgs represents the arguments for SetInspectMode in the Overlay domain.
type SetInspectModeArgs struct {
	Mode            InspectMode      `json:"mode"`                      // Set an inspection mode.
	HighlightConfig *HighlightConfig `json:"highlightConfig,omitempty"` // A descriptor for the highlight appearance of hovered-over nodes. May be omitted if `enabled == false`.
}

// NewSetInspectModeArgs initializes SetInspectModeArgs with the required arguments.
func NewSetInspectModeArgs(mode InspectMode) *SetInspectModeArgs {
	args := new(SetInspectModeArgs)
	args.Mode = mode
	return args
}

// SetHighlightConfig sets the HighlightConfig optional argument. A
// descriptor for the highlight appearance of hovered-over nodes. May
// be omitted if `enabled == false`.
func (a *SetInspectModeArgs) SetHighlightConfig(highlightConfig HighlightConfig) *SetInspectModeArgs {
	a.HighlightConfig = &highlightConfig
	return a
}

// SetShowAdHighlightsArgs represents the arguments for SetShowAdHighlights in the Overlay domain.
type SetShowAdHighlightsArgs struct {
	Show bool `json:"show"` // True for showing ad highlights
}

// NewSetShowAdHighlightsArgs initializes SetShowAdHighlightsArgs with the required arguments.
func NewSetShowAdHighlightsArgs(show bool) *SetShowAdHighlightsArgs {
	args := new(SetShowAdHighlightsArgs)
	args.Show = show
	return args
}

// SetPausedInDebuggerMessageArgs represents the arguments for SetPausedInDebuggerMessage in the Overlay domain.
type SetPausedInDebuggerMessageArgs struct {
	Message *string `json:"message,omitempty"` // The message to display, also triggers resume and step over controls.
}

// NewSetPausedInDebuggerMessageArgs initializes SetPausedInDebuggerMessageArgs with the required arguments.
func NewSetPausedInDebuggerMessageArgs() *SetPausedInDebuggerMessageArgs {
	args := new(SetPausedInDebuggerMessageArgs)

	return args
}

// SetMessage sets the Message optional argument. The message to
// display, also triggers resume and step over controls.
func (a *SetPausedInDebuggerMessageArgs) SetMessage(message string) *SetPausedInDebuggerMessageArgs {
	a.Message = &message
	return a
}

// SetShowDebugBordersArgs represents the arguments for SetShowDebugBorders in the Overlay domain.
type SetShowDebugBordersArgs struct {
	Show bool `json:"show"` // True for showing debug borders
}

// NewSetShowDebugBordersArgs initializes SetShowDebugBordersArgs with the required arguments.
func NewSetShowDebugBordersArgs(show bool) *SetShowDebugBordersArgs {
	args := new(SetShowDebugBordersArgs)
	args.Show = show
	return args
}

// SetShowFPSCounterArgs represents the arguments for SetShowFPSCounter in the Overlay domain.
type SetShowFPSCounterArgs struct {
	Show bool `json:"show"` // True for showing the FPS counter
}

// NewSetShowFPSCounterArgs initializes SetShowFPSCounterArgs with the required arguments.
func NewSetShowFPSCounterArgs(show bool) *SetShowFPSCounterArgs {
	args := new(SetShowFPSCounterArgs)
	args.Show = show
	return args
}

// SetShowGridOverlaysArgs represents the arguments for SetShowGridOverlays in the Overlay domain.
type SetShowGridOverlaysArgs struct {
	GridNodeHighlightConfigs []GridNodeHighlightConfig `json:"gridNodeHighlightConfigs"` // An array of node identifiers and descriptors for the highlight appearance.
}

// NewSetShowGridOverlaysArgs initializes SetShowGridOverlaysArgs with the required arguments.
func NewSetShowGridOverlaysArgs(gridNodeHighlightConfigs []GridNodeHighlightConfig) *SetShowGridOverlaysArgs {
	args := new(SetShowGridOverlaysArgs)
	args.GridNodeHighlightConfigs = gridNodeHighlightConfigs
	return args
}

// SetShowFlexOverlaysArgs represents the arguments for SetShowFlexOverlays in the Overlay domain.
type SetShowFlexOverlaysArgs struct {
	FlexNodeHighlightConfigs []FlexNodeHighlightConfig `json:"flexNodeHighlightConfigs"` // An array of node identifiers and descriptors for the highlight appearance.
}

// NewSetShowFlexOverlaysArgs initializes SetShowFlexOverlaysArgs with the required arguments.
func NewSetShowFlexOverlaysArgs(flexNodeHighlightConfigs []FlexNodeHighlightConfig) *SetShowFlexOverlaysArgs {
	args := new(SetShowFlexOverlaysArgs)
	args.FlexNodeHighlightConfigs = flexNodeHighlightConfigs
	return args
}

// SetShowScrollSnapOverlaysArgs represents the arguments for SetShowScrollSnapOverlays in the Overlay domain.
type SetShowScrollSnapOverlaysArgs struct {
	ScrollSnapHighlightConfigs []ScrollSnapHighlightConfig `json:"scrollSnapHighlightConfigs"` // An array of node identifiers and descriptors for the highlight appearance.
}

// NewSetShowScrollSnapOverlaysArgs initializes SetShowScrollSnapOverlaysArgs with the required arguments.
func NewSetShowScrollSnapOverlaysArgs(scrollSnapHighlightConfigs []ScrollSnapHighlightConfig) *SetShowScrollSnapOverlaysArgs {
	args := new(SetShowScrollSnapOverlaysArgs)
	args.ScrollSnapHighlightConfigs = scrollSnapHighlightConfigs
	return args
}

// SetShowContainerQueryOverlaysArgs represents the arguments for SetShowContainerQueryOverlays in the Overlay domain.
type SetShowContainerQueryOverlaysArgs struct {
	ContainerQueryHighlightConfigs []ContainerQueryHighlightConfig `json:"containerQueryHighlightConfigs"` // An array of node identifiers and descriptors for the highlight appearance.
}

// NewSetShowContainerQueryOverlaysArgs initializes SetShowContainerQueryOverlaysArgs with the required arguments.
func NewSetShowContainerQueryOverlaysArgs(containerQueryHighlightConfigs []ContainerQueryHighlightConfig) *SetShowContainerQueryOverlaysArgs {
	args := new(SetShowContainerQueryOverlaysArgs)
	args.ContainerQueryHighlightConfigs = containerQueryHighlightConfigs
	return args
}

// SetShowPaintRectsArgs represents the arguments for SetShowPaintRects in the Overlay domain.
type SetShowPaintRectsArgs struct {
	Result bool `json:"result"` // True for showing paint rectangles
}

// NewSetShowPaintRectsArgs initializes SetShowPaintRectsArgs with the required arguments.
func NewSetShowPaintRectsArgs(result bool) *SetShowPaintRectsArgs {
	args := new(SetShowPaintRectsArgs)
	args.Result = result
	return args
}

// SetShowLayoutShiftRegionsArgs represents the arguments for SetShowLayoutShiftRegions in the Overlay domain.
type SetShowLayoutShiftRegionsArgs struct {
	Result bool `json:"result"` // True for showing layout shift regions
}

// NewSetShowLayoutShiftRegionsArgs initializes SetShowLayoutShiftRegionsArgs with the required arguments.
func NewSetShowLayoutShiftRegionsArgs(result bool) *SetShowLayoutShiftRegionsArgs {
	args := new(SetShowLayoutShiftRegionsArgs)
	args.Result = result
	return args
}

// SetShowScrollBottleneckRectsArgs represents the arguments for SetShowScrollBottleneckRects in the Overlay domain.
type SetShowScrollBottleneckRectsArgs struct {
	Show bool `json:"show"` // True for showing scroll bottleneck rects
}

// NewSetShowScrollBottleneckRectsArgs initializes SetShowScrollBottleneckRectsArgs with the required arguments.
func NewSetShowScrollBottleneckRectsArgs(show bool) *SetShowScrollBottleneckRectsArgs {
	args := new(SetShowScrollBottleneckRectsArgs)
	args.Show = show
	return args
}

// SetShowHitTestBordersArgs represents the arguments for SetShowHitTestBorders in the Overlay domain.
type SetShowHitTestBordersArgs struct {
	Show bool `json:"show"` // True for showing hit-test borders
}

// NewSetShowHitTestBordersArgs initializes SetShowHitTestBordersArgs with the required arguments.
func NewSetShowHitTestBordersArgs(show bool) *SetShowHitTestBordersArgs {
	args := new(SetShowHitTestBordersArgs)
	args.Show = show
	return args
}

// SetShowWebVitalsArgs represents the arguments for SetShowWebVitals in the Overlay domain.
type SetShowWebVitalsArgs struct {
	Show bool `json:"show"` // No description.
}

// NewSetShowWebVitalsArgs initializes SetShowWebVitalsArgs with the required arguments.
func NewSetShowWebVitalsArgs(show bool) *SetShowWebVitalsArgs {
	args := new(SetShowWebVitalsArgs)
	args.Show = show
	return args
}

// SetShowViewportSizeOnResizeArgs represents the arguments for SetShowViewportSizeOnResize in the Overlay domain.
type SetShowViewportSizeOnResizeArgs struct {
	Show bool `json:"show"` // Whether to paint size or not.
}

// NewSetShowViewportSizeOnResizeArgs initializes SetShowViewportSizeOnResizeArgs with the required arguments.
func NewSetShowViewportSizeOnResizeArgs(show bool) *SetShowViewportSizeOnResizeArgs {
	args := new(SetShowViewportSizeOnResizeArgs)
	args.Show = show
	return args
}

// SetShowHingeArgs represents the arguments for SetShowHinge in the Overlay domain.
type SetShowHingeArgs struct {
	HingeConfig *HingeConfig `json:"hingeConfig,omitempty"` // hinge data, null means hideHinge
}

// NewSetShowHingeArgs initializes SetShowHingeArgs with the required arguments.
func NewSetShowHingeArgs() *SetShowHingeArgs {
	args := new(SetShowHingeArgs)

	return args
}

// SetHingeConfig sets the HingeConfig optional argument. hinge data,
// null means hideHinge
func (a *SetShowHingeArgs) SetHingeConfig(hingeConfig HingeConfig) *SetShowHingeArgs {
	a.HingeConfig = &hingeConfig
	return a
}

// SetShowIsolatedElementsArgs represents the arguments for SetShowIsolatedElements in the Overlay domain.
type SetShowIsolatedElementsArgs struct {
	IsolatedElementHighlightConfigs []IsolatedElementHighlightConfig `json:"isolatedElementHighlightConfigs"` // An array of node identifiers and descriptors for the highlight appearance.
}

// NewSetShowIsolatedElementsArgs initializes SetShowIsolatedElementsArgs with the required arguments.
func NewSetShowIsolatedElementsArgs(isolatedElementHighlightConfigs []IsolatedElementHighlightConfig) *SetShowIsolatedElementsArgs {
	args := new(SetShowIsolatedElementsArgs)
	args.IsolatedElementHighlightConfigs = isolatedElementHighlightConfigs
	return args
}

// SetShowWindowControlsOverlayArgs represents the arguments for SetShowWindowControlsOverlay in the Overlay domain.
type SetShowWindowControlsOverlayArgs struct {
	WindowControlsOverlayConfig *WindowControlsOverlayConfig `json:"windowControlsOverlayConfig,omitempty"` // Window Controls Overlay data, null means hide Window Controls Overlay
}

// NewSetShowWindowControlsOverlayArgs initializes SetShowWindowControlsOverlayArgs with the required arguments.
func NewSetShowWindowControlsOverlayArgs() *SetShowWindowControlsOverlayArgs {
	args := new(SetShowWindowControlsOverlayArgs)

	return args
}

// SetWindowControlsOverlayConfig sets the WindowControlsOverlayConfig optional argument.
// Window Controls Overlay data, null means hide Window Controls
// Overlay
func (a *SetShowWindowControlsOverlayArgs) SetWindowControlsOverlayConfig(windowControlsOverlayConfig WindowControlsOverlayConfig) *SetShowWindowControlsOverlayArgs {
	a.WindowControlsOverlayConfig = &windowControlsOverlayConfig
	return a
}
