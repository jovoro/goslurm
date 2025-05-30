/*
Slurm Rest API RO

API to access Slurm. Only GET requests are implemented.

API version: 0.0.38
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v0038

import (
	"encoding/json"
)

// checks if the V0038Partition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0038Partition{}

// V0038Partition struct for V0038Partition
type V0038Partition struct {
	// partition options
	Flags []string `json:"flags,omitempty"`
	// preemption type
	PreemptionMode []string `json:"preemption_mode,omitempty"`
	// list names of allowed allocating nodes
	AllowedAllocationNodes *string `json:"allowed_allocation_nodes,omitempty"`
	// comma delimited list of accounts
	AllowedAccounts *string `json:"allowed_accounts,omitempty"`
	// comma delimited list of groups
	AllowedGroups *string `json:"allowed_groups,omitempty"`
	// comma delimited list of qos
	AllowedQos *string `json:"allowed_qos,omitempty"`
	// name of alternate partition
	Alternative *string `json:"alternative,omitempty"`
	// TRES billing weights
	BillingWeights *string `json:"billing_weights,omitempty"`
	// default MB memory per allocated CPU
	DefaultMemoryPerCpu *int64 `json:"default_memory_per_cpu,omitempty"`
	// default time limit (minutes)
	DefaultTimeLimit *int64 `json:"default_time_limit,omitempty"`
	// comma delimited list of denied accounts
	DeniedAccounts *string `json:"denied_accounts,omitempty"`
	// comma delimited list of denied qos
	DeniedQos *string `json:"denied_qos,omitempty"`
	// preemption grace time (seconds)
	PreemptionGraceTime *int64 `json:"preemption_grace_time,omitempty"`
	// maximum allocated CPUs per node
	MaximumCpusPerNode *int32 `json:"maximum_cpus_per_node,omitempty"`
	// maximum memory per allocated node (MiB)
	MaximumMemoryPerNode *int64 `json:"maximum_memory_per_node,omitempty"`
	// Max nodes per job
	MaximumNodesPerJob *int32 `json:"maximum_nodes_per_job,omitempty"`
	// Max time limit per job
	MaxTimeLimit *int64 `json:"max_time_limit,omitempty"`
	// Min number of nodes per job
	MinNodesPerJob *int32 `json:"min_nodes_per_job,omitempty"`
	// Partition name
	Name *string `json:"name,omitempty"`
	// list names of nodes in partition
	Nodes *string `json:"nodes,omitempty"`
	// job's time limit can be exceeded by this number of minutes before cancellation
	OverTimeLimit *int32 `json:"over_time_limit,omitempty"`
	// job priority weight factor
	PriorityJobFactor *int32 `json:"priority_job_factor,omitempty"`
	// tier for scheduling and preemption
	PriorityTier *int32 `json:"priority_tier,omitempty"`
	// partition QOS name
	Qos *string `json:"qos,omitempty"`
	// Partition state
	State *string `json:"state,omitempty"`
	// Total cpus in partition
	TotalCpus *int32 `json:"total_cpus,omitempty"`
	// Total number of nodes in partition
	TotalNodes *int32 `json:"total_nodes,omitempty"`
	// configured TRES in partition
	Tres *string `json:"tres,omitempty"`
	// maximum memory per allocated CPU (MiB)
	MaximumMemoryPerCpu *int64 `json:"maximum_memory_per_cpu,omitempty"`
	// default MB memory per allocated node
	DefaultMemoryPerNode *int64 `json:"default_memory_per_node,omitempty"`
}

// NewV0038Partition instantiates a new V0038Partition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0038Partition() *V0038Partition {
	this := V0038Partition{}
	return &this
}

// NewV0038PartitionWithDefaults instantiates a new V0038Partition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0038PartitionWithDefaults() *V0038Partition {
	this := V0038Partition{}
	return &this
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0038Partition) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0038Partition) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0038Partition) SetFlags(v []string) {
	o.Flags = v
}

// GetPreemptionMode returns the PreemptionMode field value if set, zero value otherwise.
func (o *V0038Partition) GetPreemptionMode() []string {
	if o == nil || IsNil(o.PreemptionMode) {
		var ret []string
		return ret
	}
	return o.PreemptionMode
}

// GetPreemptionModeOk returns a tuple with the PreemptionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetPreemptionModeOk() ([]string, bool) {
	if o == nil || IsNil(o.PreemptionMode) {
		return nil, false
	}
	return o.PreemptionMode, true
}

// HasPreemptionMode returns a boolean if a field has been set.
func (o *V0038Partition) HasPreemptionMode() bool {
	if o != nil && !IsNil(o.PreemptionMode) {
		return true
	}

	return false
}

// SetPreemptionMode gets a reference to the given []string and assigns it to the PreemptionMode field.
func (o *V0038Partition) SetPreemptionMode(v []string) {
	o.PreemptionMode = v
}

// GetAllowedAllocationNodes returns the AllowedAllocationNodes field value if set, zero value otherwise.
func (o *V0038Partition) GetAllowedAllocationNodes() string {
	if o == nil || IsNil(o.AllowedAllocationNodes) {
		var ret string
		return ret
	}
	return *o.AllowedAllocationNodes
}

// GetAllowedAllocationNodesOk returns a tuple with the AllowedAllocationNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetAllowedAllocationNodesOk() (*string, bool) {
	if o == nil || IsNil(o.AllowedAllocationNodes) {
		return nil, false
	}
	return o.AllowedAllocationNodes, true
}

// HasAllowedAllocationNodes returns a boolean if a field has been set.
func (o *V0038Partition) HasAllowedAllocationNodes() bool {
	if o != nil && !IsNil(o.AllowedAllocationNodes) {
		return true
	}

	return false
}

// SetAllowedAllocationNodes gets a reference to the given string and assigns it to the AllowedAllocationNodes field.
func (o *V0038Partition) SetAllowedAllocationNodes(v string) {
	o.AllowedAllocationNodes = &v
}

// GetAllowedAccounts returns the AllowedAccounts field value if set, zero value otherwise.
func (o *V0038Partition) GetAllowedAccounts() string {
	if o == nil || IsNil(o.AllowedAccounts) {
		var ret string
		return ret
	}
	return *o.AllowedAccounts
}

// GetAllowedAccountsOk returns a tuple with the AllowedAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetAllowedAccountsOk() (*string, bool) {
	if o == nil || IsNil(o.AllowedAccounts) {
		return nil, false
	}
	return o.AllowedAccounts, true
}

// HasAllowedAccounts returns a boolean if a field has been set.
func (o *V0038Partition) HasAllowedAccounts() bool {
	if o != nil && !IsNil(o.AllowedAccounts) {
		return true
	}

	return false
}

// SetAllowedAccounts gets a reference to the given string and assigns it to the AllowedAccounts field.
func (o *V0038Partition) SetAllowedAccounts(v string) {
	o.AllowedAccounts = &v
}

// GetAllowedGroups returns the AllowedGroups field value if set, zero value otherwise.
func (o *V0038Partition) GetAllowedGroups() string {
	if o == nil || IsNil(o.AllowedGroups) {
		var ret string
		return ret
	}
	return *o.AllowedGroups
}

// GetAllowedGroupsOk returns a tuple with the AllowedGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetAllowedGroupsOk() (*string, bool) {
	if o == nil || IsNil(o.AllowedGroups) {
		return nil, false
	}
	return o.AllowedGroups, true
}

// HasAllowedGroups returns a boolean if a field has been set.
func (o *V0038Partition) HasAllowedGroups() bool {
	if o != nil && !IsNil(o.AllowedGroups) {
		return true
	}

	return false
}

// SetAllowedGroups gets a reference to the given string and assigns it to the AllowedGroups field.
func (o *V0038Partition) SetAllowedGroups(v string) {
	o.AllowedGroups = &v
}

// GetAllowedQos returns the AllowedQos field value if set, zero value otherwise.
func (o *V0038Partition) GetAllowedQos() string {
	if o == nil || IsNil(o.AllowedQos) {
		var ret string
		return ret
	}
	return *o.AllowedQos
}

// GetAllowedQosOk returns a tuple with the AllowedQos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetAllowedQosOk() (*string, bool) {
	if o == nil || IsNil(o.AllowedQos) {
		return nil, false
	}
	return o.AllowedQos, true
}

// HasAllowedQos returns a boolean if a field has been set.
func (o *V0038Partition) HasAllowedQos() bool {
	if o != nil && !IsNil(o.AllowedQos) {
		return true
	}

	return false
}

// SetAllowedQos gets a reference to the given string and assigns it to the AllowedQos field.
func (o *V0038Partition) SetAllowedQos(v string) {
	o.AllowedQos = &v
}

// GetAlternative returns the Alternative field value if set, zero value otherwise.
func (o *V0038Partition) GetAlternative() string {
	if o == nil || IsNil(o.Alternative) {
		var ret string
		return ret
	}
	return *o.Alternative
}

// GetAlternativeOk returns a tuple with the Alternative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetAlternativeOk() (*string, bool) {
	if o == nil || IsNil(o.Alternative) {
		return nil, false
	}
	return o.Alternative, true
}

// HasAlternative returns a boolean if a field has been set.
func (o *V0038Partition) HasAlternative() bool {
	if o != nil && !IsNil(o.Alternative) {
		return true
	}

	return false
}

// SetAlternative gets a reference to the given string and assigns it to the Alternative field.
func (o *V0038Partition) SetAlternative(v string) {
	o.Alternative = &v
}

// GetBillingWeights returns the BillingWeights field value if set, zero value otherwise.
func (o *V0038Partition) GetBillingWeights() string {
	if o == nil || IsNil(o.BillingWeights) {
		var ret string
		return ret
	}
	return *o.BillingWeights
}

// GetBillingWeightsOk returns a tuple with the BillingWeights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetBillingWeightsOk() (*string, bool) {
	if o == nil || IsNil(o.BillingWeights) {
		return nil, false
	}
	return o.BillingWeights, true
}

// HasBillingWeights returns a boolean if a field has been set.
func (o *V0038Partition) HasBillingWeights() bool {
	if o != nil && !IsNil(o.BillingWeights) {
		return true
	}

	return false
}

// SetBillingWeights gets a reference to the given string and assigns it to the BillingWeights field.
func (o *V0038Partition) SetBillingWeights(v string) {
	o.BillingWeights = &v
}

// GetDefaultMemoryPerCpu returns the DefaultMemoryPerCpu field value if set, zero value otherwise.
func (o *V0038Partition) GetDefaultMemoryPerCpu() int64 {
	if o == nil || IsNil(o.DefaultMemoryPerCpu) {
		var ret int64
		return ret
	}
	return *o.DefaultMemoryPerCpu
}

// GetDefaultMemoryPerCpuOk returns a tuple with the DefaultMemoryPerCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetDefaultMemoryPerCpuOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultMemoryPerCpu) {
		return nil, false
	}
	return o.DefaultMemoryPerCpu, true
}

// HasDefaultMemoryPerCpu returns a boolean if a field has been set.
func (o *V0038Partition) HasDefaultMemoryPerCpu() bool {
	if o != nil && !IsNil(o.DefaultMemoryPerCpu) {
		return true
	}

	return false
}

// SetDefaultMemoryPerCpu gets a reference to the given int64 and assigns it to the DefaultMemoryPerCpu field.
func (o *V0038Partition) SetDefaultMemoryPerCpu(v int64) {
	o.DefaultMemoryPerCpu = &v
}

// GetDefaultTimeLimit returns the DefaultTimeLimit field value if set, zero value otherwise.
func (o *V0038Partition) GetDefaultTimeLimit() int64 {
	if o == nil || IsNil(o.DefaultTimeLimit) {
		var ret int64
		return ret
	}
	return *o.DefaultTimeLimit
}

// GetDefaultTimeLimitOk returns a tuple with the DefaultTimeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetDefaultTimeLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultTimeLimit) {
		return nil, false
	}
	return o.DefaultTimeLimit, true
}

// HasDefaultTimeLimit returns a boolean if a field has been set.
func (o *V0038Partition) HasDefaultTimeLimit() bool {
	if o != nil && !IsNil(o.DefaultTimeLimit) {
		return true
	}

	return false
}

// SetDefaultTimeLimit gets a reference to the given int64 and assigns it to the DefaultTimeLimit field.
func (o *V0038Partition) SetDefaultTimeLimit(v int64) {
	o.DefaultTimeLimit = &v
}

// GetDeniedAccounts returns the DeniedAccounts field value if set, zero value otherwise.
func (o *V0038Partition) GetDeniedAccounts() string {
	if o == nil || IsNil(o.DeniedAccounts) {
		var ret string
		return ret
	}
	return *o.DeniedAccounts
}

// GetDeniedAccountsOk returns a tuple with the DeniedAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetDeniedAccountsOk() (*string, bool) {
	if o == nil || IsNil(o.DeniedAccounts) {
		return nil, false
	}
	return o.DeniedAccounts, true
}

// HasDeniedAccounts returns a boolean if a field has been set.
func (o *V0038Partition) HasDeniedAccounts() bool {
	if o != nil && !IsNil(o.DeniedAccounts) {
		return true
	}

	return false
}

// SetDeniedAccounts gets a reference to the given string and assigns it to the DeniedAccounts field.
func (o *V0038Partition) SetDeniedAccounts(v string) {
	o.DeniedAccounts = &v
}

// GetDeniedQos returns the DeniedQos field value if set, zero value otherwise.
func (o *V0038Partition) GetDeniedQos() string {
	if o == nil || IsNil(o.DeniedQos) {
		var ret string
		return ret
	}
	return *o.DeniedQos
}

// GetDeniedQosOk returns a tuple with the DeniedQos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetDeniedQosOk() (*string, bool) {
	if o == nil || IsNil(o.DeniedQos) {
		return nil, false
	}
	return o.DeniedQos, true
}

// HasDeniedQos returns a boolean if a field has been set.
func (o *V0038Partition) HasDeniedQos() bool {
	if o != nil && !IsNil(o.DeniedQos) {
		return true
	}

	return false
}

// SetDeniedQos gets a reference to the given string and assigns it to the DeniedQos field.
func (o *V0038Partition) SetDeniedQos(v string) {
	o.DeniedQos = &v
}

// GetPreemptionGraceTime returns the PreemptionGraceTime field value if set, zero value otherwise.
func (o *V0038Partition) GetPreemptionGraceTime() int64 {
	if o == nil || IsNil(o.PreemptionGraceTime) {
		var ret int64
		return ret
	}
	return *o.PreemptionGraceTime
}

// GetPreemptionGraceTimeOk returns a tuple with the PreemptionGraceTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetPreemptionGraceTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.PreemptionGraceTime) {
		return nil, false
	}
	return o.PreemptionGraceTime, true
}

// HasPreemptionGraceTime returns a boolean if a field has been set.
func (o *V0038Partition) HasPreemptionGraceTime() bool {
	if o != nil && !IsNil(o.PreemptionGraceTime) {
		return true
	}

	return false
}

// SetPreemptionGraceTime gets a reference to the given int64 and assigns it to the PreemptionGraceTime field.
func (o *V0038Partition) SetPreemptionGraceTime(v int64) {
	o.PreemptionGraceTime = &v
}

// GetMaximumCpusPerNode returns the MaximumCpusPerNode field value if set, zero value otherwise.
func (o *V0038Partition) GetMaximumCpusPerNode() int32 {
	if o == nil || IsNil(o.MaximumCpusPerNode) {
		var ret int32
		return ret
	}
	return *o.MaximumCpusPerNode
}

// GetMaximumCpusPerNodeOk returns a tuple with the MaximumCpusPerNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetMaximumCpusPerNodeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumCpusPerNode) {
		return nil, false
	}
	return o.MaximumCpusPerNode, true
}

// HasMaximumCpusPerNode returns a boolean if a field has been set.
func (o *V0038Partition) HasMaximumCpusPerNode() bool {
	if o != nil && !IsNil(o.MaximumCpusPerNode) {
		return true
	}

	return false
}

// SetMaximumCpusPerNode gets a reference to the given int32 and assigns it to the MaximumCpusPerNode field.
func (o *V0038Partition) SetMaximumCpusPerNode(v int32) {
	o.MaximumCpusPerNode = &v
}

// GetMaximumMemoryPerNode returns the MaximumMemoryPerNode field value if set, zero value otherwise.
func (o *V0038Partition) GetMaximumMemoryPerNode() int64 {
	if o == nil || IsNil(o.MaximumMemoryPerNode) {
		var ret int64
		return ret
	}
	return *o.MaximumMemoryPerNode
}

// GetMaximumMemoryPerNodeOk returns a tuple with the MaximumMemoryPerNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetMaximumMemoryPerNodeOk() (*int64, bool) {
	if o == nil || IsNil(o.MaximumMemoryPerNode) {
		return nil, false
	}
	return o.MaximumMemoryPerNode, true
}

// HasMaximumMemoryPerNode returns a boolean if a field has been set.
func (o *V0038Partition) HasMaximumMemoryPerNode() bool {
	if o != nil && !IsNil(o.MaximumMemoryPerNode) {
		return true
	}

	return false
}

// SetMaximumMemoryPerNode gets a reference to the given int64 and assigns it to the MaximumMemoryPerNode field.
func (o *V0038Partition) SetMaximumMemoryPerNode(v int64) {
	o.MaximumMemoryPerNode = &v
}

// GetMaximumNodesPerJob returns the MaximumNodesPerJob field value if set, zero value otherwise.
func (o *V0038Partition) GetMaximumNodesPerJob() int32 {
	if o == nil || IsNil(o.MaximumNodesPerJob) {
		var ret int32
		return ret
	}
	return *o.MaximumNodesPerJob
}

// GetMaximumNodesPerJobOk returns a tuple with the MaximumNodesPerJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetMaximumNodesPerJobOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumNodesPerJob) {
		return nil, false
	}
	return o.MaximumNodesPerJob, true
}

// HasMaximumNodesPerJob returns a boolean if a field has been set.
func (o *V0038Partition) HasMaximumNodesPerJob() bool {
	if o != nil && !IsNil(o.MaximumNodesPerJob) {
		return true
	}

	return false
}

// SetMaximumNodesPerJob gets a reference to the given int32 and assigns it to the MaximumNodesPerJob field.
func (o *V0038Partition) SetMaximumNodesPerJob(v int32) {
	o.MaximumNodesPerJob = &v
}

// GetMaxTimeLimit returns the MaxTimeLimit field value if set, zero value otherwise.
func (o *V0038Partition) GetMaxTimeLimit() int64 {
	if o == nil || IsNil(o.MaxTimeLimit) {
		var ret int64
		return ret
	}
	return *o.MaxTimeLimit
}

// GetMaxTimeLimitOk returns a tuple with the MaxTimeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetMaxTimeLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxTimeLimit) {
		return nil, false
	}
	return o.MaxTimeLimit, true
}

// HasMaxTimeLimit returns a boolean if a field has been set.
func (o *V0038Partition) HasMaxTimeLimit() bool {
	if o != nil && !IsNil(o.MaxTimeLimit) {
		return true
	}

	return false
}

// SetMaxTimeLimit gets a reference to the given int64 and assigns it to the MaxTimeLimit field.
func (o *V0038Partition) SetMaxTimeLimit(v int64) {
	o.MaxTimeLimit = &v
}

// GetMinNodesPerJob returns the MinNodesPerJob field value if set, zero value otherwise.
func (o *V0038Partition) GetMinNodesPerJob() int32 {
	if o == nil || IsNil(o.MinNodesPerJob) {
		var ret int32
		return ret
	}
	return *o.MinNodesPerJob
}

// GetMinNodesPerJobOk returns a tuple with the MinNodesPerJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetMinNodesPerJobOk() (*int32, bool) {
	if o == nil || IsNil(o.MinNodesPerJob) {
		return nil, false
	}
	return o.MinNodesPerJob, true
}

// HasMinNodesPerJob returns a boolean if a field has been set.
func (o *V0038Partition) HasMinNodesPerJob() bool {
	if o != nil && !IsNil(o.MinNodesPerJob) {
		return true
	}

	return false
}

// SetMinNodesPerJob gets a reference to the given int32 and assigns it to the MinNodesPerJob field.
func (o *V0038Partition) SetMinNodesPerJob(v int32) {
	o.MinNodesPerJob = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0038Partition) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0038Partition) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0038Partition) SetName(v string) {
	o.Name = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *V0038Partition) GetNodes() string {
	if o == nil || IsNil(o.Nodes) {
		var ret string
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetNodesOk() (*string, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *V0038Partition) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given string and assigns it to the Nodes field.
func (o *V0038Partition) SetNodes(v string) {
	o.Nodes = &v
}

// GetOverTimeLimit returns the OverTimeLimit field value if set, zero value otherwise.
func (o *V0038Partition) GetOverTimeLimit() int32 {
	if o == nil || IsNil(o.OverTimeLimit) {
		var ret int32
		return ret
	}
	return *o.OverTimeLimit
}

// GetOverTimeLimitOk returns a tuple with the OverTimeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetOverTimeLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.OverTimeLimit) {
		return nil, false
	}
	return o.OverTimeLimit, true
}

// HasOverTimeLimit returns a boolean if a field has been set.
func (o *V0038Partition) HasOverTimeLimit() bool {
	if o != nil && !IsNil(o.OverTimeLimit) {
		return true
	}

	return false
}

// SetOverTimeLimit gets a reference to the given int32 and assigns it to the OverTimeLimit field.
func (o *V0038Partition) SetOverTimeLimit(v int32) {
	o.OverTimeLimit = &v
}

// GetPriorityJobFactor returns the PriorityJobFactor field value if set, zero value otherwise.
func (o *V0038Partition) GetPriorityJobFactor() int32 {
	if o == nil || IsNil(o.PriorityJobFactor) {
		var ret int32
		return ret
	}
	return *o.PriorityJobFactor
}

// GetPriorityJobFactorOk returns a tuple with the PriorityJobFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetPriorityJobFactorOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityJobFactor) {
		return nil, false
	}
	return o.PriorityJobFactor, true
}

// HasPriorityJobFactor returns a boolean if a field has been set.
func (o *V0038Partition) HasPriorityJobFactor() bool {
	if o != nil && !IsNil(o.PriorityJobFactor) {
		return true
	}

	return false
}

// SetPriorityJobFactor gets a reference to the given int32 and assigns it to the PriorityJobFactor field.
func (o *V0038Partition) SetPriorityJobFactor(v int32) {
	o.PriorityJobFactor = &v
}

// GetPriorityTier returns the PriorityTier field value if set, zero value otherwise.
func (o *V0038Partition) GetPriorityTier() int32 {
	if o == nil || IsNil(o.PriorityTier) {
		var ret int32
		return ret
	}
	return *o.PriorityTier
}

// GetPriorityTierOk returns a tuple with the PriorityTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetPriorityTierOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityTier) {
		return nil, false
	}
	return o.PriorityTier, true
}

// HasPriorityTier returns a boolean if a field has been set.
func (o *V0038Partition) HasPriorityTier() bool {
	if o != nil && !IsNil(o.PriorityTier) {
		return true
	}

	return false
}

// SetPriorityTier gets a reference to the given int32 and assigns it to the PriorityTier field.
func (o *V0038Partition) SetPriorityTier(v int32) {
	o.PriorityTier = &v
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *V0038Partition) GetQos() string {
	if o == nil || IsNil(o.Qos) {
		var ret string
		return ret
	}
	return *o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetQosOk() (*string, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *V0038Partition) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given string and assigns it to the Qos field.
func (o *V0038Partition) SetQos(v string) {
	o.Qos = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *V0038Partition) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *V0038Partition) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *V0038Partition) SetState(v string) {
	o.State = &v
}

// GetTotalCpus returns the TotalCpus field value if set, zero value otherwise.
func (o *V0038Partition) GetTotalCpus() int32 {
	if o == nil || IsNil(o.TotalCpus) {
		var ret int32
		return ret
	}
	return *o.TotalCpus
}

// GetTotalCpusOk returns a tuple with the TotalCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetTotalCpusOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCpus) {
		return nil, false
	}
	return o.TotalCpus, true
}

// HasTotalCpus returns a boolean if a field has been set.
func (o *V0038Partition) HasTotalCpus() bool {
	if o != nil && !IsNil(o.TotalCpus) {
		return true
	}

	return false
}

// SetTotalCpus gets a reference to the given int32 and assigns it to the TotalCpus field.
func (o *V0038Partition) SetTotalCpus(v int32) {
	o.TotalCpus = &v
}

// GetTotalNodes returns the TotalNodes field value if set, zero value otherwise.
func (o *V0038Partition) GetTotalNodes() int32 {
	if o == nil || IsNil(o.TotalNodes) {
		var ret int32
		return ret
	}
	return *o.TotalNodes
}

// GetTotalNodesOk returns a tuple with the TotalNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetTotalNodesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalNodes) {
		return nil, false
	}
	return o.TotalNodes, true
}

// HasTotalNodes returns a boolean if a field has been set.
func (o *V0038Partition) HasTotalNodes() bool {
	if o != nil && !IsNil(o.TotalNodes) {
		return true
	}

	return false
}

// SetTotalNodes gets a reference to the given int32 and assigns it to the TotalNodes field.
func (o *V0038Partition) SetTotalNodes(v int32) {
	o.TotalNodes = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *V0038Partition) GetTres() string {
	if o == nil || IsNil(o.Tres) {
		var ret string
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetTresOk() (*string, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *V0038Partition) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given string and assigns it to the Tres field.
func (o *V0038Partition) SetTres(v string) {
	o.Tres = &v
}

// GetMaximumMemoryPerCpu returns the MaximumMemoryPerCpu field value if set, zero value otherwise.
func (o *V0038Partition) GetMaximumMemoryPerCpu() int64 {
	if o == nil || IsNil(o.MaximumMemoryPerCpu) {
		var ret int64
		return ret
	}
	return *o.MaximumMemoryPerCpu
}

// GetMaximumMemoryPerCpuOk returns a tuple with the MaximumMemoryPerCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetMaximumMemoryPerCpuOk() (*int64, bool) {
	if o == nil || IsNil(o.MaximumMemoryPerCpu) {
		return nil, false
	}
	return o.MaximumMemoryPerCpu, true
}

// HasMaximumMemoryPerCpu returns a boolean if a field has been set.
func (o *V0038Partition) HasMaximumMemoryPerCpu() bool {
	if o != nil && !IsNil(o.MaximumMemoryPerCpu) {
		return true
	}

	return false
}

// SetMaximumMemoryPerCpu gets a reference to the given int64 and assigns it to the MaximumMemoryPerCpu field.
func (o *V0038Partition) SetMaximumMemoryPerCpu(v int64) {
	o.MaximumMemoryPerCpu = &v
}

// GetDefaultMemoryPerNode returns the DefaultMemoryPerNode field value if set, zero value otherwise.
func (o *V0038Partition) GetDefaultMemoryPerNode() int64 {
	if o == nil || IsNil(o.DefaultMemoryPerNode) {
		var ret int64
		return ret
	}
	return *o.DefaultMemoryPerNode
}

// GetDefaultMemoryPerNodeOk returns a tuple with the DefaultMemoryPerNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Partition) GetDefaultMemoryPerNodeOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultMemoryPerNode) {
		return nil, false
	}
	return o.DefaultMemoryPerNode, true
}

// HasDefaultMemoryPerNode returns a boolean if a field has been set.
func (o *V0038Partition) HasDefaultMemoryPerNode() bool {
	if o != nil && !IsNil(o.DefaultMemoryPerNode) {
		return true
	}

	return false
}

// SetDefaultMemoryPerNode gets a reference to the given int64 and assigns it to the DefaultMemoryPerNode field.
func (o *V0038Partition) SetDefaultMemoryPerNode(v int64) {
	o.DefaultMemoryPerNode = &v
}

func (o V0038Partition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0038Partition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.PreemptionMode) {
		toSerialize["preemption_mode"] = o.PreemptionMode
	}
	if !IsNil(o.AllowedAllocationNodes) {
		toSerialize["allowed_allocation_nodes"] = o.AllowedAllocationNodes
	}
	if !IsNil(o.AllowedAccounts) {
		toSerialize["allowed_accounts"] = o.AllowedAccounts
	}
	if !IsNil(o.AllowedGroups) {
		toSerialize["allowed_groups"] = o.AllowedGroups
	}
	if !IsNil(o.AllowedQos) {
		toSerialize["allowed_qos"] = o.AllowedQos
	}
	if !IsNil(o.Alternative) {
		toSerialize["alternative"] = o.Alternative
	}
	if !IsNil(o.BillingWeights) {
		toSerialize["billing_weights"] = o.BillingWeights
	}
	if !IsNil(o.DefaultMemoryPerCpu) {
		toSerialize["default_memory_per_cpu"] = o.DefaultMemoryPerCpu
	}
	if !IsNil(o.DefaultTimeLimit) {
		toSerialize["default_time_limit"] = o.DefaultTimeLimit
	}
	if !IsNil(o.DeniedAccounts) {
		toSerialize["denied_accounts"] = o.DeniedAccounts
	}
	if !IsNil(o.DeniedQos) {
		toSerialize["denied_qos"] = o.DeniedQos
	}
	if !IsNil(o.PreemptionGraceTime) {
		toSerialize["preemption_grace_time"] = o.PreemptionGraceTime
	}
	if !IsNil(o.MaximumCpusPerNode) {
		toSerialize["maximum_cpus_per_node"] = o.MaximumCpusPerNode
	}
	if !IsNil(o.MaximumMemoryPerNode) {
		toSerialize["maximum_memory_per_node"] = o.MaximumMemoryPerNode
	}
	if !IsNil(o.MaximumNodesPerJob) {
		toSerialize["maximum_nodes_per_job"] = o.MaximumNodesPerJob
	}
	if !IsNil(o.MaxTimeLimit) {
		toSerialize["max_time_limit"] = o.MaxTimeLimit
	}
	if !IsNil(o.MinNodesPerJob) {
		toSerialize["min_nodes_per_job"] = o.MinNodesPerJob
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	if !IsNil(o.OverTimeLimit) {
		toSerialize["over_time_limit"] = o.OverTimeLimit
	}
	if !IsNil(o.PriorityJobFactor) {
		toSerialize["priority_job_factor"] = o.PriorityJobFactor
	}
	if !IsNil(o.PriorityTier) {
		toSerialize["priority_tier"] = o.PriorityTier
	}
	if !IsNil(o.Qos) {
		toSerialize["qos"] = o.Qos
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.TotalCpus) {
		toSerialize["total_cpus"] = o.TotalCpus
	}
	if !IsNil(o.TotalNodes) {
		toSerialize["total_nodes"] = o.TotalNodes
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	if !IsNil(o.MaximumMemoryPerCpu) {
		toSerialize["maximum_memory_per_cpu"] = o.MaximumMemoryPerCpu
	}
	if !IsNil(o.DefaultMemoryPerNode) {
		toSerialize["default_memory_per_node"] = o.DefaultMemoryPerNode
	}
	return toSerialize, nil
}

type NullableV0038Partition struct {
	value *V0038Partition
	isSet bool
}

func (v NullableV0038Partition) Get() *V0038Partition {
	return v.value
}

func (v *NullableV0038Partition) Set(val *V0038Partition) {
	v.value = val
	v.isSet = true
}

func (v NullableV0038Partition) IsSet() bool {
	return v.isSet
}

func (v *NullableV0038Partition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0038Partition(val *V0038Partition) *NullableV0038Partition {
	return &NullableV0038Partition{value: val, isSet: true}
}

func (v NullableV0038Partition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0038Partition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


