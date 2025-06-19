// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmmaintenancewindowtask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ssmmaintenancewindowtask/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference interface {
	cdktf.ComplexObject
	CloudwatchConfig() SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference
	CloudwatchConfigInput() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
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
	DocumentHash() *string
	SetDocumentHash(val *string)
	DocumentHashInput() *string
	DocumentHashType() *string
	SetDocumentHashType(val *string)
	DocumentHashTypeInput() *string
	DocumentVersion() *string
	SetDocumentVersion(val *string)
	DocumentVersionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters
	SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters)
	NotificationConfig() SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference
	NotificationConfigInput() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig
	OutputS3Bucket() *string
	SetOutputS3Bucket(val *string)
	OutputS3BucketInput() *string
	OutputS3KeyPrefix() *string
	SetOutputS3KeyPrefix(val *string)
	OutputS3KeyPrefixInput() *string
	Parameter() SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterList
	ParameterInput() interface{}
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	ServiceRoleArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutSeconds() *float64
	SetTimeoutSeconds(val *float64)
	TimeoutSecondsInput() *float64
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
	PutCloudwatchConfig(value *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig)
	PutNotificationConfig(value *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig)
	PutParameter(value interface{})
	ResetCloudwatchConfig()
	ResetComment()
	ResetDocumentHash()
	ResetDocumentHashType()
	ResetDocumentVersion()
	ResetNotificationConfig()
	ResetOutputS3Bucket()
	ResetOutputS3KeyPrefix()
	ResetParameter()
	ResetServiceRoleArn()
	ResetTimeoutSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference
type jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) CloudwatchConfig() SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference {
	var returns SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfigOutputReference
	_jsii_.Get(
		j,
		"cloudwatchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) CloudwatchConfigInput() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig
	_jsii_.Get(
		j,
		"cloudwatchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) DocumentHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) DocumentHashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) DocumentHashType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentHashType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) DocumentHashTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentHashTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) DocumentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) DocumentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) NotificationConfig() SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference {
	var returns SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigOutputReference
	_jsii_.Get(
		j,
		"notificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) NotificationConfigInput() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig
	_jsii_.Get(
		j,
		"notificationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) OutputS3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) OutputS3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputS3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) OutputS3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputS3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) OutputS3KeyPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputS3KeyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) Parameter() SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterList {
	var returns SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterList
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}


func NewSsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference {
	_init_.Initialize()

	if err := validateNewSsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference_Override(s SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference)SetDocumentHash(val *string) {
	if err := j.validateSetDocumentHashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentHash",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference)SetDocumentHashType(val *string) {
	if err := j.validateSetDocumentHashTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentHashType",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference)SetDocumentVersion(val *string) {
	if err := j.validateSetDocumentVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentVersion",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference)SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference)SetOutputS3Bucket(val *string) {
	if err := j.validateSetOutputS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputS3Bucket",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference)SetOutputS3KeyPrefix(val *string) {
	if err := j.validateSetOutputS3KeyPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputS3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference)SetServiceRoleArn(val *string) {
	if err := j.validateSetServiceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) PutCloudwatchConfig(value *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig) {
	if err := s.validatePutCloudwatchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCloudwatchConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) PutNotificationConfig(value *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig) {
	if err := s.validatePutNotificationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNotificationConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) PutParameter(value interface{}) {
	if err := s.validatePutParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putParameter",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetCloudwatchConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetCloudwatchConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		s,
		"resetComment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetDocumentHash() {
	_jsii_.InvokeVoid(
		s,
		"resetDocumentHash",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetDocumentHashType() {
	_jsii_.InvokeVoid(
		s,
		"resetDocumentHashType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetDocumentVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetDocumentVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetNotificationConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetNotificationConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetOutputS3Bucket() {
	_jsii_.InvokeVoid(
		s,
		"resetOutputS3Bucket",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetOutputS3KeyPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetOutputS3KeyPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetParameter() {
	_jsii_.InvokeVoid(
		s,
		"resetParameter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetServiceRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetServiceRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

