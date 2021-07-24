package database

type Tivi struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Technology  string   `json:"technology"`
	Resolution  string   `json:"recursion"`
	Type        string   `json:"type"`
	Imgs        []string `json:"imgs"`
	Description []string `json:"description"`
	Sizes       []string `json:"sizes"`
	Prices      []string `json:"prices"`
	LinkDetail  string   `json:"linkdetail"`
}
