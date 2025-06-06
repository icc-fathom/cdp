// Code generated by cdpgen. DO NOT EDIT.

// Package css implements the CSS domain. This domain exposes CSS read/write
// operations. All CSS objects (stylesheets, rules, and styles) have an
// associated `id` used in subsequent operations on the related object. Each
// object type has a specific `id` structure, and those are not interchangeable
// between objects of different kinds. CSS objects can be loaded using the
// `get*ForNode()` calls (which accept a DOM node id). A client can also keep
// track of stylesheets via the `styleSheetAdded`/`styleSheetRemoved` events
// and subsequently load the required stylesheet contents using the
// `getStyleSheet[Text]()` methods.
package css

import (
	"context"

	"github.com/icc-fathom/cdp/protocol/internal"
	"github.com/icc-fathom/cdp/rpcc"
)

// domainClient is a client for the CSS domain. This domain exposes CSS
// read/write operations. All CSS objects (stylesheets, rules, and styles) have
// an associated `id` used in subsequent operations on the related object. Each
// object type has a specific `id` structure, and those are not interchangeable
// between objects of different kinds. CSS objects can be loaded using the
// `get*ForNode()` calls (which accept a DOM node id). A client can also keep
// track of stylesheets via the `styleSheetAdded`/`styleSheetRemoved` events
// and subsequently load the required stylesheet contents using the
// `getStyleSheet[Text]()` methods.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the CSS domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// AddRule invokes the CSS method. Inserts a new rule with the given
// `ruleText` in a stylesheet with given `styleSheetId`, at the position
// specified by `location`.
func (d *domainClient) AddRule(ctx context.Context, args *AddRuleArgs) (reply *AddRuleReply, err error) {
	reply = new(AddRuleReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.addRule", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.addRule", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "AddRule", Err: err}
	}
	return
}

// CollectClassNames invokes the CSS method. Returns all class names from
// specified stylesheet.
func (d *domainClient) CollectClassNames(ctx context.Context, args *CollectClassNamesArgs) (reply *CollectClassNamesReply, err error) {
	reply = new(CollectClassNamesReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.collectClassNames", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.collectClassNames", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "CollectClassNames", Err: err}
	}
	return
}

// CreateStyleSheet invokes the CSS method. Creates a new special
// "via-inspector" stylesheet in the frame with given `frameId`.
func (d *domainClient) CreateStyleSheet(ctx context.Context, args *CreateStyleSheetArgs) (reply *CreateStyleSheetReply, err error) {
	reply = new(CreateStyleSheetReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.createStyleSheet", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.createStyleSheet", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "CreateStyleSheet", Err: err}
	}
	return
}

// Disable invokes the CSS method. Disables the CSS agent for the given page.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "CSS.disable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "Disable", Err: err}
	}
	return
}

// Enable invokes the CSS method. Enables the CSS agent for the given page.
// Clients should not assume that the CSS agent has been enabled until the
// result of this command is received.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "CSS.enable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "Enable", Err: err}
	}
	return
}

// ForcePseudoState invokes the CSS method. Ensures that the given node will
// have specified pseudo-classes whenever its style is computed by the browser.
func (d *domainClient) ForcePseudoState(ctx context.Context, args *ForcePseudoStateArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.forcePseudoState", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.forcePseudoState", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "ForcePseudoState", Err: err}
	}
	return
}

// GetBackgroundColors invokes the CSS method.
func (d *domainClient) GetBackgroundColors(ctx context.Context, args *GetBackgroundColorsArgs) (reply *GetBackgroundColorsReply, err error) {
	reply = new(GetBackgroundColorsReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.getBackgroundColors", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.getBackgroundColors", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "GetBackgroundColors", Err: err}
	}
	return
}

// GetComputedStyleForNode invokes the CSS method. Returns the computed style
// for a DOM node identified by `nodeId`.
func (d *domainClient) GetComputedStyleForNode(ctx context.Context, args *GetComputedStyleForNodeArgs) (reply *GetComputedStyleForNodeReply, err error) {
	reply = new(GetComputedStyleForNodeReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.getComputedStyleForNode", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.getComputedStyleForNode", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "GetComputedStyleForNode", Err: err}
	}
	return
}

// GetInlineStylesForNode invokes the CSS method. Returns the styles defined
// inline (explicitly in the "style" attribute and implicitly, using DOM
// attributes) for a DOM node identified by `nodeId`.
func (d *domainClient) GetInlineStylesForNode(ctx context.Context, args *GetInlineStylesForNodeArgs) (reply *GetInlineStylesForNodeReply, err error) {
	reply = new(GetInlineStylesForNodeReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.getInlineStylesForNode", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.getInlineStylesForNode", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "GetInlineStylesForNode", Err: err}
	}
	return
}

// GetMatchedStylesForNode invokes the CSS method. Returns requested styles
// for a DOM node identified by `nodeId`.
func (d *domainClient) GetMatchedStylesForNode(ctx context.Context, args *GetMatchedStylesForNodeArgs) (reply *GetMatchedStylesForNodeReply, err error) {
	reply = new(GetMatchedStylesForNodeReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.getMatchedStylesForNode", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.getMatchedStylesForNode", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "GetMatchedStylesForNode", Err: err}
	}
	return
}

// GetMediaQueries invokes the CSS method. Returns all media queries parsed by
// the rendering engine.
func (d *domainClient) GetMediaQueries(ctx context.Context) (reply *GetMediaQueriesReply, err error) {
	reply = new(GetMediaQueriesReply)
	err = rpcc.Invoke(ctx, "CSS.getMediaQueries", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "GetMediaQueries", Err: err}
	}
	return
}

// GetPlatformFontsForNode invokes the CSS method. Requests information about
// platform fonts which we used to render child TextNodes in the given node.
func (d *domainClient) GetPlatformFontsForNode(ctx context.Context, args *GetPlatformFontsForNodeArgs) (reply *GetPlatformFontsForNodeReply, err error) {
	reply = new(GetPlatformFontsForNodeReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.getPlatformFontsForNode", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.getPlatformFontsForNode", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "GetPlatformFontsForNode", Err: err}
	}
	return
}

// GetStyleSheetText invokes the CSS method. Returns the current textual
// content for a stylesheet.
func (d *domainClient) GetStyleSheetText(ctx context.Context, args *GetStyleSheetTextArgs) (reply *GetStyleSheetTextReply, err error) {
	reply = new(GetStyleSheetTextReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.getStyleSheetText", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.getStyleSheetText", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "GetStyleSheetText", Err: err}
	}
	return
}

// GetLayersForNode invokes the CSS method. Returns all layers parsed by the
// rendering engine for the tree scope of a node. Given a DOM element
// identified by nodeId, getLayersForNode returns the root layer for the
// nearest ancestor document or shadow root. The layer root contains the full
// layer tree for the tree scope and their ordering.
func (d *domainClient) GetLayersForNode(ctx context.Context, args *GetLayersForNodeArgs) (reply *GetLayersForNodeReply, err error) {
	reply = new(GetLayersForNodeReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.getLayersForNode", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.getLayersForNode", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "GetLayersForNode", Err: err}
	}
	return
}

// GetLocationForSelector invokes the CSS method. Given a CSS selector text
// and a style sheet ID, getLocationForSelector returns an array of locations
// of the CSS selector in the style sheet.
func (d *domainClient) GetLocationForSelector(ctx context.Context, args *GetLocationForSelectorArgs) (reply *GetLocationForSelectorReply, err error) {
	reply = new(GetLocationForSelectorReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.getLocationForSelector", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.getLocationForSelector", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "GetLocationForSelector", Err: err}
	}
	return
}

// TrackComputedStyleUpdates invokes the CSS method. Starts tracking the given
// computed styles for updates. The specified array of properties replaces the
// one previously specified. Pass empty array to disable tracking. Use
// takeComputedStyleUpdates to retrieve the list of nodes that had properties
// modified. The changes to computed style properties are only tracked for
// nodes pushed to the front-end by the DOM agent. If no changes to the tracked
// properties occur after the node has been pushed to the front-end, no updates
// will be issued for the node.
func (d *domainClient) TrackComputedStyleUpdates(ctx context.Context, args *TrackComputedStyleUpdatesArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.trackComputedStyleUpdates", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.trackComputedStyleUpdates", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "TrackComputedStyleUpdates", Err: err}
	}
	return
}

// TakeComputedStyleUpdates invokes the CSS method. Polls the next batch of
// computed style updates.
func (d *domainClient) TakeComputedStyleUpdates(ctx context.Context) (reply *TakeComputedStyleUpdatesReply, err error) {
	reply = new(TakeComputedStyleUpdatesReply)
	err = rpcc.Invoke(ctx, "CSS.takeComputedStyleUpdates", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "TakeComputedStyleUpdates", Err: err}
	}
	return
}

// SetEffectivePropertyValueForNode invokes the CSS method. Find a rule with
// the given active property for the given node and set the new value for this
// property
func (d *domainClient) SetEffectivePropertyValueForNode(ctx context.Context, args *SetEffectivePropertyValueForNodeArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.setEffectivePropertyValueForNode", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.setEffectivePropertyValueForNode", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "SetEffectivePropertyValueForNode", Err: err}
	}
	return
}

// SetPropertyRulePropertyName invokes the CSS method. Modifies the property
// rule property name.
func (d *domainClient) SetPropertyRulePropertyName(ctx context.Context, args *SetPropertyRulePropertyNameArgs) (reply *SetPropertyRulePropertyNameReply, err error) {
	reply = new(SetPropertyRulePropertyNameReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.setPropertyRulePropertyName", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.setPropertyRulePropertyName", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "SetPropertyRulePropertyName", Err: err}
	}
	return
}

// SetKeyframeKey invokes the CSS method. Modifies the keyframe rule key text.
func (d *domainClient) SetKeyframeKey(ctx context.Context, args *SetKeyframeKeyArgs) (reply *SetKeyframeKeyReply, err error) {
	reply = new(SetKeyframeKeyReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.setKeyframeKey", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.setKeyframeKey", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "SetKeyframeKey", Err: err}
	}
	return
}

// SetMediaText invokes the CSS method. Modifies the rule selector.
func (d *domainClient) SetMediaText(ctx context.Context, args *SetMediaTextArgs) (reply *SetMediaTextReply, err error) {
	reply = new(SetMediaTextReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.setMediaText", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.setMediaText", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "SetMediaText", Err: err}
	}
	return
}

// SetContainerQueryText invokes the CSS method. Modifies the expression of a
// container query.
func (d *domainClient) SetContainerQueryText(ctx context.Context, args *SetContainerQueryTextArgs) (reply *SetContainerQueryTextReply, err error) {
	reply = new(SetContainerQueryTextReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.setContainerQueryText", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.setContainerQueryText", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "SetContainerQueryText", Err: err}
	}
	return
}

// SetSupportsText invokes the CSS method. Modifies the expression of a
// supports at-rule.
func (d *domainClient) SetSupportsText(ctx context.Context, args *SetSupportsTextArgs) (reply *SetSupportsTextReply, err error) {
	reply = new(SetSupportsTextReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.setSupportsText", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.setSupportsText", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "SetSupportsText", Err: err}
	}
	return
}

// SetScopeText invokes the CSS method. Modifies the expression of a scope
// at-rule.
func (d *domainClient) SetScopeText(ctx context.Context, args *SetScopeTextArgs) (reply *SetScopeTextReply, err error) {
	reply = new(SetScopeTextReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.setScopeText", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.setScopeText", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "SetScopeText", Err: err}
	}
	return
}

// SetRuleSelector invokes the CSS method. Modifies the rule selector.
func (d *domainClient) SetRuleSelector(ctx context.Context, args *SetRuleSelectorArgs) (reply *SetRuleSelectorReply, err error) {
	reply = new(SetRuleSelectorReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.setRuleSelector", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.setRuleSelector", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "SetRuleSelector", Err: err}
	}
	return
}

// SetStyleSheetText invokes the CSS method. Sets the new stylesheet text.
func (d *domainClient) SetStyleSheetText(ctx context.Context, args *SetStyleSheetTextArgs) (reply *SetStyleSheetTextReply, err error) {
	reply = new(SetStyleSheetTextReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.setStyleSheetText", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.setStyleSheetText", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "SetStyleSheetText", Err: err}
	}
	return
}

// SetStyleTexts invokes the CSS method. Applies specified style edits one
// after another in the given order.
func (d *domainClient) SetStyleTexts(ctx context.Context, args *SetStyleTextsArgs) (reply *SetStyleTextsReply, err error) {
	reply = new(SetStyleTextsReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.setStyleTexts", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.setStyleTexts", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "SetStyleTexts", Err: err}
	}
	return
}

// StartRuleUsageTracking invokes the CSS method. Enables the selector
// recording.
func (d *domainClient) StartRuleUsageTracking(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "CSS.startRuleUsageTracking", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "StartRuleUsageTracking", Err: err}
	}
	return
}

// StopRuleUsageTracking invokes the CSS method. Stop tracking rule usage and
// return the list of rules that were used since last call to
// `takeCoverageDelta` (or since start of coverage instrumentation).
func (d *domainClient) StopRuleUsageTracking(ctx context.Context) (reply *StopRuleUsageTrackingReply, err error) {
	reply = new(StopRuleUsageTrackingReply)
	err = rpcc.Invoke(ctx, "CSS.stopRuleUsageTracking", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "StopRuleUsageTracking", Err: err}
	}
	return
}

// TakeCoverageDelta invokes the CSS method. Obtain list of rules that became
// used since last call to this method (or since start of coverage
// instrumentation).
func (d *domainClient) TakeCoverageDelta(ctx context.Context) (reply *TakeCoverageDeltaReply, err error) {
	reply = new(TakeCoverageDeltaReply)
	err = rpcc.Invoke(ctx, "CSS.takeCoverageDelta", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "TakeCoverageDelta", Err: err}
	}
	return
}

// SetLocalFontsEnabled invokes the CSS method. Enables/disables rendering of
// local CSS fonts (enabled by default).
func (d *domainClient) SetLocalFontsEnabled(ctx context.Context, args *SetLocalFontsEnabledArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "CSS.setLocalFontsEnabled", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CSS.setLocalFontsEnabled", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CSS", Op: "SetLocalFontsEnabled", Err: err}
	}
	return
}

func (d *domainClient) FontsUpdated(ctx context.Context) (FontsUpdatedClient, error) {
	s, err := rpcc.NewStream(ctx, "CSS.fontsUpdated", d.conn)
	if err != nil {
		return nil, err
	}
	return &fontsUpdatedClient{Stream: s}, nil
}

type fontsUpdatedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *fontsUpdatedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *fontsUpdatedClient) Recv() (*FontsUpdatedReply, error) {
	event := new(FontsUpdatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "CSS", Op: "FontsUpdated Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) MediaQueryResultChanged(ctx context.Context) (MediaQueryResultChangedClient, error) {
	s, err := rpcc.NewStream(ctx, "CSS.mediaQueryResultChanged", d.conn)
	if err != nil {
		return nil, err
	}
	return &mediaQueryResultChangedClient{Stream: s}, nil
}

type mediaQueryResultChangedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *mediaQueryResultChangedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *mediaQueryResultChangedClient) Recv() (*MediaQueryResultChangedReply, error) {
	event := new(MediaQueryResultChangedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "CSS", Op: "MediaQueryResultChanged Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) StyleSheetAdded(ctx context.Context) (StyleSheetAddedClient, error) {
	s, err := rpcc.NewStream(ctx, "CSS.styleSheetAdded", d.conn)
	if err != nil {
		return nil, err
	}
	return &styleSheetAddedClient{Stream: s}, nil
}

type styleSheetAddedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *styleSheetAddedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *styleSheetAddedClient) Recv() (*StyleSheetAddedReply, error) {
	event := new(StyleSheetAddedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "CSS", Op: "StyleSheetAdded Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) StyleSheetChanged(ctx context.Context) (StyleSheetChangedClient, error) {
	s, err := rpcc.NewStream(ctx, "CSS.styleSheetChanged", d.conn)
	if err != nil {
		return nil, err
	}
	return &styleSheetChangedClient{Stream: s}, nil
}

type styleSheetChangedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *styleSheetChangedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *styleSheetChangedClient) Recv() (*StyleSheetChangedReply, error) {
	event := new(StyleSheetChangedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "CSS", Op: "StyleSheetChanged Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) StyleSheetRemoved(ctx context.Context) (StyleSheetRemovedClient, error) {
	s, err := rpcc.NewStream(ctx, "CSS.styleSheetRemoved", d.conn)
	if err != nil {
		return nil, err
	}
	return &styleSheetRemovedClient{Stream: s}, nil
}

type styleSheetRemovedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *styleSheetRemovedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *styleSheetRemovedClient) Recv() (*StyleSheetRemovedReply, error) {
	event := new(StyleSheetRemovedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "CSS", Op: "StyleSheetRemoved Recv", Err: err}
	}
	return event, nil
}
