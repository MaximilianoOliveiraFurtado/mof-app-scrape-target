package dynamodbmodel

// ScrapeTarget representa um endpoint para scraping.
type ScrapeTarget struct {
	TargetID    string                 `dynamodbav:"target_id"`
	Method      string                 `dynamodbav:"method"`
	Description string                 `dynamodbav:"description,omitempty"`
	URL         string                 `dynamodbav:"url"`
	Domain      string                 `dynamodbav:"domain"`
	Version     string                 `dynamodbav:"version"`
	BasePath    string                 `dynamodbav:"base_path"`
	Paths       []Path                 `dynamodbav:"paths,omitemptyelem"`
	QueryParams []QueryParam           `dynamodbav:"query_params,omitemptyelem"`
	Headers     []Header               `dynamodbav:"headers"`
	PayloadBody map[string]interface{} `dynamodbav:"payload_body,omitempty"`
}

type Path struct {
	Path           string `dynamodbav:"path"`
	ValuePathParam string `dynamodbav:"value_path_param,omitempty"`
}

type QueryParam struct {
	Param string `dynamodbav:"param"`
	Value string `dynamodbav:"value"`
}

type Header struct {
	Name  string `dynamodbav:"name"`
	Value string `dynamodbav:"value"`
}
