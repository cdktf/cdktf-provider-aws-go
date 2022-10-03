package networkfirewallfirewallpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/networkfirewallfirewallpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkfirewallFirewallPolicyFirewallPolicyOutputReference interface {
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
	InternalValue() *NetworkfirewallFirewallPolicyFirewallPolicy
	SetInternalValue(val *NetworkfirewallFirewallPolicyFirewallPolicy)
	StatefulDefaultActions() *[]*string
	SetStatefulDefaultActions(val *[]*string)
	StatefulDefaultActionsInput() *[]*string
	StatefulEngineOptions() NetworkfirewallFirewallPolicyFirewallPolicyStatefulEngineOptionsOutputReference
	StatefulEngineOptionsInput() *NetworkfirewallFirewallPolicyFirewallPolicyStatefulEngineOptions
	StatefulRuleGroupReference() NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList
	StatefulRuleGroupReferenceInput() interface{}
	StatelessCustomAction() NetworkfirewallFirewallPolicyFirewallPolicyStatelessCustomActionList
	StatelessCustomActionInput() interface{}
	StatelessDefaultActions() *[]*string
	SetStatelessDefaultActions(val *[]*string)
	StatelessDefaultActionsInput() *[]*string
	StatelessFragmentDefaultActions() *[]*string
	SetStatelessFragmentDefaultActions(val *[]*string)
	StatelessFragmentDefaultActionsInput() *[]*string
	StatelessRuleGroupReference() NetworkfirewallFirewallPolicyFirewallPolicyStatelessRuleGroupReferenceList
	StatelessRuleGroupReferenceInput() interface{}
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
	PutStatefulEngineOptions(value *NetworkfirewallFirewallPolicyFirewallPolicyStatefulEngineOptions)
	PutStatefulRuleGroupReference(value interface{})
	PutStatelessCustomAction(value interface{})
	PutStatelessRuleGroupReference(value interface{})
	ResetStatefulDefaultActions()
	ResetStatefulEngineOptions()
	ResetStatefulRuleGroupReference()
	ResetStatelessCustomAction()
	ResetStatelessRuleGroupReference()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkfirewallFirewallPolicyFirewallPolicyOutputReference
type jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) InternalValue() *NetworkfirewallFirewallPolicyFirewallPolicy {
	var returns *NetworkfirewallFirewallPolicyFirewallPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatefulDefaultActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statefulDefaultActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatefulDefaultActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statefulDefaultActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatefulEngineOptions() NetworkfirewallFirewallPolicyFirewallPolicyStatefulEngineOptionsOutputReference {
	var returns NetworkfirewallFirewallPolicyFirewallPolicyStatefulEngineOptionsOutputReference
	_jsii_.Get(
		j,
		"statefulEngineOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatefulEngineOptionsInput() *NetworkfirewallFirewallPolicyFirewallPolicyStatefulEngineOptions {
	var returns *NetworkfirewallFirewallPolicyFirewallPolicyStatefulEngineOptions
	_jsii_.Get(
		j,
		"statefulEngineOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatefulRuleGroupReference() NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList {
	var returns NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList
	_jsii_.Get(
		j,
		"statefulRuleGroupReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatefulRuleGroupReferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulRuleGroupReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatelessCustomAction() NetworkfirewallFirewallPolicyFirewallPolicyStatelessCustomActionList {
	var returns NetworkfirewallFirewallPolicyFirewallPolicyStatelessCustomActionList
	_jsii_.Get(
		j,
		"statelessCustomAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatelessCustomActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statelessCustomActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatelessDefaultActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statelessDefaultActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatelessDefaultActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statelessDefaultActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatelessFragmentDefaultActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statelessFragmentDefaultActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatelessFragmentDefaultActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statelessFragmentDefaultActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatelessRuleGroupReference() NetworkfirewallFirewallPolicyFirewallPolicyStatelessRuleGroupReferenceList {
	var returns NetworkfirewallFirewallPolicyFirewallPolicyStatelessRuleGroupReferenceList
	_jsii_.Get(
		j,
		"statelessRuleGroupReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) StatelessRuleGroupReferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statelessRuleGroupReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkfirewallFirewallPolicyFirewallPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkfirewallFirewallPolicyFirewallPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkfirewallFirewallPolicyFirewallPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.networkfirewallFirewallPolicy.NetworkfirewallFirewallPolicyFirewallPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkfirewallFirewallPolicyFirewallPolicyOutputReference_Override(n NetworkfirewallFirewallPolicyFirewallPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.networkfirewallFirewallPolicy.NetworkfirewallFirewallPolicyFirewallPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference)SetInternalValue(val *NetworkfirewallFirewallPolicyFirewallPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference)SetStatefulDefaultActions(val *[]*string) {
	if err := j.validateSetStatefulDefaultActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statefulDefaultActions",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference)SetStatelessDefaultActions(val *[]*string) {
	if err := j.validateSetStatelessDefaultActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statelessDefaultActions",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference)SetStatelessFragmentDefaultActions(val *[]*string) {
	if err := j.validateSetStatelessFragmentDefaultActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statelessFragmentDefaultActions",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) PutStatefulEngineOptions(value *NetworkfirewallFirewallPolicyFirewallPolicyStatefulEngineOptions) {
	if err := n.validatePutStatefulEngineOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putStatefulEngineOptions",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) PutStatefulRuleGroupReference(value interface{}) {
	if err := n.validatePutStatefulRuleGroupReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putStatefulRuleGroupReference",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) PutStatelessCustomAction(value interface{}) {
	if err := n.validatePutStatelessCustomActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putStatelessCustomAction",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) PutStatelessRuleGroupReference(value interface{}) {
	if err := n.validatePutStatelessRuleGroupReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putStatelessRuleGroupReference",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) ResetStatefulDefaultActions() {
	_jsii_.InvokeVoid(
		n,
		"resetStatefulDefaultActions",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) ResetStatefulEngineOptions() {
	_jsii_.InvokeVoid(
		n,
		"resetStatefulEngineOptions",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) ResetStatefulRuleGroupReference() {
	_jsii_.InvokeVoid(
		n,
		"resetStatefulRuleGroupReference",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) ResetStatelessCustomAction() {
	_jsii_.InvokeVoid(
		n,
		"resetStatelessCustomAction",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) ResetStatelessRuleGroupReference() {
	_jsii_.InvokeVoid(
		n,
		"resetStatelessRuleGroupReference",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

