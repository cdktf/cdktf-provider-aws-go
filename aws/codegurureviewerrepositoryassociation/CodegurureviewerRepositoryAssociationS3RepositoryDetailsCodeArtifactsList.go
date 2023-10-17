// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codegurureviewerrepositoryassociation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/codegurureviewerrepositoryassociation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList interface {
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
	Get(index *float64) CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList
type jsiiProxy_CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList {
	_init_.Initialize()

	if err := validateNewCodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList{}

	_jsii_.Create(
		"@cdktf/provider-aws.codegurureviewerRepositoryAssociation.CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList_Override(c CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.codegurureviewerRepositoryAssociation.CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList) Get(index *float64) CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsOutputReference {
	if err := c.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationS3RepositoryDetailsCodeArtifactsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

