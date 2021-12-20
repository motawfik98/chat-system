package elasticdomain

import "chat-system/domain"

type match map[string]interface{}
type query map[string]match
type Search map[string]query

func CreateQuery(message string) Search {
	return Search{
		"query": query{
			"match": match{
				"message": message,
			},
		},
	}
}

type Hit struct {
	Index  string         `json:"_index"`
	Type   string         `json:"_type"`
	ID     string         `json:"_id"`
	Score  float64        `json:"_score"`
	Source domain.Message `json:"_source"`
}

type Shards struct {
	Total      float64 `json:"total"`
	Successful float64 `json:"successful"`
	Skipped    float64 `json:"skipped"`
	Failed     float64 `json:"failed"`
}

type TotalHits struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type SuccessHits struct {
	Total    TotalHits `json:"total"`
	MaxScore float64   `json:"max_score"`
	Hits     []Hit     `json:"hits"`
}

type SuccessResponse struct {
	Took     float64     `json:"took"`
	TimedOut bool        `json:"timed_out"`
	Shards   Shards      `json:"_shards"`
	Hits     SuccessHits `json:"hits"`
}

func (r *SuccessResponse) GetMessages(appToken string, chatNumber uint) []domain.Message {
	messages := make([]domain.Message, r.Hits.Total.Value)
	for i := 0; i < r.Hits.Total.Value; i++ {
		messages[i] = r.Hits.Hits[i].Source
		messages[i].AppToken = appToken
		messages[i].ChatNumber = chatNumber
	}
	return messages
}
