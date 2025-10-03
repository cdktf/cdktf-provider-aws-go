// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecscapacityprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ecscapacityprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference interface {
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
	Ec2InstanceProfileArn() *string
	SetEc2InstanceProfileArn(val *string)
	Ec2InstanceProfileArnInput() *string
	// Experimental.
	Fqn() *string
	InstanceRequirements() EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirementsOutputReference
	InstanceRequirementsInput() *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirements
	InternalValue() *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate
	SetInternalValue(val *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate)
	Monitoring() *string
	SetMonitoring(val *string)
	MonitoringInput() *string
	NetworkConfiguration() EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateNetworkConfigurationOutputReference
	NetworkConfigurationInput() *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateNetworkConfiguration
	StorageConfiguration() EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateStorageConfigurationOutputReference
	StorageConfigurationInput() *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateStorageConfiguration
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
	PutInstanceRequirements(value *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirements)
	PutNetworkConfiguration(value *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateNetworkConfiguration)
	PutStorageConfiguration(value *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateStorageConfiguration)
	ResetInstanceRequirements()
	ResetMonitoring()
	ResetStorageConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference
type jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) Ec2InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) Ec2InstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) InstanceRequirements() EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirementsOutputReference {
	var returns EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirementsOutputReference
	_jsii_.Get(
		j,
		"instanceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) InstanceRequirementsInput() *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirements {
	var returns *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirements
	_jsii_.Get(
		j,
		"instanceRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) InternalValue() *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate {
	var returns *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) Monitoring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) MonitoringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) NetworkConfiguration() EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateNetworkConfigurationOutputReference {
	var returns EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) NetworkConfigurationInput() *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateNetworkConfiguration {
	var returns *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateNetworkConfiguration
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) StorageConfiguration() EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateStorageConfigurationOutputReference {
	var returns EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateStorageConfigurationOutputReference
	_jsii_.Get(
		j,
		"storageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) StorageConfigurationInput() *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateStorageConfiguration {
	var returns *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateStorageConfiguration
	_jsii_.Get(
		j,
		"storageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference {
	_init_.Initialize()

	if err := validateNewEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ecsCapacityProvider.EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference_Override(e EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ecsCapacityProvider.EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference)SetEc2InstanceProfileArn(val *string) {
	if err := j.validateSetEc2InstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2InstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference)SetInternalValue(val *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference)SetMonitoring(val *string) {
	if err := j.validateSetMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoring",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) PutInstanceRequirements(value *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirements) {
	if err := e.validatePutInstanceRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInstanceRequirements",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) PutNetworkConfiguration(value *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateNetworkConfiguration) {
	if err := e.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) PutStorageConfiguration(value *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateStorageConfiguration) {
	if err := e.validatePutStorageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStorageConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) ResetInstanceRequirements() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceRequirements",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) ResetMonitoring() {
	_jsii_.InvokeVoid(
		e,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) ResetStorageConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetStorageConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

