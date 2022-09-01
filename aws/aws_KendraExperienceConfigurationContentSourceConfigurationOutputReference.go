// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KendraExperienceConfigurationContentSourceConfigurationOutputReference interface {
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
	DataSourceIds() *[]*string
	SetDataSourceIds(val *[]*string)
	DataSourceIdsInput() *[]*string
	DirectPutContent() interface{}
	SetDirectPutContent(val interface{})
	DirectPutContentInput() interface{}
	FaqIds() *[]*string
	SetFaqIds(val *[]*string)
	FaqIdsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KendraExperienceConfigurationContentSourceConfiguration
	SetInternalValue(val *KendraExperienceConfigurationContentSourceConfiguration)
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
	ResetDataSourceIds()
	ResetDirectPutContent()
	ResetFaqIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraExperienceConfigurationContentSourceConfigurationOutputReference
type jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) DataSourceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataSourceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) DataSourceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataSourceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) DirectPutContent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directPutContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) DirectPutContentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directPutContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) FaqIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"faqIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) FaqIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"faqIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) InternalValue() *KendraExperienceConfigurationContentSourceConfiguration {
	var returns *KendraExperienceConfigurationContentSourceConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKendraExperienceConfigurationContentSourceConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KendraExperienceConfigurationContentSourceConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKendraExperienceConfigurationContentSourceConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.KendraExperienceConfigurationContentSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraExperienceConfigurationContentSourceConfigurationOutputReference_Override(k KendraExperienceConfigurationContentSourceConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.KendraExperienceConfigurationContentSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference)SetDataSourceIds(val *[]*string) {
	if err := j.validateSetDataSourceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceIds",
		val,
	)
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference)SetDirectPutContent(val interface{}) {
	if err := j.validateSetDirectPutContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directPutContent",
		val,
	)
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference)SetFaqIds(val *[]*string) {
	if err := j.validateSetFaqIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"faqIds",
		val,
	)
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference)SetInternalValue(val *KendraExperienceConfigurationContentSourceConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) ResetDataSourceIds() {
	_jsii_.InvokeVoid(
		k,
		"resetDataSourceIds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) ResetDirectPutContent() {
	_jsii_.InvokeVoid(
		k,
		"resetDirectPutContent",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) ResetFaqIds() {
	_jsii_.InvokeVoid(
		k,
		"resetFaqIds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KendraExperienceConfigurationContentSourceConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

