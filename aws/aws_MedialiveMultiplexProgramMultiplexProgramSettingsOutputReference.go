// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference interface {
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
	InternalValue() *MedialiveMultiplexProgramMultiplexProgramSettings
	SetInternalValue(val *MedialiveMultiplexProgramMultiplexProgramSettings)
	PreferredChannelPipeline() *string
	SetPreferredChannelPipeline(val *string)
	PreferredChannelPipelineInput() *string
	ProgramNumber() *float64
	SetProgramNumber(val *float64)
	ProgramNumberInput() *float64
	ServiceDescriptor() MedialiveMultiplexProgramMultiplexProgramSettingsServiceDescriptorOutputReference
	ServiceDescriptorInput() *MedialiveMultiplexProgramMultiplexProgramSettingsServiceDescriptor
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VideoSettings() MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsOutputReference
	VideoSettingsInput() *MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettings
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
	PutServiceDescriptor(value *MedialiveMultiplexProgramMultiplexProgramSettingsServiceDescriptor)
	PutVideoSettings(value *MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettings)
	ResetServiceDescriptor()
	ResetVideoSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference
type jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) InternalValue() *MedialiveMultiplexProgramMultiplexProgramSettings {
	var returns *MedialiveMultiplexProgramMultiplexProgramSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) PreferredChannelPipeline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredChannelPipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) PreferredChannelPipelineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredChannelPipelineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) ProgramNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) ProgramNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) ServiceDescriptor() MedialiveMultiplexProgramMultiplexProgramSettingsServiceDescriptorOutputReference {
	var returns MedialiveMultiplexProgramMultiplexProgramSettingsServiceDescriptorOutputReference
	_jsii_.Get(
		j,
		"serviceDescriptor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) ServiceDescriptorInput() *MedialiveMultiplexProgramMultiplexProgramSettingsServiceDescriptor {
	var returns *MedialiveMultiplexProgramMultiplexProgramSettingsServiceDescriptor
	_jsii_.Get(
		j,
		"serviceDescriptorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) VideoSettings() MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsOutputReference {
	var returns MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsOutputReference
	_jsii_.Get(
		j,
		"videoSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) VideoSettingsInput() *MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettings {
	var returns *MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettings
	_jsii_.Get(
		j,
		"videoSettingsInput",
		&returns,
	)
	return returns
}


func NewMedialiveMultiplexProgramMultiplexProgramSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveMultiplexProgramMultiplexProgramSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveMultiplexProgramMultiplexProgramSettingsOutputReference_Override(m MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference)SetInternalValue(val *MedialiveMultiplexProgramMultiplexProgramSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference)SetPreferredChannelPipeline(val *string) {
	if err := j.validateSetPreferredChannelPipelineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredChannelPipeline",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference)SetProgramNumber(val *float64) {
	if err := j.validateSetProgramNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programNumber",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) PutServiceDescriptor(value *MedialiveMultiplexProgramMultiplexProgramSettingsServiceDescriptor) {
	if err := m.validatePutServiceDescriptorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putServiceDescriptor",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) PutVideoSettings(value *MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettings) {
	if err := m.validatePutVideoSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putVideoSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) ResetServiceDescriptor() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceDescriptor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) ResetVideoSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetVideoSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

