package summarize

//go:generate mockgen -destination=mock_summarize.go -package=summarize . Summarize

type Summarize interface {
	Summarize(a, b int) int
}
