package types

type Server struct {
	Version      string
	VersionMajor uint
	VersionMinor uint
	StartTime    JSONTime
	CurrentTime  JSONTime
	BuildNumber  string
	BuildDate    JSONTime
}
