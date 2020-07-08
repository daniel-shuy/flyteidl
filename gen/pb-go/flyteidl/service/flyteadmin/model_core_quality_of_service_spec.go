/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Represents customized execution run-time attributes.
type CoreQualityOfServiceSpec struct {
	// Indicates how much queueing delay an execution can tolerate.
	QueueingBudget string `json:"queueing_budget,omitempty"`
}
