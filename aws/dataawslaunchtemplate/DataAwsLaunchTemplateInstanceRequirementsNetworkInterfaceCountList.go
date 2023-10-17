// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawslaunchtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/dataawslaunchtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList interface {
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
	Get(index *float64) DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList
type jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList {
	_init_.Initialize()

	if err := validateNewDataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsLaunchTemplate.DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList_Override(d DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsLaunchTemplate.DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList) Get(index *float64) DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsLaunchTemplateInstanceRequirementsNetworkInterfaceCountList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

