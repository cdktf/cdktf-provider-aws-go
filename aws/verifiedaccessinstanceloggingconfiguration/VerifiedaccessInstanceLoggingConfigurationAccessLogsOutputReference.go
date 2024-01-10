// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedaccessinstanceloggingconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/verifiedaccessinstanceloggingconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference interface {
	cdktf.ComplexObject
	CloudwatchLogs() VerifiedaccessInstanceLoggingConfigurationAccessLogsCloudwatchLogsOutputReference
	CloudwatchLogsInput() *VerifiedaccessInstanceLoggingConfigurationAccessLogsCloudwatchLogs
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
	IncludeTrustContext() interface{}
	SetIncludeTrustContext(val interface{})
	IncludeTrustContextInput() interface{}
	InternalValue() *VerifiedaccessInstanceLoggingConfigurationAccessLogs
	SetInternalValue(val *VerifiedaccessInstanceLoggingConfigurationAccessLogs)
	KinesisDataFirehose() VerifiedaccessInstanceLoggingConfigurationAccessLogsKinesisDataFirehoseOutputReference
	KinesisDataFirehoseInput() *VerifiedaccessInstanceLoggingConfigurationAccessLogsKinesisDataFirehose
	LogVersion() *string
	SetLogVersion(val *string)
	LogVersionInput() *string
	S3() VerifiedaccessInstanceLoggingConfigurationAccessLogsS3OutputReference
	S3Input() *VerifiedaccessInstanceLoggingConfigurationAccessLogsS3
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
	PutCloudwatchLogs(value *VerifiedaccessInstanceLoggingConfigurationAccessLogsCloudwatchLogs)
	PutKinesisDataFirehose(value *VerifiedaccessInstanceLoggingConfigurationAccessLogsKinesisDataFirehose)
	PutS3(value *VerifiedaccessInstanceLoggingConfigurationAccessLogsS3)
	ResetCloudwatchLogs()
	ResetIncludeTrustContext()
	ResetKinesisDataFirehose()
	ResetLogVersion()
	ResetS3()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference
type jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) CloudwatchLogs() VerifiedaccessInstanceLoggingConfigurationAccessLogsCloudwatchLogsOutputReference {
	var returns VerifiedaccessInstanceLoggingConfigurationAccessLogsCloudwatchLogsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) CloudwatchLogsInput() *VerifiedaccessInstanceLoggingConfigurationAccessLogsCloudwatchLogs {
	var returns *VerifiedaccessInstanceLoggingConfigurationAccessLogsCloudwatchLogs
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) IncludeTrustContext() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTrustContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) IncludeTrustContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTrustContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) InternalValue() *VerifiedaccessInstanceLoggingConfigurationAccessLogs {
	var returns *VerifiedaccessInstanceLoggingConfigurationAccessLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) KinesisDataFirehose() VerifiedaccessInstanceLoggingConfigurationAccessLogsKinesisDataFirehoseOutputReference {
	var returns VerifiedaccessInstanceLoggingConfigurationAccessLogsKinesisDataFirehoseOutputReference
	_jsii_.Get(
		j,
		"kinesisDataFirehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) KinesisDataFirehoseInput() *VerifiedaccessInstanceLoggingConfigurationAccessLogsKinesisDataFirehose {
	var returns *VerifiedaccessInstanceLoggingConfigurationAccessLogsKinesisDataFirehose
	_jsii_.Get(
		j,
		"kinesisDataFirehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) LogVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) LogVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) S3() VerifiedaccessInstanceLoggingConfigurationAccessLogsS3OutputReference {
	var returns VerifiedaccessInstanceLoggingConfigurationAccessLogsS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) S3Input() *VerifiedaccessInstanceLoggingConfigurationAccessLogsS3 {
	var returns *VerifiedaccessInstanceLoggingConfigurationAccessLogsS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference {
	_init_.Initialize()

	if err := validateNewVerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.verifiedaccessInstanceLoggingConfiguration.VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference_Override(v VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.verifiedaccessInstanceLoggingConfiguration.VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference)SetIncludeTrustContext(val interface{}) {
	if err := j.validateSetIncludeTrustContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTrustContext",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference)SetInternalValue(val *VerifiedaccessInstanceLoggingConfigurationAccessLogs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference)SetLogVersion(val *string) {
	if err := j.validateSetLogVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logVersion",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) PutCloudwatchLogs(value *VerifiedaccessInstanceLoggingConfigurationAccessLogsCloudwatchLogs) {
	if err := v.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) PutKinesisDataFirehose(value *VerifiedaccessInstanceLoggingConfigurationAccessLogsKinesisDataFirehose) {
	if err := v.validatePutKinesisDataFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putKinesisDataFirehose",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) PutS3(value *VerifiedaccessInstanceLoggingConfigurationAccessLogsS3) {
	if err := v.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putS3",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		v,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) ResetIncludeTrustContext() {
	_jsii_.InvokeVoid(
		v,
		"resetIncludeTrustContext",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) ResetKinesisDataFirehose() {
	_jsii_.InvokeVoid(
		v,
		"resetKinesisDataFirehose",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) ResetLogVersion() {
	_jsii_.InvokeVoid(
		v,
		"resetLogVersion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		v,
		"resetS3",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessInstanceLoggingConfigurationAccessLogsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

