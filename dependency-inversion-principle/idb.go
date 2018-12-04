package dependency

type IDb interface {
	Ping() string
}
