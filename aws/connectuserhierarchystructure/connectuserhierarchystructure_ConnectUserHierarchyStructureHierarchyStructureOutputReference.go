package connectuserhierarchystructure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/connectuserhierarchystructure/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConnectUserHierarchyStructureHierarchyStructureOutputReference interface {
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
	InternalValue() *ConnectUserHierarchyStructureHierarchyStructure
	SetInternalValue(val *ConnectUserHierarchyStructureHierarchyStructure)
	LevelFive() ConnectUserHierarchyStructureHierarchyStructureLevelFiveOutputReference
	LevelFiveInput() *ConnectUserHierarchyStructureHierarchyStructureLevelFive
	LevelFour() ConnectUserHierarchyStructureHierarchyStructureLevelFourOutputReference
	LevelFourInput() *ConnectUserHierarchyStructureHierarchyStructureLevelFour
	LevelOne() ConnectUserHierarchyStructureHierarchyStructureLevelOneOutputReference
	LevelOneInput() *ConnectUserHierarchyStructureHierarchyStructureLevelOne
	LevelThree() ConnectUserHierarchyStructureHierarchyStructureLevelThreeOutputReference
	LevelThreeInput() *ConnectUserHierarchyStructureHierarchyStructureLevelThree
	LevelTwo() ConnectUserHierarchyStructureHierarchyStructureLevelTwoOutputReference
	LevelTwoInput() *ConnectUserHierarchyStructureHierarchyStructureLevelTwo
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
	PutLevelFive(value *ConnectUserHierarchyStructureHierarchyStructureLevelFive)
	PutLevelFour(value *ConnectUserHierarchyStructureHierarchyStructureLevelFour)
	PutLevelOne(value *ConnectUserHierarchyStructureHierarchyStructureLevelOne)
	PutLevelThree(value *ConnectUserHierarchyStructureHierarchyStructureLevelThree)
	PutLevelTwo(value *ConnectUserHierarchyStructureHierarchyStructureLevelTwo)
	ResetLevelFive()
	ResetLevelFour()
	ResetLevelOne()
	ResetLevelThree()
	ResetLevelTwo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectUserHierarchyStructureHierarchyStructureOutputReference
type jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) InternalValue() *ConnectUserHierarchyStructureHierarchyStructure {
	var returns *ConnectUserHierarchyStructureHierarchyStructure
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) LevelFive() ConnectUserHierarchyStructureHierarchyStructureLevelFiveOutputReference {
	var returns ConnectUserHierarchyStructureHierarchyStructureLevelFiveOutputReference
	_jsii_.Get(
		j,
		"levelFive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) LevelFiveInput() *ConnectUserHierarchyStructureHierarchyStructureLevelFive {
	var returns *ConnectUserHierarchyStructureHierarchyStructureLevelFive
	_jsii_.Get(
		j,
		"levelFiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) LevelFour() ConnectUserHierarchyStructureHierarchyStructureLevelFourOutputReference {
	var returns ConnectUserHierarchyStructureHierarchyStructureLevelFourOutputReference
	_jsii_.Get(
		j,
		"levelFour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) LevelFourInput() *ConnectUserHierarchyStructureHierarchyStructureLevelFour {
	var returns *ConnectUserHierarchyStructureHierarchyStructureLevelFour
	_jsii_.Get(
		j,
		"levelFourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) LevelOne() ConnectUserHierarchyStructureHierarchyStructureLevelOneOutputReference {
	var returns ConnectUserHierarchyStructureHierarchyStructureLevelOneOutputReference
	_jsii_.Get(
		j,
		"levelOne",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) LevelOneInput() *ConnectUserHierarchyStructureHierarchyStructureLevelOne {
	var returns *ConnectUserHierarchyStructureHierarchyStructureLevelOne
	_jsii_.Get(
		j,
		"levelOneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) LevelThree() ConnectUserHierarchyStructureHierarchyStructureLevelThreeOutputReference {
	var returns ConnectUserHierarchyStructureHierarchyStructureLevelThreeOutputReference
	_jsii_.Get(
		j,
		"levelThree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) LevelThreeInput() *ConnectUserHierarchyStructureHierarchyStructureLevelThree {
	var returns *ConnectUserHierarchyStructureHierarchyStructureLevelThree
	_jsii_.Get(
		j,
		"levelThreeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) LevelTwo() ConnectUserHierarchyStructureHierarchyStructureLevelTwoOutputReference {
	var returns ConnectUserHierarchyStructureHierarchyStructureLevelTwoOutputReference
	_jsii_.Get(
		j,
		"levelTwo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) LevelTwoInput() *ConnectUserHierarchyStructureHierarchyStructureLevelTwo {
	var returns *ConnectUserHierarchyStructureHierarchyStructureLevelTwo
	_jsii_.Get(
		j,
		"levelTwoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectUserHierarchyStructureHierarchyStructureOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConnectUserHierarchyStructureHierarchyStructureOutputReference {
	_init_.Initialize()

	if err := validateNewConnectUserHierarchyStructureHierarchyStructureOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.connectUserHierarchyStructure.ConnectUserHierarchyStructureHierarchyStructureOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectUserHierarchyStructureHierarchyStructureOutputReference_Override(c ConnectUserHierarchyStructureHierarchyStructureOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.connectUserHierarchyStructure.ConnectUserHierarchyStructureHierarchyStructureOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference)SetInternalValue(val *ConnectUserHierarchyStructureHierarchyStructure) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) PutLevelFive(value *ConnectUserHierarchyStructureHierarchyStructureLevelFive) {
	if err := c.validatePutLevelFiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLevelFive",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) PutLevelFour(value *ConnectUserHierarchyStructureHierarchyStructureLevelFour) {
	if err := c.validatePutLevelFourParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLevelFour",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) PutLevelOne(value *ConnectUserHierarchyStructureHierarchyStructureLevelOne) {
	if err := c.validatePutLevelOneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLevelOne",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) PutLevelThree(value *ConnectUserHierarchyStructureHierarchyStructureLevelThree) {
	if err := c.validatePutLevelThreeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLevelThree",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) PutLevelTwo(value *ConnectUserHierarchyStructureHierarchyStructureLevelTwo) {
	if err := c.validatePutLevelTwoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLevelTwo",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) ResetLevelFive() {
	_jsii_.InvokeVoid(
		c,
		"resetLevelFive",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) ResetLevelFour() {
	_jsii_.InvokeVoid(
		c,
		"resetLevelFour",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) ResetLevelOne() {
	_jsii_.InvokeVoid(
		c,
		"resetLevelOne",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) ResetLevelThree() {
	_jsii_.InvokeVoid(
		c,
		"resetLevelThree",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) ResetLevelTwo() {
	_jsii_.InvokeVoid(
		c,
		"resetLevelTwo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConnectUserHierarchyStructureHierarchyStructureOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

