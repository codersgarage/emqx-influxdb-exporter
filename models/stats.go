package models

type Stats struct {
	CName              string
	Node               string `json:"node"`
	SubscriptionsCount int64  `json:"subscriptions/count"`
	SubscriptionsMax   int64  `json:"subscriptions/max"`
	TopicsCount        int64  `json:"topics/count"`
	TopicsMax          int64  `json:"topics/max"`
	RoutesCount        int64  `json:"routes/count"`
	RoutesMax          int64  `json:"routes/max"`
	SubscribersCount   int64  `json:"subscribers/count"`
	SubscribersMax     int64  `json:"subscribers/max"`
	ConnectionsCount   int64  `json:"connections/count"`
	ConnectionsMax     int64  `json:"connections/max"`
}

func (s *Stats) CollectionName() string {
	return s.CName
}

func (s *Stats) SetCollectionName(name string) {
	s.CName = name
}
