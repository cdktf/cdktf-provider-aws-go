// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveMultiplexMultiplexSettingsOutputReference interface {
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
	InternalValue() *MedialiveMultiplexMultiplexSettings
	SetInternalValue(val *MedialiveMultiplexMultiplexSettings)
	MaximumVideoBufferDelayMilliseconds() *float64
	SetMaximumVideoBufferDelayMilliseconds(val *float64)
	MaximumVideoBufferDelayMillisecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransportStreamBitrate() *float64
	SetTransportStreamBitrate(val *float64)
	TransportStreamBitrateInput() *float64
	TransportStreamId() *float64
	SetTransportStreamId(val *float64)
	TransportStreamIdInput() *float64
	TransportStreamReservedBitrate() *float64
	SetTransportStreamReservedBitrate(val *float64)
	TransportStreamReservedBitrateInput() *float64
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
	ResetMaximumVideoBufferDelayMilliseconds()
	ResetTransportStreamReservedBitrate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveMultiplexMultiplexSettingsOutputReference
type jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) InternalValue() *MedialiveMultiplexMultiplexSettings {
	var returns *MedialiveMultiplexMultiplexSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) MaximumVideoBufferDelayMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumVideoBufferDelayMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) MaximumVideoBufferDelayMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumVideoBufferDelayMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) TransportStreamBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) TransportStreamBitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamBitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) TransportStreamId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) TransportStreamIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) TransportStreamReservedBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamReservedBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) TransportStreamReservedBitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamReservedBitrateInput",
		&returns,
	)
	return returns
}


func NewMedialiveMultiplexMultiplexSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveMultiplexMultiplexSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveMultiplexMultiplexSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.MedialiveMultiplexMultiplexSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveMultiplexMultiplexSettingsOutputReference_Override(m MedialiveMultiplexMultiplexSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.MedialiveMultiplexMultiplexSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference)SetInternalValue(val *MedialiveMultiplexMultiplexSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference)SetMaximumVideoBufferDelayMilliseconds(val *float64) {
	if err := j.validateSetMaximumVideoBufferDelayMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumVideoBufferDelayMilliseconds",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference)SetTransportStreamBitrate(val *float64) {
	if err := j.validateSetTransportStreamBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportStreamBitrate",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference)SetTransportStreamId(val *float64) {
	if err := j.validateSetTransportStreamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportStreamId",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference)SetTransportStreamReservedBitrate(val *float64) {
	if err := j.validateSetTransportStreamReservedBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportStreamReservedBitrate",
		val,
	)
}

func (m *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) ResetMaximumVideoBufferDelayMilliseconds() {
	_jsii_.InvokeVoid(
		m,
		"resetMaximumVideoBufferDelayMilliseconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) ResetTransportStreamReservedBitrate() {
	_jsii_.InvokeVoid(
		m,
		"resetTransportStreamReservedBitrate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveMultiplexMultiplexSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

