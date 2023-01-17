package ec2networkinsightsanalysis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/ec2networkinsightsanalysis/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference interface {
	cdktf.ComplexObject
	AclRule() Ec2NetworkInsightsAnalysisReturnPathComponentsAclRuleList
	AdditionalDetails() Ec2NetworkInsightsAnalysisReturnPathComponentsAdditionalDetailsList
	AttachedTo() Ec2NetworkInsightsAnalysisReturnPathComponentsAttachedToList
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
	Component() Ec2NetworkInsightsAnalysisReturnPathComponentsComponentList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DestinationVpc() Ec2NetworkInsightsAnalysisReturnPathComponentsDestinationVpcList
	// Experimental.
	Fqn() *string
	InboundHeader() Ec2NetworkInsightsAnalysisReturnPathComponentsInboundHeaderList
	InternalValue() *Ec2NetworkInsightsAnalysisReturnPathComponents
	SetInternalValue(val *Ec2NetworkInsightsAnalysisReturnPathComponents)
	OutboundHeader() Ec2NetworkInsightsAnalysisReturnPathComponentsOutboundHeaderList
	RouteTableRoute() Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteList
	SecurityGroupRule() Ec2NetworkInsightsAnalysisReturnPathComponentsSecurityGroupRuleList
	SequenceNumber() *float64
	SourceVpc() Ec2NetworkInsightsAnalysisReturnPathComponentsSourceVpcList
	Subnet() Ec2NetworkInsightsAnalysisReturnPathComponentsSubnetList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransitGateway() Ec2NetworkInsightsAnalysisReturnPathComponentsTransitGatewayList
	TransitGatewayRouteTableRoute() Ec2NetworkInsightsAnalysisReturnPathComponentsTransitGatewayRouteTableRouteList
	Vpc() Ec2NetworkInsightsAnalysisReturnPathComponentsVpcList
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

// The jsii proxy struct for Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference
type jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) AclRule() Ec2NetworkInsightsAnalysisReturnPathComponentsAclRuleList {
	var returns Ec2NetworkInsightsAnalysisReturnPathComponentsAclRuleList
	_jsii_.Get(
		j,
		"aclRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) AdditionalDetails() Ec2NetworkInsightsAnalysisReturnPathComponentsAdditionalDetailsList {
	var returns Ec2NetworkInsightsAnalysisReturnPathComponentsAdditionalDetailsList
	_jsii_.Get(
		j,
		"additionalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) AttachedTo() Ec2NetworkInsightsAnalysisReturnPathComponentsAttachedToList {
	var returns Ec2NetworkInsightsAnalysisReturnPathComponentsAttachedToList
	_jsii_.Get(
		j,
		"attachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) Component() Ec2NetworkInsightsAnalysisReturnPathComponentsComponentList {
	var returns Ec2NetworkInsightsAnalysisReturnPathComponentsComponentList
	_jsii_.Get(
		j,
		"component",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) DestinationVpc() Ec2NetworkInsightsAnalysisReturnPathComponentsDestinationVpcList {
	var returns Ec2NetworkInsightsAnalysisReturnPathComponentsDestinationVpcList
	_jsii_.Get(
		j,
		"destinationVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) InboundHeader() Ec2NetworkInsightsAnalysisReturnPathComponentsInboundHeaderList {
	var returns Ec2NetworkInsightsAnalysisReturnPathComponentsInboundHeaderList
	_jsii_.Get(
		j,
		"inboundHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) InternalValue() *Ec2NetworkInsightsAnalysisReturnPathComponents {
	var returns *Ec2NetworkInsightsAnalysisReturnPathComponents
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) OutboundHeader() Ec2NetworkInsightsAnalysisReturnPathComponentsOutboundHeaderList {
	var returns Ec2NetworkInsightsAnalysisReturnPathComponentsOutboundHeaderList
	_jsii_.Get(
		j,
		"outboundHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) RouteTableRoute() Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteList {
	var returns Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteList
	_jsii_.Get(
		j,
		"routeTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) SecurityGroupRule() Ec2NetworkInsightsAnalysisReturnPathComponentsSecurityGroupRuleList {
	var returns Ec2NetworkInsightsAnalysisReturnPathComponentsSecurityGroupRuleList
	_jsii_.Get(
		j,
		"securityGroupRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) SequenceNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sequenceNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) SourceVpc() Ec2NetworkInsightsAnalysisReturnPathComponentsSourceVpcList {
	var returns Ec2NetworkInsightsAnalysisReturnPathComponentsSourceVpcList
	_jsii_.Get(
		j,
		"sourceVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) Subnet() Ec2NetworkInsightsAnalysisReturnPathComponentsSubnetList {
	var returns Ec2NetworkInsightsAnalysisReturnPathComponentsSubnetList
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) TransitGateway() Ec2NetworkInsightsAnalysisReturnPathComponentsTransitGatewayList {
	var returns Ec2NetworkInsightsAnalysisReturnPathComponentsTransitGatewayList
	_jsii_.Get(
		j,
		"transitGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) TransitGatewayRouteTableRoute() Ec2NetworkInsightsAnalysisReturnPathComponentsTransitGatewayRouteTableRouteList {
	var returns Ec2NetworkInsightsAnalysisReturnPathComponentsTransitGatewayRouteTableRouteList
	_jsii_.Get(
		j,
		"transitGatewayRouteTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) Vpc() Ec2NetworkInsightsAnalysisReturnPathComponentsVpcList {
	var returns Ec2NetworkInsightsAnalysisReturnPathComponentsVpcList
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2NetworkInsightsAnalysisReturnPathComponentsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2NetworkInsightsAnalysis.Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference_Override(e Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2NetworkInsightsAnalysis.Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference)SetInternalValue(val *Ec2NetworkInsightsAnalysisReturnPathComponents) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

