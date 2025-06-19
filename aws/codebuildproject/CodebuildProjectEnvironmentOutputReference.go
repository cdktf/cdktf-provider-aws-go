// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/codebuildproject/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodebuildProjectEnvironmentOutputReference interface {
	cdktf.ComplexObject
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
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
	ComputeType() *string
	SetComputeType(val *string)
	ComputeTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnvironmentVariable() CodebuildProjectEnvironmentEnvironmentVariableList
	EnvironmentVariableInput() interface{}
	Fleet() CodebuildProjectEnvironmentFleetOutputReference
	FleetInput() *CodebuildProjectEnvironmentFleet
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	ImagePullCredentialsType() *string
	SetImagePullCredentialsType(val *string)
	ImagePullCredentialsTypeInput() *string
	InternalValue() *CodebuildProjectEnvironment
	SetInternalValue(val *CodebuildProjectEnvironment)
	PrivilegedMode() interface{}
	SetPrivilegedMode(val interface{})
	PrivilegedModeInput() interface{}
	RegistryCredential() CodebuildProjectEnvironmentRegistryCredentialOutputReference
	RegistryCredentialInput() *CodebuildProjectEnvironmentRegistryCredential
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutEnvironmentVariable(value interface{})
	PutFleet(value *CodebuildProjectEnvironmentFleet)
	PutRegistryCredential(value *CodebuildProjectEnvironmentRegistryCredential)
	ResetCertificate()
	ResetEnvironmentVariable()
	ResetFleet()
	ResetImagePullCredentialsType()
	ResetPrivilegedMode()
	ResetRegistryCredential()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodebuildProjectEnvironmentOutputReference
type jsiiProxy_CodebuildProjectEnvironmentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ComputeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ComputeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) EnvironmentVariable() CodebuildProjectEnvironmentEnvironmentVariableList {
	var returns CodebuildProjectEnvironmentEnvironmentVariableList
	_jsii_.Get(
		j,
		"environmentVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) EnvironmentVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) Fleet() CodebuildProjectEnvironmentFleetOutputReference {
	var returns CodebuildProjectEnvironmentFleetOutputReference
	_jsii_.Get(
		j,
		"fleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) FleetInput() *CodebuildProjectEnvironmentFleet {
	var returns *CodebuildProjectEnvironmentFleet
	_jsii_.Get(
		j,
		"fleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ImagePullCredentialsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullCredentialsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ImagePullCredentialsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullCredentialsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) InternalValue() *CodebuildProjectEnvironment {
	var returns *CodebuildProjectEnvironment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) PrivilegedMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) PrivilegedModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) RegistryCredential() CodebuildProjectEnvironmentRegistryCredentialOutputReference {
	var returns CodebuildProjectEnvironmentRegistryCredentialOutputReference
	_jsii_.Get(
		j,
		"registryCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) RegistryCredentialInput() *CodebuildProjectEnvironmentRegistryCredential {
	var returns *CodebuildProjectEnvironmentRegistryCredential
	_jsii_.Get(
		j,
		"registryCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCodebuildProjectEnvironmentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CodebuildProjectEnvironmentOutputReference {
	_init_.Initialize()

	if err := validateNewCodebuildProjectEnvironmentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodebuildProjectEnvironmentOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.codebuildProject.CodebuildProjectEnvironmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCodebuildProjectEnvironmentOutputReference_Override(c CodebuildProjectEnvironmentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.codebuildProject.CodebuildProjectEnvironmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference)SetCertificate(val *string) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference)SetComputeType(val *string) {
	if err := j.validateSetComputeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeType",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference)SetImagePullCredentialsType(val *string) {
	if err := j.validateSetImagePullCredentialsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePullCredentialsType",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference)SetInternalValue(val *CodebuildProjectEnvironment) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference)SetPrivilegedMode(val interface{}) {
	if err := j.validateSetPrivilegedModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilegedMode",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) PutEnvironmentVariable(value interface{}) {
	if err := c.validatePutEnvironmentVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEnvironmentVariable",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) PutFleet(value *CodebuildProjectEnvironmentFleet) {
	if err := c.validatePutFleetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFleet",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) PutRegistryCredential(value *CodebuildProjectEnvironmentRegistryCredential) {
	if err := c.validatePutRegistryCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRegistryCredential",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		c,
		"resetCertificate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ResetEnvironmentVariable() {
	_jsii_.InvokeVoid(
		c,
		"resetEnvironmentVariable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ResetFleet() {
	_jsii_.InvokeVoid(
		c,
		"resetFleet",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ResetImagePullCredentialsType() {
	_jsii_.InvokeVoid(
		c,
		"resetImagePullCredentialsType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ResetPrivilegedMode() {
	_jsii_.InvokeVoid(
		c,
		"resetPrivilegedMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ResetRegistryCredential() {
	_jsii_.InvokeVoid(
		c,
		"resetRegistryCredential",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

