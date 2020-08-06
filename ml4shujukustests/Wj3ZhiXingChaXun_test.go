package ml4shujukustests
import(
  "testing"
  "xm1shengcheng/ml3shujukus"
  "xm1shengcheng/ml1changliangs"
  "xm1shengcheng/ml2moxings"
  "log"
)

func TestZhiXingChaXun(t *testing.T){
  canShu:=ml2moxings.CanShu{
    YEWU:ml1changliangs.JY0000000001,
    SHUJU:map[string]string{},
  }
  ret := ml3shujukus.ZhiXingChaXun(canShu)
  log.Println("Wj3ZhiXingChaXun_test.go---TestZhiXingChaXun,ret",ret)
}