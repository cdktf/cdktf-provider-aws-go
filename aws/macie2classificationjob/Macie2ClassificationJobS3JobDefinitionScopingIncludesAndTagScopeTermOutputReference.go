package macie2classificationjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/macie2classificationjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference interface {
	cdktf.ComplexObject
	Comparator() *string
	SetComparator(val *string)
	ComparatorInput() *string
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
	InternalValue() *Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTerm
	SetInternalValue(val *Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTerm)
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	TagValues() Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermTagValuesList
	TagValuesInput() interface{}
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
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
	PutTagValues(value interface{})
	ResetComparator()
	ResetKey()
	ResetTagValues()
	ResetTarget()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference
type jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) Comparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) ComparatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) InternalValue() *Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTerm {
	var returns *Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTerm
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) TagValues() Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermTagValuesList {
	var returns Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermTagValuesList
	_jsii_.Get(
		j,
		"tagValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) TagValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMacie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference {
	_init_.Initialize()

	if err := validateNewMacie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.macie2ClassificationJob.Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMacie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference_Override(m Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.macie2ClassificationJob.Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference)SetComparator(val *string) {
	if err := j.validateSetComparatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparator",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference)SetInternalValue(val *Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTerm) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) PutTagValues(value interface{}) {
	if err := m.validatePutTagValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTagValues",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) ResetComparator() {
	_jsii_.InvokeVoid(
		m,
		"resetComparator",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		m,
		"resetKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) ResetTagValues() {
	_jsii_.InvokeVoid(
		m,
		"resetTagValues",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		m,
		"resetTarget",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingIncludesAndTagScopeTermOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

