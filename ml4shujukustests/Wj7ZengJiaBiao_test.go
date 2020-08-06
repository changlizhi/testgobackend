package ml4shujukustests
import(
  "testing"
  "xm1shengcheng/ml3shujukus"
  "log"
)

func TestZengJiaBiao(t *testing.T){
  sql := "create table abc(bcd text)"
  ret := ml3shujukus.ZengJiaBiao(sql)
  log.Println("Wj3ZengJiaBiao_test.go---TestZengJiaBiao,ret",ret)
}
func TestZengJiaBsj(t *testing.T){
  jiShu:="3"
  ret := ml3shujukus.ZengJiaBsj(jiShu)
  log.Println("Wj3ZengJiaBsj_test.go---TestZengJiaBsj,ret",ret)
}