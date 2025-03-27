package types

type StatDomainReq_Stats struct {
	Domain  string      `json:"domain"`
	Stats   DomainState `json:"stats"`
	Message string      `json:"message,omitempty"`
}

type StatDomainReq struct {
	IP    string                 `json:"ip"`
	Stats []*StatDomainReq_Stats `json:"stats"`
}
