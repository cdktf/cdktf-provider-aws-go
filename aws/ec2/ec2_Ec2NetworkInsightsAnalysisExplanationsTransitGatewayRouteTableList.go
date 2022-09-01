package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/ec2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList
type jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEc2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList {
	_init_.Initialize()

	if err := validateNewEc2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2.Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEc2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList_Override(e Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2.Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList) Get(index *float64) Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableOutputReference {
	if err := e.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsTransitGatewayRouteTableList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

