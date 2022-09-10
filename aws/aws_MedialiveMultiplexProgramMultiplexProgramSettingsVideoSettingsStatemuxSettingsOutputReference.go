// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference interface {
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
	InternalValue() *MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettings
	SetInternalValue(val *MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettings)
	MaximumBitrate() *float64
	SetMaximumBitrate(val *float64)
	MaximumBitrateInput() *float64
	MinimumBitrate() *float64
	SetMinimumBitrate(val *float64)
	MinimumBitrateInput() *float64
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
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
	ResetMaximumBitrate()
	ResetMinimumBitrate()
	ResetPriority()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference
type jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) InternalValue() *MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettings {
	var returns *MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) MaximumBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) MaximumBitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) MinimumBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) MinimumBitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumBitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference_Override(m MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference)SetInternalValue(val *MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference)SetMaximumBitrate(val *float64) {
	if err := j.validateSetMaximumBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBitrate",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference)SetMinimumBitrate(val *float64) {
	if err := j.validateSetMinimumBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumBitrate",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) ResetMaximumBitrate() {
	_jsii_.InvokeVoid(
		m,
		"resetMaximumBitrate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) ResetMinimumBitrate() {
	_jsii_.InvokeVoid(
		m,
		"resetMinimumBitrate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		m,
		"resetPriority",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

