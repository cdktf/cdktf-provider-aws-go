// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontresponseheaderspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/cloudfrontresponseheaderspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference interface {
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
	InternalValue() *CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection
	SetInternalValue(val *CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection)
	ModeBlock() interface{}
	SetModeBlock(val interface{})
	ModeBlockInput() interface{}
	Override() interface{}
	SetOverride(val interface{})
	OverrideInput() interface{}
	Protection() interface{}
	SetProtection(val interface{})
	ProtectionInput() interface{}
	ReportUri() *string
	SetReportUri(val *string)
	ReportUriInput() *string
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
	ResetModeBlock()
	ResetReportUri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) InternalValue() *CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection {
	var returns *CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ModeBlock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modeBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ModeBlockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modeBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) Override() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) OverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) Protection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ReportUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ReportUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontResponseHeadersPolicy.CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference_Override(c CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontResponseHeadersPolicy.CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference)SetInternalValue(val *CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference)SetModeBlock(val interface{}) {
	if err := j.validateSetModeBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modeBlock",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference)SetOverride(val interface{}) {
	if err := j.validateSetOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"override",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference)SetProtection(val interface{}) {
	if err := j.validateSetProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protection",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference)SetReportUri(val *string) {
	if err := j.validateSetReportUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportUri",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ResetModeBlock() {
	_jsii_.InvokeVoid(
		c,
		"resetModeBlock",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ResetReportUri() {
	_jsii_.InvokeVoid(
		c,
		"resetReportUri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

