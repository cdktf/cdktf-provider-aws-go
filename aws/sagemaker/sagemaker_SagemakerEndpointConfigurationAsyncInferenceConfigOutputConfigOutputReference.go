package sagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/sagemaker/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfig
	SetInternalValue(val *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfig)
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	NotificationConfig() SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference
	NotificationConfigInput() *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfig
	S3OutputPath() *string
	SetS3OutputPath(val *string)
	S3OutputPathInput() *string
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
	PutNotificationConfig(value *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfig)
	ResetKmsKeyId()
	ResetNotificationConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference
type jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) InternalValue() *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfig {
	var returns *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) NotificationConfig() SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference {
	var returns SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference
	_jsii_.Get(
		j,
		"notificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) NotificationConfigInput() *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfig {
	var returns *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfig
	_jsii_.Get(
		j,
		"notificationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) S3OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3OutputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) S3OutputPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3OutputPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemaker.SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference_Override(s SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemaker.SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference)SetInternalValue(val *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference)SetS3OutputPath(val *string) {
	if err := j.validateSetS3OutputPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3OutputPath",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) PutNotificationConfig(value *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfig) {
	if err := s.validatePutNotificationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNotificationConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) ResetNotificationConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetNotificationConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

