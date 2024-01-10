// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerappimageconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/sagemakerappimageconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference interface {
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
	DefaultGid() *float64
	SetDefaultGid(val *float64)
	DefaultGidInput() *float64
	DefaultUid() *float64
	SetDefaultUid(val *float64)
	DefaultUidInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfig
	SetInternalValue(val *SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfig)
	MountPath() *string
	SetMountPath(val *string)
	MountPathInput() *string
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
	ResetDefaultGid()
	ResetDefaultUid()
	ResetMountPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference
type jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) DefaultGid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultGid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) DefaultGidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultGidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) DefaultUid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) DefaultUidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultUidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) InternalValue() *SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfig {
	var returns *SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) MountPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) MountPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerAppImageConfig.SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference_Override(s SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerAppImageConfig.SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference)SetDefaultGid(val *float64) {
	if err := j.validateSetDefaultGidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultGid",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference)SetDefaultUid(val *float64) {
	if err := j.validateSetDefaultUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultUid",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference)SetInternalValue(val *SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference)SetMountPath(val *string) {
	if err := j.validateSetMountPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mountPath",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) ResetDefaultGid() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultGid",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) ResetDefaultUid() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultUid",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) ResetMountPath() {
	_jsii_.InvokeVoid(
		s,
		"resetMountPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

