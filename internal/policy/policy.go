// Package policy defines the access-policy evaluation contract.
// Interfaces only for now; implementations come in a later iteration.
package policy

import "context"

// Decision is the result of evaluating a policy for a retrieval request.
type Decision struct {
	Allow  bool
	Reason string
}

// Evaluator decides whether a subject may access a resource
// (e.g. a document chunk or a vector DB collection).
type Evaluator interface {
	Evaluate(ctx context.Context, subject, resource string) (Decision, error)
}
