// Code generated by ent, DO NOT EDIT.

package bugs

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Bugs {
	return predicate.Bugs(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Bugs {
	return predicate.Bugs(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Bugs {
	return predicate.Bugs(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Bugs {
	return predicate.Bugs(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Bugs {
	return predicate.Bugs(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Bugs {
	return predicate.Bugs(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Bugs {
	return predicate.Bugs(sql.FieldLTE(FieldID, id))
}

// JiraKey applies equality check predicate on the "jira_key" field. It's identical to JiraKeyEQ.
func JiraKey(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldJiraKey, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldUpdatedAt, v))
}

// ResolvedAt applies equality check predicate on the "resolved_at" field. It's identical to ResolvedAtEQ.
func ResolvedAt(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldResolvedAt, v))
}

// Resolved applies equality check predicate on the "resolved" field. It's identical to ResolvedEQ.
func Resolved(v bool) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldResolved, v))
}

// Priority applies equality check predicate on the "priority" field. It's identical to PriorityEQ.
func Priority(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldPriority, v))
}

// ResolutionTime applies equality check predicate on the "resolution_time" field. It's identical to ResolutionTimeEQ.
func ResolutionTime(v float64) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldResolutionTime, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldStatus, v))
}

// Summary applies equality check predicate on the "summary" field. It's identical to SummaryEQ.
func Summary(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldSummary, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldURL, v))
}

// JiraKeyEQ applies the EQ predicate on the "jira_key" field.
func JiraKeyEQ(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldJiraKey, v))
}

// JiraKeyNEQ applies the NEQ predicate on the "jira_key" field.
func JiraKeyNEQ(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldNEQ(FieldJiraKey, v))
}

// JiraKeyIn applies the In predicate on the "jira_key" field.
func JiraKeyIn(vs ...string) predicate.Bugs {
	return predicate.Bugs(sql.FieldIn(FieldJiraKey, vs...))
}

// JiraKeyNotIn applies the NotIn predicate on the "jira_key" field.
func JiraKeyNotIn(vs ...string) predicate.Bugs {
	return predicate.Bugs(sql.FieldNotIn(FieldJiraKey, vs...))
}

// JiraKeyGT applies the GT predicate on the "jira_key" field.
func JiraKeyGT(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldGT(FieldJiraKey, v))
}

// JiraKeyGTE applies the GTE predicate on the "jira_key" field.
func JiraKeyGTE(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldGTE(FieldJiraKey, v))
}

// JiraKeyLT applies the LT predicate on the "jira_key" field.
func JiraKeyLT(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldLT(FieldJiraKey, v))
}

// JiraKeyLTE applies the LTE predicate on the "jira_key" field.
func JiraKeyLTE(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldLTE(FieldJiraKey, v))
}

// JiraKeyContains applies the Contains predicate on the "jira_key" field.
func JiraKeyContains(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldContains(FieldJiraKey, v))
}

// JiraKeyHasPrefix applies the HasPrefix predicate on the "jira_key" field.
func JiraKeyHasPrefix(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldHasPrefix(FieldJiraKey, v))
}

// JiraKeyHasSuffix applies the HasSuffix predicate on the "jira_key" field.
func JiraKeyHasSuffix(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldHasSuffix(FieldJiraKey, v))
}

// JiraKeyEqualFold applies the EqualFold predicate on the "jira_key" field.
func JiraKeyEqualFold(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldEqualFold(FieldJiraKey, v))
}

// JiraKeyContainsFold applies the ContainsFold predicate on the "jira_key" field.
func JiraKeyContainsFold(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldContainsFold(FieldJiraKey, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldLTE(FieldUpdatedAt, v))
}

// ResolvedAtEQ applies the EQ predicate on the "resolved_at" field.
func ResolvedAtEQ(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldResolvedAt, v))
}

// ResolvedAtNEQ applies the NEQ predicate on the "resolved_at" field.
func ResolvedAtNEQ(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldNEQ(FieldResolvedAt, v))
}

// ResolvedAtIn applies the In predicate on the "resolved_at" field.
func ResolvedAtIn(vs ...time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldIn(FieldResolvedAt, vs...))
}

// ResolvedAtNotIn applies the NotIn predicate on the "resolved_at" field.
func ResolvedAtNotIn(vs ...time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldNotIn(FieldResolvedAt, vs...))
}

// ResolvedAtGT applies the GT predicate on the "resolved_at" field.
func ResolvedAtGT(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldGT(FieldResolvedAt, v))
}

// ResolvedAtGTE applies the GTE predicate on the "resolved_at" field.
func ResolvedAtGTE(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldGTE(FieldResolvedAt, v))
}

// ResolvedAtLT applies the LT predicate on the "resolved_at" field.
func ResolvedAtLT(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldLT(FieldResolvedAt, v))
}

// ResolvedAtLTE applies the LTE predicate on the "resolved_at" field.
func ResolvedAtLTE(v time.Time) predicate.Bugs {
	return predicate.Bugs(sql.FieldLTE(FieldResolvedAt, v))
}

// ResolvedEQ applies the EQ predicate on the "resolved" field.
func ResolvedEQ(v bool) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldResolved, v))
}

// ResolvedNEQ applies the NEQ predicate on the "resolved" field.
func ResolvedNEQ(v bool) predicate.Bugs {
	return predicate.Bugs(sql.FieldNEQ(FieldResolved, v))
}

// PriorityEQ applies the EQ predicate on the "priority" field.
func PriorityEQ(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldPriority, v))
}

// PriorityNEQ applies the NEQ predicate on the "priority" field.
func PriorityNEQ(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldNEQ(FieldPriority, v))
}

// PriorityIn applies the In predicate on the "priority" field.
func PriorityIn(vs ...string) predicate.Bugs {
	return predicate.Bugs(sql.FieldIn(FieldPriority, vs...))
}

// PriorityNotIn applies the NotIn predicate on the "priority" field.
func PriorityNotIn(vs ...string) predicate.Bugs {
	return predicate.Bugs(sql.FieldNotIn(FieldPriority, vs...))
}

// PriorityGT applies the GT predicate on the "priority" field.
func PriorityGT(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldGT(FieldPriority, v))
}

// PriorityGTE applies the GTE predicate on the "priority" field.
func PriorityGTE(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldGTE(FieldPriority, v))
}

// PriorityLT applies the LT predicate on the "priority" field.
func PriorityLT(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldLT(FieldPriority, v))
}

// PriorityLTE applies the LTE predicate on the "priority" field.
func PriorityLTE(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldLTE(FieldPriority, v))
}

// PriorityContains applies the Contains predicate on the "priority" field.
func PriorityContains(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldContains(FieldPriority, v))
}

// PriorityHasPrefix applies the HasPrefix predicate on the "priority" field.
func PriorityHasPrefix(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldHasPrefix(FieldPriority, v))
}

// PriorityHasSuffix applies the HasSuffix predicate on the "priority" field.
func PriorityHasSuffix(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldHasSuffix(FieldPriority, v))
}

// PriorityEqualFold applies the EqualFold predicate on the "priority" field.
func PriorityEqualFold(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldEqualFold(FieldPriority, v))
}

// PriorityContainsFold applies the ContainsFold predicate on the "priority" field.
func PriorityContainsFold(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldContainsFold(FieldPriority, v))
}

// ResolutionTimeEQ applies the EQ predicate on the "resolution_time" field.
func ResolutionTimeEQ(v float64) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldResolutionTime, v))
}

// ResolutionTimeNEQ applies the NEQ predicate on the "resolution_time" field.
func ResolutionTimeNEQ(v float64) predicate.Bugs {
	return predicate.Bugs(sql.FieldNEQ(FieldResolutionTime, v))
}

// ResolutionTimeIn applies the In predicate on the "resolution_time" field.
func ResolutionTimeIn(vs ...float64) predicate.Bugs {
	return predicate.Bugs(sql.FieldIn(FieldResolutionTime, vs...))
}

// ResolutionTimeNotIn applies the NotIn predicate on the "resolution_time" field.
func ResolutionTimeNotIn(vs ...float64) predicate.Bugs {
	return predicate.Bugs(sql.FieldNotIn(FieldResolutionTime, vs...))
}

// ResolutionTimeGT applies the GT predicate on the "resolution_time" field.
func ResolutionTimeGT(v float64) predicate.Bugs {
	return predicate.Bugs(sql.FieldGT(FieldResolutionTime, v))
}

// ResolutionTimeGTE applies the GTE predicate on the "resolution_time" field.
func ResolutionTimeGTE(v float64) predicate.Bugs {
	return predicate.Bugs(sql.FieldGTE(FieldResolutionTime, v))
}

// ResolutionTimeLT applies the LT predicate on the "resolution_time" field.
func ResolutionTimeLT(v float64) predicate.Bugs {
	return predicate.Bugs(sql.FieldLT(FieldResolutionTime, v))
}

// ResolutionTimeLTE applies the LTE predicate on the "resolution_time" field.
func ResolutionTimeLTE(v float64) predicate.Bugs {
	return predicate.Bugs(sql.FieldLTE(FieldResolutionTime, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Bugs {
	return predicate.Bugs(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Bugs {
	return predicate.Bugs(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldContainsFold(FieldStatus, v))
}

// SummaryEQ applies the EQ predicate on the "summary" field.
func SummaryEQ(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldSummary, v))
}

// SummaryNEQ applies the NEQ predicate on the "summary" field.
func SummaryNEQ(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldNEQ(FieldSummary, v))
}

// SummaryIn applies the In predicate on the "summary" field.
func SummaryIn(vs ...string) predicate.Bugs {
	return predicate.Bugs(sql.FieldIn(FieldSummary, vs...))
}

// SummaryNotIn applies the NotIn predicate on the "summary" field.
func SummaryNotIn(vs ...string) predicate.Bugs {
	return predicate.Bugs(sql.FieldNotIn(FieldSummary, vs...))
}

// SummaryGT applies the GT predicate on the "summary" field.
func SummaryGT(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldGT(FieldSummary, v))
}

// SummaryGTE applies the GTE predicate on the "summary" field.
func SummaryGTE(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldGTE(FieldSummary, v))
}

// SummaryLT applies the LT predicate on the "summary" field.
func SummaryLT(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldLT(FieldSummary, v))
}

// SummaryLTE applies the LTE predicate on the "summary" field.
func SummaryLTE(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldLTE(FieldSummary, v))
}

// SummaryContains applies the Contains predicate on the "summary" field.
func SummaryContains(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldContains(FieldSummary, v))
}

// SummaryHasPrefix applies the HasPrefix predicate on the "summary" field.
func SummaryHasPrefix(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldHasPrefix(FieldSummary, v))
}

// SummaryHasSuffix applies the HasSuffix predicate on the "summary" field.
func SummaryHasSuffix(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldHasSuffix(FieldSummary, v))
}

// SummaryEqualFold applies the EqualFold predicate on the "summary" field.
func SummaryEqualFold(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldEqualFold(FieldSummary, v))
}

// SummaryContainsFold applies the ContainsFold predicate on the "summary" field.
func SummaryContainsFold(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldContainsFold(FieldSummary, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Bugs {
	return predicate.Bugs(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Bugs {
	return predicate.Bugs(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Bugs {
	return predicate.Bugs(sql.FieldContainsFold(FieldURL, v))
}

// HasBugs applies the HasEdge predicate on the "bugs" edge.
func HasBugs() predicate.Bugs {
	return predicate.Bugs(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BugsTable, BugsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBugsWith applies the HasEdge predicate on the "bugs" edge with a given conditions (other predicates).
func HasBugsWith(preds ...predicate.Teams) predicate.Bugs {
	return predicate.Bugs(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BugsInverseTable, TeamsFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BugsTable, BugsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Bugs) predicate.Bugs {
	return predicate.Bugs(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Bugs) predicate.Bugs {
	return predicate.Bugs(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Bugs) predicate.Bugs {
	return predicate.Bugs(func(s *sql.Selector) {
		p(s.Not())
	})
}
