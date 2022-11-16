package macie2classificationjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/macie2classificationjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SimpleScopeTerm() Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference
	SimpleScopeTermInput() *Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTerm
	TagScopeTerm() Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference
	TagScopeTermInput() *Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTerm
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
	PutSimpleScopeTerm(value *Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTerm)
	PutTagScopeTerm(value *Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTerm)
	ResetSimpleScopeTerm()
	ResetTagScopeTerm()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference
type jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) SimpleScopeTerm() Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference {
	var returns Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTermOutputReference
	_jsii_.Get(
		j,
		"simpleScopeTerm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) SimpleScopeTermInput() *Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTerm {
	var returns *Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTerm
	_jsii_.Get(
		j,
		"simpleScopeTermInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) TagScopeTerm() Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference {
	var returns Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTermOutputReference
	_jsii_.Get(
		j,
		"tagScopeTerm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) TagScopeTermInput() *Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTerm {
	var returns *Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTerm
	_jsii_.Get(
		j,
		"tagScopeTermInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMacie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference {
	_init_.Initialize()

	if err := validateNewMacie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.macie2ClassificationJob.Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMacie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference_Override(m Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.macie2ClassificationJob.Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) PutSimpleScopeTerm(value *Macie2ClassificationJobS3JobDefinitionScopingExcludesAndSimpleScopeTerm) {
	if err := m.validatePutSimpleScopeTermParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSimpleScopeTerm",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) PutTagScopeTerm(value *Macie2ClassificationJobS3JobDefinitionScopingExcludesAndTagScopeTerm) {
	if err := m.validatePutTagScopeTermParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTagScopeTerm",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) ResetSimpleScopeTerm() {
	_jsii_.InvokeVoid(
		m,
		"resetSimpleScopeTerm",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) ResetTagScopeTerm() {
	_jsii_.InvokeVoid(
		m,
		"resetTagScopeTerm",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionScopingExcludesAndOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

