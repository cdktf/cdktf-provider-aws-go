package mskconnectconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v13/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v13/mskconnectconnector/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MskconnectConnectorCapacityAutoscalingOutputReference interface {
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
	InternalValue() *MskconnectConnectorCapacityAutoscaling
	SetInternalValue(val *MskconnectConnectorCapacityAutoscaling)
	MaxWorkerCount() *float64
	SetMaxWorkerCount(val *float64)
	MaxWorkerCountInput() *float64
	McuCount() *float64
	SetMcuCount(val *float64)
	McuCountInput() *float64
	MinWorkerCount() *float64
	SetMinWorkerCount(val *float64)
	MinWorkerCountInput() *float64
	ScaleInPolicy() MskconnectConnectorCapacityAutoscalingScaleInPolicyOutputReference
	ScaleInPolicyInput() *MskconnectConnectorCapacityAutoscalingScaleInPolicy
	ScaleOutPolicy() MskconnectConnectorCapacityAutoscalingScaleOutPolicyOutputReference
	ScaleOutPolicyInput() *MskconnectConnectorCapacityAutoscalingScaleOutPolicy
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
	PutScaleInPolicy(value *MskconnectConnectorCapacityAutoscalingScaleInPolicy)
	PutScaleOutPolicy(value *MskconnectConnectorCapacityAutoscalingScaleOutPolicy)
	ResetMcuCount()
	ResetScaleInPolicy()
	ResetScaleOutPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MskconnectConnectorCapacityAutoscalingOutputReference
type jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) InternalValue() *MskconnectConnectorCapacityAutoscaling {
	var returns *MskconnectConnectorCapacityAutoscaling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) MaxWorkerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) MaxWorkerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkerCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) McuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mcuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) McuCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mcuCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) MinWorkerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWorkerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) MinWorkerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWorkerCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) ScaleInPolicy() MskconnectConnectorCapacityAutoscalingScaleInPolicyOutputReference {
	var returns MskconnectConnectorCapacityAutoscalingScaleInPolicyOutputReference
	_jsii_.Get(
		j,
		"scaleInPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) ScaleInPolicyInput() *MskconnectConnectorCapacityAutoscalingScaleInPolicy {
	var returns *MskconnectConnectorCapacityAutoscalingScaleInPolicy
	_jsii_.Get(
		j,
		"scaleInPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) ScaleOutPolicy() MskconnectConnectorCapacityAutoscalingScaleOutPolicyOutputReference {
	var returns MskconnectConnectorCapacityAutoscalingScaleOutPolicyOutputReference
	_jsii_.Get(
		j,
		"scaleOutPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) ScaleOutPolicyInput() *MskconnectConnectorCapacityAutoscalingScaleOutPolicy {
	var returns *MskconnectConnectorCapacityAutoscalingScaleOutPolicy
	_jsii_.Get(
		j,
		"scaleOutPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMskconnectConnectorCapacityAutoscalingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MskconnectConnectorCapacityAutoscalingOutputReference {
	_init_.Initialize()

	if err := validateNewMskconnectConnectorCapacityAutoscalingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.mskconnectConnector.MskconnectConnectorCapacityAutoscalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMskconnectConnectorCapacityAutoscalingOutputReference_Override(m MskconnectConnectorCapacityAutoscalingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.mskconnectConnector.MskconnectConnectorCapacityAutoscalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference)SetInternalValue(val *MskconnectConnectorCapacityAutoscaling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference)SetMaxWorkerCount(val *float64) {
	if err := j.validateSetMaxWorkerCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWorkerCount",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference)SetMcuCount(val *float64) {
	if err := j.validateSetMcuCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mcuCount",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference)SetMinWorkerCount(val *float64) {
	if err := j.validateSetMinWorkerCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minWorkerCount",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) PutScaleInPolicy(value *MskconnectConnectorCapacityAutoscalingScaleInPolicy) {
	if err := m.validatePutScaleInPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putScaleInPolicy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) PutScaleOutPolicy(value *MskconnectConnectorCapacityAutoscalingScaleOutPolicy) {
	if err := m.validatePutScaleOutPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putScaleOutPolicy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) ResetMcuCount() {
	_jsii_.InvokeVoid(
		m,
		"resetMcuCount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) ResetScaleInPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetScaleInPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) ResetScaleOutPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetScaleOutPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MskconnectConnectorCapacityAutoscalingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

