// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference interface {
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
	DateValue() *string
	SetDateValue(val *string)
	DateValueInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValue
	SetInternalValue(val *KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValue)
	LongValue() *float64
	SetLongValue(val *float64)
	LongValueInput() *float64
	StringListValue() *[]*string
	SetStringListValue(val *[]*string)
	StringListValueInput() *[]*string
	StringValue() *string
	SetStringValue(val *string)
	StringValueInput() *string
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
	ResetDateValue()
	ResetLongValue()
	ResetStringListValue()
	ResetStringValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference
type jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) DateValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) DateValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) InternalValue() *KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValue {
	var returns *KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) LongValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) LongValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) StringListValue() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stringListValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) StringListValueInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stringListValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) StringValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference {
	_init_.Initialize()

	if err := validateNewKendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference_Override(k KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetDateValue(val *string) {
	if err := j.validateSetDateValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetInternalValue(val *KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValue) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetLongValue(val *float64) {
	if err := j.validateSetLongValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetStringListValue(val *[]*string) {
	if err := j.validateSetStringListValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringListValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetStringValue(val *string) {
	if err := j.validateSetStringValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) ResetDateValue() {
	_jsii_.InvokeVoid(
		k,
		"resetDateValue",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) ResetLongValue() {
	_jsii_.InvokeVoid(
		k,
		"resetLongValue",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) ResetStringListValue() {
	_jsii_.InvokeVoid(
		k,
		"resetStringListValue",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) ResetStringValue() {
	_jsii_.InvokeVoid(
		k,
		"resetStringValue",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

