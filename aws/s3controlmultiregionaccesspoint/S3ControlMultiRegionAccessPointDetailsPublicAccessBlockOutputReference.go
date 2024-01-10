// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3controlmultiregionaccesspoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/s3controlmultiregionaccesspoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference interface {
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
	InternalValue() *S3ControlMultiRegionAccessPointDetailsPublicAccessBlock
	SetInternalValue(val *S3ControlMultiRegionAccessPointDetailsPublicAccessBlock)
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

// The jsii proxy struct for S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference
type jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) BlockPublicAcls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPublicAcls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) BlockPublicAclsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPublicAclsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) BlockPublicPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPublicPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) BlockPublicPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockPublicPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) IgnorePublicAcls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignorePublicAcls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) IgnorePublicAclsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignorePublicAclsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) InternalValue() *S3ControlMultiRegionAccessPointDetailsPublicAccessBlock {
	var returns *S3ControlMultiRegionAccessPointDetailsPublicAccessBlock
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) RestrictPublicBuckets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictPublicBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) RestrictPublicBucketsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictPublicBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference {
	_init_.Initialize()

	if err := validateNewS3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.s3ControlMultiRegionAccessPoint.S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference_Override(s S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.s3ControlMultiRegionAccessPoint.S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference)SetBlockPublicAcls(val interface{}) {
	if err := j.validateSetBlockPublicAclsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockPublicAcls",
		val,
	)
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference)SetBlockPublicPolicy(val interface{}) {
	if err := j.validateSetBlockPublicPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockPublicPolicy",
		val,
	)
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference)SetIgnorePublicAcls(val interface{}) {
	if err := j.validateSetIgnorePublicAclsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignorePublicAcls",
		val,
	)
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference)SetInternalValue(val *S3ControlMultiRegionAccessPointDetailsPublicAccessBlock) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference)SetRestrictPublicBuckets(val interface{}) {
	if err := j.validateSetRestrictPublicBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictPublicBuckets",
		val,
	)
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) ResetBlockPublicAcls() {
	_jsii_.InvokeVoid(
		s,
		"resetBlockPublicAcls",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) ResetBlockPublicPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetBlockPublicPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) ResetIgnorePublicAcls() {
	_jsii_.InvokeVoid(
		s,
		"resetIgnorePublicAcls",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) ResetRestrictPublicBuckets() {
	_jsii_.InvokeVoid(
		s,
		"resetRestrictPublicBuckets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3ControlMultiRegionAccessPointDetailsPublicAccessBlockOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

