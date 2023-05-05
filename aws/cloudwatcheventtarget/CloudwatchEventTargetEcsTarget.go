package cloudwatcheventtarget


type CloudwatchEventTargetEcsTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cloudwatch_event_target#task_definition_arn CloudwatchEventTarget#task_definition_arn}.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// capacity_provider_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cloudwatch_event_target#capacity_provider_strategy CloudwatchEventTarget#capacity_provider_strategy}
	CapacityProviderStrategy interface{} `field:"optional" json:"capacityProviderStrategy" yaml:"capacityProviderStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cloudwatch_event_target#enable_ecs_managed_tags CloudwatchEventTarget#enable_ecs_managed_tags}.
	EnableEcsManagedTags interface{} `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cloudwatch_event_target#enable_execute_command CloudwatchEventTarget#enable_execute_command}.
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cloudwatch_event_target#group CloudwatchEventTarget#group}.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cloudwatch_event_target#launch_type CloudwatchEventTarget#launch_type}.
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cloudwatch_event_target#network_configuration CloudwatchEventTarget#network_configuration}
	NetworkConfiguration *CloudwatchEventTargetEcsTargetNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// ordered_placement_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cloudwatch_event_target#ordered_placement_strategy CloudwatchEventTarget#ordered_placement_strategy}
	OrderedPlacementStrategy interface{} `field:"optional" json:"orderedPlacementStrategy" yaml:"orderedPlacementStrategy"`
	// placement_constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cloudwatch_event_target#placement_constraint CloudwatchEventTarget#placement_constraint}
	PlacementConstraint interface{} `field:"optional" json:"placementConstraint" yaml:"placementConstraint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cloudwatch_event_target#platform_version CloudwatchEventTarget#platform_version}.
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cloudwatch_event_target#propagate_tags CloudwatchEventTarget#propagate_tags}.
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cloudwatch_event_target#tags CloudwatchEventTarget#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cloudwatch_event_target#task_count CloudwatchEventTarget#task_count}.
	TaskCount *float64 `field:"optional" json:"taskCount" yaml:"taskCount"`
}

