package main

import "github.com/urfave/cli/v2"

var CliFlags = []cli.Flag{
	CliFlagAuthUsernamePassword,
	CliFlagWindowSize,
	CliFlagLocalAddress,
	CliFlagConcurrentRequestNumber,
	CliFlagCookieNameValue,
	CliFlagDontDisplayPercentageTable,
	CliFlagCSVFile,
	CliFlagProtocol,
	CliFlagGnuplotFile,
	// CliFlagDisplayUsageInformation,
	CliFlagCustomHeader,
	CliFlagDoHEADInsteadOfGET,
	CliFlagEnableHTTPKeepAliveFeature,
	CliFlagDoNotReportResponseLengthErrors,
	CliFlagHTTPMethod,
	CliFlagRequestNumber,
	CliFlagPOSTFile,
	CliFlagProxyAuthUsernamePassword,
	CliFlagDoNotDisplayProgressMessages,
	CliFlagDoNotExitOnSocketReceiveErrors,
	CliFlagTimeoutSeconds,
	CliFlagDoNotDisplayMedianAndStdDevValues,
	CliFlagTimelimitSeconds,
	CliFlagContentType,
	CliFlagPUTFile,
	CliFlagVerbosityLevel,
	CliFlagDisplayVersionNumberAndExit,
	CliFlagPrintHTMLTables,
	CliFlagTableAttributes,
	CliFlagProxyAddress,
	CliFlagTRAttributes,
	CliFlagTDAttributes,
	CliFlagCipherSuite,
}

var CliFlagAuthUsernamePassword = &cli.StringFlag{
	Name:  "A",
	Value: "",
	Usage: "auth-username:password",
}

var CliFlagWindowSize = &cli.IntFlag{
	Name:  "b",
	Value: 0,
	Usage: "windowsize",
}

var CliFlagLocalAddress = &cli.StringFlag{
	Name:  "B",
	Value: "",
	Usage: "local-address",
}

var CliFlagConcurrentRequestNumber = &cli.IntFlag{
	Name:  "c",
	Value: 1,
	Usage: "number of concurrent requests",
}

var CliFlagCookieNameValue = &cli.StringFlag{
	Name:  "C",
	Value: "",
	Usage: "cookie-name=value",
}

var CliFlagDontDisplayPercentageTable = &cli.BoolFlag{
	Name:  "d",
	Usage: "do not display the percentage table",
	Value: false,
}

var CliFlagCSVFile = &cli.StringFlag{
	Name:  "e",
	Usage: "csv-file",
	Value: "",
}

var CliFlagProtocol = &cli.StringFlag{
	Name:  "f",
	Usage: "protocol",
	Value: "",
}

var CliFlagGnuplotFile = &cli.StringFlag{
	Name:  "g",
	Usage: "gnuplot-file",
	Value: "",
}

/*var CliFlagDisplayUsageInformation = &cli.BoolFlag{
	Name:  "h",
	Usage: "display usage information",
	Value: false,
}*/

var CliFlagCustomHeader = &cli.StringSliceFlag{
	Name:  "H",
	Value: &cli.StringSlice{},
	Usage: "custom-header",
}

var CliFlagDoHEADInsteadOfGET = &cli.BoolFlag{
	Name:  "i",
	Usage: "do HEAD requests instead of GET",
	Value: false,
}

var CliFlagEnableHTTPKeepAliveFeature = &cli.BoolFlag{
	Name:  "k",
	Usage: "enable HTTP KeepAlive feature",
	Value: false,
}

var CliFlagDoNotReportResponseLengthErrors = &cli.BoolFlag{
	Name:  "l",
	Usage: "do not report errors if the length of the responses is not constant",
	Value: false,
}

var CliFlagHTTPMethod = &cli.StringFlag{
	Name:  "m",
	Value: "GET",
	Usage: "HTTP method",
}

var CliFlagRequestNumber = &cli.IntFlag{
	Name:  "n",
	Value: 1,
	Usage: "number of requests",
}

var CliFlagPOSTFile = &cli.StringFlag{
	Name:  "p",
	Value: "",
	Usage: "POST-file",
}

var CliFlagProxyAuthUsernamePassword = &cli.StringFlag{
	Name:  "P",
	Value: "",
	Usage: "proxy-auth-username:password",
}

var CliFlagDoNotDisplayProgressMessages = &cli.BoolFlag{
	Name:  "q",
	Usage: "do not display progress messages",
	Value: false,
}

var CliFlagDoNotExitOnSocketReceiveErrors = &cli.BoolFlag{
	Name:  "r",
	Usage: "don't exit on socket receive errors",
	Value: false,
}

var CliFlagTimeoutSeconds = &cli.IntFlag{
	Name:  "s",
	Value: 30,
	Usage: "timeout in seconds",
}

var CliFlagDoNotDisplayMedianAndStdDevValues = &cli.BoolFlag{
	Name:  "S",
	Usage: "do not display the median and standard deviation values",
	Value: false,
}

var CliFlagTimelimitSeconds = &cli.IntFlag{
	Name:  "t",
	Value: 0,
	Usage: "timelimit in seconds",
}

var CliFlagContentType = &cli.StringFlag{
	Name:  "T",
	Value: "",
	Usage: "content-type",
}

var CliFlagPUTFile = &cli.StringFlag{
	Name:  "u",
	Value: "",
	Usage: "PUT-file",
}

var CliFlagVerbosityLevel = &cli.IntFlag{
	Name:  "v",
	Value: 0,
	Usage: "verbosity level",
}

var CliFlagDisplayVersionNumberAndExit = &cli.BoolFlag{
	Name:  "V",
	Usage: "display version number and exit",
	Value: false,
}

var CliFlagPrintHTMLTables = &cli.BoolFlag{
	Name:  "w",
	Usage: "print out results in HTML tables",
	Value: false,
}

var CliFlagTableAttributes = &cli.StringFlag{
	Name:  "x",
	Value: "",
	Usage: "<table>-attributes",
}

var CliFlagProxyAddress = &cli.StringFlag{
	Name:  "X",
	Value: "",
	Usage: "proxy[:port]",
}

var CliFlagTRAttributes = &cli.StringFlag{
	Name:  "y",
	Value: "",
	Usage: "<tr>-attributes",
}

var CliFlagTDAttributes = &cli.StringFlag{
	Name:  "z",
	Value: "",
	Usage: "<td>-attributes",
}

var CliFlagCipherSuite = &cli.StringFlag{
	Name:  "Z",
	Value: "",
	Usage: "ciphersuite",
}
