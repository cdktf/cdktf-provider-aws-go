// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connectuserhierarchystructure


type ConnectUserHierarchyStructureHierarchyStructure struct {
	// level_five block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/connect_user_hierarchy_structure#level_five ConnectUserHierarchyStructure#level_five}
	LevelFive *ConnectUserHierarchyStructureHierarchyStructureLevelFive `field:"optional" json:"levelFive" yaml:"levelFive"`
	// level_four block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/connect_user_hierarchy_structure#level_four ConnectUserHierarchyStructure#level_four}
	LevelFour *ConnectUserHierarchyStructureHierarchyStructureLevelFour `field:"optional" json:"levelFour" yaml:"levelFour"`
	// level_one block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/connect_user_hierarchy_structure#level_one ConnectUserHierarchyStructure#level_one}
	LevelOne *ConnectUserHierarchyStructureHierarchyStructureLevelOne `field:"optional" json:"levelOne" yaml:"levelOne"`
	// level_three block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/connect_user_hierarchy_structure#level_three ConnectUserHierarchyStructure#level_three}
	LevelThree *ConnectUserHierarchyStructureHierarchyStructureLevelThree `field:"optional" json:"levelThree" yaml:"levelThree"`
	// level_two block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/connect_user_hierarchy_structure#level_two ConnectUserHierarchyStructure#level_two}
	LevelTwo *ConnectUserHierarchyStructureHierarchyStructureLevelTwo `field:"optional" json:"levelTwo" yaml:"levelTwo"`
}

