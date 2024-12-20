// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontvpcorigin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/cloudfrontvpcorigin/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	HttpPort() *float64
	SetHttpPort(val *float64)
	HttpPortInput() *float64
	HttpsPort() *float64
	SetHttpsPort(val *float64)
	HttpsPortInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	OriginProtocolPolicy() *string
	SetOriginProtocolPolicy(val *string)
	OriginProtocolPolicyInput() *string
	OriginSslProtocols() CloudfrontVpcOriginVpcOriginEndpointConfigOriginSslProtocolsList
	OriginSslProtocolsInput() interface{}
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
	PutOriginSslProtocols(value interface{})
	ResetOriginSslProtocols()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference
type jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) HttpPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) HttpPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) HttpsPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpsPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) HttpsPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpsPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) OriginProtocolPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originProtocolPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) OriginProtocolPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originProtocolPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) OriginSslProtocols() CloudfrontVpcOriginVpcOriginEndpointConfigOriginSslProtocolsList {
	var returns CloudfrontVpcOriginVpcOriginEndpointConfigOriginSslProtocolsList
	_jsii_.Get(
		j,
		"originSslProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) OriginSslProtocolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originSslProtocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudfrontVpcOriginVpcOriginEndpointConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontVpcOriginVpcOriginEndpointConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontVpcOrigin.CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudfrontVpcOriginVpcOriginEndpointConfigOutputReference_Override(c CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontVpcOrigin.CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference)SetArn(val *string) {
	if err := j.validateSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference)SetHttpPort(val *float64) {
	if err := j.validateSetHttpPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpPort",
		val,
	)
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference)SetHttpsPort(val *float64) {
	if err := j.validateSetHttpsPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsPort",
		val,
	)
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference)SetOriginProtocolPolicy(val *string) {
	if err := j.validateSetOriginProtocolPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originProtocolPolicy",
		val,
	)
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) PutOriginSslProtocols(value interface{}) {
	if err := c.validatePutOriginSslProtocolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOriginSslProtocols",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) ResetOriginSslProtocols() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginSslProtocols",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudfrontVpcOriginVpcOriginEndpointConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

