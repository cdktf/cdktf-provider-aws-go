package macie2classificationjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/macie2classificationjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Macie2ClassificationJobS3JobDefinitionOutputReference interface {
	cdktf.ComplexObject
	BucketCriteria() Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference
	BucketCriteriaInput() *Macie2ClassificationJobS3JobDefinitionBucketCriteria
	BucketDefinitions() Macie2ClassificationJobS3JobDefinitionBucketDefinitionsList
	BucketDefinitionsInput() interface{}
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
	InternalValue() *Macie2ClassificationJobS3JobDefinition
	SetInternalValue(val *Macie2ClassificationJobS3JobDefinition)
	Scoping() Macie2ClassificationJobS3JobDefinitionScopingOutputReference
	ScopingInput() *Macie2ClassificationJobS3JobDefinitionScoping
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
	PutBucketCriteria(value *Macie2ClassificationJobS3JobDefinitionBucketCriteria)
	PutBucketDefinitions(value interface{})
	PutScoping(value *Macie2ClassificationJobS3JobDefinitionScoping)
	ResetBucketCriteria()
	ResetBucketDefinitions()
	ResetScoping()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Macie2ClassificationJobS3JobDefinitionOutputReference
type jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) BucketCriteria() Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference {
	var returns Macie2ClassificationJobS3JobDefinitionBucketCriteriaOutputReference
	_jsii_.Get(
		j,
		"bucketCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) BucketCriteriaInput() *Macie2ClassificationJobS3JobDefinitionBucketCriteria {
	var returns *Macie2ClassificationJobS3JobDefinitionBucketCriteria
	_jsii_.Get(
		j,
		"bucketCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) BucketDefinitions() Macie2ClassificationJobS3JobDefinitionBucketDefinitionsList {
	var returns Macie2ClassificationJobS3JobDefinitionBucketDefinitionsList
	_jsii_.Get(
		j,
		"bucketDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) BucketDefinitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bucketDefinitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) InternalValue() *Macie2ClassificationJobS3JobDefinition {
	var returns *Macie2ClassificationJobS3JobDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) Scoping() Macie2ClassificationJobS3JobDefinitionScopingOutputReference {
	var returns Macie2ClassificationJobS3JobDefinitionScopingOutputReference
	_jsii_.Get(
		j,
		"scoping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) ScopingInput() *Macie2ClassificationJobS3JobDefinitionScoping {
	var returns *Macie2ClassificationJobS3JobDefinitionScoping
	_jsii_.Get(
		j,
		"scopingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMacie2ClassificationJobS3JobDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Macie2ClassificationJobS3JobDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewMacie2ClassificationJobS3JobDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.macie2ClassificationJob.Macie2ClassificationJobS3JobDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMacie2ClassificationJobS3JobDefinitionOutputReference_Override(m Macie2ClassificationJobS3JobDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.macie2ClassificationJob.Macie2ClassificationJobS3JobDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference)SetInternalValue(val *Macie2ClassificationJobS3JobDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) PutBucketCriteria(value *Macie2ClassificationJobS3JobDefinitionBucketCriteria) {
	if err := m.validatePutBucketCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBucketCriteria",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) PutBucketDefinitions(value interface{}) {
	if err := m.validatePutBucketDefinitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBucketDefinitions",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) PutScoping(value *Macie2ClassificationJobS3JobDefinitionScoping) {
	if err := m.validatePutScopingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putScoping",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) ResetBucketCriteria() {
	_jsii_.InvokeVoid(
		m,
		"resetBucketCriteria",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) ResetBucketDefinitions() {
	_jsii_.InvokeVoid(
		m,
		"resetBucketDefinitions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) ResetScoping() {
	_jsii_.InvokeVoid(
		m,
		"resetScoping",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_Macie2ClassificationJobS3JobDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

