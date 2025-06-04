// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference interface {
	cdktf.ComplexObject
	Alignment() *string
	SetAlignment(val *string)
	AlignmentInput() *string
	BackgroundColor() *string
	SetBackgroundColor(val *string)
	BackgroundColorInput() *string
	BackgroundOpacity() *float64
	SetBackgroundOpacity(val *float64)
	BackgroundOpacityInput() *float64
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Font() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsFontOutputReference
	FontColor() *string
	SetFontColor(val *string)
	FontColorInput() *string
	FontInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsFont
	FontOpacity() *float64
	SetFontOpacity(val *float64)
	FontOpacityInput() *float64
	FontResolution() *float64
	SetFontResolution(val *float64)
	FontResolutionInput() *float64
	FontSize() *string
	SetFontSize(val *string)
	FontSizeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettings
	SetInternalValue(val *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettings)
	OutlineColor() *string
	SetOutlineColor(val *string)
	OutlineColorInput() *string
	OutlineSize() *float64
	SetOutlineSize(val *float64)
	OutlineSizeInput() *float64
	ShadowColor() *string
	SetShadowColor(val *string)
	ShadowColorInput() *string
	ShadowOpacity() *float64
	SetShadowOpacity(val *float64)
	ShadowOpacityInput() *float64
	ShadowXOffset() *float64
	SetShadowXOffset(val *float64)
	ShadowXOffsetInput() *float64
	ShadowYOffset() *float64
	SetShadowYOffset(val *float64)
	ShadowYOffsetInput() *float64
	TeletextGridControl() *string
	SetTeletextGridControl(val *string)
	TeletextGridControlInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	XPosition() *float64
	SetXPosition(val *float64)
	XPositionInput() *float64
	YPosition() *float64
	SetYPosition(val *float64)
	YPositionInput() *float64
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutFont(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsFont)
	ResetAlignment()
	ResetBackgroundColor()
	ResetBackgroundOpacity()
	ResetFont()
	ResetFontColor()
	ResetFontOpacity()
	ResetFontResolution()
	ResetFontSize()
	ResetOutlineColor()
	ResetOutlineSize()
	ResetShadowColor()
	ResetShadowOpacity()
	ResetShadowXOffset()
	ResetShadowYOffset()
	ResetTeletextGridControl()
	ResetXPosition()
	ResetYPosition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference
type jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) Alignment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) AlignmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alignmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) BackgroundColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) BackgroundColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) BackgroundOpacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backgroundOpacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) BackgroundOpacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backgroundOpacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) Font() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsFontOutputReference {
	var returns MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsFontOutputReference
	_jsii_.Get(
		j,
		"font",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) FontColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) FontColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) FontInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsFont {
	var returns *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsFont
	_jsii_.Get(
		j,
		"fontInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) FontOpacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontOpacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) FontOpacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontOpacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) FontResolution() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) FontResolutionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fontResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) FontSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) FontSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) InternalValue() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettings {
	var returns *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) OutlineColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outlineColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) OutlineColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outlineColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) OutlineSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outlineSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) OutlineSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outlineSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ShadowColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shadowColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ShadowColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shadowColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ShadowOpacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowOpacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ShadowOpacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowOpacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ShadowXOffset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowXOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ShadowXOffsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowXOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ShadowYOffset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowYOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ShadowYOffsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowYOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) TeletextGridControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teletextGridControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) TeletextGridControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teletextGridControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) XPosition() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"xPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) XPositionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"xPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) YPosition() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) YPositionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yPositionInput",
		&returns,
	)
	return returns
}


func NewMedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference_Override(m MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetAlignment(val *string) {
	if err := j.validateSetAlignmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alignment",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetBackgroundColor(val *string) {
	if err := j.validateSetBackgroundColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backgroundColor",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetBackgroundOpacity(val *float64) {
	if err := j.validateSetBackgroundOpacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backgroundOpacity",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetFontColor(val *string) {
	if err := j.validateSetFontColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontColor",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetFontOpacity(val *float64) {
	if err := j.validateSetFontOpacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontOpacity",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetFontResolution(val *float64) {
	if err := j.validateSetFontResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontResolution",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetFontSize(val *string) {
	if err := j.validateSetFontSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontSize",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetInternalValue(val *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetOutlineColor(val *string) {
	if err := j.validateSetOutlineColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outlineColor",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetOutlineSize(val *float64) {
	if err := j.validateSetOutlineSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outlineSize",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetShadowColor(val *string) {
	if err := j.validateSetShadowColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowColor",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetShadowOpacity(val *float64) {
	if err := j.validateSetShadowOpacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowOpacity",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetShadowXOffset(val *float64) {
	if err := j.validateSetShadowXOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowXOffset",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetShadowYOffset(val *float64) {
	if err := j.validateSetShadowYOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowYOffset",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetTeletextGridControl(val *string) {
	if err := j.validateSetTeletextGridControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teletextGridControl",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetXPosition(val *float64) {
	if err := j.validateSetXPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xPosition",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference)SetYPosition(val *float64) {
	if err := j.validateSetYPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yPosition",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) PutFont(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsFont) {
	if err := m.validatePutFontParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFont",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ResetAlignment() {
	_jsii_.InvokeVoid(
		m,
		"resetAlignment",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ResetBackgroundColor() {
	_jsii_.InvokeVoid(
		m,
		"resetBackgroundColor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ResetBackgroundOpacity() {
	_jsii_.InvokeVoid(
		m,
		"resetBackgroundOpacity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ResetFont() {
	_jsii_.InvokeVoid(
		m,
		"resetFont",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ResetFontColor() {
	_jsii_.InvokeVoid(
		m,
		"resetFontColor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ResetFontOpacity() {
	_jsii_.InvokeVoid(
		m,
		"resetFontOpacity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ResetFontResolution() {
	_jsii_.InvokeVoid(
		m,
		"resetFontResolution",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ResetFontSize() {
	_jsii_.InvokeVoid(
		m,
		"resetFontSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ResetOutlineColor() {
	_jsii_.InvokeVoid(
		m,
		"resetOutlineColor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ResetOutlineSize() {
	_jsii_.InvokeVoid(
		m,
		"resetOutlineSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ResetShadowColor() {
	_jsii_.InvokeVoid(
		m,
		"resetShadowColor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ResetShadowOpacity() {
	_jsii_.InvokeVoid(
		m,
		"resetShadowOpacity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ResetShadowXOffset() {
	_jsii_.InvokeVoid(
		m,
		"resetShadowXOffset",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ResetShadowYOffset() {
	_jsii_.InvokeVoid(
		m,
		"resetShadowYOffset",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ResetTeletextGridControl() {
	_jsii_.InvokeVoid(
		m,
		"resetTeletextGridControl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ResetXPosition() {
	_jsii_.InvokeVoid(
		m,
		"resetXPosition",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ResetYPosition() {
	_jsii_.InvokeVoid(
		m,
		"resetYPosition",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

