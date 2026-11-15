package factory

import (
 "fmt"
 "io"
 "os"
 "os/exec"
 "strings"
)

// checkTools never launches third-party executables. Presence is not compatibility.
func checkTools(out io.Writer) error {
 fmt.Fprintln(out,"Factory tooling diagnostics (availability only; no commands executed)")
 for _,name:=range []string{"git","gh","opencode","claude","bmad","openspec"} {
  path,err:=exec.LookPath(name)
  if err!=nil {fmt.Fprintf(out,"%-14s missing\n",name);continue}
  fmt.Fprintf(out,"%-14s found    %s\n",name,path)
 }
 custom:=strings.TrimSpace(os.Getenv("FACTORY_ISSUE_CLI"))
 if custom=="" {
  fmt.Fprintln(out,"issue-cli      not configured (set FACTORY_ISSUE_CLI to an approved executable)")
 } else {
  path,err:=exec.LookPath(custom)
  if err!=nil {fmt.Fprintln(out,"issue-cli      configured, not found");return nil}
  fmt.Fprintf(out,"issue-cli      found    %s\n",path)
 }
 fmt.Fprintln(out,"No integrations, model access, permissions, version constraints or process semantics have been validated.")
 return nil
}
