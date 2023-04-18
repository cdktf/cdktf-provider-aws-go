package cognitoriskconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/cognitoriskconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference interface {
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
	HighAction() CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference
	HighActionInput() *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighAction
	InternalValue() *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActions
	SetInternalValue(val *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActions)
	LowAction() CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsLowActionOutputReference
	LowActionInput() *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsLowAction
	MediumAction() CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsMediumActionOutputReference
	MediumActionInput() *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsMediumAction
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
	PutHighAction(value *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighAction)
	PutLowAction(value *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsLowAction)
	PutMediumAction(value *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsMediumAction)
	ResetHighAction()
	ResetLowAction()
	ResetMediumAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference
type jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) HighAction() CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference {
	var returns CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighActionOutputReference
	_jsii_.Get(
		j,
		"highAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) HighActionInput() *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighAction {
	var returns *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighAction
	_jsii_.Get(
		j,
		"highActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) InternalValue() *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActions {
	var returns *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) LowAction() CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsLowActionOutputReference {
	var returns CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsLowActionOutputReference
	_jsii_.Get(
		j,
		"lowAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) LowActionInput() *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsLowAction {
	var returns *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsLowAction
	_jsii_.Get(
		j,
		"lowActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) MediumAction() CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsMediumActionOutputReference {
	var returns CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsMediumActionOutputReference
	_jsii_.Get(
		j,
		"mediumAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) MediumActionInput() *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsMediumAction {
	var returns *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsMediumAction
	_jsii_.Get(
		j,
		"mediumActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference {
	_init_.Initialize()

	if err := validateNewCognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoRiskConfiguration.CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference_Override(c CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoRiskConfiguration.CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference)SetInternalValue(val *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) PutHighAction(value *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsHighAction) {
	if err := c.validatePutHighActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHighAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) PutLowAction(value *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsLowAction) {
	if err := c.validatePutLowActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLowAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) PutMediumAction(value *CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsMediumAction) {
	if err := c.validatePutMediumActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMediumAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) ResetHighAction() {
	_jsii_.InvokeVoid(
		c,
		"resetHighAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) ResetLowAction() {
	_jsii_.InvokeVoid(
		c,
		"resetLowAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) ResetMediumAction() {
	_jsii_.InvokeVoid(
		c,
		"resetMediumAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CognitoRiskConfigurationAccountTakeoverRiskConfigurationActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

