// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3accesspoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/s3accesspoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3AccessPointPublicAccessBlockConfigurationOutputReference interface {
	cdktf.ComplexObject
	BlockPublicAcls() interface{}
	SetBlockPublicAcls(val interface{})
	BlockPublicAclsInput() interface{}
	BlockPublicPolicy() interface{}
	SetBlockPublicPolicy(val interface{})
	BlockPublicPolicyInput() interface{}
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
	IgnorePublicAcls() interface{}
	SetIgnorePublicAcls(val interface{})
	IgnorePublicAclsInput() interface{}
	InternalValue() *S3AccessPointPublicAccessBlockConfiguration
	SetInternalValue(val *S3AccessPointPublicAccessBlockConfiguration)
	RestrictPublicBuckets() interface{}
	SetRestrictPublicBuckets(val interface{})
	RestrictPublicBucketsInput() interface{}
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
	ResetBlockPublicAcls()
	ResetBlockPublicPolicy()
	ResetIgnorePublicAcls()
	ResetRestrictPublicBuckets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3AccessPointPublicAccessBlockConfigurationOutputReference
type jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) BlockPublicAcls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPublicAcls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) BlockPublicAclsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPublicAclsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) BlockPublicPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPublicPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) BlockPublicPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPublicPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) IgnorePublicAcls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignorePublicAcls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) IgnorePublicAclsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignorePublicAclsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) InternalValue() *S3AccessPointPublicAccessBlockConfiguration {
	var returns *S3AccessPointPublicAccessBlockConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) RestrictPublicBuckets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictPublicBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) RestrictPublicBucketsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictPublicBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3AccessPointPublicAccessBlockConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) S3AccessPointPublicAccessBlockConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewS3AccessPointPublicAccessBlockConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.s3AccessPoint.S3AccessPointPublicAccessBlockConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3AccessPointPublicAccessBlockConfigurationOutputReference_Override(s S3AccessPointPublicAccessBlockConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.s3AccessPoint.S3AccessPointPublicAccessBlockConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference)SetBlockPublicAcls(val interface{}) {
	if err := j.validateSetBlockPublicAclsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockPublicAcls",
		val,
	)
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference)SetBlockPublicPolicy(val interface{}) {
	if err := j.validateSetBlockPublicPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockPublicPolicy",
		val,
	)
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference)SetIgnorePublicAcls(val interface{}) {
	if err := j.validateSetIgnorePublicAclsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignorePublicAcls",
		val,
	)
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference)SetInternalValue(val *S3AccessPointPublicAccessBlockConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference)SetRestrictPublicBuckets(val interface{}) {
	if err := j.validateSetRestrictPublicBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictPublicBuckets",
		val,
	)
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) ResetBlockPublicAcls() {
	_jsii_.InvokeVoid(
		s,
		"resetBlockPublicAcls",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) ResetBlockPublicPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetBlockPublicPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) ResetIgnorePublicAcls() {
	_jsii_.InvokeVoid(
		s,
		"resetIgnorePublicAcls",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) ResetRestrictPublicBuckets() {
	_jsii_.InvokeVoid(
		s,
		"resetRestrictPublicBuckets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3AccessPointPublicAccessBlockConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

