// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rumappmonitor


type RumAppMonitorAppMonitorConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/rum_app_monitor#allow_cookies RumAppMonitor#allow_cookies}.
	AllowCookies interface{} `field:"optional" json:"allowCookies" yaml:"allowCookies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/rum_app_monitor#enable_xray RumAppMonitor#enable_xray}.
	EnableXray interface{} `field:"optional" json:"enableXray" yaml:"enableXray"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/rum_app_monitor#excluded_pages RumAppMonitor#excluded_pages}.
	ExcludedPages *[]*string `field:"optional" json:"excludedPages" yaml:"excludedPages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/rum_app_monitor#favorite_pages RumAppMonitor#favorite_pages}.
	FavoritePages *[]*string `field:"optional" json:"favoritePages" yaml:"favoritePages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/rum_app_monitor#guest_role_arn RumAppMonitor#guest_role_arn}.
	GuestRoleArn *string `field:"optional" json:"guestRoleArn" yaml:"guestRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/rum_app_monitor#identity_pool_id RumAppMonitor#identity_pool_id}.
	IdentityPoolId *string `field:"optional" json:"identityPoolId" yaml:"identityPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/rum_app_monitor#included_pages RumAppMonitor#included_pages}.
	IncludedPages *[]*string `field:"optional" json:"includedPages" yaml:"includedPages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/rum_app_monitor#session_sample_rate RumAppMonitor#session_sample_rate}.
	SessionSampleRate *float64 `field:"optional" json:"sessionSampleRate" yaml:"sessionSampleRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/rum_app_monitor#telemetries RumAppMonitor#telemetries}.
	Telemetries *[]*string `field:"optional" json:"telemetries" yaml:"telemetries"`
}

