// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3controlstoragelensconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/s3controlstoragelensconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference interface {
	cdktf.ComplexObject
	ActivityMetrics() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetricsOutputReference
	ActivityMetricsInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetrics
	AdvancedCostOptimizationMetrics() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsOutputReference
	AdvancedCostOptimizationMetricsInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetrics
	AdvancedDataProtectionMetrics() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsOutputReference
	AdvancedDataProtectionMetricsInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedDataProtectionMetrics
	BucketLevel() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference
	BucketLevelInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevel
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
	DetailedStatusCodeMetrics() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelDetailedStatusCodeMetricsOutputReference
	DetailedStatusCodeMetricsInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelDetailedStatusCodeMetrics
	// Experimental.
	Fqn() *string
	InternalValue() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevel
	SetInternalValue(val *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevel)
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
	PutActivityMetrics(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetrics)
	PutAdvancedCostOptimizationMetrics(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetrics)
	PutAdvancedDataProtectionMetrics(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedDataProtectionMetrics)
	PutBucketLevel(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevel)
	PutDetailedStatusCodeMetrics(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelDetailedStatusCodeMetrics)
	ResetActivityMetrics()
	ResetAdvancedCostOptimizationMetrics()
	ResetAdvancedDataProtectionMetrics()
	ResetDetailedStatusCodeMetrics()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference
type jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) ActivityMetrics() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetricsOutputReference {
	var returns S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetricsOutputReference
	_jsii_.Get(
		j,
		"activityMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) ActivityMetricsInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetrics {
	var returns *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetrics
	_jsii_.Get(
		j,
		"activityMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) AdvancedCostOptimizationMetrics() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsOutputReference {
	var returns S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsOutputReference
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) AdvancedCostOptimizationMetricsInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetrics {
	var returns *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetrics
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) AdvancedDataProtectionMetrics() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsOutputReference {
	var returns S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsOutputReference
	_jsii_.Get(
		j,
		"advancedDataProtectionMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) AdvancedDataProtectionMetricsInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedDataProtectionMetrics {
	var returns *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedDataProtectionMetrics
	_jsii_.Get(
		j,
		"advancedDataProtectionMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) BucketLevel() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference {
	var returns S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference
	_jsii_.Get(
		j,
		"bucketLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) BucketLevelInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevel {
	var returns *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevel
	_jsii_.Get(
		j,
		"bucketLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) DetailedStatusCodeMetrics() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelDetailedStatusCodeMetricsOutputReference {
	var returns S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelDetailedStatusCodeMetricsOutputReference
	_jsii_.Get(
		j,
		"detailedStatusCodeMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) DetailedStatusCodeMetricsInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelDetailedStatusCodeMetrics {
	var returns *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelDetailedStatusCodeMetrics
	_jsii_.Get(
		j,
		"detailedStatusCodeMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) InternalValue() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevel {
	var returns *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevel
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference {
	_init_.Initialize()

	if err := validateNewS3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.s3ControlStorageLensConfiguration.S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference_Override(s S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.s3ControlStorageLensConfiguration.S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference)SetInternalValue(val *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevel) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) PutActivityMetrics(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetrics) {
	if err := s.validatePutActivityMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putActivityMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) PutAdvancedCostOptimizationMetrics(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetrics) {
	if err := s.validatePutAdvancedCostOptimizationMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAdvancedCostOptimizationMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) PutAdvancedDataProtectionMetrics(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedDataProtectionMetrics) {
	if err := s.validatePutAdvancedDataProtectionMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAdvancedDataProtectionMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) PutBucketLevel(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevel) {
	if err := s.validatePutBucketLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBucketLevel",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) PutDetailedStatusCodeMetrics(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelDetailedStatusCodeMetrics) {
	if err := s.validatePutDetailedStatusCodeMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDetailedStatusCodeMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) ResetActivityMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetActivityMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) ResetAdvancedCostOptimizationMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetAdvancedCostOptimizationMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) ResetAdvancedDataProtectionMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetAdvancedDataProtectionMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) ResetDetailedStatusCodeMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetDetailedStatusCodeMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

