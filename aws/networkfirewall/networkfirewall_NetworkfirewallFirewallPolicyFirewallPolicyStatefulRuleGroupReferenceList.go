package networkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/networkfirewall/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList
type jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewNetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList {
	_init_.Initialize()

	if err := validateNewNetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList{}

	_jsii_.Create(
		"@cdktf/provider-aws.networkfirewall.NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewNetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList_Override(n NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.networkfirewall.NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		n,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList) Get(index *float64) NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceOutputReference {
	if err := n.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceOutputReference

	_jsii_.Invoke(
		n,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
