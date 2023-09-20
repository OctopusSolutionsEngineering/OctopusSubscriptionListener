package events

import "time"

type SubscriptionEvent struct {
	Timestamp time.Time `json:"Timestamp"`
	EventType string    `json:"EventType"`
	Payload   struct {
		ServerUri      string `json:"ServerUri"`
		ServerAuditUri string `json:"ServerAuditUri"`
		Subscription   struct {
		} `json:"Subscription"`
		Event struct {
			Category          string `json:"Category"`
			Message           string `json:"Message"`
			MessageReferences []struct {
				ReferencedDocumentId string `json:"ReferencedDocumentId"`
			} `json:"MessageReferences"`
			RelatedDocumentIds []string `json:"RelatedDocumentIds"`
			SpaceId            string   `json:"SpaceId"`
		} `json:"Event"`
		BatchProcessingDate time.Time `json:"BatchProcessingDate"`
		BatchId             string    `json:"BatchId"`
		TotalEventsInBatch  string    `json:"TotalEventsInBatch"`
		EventNumberInBatch  string    `json:"EventNumberInBatch"`
	} `json:"Payload"`
}
