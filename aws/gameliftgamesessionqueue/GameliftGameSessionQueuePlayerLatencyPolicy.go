package gameliftgamesessionqueue


type GameliftGameSessionQueuePlayerLatencyPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/gamelift_game_session_queue#maximum_individual_player_latency_milliseconds GameliftGameSessionQueue#maximum_individual_player_latency_milliseconds}.
	MaximumIndividualPlayerLatencyMilliseconds *float64 `field:"required" json:"maximumIndividualPlayerLatencyMilliseconds" yaml:"maximumIndividualPlayerLatencyMilliseconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/gamelift_game_session_queue#policy_duration_seconds GameliftGameSessionQueue#policy_duration_seconds}.
	PolicyDurationSeconds *float64 `field:"optional" json:"policyDurationSeconds" yaml:"policyDurationSeconds"`
}

