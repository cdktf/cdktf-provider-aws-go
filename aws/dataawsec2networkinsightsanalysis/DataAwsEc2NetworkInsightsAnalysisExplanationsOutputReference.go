package dataawsec2networkinsightsanalysis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/dataawsec2networkinsightsanalysis/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference interface {
	cdktf.ComplexObject
	Acl() DataAwsEc2NetworkInsightsAnalysisExplanationsAclList
	AclRule() DataAwsEc2NetworkInsightsAnalysisExplanationsAclRuleList
	Address() *string
	Addresses() *[]*string
	AttachedTo() DataAwsEc2NetworkInsightsAnalysisExplanationsAttachedToList
	AvailabilityZones() *[]*string
	Cidrs() *[]*string
	ClassicLoadBalancerListener() DataAwsEc2NetworkInsightsAnalysisExplanationsClassicLoadBalancerListenerList
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
	Component() DataAwsEc2NetworkInsightsAnalysisExplanationsComponentList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomerGateway() DataAwsEc2NetworkInsightsAnalysisExplanationsCustomerGatewayList
	Destination() DataAwsEc2NetworkInsightsAnalysisExplanationsDestinationList
	DestinationVpc() DataAwsEc2NetworkInsightsAnalysisExplanationsDestinationVpcList
	Direction() *string
	ElasticLoadBalancerListener() DataAwsEc2NetworkInsightsAnalysisExplanationsElasticLoadBalancerListenerList
	ExplanationCode() *string
	// Experimental.
	Fqn() *string
	IngressRouteTable() DataAwsEc2NetworkInsightsAnalysisExplanationsIngressRouteTableList
	InternalValue() *DataAwsEc2NetworkInsightsAnalysisExplanations
	SetInternalValue(val *DataAwsEc2NetworkInsightsAnalysisExplanations)
	InternetGateway() DataAwsEc2NetworkInsightsAnalysisExplanationsInternetGatewayList
	LoadBalancerArn() *string
	LoadBalancerListenerPort() *float64
	LoadBalancerTargetGroup() DataAwsEc2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupList
	LoadBalancerTargetGroups() DataAwsEc2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupsList
	LoadBalancerTargetPort() *float64
	MissingComponent() *string
	NatGateway() DataAwsEc2NetworkInsightsAnalysisExplanationsNatGatewayList
	NetworkInterface() DataAwsEc2NetworkInsightsAnalysisExplanationsNetworkInterfaceList
	PacketField() *string
	Port() *float64
	PortRanges() DataAwsEc2NetworkInsightsAnalysisExplanationsPortRangesList
	PrefixList() DataAwsEc2NetworkInsightsAnalysisExplanationsPrefixListList
	Protocols() *[]*string
	RouteTable() DataAwsEc2NetworkInsightsAnalysisExplanationsRouteTableList
	RouteTableRoute() DataAwsEc2NetworkInsightsAnalysisExplanationsRouteTableRouteList
	SecurityGroup() DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupList
	SecurityGroupRule() DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRuleList
	SecurityGroups() DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupsList
	SourceVpc() DataAwsEc2NetworkInsightsAnalysisExplanationsSourceVpcList
	State() *string
	Subnet() DataAwsEc2NetworkInsightsAnalysisExplanationsSubnetList
	SubnetRouteTable() DataAwsEc2NetworkInsightsAnalysisExplanationsSubnetRouteTableList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransitGateway() DataAwsEc2NetworkInsightsAnalysisExplanationsTransitGatewayList
	TransitGatewayAttachment() DataAwsEc2NetworkInsightsAnalysisExplanationsTransitGatewayAttachmentList
	TransitGatewayRouteTable() DataAwsEc2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList
	TransitGatewayRouteTableRoute() DataAwsEc2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableRouteList
	Vpc() DataAwsEc2NetworkInsightsAnalysisExplanationsVpcList
	VpcEndpoint() DataAwsEc2NetworkInsightsAnalysisExplanationsVpcEndpointList
	VpcPeeringConnection() DataAwsEc2NetworkInsightsAnalysisExplanationsVpcPeeringConnectionList
	VpnConnection() DataAwsEc2NetworkInsightsAnalysisExplanationsVpnConnectionList
	VpnGateway() DataAwsEc2NetworkInsightsAnalysisExplanationsVpnGatewayList
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

// The jsii proxy struct for DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference
type jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) Acl() DataAwsEc2NetworkInsightsAnalysisExplanationsAclList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsAclList
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) AclRule() DataAwsEc2NetworkInsightsAnalysisExplanationsAclRuleList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsAclRuleList
	_jsii_.Get(
		j,
		"aclRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) AttachedTo() DataAwsEc2NetworkInsightsAnalysisExplanationsAttachedToList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsAttachedToList
	_jsii_.Get(
		j,
		"attachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) Cidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) ClassicLoadBalancerListener() DataAwsEc2NetworkInsightsAnalysisExplanationsClassicLoadBalancerListenerList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsClassicLoadBalancerListenerList
	_jsii_.Get(
		j,
		"classicLoadBalancerListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) Component() DataAwsEc2NetworkInsightsAnalysisExplanationsComponentList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsComponentList
	_jsii_.Get(
		j,
		"component",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) CustomerGateway() DataAwsEc2NetworkInsightsAnalysisExplanationsCustomerGatewayList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsCustomerGatewayList
	_jsii_.Get(
		j,
		"customerGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) Destination() DataAwsEc2NetworkInsightsAnalysisExplanationsDestinationList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsDestinationList
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) DestinationVpc() DataAwsEc2NetworkInsightsAnalysisExplanationsDestinationVpcList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsDestinationVpcList
	_jsii_.Get(
		j,
		"destinationVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) ElasticLoadBalancerListener() DataAwsEc2NetworkInsightsAnalysisExplanationsElasticLoadBalancerListenerList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsElasticLoadBalancerListenerList
	_jsii_.Get(
		j,
		"elasticLoadBalancerListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) ExplanationCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"explanationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) IngressRouteTable() DataAwsEc2NetworkInsightsAnalysisExplanationsIngressRouteTableList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsIngressRouteTableList
	_jsii_.Get(
		j,
		"ingressRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) InternalValue() *DataAwsEc2NetworkInsightsAnalysisExplanations {
	var returns *DataAwsEc2NetworkInsightsAnalysisExplanations
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) InternetGateway() DataAwsEc2NetworkInsightsAnalysisExplanationsInternetGatewayList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsInternetGatewayList
	_jsii_.Get(
		j,
		"internetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerListenerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadBalancerListenerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerTargetGroup() DataAwsEc2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupList
	_jsii_.Get(
		j,
		"loadBalancerTargetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerTargetGroups() DataAwsEc2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupsList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsLoadBalancerTargetGroupsList
	_jsii_.Get(
		j,
		"loadBalancerTargetGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) LoadBalancerTargetPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadBalancerTargetPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) MissingComponent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"missingComponent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) NatGateway() DataAwsEc2NetworkInsightsAnalysisExplanationsNatGatewayList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsNatGatewayList
	_jsii_.Get(
		j,
		"natGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) NetworkInterface() DataAwsEc2NetworkInsightsAnalysisExplanationsNetworkInterfaceList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) PacketField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packetField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) PortRanges() DataAwsEc2NetworkInsightsAnalysisExplanationsPortRangesList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsPortRangesList
	_jsii_.Get(
		j,
		"portRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) PrefixList() DataAwsEc2NetworkInsightsAnalysisExplanationsPrefixListList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsPrefixListList
	_jsii_.Get(
		j,
		"prefixList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) RouteTable() DataAwsEc2NetworkInsightsAnalysisExplanationsRouteTableList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsRouteTableList
	_jsii_.Get(
		j,
		"routeTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) RouteTableRoute() DataAwsEc2NetworkInsightsAnalysisExplanationsRouteTableRouteList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsRouteTableRouteList
	_jsii_.Get(
		j,
		"routeTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) SecurityGroup() DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupList
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) SecurityGroupRule() DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRuleList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRuleList
	_jsii_.Get(
		j,
		"securityGroupRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) SecurityGroups() DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupsList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupsList
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) SourceVpc() DataAwsEc2NetworkInsightsAnalysisExplanationsSourceVpcList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsSourceVpcList
	_jsii_.Get(
		j,
		"sourceVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) Subnet() DataAwsEc2NetworkInsightsAnalysisExplanationsSubnetList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsSubnetList
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) SubnetRouteTable() DataAwsEc2NetworkInsightsAnalysisExplanationsSubnetRouteTableList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsSubnetRouteTableList
	_jsii_.Get(
		j,
		"subnetRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) TransitGateway() DataAwsEc2NetworkInsightsAnalysisExplanationsTransitGatewayList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsTransitGatewayList
	_jsii_.Get(
		j,
		"transitGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) TransitGatewayAttachment() DataAwsEc2NetworkInsightsAnalysisExplanationsTransitGatewayAttachmentList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsTransitGatewayAttachmentList
	_jsii_.Get(
		j,
		"transitGatewayAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) TransitGatewayRouteTable() DataAwsEc2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList
	_jsii_.Get(
		j,
		"transitGatewayRouteTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) TransitGatewayRouteTableRoute() DataAwsEc2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableRouteList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableRouteList
	_jsii_.Get(
		j,
		"transitGatewayRouteTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) Vpc() DataAwsEc2NetworkInsightsAnalysisExplanationsVpcList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsVpcList
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) VpcEndpoint() DataAwsEc2NetworkInsightsAnalysisExplanationsVpcEndpointList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsVpcEndpointList
	_jsii_.Get(
		j,
		"vpcEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) VpcPeeringConnection() DataAwsEc2NetworkInsightsAnalysisExplanationsVpcPeeringConnectionList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsVpcPeeringConnectionList
	_jsii_.Get(
		j,
		"vpcPeeringConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) VpnConnection() DataAwsEc2NetworkInsightsAnalysisExplanationsVpnConnectionList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsVpnConnectionList
	_jsii_.Get(
		j,
		"vpnConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) VpnGateway() DataAwsEc2NetworkInsightsAnalysisExplanationsVpnGatewayList {
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsVpnGatewayList
	_jsii_.Get(
		j,
		"vpnGateway",
		&returns,
	)
	return returns
}


func NewDataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsEc2NetworkInsightsAnalysisExplanationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsEc2NetworkInsightsAnalysis.DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference_Override(d DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsEc2NetworkInsightsAnalysis.DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference)SetInternalValue(val *DataAwsEc2NetworkInsightsAnalysisExplanations) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

