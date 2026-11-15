package factory

import (
 "bytes"
 "os"
 "path/filepath"
 "strings"
 "testing"
)

func TestHelp(t *testing.T){
 var out bytes.Buffer
 if err:=Execute(nil,&out,&bytes.Buffer{});err!=nil{t.Fatal(err)}
 if !strings.Contains(out.String(),"factory scan"){t.Fatal(out.String())}
}

func TestScanNoMutation(t *testing.T){
 dir:=t.TempDir()
 var out bytes.Buffer
 if err:=Execute([]string{"scan",dir},&out,&bytes.Buffer{});err!=nil{t.Fatal(err)}
 if !strings.Contains(out.String(),"BMAD"){t.Fatal(out.String())}
 entries,err:=os.ReadDir(dir);if err!=nil{t.Fatal(err)}
 if len(entries)!=0{t.Fatalf("scan wrote %d files",len(entries))}
}

func TestProcessDiscovery(t *testing.T){
 dir:=t.TempDir()
 p:=filepath.Join(dir,".factory","PROCESS.md")
 if err:=os.MkdirAll(filepath.Dir(p),0755);err!=nil{t.Fatal(err)}
 if err:=os.WriteFile(p,[]byte("---\nname: example\n---\n# Instructions\n"),0644);err!=nil{t.Fatal(err)}
 var out bytes.Buffer
 if err:=Execute([]string{"processes",dir},&out,&bytes.Buffer{});err!=nil{t.Fatal(err)}
 if !strings.Contains(out.String(),"example"){t.Fatal(out.String())}
 if err:=Execute([]string{"doctor",dir},&bytes.Buffer{},&bytes.Buffer{});err!=nil{t.Fatal(err)}
}

func TestUnknownCommand(t *testing.T){
 if err:=Execute([]string{"run"},&bytes.Buffer{},&bytes.Buffer{});err==nil{t.Fatal("unimplemented run should fail")}
}
