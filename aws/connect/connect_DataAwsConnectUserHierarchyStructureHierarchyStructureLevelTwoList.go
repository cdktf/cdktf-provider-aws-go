package connect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/connect/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList interface {
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
	Get(index *float64) DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList
type jsiiProxy_DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList {
	_init_.Initialize()

	if err := validateNewDataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList{}

	_jsii_.Create(
		"@cdktf/provider-aws.connect.DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList_Override(d DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.connect.DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList) Get(index *float64) DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsConnectUserHierarchyStructureHierarchyStructureLevelTwoList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

