// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmonitoringsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/cloudfrontmonitoringsubscription/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference interface {
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
	InternalValue() *CloudfrontMonitoringSubscriptionMonitoringSubscription
	SetInternalValue(val *CloudfrontMonitoringSubscriptionMonitoringSubscription)
	RealtimeMetricsSubscriptionConfig() CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference
	RealtimeMetricsSubscriptionConfigInput() *CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig
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
	PutRealtimeMetricsSubscriptionConfig(value *CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference
type jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) InternalValue() *CloudfrontMonitoringSubscriptionMonitoringSubscription {
	var returns *CloudfrontMonitoringSubscriptionMonitoringSubscription
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) RealtimeMetricsSubscriptionConfig() CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference {
	var returns CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference
	_jsii_.Get(
		j,
		"realtimeMetricsSubscriptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) RealtimeMetricsSubscriptionConfigInput() *CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig {
	var returns *CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig
	_jsii_.Get(
		j,
		"realtimeMetricsSubscriptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontMonitoringSubscription.CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference_Override(c CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontMonitoringSubscription.CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference)SetInternalValue(val *CloudfrontMonitoringSubscriptionMonitoringSubscription) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) PutRealtimeMetricsSubscriptionConfig(value *CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig) {
	if err := c.validatePutRealtimeMetricsSubscriptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRealtimeMetricsSubscriptionConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

