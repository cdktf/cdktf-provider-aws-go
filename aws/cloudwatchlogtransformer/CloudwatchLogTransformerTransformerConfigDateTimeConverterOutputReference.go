// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogtransformer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/cloudwatchlogtransformer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference interface {
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
	Locale() *string
	SetLocale(val *string)
	LocaleInput() *string
	MatchPatterns() *[]*string
	SetMatchPatterns(val *[]*string)
	MatchPatternsInput() *[]*string
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	SourceTimezone() *string
	SetSourceTimezone(val *string)
	SourceTimezoneInput() *string
	Target() *string
	SetTarget(val *string)
	TargetFormat() *string
	SetTargetFormat(val *string)
	TargetFormatInput() *string
	TargetInput() *string
	TargetTimezone() *string
	SetTargetTimezone(val *string)
	TargetTimezoneInput() *string
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
	ResetLocale()
	ResetSourceTimezone()
	ResetTargetFormat()
	ResetTargetTimezone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference
type jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) Locale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) LocaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) MatchPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) MatchPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) SourceTimezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) SourceTimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) TargetFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) TargetFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) TargetTimezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) TargetTimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference {
	_init_.Initialize()

	if err := validateNewCloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudwatchLogTransformer.CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference_Override(c CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudwatchLogTransformer.CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference)SetLocale(val *string) {
	if err := j.validateSetLocaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locale",
		val,
	)
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference)SetMatchPatterns(val *[]*string) {
	if err := j.validateSetMatchPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchPatterns",
		val,
	)
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference)SetSourceTimezone(val *string) {
	if err := j.validateSetSourceTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTimezone",
		val,
	)
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference)SetTargetFormat(val *string) {
	if err := j.validateSetTargetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetFormat",
		val,
	)
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference)SetTargetTimezone(val *string) {
	if err := j.validateSetTargetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetTimezone",
		val,
	)
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) ResetLocale() {
	_jsii_.InvokeVoid(
		c,
		"resetLocale",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) ResetSourceTimezone() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceTimezone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) ResetTargetFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) ResetTargetTimezone() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetTimezone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigDateTimeConverterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

