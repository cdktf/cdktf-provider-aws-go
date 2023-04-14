package s3controlstoragelensconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v13/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v13/s3controlstoragelensconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference interface {
	cdktf.ComplexObject
	ActivityMetrics() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetricsOutputReference
	ActivityMetricsInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetrics
	AdvancedCostOptimizationMetrics() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsOutputReference
	AdvancedCostOptimizationMetricsInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetrics
	AdvancedDataProtectionMetrics() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsOutputReference
	AdvancedDataProtectionMetricsInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetrics
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
	DetailedStatusCodeMetrics() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsOutputReference
	DetailedStatusCodeMetricsInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetrics
	// Experimental.
	Fqn() *string
	InternalValue() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevel
	SetInternalValue(val *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevel)
	PrefixLevel() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevelOutputReference
	PrefixLevelInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevel
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
	PutActivityMetrics(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetrics)
	PutAdvancedCostOptimizationMetrics(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetrics)
	PutAdvancedDataProtectionMetrics(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetrics)
	PutDetailedStatusCodeMetrics(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetrics)
	PutPrefixLevel(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevel)
	ResetActivityMetrics()
	ResetAdvancedCostOptimizationMetrics()
	ResetAdvancedDataProtectionMetrics()
	ResetDetailedStatusCodeMetrics()
	ResetPrefixLevel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference
type jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) ActivityMetrics() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetricsOutputReference {
	var returns S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetricsOutputReference
	_jsii_.Get(
		j,
		"activityMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) ActivityMetricsInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetrics {
	var returns *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetrics
	_jsii_.Get(
		j,
		"activityMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) AdvancedCostOptimizationMetrics() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsOutputReference {
	var returns S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsOutputReference
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) AdvancedCostOptimizationMetricsInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetrics {
	var returns *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetrics
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) AdvancedDataProtectionMetrics() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsOutputReference {
	var returns S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsOutputReference
	_jsii_.Get(
		j,
		"advancedDataProtectionMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) AdvancedDataProtectionMetricsInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetrics {
	var returns *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetrics
	_jsii_.Get(
		j,
		"advancedDataProtectionMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) DetailedStatusCodeMetrics() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsOutputReference {
	var returns S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsOutputReference
	_jsii_.Get(
		j,
		"detailedStatusCodeMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) DetailedStatusCodeMetricsInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetrics {
	var returns *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetrics
	_jsii_.Get(
		j,
		"detailedStatusCodeMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) InternalValue() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevel {
	var returns *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevel
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) PrefixLevel() S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevelOutputReference {
	var returns S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevelOutputReference
	_jsii_.Get(
		j,
		"prefixLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) PrefixLevelInput() *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevel {
	var returns *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevel
	_jsii_.Get(
		j,
		"prefixLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference {
	_init_.Initialize()

	if err := validateNewS3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.s3ControlStorageLensConfiguration.S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference_Override(s S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.s3ControlStorageLensConfiguration.S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference)SetInternalValue(val *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevel) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) PutActivityMetrics(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetrics) {
	if err := s.validatePutActivityMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putActivityMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) PutAdvancedCostOptimizationMetrics(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetrics) {
	if err := s.validatePutAdvancedCostOptimizationMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAdvancedCostOptimizationMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) PutAdvancedDataProtectionMetrics(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetrics) {
	if err := s.validatePutAdvancedDataProtectionMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAdvancedDataProtectionMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) PutDetailedStatusCodeMetrics(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetrics) {
	if err := s.validatePutDetailedStatusCodeMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDetailedStatusCodeMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) PutPrefixLevel(value *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevel) {
	if err := s.validatePutPrefixLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPrefixLevel",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) ResetActivityMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetActivityMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) ResetAdvancedCostOptimizationMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetAdvancedCostOptimizationMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) ResetAdvancedDataProtectionMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetAdvancedDataProtectionMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) ResetDetailedStatusCodeMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetDetailedStatusCodeMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) ResetPrefixLevel() {
	_jsii_.InvokeVoid(
		s,
		"resetPrefixLevel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

