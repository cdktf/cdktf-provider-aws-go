// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ecstaskdefinition/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference interface {
	cdktf.ComplexObject
	AuthorizationConfig() EcsTaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfigOutputReference
	AuthorizationConfigInput() *EcsTaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfig
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
	FileSystemId() *string
	SetFileSystemId(val *string)
	FileSystemIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *EcsTaskDefinitionVolumeEfsVolumeConfiguration
	SetInternalValue(val *EcsTaskDefinitionVolumeEfsVolumeConfiguration)
	RootDirectory() *string
	SetRootDirectory(val *string)
	RootDirectoryInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransitEncryption() *string
	SetTransitEncryption(val *string)
	TransitEncryptionInput() *string
	TransitEncryptionPort() *float64
	SetTransitEncryptionPort(val *float64)
	TransitEncryptionPortInput() *float64
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
	PutAuthorizationConfig(value *EcsTaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfig)
	ResetAuthorizationConfig()
	ResetRootDirectory()
	ResetTransitEncryption()
	ResetTransitEncryptionPort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference
type jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) AuthorizationConfig() EcsTaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfigOutputReference {
	var returns EcsTaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfigOutputReference
	_jsii_.Get(
		j,
		"authorizationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) AuthorizationConfigInput() *EcsTaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfig {
	var returns *EcsTaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfig
	_jsii_.Get(
		j,
		"authorizationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) FileSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) InternalValue() *EcsTaskDefinitionVolumeEfsVolumeConfiguration {
	var returns *EcsTaskDefinitionVolumeEfsVolumeConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) RootDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) RootDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) TransitEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) TransitEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) TransitEncryptionPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transitEncryptionPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) TransitEncryptionPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transitEncryptionPortInput",
		&returns,
	)
	return returns
}


func NewEcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewEcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ecsTaskDefinition.EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference_Override(e EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ecsTaskDefinition.EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference)SetFileSystemId(val *string) {
	if err := j.validateSetFileSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemId",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference)SetInternalValue(val *EcsTaskDefinitionVolumeEfsVolumeConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference)SetRootDirectory(val *string) {
	if err := j.validateSetRootDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootDirectory",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference)SetTransitEncryption(val *string) {
	if err := j.validateSetTransitEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitEncryption",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference)SetTransitEncryptionPort(val *float64) {
	if err := j.validateSetTransitEncryptionPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitEncryptionPort",
		val,
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) PutAuthorizationConfig(value *EcsTaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfig) {
	if err := e.validatePutAuthorizationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAuthorizationConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) ResetAuthorizationConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetAuthorizationConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) ResetRootDirectory() {
	_jsii_.InvokeVoid(
		e,
		"resetRootDirectory",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) ResetTransitEncryption() {
	_jsii_.InvokeVoid(
		e,
		"resetTransitEncryption",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) ResetTransitEncryptionPort() {
	_jsii_.InvokeVoid(
		e,
		"resetTransitEncryptionPort",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcsTaskDefinitionVolumeEfsVolumeConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

