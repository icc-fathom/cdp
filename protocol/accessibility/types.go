// Code generated by cdpgen. DO NOT EDIT.

package accessibility

import (
	"encoding/json"

	"github.com/icc-fathom/cdp/protocol/dom"
	"github.com/icc-fathom/cdp/protocol/page"
)

// AXNodeID Unique accessibility node identifier.
type AXNodeID string

// AXValueType Enum of possible property types.
type AXValueType string

// AXValueType as enums.
const (
	AXValueTypeNotSet             AXValueType = ""
	AXValueTypeBoolean            AXValueType = "boolean"
	AXValueTypeTristate           AXValueType = "tristate"
	AXValueTypeBooleanOrUndefined AXValueType = "booleanOrUndefined"
	AXValueTypeIDRef              AXValueType = "idref"
	AXValueTypeIDRefList          AXValueType = "idrefList"
	AXValueTypeInteger            AXValueType = "integer"
	AXValueTypeNode               AXValueType = "node"
	AXValueTypeNodeList           AXValueType = "nodeList"
	AXValueTypeNumber             AXValueType = "number"
	AXValueTypeString             AXValueType = "string"
	AXValueTypeComputedString     AXValueType = "computedString"
	AXValueTypeToken              AXValueType = "token"
	AXValueTypeTokenList          AXValueType = "tokenList"
	AXValueTypeDOMRelation        AXValueType = "domRelation"
	AXValueTypeRole               AXValueType = "role"
	AXValueTypeInternalRole       AXValueType = "internalRole"
	AXValueTypeValueUndefined     AXValueType = "valueUndefined"
)

func (e AXValueType) Valid() bool {
	switch e {
	case "boolean", "tristate", "booleanOrUndefined", "idref", "idrefList", "integer", "node", "nodeList", "number", "string", "computedString", "token", "tokenList", "domRelation", "role", "internalRole", "valueUndefined":
		return true
	default:
		return false
	}
}

func (e AXValueType) String() string {
	return string(e)
}

// AXValueSourceType Enum of possible property sources.
type AXValueSourceType string

// AXValueSourceType as enums.
const (
	AXValueSourceTypeNotSet         AXValueSourceType = ""
	AXValueSourceTypeAttribute      AXValueSourceType = "attribute"
	AXValueSourceTypeImplicit       AXValueSourceType = "implicit"
	AXValueSourceTypeStyle          AXValueSourceType = "style"
	AXValueSourceTypeContents       AXValueSourceType = "contents"
	AXValueSourceTypePlaceholder    AXValueSourceType = "placeholder"
	AXValueSourceTypeRelatedElement AXValueSourceType = "relatedElement"
)

func (e AXValueSourceType) Valid() bool {
	switch e {
	case "attribute", "implicit", "style", "contents", "placeholder", "relatedElement":
		return true
	default:
		return false
	}
}

func (e AXValueSourceType) String() string {
	return string(e)
}

// AXValueNativeSourceType Enum of possible native property sources (as a
// subtype of a particular AXValueSourceType).
type AXValueNativeSourceType string

// AXValueNativeSourceType as enums.
const (
	AXValueNativeSourceTypeNotSet         AXValueNativeSourceType = ""
	AXValueNativeSourceTypeDescription    AXValueNativeSourceType = "description"
	AXValueNativeSourceTypeFigcaption     AXValueNativeSourceType = "figcaption"
	AXValueNativeSourceTypeLabel          AXValueNativeSourceType = "label"
	AXValueNativeSourceTypeLabelfor       AXValueNativeSourceType = "labelfor"
	AXValueNativeSourceTypeLabelwrapped   AXValueNativeSourceType = "labelwrapped"
	AXValueNativeSourceTypeLegend         AXValueNativeSourceType = "legend"
	AXValueNativeSourceTypeRubyannotation AXValueNativeSourceType = "rubyannotation"
	AXValueNativeSourceTypeTablecaption   AXValueNativeSourceType = "tablecaption"
	AXValueNativeSourceTypeTitle          AXValueNativeSourceType = "title"
	AXValueNativeSourceTypeOther          AXValueNativeSourceType = "other"
)

func (e AXValueNativeSourceType) Valid() bool {
	switch e {
	case "description", "figcaption", "label", "labelfor", "labelwrapped", "legend", "rubyannotation", "tablecaption", "title", "other":
		return true
	default:
		return false
	}
}

func (e AXValueNativeSourceType) String() string {
	return string(e)
}

// AXValueSource A single source for a computed AX property.
type AXValueSource struct {
	Type              AXValueSourceType       `json:"type"`                        // What type of source this is.
	Value             *AXValue                `json:"value,omitempty"`             // The value of this property source.
	Attribute         *string                 `json:"attribute,omitempty"`         // The name of the relevant attribute, if any.
	AttributeValue    *AXValue                `json:"attributeValue,omitempty"`    // The value of the relevant attribute, if any.
	Superseded        *bool                   `json:"superseded,omitempty"`        // Whether this source is superseded by a higher priority source.
	NativeSource      AXValueNativeSourceType `json:"nativeSource,omitempty"`      // The native markup source for this value, e.g. a `<label>` element.
	NativeSourceValue *AXValue                `json:"nativeSourceValue,omitempty"` // The value, such as a node or node list, of the native source.
	Invalid           *bool                   `json:"invalid,omitempty"`           // Whether the value for this property is invalid.
	InvalidReason     *string                 `json:"invalidReason,omitempty"`     // Reason for the value being invalid, if it is.
}

// AXRelatedNode
type AXRelatedNode struct {
	BackendDOMNodeID dom.BackendNodeID `json:"backendDOMNodeId"` // The BackendNodeId of the related DOM node.
	IDRef            *string           `json:"idref,omitempty"`  // The IDRef value provided, if any.
	Text             *string           `json:"text,omitempty"`   // The text alternative of this node in the current context.
}

// AXProperty
type AXProperty struct {
	Name  AXPropertyName `json:"name"`  // The name of this property.
	Value AXValue        `json:"value"` // The value of this property.
}

// AXValue A single computed AX property.
type AXValue struct {
	Type         AXValueType     `json:"type"`                   // The type of this value.
	Value        json.RawMessage `json:"value,omitempty"`        // The computed value of this property.
	RelatedNodes []AXRelatedNode `json:"relatedNodes,omitempty"` // One or more related nodes, if applicable.
	Sources      []AXValueSource `json:"sources,omitempty"`      // The sources which contributed to the computation of this property.
}

// AXPropertyName Values of AXProperty name: - from 'busy' to
// 'roledescription': states which apply to every AX node - from 'live' to
// 'root': attributes which apply to nodes in live regions - from
// 'autocomplete' to 'valuetext': attributes which apply to widgets - from
// 'checked' to 'selected': states which apply to widgets - from
// 'activedescendant' to 'owns' - relationships between elements other than
// parent/child/sibling.
type AXPropertyName string

// AXPropertyName as enums.
const (
	AXPropertyNameNotSet           AXPropertyName = ""
	AXPropertyNameBusy             AXPropertyName = "busy"
	AXPropertyNameDisabled         AXPropertyName = "disabled"
	AXPropertyNameEditable         AXPropertyName = "editable"
	AXPropertyNameFocusable        AXPropertyName = "focusable"
	AXPropertyNameFocused          AXPropertyName = "focused"
	AXPropertyNameHidden           AXPropertyName = "hidden"
	AXPropertyNameHiddenRoot       AXPropertyName = "hiddenRoot"
	AXPropertyNameInvalid          AXPropertyName = "invalid"
	AXPropertyNameKeyshortcuts     AXPropertyName = "keyshortcuts"
	AXPropertyNameSettable         AXPropertyName = "settable"
	AXPropertyNameRoledescription  AXPropertyName = "roledescription"
	AXPropertyNameLive             AXPropertyName = "live"
	AXPropertyNameAtomic           AXPropertyName = "atomic"
	AXPropertyNameRelevant         AXPropertyName = "relevant"
	AXPropertyNameRoot             AXPropertyName = "root"
	AXPropertyNameAutocomplete     AXPropertyName = "autocomplete"
	AXPropertyNameHasPopup         AXPropertyName = "hasPopup"
	AXPropertyNameLevel            AXPropertyName = "level"
	AXPropertyNameMultiselectable  AXPropertyName = "multiselectable"
	AXPropertyNameOrientation      AXPropertyName = "orientation"
	AXPropertyNameMultiline        AXPropertyName = "multiline"
	AXPropertyNameReadonly         AXPropertyName = "readonly"
	AXPropertyNameRequired         AXPropertyName = "required"
	AXPropertyNameValuemin         AXPropertyName = "valuemin"
	AXPropertyNameValuemax         AXPropertyName = "valuemax"
	AXPropertyNameValuetext        AXPropertyName = "valuetext"
	AXPropertyNameChecked          AXPropertyName = "checked"
	AXPropertyNameExpanded         AXPropertyName = "expanded"
	AXPropertyNameModal            AXPropertyName = "modal"
	AXPropertyNamePressed          AXPropertyName = "pressed"
	AXPropertyNameSelected         AXPropertyName = "selected"
	AXPropertyNameActivedescendant AXPropertyName = "activedescendant"
	AXPropertyNameControls         AXPropertyName = "controls"
	AXPropertyNameDescribedby      AXPropertyName = "describedby"
	AXPropertyNameDetails          AXPropertyName = "details"
	AXPropertyNameErrormessage     AXPropertyName = "errormessage"
	AXPropertyNameFlowto           AXPropertyName = "flowto"
	AXPropertyNameLabelledby       AXPropertyName = "labelledby"
	AXPropertyNameOwns             AXPropertyName = "owns"
	AXPropertyNameURL              AXPropertyName = "url"
)

func (e AXPropertyName) Valid() bool {
	switch e {
	case "busy", "disabled", "editable", "focusable", "focused", "hidden", "hiddenRoot", "invalid", "keyshortcuts", "settable", "roledescription", "live", "atomic", "relevant", "root", "autocomplete", "hasPopup", "level", "multiselectable", "orientation", "multiline", "readonly", "required", "valuemin", "valuemax", "valuetext", "checked", "expanded", "modal", "pressed", "selected", "activedescendant", "controls", "describedby", "details", "errormessage", "flowto", "labelledby", "owns", "url":
		return true
	default:
		return false
	}
}

func (e AXPropertyName) String() string {
	return string(e)
}

// AXNode A node in the accessibility tree.
type AXNode struct {
	NodeID           AXNodeID           `json:"nodeId"`                     // Unique identifier for this node.
	Ignored          bool               `json:"ignored"`                    // Whether this node is ignored for accessibility
	IgnoredReasons   []AXProperty       `json:"ignoredReasons,omitempty"`   // Collection of reasons why this node is hidden.
	Role             *AXValue           `json:"role,omitempty"`             // This `Node`'s role, whether explicit or implicit.
	ChromeRole       *AXValue           `json:"chromeRole,omitempty"`       // This `Node`'s Chrome raw role.
	Name             *AXValue           `json:"name,omitempty"`             // The accessible name for this `Node`.
	Description      *AXValue           `json:"description,omitempty"`      // The accessible description for this `Node`.
	Value            *AXValue           `json:"value,omitempty"`            // The value for this `Node`.
	Properties       []AXProperty       `json:"properties,omitempty"`       // All other properties
	ParentID         *AXNodeID          `json:"parentId,omitempty"`         // ID for this node's parent.
	ChildIDs         []AXNodeID         `json:"childIds,omitempty"`         // IDs for each of this node's child nodes.
	BackendDOMNodeID *dom.BackendNodeID `json:"backendDOMNodeId,omitempty"` // The backend ID for the associated DOM node, if any.
	FrameID          *page.FrameID      `json:"frameId,omitempty"`          // The frame ID for the frame associated with this nodes document.
}
