package factory

import(
 "bytes"
 "os"
 "path/filepath"
 "strings"
 "testing"
)

func TestPilotDryRunNeverCallsExternalCommands(t *testing.T){
 root:=t.TempDir()
 var out bytes.Buffer
 if err:=Execute([]string{"pilot","--project",root,"--repo","org/repo","--issue","5"},&out,&bytes.Buffer{});err!=nil{t.Fatal(err)}
 if !strings.Contains(out.String(),"DRY RUN"){t.Fatal(out.String())}
 if files,_:=os.ReadDir(root);len(files)!=0{t.Fatalf("modified project: %v",files)}
}

func TestPilotWithMockGitHubAndOpenCode(t *testing.T){
 root:=t.TempDir()
 bin:=t.TempDir()
 issueLog:=filepath.Join(root,"issue.args")
 promptLog:=filepath.Join(root,"prompt.args")
 ghScript:="#!/bin/sh\nprintf '%s' \"$*\" > "+issueLog+"\nprintf '%s\\n' '{\"number\":7,\"title\":\"Test\",\"body\":\"Please assess\",\"url\":\"https://github.com/example/repo/issues/7\",\"state\":\"OPEN\"}'\n"
 ocScript:="#!/bin/sh\nprintf '%s' \"$*\" > "+promptLog+"\nprintf 'mock opencode response\\n'\n"
 for name,script:=range map[string]string{"gh":ghScript,"opencode":ocScript}{
  if err:=os.WriteFile(filepath.Join(bin,name),[]byte(script),0700);err!=nil{t.Fatal(err)}
 }
 t.Setenv("PATH",bin+string(os.PathListSeparator)+os.Getenv("PATH"))
 var out bytes.Buffer
 err:=Execute([]string{"pilot","--project",root,"--repo","example/repo","--issue","7","--execute"},&out,&bytes.Buffer{})
 if err!=nil{t.Fatal(err)}
 if !strings.Contains(out.String(),"mock opencode response"){t.Fatal(out.String())}
 b,_:=os.ReadFile(issueLog)
 if !strings.Contains(string(b),"issue view 7 -R example/repo"){t.Fatal(string(b))}
 b,_=os.ReadFile(promptLog)
 if !strings.Contains(string(b),"--agent plan")||!strings.Contains(string(b),"Please assess"){t.Fatal(string(b))}
}

func TestPilotRejectsBadInputs(t *testing.T){
 for _,args:=range [][]string{
  {"pilot","--repo","../../bad","--issue","1"},
  {"pilot","--repo","a/b","--issue","0"},
  {"pilot","--repo","a/b","--issue","1","--agent","-auto"},
 }{
  if err:=Execute(args,&bytes.Buffer{},&bytes.Buffer{});err==nil{t.Fatalf("expected error: %v",args)}
 }
}
