// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3controlbucketlifecycleconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/s3controlbucketlifecycleconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3ControlBucketLifecycleConfigurationRuleOutputReference interface {
	cdktf.ComplexObject
	AbortIncompleteMultipartUpload() S3ControlBucketLifecycleConfigurationRuleAbortIncompleteMultipartUploadOutputReference
	AbortIncompleteMultipartUploadInput() *S3ControlBucketLifecycleConfigurationRuleAbortIncompleteMultipartUpload
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
	Expiration() S3ControlBucketLifecycleConfigurationRuleExpirationOutputReference
	ExpirationInput() *S3ControlBucketLifecycleConfigurationRuleExpiration
	Filter() S3ControlBucketLifecycleConfigurationRuleFilterOutputReference
	FilterInput() *S3ControlBucketLifecycleConfigurationRuleFilter
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
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
	PutAbortIncompleteMultipartUpload(value *S3ControlBucketLifecycleConfigurationRuleAbortIncompleteMultipartUpload)
	PutExpiration(value *S3ControlBucketLifecycleConfigurationRuleExpiration)
	PutFilter(value *S3ControlBucketLifecycleConfigurationRuleFilter)
	ResetAbortIncompleteMultipartUpload()
	ResetExpiration()
	ResetFilter()
	ResetStatus()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3ControlBucketLifecycleConfigurationRuleOutputReference
type jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) AbortIncompleteMultipartUpload() S3ControlBucketLifecycleConfigurationRuleAbortIncompleteMultipartUploadOutputReference {
	var returns S3ControlBucketLifecycleConfigurationRuleAbortIncompleteMultipartUploadOutputReference
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUpload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) AbortIncompleteMultipartUploadInput() *S3ControlBucketLifecycleConfigurationRuleAbortIncompleteMultipartUpload {
	var returns *S3ControlBucketLifecycleConfigurationRuleAbortIncompleteMultipartUpload
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUploadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) Expiration() S3ControlBucketLifecycleConfigurationRuleExpirationOutputReference {
	var returns S3ControlBucketLifecycleConfigurationRuleExpirationOutputReference
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) ExpirationInput() *S3ControlBucketLifecycleConfigurationRuleExpiration {
	var returns *S3ControlBucketLifecycleConfigurationRuleExpiration
	_jsii_.Get(
		j,
		"expirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) Filter() S3ControlBucketLifecycleConfigurationRuleFilterOutputReference {
	var returns S3ControlBucketLifecycleConfigurationRuleFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) FilterInput() *S3ControlBucketLifecycleConfigurationRuleFilter {
	var returns *S3ControlBucketLifecycleConfigurationRuleFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3ControlBucketLifecycleConfigurationRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) S3ControlBucketLifecycleConfigurationRuleOutputReference {
	_init_.Initialize()

	if err := validateNewS3ControlBucketLifecycleConfigurationRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.s3ControlBucketLifecycleConfiguration.S3ControlBucketLifecycleConfigurationRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewS3ControlBucketLifecycleConfigurationRuleOutputReference_Override(s S3ControlBucketLifecycleConfigurationRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.s3ControlBucketLifecycleConfiguration.S3ControlBucketLifecycleConfigurationRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) PutAbortIncompleteMultipartUpload(value *S3ControlBucketLifecycleConfigurationRuleAbortIncompleteMultipartUpload) {
	if err := s.validatePutAbortIncompleteMultipartUploadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAbortIncompleteMultipartUpload",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) PutExpiration(value *S3ControlBucketLifecycleConfigurationRuleExpiration) {
	if err := s.validatePutExpirationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExpiration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) PutFilter(value *S3ControlBucketLifecycleConfigurationRuleFilter) {
	if err := s.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFilter",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) ResetAbortIncompleteMultipartUpload() {
	_jsii_.InvokeVoid(
		s,
		"resetAbortIncompleteMultipartUpload",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) ResetExpiration() {
	_jsii_.InvokeVoid(
		s,
		"resetExpiration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		s,
		"resetFilter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3ControlBucketLifecycleConfigurationRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

