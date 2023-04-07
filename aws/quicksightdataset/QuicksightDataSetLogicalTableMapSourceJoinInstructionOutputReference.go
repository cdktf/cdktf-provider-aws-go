package quicksightdataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/quicksightdataset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference interface {
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
	InternalValue() *QuicksightDataSetLogicalTableMapSourceJoinInstruction
	SetInternalValue(val *QuicksightDataSetLogicalTableMapSourceJoinInstruction)
	LeftJoinKeyProperties() QuicksightDataSetLogicalTableMapSourceJoinInstructionLeftJoinKeyPropertiesOutputReference
	LeftJoinKeyPropertiesInput() *QuicksightDataSetLogicalTableMapSourceJoinInstructionLeftJoinKeyProperties
	LeftOperand() *string
	SetLeftOperand(val *string)
	LeftOperandInput() *string
	OnClause() *string
	SetOnClause(val *string)
	OnClauseInput() *string
	RightJoinKeyProperties() QuicksightDataSetLogicalTableMapSourceJoinInstructionRightJoinKeyPropertiesOutputReference
	RightJoinKeyPropertiesInput() *QuicksightDataSetLogicalTableMapSourceJoinInstructionRightJoinKeyProperties
	RightOperand() *string
	SetRightOperand(val *string)
	RightOperandInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutLeftJoinKeyProperties(value *QuicksightDataSetLogicalTableMapSourceJoinInstructionLeftJoinKeyProperties)
	PutRightJoinKeyProperties(value *QuicksightDataSetLogicalTableMapSourceJoinInstructionRightJoinKeyProperties)
	ResetLeftJoinKeyProperties()
	ResetRightJoinKeyProperties()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference
type jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) InternalValue() *QuicksightDataSetLogicalTableMapSourceJoinInstruction {
	var returns *QuicksightDataSetLogicalTableMapSourceJoinInstruction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) LeftJoinKeyProperties() QuicksightDataSetLogicalTableMapSourceJoinInstructionLeftJoinKeyPropertiesOutputReference {
	var returns QuicksightDataSetLogicalTableMapSourceJoinInstructionLeftJoinKeyPropertiesOutputReference
	_jsii_.Get(
		j,
		"leftJoinKeyProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) LeftJoinKeyPropertiesInput() *QuicksightDataSetLogicalTableMapSourceJoinInstructionLeftJoinKeyProperties {
	var returns *QuicksightDataSetLogicalTableMapSourceJoinInstructionLeftJoinKeyProperties
	_jsii_.Get(
		j,
		"leftJoinKeyPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) LeftOperand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leftOperand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) LeftOperandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leftOperandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) OnClause() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onClause",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) OnClauseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onClauseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) RightJoinKeyProperties() QuicksightDataSetLogicalTableMapSourceJoinInstructionRightJoinKeyPropertiesOutputReference {
	var returns QuicksightDataSetLogicalTableMapSourceJoinInstructionRightJoinKeyPropertiesOutputReference
	_jsii_.Get(
		j,
		"rightJoinKeyProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) RightJoinKeyPropertiesInput() *QuicksightDataSetLogicalTableMapSourceJoinInstructionRightJoinKeyProperties {
	var returns *QuicksightDataSetLogicalTableMapSourceJoinInstructionRightJoinKeyProperties
	_jsii_.Get(
		j,
		"rightJoinKeyPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) RightOperand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rightOperand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) RightOperandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rightOperandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewQuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightDataSet.QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference_Override(q QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightDataSet.QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference)SetInternalValue(val *QuicksightDataSetLogicalTableMapSourceJoinInstruction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference)SetLeftOperand(val *string) {
	if err := j.validateSetLeftOperandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leftOperand",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference)SetOnClause(val *string) {
	if err := j.validateSetOnClauseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onClause",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference)SetRightOperand(val *string) {
	if err := j.validateSetRightOperandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rightOperand",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) PutLeftJoinKeyProperties(value *QuicksightDataSetLogicalTableMapSourceJoinInstructionLeftJoinKeyProperties) {
	if err := q.validatePutLeftJoinKeyPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putLeftJoinKeyProperties",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) PutRightJoinKeyProperties(value *QuicksightDataSetLogicalTableMapSourceJoinInstructionRightJoinKeyProperties) {
	if err := q.validatePutRightJoinKeyPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putRightJoinKeyProperties",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) ResetLeftJoinKeyProperties() {
	_jsii_.InvokeVoid(
		q,
		"resetLeftJoinKeyProperties",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) ResetRightJoinKeyProperties() {
	_jsii_.InvokeVoid(
		q,
		"resetRightJoinKeyProperties",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := q.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSetLogicalTableMapSourceJoinInstructionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

