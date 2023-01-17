package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference interface {
	cdktf.ComplexObject
	BlackDetectThreshold() *float64
	SetBlackDetectThreshold(val *float64)
	BlackDetectThresholdInput() *float64
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
	InternalValue() *MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettings
	SetInternalValue(val *MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettings)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VideoBlackThresholdMsec() *float64
	SetVideoBlackThresholdMsec(val *float64)
	VideoBlackThresholdMsecInput() *float64
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
	ResetBlackDetectThreshold()
	ResetVideoBlackThresholdMsec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference
type jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) BlackDetectThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blackDetectThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) BlackDetectThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blackDetectThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) InternalValue() *MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettings {
	var returns *MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) VideoBlackThresholdMsec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"videoBlackThresholdMsec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) VideoBlackThresholdMsecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"videoBlackThresholdMsecInput",
		&returns,
	)
	return returns
}


func NewMedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference_Override(m MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference)SetBlackDetectThreshold(val *float64) {
	if err := j.validateSetBlackDetectThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blackDetectThreshold",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference)SetInternalValue(val *MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference)SetVideoBlackThresholdMsec(val *float64) {
	if err := j.validateSetVideoBlackThresholdMsecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"videoBlackThresholdMsec",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) ResetBlackDetectThreshold() {
	_jsii_.InvokeVoid(
		m,
		"resetBlackDetectThreshold",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) ResetVideoBlackThresholdMsec() {
	_jsii_.InvokeVoid(
		m,
		"resetVideoBlackThresholdMsec",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

