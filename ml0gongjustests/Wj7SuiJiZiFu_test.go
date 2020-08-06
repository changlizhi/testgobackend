package ml0gongjustests
import(
  "log"
  "testing"
  "xm1shengcheng/ml0gongjus"
)
func TestSuiJiShuZi(t *testing.T){
  ret := ml0gongjus.SuiJiShuZi(50)
  log.Println(ret,"\n",len(ret))
}
func TestSuiJiDaXieZiMuHeShuZi(t *testing.T){
  ret := ml0gongjus.SuiJiDaXieZiMuHeShuZi(10)
  log.Println(ret,"\n",len(ret))
}
func TestSuiJiXiaoXieZiMuHeShuZi(t *testing.T){
  ret := ml0gongjus.SuiJiXiaoXieZiMuHeShuZi(10)
  log.Println(ret,"\n",len(ret))
}
func TestSuiJiZiMuHeShuZi(t *testing.T){
  ret := ml0gongjus.SuiJiZiMuHeShuZi(10)
  log.Println(ret,"\n",len(ret))
}
