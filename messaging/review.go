package messaging

const (
	ContentTypeReview = "application/x-review"
)

// Review - maybe add Task??
type Review struct {
	Duration string
}

func NewReviewMessage(review *Review) *Message {
	m := NewMessage(ChannelControl, ConfigEvent)
	m.SetContent(ContentTypeReview, review)
	return m
}

func ReviewContent(m *Message) *Review {
	if m.Name() != ConfigEvent || m.ContentType() != ContentTypeReview {
		return nil
	}
	if v, ok := m.Body.(*Review); ok {
		return v
	}
	return nil
}
