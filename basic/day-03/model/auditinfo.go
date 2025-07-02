package model

import "time"

// AuditInfo holds information about the audit of a transaction.
type AuditInfo struct {
	byUserID int       // ID of the user who performed the transaction
	byUser   string    // Name of the user who performed the transaction
	atTime   time.Time // Timestamp of when the transaction was performed
	notes    string    // Additional notes about the transaction
}

// NewAuditInfo creates a new AuditInfo instance with the current time.
func NewAuditInfo(byUserID int, byUser string, note string) *AuditInfo {
	return &AuditInfo{
		byUserID: byUserID,
		byUser:   byUser,
		atTime:   time.Now(),
		notes:    note,
	}
}

func (a *AuditInfo) Summary() string {
	return a.byUser + " performed a transaction at " + a.atTime.Format(time.RFC3339) + ". Notes: " + a.notes
}
