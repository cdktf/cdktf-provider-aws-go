// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitoriskconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/cognitoriskconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference interface {
	cdktf.ComplexObject
	Actions() CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationActionsOutputReference
	ActionsInput() *CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationActions
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
	EventFilter() *[]*string
	SetEventFilter(val *[]*string)
	EventFilterInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CognitoRiskConfigurationCompromisedCredentialsRiskConfiguration
	SetInternalValue(val *CognitoRiskConfigurationCompromisedCredentialsRiskConfiguration)
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
	PutActions(value *CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationActions)
	ResetEventFilter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference
type jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) Actions() CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationActionsOutputReference {
	var returns CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationActionsOutputReference
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) ActionsInput() *CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationActions {
	var returns *CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationActions
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) EventFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) EventFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) InternalValue() *CognitoRiskConfigurationCompromisedCredentialsRiskConfiguration {
	var returns *CognitoRiskConfigurationCompromisedCredentialsRiskConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewCognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoRiskConfiguration.CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference_Override(c CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoRiskConfiguration.CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference)SetEventFilter(val *[]*string) {
	if err := j.validateSetEventFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventFilter",
		val,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference)SetInternalValue(val *CognitoRiskConfigurationCompromisedCredentialsRiskConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) PutActions(value *CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationActions) {
	if err := c.validatePutActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putActions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) ResetEventFilter() {
	_jsii_.InvokeVoid(
		c,
		"resetEventFilter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CognitoRiskConfigurationCompromisedCredentialsRiskConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

