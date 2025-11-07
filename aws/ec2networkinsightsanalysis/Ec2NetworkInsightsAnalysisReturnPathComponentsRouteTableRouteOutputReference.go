// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2networkinsightsanalysis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ec2networkinsightsanalysis/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference interface {
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
	DestinationCidr() *string
	DestinationPrefixListId() *string
	EgressOnlyInternetGatewayId() *string
	// Experimental.
	Fqn() *string
	GatewayId() *string
	InstanceId() *string
	InternalValue() *Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRoute
	SetInternalValue(val *Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRoute)
	NatGatewayId() *string
	NetworkInterfaceId() *string
	Origin() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransitGatewayId() *string
	VpcPeeringConnectionId() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference
type jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) DestinationCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) DestinationPrefixListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPrefixListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) EgressOnlyInternetGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egressOnlyInternetGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) GatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) InternalValue() *Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRoute {
	var returns *Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRoute
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) NatGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) NetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) Origin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) TransitGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) VpcPeeringConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcPeeringConnectionId",
		&returns,
	)
	return returns
}


func NewEc2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference {
	_init_.Initialize()

	if err := validateNewEc2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2NetworkInsightsAnalysis.Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference_Override(e Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2NetworkInsightsAnalysis.Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference)SetInternalValue(val *Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRoute) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisReturnPathComponentsRouteTableRouteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

