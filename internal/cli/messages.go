package cli

var Stderr = struct {
	NoAuthFile         string
	ResponseNotWritten string
	PageBytesWritten   string
}{
	NoAuthFile:         "no auth.json file present, will not be able to access the Battle.Net APIs without it",
	ResponseNotWritten: "could not write response body, %s",
	PageBytesWritten:   "number of bytes written %s",
}
