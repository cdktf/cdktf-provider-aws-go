package macie2classificationjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/macie2classificationjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference interface {
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
	Excludes() Macie2ClassificationJobS3JobDefinitionBucketCriteriaExcludesOutputReference
	ExcludesInput() *Macie2ClassificationJobS3JobDefinitionBucketCriteriaExcludes
	// Experimental.
	Fqn() *string
	Includes() Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesOutputReference
	IncludesInput() *Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludes
	InternalValue() *Macie2ClassificationJobS3JobDefinitionBucketCriteria
	SetInternalValue(val *Macie2ClassificationJobS3JobDefinitionBucketCriteria)
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
	PutExcludes(value *Macie2ClassificationJobS3JobDefinitionBucketCriteriaExcludes)
	PutIncludes(value *Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludes)
	ResetExcludes()
	ResetIncludes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference
type jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) Excludes() Macie2ClassificationJobS3JobDefinitionBucketCriteriaExcludesOutputReference {
	var returns Macie2ClassificationJobS3JobDefinitionBucketCriteriaExcludesOutputReference
	_jsii_.Get(
		j,
		"excludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) ExcludesInput() *Macie2ClassificationJobS3JobDefinitionBucketCriteriaExcludes {
	var returns *Macie2ClassificationJobS3JobDefinitionBucketCriteriaExcludes
	_jsii_.Get(
		j,
		"excludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) Includes() Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesOutputReference {
	var returns Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludesOutputReference
	_jsii_.Get(
		j,
		"includes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) IncludesInput() *Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludes {
	var returns *Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludes
	_jsii_.Get(
		j,
		"includesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) InternalValue() *Macie2ClassificationJobS3JobDefinitionBucketCriteria {
	var returns *Macie2ClassificationJobS3JobDefinitionBucketCriteria
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMacie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewMacie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.macie2ClassificationJob.Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMacie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference_Override(m Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.macie2ClassificationJob.Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference)SetInternalValue(val *Macie2ClassificationJobS3JobDefinitionBucketCriteria) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) PutExcludes(value *Macie2ClassificationJobS3JobDefinitionBucketCriteriaExcludes) {
	if err := m.validatePutExcludesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putExcludes",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) PutIncludes(value *Macie2ClassificationJobS3JobDefinitionBucketCriteriaIncludes) {
	if err := m.validatePutIncludesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putIncludes",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) ResetExcludes() {
	_jsii_.InvokeVoid(
		m,
		"resetExcludes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) ResetIncludes() {
	_jsii_.InvokeVoid(
		m,
		"resetIncludes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

