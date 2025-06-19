// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2networkinsightsanalysis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ec2networkinsightsanalysis/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2NetworkInsightsAnalysisExplanationsOutputReference interface {
	cdktf.ComplexObject
	Acl() Ec2NetworkInsightsAnalysisExplanationsAclList
	AclRule() Ec2NetworkInsightsAnalysisExplanationsAclRuleList
	Address() *string
	Addresses() *[]*string
	AttachedTo() Ec2NetworkInsightsAnalysisExplanationsAttachedToList
	AvailabilityZones() *[]*string
	Cidrs() *[]*string
	ClassicLoadBalancerListener() Ec2NetworkInsightsAnalysisExplanationsClassicLoadBalancerListenerList
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
	Component() Ec2NetworkInsightsAnalysisExplanationsComponentList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomerGateway() Ec2NetworkInsightsAnalysisExplanationsCustomerGatewayList
	Destination() Ec2NetworkInsightsAnalysisExplanationsDestinationList
	DestinationVpc() Ec2NetworkInsightsAnalysisExplanationsDestinationVpcList
	Direction() *string
	ElasticLoadBalancerListener() Ec2NetworkInsightsAnalysisExplanationsElasticLoadBalancerListenerList
	ExplanationCode() *string
	// Experimental.
	Fqn() *string
	IngressRouteTable() Ec2NetworkInsightsAnalysisExplanationsIngressRouteTableList
	InternalValue() *Ec2NetworkInsightsAnalysisExplanations
	SetInternalValue(val *Ec2NetworkInsightsAnalysisExplanations)
	InternetGateway() Ec2NetworkInsightsAnalysisExplanationsInternetGatewayList
	LoadBalancerArn() *string
	LoadBalancerListenerPort() *float64
	LoadBalancerTargetGroup() Ec2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupList
	LoadBalancerTargetGroups() Ec2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupsList
	LoadBalancerTargetPort() *float64
	MissingComponent() *string
	NatGateway() Ec2NetworkInsightsAnalysisExplanationsNatGatewayList
	NetworkInterface() Ec2NetworkInsightsAnalysisExplanationsNetworkInterfaceList
	PacketField() *string
	Port() *float64
	PortRanges() Ec2NetworkInsightsAnalysisExplanationsPortRangesList
	PrefixList() Ec2NetworkInsightsAnalysisExplanationsPrefixListStructList
	Protocols() *[]*string
	RouteTable() Ec2NetworkInsightsAnalysisExplanationsRouteTableList
	RouteTableRoute() Ec2NetworkInsightsAnalysisExplanationsRouteTableRouteList
	SecurityGroup() Ec2NetworkInsightsAnalysisExplanationsSecurityGroupList
	SecurityGroupRule() Ec2NetworkInsightsAnalysisExplanationsSecurityGroupRuleList
	SecurityGroups() Ec2NetworkInsightsAnalysisExplanationsSecurityGroupsList
	SourceVpc() Ec2NetworkInsightsAnalysisExplanationsSourceVpcList
	State() *string
	Subnet() Ec2NetworkInsightsAnalysisExplanationsSubnetList
	SubnetRouteTable() Ec2NetworkInsightsAnalysisExplanationsSubnetRouteTableList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransitGateway() Ec2NetworkInsightsAnalysisExplanationsTransitGatewayList
	TransitGatewayAttachment() Ec2NetworkInsightsAnalysisExplanationsTransitGatewayAttachmentList
	TransitGatewayRouteTable() Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList
	TransitGatewayRouteTableRoute() Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableRouteList
	Vpc() Ec2NetworkInsightsAnalysisExplanationsVpcList
	VpcEndpoint() Ec2NetworkInsightsAnalysisExplanationsVpcEndpointList
	VpcPeeringConnection() Ec2NetworkInsightsAnalysisExplanationsVpcPeeringConnectionList
	VpnConnection() Ec2NetworkInsightsAnalysisExplanationsVpnConnectionList
	VpnGateway() Ec2NetworkInsightsAnalysisExplanationsVpnGatewayList
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

// The jsii proxy struct for Ec2NetworkInsightsAnalysisExplanationsOutputReference
type jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Acl() Ec2NetworkInsightsAnalysisExplanationsAclList {
	var returns Ec2NetworkInsightsAnalysisExplanationsAclList
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) AclRule() Ec2NetworkInsightsAnalysisExplanationsAclRuleList {
	var returns Ec2NetworkInsightsAnalysisExplanationsAclRuleList
	_jsii_.Get(
		j,
		"aclRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) AttachedTo() Ec2NetworkInsightsAnalysisExplanationsAttachedToList {
	var returns Ec2NetworkInsightsAnalysisExplanationsAttachedToList
	_jsii_.Get(
		j,
		"attachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Cidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) ClassicLoadBalancerListener() Ec2NetworkInsightsAnalysisExplanationsClassicLoadBalancerListenerList {
	var returns Ec2NetworkInsightsAnalysisExplanationsClassicLoadBalancerListenerList
	_jsii_.Get(
		j,
		"classicLoadBalancerListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Component() Ec2NetworkInsightsAnalysisExplanationsComponentList {
	var returns Ec2NetworkInsightsAnalysisExplanationsComponentList
	_jsii_.Get(
		j,
		"component",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) CustomerGateway() Ec2NetworkInsightsAnalysisExplanationsCustomerGatewayList {
	var returns Ec2NetworkInsightsAnalysisExplanationsCustomerGatewayList
	_jsii_.Get(
		j,
		"customerGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Destination() Ec2NetworkInsightsAnalysisExplanationsDestinationList {
	var returns Ec2NetworkInsightsAnalysisExplanationsDestinationList
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) DestinationVpc() Ec2NetworkInsightsAnalysisExplanationsDestinationVpcList {
	var returns Ec2NetworkInsightsAnalysisExplanationsDestinationVpcList
	_jsii_.Get(
		j,
		"destinationVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) ElasticLoadBalancerListener() Ec2NetworkInsightsAnalysisExplanationsElasticLoadBalancerListenerList {
	var returns Ec2NetworkInsightsAnalysisExplanationsElasticLoadBalancerListenerList
	_jsii_.Get(
		j,
		"elasticLoadBalancerListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) ExplanationCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"explanationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) IngressRouteTable() Ec2NetworkInsightsAnalysisExplanationsIngressRouteTableList {
	var returns Ec2NetworkInsightsAnalysisExplanationsIngressRouteTableList
	_jsii_.Get(
		j,
		"ingressRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) InternalValue() *Ec2NetworkInsightsAnalysisExplanations {
	var returns *Ec2NetworkInsightsAnalysisExplanations
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) InternetGateway() Ec2NetworkInsightsAnalysisExplanationsInternetGatewayList {
	var returns Ec2NetworkInsightsAnalysisExplanationsInternetGatewayList
	_jsii_.Get(
		j,
		"internetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerListenerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadBalancerListenerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerTargetGroup() Ec2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupList {
	var returns Ec2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupList
	_jsii_.Get(
		j,
		"loadBalancerTargetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerTargetGroups() Ec2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupsList {
	var returns Ec2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupsList
	_jsii_.Get(
		j,
		"loadBalancerTargetGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerTargetPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadBalancerTargetPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) MissingComponent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"missingComponent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) NatGateway() Ec2NetworkInsightsAnalysisExplanationsNatGatewayList {
	var returns Ec2NetworkInsightsAnalysisExplanationsNatGatewayList
	_jsii_.Get(
		j,
		"natGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) NetworkInterface() Ec2NetworkInsightsAnalysisExplanationsNetworkInterfaceList {
	var returns Ec2NetworkInsightsAnalysisExplanationsNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) PacketField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packetField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) PortRanges() Ec2NetworkInsightsAnalysisExplanationsPortRangesList {
	var returns Ec2NetworkInsightsAnalysisExplanationsPortRangesList
	_jsii_.Get(
		j,
		"portRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) PrefixList() Ec2NetworkInsightsAnalysisExplanationsPrefixListStructList {
	var returns Ec2NetworkInsightsAnalysisExplanationsPrefixListStructList
	_jsii_.Get(
		j,
		"prefixList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) RouteTable() Ec2NetworkInsightsAnalysisExplanationsRouteTableList {
	var returns Ec2NetworkInsightsAnalysisExplanationsRouteTableList
	_jsii_.Get(
		j,
		"routeTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) RouteTableRoute() Ec2NetworkInsightsAnalysisExplanationsRouteTableRouteList {
	var returns Ec2NetworkInsightsAnalysisExplanationsRouteTableRouteList
	_jsii_.Get(
		j,
		"routeTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) SecurityGroup() Ec2NetworkInsightsAnalysisExplanationsSecurityGroupList {
	var returns Ec2NetworkInsightsAnalysisExplanationsSecurityGroupList
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) SecurityGroupRule() Ec2NetworkInsightsAnalysisExplanationsSecurityGroupRuleList {
	var returns Ec2NetworkInsightsAnalysisExplanationsSecurityGroupRuleList
	_jsii_.Get(
		j,
		"securityGroupRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) SecurityGroups() Ec2NetworkInsightsAnalysisExplanationsSecurityGroupsList {
	var returns Ec2NetworkInsightsAnalysisExplanationsSecurityGroupsList
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) SourceVpc() Ec2NetworkInsightsAnalysisExplanationsSourceVpcList {
	var returns Ec2NetworkInsightsAnalysisExplanationsSourceVpcList
	_jsii_.Get(
		j,
		"sourceVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Subnet() Ec2NetworkInsightsAnalysisExplanationsSubnetList {
	var returns Ec2NetworkInsightsAnalysisExplanationsSubnetList
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) SubnetRouteTable() Ec2NetworkInsightsAnalysisExplanationsSubnetRouteTableList {
	var returns Ec2NetworkInsightsAnalysisExplanationsSubnetRouteTableList
	_jsii_.Get(
		j,
		"subnetRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) TransitGateway() Ec2NetworkInsightsAnalysisExplanationsTransitGatewayList {
	var returns Ec2NetworkInsightsAnalysisExplanationsTransitGatewayList
	_jsii_.Get(
		j,
		"transitGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) TransitGatewayAttachment() Ec2NetworkInsightsAnalysisExplanationsTransitGatewayAttachmentList {
	var returns Ec2NetworkInsightsAnalysisExplanationsTransitGatewayAttachmentList
	_jsii_.Get(
		j,
		"transitGatewayAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) TransitGatewayRouteTable() Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList {
	var returns Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList
	_jsii_.Get(
		j,
		"transitGatewayRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) TransitGatewayRouteTableRoute() Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableRouteList {
	var returns Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableRouteList
	_jsii_.Get(
		j,
		"transitGatewayRouteTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Vpc() Ec2NetworkInsightsAnalysisExplanationsVpcList {
	var returns Ec2NetworkInsightsAnalysisExplanationsVpcList
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) VpcEndpoint() Ec2NetworkInsightsAnalysisExplanationsVpcEndpointList {
	var returns Ec2NetworkInsightsAnalysisExplanationsVpcEndpointList
	_jsii_.Get(
		j,
		"vpcEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) VpcPeeringConnection() Ec2NetworkInsightsAnalysisExplanationsVpcPeeringConnectionList {
	var returns Ec2NetworkInsightsAnalysisExplanationsVpcPeeringConnectionList
	_jsii_.Get(
		j,
		"vpcPeeringConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) VpnConnection() Ec2NetworkInsightsAnalysisExplanationsVpnConnectionList {
	var returns Ec2NetworkInsightsAnalysisExplanationsVpnConnectionList
	_jsii_.Get(
		j,
		"vpnConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) VpnGateway() Ec2NetworkInsightsAnalysisExplanationsVpnGatewayList {
	var returns Ec2NetworkInsightsAnalysisExplanationsVpnGatewayList
	_jsii_.Get(
		j,
		"vpnGateway",
		&returns,
	)
	return returns
}


func NewEc2NetworkInsightsAnalysisExplanationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2NetworkInsightsAnalysisExplanationsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2NetworkInsightsAnalysisExplanationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2NetworkInsightsAnalysis.Ec2NetworkInsightsAnalysisExplanationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2NetworkInsightsAnalysisExplanationsOutputReference_Override(e Ec2NetworkInsightsAnalysisExplanationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2NetworkInsightsAnalysis.Ec2NetworkInsightsAnalysisExplanationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference)SetInternalValue(val *Ec2NetworkInsightsAnalysisExplanations) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

