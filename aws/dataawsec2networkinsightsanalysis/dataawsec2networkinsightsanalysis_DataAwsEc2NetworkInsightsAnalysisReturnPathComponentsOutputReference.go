package dataawsec2networkinsightsanalysis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/dataawsec2networkinsightsanalysis/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference interface {
	cdktf.ComplexObject
	AclRule() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsAclRuleList
	AdditionalDetails() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsAdditionalDetailsList
	AttachedTo() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsAttachedToList
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
	Component() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsComponentList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DestinationVpc() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsDestinationVpcList
	// Experimental.
	Fqn() *string
	InboundHeader() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsInboundHeaderList
	InternalValue() *DataAwsEc2NetworkInsightsAnalysisReturnPathComponents
	SetInternalValue(val *DataAwsEc2NetworkInsightsAnalysisReturnPathComponents)
	OutboundHeader() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutboundHeaderList
	RouteTableRoute() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteList
	SecurityGroupRule() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsSecurityGroupRuleList
	SequenceNumber() *float64
	SourceVpc() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsSourceVpcList
	Subnet() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsSubnetList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransitGateway() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsTransitGatewayList
	TransitGatewayRouteTableRoute() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsTransitGatewayRouteTableRouteList
	Vpc() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsVpcList
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

// The jsii proxy struct for DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference
type jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) AclRule() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsAclRuleList {
	var returns DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsAclRuleList
	_jsii_.Get(
		j,
		"aclRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) AdditionalDetails() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsAdditionalDetailsList {
	var returns DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsAdditionalDetailsList
	_jsii_.Get(
		j,
		"additionalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) AttachedTo() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsAttachedToList {
	var returns DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsAttachedToList
	_jsii_.Get(
		j,
		"attachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) Component() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsComponentList {
	var returns DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsComponentList
	_jsii_.Get(
		j,
		"component",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) DestinationVpc() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsDestinationVpcList {
	var returns DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsDestinationVpcList
	_jsii_.Get(
		j,
		"destinationVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) InboundHeader() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsInboundHeaderList {
	var returns DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsInboundHeaderList
	_jsii_.Get(
		j,
		"inboundHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) InternalValue() *DataAwsEc2NetworkInsightsAnalysisReturnPathComponents {
	var returns *DataAwsEc2NetworkInsightsAnalysisReturnPathComponents
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) OutboundHeader() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutboundHeaderList {
	var returns DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutboundHeaderList
	_jsii_.Get(
		j,
		"outboundHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) RouteTableRoute() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteList {
	var returns DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteList
	_jsii_.Get(
		j,
		"routeTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) SecurityGroupRule() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsSecurityGroupRuleList {
	var returns DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsSecurityGroupRuleList
	_jsii_.Get(
		j,
		"securityGroupRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) SequenceNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sequenceNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) SourceVpc() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsSourceVpcList {
	var returns DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsSourceVpcList
	_jsii_.Get(
		j,
		"sourceVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) Subnet() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsSubnetList {
	var returns DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsSubnetList
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) TransitGateway() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsTransitGatewayList {
	var returns DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsTransitGatewayList
	_jsii_.Get(
		j,
		"transitGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) TransitGatewayRouteTableRoute() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsTransitGatewayRouteTableRouteList {
	var returns DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsTransitGatewayRouteTableRouteList
	_jsii_.Get(
		j,
		"transitGatewayRouteTableRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) Vpc() DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsVpcList {
	var returns DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsVpcList
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewDataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsEc2NetworkInsightsAnalysis.DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference_Override(d DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsEc2NetworkInsightsAnalysis.DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference)SetInternalValue(val *DataAwsEc2NetworkInsightsAnalysisReturnPathComponents) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisReturnPathComponentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

