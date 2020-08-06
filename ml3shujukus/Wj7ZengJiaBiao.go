package ml3shujukus

import(
  "log"
  "xm1shengcheng/ml2moxings"
)

func ZengJiaBiao(sql string)ml2moxings.FanHui{
  err := LianJieChi().Exec(sql)
  log.Println(err)
  ret := ml2moxings.FanHui{
    ZhuangTai:"00",
  }
  return ret
}
func ZengJiaBsj(jiShu string)ml2moxings.FanHui{
  sql := "CREATE TABLE Bsj" + jiShu + " (ZhuJian TEXT NOT NULL DEFAULT ('hfxmoren'))"
  ZengJiaBiao(sql)
  ret := ml2moxings.FanHui{
    ZhuangTai:"00",
  }
  return ret
}