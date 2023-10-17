// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chimesdkmediapipelinesmediainsightspipelineconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/chimesdkmediapipelinesmediainsightspipelineconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference interface {
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
	Destination() *string
	SetDestination(val *string)
	DestinationInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfiguration
	SetInternalValue(val *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfiguration)
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
	ResetDestination()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference
type jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) DestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) InternalValue() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfiguration {
	var returns *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.chimesdkmediapipelinesMediaInsightsPipelineConfiguration.ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference_Override(c ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.chimesdkmediapipelinesMediaInsightsPipelineConfiguration.ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference)SetDestination(val *string) {
	if err := j.validateSetDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destination",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference)SetInternalValue(val *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) ResetDestination() {
	_jsii_.InvokeVoid(
		c,
		"resetDestination",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

