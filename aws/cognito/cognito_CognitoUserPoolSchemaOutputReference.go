package cognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/cognito/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CognitoUserPoolSchemaOutputReference interface {
	cdktf.ComplexObject
	AttributeDataType() *string
	SetAttributeDataType(val *string)
	AttributeDataTypeInput() *string
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
	DeveloperOnlyAttribute() interface{}
	SetDeveloperOnlyAttribute(val interface{})
	DeveloperOnlyAttributeInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Mutable() interface{}
	SetMutable(val interface{})
	MutableInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NumberAttributeConstraints() CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference
	NumberAttributeConstraintsInput() *CognitoUserPoolSchemaNumberAttributeConstraints
	Required() interface{}
	SetRequired(val interface{})
	RequiredInput() interface{}
	StringAttributeConstraints() CognitoUserPoolSchemaStringAttributeConstraintsOutputReference
	StringAttributeConstraintsInput() *CognitoUserPoolSchemaStringAttributeConstraints
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
	PutNumberAttributeConstraints(value *CognitoUserPoolSchemaNumberAttributeConstraints)
	PutStringAttributeConstraints(value *CognitoUserPoolSchemaStringAttributeConstraints)
	ResetDeveloperOnlyAttribute()
	ResetMutable()
	ResetNumberAttributeConstraints()
	ResetRequired()
	ResetStringAttributeConstraints()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CognitoUserPoolSchemaOutputReference
type jsiiProxy_CognitoUserPoolSchemaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) AttributeDataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeDataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) AttributeDataTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeDataTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) DeveloperOnlyAttribute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"developerOnlyAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) DeveloperOnlyAttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"developerOnlyAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) Mutable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) MutableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mutableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) NumberAttributeConstraints() CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference {
	var returns CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference
	_jsii_.Get(
		j,
		"numberAttributeConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) NumberAttributeConstraintsInput() *CognitoUserPoolSchemaNumberAttributeConstraints {
	var returns *CognitoUserPoolSchemaNumberAttributeConstraints
	_jsii_.Get(
		j,
		"numberAttributeConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) Required() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) RequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) StringAttributeConstraints() CognitoUserPoolSchemaStringAttributeConstraintsOutputReference {
	var returns CognitoUserPoolSchemaStringAttributeConstraintsOutputReference
	_jsii_.Get(
		j,
		"stringAttributeConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) StringAttributeConstraintsInput() *CognitoUserPoolSchemaStringAttributeConstraints {
	var returns *CognitoUserPoolSchemaStringAttributeConstraints
	_jsii_.Get(
		j,
		"stringAttributeConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCognitoUserPoolSchemaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CognitoUserPoolSchemaOutputReference {
	_init_.Initialize()

	if err := validateNewCognitoUserPoolSchemaOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoUserPoolSchemaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cognito.CognitoUserPoolSchemaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCognitoUserPoolSchemaOutputReference_Override(c CognitoUserPoolSchemaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cognito.CognitoUserPoolSchemaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference)SetAttributeDataType(val *string) {
	if err := j.validateSetAttributeDataTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeDataType",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference)SetDeveloperOnlyAttribute(val interface{}) {
	if err := j.validateSetDeveloperOnlyAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developerOnlyAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference)SetMutable(val interface{}) {
	if err := j.validateSetMutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mutable",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference)SetRequired(val interface{}) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolSchemaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) PutNumberAttributeConstraints(value *CognitoUserPoolSchemaNumberAttributeConstraints) {
	if err := c.validatePutNumberAttributeConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNumberAttributeConstraints",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) PutStringAttributeConstraints(value *CognitoUserPoolSchemaStringAttributeConstraints) {
	if err := c.validatePutStringAttributeConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStringAttributeConstraints",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) ResetDeveloperOnlyAttribute() {
	_jsii_.InvokeVoid(
		c,
		"resetDeveloperOnlyAttribute",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) ResetMutable() {
	_jsii_.InvokeVoid(
		c,
		"resetMutable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) ResetNumberAttributeConstraints() {
	_jsii_.InvokeVoid(
		c,
		"resetNumberAttributeConstraints",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) ResetRequired() {
	_jsii_.InvokeVoid(
		c,
		"resetRequired",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) ResetStringAttributeConstraints() {
	_jsii_.InvokeVoid(
		c,
		"resetStringAttributeConstraints",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CognitoUserPoolSchemaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

