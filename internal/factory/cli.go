package factory

import (
 "errors"
 "fmt"
 "io"
 "os"
 "path/filepath"
 "sort"
 "strings"
)

const usage = `factory - process-driven AI work orchestration (early prototype)

Usage:
  factory                    Show help
  factory scan [path]        Inspect a local project without changing it
  factory doctor [path]      Check required project configuration
  factory processes [path]   List discovered PROCESS.md definitions
  factory status [path]      Show project setup status
  factory version            Print version
  factory help               Show help

Execution, scheduling and integrations are NOT implemented yet.
`

type Finding struct {
 Name string
 Present bool
 Path string
}

func Execute(args []string, out, errOut io.Writer) error {
 _ = errOut
 if len(args)==0 { _,e:=io.WriteString(out,usage); return e }
 command:=args[0]
 if command=="help" || command=="--help" || command=="-h" { _,e:=io.WriteString(out,usage); return e }
 if command=="version" { _,e:=fmt.Fprintln(out,"factory 0.1.0-dev"); return e }
 if command!="scan" && command!="doctor" && command!="processes" && command!="status" {
  return fmt.Errorf("unknown or not yet implemented command %q (run factory help)",command)
 }
 if len(args)>2 { return errors.New("expected at most one project path") }
 root:="."
 if len(args)==2 {root=args[1]}
 abs,e:=filepath.Abs(root);if e!=nil{return e}
 info,e:=os.Stat(abs);if e!=nil{return e};if !info.IsDir(){return errors.New("project path must be a directory")}
 findings:=inspect(abs)
 switch command {
 case "scan":
  fmt.Fprintf(out,"Project: %s\n",abs)
  for _,f:=range findings { status:="missing";if f.Present{status="found"};fmt.Fprintf(out,"%-16s %-8s %s\n",f.Name,status,f.Path) }
  return nil
 case "doctor":
  fmt.Fprintf(out,"Project: %s\n",abs)
  has:=false
  for _,f:=range findings {if f.Name=="PROCESS.md" && f.Present {has=true}}
  if !has {fmt.Fprintln(out,"CHECK: no PROCESS.md found; run factory scan to inspect the project");return errors.New("process configuration not found")}
  fmt.Fprintln(out,"CHECK: process found; runtime and provider validation not implemented")
  return nil
 case "status":
  fmt.Fprintf(out,"Project: %s\n",abs)
  fmt.Fprintln(out,"Supervisor: not running (execution engine not implemented)")
  return nil
 case "processes":
  paths,e:=discoverProcesses(abs);if e!=nil{return e}
  if len(paths)==0 {fmt.Fprintln(out,"No PROCESS.md files found");return nil}
  for _,p:=range paths {name:=processName(p);fmt.Fprintf(out,"%s\t%s\n",name,p)}
  return nil
 }
 return nil
}

func inspect(root string) []Finding {
 items:=[]struct{name,path string}{
  {"PROCESS.md","PROCESS.md"},
  {"BMAD","_bmad"},
  {"BMAD skills",".agents/skills"},
  {"OpenSpec","openspec"},
  {"AGENTS.md","AGENTS.md"},
  {"Factory config",".factory"},
 }
 out:=make([]Finding,0,len(items))
 for _,x:=range items {
  p:=filepath.Join(root,x.path)
  _,e:=os.Stat(p)
  out=append(out,Finding{Name:x.name,Present:e==nil,Path:x.path})
 }
 return out
}

func discoverProcesses(root string)([]string,error){
 var paths []string
 err:=filepath.WalkDir(root,func(path string,d os.DirEntry,e error)error{
  if e!=nil{return e}
  if d.IsDir(){
   if path!=root && (strings.HasPrefix(d.Name(),".git") || d.Name()=="node_modules" || d.Name()=="vendor" || d.Name()=="_bmad" || d.Name()=="_bmad-output"){return filepath.SkipDir}
   rel,_:=filepath.Rel(root,path)
   if len(strings.Split(rel,string(filepath.Separator)))>5{return filepath.SkipDir}
   return nil
  }
  if d.Name()=="PROCESS.md"{paths=append(paths,path)}
  return nil
 })
 sort.Strings(paths)
 return paths,err
}

func processName(path string)string{
 b,e:=os.ReadFile(path);if e!=nil{return "unknown"}
 data:=string(b)
 if !strings.HasPrefix(data,"---\n"){return filepath.Base(filepath.Dir(path))}
 lines:=strings.Split(data,"\n")
 for _,line:=range lines[1:] {
  if strings.TrimSpace(line)=="---"{break}
  if strings.HasPrefix(line,"name:"){return strings.Trim(strings.TrimSpace(strings.TrimPrefix(line,"name:")),"\"'")}
 }
 return filepath.Base(filepath.Dir(path))
}
