package ec2networkinsightsanalysis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/ec2networkinsightsanalysis/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference interface {
	cdktf.ComplexObject
	AclRule() Ec2NetworkInsightsAnalysisForwardPathComponentsAclRuleList
	AdditionalDetails() Ec2NetworkInsightsAnalysisForwardPathComponentsAdditionalDetailsList
	AttachedTo() Ec2NetworkInsightsAnalysisForwardPathComponentsAttachedToList
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
	Component() Ec2NetworkInsightsAnalysisForwardPathComponentsComponentList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DestinationVpc() Ec2NetworkInsightsAnalysisForwardPathComponentsDestinationVpcList
	// Experimental.
	Fqn() *string
	InboundHeader() Ec2NetworkInsightsAnalysisForwardPathComponentsInboundHeaderList
	InternalValue() *Ec2NetworkInsightsAnalysisForwardPathComponents
	SetInternalValue(val *Ec2NetworkInsightsAnalysisForwardPathComponents)
	OutboundHeader() Ec2NetworkInsightsAnalysisForwardPathComponentsOutboundHeaderList
	RouteTableRoute() Ec2NetworkInsightsAnalysisForwardPathComponentsRouteTableRouteList
	SecurityGroupRule() Ec2NetworkInsightsAnalysisForwardPathComponentsSecurityGroupRuleList
	SequenceNumber() *float64
	SourceVpc() Ec2NetworkInsightsAnalysisForwardPathComponentsSourceVpcList
	Subnet() Ec2NetworkInsightsAnalysisForwardPathComponentsSubnetList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransitGateway() Ec2NetworkInsightsAnalysisForwardPathComponentsTransitGatewayList
	TransitGatewayRouteTableRoute() Ec2NetworkInsightsAnalysisForwardPathComponentsTransitGatewayRouteTableRouteList
	Vpc() Ec2NetworkInsightsAnalysisForwardPathComponentsVpcList
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

// The jsii proxy struct for Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference
type jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) AclRule() Ec2NetworkInsightsAnalysisForwardPathComponentsAclRuleList {
	var returns Ec2NetworkInsightsAnalysisForwardPathComponentsAclRuleList
	_jsii_.Get(
		j,
		"aclRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) AdditionalDetails() Ec2NetworkInsightsAnalysisForwardPathComponentsAdditionalDetailsList {
	var returns Ec2NetworkInsightsAnalysisForwardPathComponentsAdditionalDetailsList
	_jsii_.Get(
		j,
		"additionalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) AttachedTo() Ec2NetworkInsightsAnalysisForwardPathComponentsAttachedToList {
	var returns Ec2NetworkInsightsAnalysisForwardPathComponentsAttachedToList
	_jsii_.Get(
		j,
		"attachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) Component() Ec2NetworkInsightsAnalysisForwardPathComponentsComponentList {
	var returns Ec2NetworkInsightsAnalysisForwardPathComponentsComponentList
	_jsii_.Get(
		j,
		"component",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) DestinationVpc() Ec2NetworkInsightsAnalysisForwardPathComponentsDestinationVpcList {
	var returns Ec2NetworkInsightsAnalysisForwardPathComponentsDestinationVpcList
	_jsii_.Get(
		j,
		"destinationVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) InboundHeader() Ec2NetworkInsightsAnalysisForwardPathComponentsInboundHeaderList {
	var returns Ec2NetworkInsightsAnalysisForwardPathComponentsInboundHeaderList
	_jsii_.Get(
		j,
		"inboundHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) InternalValue() *Ec2NetworkInsightsAnalysisForwardPathComponents {
	var returns *Ec2NetworkInsightsAnalysisForwardPathComponents
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) OutboundHeader() Ec2NetworkInsightsAnalysisForwardPathComponentsOutboundHeaderList {
	var returns Ec2NetworkInsightsAnalysisForwardPathComponentsOutboundHeaderList
	_jsii_.Get(
		j,
		"outboundHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) RouteTableRoute() Ec2NetworkInsightsAnalysisForwardPathComponentsRouteTableRouteList {
	var returns Ec2NetworkInsightsAnalysisForwardPathComponentsRouteTableRouteList
	_jsii_.Get(
		j,
		"routeTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) SecurityGroupRule() Ec2NetworkInsightsAnalysisForwardPathComponentsSecurityGroupRuleList {
	var returns Ec2NetworkInsightsAnalysisForwardPathComponentsSecurityGroupRuleList
	_jsii_.Get(
		j,
		"securityGroupRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) SequenceNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sequenceNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) SourceVpc() Ec2NetworkInsightsAnalysisForwardPathComponentsSourceVpcList {
	var returns Ec2NetworkInsightsAnalysisForwardPathComponentsSourceVpcList
	_jsii_.Get(
		j,
		"sourceVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) Subnet() Ec2NetworkInsightsAnalysisForwardPathComponentsSubnetList {
	var returns Ec2NetworkInsightsAnalysisForwardPathComponentsSubnetList
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) TransitGateway() Ec2NetworkInsightsAnalysisForwardPathComponentsTransitGatewayList {
	var returns Ec2NetworkInsightsAnalysisForwardPathComponentsTransitGatewayList
	_jsii_.Get(
		j,
		"transitGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) TransitGatewayRouteTableRoute() Ec2NetworkInsightsAnalysisForwardPathComponentsTransitGatewayRouteTableRouteList {
	var returns Ec2NetworkInsightsAnalysisForwardPathComponentsTransitGatewayRouteTableRouteList
	_jsii_.Get(
		j,
		"transitGatewayRouteTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) Vpc() Ec2NetworkInsightsAnalysisForwardPathComponentsVpcList {
	var returns Ec2NetworkInsightsAnalysisForwardPathComponentsVpcList
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2NetworkInsightsAnalysisForwardPathComponentsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2NetworkInsightsAnalysis.Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2NetworkInsightsAnalysisForwardPathComponentsOutputReference_Override(e Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2NetworkInsightsAnalysis.Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference)SetInternalValue(val *Ec2NetworkInsightsAnalysisForwardPathComponents) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisForwardPathComponentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

