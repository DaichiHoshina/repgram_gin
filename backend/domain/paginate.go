package domain

type Paginate struct {
	Page int `json:"page"`
	Per  int `json:"per"`
}
