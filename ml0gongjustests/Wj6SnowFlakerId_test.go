package ml0gongjustests
import(
  "log"
  "testing"
  "xm1shengcheng/ml0gongjus"
)
func TestHuoquId(t *testing.T){
  ret := ml0gongjus.HuoquId()
  log.Println(ret)
}
func TestHuoQuIdZiFu(t *testing.T){
  for i:=0;i<40;i++{
    ret := ml0gongjus.HuoQuIdZiFu()
    log.Println(ret,len(ret))
  }
}
