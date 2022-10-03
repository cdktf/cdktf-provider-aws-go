package dataawsiampolicydocument

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/dataawsiampolicydocument/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsIamPolicyDocumentStatementOutputReference interface {
	cdktf.ComplexObject
	Actions() *[]*string
	SetActions(val *[]*string)
	ActionsInput() *[]*string
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
	Condition() DataAwsIamPolicyDocumentStatementConditionList
	ConditionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Effect() *string
	SetEffect(val *string)
	EffectInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NotActions() *[]*string
	SetNotActions(val *[]*string)
	NotActionsInput() *[]*string
	NotPrincipals() DataAwsIamPolicyDocumentStatementNotPrincipalsList
	NotPrincipalsInput() interface{}
	NotResources() *[]*string
	SetNotResources(val *[]*string)
	NotResourcesInput() *[]*string
	Principals() DataAwsIamPolicyDocumentStatementPrincipalsList
	PrincipalsInput() interface{}
	Resources() *[]*string
	SetResources(val *[]*string)
	ResourcesInput() *[]*string
	Sid() *string
	SetSid(val *string)
	SidInput() *string
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
	PutCondition(value interface{})
	PutNotPrincipals(value interface{})
	PutPrincipals(value interface{})
	ResetActions()
	ResetCondition()
	ResetEffect()
	ResetNotActions()
	ResetNotPrincipals()
	ResetNotResources()
	ResetPrincipals()
	ResetResources()
	ResetSid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsIamPolicyDocumentStatementOutputReference
type jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) Actions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) ActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) Condition() DataAwsIamPolicyDocumentStatementConditionList {
	var returns DataAwsIamPolicyDocumentStatementConditionList
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) ConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) Effect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) EffectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) NotActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) NotActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) NotPrincipals() DataAwsIamPolicyDocumentStatementNotPrincipalsList {
	var returns DataAwsIamPolicyDocumentStatementNotPrincipalsList
	_jsii_.Get(
		j,
		"notPrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) NotPrincipalsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notPrincipalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) NotResources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) NotResourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) Principals() DataAwsIamPolicyDocumentStatementPrincipalsList {
	var returns DataAwsIamPolicyDocumentStatementPrincipalsList
	_jsii_.Get(
		j,
		"principals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) PrincipalsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"principalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) Resources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) ResourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) Sid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) SidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsIamPolicyDocumentStatementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsIamPolicyDocumentStatementOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsIamPolicyDocumentStatementOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocumentStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsIamPolicyDocumentStatementOutputReference_Override(d DataAwsIamPolicyDocumentStatementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsIamPolicyDocument.DataAwsIamPolicyDocumentStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference)SetActions(val *[]*string) {
	if err := j.validateSetActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference)SetEffect(val *string) {
	if err := j.validateSetEffectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effect",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference)SetNotActions(val *[]*string) {
	if err := j.validateSetNotActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notActions",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference)SetNotResources(val *[]*string) {
	if err := j.validateSetNotResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notResources",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference)SetResources(val *[]*string) {
	if err := j.validateSetResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resources",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference)SetSid(val *string) {
	if err := j.validateSetSidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sid",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) PutCondition(value interface{}) {
	if err := d.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCondition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) PutNotPrincipals(value interface{}) {
	if err := d.validatePutNotPrincipalsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNotPrincipals",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) PutPrincipals(value interface{}) {
	if err := d.validatePutPrincipalsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPrincipals",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) ResetActions() {
	_jsii_.InvokeVoid(
		d,
		"resetActions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		d,
		"resetCondition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) ResetEffect() {
	_jsii_.InvokeVoid(
		d,
		"resetEffect",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) ResetNotActions() {
	_jsii_.InvokeVoid(
		d,
		"resetNotActions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) ResetNotPrincipals() {
	_jsii_.InvokeVoid(
		d,
		"resetNotPrincipals",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) ResetNotResources() {
	_jsii_.InvokeVoid(
		d,
		"resetNotResources",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) ResetPrincipals() {
	_jsii_.InvokeVoid(
		d,
		"resetPrincipals",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		d,
		"resetResources",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) ResetSid() {
	_jsii_.InvokeVoid(
		d,
		"resetSid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

