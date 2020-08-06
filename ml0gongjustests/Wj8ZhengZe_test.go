package ml0gongjustests
import(
  "log"
  "testing"
  "xm1shengcheng/ml0gongjus"
)
func TestShouJiHaoPiPei(t *testing.T){
  ret := ml0gongjus.ShouJiHaoPiPei("1112222ff3333")
  log.Println(ret,len(ret))
}
