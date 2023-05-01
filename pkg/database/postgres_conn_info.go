package database

type FlagSet interface {
	StringVar(p *string, name, value, desc string)
}

var (
	ModeDisable string = "disable"
)

type PostgresConnInfo struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func Export(flag FlagSet, defaultConnInfo PostgresConnInfo) *PostgresConnInfo {
	var info PostgresConnInfo

	flag.StringVar(&info.Host,     "host",     defaultConnInfo.Host,     "database host")
	flag.StringVar(&info.Port,     "port",     defaultConnInfo.Port,     "database port")
	flag.StringVar(&info.User,     "user",     defaultConnInfo.User,     "database user")
	flag.StringVar(&info.Password, "password", defaultConnInfo.Password, "database password")
	flag.StringVar(&info.Name,     "name",     defaultConnInfo.Name,     "database name")
	flag.StringVar(&info.SSLMode,  "SSLMode",  defaultConnInfo.SSLMode,  "database SSLMode")

	return &info
}
