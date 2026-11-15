package factory

import (
 "bytes"
 "strings"
 "testing"
)
func TestToolsReadOnly(t *testing.T) {
 t.Setenv("FACTORY_ISSUE_CLI","")
 var out bytes.Buffer
 if err:=Execute([]string{"tools"},&out,&bytes.Buffer{});err!=nil{t.Fatal(err)}
 if !strings.Contains(out.String(),"issue-cli") || !strings.Contains(out.String(),"availability only"){t.Fatal(out.String())}
}
func TestVersionVariable(t *testing.T) {
 var out bytes.Buffer
 if err:=Execute([]string{"version"},&out,&bytes.Buffer{});err!=nil{t.Fatal(err)}
 if !strings.Contains(out.String(),"factory "+Version){t.Fatal(out.String())}
}
