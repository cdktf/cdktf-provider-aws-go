package s3controlobjectlambdaaccesspoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/s3controlobjectlambdaaccesspoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3ControlObjectLambdaAccessPointConfigurationOutputReference interface {
	cdktf.ComplexObject
	AllowedFeatures() *[]*string
	SetAllowedFeatures(val *[]*string)
	AllowedFeaturesInput() *[]*string
	CloudWatchMetricsEnabled() interface{}
	SetCloudWatchMetricsEnabled(val interface{})
	CloudWatchMetricsEnabledInput() interface{}
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
	InternalValue() *S3ControlObjectLambdaAccessPointConfiguration
	SetInternalValue(val *S3ControlObjectLambdaAccessPointConfiguration)
	SupportingAccessPoint() *string
	SetSupportingAccessPoint(val *string)
	SupportingAccessPointInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransformationConfiguration() S3ControlObjectLambdaAccessPointConfigurationTransformationConfigurationList
	TransformationConfigurationInput() interface{}
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
	PutTransformationConfiguration(value interface{})
	ResetAllowedFeatures()
	ResetCloudWatchMetricsEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3ControlObjectLambdaAccessPointConfigurationOutputReference
type jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) AllowedFeatures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) AllowedFeaturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) CloudWatchMetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudWatchMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) CloudWatchMetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudWatchMetricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) InternalValue() *S3ControlObjectLambdaAccessPointConfiguration {
	var returns *S3ControlObjectLambdaAccessPointConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) SupportingAccessPoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportingAccessPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) SupportingAccessPointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportingAccessPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) TransformationConfiguration() S3ControlObjectLambdaAccessPointConfigurationTransformationConfigurationList {
	var returns S3ControlObjectLambdaAccessPointConfigurationTransformationConfigurationList
	_jsii_.Get(
		j,
		"transformationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) TransformationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformationConfigurationInput",
		&returns,
	)
	return returns
}


func NewS3ControlObjectLambdaAccessPointConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) S3ControlObjectLambdaAccessPointConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewS3ControlObjectLambdaAccessPointConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.s3ControlObjectLambdaAccessPoint.S3ControlObjectLambdaAccessPointConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3ControlObjectLambdaAccessPointConfigurationOutputReference_Override(s S3ControlObjectLambdaAccessPointConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.s3ControlObjectLambdaAccessPoint.S3ControlObjectLambdaAccessPointConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference)SetAllowedFeatures(val *[]*string) {
	if err := j.validateSetAllowedFeaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedFeatures",
		val,
	)
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference)SetCloudWatchMetricsEnabled(val interface{}) {
	if err := j.validateSetCloudWatchMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudWatchMetricsEnabled",
		val,
	)
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference)SetInternalValue(val *S3ControlObjectLambdaAccessPointConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference)SetSupportingAccessPoint(val *string) {
	if err := j.validateSetSupportingAccessPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportingAccessPoint",
		val,
	)
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) PutTransformationConfiguration(value interface{}) {
	if err := s.validatePutTransformationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTransformationConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) ResetAllowedFeatures() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedFeatures",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) ResetCloudWatchMetricsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetCloudWatchMetricsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3ControlObjectLambdaAccessPointConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

