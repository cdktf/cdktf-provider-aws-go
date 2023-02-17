package wafv2webaclloggingconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/wafv2webaclloggingconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference interface {
	cdktf.ComplexObject
	ActionCondition() Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionActionConditionOutputReference
	ActionConditionInput() *Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionActionCondition
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
	LabelNameCondition() Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameConditionOutputReference
	LabelNameConditionInput() *Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameCondition
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
	PutActionCondition(value *Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionActionCondition)
	PutLabelNameCondition(value *Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameCondition)
	ResetActionCondition()
	ResetLabelNameCondition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference
type jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) ActionCondition() Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionActionConditionOutputReference {
	var returns Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionActionConditionOutputReference
	_jsii_.Get(
		j,
		"actionCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) ActionConditionInput() *Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionActionCondition {
	var returns *Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionActionCondition
	_jsii_.Get(
		j,
		"actionConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) LabelNameCondition() Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameConditionOutputReference {
	var returns Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameConditionOutputReference
	_jsii_.Get(
		j,
		"labelNameCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) LabelNameConditionInput() *Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameCondition {
	var returns *Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameCondition
	_jsii_.Get(
		j,
		"labelNameConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAclLoggingConfiguration.Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference_Override(w Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAclLoggingConfiguration.Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) PutActionCondition(value *Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionActionCondition) {
	if err := w.validatePutActionConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putActionCondition",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) PutLabelNameCondition(value *Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameCondition) {
	if err := w.validatePutLabelNameConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putLabelNameCondition",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) ResetActionCondition() {
	_jsii_.InvokeVoid(
		w,
		"resetActionCondition",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) ResetLabelNameCondition() {
	_jsii_.InvokeVoid(
		w,
		"resetLabelNameCondition",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclLoggingConfigurationLoggingFilterFilterConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

