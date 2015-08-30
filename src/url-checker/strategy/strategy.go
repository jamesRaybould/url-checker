package strategy

type UrlStrategy interface {
	Get() ([]string, error)
}
