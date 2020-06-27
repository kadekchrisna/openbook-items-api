package queries

type EsQuery struct {
	Equal []FieldValue `json:"equal"`
}

type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}
