package data

type State struct {
	Name     string `json:"name"`
	Capital  string `json:"capital"`
	Language string `json:"language"`
}
type Info struct {
	Name   string `json:"name"`
	Schema string `json:"schema"`
	Files  []Form `json:"files"`
}
type Form struct {
	Name     string `json:"name"`
	Contents string `json:"contents"`
}
