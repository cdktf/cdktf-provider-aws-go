package backupselection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/backupselection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupSelectionConditionOutputReference interface {
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
	StringEquals() BackupSelectionConditionStringEqualsList
	StringEqualsInput() interface{}
	StringLike() BackupSelectionConditionStringLikeList
	StringLikeInput() interface{}
	StringNotEquals() BackupSelectionConditionStringNotEqualsList
	StringNotEqualsInput() interface{}
	StringNotLike() BackupSelectionConditionStringNotLikeList
	StringNotLikeInput() interface{}
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
	PutStringEquals(value interface{})
	PutStringLike(value interface{})
	PutStringNotEquals(value interface{})
	PutStringNotLike(value interface{})
	ResetStringEquals()
	ResetStringLike()
	ResetStringNotEquals()
	ResetStringNotLike()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupSelectionConditionOutputReference
type jsiiProxy_BackupSelectionConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference) StringEquals() BackupSelectionConditionStringEqualsList {
	var returns BackupSelectionConditionStringEqualsList
	_jsii_.Get(
		j,
		"stringEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference) StringEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference) StringLike() BackupSelectionConditionStringLikeList {
	var returns BackupSelectionConditionStringLikeList
	_jsii_.Get(
		j,
		"stringLike",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference) StringLikeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringLikeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference) StringNotEquals() BackupSelectionConditionStringNotEqualsList {
	var returns BackupSelectionConditionStringNotEqualsList
	_jsii_.Get(
		j,
		"stringNotEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference) StringNotEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference) StringNotLike() BackupSelectionConditionStringNotLikeList {
	var returns BackupSelectionConditionStringNotLikeList
	_jsii_.Get(
		j,
		"stringNotLike",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference) StringNotLikeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotLikeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBackupSelectionConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BackupSelectionConditionOutputReference {
	_init_.Initialize()

	if err := validateNewBackupSelectionConditionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupSelectionConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.backupSelection.BackupSelectionConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBackupSelectionConditionOutputReference_Override(b BackupSelectionConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.backupSelection.BackupSelectionConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupSelectionConditionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) PutStringEquals(value interface{}) {
	if err := b.validatePutStringEqualsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStringEquals",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) PutStringLike(value interface{}) {
	if err := b.validatePutStringLikeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStringLike",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) PutStringNotEquals(value interface{}) {
	if err := b.validatePutStringNotEqualsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStringNotEquals",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) PutStringNotLike(value interface{}) {
	if err := b.validatePutStringNotLikeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStringNotLike",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) ResetStringEquals() {
	_jsii_.InvokeVoid(
		b,
		"resetStringEquals",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) ResetStringLike() {
	_jsii_.InvokeVoid(
		b,
		"resetStringLike",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) ResetStringNotEquals() {
	_jsii_.InvokeVoid(
		b,
		"resetStringNotEquals",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) ResetStringNotLike() {
	_jsii_.InvokeVoid(
		b,
		"resetStringNotLike",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupSelectionConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

