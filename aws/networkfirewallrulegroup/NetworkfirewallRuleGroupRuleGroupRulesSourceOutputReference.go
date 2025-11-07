// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/networkfirewallrulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference interface {
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
	InternalValue() *NetworkfirewallRuleGroupRuleGroupRulesSource
	SetInternalValue(val *NetworkfirewallRuleGroupRuleGroupRulesSource)
	RulesSourceList() NetworkfirewallRuleGroupRuleGroupRulesSourceRulesSourceListStructOutputReference
	RulesSourceListInput() *NetworkfirewallRuleGroupRuleGroupRulesSourceRulesSourceListStruct
	RulesString() *string
	SetRulesString(val *string)
	RulesStringInput() *string
	StatefulRule() NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRuleList
	StatefulRuleInput() interface{}
	StatelessRulesAndCustomActions() NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsOutputReference
	StatelessRulesAndCustomActionsInput() *NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActions
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutRulesSourceList(value *NetworkfirewallRuleGroupRuleGroupRulesSourceRulesSourceListStruct)
	PutStatefulRule(value interface{})
	PutStatelessRulesAndCustomActions(value *NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActions)
	ResetRulesSourceList()
	ResetRulesString()
	ResetStatefulRule()
	ResetStatelessRulesAndCustomActions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference
type jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) InternalValue() *NetworkfirewallRuleGroupRuleGroupRulesSource {
	var returns *NetworkfirewallRuleGroupRuleGroupRulesSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) RulesSourceList() NetworkfirewallRuleGroupRuleGroupRulesSourceRulesSourceListStructOutputReference {
	var returns NetworkfirewallRuleGroupRuleGroupRulesSourceRulesSourceListStructOutputReference
	_jsii_.Get(
		j,
		"rulesSourceList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) RulesSourceListInput() *NetworkfirewallRuleGroupRuleGroupRulesSourceRulesSourceListStruct {
	var returns *NetworkfirewallRuleGroupRuleGroupRulesSourceRulesSourceListStruct
	_jsii_.Get(
		j,
		"rulesSourceListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) RulesString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rulesString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) RulesStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rulesStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) StatefulRule() NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRuleList {
	var returns NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRuleList
	_jsii_.Get(
		j,
		"statefulRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) StatefulRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) StatelessRulesAndCustomActions() NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsOutputReference {
	var returns NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsOutputReference
	_jsii_.Get(
		j,
		"statelessRulesAndCustomActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) StatelessRulesAndCustomActionsInput() *NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActions {
	var returns *NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActions
	_jsii_.Get(
		j,
		"statelessRulesAndCustomActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkfirewallRuleGroupRuleGroupRulesSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.networkfirewallRuleGroup.NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference_Override(n NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.networkfirewallRuleGroup.NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference)SetInternalValue(val *NetworkfirewallRuleGroupRuleGroupRulesSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference)SetRulesString(val *string) {
	if err := j.validateSetRulesStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rulesString",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) PutRulesSourceList(value *NetworkfirewallRuleGroupRuleGroupRulesSourceRulesSourceListStruct) {
	if err := n.validatePutRulesSourceListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putRulesSourceList",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) PutStatefulRule(value interface{}) {
	if err := n.validatePutStatefulRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putStatefulRule",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) PutStatelessRulesAndCustomActions(value *NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActions) {
	if err := n.validatePutStatelessRulesAndCustomActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putStatelessRulesAndCustomActions",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) ResetRulesSourceList() {
	_jsii_.InvokeVoid(
		n,
		"resetRulesSourceList",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) ResetRulesString() {
	_jsii_.InvokeVoid(
		n,
		"resetRulesString",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) ResetStatefulRule() {
	_jsii_.InvokeVoid(
		n,
		"resetStatefulRule",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) ResetStatelessRulesAndCustomActions() {
	_jsii_.InvokeVoid(
		n,
		"resetStatelessRulesAndCustomActions",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallRuleGroupRuleGroupRulesSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

