// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eksnodegroup


type EksNodeGroupNodeRepairConfigNodeRepairConfigOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/eks_node_group#min_repair_wait_time_mins EksNodeGroup#min_repair_wait_time_mins}.
	MinRepairWaitTimeMins *float64 `field:"required" json:"minRepairWaitTimeMins" yaml:"minRepairWaitTimeMins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/eks_node_group#node_monitoring_condition EksNodeGroup#node_monitoring_condition}.
	NodeMonitoringCondition *string `field:"required" json:"nodeMonitoringCondition" yaml:"nodeMonitoringCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/eks_node_group#node_unhealthy_reason EksNodeGroup#node_unhealthy_reason}.
	NodeUnhealthyReason *string `field:"required" json:"nodeUnhealthyReason" yaml:"nodeUnhealthyReason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/eks_node_group#repair_action EksNodeGroup#repair_action}.
	RepairAction *string `field:"required" json:"repairAction" yaml:"repairAction"`
}

