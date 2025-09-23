// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package macie2classificationjob


type Macie2ClassificationJobScheduleFrequency struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/macie2_classification_job#daily_schedule Macie2ClassificationJob#daily_schedule}.
	DailySchedule interface{} `field:"optional" json:"dailySchedule" yaml:"dailySchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/macie2_classification_job#monthly_schedule Macie2ClassificationJob#monthly_schedule}.
	MonthlySchedule *float64 `field:"optional" json:"monthlySchedule" yaml:"monthlySchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/macie2_classification_job#weekly_schedule Macie2ClassificationJob#weekly_schedule}.
	WeeklySchedule *string `field:"optional" json:"weeklySchedule" yaml:"weeklySchedule"`
}

