package cognitoriskconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/cognitoriskconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference interface {
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
	EventAction() *string
	SetEventAction(val *string)
	EventActionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighAction
	SetInternalValue(val *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighAction)
	Notify() interface{}
	SetNotify(val interface{})
	NotifyInput() interface{}
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference
type jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) EventAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) EventActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) InternalValue() *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighAction {
	var returns *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) Notify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) NotifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference {
	_init_.Initialize()

	if err := validateNewCognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoRiskConfiguration.CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference_Override(c CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoRiskConfiguration.CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference)SetEventAction(val *string) {
	if err := j.validateSetEventActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventAction",
		val,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference)SetInternalValue(val *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference)SetNotify(val interface{}) {
	if err := j.validateSetNotifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notify",
		val,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

