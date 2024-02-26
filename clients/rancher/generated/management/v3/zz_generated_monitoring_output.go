package client

const (
	MonitoringOutputType                  = "monitoringOutput"
	MonitoringOutputFieldAnswers          = "answers"
	MonitoringOutputFieldAnswersSetString = "answersSetString"
	MonitoringOutputFieldExtraAnswers     = "extraAnswers"
	MonitoringOutputFieldValuesYaml       = "valuesYaml"
	MonitoringOutputFieldVersion          = "version"
)

type MonitoringOutput struct {
	Answers          map[string]string `json:"answers,omitempty" yaml:"answers,omitempty"`
	AnswersSetString map[string]string `json:"answersSetString,omitempty" yaml:"answersSetString,omitempty"`
	ExtraAnswers     map[string]string `json:"extraAnswers,omitempty" yaml:"extraAnswers,omitempty"`
	ValuesYaml       string            `json:"valuesYaml,omitempty" yaml:"valuesYaml,omitempty"`
	Version          string            `json:"version,omitempty" yaml:"version,omitempty"`
}
