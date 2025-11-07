// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/pipespipe/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipesPipeLogConfigurationOutputReference interface {
	cdktf.ComplexObject
	CloudwatchLogsLogDestination() PipesPipeLogConfigurationCloudwatchLogsLogDestinationOutputReference
	CloudwatchLogsLogDestinationInput() *PipesPipeLogConfigurationCloudwatchLogsLogDestination
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
	FirehoseLogDestination() PipesPipeLogConfigurationFirehoseLogDestinationOutputReference
	FirehoseLogDestinationInput() *PipesPipeLogConfigurationFirehoseLogDestination
	// Experimental.
	Fqn() *string
	IncludeExecutionData() *[]*string
	SetIncludeExecutionData(val *[]*string)
	IncludeExecutionDataInput() *[]*string
	InternalValue() *PipesPipeLogConfiguration
	SetInternalValue(val *PipesPipeLogConfiguration)
	Level() *string
	SetLevel(val *string)
	LevelInput() *string
	S3LogDestination() PipesPipeLogConfigurationS3LogDestinationOutputReference
	S3LogDestinationInput() *PipesPipeLogConfigurationS3LogDestination
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
	PutCloudwatchLogsLogDestination(value *PipesPipeLogConfigurationCloudwatchLogsLogDestination)
	PutFirehoseLogDestination(value *PipesPipeLogConfigurationFirehoseLogDestination)
	PutS3LogDestination(value *PipesPipeLogConfigurationS3LogDestination)
	ResetCloudwatchLogsLogDestination()
	ResetFirehoseLogDestination()
	ResetIncludeExecutionData()
	ResetS3LogDestination()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipesPipeLogConfigurationOutputReference
type jsiiProxy_PipesPipeLogConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference) CloudwatchLogsLogDestination() PipesPipeLogConfigurationCloudwatchLogsLogDestinationOutputReference {
	var returns PipesPipeLogConfigurationCloudwatchLogsLogDestinationOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogsLogDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference) CloudwatchLogsLogDestinationInput() *PipesPipeLogConfigurationCloudwatchLogsLogDestination {
	var returns *PipesPipeLogConfigurationCloudwatchLogsLogDestination
	_jsii_.Get(
		j,
		"cloudwatchLogsLogDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference) FirehoseLogDestination() PipesPipeLogConfigurationFirehoseLogDestinationOutputReference {
	var returns PipesPipeLogConfigurationFirehoseLogDestinationOutputReference
	_jsii_.Get(
		j,
		"firehoseLogDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference) FirehoseLogDestinationInput() *PipesPipeLogConfigurationFirehoseLogDestination {
	var returns *PipesPipeLogConfigurationFirehoseLogDestination
	_jsii_.Get(
		j,
		"firehoseLogDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference) IncludeExecutionData() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeExecutionData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference) IncludeExecutionDataInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeExecutionDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference) InternalValue() *PipesPipeLogConfiguration {
	var returns *PipesPipeLogConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference) S3LogDestination() PipesPipeLogConfigurationS3LogDestinationOutputReference {
	var returns PipesPipeLogConfigurationS3LogDestinationOutputReference
	_jsii_.Get(
		j,
		"s3LogDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference) S3LogDestinationInput() *PipesPipeLogConfigurationS3LogDestination {
	var returns *PipesPipeLogConfigurationS3LogDestination
	_jsii_.Get(
		j,
		"s3LogDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipesPipeLogConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PipesPipeLogConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewPipesPipeLogConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipesPipeLogConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeLogConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipesPipeLogConfigurationOutputReference_Override(p PipesPipeLogConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.pipesPipe.PipesPipeLogConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference)SetIncludeExecutionData(val *[]*string) {
	if err := j.validateSetIncludeExecutionDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeExecutionData",
		val,
	)
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference)SetInternalValue(val *PipesPipeLogConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipesPipeLogConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) PutCloudwatchLogsLogDestination(value *PipesPipeLogConfigurationCloudwatchLogsLogDestination) {
	if err := p.validatePutCloudwatchLogsLogDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCloudwatchLogsLogDestination",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) PutFirehoseLogDestination(value *PipesPipeLogConfigurationFirehoseLogDestination) {
	if err := p.validatePutFirehoseLogDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFirehoseLogDestination",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) PutS3LogDestination(value *PipesPipeLogConfigurationS3LogDestination) {
	if err := p.validatePutS3LogDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putS3LogDestination",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) ResetCloudwatchLogsLogDestination() {
	_jsii_.InvokeVoid(
		p,
		"resetCloudwatchLogsLogDestination",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) ResetFirehoseLogDestination() {
	_jsii_.InvokeVoid(
		p,
		"resetFirehoseLogDestination",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) ResetIncludeExecutionData() {
	_jsii_.InvokeVoid(
		p,
		"resetIncludeExecutionData",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) ResetS3LogDestination() {
	_jsii_.InvokeVoid(
		p,
		"resetS3LogDestination",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeLogConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

