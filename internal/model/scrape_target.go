package model

// ScrapeTarget representa um endpoint para scraping.
type ScrapeTarget struct {
	TargetID    string                 `json:"targetId"`
	Method      string                 `json:"method"`
	Description string                 `json:"description,omitempty"`
	URL         string                 `json:"url"`
	Domain      string                 `json:"domain"`
	Version     string                 `json:"version"`
	BasePath    string                 `json:"base_path"`
	Paths       []Path                 `json:"paths,omitemptyelem"`
	QueryParams []QueryParam           `json:"query_params,omitemptyelem"`
	Headers     []Header               `json:"headers"`
	PayloadBody map[string]interface{} `json:"payload_body,omitempty"` // A estrutura pode variar
}

type Path struct {
	Path           string `json:"path"`
	ValuePathParam string `json:"value_path_param,omitempty"`
}

type QueryParam struct {
	Param string `json:"param"`
	Value string `json:"value"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
