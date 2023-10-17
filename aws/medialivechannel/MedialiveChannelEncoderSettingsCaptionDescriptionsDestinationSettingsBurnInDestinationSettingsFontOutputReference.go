// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference interface {
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
	InternalValue() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFont
	SetInternalValue(val *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFont)
	PasswordParam() *string
	SetPasswordParam(val *string)
	PasswordParamInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() *string
	SetUri(val *string)
	UriInput() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	ResetPasswordParam()
	ResetUsername()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference
type jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) InternalValue() *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFont {
	var returns *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFont
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) PasswordParam() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) PasswordParamInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordParamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewMedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference_Override(m MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference)SetInternalValue(val *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFont) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference)SetPasswordParam(val *string) {
	if err := j.validateSetPasswordParamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordParam",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) ResetPasswordParam() {
	_jsii_.InvokeVoid(
		m,
		"resetPasswordParam",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		m,
		"resetUsername",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettingsBurnInDestinationSettingsFontOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

