// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AgentStatus undocumented
type AgentStatus int

const (
	// AgentStatusVActive undocumented
	AgentStatusVActive AgentStatus = 0
	// AgentStatusVInactive undocumented
	AgentStatusVInactive AgentStatus = 1
)

// AgentStatusPActive returns a pointer to AgentStatusVActive
func AgentStatusPActive() *AgentStatus {
	v := AgentStatusVActive
	return &v
}

// AgentStatusPInactive returns a pointer to AgentStatusVInactive
func AgentStatusPInactive() *AgentStatus {
	v := AgentStatusVInactive
	return &v
}