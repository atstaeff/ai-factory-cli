package factory

import (
 "bytes"
 "strings"
 "testing"
)

func TestDemoScenarios(t *testing.T) {
 for _,name:=range []string{"running","approvals","incidents","multi"}{
  text:=renderDemo(name,demoScenarios[name],0)
  for _,required:=range []string{"DEMO MODE - SIMULATED DATA","Actual jobs: 0",demoScenarios[name].Title}{
   if !strings.Contains(text,required){t.Errorf("%s missing %q",name,required)}
  }
 }
}

func TestFiniteDemo(t *testing.T) {
 var out bytes.Buffer
 err:=Execute([]string{"demo","--scenario","running","--frames","1"},&out,&bytes.Buffer{})
 if err!=nil{t.Fatal(err)}
 if !strings.Contains(out.String(),"SIMULATED DATA"){t.Fatal("missing demo disclaimer")}
}

func TestUnknownScenarioRejected(t *testing.T) {
 if err:=Execute([]string{"demo","--scenario","unknown","--frames","1"},&bytes.Buffer{},&bytes.Buffer{});err==nil{t.Fatal("expected error")}
}
