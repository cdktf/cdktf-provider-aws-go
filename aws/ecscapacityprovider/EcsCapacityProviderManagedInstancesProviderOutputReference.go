// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecscapacityprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ecscapacityprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsCapacityProviderManagedInstancesProviderOutputReference interface {
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
	InfrastructureRoleArn() *string
	SetInfrastructureRoleArn(val *string)
	InfrastructureRoleArnInput() *string
	InstanceLaunchTemplate() EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference
	InstanceLaunchTemplateInput() *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate
	InternalValue() *EcsCapacityProviderManagedInstancesProvider
	SetInternalValue(val *EcsCapacityProviderManagedInstancesProvider)
	PropagateTags() *string
	SetPropagateTags(val *string)
	PropagateTagsInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutInstanceLaunchTemplate(value *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate)
	ResetPropagateTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsCapacityProviderManagedInstancesProviderOutputReference
type jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) InfrastructureRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) InfrastructureRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) InstanceLaunchTemplate() EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference {
	var returns EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference
	_jsii_.Get(
		j,
		"instanceLaunchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) InstanceLaunchTemplateInput() *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate {
	var returns *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate
	_jsii_.Get(
		j,
		"instanceLaunchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) InternalValue() *EcsCapacityProviderManagedInstancesProvider {
	var returns *EcsCapacityProviderManagedInstancesProvider
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) PropagateTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) PropagateTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEcsCapacityProviderManagedInstancesProviderOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EcsCapacityProviderManagedInstancesProviderOutputReference {
	_init_.Initialize()

	if err := validateNewEcsCapacityProviderManagedInstancesProviderOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ecsCapacityProvider.EcsCapacityProviderManagedInstancesProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsCapacityProviderManagedInstancesProviderOutputReference_Override(e EcsCapacityProviderManagedInstancesProviderOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ecsCapacityProvider.EcsCapacityProviderManagedInstancesProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference)SetInfrastructureRoleArn(val *string) {
	if err := j.validateSetInfrastructureRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"infrastructureRoleArn",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference)SetInternalValue(val *EcsCapacityProviderManagedInstancesProvider) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference)SetPropagateTags(val *string) {
	if err := j.validateSetPropagateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) PutInstanceLaunchTemplate(value *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate) {
	if err := e.validatePutInstanceLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInstanceLaunchTemplate",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) ResetPropagateTags() {
	_jsii_.InvokeVoid(
		e,
		"resetPropagateTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

