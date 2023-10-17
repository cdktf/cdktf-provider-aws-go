// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/schedulerschedule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SchedulerScheduleTargetOutputReference interface {
	cdktf.ComplexObject
	Arn() *string
	SetArn(val *string)
	ArnInput() *string
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
	DeadLetterConfig() SchedulerScheduleTargetDeadLetterConfigOutputReference
	DeadLetterConfigInput() *SchedulerScheduleTargetDeadLetterConfig
	EcsParameters() SchedulerScheduleTargetEcsParametersOutputReference
	EcsParametersInput() *SchedulerScheduleTargetEcsParameters
	EventbridgeParameters() SchedulerScheduleTargetEventbridgeParametersOutputReference
	EventbridgeParametersInput() *SchedulerScheduleTargetEventbridgeParameters
	// Experimental.
	Fqn() *string
	Input() *string
	SetInput(val *string)
	InputInput() *string
	InternalValue() *SchedulerScheduleTarget
	SetInternalValue(val *SchedulerScheduleTarget)
	KinesisParameters() SchedulerScheduleTargetKinesisParametersOutputReference
	KinesisParametersInput() *SchedulerScheduleTargetKinesisParameters
	RetryPolicy() SchedulerScheduleTargetRetryPolicyOutputReference
	RetryPolicyInput() *SchedulerScheduleTargetRetryPolicy
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	SagemakerPipelineParameters() SchedulerScheduleTargetSagemakerPipelineParametersOutputReference
	SagemakerPipelineParametersInput() *SchedulerScheduleTargetSagemakerPipelineParameters
	SqsParameters() SchedulerScheduleTargetSqsParametersOutputReference
	SqsParametersInput() *SchedulerScheduleTargetSqsParameters
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
	PutDeadLetterConfig(value *SchedulerScheduleTargetDeadLetterConfig)
	PutEcsParameters(value *SchedulerScheduleTargetEcsParameters)
	PutEventbridgeParameters(value *SchedulerScheduleTargetEventbridgeParameters)
	PutKinesisParameters(value *SchedulerScheduleTargetKinesisParameters)
	PutRetryPolicy(value *SchedulerScheduleTargetRetryPolicy)
	PutSagemakerPipelineParameters(value *SchedulerScheduleTargetSagemakerPipelineParameters)
	PutSqsParameters(value *SchedulerScheduleTargetSqsParameters)
	ResetDeadLetterConfig()
	ResetEcsParameters()
	ResetEventbridgeParameters()
	ResetInput()
	ResetKinesisParameters()
	ResetRetryPolicy()
	ResetSagemakerPipelineParameters()
	ResetSqsParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SchedulerScheduleTargetOutputReference
type jsiiProxy_SchedulerScheduleTargetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) DeadLetterConfig() SchedulerScheduleTargetDeadLetterConfigOutputReference {
	var returns SchedulerScheduleTargetDeadLetterConfigOutputReference
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) DeadLetterConfigInput() *SchedulerScheduleTargetDeadLetterConfig {
	var returns *SchedulerScheduleTargetDeadLetterConfig
	_jsii_.Get(
		j,
		"deadLetterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) EcsParameters() SchedulerScheduleTargetEcsParametersOutputReference {
	var returns SchedulerScheduleTargetEcsParametersOutputReference
	_jsii_.Get(
		j,
		"ecsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) EcsParametersInput() *SchedulerScheduleTargetEcsParameters {
	var returns *SchedulerScheduleTargetEcsParameters
	_jsii_.Get(
		j,
		"ecsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) EventbridgeParameters() SchedulerScheduleTargetEventbridgeParametersOutputReference {
	var returns SchedulerScheduleTargetEventbridgeParametersOutputReference
	_jsii_.Get(
		j,
		"eventbridgeParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) EventbridgeParametersInput() *SchedulerScheduleTargetEventbridgeParameters {
	var returns *SchedulerScheduleTargetEventbridgeParameters
	_jsii_.Get(
		j,
		"eventbridgeParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) InputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) InternalValue() *SchedulerScheduleTarget {
	var returns *SchedulerScheduleTarget
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) KinesisParameters() SchedulerScheduleTargetKinesisParametersOutputReference {
	var returns SchedulerScheduleTargetKinesisParametersOutputReference
	_jsii_.Get(
		j,
		"kinesisParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) KinesisParametersInput() *SchedulerScheduleTargetKinesisParameters {
	var returns *SchedulerScheduleTargetKinesisParameters
	_jsii_.Get(
		j,
		"kinesisParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) RetryPolicy() SchedulerScheduleTargetRetryPolicyOutputReference {
	var returns SchedulerScheduleTargetRetryPolicyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) RetryPolicyInput() *SchedulerScheduleTargetRetryPolicy {
	var returns *SchedulerScheduleTargetRetryPolicy
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) SagemakerPipelineParameters() SchedulerScheduleTargetSagemakerPipelineParametersOutputReference {
	var returns SchedulerScheduleTargetSagemakerPipelineParametersOutputReference
	_jsii_.Get(
		j,
		"sagemakerPipelineParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) SagemakerPipelineParametersInput() *SchedulerScheduleTargetSagemakerPipelineParameters {
	var returns *SchedulerScheduleTargetSagemakerPipelineParameters
	_jsii_.Get(
		j,
		"sagemakerPipelineParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) SqsParameters() SchedulerScheduleTargetSqsParametersOutputReference {
	var returns SchedulerScheduleTargetSqsParametersOutputReference
	_jsii_.Get(
		j,
		"sqsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) SqsParametersInput() *SchedulerScheduleTargetSqsParameters {
	var returns *SchedulerScheduleTargetSqsParameters
	_jsii_.Get(
		j,
		"sqsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSchedulerScheduleTargetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SchedulerScheduleTargetOutputReference {
	_init_.Initialize()

	if err := validateNewSchedulerScheduleTargetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SchedulerScheduleTargetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.schedulerSchedule.SchedulerScheduleTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSchedulerScheduleTargetOutputReference_Override(s SchedulerScheduleTargetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.schedulerSchedule.SchedulerScheduleTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference)SetArn(val *string) {
	if err := j.validateSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference)SetInput(val *string) {
	if err := j.validateSetInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference)SetInternalValue(val *SchedulerScheduleTarget) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) PutDeadLetterConfig(value *SchedulerScheduleTargetDeadLetterConfig) {
	if err := s.validatePutDeadLetterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDeadLetterConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) PutEcsParameters(value *SchedulerScheduleTargetEcsParameters) {
	if err := s.validatePutEcsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEcsParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) PutEventbridgeParameters(value *SchedulerScheduleTargetEventbridgeParameters) {
	if err := s.validatePutEventbridgeParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEventbridgeParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) PutKinesisParameters(value *SchedulerScheduleTargetKinesisParameters) {
	if err := s.validatePutKinesisParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putKinesisParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) PutRetryPolicy(value *SchedulerScheduleTargetRetryPolicy) {
	if err := s.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) PutSagemakerPipelineParameters(value *SchedulerScheduleTargetSagemakerPipelineParameters) {
	if err := s.validatePutSagemakerPipelineParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSagemakerPipelineParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) PutSqsParameters(value *SchedulerScheduleTargetSqsParameters) {
	if err := s.validatePutSqsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSqsParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) ResetDeadLetterConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetDeadLetterConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) ResetEcsParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetEcsParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) ResetEventbridgeParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetEventbridgeParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) ResetInput() {
	_jsii_.InvokeVoid(
		s,
		"resetInput",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) ResetKinesisParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetKinesisParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) ResetSagemakerPipelineParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetSagemakerPipelineParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) ResetSqsParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetSqsParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SchedulerScheduleTargetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

