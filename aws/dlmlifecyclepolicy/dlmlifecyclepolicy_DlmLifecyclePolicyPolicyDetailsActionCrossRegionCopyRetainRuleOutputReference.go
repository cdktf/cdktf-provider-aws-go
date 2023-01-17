package dlmlifecyclepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/dlmlifecyclepolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference interface {
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
	InternalValue() *DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRule
	SetInternalValue(val *DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRule)
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	IntervalUnit() *string
	SetIntervalUnit(val *string)
	IntervalUnitInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference
type jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) InternalValue() *DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRule {
	var returns *DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) IntervalUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) IntervalUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference {
	_init_.Initialize()

	if err := validateNewDlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dlmLifecyclePolicy.DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference_Override(d DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dlmLifecyclePolicy.DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference)SetInternalValue(val *DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference)SetInterval(val *float64) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference)SetIntervalUnit(val *string) {
	if err := j.validateSetIntervalUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalUnit",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

