package models

type Metrics struct {
	CName   string
	Node    string `json:"node"`
	Metrics Vol    `json:"metrics"`
}

type Vol struct {
	BytesReceived          int64 `json:"bytes/received"`
	BytesSent              int64 `json:"bytes/sent"`
	MessagesExpired        int64 `json:"messages/expired"`
	PacketsPublishReceived int64 `json:"packets/publish/received"`
	PacketsPublishSent     int64 `json:"packets/publish/sent"`
	MessagesQos1Received   int64 `json:"messages/qos1/received"`
	MessagesQos1Sent       int64 `json:"messages/qos1/sent"`
	MessagesSent           int64 `json:"messages/sent"`
	MessagesReceived       int64 `json:"messages/received"`
	MessagesDropped        int64 `json:"messages/dropped"`
	MessagesForward        int64 `json:"messages/forward"`
}

func (m *Metrics) CollectionName() string {
	return m.CName
}

func (m *Metrics) SetCollectionName(name string) {
	m.CName = name
}
