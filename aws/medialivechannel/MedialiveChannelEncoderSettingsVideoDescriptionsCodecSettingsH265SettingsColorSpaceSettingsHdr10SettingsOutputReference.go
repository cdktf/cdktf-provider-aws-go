// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10Settings
	SetInternalValue(val *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10Settings)
	MaxCll() *float64
	SetMaxCll(val *float64)
	MaxCllInput() *float64
	MaxFall() *float64
	SetMaxFall(val *float64)
	MaxFallInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetMaxCll()
	ResetMaxFall()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference
type jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) InternalValue() *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10Settings {
	var returns *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10Settings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) MaxCll() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) MaxCllInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) MaxFall() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) MaxFallInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference_Override(m MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference)SetInternalValue(val *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10Settings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference)SetMaxCll(val *float64) {
	if err := j.validateSetMaxCllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCll",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference)SetMaxFall(val *float64) {
	if err := j.validateSetMaxFallParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFall",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) ResetMaxCll() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxCll",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) ResetMaxFall() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxFall",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10SettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

