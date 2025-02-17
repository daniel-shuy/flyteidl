/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

import (
	"time"
)

type EventNodeExecutionEvent struct {
	Id *CoreNodeExecutionIdentifier `json:"id,omitempty"`
	ProducerId string `json:"producer_id,omitempty"`
	Phase *CoreNodeExecutionPhase `json:"phase,omitempty"`
	// This timestamp represents when the original event occurred, it is generated by the executor of the node.
	OccurredAt time.Time `json:"occurred_at,omitempty"`
	InputUri string `json:"input_uri,omitempty"`
	// URL to the output of the execution, it encodes all the information including Cloud source provider. ie., s3://...
	OutputUri string `json:"output_uri,omitempty"`
	Error_ *CoreExecutionError `json:"error,omitempty"`
	// Raw output data produced by this node execution.
	OutputData *CoreLiteralMap `json:"output_data,omitempty"`
	WorkflowNodeMetadata *FlyteidleventWorkflowNodeMetadata `json:"workflow_node_metadata,omitempty"`
	TaskNodeMetadata *FlyteidleventTaskNodeMetadata `json:"task_node_metadata,omitempty"`
	// [To be deprecated] Specifies which task (if any) launched this node.
	ParentTaskMetadata *EventParentTaskExecutionMetadata `json:"parent_task_metadata,omitempty"`
	// Specifies the parent node of the current node execution. Node executions at level zero will not have a parent node.
	ParentNodeMetadata *EventParentNodeExecutionMetadata `json:"parent_node_metadata,omitempty"`
	RetryGroup string `json:"retry_group,omitempty"`
	SpecNodeId string `json:"spec_node_id,omitempty"`
	NodeName string `json:"node_name,omitempty"`
	EventVersion int32 `json:"event_version,omitempty"`
	// Whether this node launched a subworkflow.
	IsParent bool `json:"is_parent,omitempty"`
	// Whether this node yielded a dynamic workflow.
	IsDynamic bool `json:"is_dynamic,omitempty"`
	DeckUri string `json:"deck_uri,omitempty"`
}
