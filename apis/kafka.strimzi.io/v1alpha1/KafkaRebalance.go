// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package v1alpha1

// KafkaRebalance
type KafkaRebalance struct {
	// The specification of the Kafka rebalance.
	Spec *KafkaRebalanceSpec `json:"spec,omitempty"`

	// The status of the Kafka rebalance.
	Status *KafkaRebalanceStatus `json:"status,omitempty"`
}

// The specification of the Kafka rebalance.
type KafkaRebalanceSpec struct {
	// The upper bound of ongoing partition replica movements between disks within
	// each broker. Default is 2.
	ConcurrentIntraBrokerPartitionMovements *int `json:"concurrentIntraBrokerPartitionMovements,omitempty"`

	// The upper bound of ongoing partition leadership movements. Default is 1000.
	ConcurrentLeaderMovements *int `json:"concurrentLeaderMovements,omitempty"`

	// The upper bound of ongoing partition replica movements going into/out of each
	// broker. Default is 5.
	ConcurrentPartitionMovementsPerBroker *int `json:"concurrentPartitionMovementsPerBroker,omitempty"`

	// A regular expression where any matching topics will be excluded from the
	// calculation of optimization proposals. This expression will be parsed by the
	// java.util.regex.Pattern class; for more information on the supported formar
	// consult the documentation for that class.
	ExcludedTopics *string `json:"excludedTopics,omitempty"`

	// A list of goals, ordered by decreasing priority, to use for generating and
	// executing the rebalance proposal. The supported goals are available at
	// https://github.com/linkedin/cruise-control#goals. If an empty goals list is
	// provided, the goals declared in the default.goals Cruise Control configuration
	// parameter are used.
	Goals []string `json:"goals,omitempty"`

	// A list of strategy class names used to determine the execution order for the
	// replica movements in the generated optimization proposal. By default
	// BaseReplicaMovementStrategy is used, which will execute the replica movements
	// in the order that they were generated.
	ReplicaMovementStrategies []string `json:"replicaMovementStrategies,omitempty"`

	// The upper bound, in bytes per second, on the bandwidth used to move replicas.
	// There is no limit by default.
	ReplicationThrottle *int `json:"replicationThrottle,omitempty"`

	// Whether to allow the hard goals specified in the Kafka CR to be skipped in
	// optimization proposal generation. This can be useful when some of those hard
	// goals are preventing a balance solution being found. Default is false.
	SkipHardGoalCheck *bool `json:"skipHardGoalCheck,omitempty"`
}

// The status of the Kafka rebalance.
type KafkaRebalanceStatus struct {
	// List of status conditions.
	Conditions []KafkaRebalanceStatusConditionsElem `json:"conditions,omitempty"`

	// The generation of the CRD that was last reconciled by the operator.
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	// A JSON object describing the optimization result.
	OptimizationResult KafkaRebalanceStatusOptimizationResult `json:"optimizationResult,omitempty"`

	// The session identifier for requests to Cruise Control pertaining to this
	// KafkaRebalance resource. This is used by the Kafka Rebalance operator to track
	// the status of ongoing rebalancing operations.
	SessionId *string `json:"sessionId,omitempty"`
}

type KafkaRebalanceStatusConditionsElem struct {
	// Last time the condition of a type changed from one status to another. The
	// required format is 'yyyy-MM-ddTHH:mm:ssZ', in the UTC time zone.
	LastTransitionTime *string `json:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about the condition's last
	// transition.
	Message *string `json:"message,omitempty"`

	// The reason for the condition's last transition (a single word in CamelCase).
	Reason *string `json:"reason,omitempty"`

	// The status of the condition, either True, False or Unknown.
	Status *string `json:"status,omitempty"`

	// The unique identifier of a condition, used to distinguish between other
	// conditions in the resource.
	Type *string `json:"type,omitempty"`
}

// A JSON object describing the optimization result.
type KafkaRebalanceStatusOptimizationResult map[string]interface{}
