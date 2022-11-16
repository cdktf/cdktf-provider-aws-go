package sagemakerendpointconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/sagemakerendpointconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference interface {
	cdktf.ComplexObject
	ClientConfig() SagemakerEndpointConfigurationAsyncInferenceConfigClientConfigOutputReference
	ClientConfigInput() *SagemakerEndpointConfigurationAsyncInferenceConfigClientConfig
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
	InternalValue() *SagemakerEndpointConfigurationAsyncInferenceConfig
	SetInternalValue(val *SagemakerEndpointConfigurationAsyncInferenceConfig)
	OutputConfig() SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference
	OutputConfigInput() *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfig
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
	PutClientConfig(value *SagemakerEndpointConfigurationAsyncInferenceConfigClientConfig)
	PutOutputConfig(value *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfig)
	ResetClientConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference
type jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) ClientConfig() SagemakerEndpointConfigurationAsyncInferenceConfigClientConfigOutputReference {
	var returns SagemakerEndpointConfigurationAsyncInferenceConfigClientConfigOutputReference
	_jsii_.Get(
		j,
		"clientConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) ClientConfigInput() *SagemakerEndpointConfigurationAsyncInferenceConfigClientConfig {
	var returns *SagemakerEndpointConfigurationAsyncInferenceConfigClientConfig
	_jsii_.Get(
		j,
		"clientConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) InternalValue() *SagemakerEndpointConfigurationAsyncInferenceConfig {
	var returns *SagemakerEndpointConfigurationAsyncInferenceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) OutputConfig() SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference {
	var returns SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference
	_jsii_.Get(
		j,
		"outputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) OutputConfigInput() *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfig {
	var returns *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfig
	_jsii_.Get(
		j,
		"outputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerEndpointConfigurationAsyncInferenceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerEndpointConfigurationAsyncInferenceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerEndpointConfiguration.SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerEndpointConfigurationAsyncInferenceConfigOutputReference_Override(s SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerEndpointConfiguration.SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference)SetInternalValue(val *SagemakerEndpointConfigurationAsyncInferenceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) PutClientConfig(value *SagemakerEndpointConfigurationAsyncInferenceConfigClientConfig) {
	if err := s.validatePutClientConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putClientConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) PutOutputConfig(value *SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfig) {
	if err := s.validatePutOutputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOutputConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) ResetClientConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetClientConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

