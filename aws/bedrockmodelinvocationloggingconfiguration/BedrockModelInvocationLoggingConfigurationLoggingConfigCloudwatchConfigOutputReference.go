// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockmodelinvocationloggingconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/bedrockmodelinvocationloggingconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LargeDataDeliveryS3Config() BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3ConfigList
	LargeDataDeliveryS3ConfigInput() interface{}
	LogGroupName() *string
	SetLogGroupName(val *string)
	LogGroupNameInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
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
	PutLargeDataDeliveryS3Config(value interface{})
	ResetLargeDataDeliveryS3Config()
	ResetLogGroupName()
	ResetRoleArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference
type jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) LargeDataDeliveryS3Config() BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3ConfigList {
	var returns BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3ConfigList
	_jsii_.Get(
		j,
		"largeDataDeliveryS3Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) LargeDataDeliveryS3ConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"largeDataDeliveryS3ConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) LogGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockModelInvocationLoggingConfiguration.BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference_Override(b BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockModelInvocationLoggingConfiguration.BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference)SetLogGroupName(val *string) {
	if err := j.validateSetLogGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroupName",
		val,
	)
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) PutLargeDataDeliveryS3Config(value interface{}) {
	if err := b.validatePutLargeDataDeliveryS3ConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLargeDataDeliveryS3Config",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) ResetLargeDataDeliveryS3Config() {
	_jsii_.InvokeVoid(
		b,
		"resetLargeDataDeliveryS3Config",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) ResetLogGroupName() {
	_jsii_.InvokeVoid(
		b,
		"resetLogGroupName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		b,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

