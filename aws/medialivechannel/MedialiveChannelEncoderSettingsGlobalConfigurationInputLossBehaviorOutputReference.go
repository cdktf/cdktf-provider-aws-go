package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference interface {
	cdktf.ComplexObject
	BlackFrameMsec() *float64
	SetBlackFrameMsec(val *float64)
	BlackFrameMsecInput() *float64
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
	InputLossImageColor() *string
	SetInputLossImageColor(val *string)
	InputLossImageColorInput() *string
	InputLossImageSlate() MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorInputLossImageSlateOutputReference
	InputLossImageSlateInput() *MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorInputLossImageSlate
	InputLossImageType() *string
	SetInputLossImageType(val *string)
	InputLossImageTypeInput() *string
	InternalValue() *MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehavior
	SetInternalValue(val *MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehavior)
	RepeatFrameMsec() *float64
	SetRepeatFrameMsec(val *float64)
	RepeatFrameMsecInput() *float64
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
	PutInputLossImageSlate(value *MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorInputLossImageSlate)
	ResetBlackFrameMsec()
	ResetInputLossImageColor()
	ResetInputLossImageSlate()
	ResetInputLossImageType()
	ResetRepeatFrameMsec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference
type jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) BlackFrameMsec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blackFrameMsec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) BlackFrameMsecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blackFrameMsecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) InputLossImageColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossImageColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) InputLossImageColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossImageColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) InputLossImageSlate() MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorInputLossImageSlateOutputReference {
	var returns MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorInputLossImageSlateOutputReference
	_jsii_.Get(
		j,
		"inputLossImageSlate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) InputLossImageSlateInput() *MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorInputLossImageSlate {
	var returns *MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorInputLossImageSlate
	_jsii_.Get(
		j,
		"inputLossImageSlateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) InputLossImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossImageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) InputLossImageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossImageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) InternalValue() *MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehavior {
	var returns *MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehavior
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) RepeatFrameMsec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"repeatFrameMsec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) RepeatFrameMsecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"repeatFrameMsecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference_Override(m MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference)SetBlackFrameMsec(val *float64) {
	if err := j.validateSetBlackFrameMsecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blackFrameMsec",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference)SetInputLossImageColor(val *string) {
	if err := j.validateSetInputLossImageColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputLossImageColor",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference)SetInputLossImageType(val *string) {
	if err := j.validateSetInputLossImageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputLossImageType",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference)SetInternalValue(val *MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehavior) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference)SetRepeatFrameMsec(val *float64) {
	if err := j.validateSetRepeatFrameMsecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repeatFrameMsec",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) PutInputLossImageSlate(value *MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorInputLossImageSlate) {
	if err := m.validatePutInputLossImageSlateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putInputLossImageSlate",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) ResetBlackFrameMsec() {
	_jsii_.InvokeVoid(
		m,
		"resetBlackFrameMsec",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) ResetInputLossImageColor() {
	_jsii_.InvokeVoid(
		m,
		"resetInputLossImageColor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) ResetInputLossImageSlate() {
	_jsii_.InvokeVoid(
		m,
		"resetInputLossImageSlate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) ResetInputLossImageType() {
	_jsii_.InvokeVoid(
		m,
		"resetInputLossImageType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) ResetRepeatFrameMsec() {
	_jsii_.InvokeVoid(
		m,
		"resetRepeatFrameMsec",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

