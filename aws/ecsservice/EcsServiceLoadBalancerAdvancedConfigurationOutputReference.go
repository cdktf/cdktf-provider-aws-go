// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ecsservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsServiceLoadBalancerAdvancedConfigurationOutputReference interface {
	cdktf.ComplexObject
	AlternateTargetGroupArn() *string
	SetAlternateTargetGroupArn(val *string)
	AlternateTargetGroupArnInput() *string
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
	InternalValue() *EcsServiceLoadBalancerAdvancedConfiguration
	SetInternalValue(val *EcsServiceLoadBalancerAdvancedConfiguration)
	ProductionListenerRule() *string
	SetProductionListenerRule(val *string)
	ProductionListenerRuleInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TestListenerRule() *string
	SetTestListenerRule(val *string)
	TestListenerRuleInput() *string
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
	ResetTestListenerRule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsServiceLoadBalancerAdvancedConfigurationOutputReference
type jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) AlternateTargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alternateTargetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) AlternateTargetGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alternateTargetGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) InternalValue() *EcsServiceLoadBalancerAdvancedConfiguration {
	var returns *EcsServiceLoadBalancerAdvancedConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) ProductionListenerRule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productionListenerRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) ProductionListenerRuleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productionListenerRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) TestListenerRule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testListenerRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) TestListenerRuleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testListenerRuleInput",
		&returns,
	)
	return returns
}


func NewEcsServiceLoadBalancerAdvancedConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EcsServiceLoadBalancerAdvancedConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewEcsServiceLoadBalancerAdvancedConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ecsService.EcsServiceLoadBalancerAdvancedConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsServiceLoadBalancerAdvancedConfigurationOutputReference_Override(e EcsServiceLoadBalancerAdvancedConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ecsService.EcsServiceLoadBalancerAdvancedConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference)SetAlternateTargetGroupArn(val *string) {
	if err := j.validateSetAlternateTargetGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alternateTargetGroupArn",
		val,
	)
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference)SetInternalValue(val *EcsServiceLoadBalancerAdvancedConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference)SetProductionListenerRule(val *string) {
	if err := j.validateSetProductionListenerRuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productionListenerRule",
		val,
	)
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference)SetTestListenerRule(val *string) {
	if err := j.validateSetTestListenerRuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"testListenerRule",
		val,
	)
}

func (e *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) ResetTestListenerRule() {
	_jsii_.InvokeVoid(
		e,
		"resetTestListenerRule",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceLoadBalancerAdvancedConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

