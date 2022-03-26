// Package css provides the Chrome DevTools Protocol
// commands, types, and events for the CSS domain.
//
// This domain exposes CSS read/write operations. All CSS objects
// (stylesheets, rules, and styles) have an associated id used in subsequent
// operations on the related object. Each object type has a specific id
// structure, and those are not interchangeable between objects of different
// kinds. CSS objects can be loaded using the get*ForNode() calls (which accept
// a DOM node id). A client can also keep track of stylesheets via the
// styleSheetAdded/styleSheetRemoved events and subsequently load the required
// stylesheet contents using the getStyleSheet[Text]() methods.
//
// Generated by the cdproto-gen command.
package css

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// AddRuleParams inserts a new rule with the given ruleText in a stylesheet
// with given styleSheetId, at the position specified by location.
type AddRuleParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // The css style sheet identifier where a new rule should be inserted.
	RuleText     string       `json:"ruleText"`     // The text of a new rule.
	Location     *SourceRange `json:"location"`     // Text position of a new rule in the target style sheet.
}

// AddRule inserts a new rule with the given ruleText in a stylesheet with
// given styleSheetId, at the position specified by location.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-addRule
//
// parameters:
//   styleSheetID - The css style sheet identifier where a new rule should be inserted.
//   ruleText - The text of a new rule.
//   location - Text position of a new rule in the target style sheet.
func AddRule(styleSheetID StyleSheetID, ruleText string, location *SourceRange) *AddRuleParams {
	return &AddRuleParams{
		StyleSheetID: styleSheetID,
		RuleText:     ruleText,
		Location:     location,
	}
}

// AddRuleReturns return values.
type AddRuleReturns struct {
	Rule *Rule `json:"rule,omitempty"` // The newly created rule.
}

// Do executes CSS.addRule against the provided context.
//
// returns:
//   rule - The newly created rule.
func (p *AddRuleParams) Do(ctx context.Context) (rule *Rule, err error) {
	// execute
	var res AddRuleReturns
	err = cdp.Execute(ctx, CommandAddRule, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Rule, nil
}

// CollectClassNamesParams returns all class names from specified stylesheet.
type CollectClassNamesParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
}

// CollectClassNames returns all class names from specified stylesheet.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-collectClassNames
//
// parameters:
//   styleSheetID
func CollectClassNames(styleSheetID StyleSheetID) *CollectClassNamesParams {
	return &CollectClassNamesParams{
		StyleSheetID: styleSheetID,
	}
}

// CollectClassNamesReturns return values.
type CollectClassNamesReturns struct {
	ClassNames []string `json:"classNames,omitempty"` // Class name list.
}

// Do executes CSS.collectClassNames against the provided context.
//
// returns:
//   classNames - Class name list.
func (p *CollectClassNamesParams) Do(ctx context.Context) (classNames []string, err error) {
	// execute
	var res CollectClassNamesReturns
	err = cdp.Execute(ctx, CommandCollectClassNames, p, &res)
	if err != nil {
		return nil, err
	}

	return res.ClassNames, nil
}

// CreateStyleSheetParams creates a new special "via-inspector" stylesheet in
// the frame with given frameId.
type CreateStyleSheetParams struct {
	FrameID cdp.FrameID `json:"frameId"` // Identifier of the frame where "via-inspector" stylesheet should be created.
}

// CreateStyleSheet creates a new special "via-inspector" stylesheet in the
// frame with given frameId.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-createStyleSheet
//
// parameters:
//   frameID - Identifier of the frame where "via-inspector" stylesheet should be created.
func CreateStyleSheet(frameID cdp.FrameID) *CreateStyleSheetParams {
	return &CreateStyleSheetParams{
		FrameID: frameID,
	}
}

// CreateStyleSheetReturns return values.
type CreateStyleSheetReturns struct {
	StyleSheetID StyleSheetID `json:"styleSheetId,omitempty"` // Identifier of the created "via-inspector" stylesheet.
}

// Do executes CSS.createStyleSheet against the provided context.
//
// returns:
//   styleSheetID - Identifier of the created "via-inspector" stylesheet.
func (p *CreateStyleSheetParams) Do(ctx context.Context) (styleSheetID StyleSheetID, err error) {
	// execute
	var res CreateStyleSheetReturns
	err = cdp.Execute(ctx, CommandCreateStyleSheet, p, &res)
	if err != nil {
		return "", err
	}

	return res.StyleSheetID, nil
}

// DisableParams disables the CSS agent for the given page.
type DisableParams struct{}

// Disable disables the CSS agent for the given page.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-disable
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes CSS.disable against the provided context.
func (p *DisableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisable, nil, nil)
}

// EnableParams enables the CSS agent for the given page. Clients should not
// assume that the CSS agent has been enabled until the result of this command
// is received.
type EnableParams struct{}

// Enable enables the CSS agent for the given page. Clients should not assume
// that the CSS agent has been enabled until the result of this command is
// received.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-enable
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes CSS.enable against the provided context.
func (p *EnableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandEnable, nil, nil)
}

// ForcePseudoStateParams ensures that the given node will have specified
// pseudo-classes whenever its style is computed by the browser.
type ForcePseudoStateParams struct {
	NodeID              cdp.NodeID `json:"nodeId"`              // The element id for which to force the pseudo state.
	ForcedPseudoClasses []string   `json:"forcedPseudoClasses"` // Element pseudo classes to force when computing the element's style.
}

// ForcePseudoState ensures that the given node will have specified
// pseudo-classes whenever its style is computed by the browser.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-forcePseudoState
//
// parameters:
//   nodeID - The element id for which to force the pseudo state.
//   forcedPseudoClasses - Element pseudo classes to force when computing the element's style.
func ForcePseudoState(nodeID cdp.NodeID, forcedPseudoClasses []string) *ForcePseudoStateParams {
	return &ForcePseudoStateParams{
		NodeID:              nodeID,
		ForcedPseudoClasses: forcedPseudoClasses,
	}
}

// Do executes CSS.forcePseudoState against the provided context.
func (p *ForcePseudoStateParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandForcePseudoState, p, nil)
}

// GetBackgroundColorsParams [no description].
type GetBackgroundColorsParams struct {
	NodeID cdp.NodeID `json:"nodeId"` // Id of the node to get background colors for.
}

// GetBackgroundColors [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-getBackgroundColors
//
// parameters:
//   nodeID - Id of the node to get background colors for.
func GetBackgroundColors(nodeID cdp.NodeID) *GetBackgroundColorsParams {
	return &GetBackgroundColorsParams{
		NodeID: nodeID,
	}
}

// GetBackgroundColorsReturns return values.
type GetBackgroundColorsReturns struct {
	BackgroundColors   []string `json:"backgroundColors,omitempty"`   // The range of background colors behind this element, if it contains any visible text. If no visible text is present, this will be undefined. In the case of a flat background color, this will consist of simply that color. In the case of a gradient, this will consist of each of the color stops. For anything more complicated, this will be an empty array. Images will be ignored (as if the image had failed to load).
	ComputedFontSize   string   `json:"computedFontSize,omitempty"`   // The computed font size for this node, as a CSS computed value string (e.g. '12px').
	ComputedFontWeight string   `json:"computedFontWeight,omitempty"` // The computed font weight for this node, as a CSS computed value string (e.g. 'normal' or '100').
}

// Do executes CSS.getBackgroundColors against the provided context.
//
// returns:
//   backgroundColors - The range of background colors behind this element, if it contains any visible text. If no visible text is present, this will be undefined. In the case of a flat background color, this will consist of simply that color. In the case of a gradient, this will consist of each of the color stops. For anything more complicated, this will be an empty array. Images will be ignored (as if the image had failed to load).
//   computedFontSize - The computed font size for this node, as a CSS computed value string (e.g. '12px').
//   computedFontWeight - The computed font weight for this node, as a CSS computed value string (e.g. 'normal' or '100').
func (p *GetBackgroundColorsParams) Do(ctx context.Context) (backgroundColors []string, computedFontSize string, computedFontWeight string, err error) {
	// execute
	var res GetBackgroundColorsReturns
	err = cdp.Execute(ctx, CommandGetBackgroundColors, p, &res)
	if err != nil {
		return nil, "", "", err
	}

	return res.BackgroundColors, res.ComputedFontSize, res.ComputedFontWeight, nil
}

// GetComputedStyleForNodeParams returns the computed style for a DOM node
// identified by nodeId.
type GetComputedStyleForNodeParams struct {
	NodeID cdp.NodeID `json:"nodeId"`
}

// GetComputedStyleForNode returns the computed style for a DOM node
// identified by nodeId.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-getComputedStyleForNode
//
// parameters:
//   nodeID
func GetComputedStyleForNode(nodeID cdp.NodeID) *GetComputedStyleForNodeParams {
	return &GetComputedStyleForNodeParams{
		NodeID: nodeID,
	}
}

// GetComputedStyleForNodeReturns return values.
type GetComputedStyleForNodeReturns struct {
	ComputedStyle []*ComputedStyleProperty `json:"computedStyle,omitempty"` // Computed style for the specified DOM node.
}

// Do executes CSS.getComputedStyleForNode against the provided context.
//
// returns:
//   computedStyle - Computed style for the specified DOM node.
func (p *GetComputedStyleForNodeParams) Do(ctx context.Context) (computedStyle []*ComputedStyleProperty, err error) {
	// execute
	var res GetComputedStyleForNodeReturns
	err = cdp.Execute(ctx, CommandGetComputedStyleForNode, p, &res)
	if err != nil {
		return nil, err
	}

	return res.ComputedStyle, nil
}

// GetInlineStylesForNodeParams returns the styles defined inline (explicitly
// in the "style" attribute and implicitly, using DOM attributes) for a DOM node
// identified by nodeId.
type GetInlineStylesForNodeParams struct {
	NodeID cdp.NodeID `json:"nodeId"`
}

// GetInlineStylesForNode returns the styles defined inline (explicitly in
// the "style" attribute and implicitly, using DOM attributes) for a DOM node
// identified by nodeId.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-getInlineStylesForNode
//
// parameters:
//   nodeID
func GetInlineStylesForNode(nodeID cdp.NodeID) *GetInlineStylesForNodeParams {
	return &GetInlineStylesForNodeParams{
		NodeID: nodeID,
	}
}

// GetInlineStylesForNodeReturns return values.
type GetInlineStylesForNodeReturns struct {
	InlineStyle     *Style `json:"inlineStyle,omitempty"`     // Inline style for the specified DOM node.
	AttributesStyle *Style `json:"attributesStyle,omitempty"` // Attribute-defined element style (e.g. resulting from "width=20 height=100%").
}

// Do executes CSS.getInlineStylesForNode against the provided context.
//
// returns:
//   inlineStyle - Inline style for the specified DOM node.
//   attributesStyle - Attribute-defined element style (e.g. resulting from "width=20 height=100%").
func (p *GetInlineStylesForNodeParams) Do(ctx context.Context) (inlineStyle *Style, attributesStyle *Style, err error) {
	// execute
	var res GetInlineStylesForNodeReturns
	err = cdp.Execute(ctx, CommandGetInlineStylesForNode, p, &res)
	if err != nil {
		return nil, nil, err
	}

	return res.InlineStyle, res.AttributesStyle, nil
}

// GetMatchedStylesForNodeParams returns requested styles for a DOM node
// identified by nodeId.
type GetMatchedStylesForNodeParams struct {
	NodeID cdp.NodeID `json:"nodeId"`
}

// GetMatchedStylesForNode returns requested styles for a DOM node identified
// by nodeId.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-getMatchedStylesForNode
//
// parameters:
//   nodeID
func GetMatchedStylesForNode(nodeID cdp.NodeID) *GetMatchedStylesForNodeParams {
	return &GetMatchedStylesForNodeParams{
		NodeID: nodeID,
	}
}

// GetMatchedStylesForNodeReturns return values.
type GetMatchedStylesForNodeReturns struct {
	InlineStyle             *Style                           `json:"inlineStyle,omitempty"`             // Inline style for the specified DOM node.
	AttributesStyle         *Style                           `json:"attributesStyle,omitempty"`         // Attribute-defined element style (e.g. resulting from "width=20 height=100%").
	MatchedCSSRules         []*RuleMatch                     `json:"matchedCSSRules,omitempty"`         // CSS rules matching this node, from all applicable stylesheets.
	PseudoElements          []*PseudoElementMatches          `json:"pseudoElements,omitempty"`          // Pseudo style matches for this node.
	Inherited               []*InheritedStyleEntry           `json:"inherited,omitempty"`               // A chain of inherited styles (from the immediate node parent up to the DOM tree root).
	InheritedPseudoElements []*InheritedPseudoElementMatches `json:"inheritedPseudoElements,omitempty"` // A chain of inherited pseudo element styles (from the immediate node parent up to the DOM tree root).
	CSSKeyframesRules       []*KeyframesRule                 `json:"cssKeyframesRules,omitempty"`       // A list of CSS keyframed animations matching this node.
}

// Do executes CSS.getMatchedStylesForNode against the provided context.
//
// returns:
//   inlineStyle - Inline style for the specified DOM node.
//   attributesStyle - Attribute-defined element style (e.g. resulting from "width=20 height=100%").
//   matchedCSSRules - CSS rules matching this node, from all applicable stylesheets.
//   pseudoElements - Pseudo style matches for this node.
//   inherited - A chain of inherited styles (from the immediate node parent up to the DOM tree root).
//   inheritedPseudoElements - A chain of inherited pseudo element styles (from the immediate node parent up to the DOM tree root).
//   cssKeyframesRules - A list of CSS keyframed animations matching this node.
func (p *GetMatchedStylesForNodeParams) Do(ctx context.Context) (inlineStyle *Style, attributesStyle *Style, matchedCSSRules []*RuleMatch, pseudoElements []*PseudoElementMatches, inherited []*InheritedStyleEntry, inheritedPseudoElements []*InheritedPseudoElementMatches, cssKeyframesRules []*KeyframesRule, err error) {
	// execute
	var res GetMatchedStylesForNodeReturns
	err = cdp.Execute(ctx, CommandGetMatchedStylesForNode, p, &res)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}

	return res.InlineStyle, res.AttributesStyle, res.MatchedCSSRules, res.PseudoElements, res.Inherited, res.InheritedPseudoElements, res.CSSKeyframesRules, nil
}

// GetMediaQueriesParams returns all media queries parsed by the rendering
// engine.
type GetMediaQueriesParams struct{}

// GetMediaQueries returns all media queries parsed by the rendering engine.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-getMediaQueries
func GetMediaQueries() *GetMediaQueriesParams {
	return &GetMediaQueriesParams{}
}

// GetMediaQueriesReturns return values.
type GetMediaQueriesReturns struct {
	Medias []*Media `json:"medias,omitempty"`
}

// Do executes CSS.getMediaQueries against the provided context.
//
// returns:
//   medias
func (p *GetMediaQueriesParams) Do(ctx context.Context) (medias []*Media, err error) {
	// execute
	var res GetMediaQueriesReturns
	err = cdp.Execute(ctx, CommandGetMediaQueries, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Medias, nil
}

// GetPlatformFontsForNodeParams requests information about platform fonts
// which we used to render child TextNodes in the given node.
type GetPlatformFontsForNodeParams struct {
	NodeID cdp.NodeID `json:"nodeId"`
}

// GetPlatformFontsForNode requests information about platform fonts which we
// used to render child TextNodes in the given node.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-getPlatformFontsForNode
//
// parameters:
//   nodeID
func GetPlatformFontsForNode(nodeID cdp.NodeID) *GetPlatformFontsForNodeParams {
	return &GetPlatformFontsForNodeParams{
		NodeID: nodeID,
	}
}

// GetPlatformFontsForNodeReturns return values.
type GetPlatformFontsForNodeReturns struct {
	Fonts []*PlatformFontUsage `json:"fonts,omitempty"` // Usage statistics for every employed platform font.
}

// Do executes CSS.getPlatformFontsForNode against the provided context.
//
// returns:
//   fonts - Usage statistics for every employed platform font.
func (p *GetPlatformFontsForNodeParams) Do(ctx context.Context) (fonts []*PlatformFontUsage, err error) {
	// execute
	var res GetPlatformFontsForNodeReturns
	err = cdp.Execute(ctx, CommandGetPlatformFontsForNode, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Fonts, nil
}

// GetStyleSheetTextParams returns the current textual content for a
// stylesheet.
type GetStyleSheetTextParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
}

// GetStyleSheetText returns the current textual content for a stylesheet.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-getStyleSheetText
//
// parameters:
//   styleSheetID
func GetStyleSheetText(styleSheetID StyleSheetID) *GetStyleSheetTextParams {
	return &GetStyleSheetTextParams{
		StyleSheetID: styleSheetID,
	}
}

// GetStyleSheetTextReturns return values.
type GetStyleSheetTextReturns struct {
	Text string `json:"text,omitempty"` // The stylesheet text.
}

// Do executes CSS.getStyleSheetText against the provided context.
//
// returns:
//   text - The stylesheet text.
func (p *GetStyleSheetTextParams) Do(ctx context.Context) (text string, err error) {
	// execute
	var res GetStyleSheetTextReturns
	err = cdp.Execute(ctx, CommandGetStyleSheetText, p, &res)
	if err != nil {
		return "", err
	}

	return res.Text, nil
}

// GetLayersForNodeParams returns all layers parsed by the rendering engine
// for the tree scope of a node. Given a DOM element identified by nodeId,
// getLayersForNode returns the root layer for the nearest ancestor document or
// shadow root. The layer root contains the full layer tree for the tree scope
// and their ordering.
type GetLayersForNodeParams struct {
	NodeID cdp.NodeID `json:"nodeId"`
}

// GetLayersForNode returns all layers parsed by the rendering engine for the
// tree scope of a node. Given a DOM element identified by nodeId,
// getLayersForNode returns the root layer for the nearest ancestor document or
// shadow root. The layer root contains the full layer tree for the tree scope
// and their ordering.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-getLayersForNode
//
// parameters:
//   nodeID
func GetLayersForNode(nodeID cdp.NodeID) *GetLayersForNodeParams {
	return &GetLayersForNodeParams{
		NodeID: nodeID,
	}
}

// GetLayersForNodeReturns return values.
type GetLayersForNodeReturns struct {
	RootLayer *LayerData `json:"rootLayer,omitempty"`
}

// Do executes CSS.getLayersForNode against the provided context.
//
// returns:
//   rootLayer
func (p *GetLayersForNodeParams) Do(ctx context.Context) (rootLayer *LayerData, err error) {
	// execute
	var res GetLayersForNodeReturns
	err = cdp.Execute(ctx, CommandGetLayersForNode, p, &res)
	if err != nil {
		return nil, err
	}

	return res.RootLayer, nil
}

// TrackComputedStyleUpdatesParams starts tracking the given computed styles
// for updates. The specified array of properties replaces the one previously
// specified. Pass empty array to disable tracking. Use takeComputedStyleUpdates
// to retrieve the list of nodes that had properties modified. The changes to
// computed style properties are only tracked for nodes pushed to the front-end
// by the DOM agent. If no changes to the tracked properties occur after the
// node has been pushed to the front-end, no updates will be issued for the
// node.
type TrackComputedStyleUpdatesParams struct {
	PropertiesToTrack []*ComputedStyleProperty `json:"propertiesToTrack"`
}

// TrackComputedStyleUpdates starts tracking the given computed styles for
// updates. The specified array of properties replaces the one previously
// specified. Pass empty array to disable tracking. Use takeComputedStyleUpdates
// to retrieve the list of nodes that had properties modified. The changes to
// computed style properties are only tracked for nodes pushed to the front-end
// by the DOM agent. If no changes to the tracked properties occur after the
// node has been pushed to the front-end, no updates will be issued for the
// node.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-trackComputedStyleUpdates
//
// parameters:
//   propertiesToTrack
func TrackComputedStyleUpdates(propertiesToTrack []*ComputedStyleProperty) *TrackComputedStyleUpdatesParams {
	return &TrackComputedStyleUpdatesParams{
		PropertiesToTrack: propertiesToTrack,
	}
}

// Do executes CSS.trackComputedStyleUpdates against the provided context.
func (p *TrackComputedStyleUpdatesParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandTrackComputedStyleUpdates, p, nil)
}

// TakeComputedStyleUpdatesParams polls the next batch of computed style
// updates.
type TakeComputedStyleUpdatesParams struct{}

// TakeComputedStyleUpdates polls the next batch of computed style updates.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-takeComputedStyleUpdates
func TakeComputedStyleUpdates() *TakeComputedStyleUpdatesParams {
	return &TakeComputedStyleUpdatesParams{}
}

// TakeComputedStyleUpdatesReturns return values.
type TakeComputedStyleUpdatesReturns struct {
	NodeIDs []cdp.NodeID `json:"nodeIds,omitempty"` // The list of node Ids that have their tracked computed styles updated
}

// Do executes CSS.takeComputedStyleUpdates against the provided context.
//
// returns:
//   nodeIDs - The list of node Ids that have their tracked computed styles updated
func (p *TakeComputedStyleUpdatesParams) Do(ctx context.Context) (nodeIDs []cdp.NodeID, err error) {
	// execute
	var res TakeComputedStyleUpdatesReturns
	err = cdp.Execute(ctx, CommandTakeComputedStyleUpdates, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.NodeIDs, nil
}

// SetEffectivePropertyValueForNodeParams find a rule with the given active
// property for the given node and set the new value for this property.
type SetEffectivePropertyValueForNodeParams struct {
	NodeID       cdp.NodeID `json:"nodeId"` // The element id for which to set property.
	PropertyName string     `json:"propertyName"`
	Value        string     `json:"value"`
}

// SetEffectivePropertyValueForNode find a rule with the given active
// property for the given node and set the new value for this property.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-setEffectivePropertyValueForNode
//
// parameters:
//   nodeID - The element id for which to set property.
//   propertyName
//   value
func SetEffectivePropertyValueForNode(nodeID cdp.NodeID, propertyName string, value string) *SetEffectivePropertyValueForNodeParams {
	return &SetEffectivePropertyValueForNodeParams{
		NodeID:       nodeID,
		PropertyName: propertyName,
		Value:        value,
	}
}

// Do executes CSS.setEffectivePropertyValueForNode against the provided context.
func (p *SetEffectivePropertyValueForNodeParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetEffectivePropertyValueForNode, p, nil)
}

// SetKeyframeKeyParams modifies the keyframe rule key text.
type SetKeyframeKeyParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Range        *SourceRange `json:"range"`
	KeyText      string       `json:"keyText"`
}

// SetKeyframeKey modifies the keyframe rule key text.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-setKeyframeKey
//
// parameters:
//   styleSheetID
//   range
//   keyText
func SetKeyframeKey(styleSheetID StyleSheetID, rangeVal *SourceRange, keyText string) *SetKeyframeKeyParams {
	return &SetKeyframeKeyParams{
		StyleSheetID: styleSheetID,
		Range:        rangeVal,
		KeyText:      keyText,
	}
}

// SetKeyframeKeyReturns return values.
type SetKeyframeKeyReturns struct {
	KeyText *Value `json:"keyText,omitempty"` // The resulting key text after modification.
}

// Do executes CSS.setKeyframeKey against the provided context.
//
// returns:
//   keyText - The resulting key text after modification.
func (p *SetKeyframeKeyParams) Do(ctx context.Context) (keyText *Value, err error) {
	// execute
	var res SetKeyframeKeyReturns
	err = cdp.Execute(ctx, CommandSetKeyframeKey, p, &res)
	if err != nil {
		return nil, err
	}

	return res.KeyText, nil
}

// SetMediaTextParams modifies the rule selector.
type SetMediaTextParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Range        *SourceRange `json:"range"`
	Text         string       `json:"text"`
}

// SetMediaText modifies the rule selector.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-setMediaText
//
// parameters:
//   styleSheetID
//   range
//   text
func SetMediaText(styleSheetID StyleSheetID, rangeVal *SourceRange, text string) *SetMediaTextParams {
	return &SetMediaTextParams{
		StyleSheetID: styleSheetID,
		Range:        rangeVal,
		Text:         text,
	}
}

// SetMediaTextReturns return values.
type SetMediaTextReturns struct {
	Media *Media `json:"media,omitempty"` // The resulting CSS media rule after modification.
}

// Do executes CSS.setMediaText against the provided context.
//
// returns:
//   media - The resulting CSS media rule after modification.
func (p *SetMediaTextParams) Do(ctx context.Context) (media *Media, err error) {
	// execute
	var res SetMediaTextReturns
	err = cdp.Execute(ctx, CommandSetMediaText, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Media, nil
}

// SetContainerQueryTextParams modifies the expression of a container query.
type SetContainerQueryTextParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Range        *SourceRange `json:"range"`
	Text         string       `json:"text"`
}

// SetContainerQueryText modifies the expression of a container query.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-setContainerQueryText
//
// parameters:
//   styleSheetID
//   range
//   text
func SetContainerQueryText(styleSheetID StyleSheetID, rangeVal *SourceRange, text string) *SetContainerQueryTextParams {
	return &SetContainerQueryTextParams{
		StyleSheetID: styleSheetID,
		Range:        rangeVal,
		Text:         text,
	}
}

// SetContainerQueryTextReturns return values.
type SetContainerQueryTextReturns struct {
	ContainerQuery *ContainerQuery `json:"containerQuery,omitempty"` // The resulting CSS container query rule after modification.
}

// Do executes CSS.setContainerQueryText against the provided context.
//
// returns:
//   containerQuery - The resulting CSS container query rule after modification.
func (p *SetContainerQueryTextParams) Do(ctx context.Context) (containerQuery *ContainerQuery, err error) {
	// execute
	var res SetContainerQueryTextReturns
	err = cdp.Execute(ctx, CommandSetContainerQueryText, p, &res)
	if err != nil {
		return nil, err
	}

	return res.ContainerQuery, nil
}

// SetSupportsTextParams modifies the expression of a supports at-rule.
type SetSupportsTextParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Range        *SourceRange `json:"range"`
	Text         string       `json:"text"`
}

// SetSupportsText modifies the expression of a supports at-rule.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-setSupportsText
//
// parameters:
//   styleSheetID
//   range
//   text
func SetSupportsText(styleSheetID StyleSheetID, rangeVal *SourceRange, text string) *SetSupportsTextParams {
	return &SetSupportsTextParams{
		StyleSheetID: styleSheetID,
		Range:        rangeVal,
		Text:         text,
	}
}

// SetSupportsTextReturns return values.
type SetSupportsTextReturns struct {
	Supports *Supports `json:"supports,omitempty"` // The resulting CSS Supports rule after modification.
}

// Do executes CSS.setSupportsText against the provided context.
//
// returns:
//   supports - The resulting CSS Supports rule after modification.
func (p *SetSupportsTextParams) Do(ctx context.Context) (supports *Supports, err error) {
	// execute
	var res SetSupportsTextReturns
	err = cdp.Execute(ctx, CommandSetSupportsText, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Supports, nil
}

// SetRuleSelectorParams modifies the rule selector.
type SetRuleSelectorParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Range        *SourceRange `json:"range"`
	Selector     string       `json:"selector"`
}

// SetRuleSelector modifies the rule selector.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-setRuleSelector
//
// parameters:
//   styleSheetID
//   range
//   selector
func SetRuleSelector(styleSheetID StyleSheetID, rangeVal *SourceRange, selector string) *SetRuleSelectorParams {
	return &SetRuleSelectorParams{
		StyleSheetID: styleSheetID,
		Range:        rangeVal,
		Selector:     selector,
	}
}

// SetRuleSelectorReturns return values.
type SetRuleSelectorReturns struct {
	SelectorList *SelectorList `json:"selectorList,omitempty"` // The resulting selector list after modification.
}

// Do executes CSS.setRuleSelector against the provided context.
//
// returns:
//   selectorList - The resulting selector list after modification.
func (p *SetRuleSelectorParams) Do(ctx context.Context) (selectorList *SelectorList, err error) {
	// execute
	var res SetRuleSelectorReturns
	err = cdp.Execute(ctx, CommandSetRuleSelector, p, &res)
	if err != nil {
		return nil, err
	}

	return res.SelectorList, nil
}

// SetStyleSheetTextParams sets the new stylesheet text.
type SetStyleSheetTextParams struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"`
	Text         string       `json:"text"`
}

// SetStyleSheetText sets the new stylesheet text.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-setStyleSheetText
//
// parameters:
//   styleSheetID
//   text
func SetStyleSheetText(styleSheetID StyleSheetID, text string) *SetStyleSheetTextParams {
	return &SetStyleSheetTextParams{
		StyleSheetID: styleSheetID,
		Text:         text,
	}
}

// SetStyleSheetTextReturns return values.
type SetStyleSheetTextReturns struct {
	SourceMapURL string `json:"sourceMapURL,omitempty"` // URL of source map associated with script (if any).
}

// Do executes CSS.setStyleSheetText against the provided context.
//
// returns:
//   sourceMapURL - URL of source map associated with script (if any).
func (p *SetStyleSheetTextParams) Do(ctx context.Context) (sourceMapURL string, err error) {
	// execute
	var res SetStyleSheetTextReturns
	err = cdp.Execute(ctx, CommandSetStyleSheetText, p, &res)
	if err != nil {
		return "", err
	}

	return res.SourceMapURL, nil
}

// SetStyleTextsParams applies specified style edits one after another in the
// given order.
type SetStyleTextsParams struct {
	Edits []*StyleDeclarationEdit `json:"edits"`
}

// SetStyleTexts applies specified style edits one after another in the given
// order.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-setStyleTexts
//
// parameters:
//   edits
func SetStyleTexts(edits []*StyleDeclarationEdit) *SetStyleTextsParams {
	return &SetStyleTextsParams{
		Edits: edits,
	}
}

// SetStyleTextsReturns return values.
type SetStyleTextsReturns struct {
	Styles []*Style `json:"styles,omitempty"` // The resulting styles after modification.
}

// Do executes CSS.setStyleTexts against the provided context.
//
// returns:
//   styles - The resulting styles after modification.
func (p *SetStyleTextsParams) Do(ctx context.Context) (styles []*Style, err error) {
	// execute
	var res SetStyleTextsReturns
	err = cdp.Execute(ctx, CommandSetStyleTexts, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Styles, nil
}

// StartRuleUsageTrackingParams enables the selector recording.
type StartRuleUsageTrackingParams struct{}

// StartRuleUsageTracking enables the selector recording.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-startRuleUsageTracking
func StartRuleUsageTracking() *StartRuleUsageTrackingParams {
	return &StartRuleUsageTrackingParams{}
}

// Do executes CSS.startRuleUsageTracking against the provided context.
func (p *StartRuleUsageTrackingParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandStartRuleUsageTracking, nil, nil)
}

// StopRuleUsageTrackingParams stop tracking rule usage and return the list
// of rules that were used since last call to takeCoverageDelta (or since start
// of coverage instrumentation).
type StopRuleUsageTrackingParams struct{}

// StopRuleUsageTracking stop tracking rule usage and return the list of
// rules that were used since last call to takeCoverageDelta (or since start of
// coverage instrumentation).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-stopRuleUsageTracking
func StopRuleUsageTracking() *StopRuleUsageTrackingParams {
	return &StopRuleUsageTrackingParams{}
}

// StopRuleUsageTrackingReturns return values.
type StopRuleUsageTrackingReturns struct {
	RuleUsage []*RuleUsage `json:"ruleUsage,omitempty"`
}

// Do executes CSS.stopRuleUsageTracking against the provided context.
//
// returns:
//   ruleUsage
func (p *StopRuleUsageTrackingParams) Do(ctx context.Context) (ruleUsage []*RuleUsage, err error) {
	// execute
	var res StopRuleUsageTrackingReturns
	err = cdp.Execute(ctx, CommandStopRuleUsageTracking, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.RuleUsage, nil
}

// TakeCoverageDeltaParams obtain list of rules that became used since last
// call to this method (or since start of coverage instrumentation).
type TakeCoverageDeltaParams struct{}

// TakeCoverageDelta obtain list of rules that became used since last call to
// this method (or since start of coverage instrumentation).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-takeCoverageDelta
func TakeCoverageDelta() *TakeCoverageDeltaParams {
	return &TakeCoverageDeltaParams{}
}

// TakeCoverageDeltaReturns return values.
type TakeCoverageDeltaReturns struct {
	Coverage  []*RuleUsage `json:"coverage,omitempty"`
	Timestamp float64      `json:"timestamp,omitempty"` // Monotonically increasing time, in seconds.
}

// Do executes CSS.takeCoverageDelta against the provided context.
//
// returns:
//   coverage
//   timestamp - Monotonically increasing time, in seconds.
func (p *TakeCoverageDeltaParams) Do(ctx context.Context) (coverage []*RuleUsage, timestamp float64, err error) {
	// execute
	var res TakeCoverageDeltaReturns
	err = cdp.Execute(ctx, CommandTakeCoverageDelta, nil, &res)
	if err != nil {
		return nil, 0, err
	}

	return res.Coverage, res.Timestamp, nil
}

// SetLocalFontsEnabledParams enables/disables rendering of local CSS fonts
// (enabled by default).
type SetLocalFontsEnabledParams struct {
	Enabled bool `json:"enabled"` // Whether rendering of local fonts is enabled.
}

// SetLocalFontsEnabled enables/disables rendering of local CSS fonts
// (enabled by default).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/CSS#method-setLocalFontsEnabled
//
// parameters:
//   enabled - Whether rendering of local fonts is enabled.
func SetLocalFontsEnabled(enabled bool) *SetLocalFontsEnabledParams {
	return &SetLocalFontsEnabledParams{
		Enabled: enabled,
	}
}

// Do executes CSS.setLocalFontsEnabled against the provided context.
func (p *SetLocalFontsEnabledParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetLocalFontsEnabled, p, nil)
}

// Command names.
const (
	CommandAddRule                          = "CSS.addRule"
	CommandCollectClassNames                = "CSS.collectClassNames"
	CommandCreateStyleSheet                 = "CSS.createStyleSheet"
	CommandDisable                          = "CSS.disable"
	CommandEnable                           = "CSS.enable"
	CommandForcePseudoState                 = "CSS.forcePseudoState"
	CommandGetBackgroundColors              = "CSS.getBackgroundColors"
	CommandGetComputedStyleForNode          = "CSS.getComputedStyleForNode"
	CommandGetInlineStylesForNode           = "CSS.getInlineStylesForNode"
	CommandGetMatchedStylesForNode          = "CSS.getMatchedStylesForNode"
	CommandGetMediaQueries                  = "CSS.getMediaQueries"
	CommandGetPlatformFontsForNode          = "CSS.getPlatformFontsForNode"
	CommandGetStyleSheetText                = "CSS.getStyleSheetText"
	CommandGetLayersForNode                 = "CSS.getLayersForNode"
	CommandTrackComputedStyleUpdates        = "CSS.trackComputedStyleUpdates"
	CommandTakeComputedStyleUpdates         = "CSS.takeComputedStyleUpdates"
	CommandSetEffectivePropertyValueForNode = "CSS.setEffectivePropertyValueForNode"
	CommandSetKeyframeKey                   = "CSS.setKeyframeKey"
	CommandSetMediaText                     = "CSS.setMediaText"
	CommandSetContainerQueryText            = "CSS.setContainerQueryText"
	CommandSetSupportsText                  = "CSS.setSupportsText"
	CommandSetRuleSelector                  = "CSS.setRuleSelector"
	CommandSetStyleSheetText                = "CSS.setStyleSheetText"
	CommandSetStyleTexts                    = "CSS.setStyleTexts"
	CommandStartRuleUsageTracking           = "CSS.startRuleUsageTracking"
	CommandStopRuleUsageTracking            = "CSS.stopRuleUsageTracking"
	CommandTakeCoverageDelta                = "CSS.takeCoverageDelta"
	CommandSetLocalFontsEnabled             = "CSS.setLocalFontsEnabled"
)
