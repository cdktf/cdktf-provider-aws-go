// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference interface {
	cdktf.ComplexObject
	AribDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsAribDestinationSettingsOutputReference
	AribDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsAribDestinationSettings
	BurnInDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsOutputReference
	BurnInDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettings
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
	DvbSubDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference
	DvbSubDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettings
	EbuTtDDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEbuTtDDestinationSettingsOutputReference
	EbuTtDDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEbuTtDDestinationSettings
	EmbeddedDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedDestinationSettingsOutputReference
	EmbeddedDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedDestinationSettings
	EmbeddedPlusScte20DestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedPlusScte20DestinationSettingsOutputReference
	EmbeddedPlusScte20DestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedPlusScte20DestinationSettings
	// Experimental.
	Fqn() *string
	InternalValue() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettings
	SetInternalValue(val *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettings)
	RtmpCaptionInfoDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsRtmpCaptionInfoDestinationSettingsOutputReference
	RtmpCaptionInfoDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsRtmpCaptionInfoDestinationSettings
	Scte20PlusEmbeddedDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte20PlusEmbeddedDestinationSettingsOutputReference
	Scte20PlusEmbeddedDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte20PlusEmbeddedDestinationSettings
	Scte27DestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte27DestinationSettingsOutputReference
	Scte27DestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte27DestinationSettings
	SmpteTtDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsSmpteTtDestinationSettingsOutputReference
	SmpteTtDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsSmpteTtDestinationSettings
	TeletextDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTeletextDestinationSettingsOutputReference
	TeletextDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTeletextDestinationSettings
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TtmlDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTtmlDestinationSettingsOutputReference
	TtmlDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTtmlDestinationSettings
	WebvttDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsWebvttDestinationSettingsOutputReference
	WebvttDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsWebvttDestinationSettings
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
	PutAribDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsAribDestinationSettings)
	PutBurnInDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettings)
	PutDvbSubDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettings)
	PutEbuTtDDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEbuTtDDestinationSettings)
	PutEmbeddedDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedDestinationSettings)
	PutEmbeddedPlusScte20DestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedPlusScte20DestinationSettings)
	PutRtmpCaptionInfoDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsRtmpCaptionInfoDestinationSettings)
	PutScte20PlusEmbeddedDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte20PlusEmbeddedDestinationSettings)
	PutScte27DestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte27DestinationSettings)
	PutSmpteTtDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsSmpteTtDestinationSettings)
	PutTeletextDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTeletextDestinationSettings)
	PutTtmlDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTtmlDestinationSettings)
	PutWebvttDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsWebvttDestinationSettings)
	ResetAribDestinationSettings()
	ResetBurnInDestinationSettings()
	ResetDvbSubDestinationSettings()
	ResetEbuTtDDestinationSettings()
	ResetEmbeddedDestinationSettings()
	ResetEmbeddedPlusScte20DestinationSettings()
	ResetRtmpCaptionInfoDestinationSettings()
	ResetScte20PlusEmbeddedDestinationSettings()
	ResetScte27DestinationSettings()
	ResetSmpteTtDestinationSettings()
	ResetTeletextDestinationSettings()
	ResetTtmlDestinationSettings()
	ResetWebvttDestinationSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference
type jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) AribDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsAribDestinationSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsAribDestinationSettingsOutputReference
	_jsii_.Get(
		j,
		"aribDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) AribDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsAribDestinationSettings {
	var returns *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsAribDestinationSettings
	_jsii_.Get(
		j,
		"aribDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) BurnInDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsOutputReference
	_jsii_.Get(
		j,
		"burnInDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) BurnInDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettings {
	var returns *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettings
	_jsii_.Get(
		j,
		"burnInDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) DvbSubDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsOutputReference
	_jsii_.Get(
		j,
		"dvbSubDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) DvbSubDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettings {
	var returns *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettings
	_jsii_.Get(
		j,
		"dvbSubDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) EbuTtDDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEbuTtDDestinationSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEbuTtDDestinationSettingsOutputReference
	_jsii_.Get(
		j,
		"ebuTtDDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) EbuTtDDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEbuTtDDestinationSettings {
	var returns *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEbuTtDDestinationSettings
	_jsii_.Get(
		j,
		"ebuTtDDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) EmbeddedDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedDestinationSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedDestinationSettingsOutputReference
	_jsii_.Get(
		j,
		"embeddedDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) EmbeddedDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedDestinationSettings {
	var returns *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedDestinationSettings
	_jsii_.Get(
		j,
		"embeddedDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) EmbeddedPlusScte20DestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedPlusScte20DestinationSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedPlusScte20DestinationSettingsOutputReference
	_jsii_.Get(
		j,
		"embeddedPlusScte20DestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) EmbeddedPlusScte20DestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedPlusScte20DestinationSettings {
	var returns *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedPlusScte20DestinationSettings
	_jsii_.Get(
		j,
		"embeddedPlusScte20DestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) InternalValue() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettings {
	var returns *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) RtmpCaptionInfoDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsRtmpCaptionInfoDestinationSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsRtmpCaptionInfoDestinationSettingsOutputReference
	_jsii_.Get(
		j,
		"rtmpCaptionInfoDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) RtmpCaptionInfoDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsRtmpCaptionInfoDestinationSettings {
	var returns *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsRtmpCaptionInfoDestinationSettings
	_jsii_.Get(
		j,
		"rtmpCaptionInfoDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) Scte20PlusEmbeddedDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte20PlusEmbeddedDestinationSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte20PlusEmbeddedDestinationSettingsOutputReference
	_jsii_.Get(
		j,
		"scte20PlusEmbeddedDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) Scte20PlusEmbeddedDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte20PlusEmbeddedDestinationSettings {
	var returns *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte20PlusEmbeddedDestinationSettings
	_jsii_.Get(
		j,
		"scte20PlusEmbeddedDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) Scte27DestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte27DestinationSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte27DestinationSettingsOutputReference
	_jsii_.Get(
		j,
		"scte27DestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) Scte27DestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte27DestinationSettings {
	var returns *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte27DestinationSettings
	_jsii_.Get(
		j,
		"scte27DestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) SmpteTtDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsSmpteTtDestinationSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsSmpteTtDestinationSettingsOutputReference
	_jsii_.Get(
		j,
		"smpteTtDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) SmpteTtDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsSmpteTtDestinationSettings {
	var returns *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsSmpteTtDestinationSettings
	_jsii_.Get(
		j,
		"smpteTtDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) TeletextDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTeletextDestinationSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTeletextDestinationSettingsOutputReference
	_jsii_.Get(
		j,
		"teletextDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) TeletextDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTeletextDestinationSettings {
	var returns *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTeletextDestinationSettings
	_jsii_.Get(
		j,
		"teletextDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) TtmlDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTtmlDestinationSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTtmlDestinationSettingsOutputReference
	_jsii_.Get(
		j,
		"ttmlDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) TtmlDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTtmlDestinationSettings {
	var returns *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTtmlDestinationSettings
	_jsii_.Get(
		j,
		"ttmlDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) WebvttDestinationSettings() MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsWebvttDestinationSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsWebvttDestinationSettingsOutputReference
	_jsii_.Get(
		j,
		"webvttDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) WebvttDestinationSettingsInput() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsWebvttDestinationSettings {
	var returns *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsWebvttDestinationSettings
	_jsii_.Get(
		j,
		"webvttDestinationSettingsInput",
		&returns,
	)
	return returns
}


func NewMedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference_Override(m MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference)SetInternalValue(val *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) PutAribDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsAribDestinationSettings) {
	if err := m.validatePutAribDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAribDestinationSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) PutBurnInDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettings) {
	if err := m.validatePutBurnInDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBurnInDestinationSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) PutDvbSubDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettings) {
	if err := m.validatePutDvbSubDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDvbSubDestinationSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) PutEbuTtDDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEbuTtDDestinationSettings) {
	if err := m.validatePutEbuTtDDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEbuTtDDestinationSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) PutEmbeddedDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedDestinationSettings) {
	if err := m.validatePutEmbeddedDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEmbeddedDestinationSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) PutEmbeddedPlusScte20DestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsEmbeddedPlusScte20DestinationSettings) {
	if err := m.validatePutEmbeddedPlusScte20DestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEmbeddedPlusScte20DestinationSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) PutRtmpCaptionInfoDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsRtmpCaptionInfoDestinationSettings) {
	if err := m.validatePutRtmpCaptionInfoDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRtmpCaptionInfoDestinationSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) PutScte20PlusEmbeddedDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte20PlusEmbeddedDestinationSettings) {
	if err := m.validatePutScte20PlusEmbeddedDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putScte20PlusEmbeddedDestinationSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) PutScte27DestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsScte27DestinationSettings) {
	if err := m.validatePutScte27DestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putScte27DestinationSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) PutSmpteTtDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsSmpteTtDestinationSettings) {
	if err := m.validatePutSmpteTtDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSmpteTtDestinationSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) PutTeletextDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTeletextDestinationSettings) {
	if err := m.validatePutTeletextDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTeletextDestinationSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) PutTtmlDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsTtmlDestinationSettings) {
	if err := m.validatePutTtmlDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTtmlDestinationSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) PutWebvttDestinationSettings(value *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsWebvttDestinationSettings) {
	if err := m.validatePutWebvttDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putWebvttDestinationSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) ResetAribDestinationSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetAribDestinationSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) ResetBurnInDestinationSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetBurnInDestinationSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) ResetDvbSubDestinationSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetDvbSubDestinationSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) ResetEbuTtDDestinationSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetEbuTtDDestinationSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) ResetEmbeddedDestinationSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetEmbeddedDestinationSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) ResetEmbeddedPlusScte20DestinationSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetEmbeddedPlusScte20DestinationSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) ResetRtmpCaptionInfoDestinationSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetRtmpCaptionInfoDestinationSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) ResetScte20PlusEmbeddedDestinationSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetScte20PlusEmbeddedDestinationSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) ResetScte27DestinationSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetScte27DestinationSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) ResetSmpteTtDestinationSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetSmpteTtDestinationSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) ResetTeletextDestinationSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetTeletextDestinationSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) ResetTtmlDestinationSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetTtmlDestinationSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) ResetWebvttDestinationSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetWebvttDestinationSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

