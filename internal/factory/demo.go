package factory

import (
 "flag"
 "fmt"
 "io"
 "strings"
 "time"
)

// The demo displays live-refreshing, entirely synthetic job telemetry.
// No third-party agent, issue provider or model is ever invoked.
type demoJob struct { Project, Issue, Method, Step, Runtime, State string }
type demoScenario struct { Title, Hint string; Jobs []demoJob; Logs, Details []string }

var demoScenarios = map[string]demoScenario{
 "running": {
  Title:"Active execution", Hint:"Simulated jobs at different engineering stages",
  Jobs:[]demoJob{
   {"document-api","DU-142","BMAD","implementation","OpenCode","RUNNING"},
   {"research-lab","RES-24","Custom","research","Claude CLI","RUNNING"},
   {"web-client","WEB-71","OpenSpec","review","OpenCode","QUEUED"}},
  Logs:[]string{"DU-142  BMAD specification complete (simulated)","RES-24  Research draft generated (simulated)","DU-142  Unit tests entered execution phase (simulated)","WEB-71  Waiting for simulated worker"},
  Details:[]string{"DU-142 steps: intake > spec > implementation > quality > review","       [DONE]    [DONE]   [RUNNING]       [PENDING]  [PENDING]","Runtime: mock OpenCode    Model: demo-model    Retries: 0"},
 },
 "approvals": {
  Title:"Human-in-the-loop", Hint:"Simulated pause on a required decision; no auto-approval",
  Jobs:[]demoJob{
   {"customer-api","API-81","BMAD","security approval","OpenCode","WAITING"},
   {"product-team","PM-22","Custom","scope review","Claude CLI","WAITING"},
   {"document-api","DU-145","BMAD","specification","OpenCode","RUNNING"}},
  Logs:[]string{"API-81  Human security approval required (simulated)","PM-22  Scope decision requested (simulated)","API-81  Agent suspended (simulated)","DU-145  Specification draft ready (simulated)"},
  Details:[]string{"API-81: approval required for privileged change","Owner: security-reviewer    Decision: PENDING","A normal issue comment must not count as approval."},
 },
 "incidents": {
  Title:"Failures, retries and stale jobs", Hint:"Simulated errors, backoff and heartbeat alerts",
  Jobs:[]demoJob{
   {"customer-api","API-91","Custom","integration","OpenCode","FAILED"},
   {"document-api","DU-149","BMAD","quality gates","OpenCode","RETRYING"},
   {"web-client","WEB-89","OpenSpec","implementation","Claude CLI","STALLED"}},
  Logs:[]string{"API-91  Mock enterprise CLI error 4","DU-149  Simulated quality gate failed (retry 2/3)","WEB-89  Simulated heartbeat timeout","API-91  No real credentials or systems accessed"},
  Details:[]string{"API-91 ERROR: adapter failure; no policy bypass","DU-149 Retry: 2/3   Backoff: 30s (simulated)","WEB-89 Heartbeat: STALE    Operator review recommended"},
 },
 "multi": {
  Title:"Multi-project orchestration", Hint:"BMAD / OpenSpec / Custom across separate demo projects",
  Jobs:[]demoJob{
   {"platform","PLAT-12","BMAD","design","OpenCode","RUNNING"},
   {"marketing","MKT-08","Custom","content QA","Claude CLI","RUNNING"},
   {"frontend","FE-34","OpenSpec","implementation","OpenCode","QUEUED"},
   {"customer-api","API-76","Custom","human review","Internal CLI","WAITING"},
   {"research","RES-33","Custom","discovery","Claude CLI","RUNNING"}},
  Logs:[]string{"PLAT-12  Simulated architecture checklist active","MKT-08  Content QA stage reached (simulated)","API-76  Mock issue approval waiting","RES-33  Mock research step 3/5 complete"},
  Details:[]string{"Workspace isolation: NOT ENFORCED by demo","Work sources: mock GitHub / mock enterprise CLI","No real coding CLI, BMAD or issue system executed"},
 },
}

// renderDemo returns an ANSI dashboard frame, independent of the terminal backend.
func renderDemo(name string, s demoScenario, tick int) string {
 var b strings.Builder
 rule:=strings.Repeat("=",100)
 fmt.Fprintf(&b,"\033[1;36m%s\033[0m\n",rule)
 fmt.Fprintln(&b," FACTORY / OBSERVABILITY                       \033[1;33mDEMO MODE - SIMULATED DATA\033[0m")
 fmt.Fprintf(&b," Scenario: %-16s %-40s\n",name,s.Title)
 fmt.Fprintf(&b," Uptime: %02d:%02d    Refresh: 1s    Work items: %d    Workers: SIMULATED    Tick: %04d\n",tick/60,tick%60,len(s.Jobs),tick)
 fmt.Fprintln(&b,rule)
 counts:=map[string]int{}
 for _,j:=range s.Jobs{counts[j.State]++}
 fmt.Fprintf(&b," RUNNING %d  |  QUEUED %d  |  WAITING %d  |  RETRYING %d  |  FAILED %d  |  STALLED %d\n",counts["RUNNING"],counts["QUEUED"],counts["WAITING"],counts["RETRYING"],counts["FAILED"],counts["STALLED"])
 fmt.Fprintf(&b,"\n %-18s %-11s %-12s %-22s %-14s %-11s\n","PROJECT","WORK ITEM","METHOD","PROCESS STEP","RUNTIME","STATUS")
 fmt.Fprintln(&b," "+strings.Repeat("-",96))
 for _,j:=range s.Jobs{
  color:="\033[32m"
  switch j.State{case "WAITING","QUEUED":color="\033[33m";case "FAILED","STALLED":color="\033[31m";case "RETRYING":color="\033[35m"}
  fmt.Fprintf(&b," %-18s %-11s %-12s %-22s %-14s %s%-11s\033[0m\n",j.Project,j.Issue,j.Method,j.Step,j.Runtime,color,j.State)
 }
 fmt.Fprintln(&b,"\n \033[1;36mLIVE EVENTS (SYNTHETIC)\033[0m")
 fmt.Fprintln(&b," "+strings.Repeat("-",96))
 for i:=0;i<4;i++{
  idx:=(tick/2+i)%len(s.Logs)
  fmt.Fprintf(&b," [%02d:%02d:%02d] INFO  %s\n",(tick/3600)%24,(tick/60)%60,tick%60,s.Logs[idx])
 }
 fmt.Fprintln(&b,"\n \033[1;36mSELECTED JOB / DIAGNOSTICS\033[0m")
 fmt.Fprintln(&b," "+strings.Repeat("-",96))
 for _,line:=range s.Details{fmt.Fprintln(&b," "+line)}
 fmt.Fprintln(&b,"\n \033[1;36mSYSTEM\033[0m  Source: DEMO GENERATOR   |  Actual jobs: 0   |  Issues synced: 0")
 fmt.Fprintln(&b," "+s.Hint)
 fmt.Fprintln(&b,"\n Other scenarios: factory demo --scenario running|approvals|incidents|multi")
 fmt.Fprintln(&b," Exit: Ctrl+C    |    The demo does NOT execute agents or query real systems.")
 return b.String()
}

func runDemo(args []string, out io.Writer) error {
 flags:=flag.NewFlagSet("demo",flag.ContinueOnError)
 flags.SetOutput(io.Discard)
 scenario:=flags.String("scenario","running","scenario name")
 frames:=flags.Int("frames",0,"number of frames (0 = until interrupted)")
 interval:=flags.Duration("interval",time.Second,"refresh interval")
 if err:=flags.Parse(args);err!=nil{return fmt.Errorf("invalid demo arguments: %w",err)}
 if flags.NArg()!=0{return fmt.Errorf("unexpected arguments: %v",flags.Args())}
 s,ok:=demoScenarios[*scenario];if !ok{return fmt.Errorf("unknown scenario %q; use running, approvals, incidents or multi",*scenario)}
 if *frames<0 || *interval<100*time.Millisecond{return fmt.Errorf("frames must be >= 0 and interval >= 100ms")}
 for tick:=0;*frames==0||tick<*frames;tick++{
  if _,err:=fmt.Fprint(out,"\033[2J\033[H"+renderDemo(*scenario,s,tick));err!=nil{return err}
  if *frames>0&&tick==*frames-1{break}
  time.Sleep(*interval)
 }
 return nil
}
