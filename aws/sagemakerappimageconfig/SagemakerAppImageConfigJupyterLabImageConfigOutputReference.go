// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerappimageconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/sagemakerappimageconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerAppImageConfigJupyterLabImageConfigOutputReference interface {
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
	ContainerConfig() SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference
	ContainerConfigInput() *SagemakerAppImageConfigJupyterLabImageConfigContainerConfig
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FileSystemConfig() SagemakerAppImageConfigJupyterLabImageConfigFileSystemConfigOutputReference
	FileSystemConfigInput() *SagemakerAppImageConfigJupyterLabImageConfigFileSystemConfig
	// Experimental.
	Fqn() *string
	InternalValue() *SagemakerAppImageConfigJupyterLabImageConfig
	SetInternalValue(val *SagemakerAppImageConfigJupyterLabImageConfig)
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
	PutContainerConfig(value *SagemakerAppImageConfigJupyterLabImageConfigContainerConfig)
	PutFileSystemConfig(value *SagemakerAppImageConfigJupyterLabImageConfigFileSystemConfig)
	ResetContainerConfig()
	ResetFileSystemConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerAppImageConfigJupyterLabImageConfigOutputReference
type jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) ContainerConfig() SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference {
	var returns SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference
	_jsii_.Get(
		j,
		"containerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) ContainerConfigInput() *SagemakerAppImageConfigJupyterLabImageConfigContainerConfig {
	var returns *SagemakerAppImageConfigJupyterLabImageConfigContainerConfig
	_jsii_.Get(
		j,
		"containerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) FileSystemConfig() SagemakerAppImageConfigJupyterLabImageConfigFileSystemConfigOutputReference {
	var returns SagemakerAppImageConfigJupyterLabImageConfigFileSystemConfigOutputReference
	_jsii_.Get(
		j,
		"fileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) FileSystemConfigInput() *SagemakerAppImageConfigJupyterLabImageConfigFileSystemConfig {
	var returns *SagemakerAppImageConfigJupyterLabImageConfigFileSystemConfig
	_jsii_.Get(
		j,
		"fileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) InternalValue() *SagemakerAppImageConfigJupyterLabImageConfig {
	var returns *SagemakerAppImageConfigJupyterLabImageConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerAppImageConfigJupyterLabImageConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerAppImageConfigJupyterLabImageConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerAppImageConfigJupyterLabImageConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerAppImageConfig.SagemakerAppImageConfigJupyterLabImageConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerAppImageConfigJupyterLabImageConfigOutputReference_Override(s SagemakerAppImageConfigJupyterLabImageConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerAppImageConfig.SagemakerAppImageConfigJupyterLabImageConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference)SetInternalValue(val *SagemakerAppImageConfigJupyterLabImageConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) PutContainerConfig(value *SagemakerAppImageConfigJupyterLabImageConfigContainerConfig) {
	if err := s.validatePutContainerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putContainerConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) PutFileSystemConfig(value *SagemakerAppImageConfigJupyterLabImageConfigFileSystemConfig) {
	if err := s.validatePutFileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFileSystemConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) ResetContainerConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetContainerConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) ResetFileSystemConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetFileSystemConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

