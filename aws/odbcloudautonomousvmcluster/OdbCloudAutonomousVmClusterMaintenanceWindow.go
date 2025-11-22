// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package odbcloudautonomousvmcluster


type OdbCloudAutonomousVmClusterMaintenanceWindow struct {
	// The preference for the maintenance window scheduling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/odb_cloud_autonomous_vm_cluster#preference OdbCloudAutonomousVmCluster#preference}
	Preference *string `field:"required" json:"preference" yaml:"preference"`
	// The days of the week when maintenance can be performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/odb_cloud_autonomous_vm_cluster#days_of_week OdbCloudAutonomousVmCluster#days_of_week}
	DaysOfWeek interface{} `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// The hours of the day when maintenance can be performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/odb_cloud_autonomous_vm_cluster#hours_of_day OdbCloudAutonomousVmCluster#hours_of_day}
	HoursOfDay *[]*float64 `field:"optional" json:"hoursOfDay" yaml:"hoursOfDay"`
	// The lead time in weeks before the maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/odb_cloud_autonomous_vm_cluster#lead_time_in_weeks OdbCloudAutonomousVmCluster#lead_time_in_weeks}
	LeadTimeInWeeks *float64 `field:"optional" json:"leadTimeInWeeks" yaml:"leadTimeInWeeks"`
	// The months when maintenance can be performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/odb_cloud_autonomous_vm_cluster#months OdbCloudAutonomousVmCluster#months}
	Months interface{} `field:"optional" json:"months" yaml:"months"`
	// Indicates whether to skip release updates during maintenance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/odb_cloud_autonomous_vm_cluster#weeks_of_month OdbCloudAutonomousVmCluster#weeks_of_month}
	WeeksOfMonth *[]*float64 `field:"optional" json:"weeksOfMonth" yaml:"weeksOfMonth"`
}

