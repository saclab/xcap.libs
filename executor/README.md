# Package Executor

## Run:
Run : Run a function and return the std output and error
Accepts two inputs, the fully qualified path of a binary (string) and concatinated list of flags (string)
Executes the binary with the given flags and
returns stdOut (string) and stdErr (string)
func Run( string,  string) (string, string)

## Shellout
Shellout : Run and stream the output to std display, also retuen the std out and err
Accepts two inputs, the fully qualified path of a binary (string) and concatinated list of flags (string)
Executes the binary with the given flags and
returns stdOut (string) and stdErr (string)
func Shellout( string,  string) (string, string)
