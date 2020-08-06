package ml7yewus

import (
  "xm1shengcheng/ml2moxings"
)

func YeWuFenPei(canShu ml2moxings.CanShu)ml2moxings.FanHui{
  // 1.校验jwt,
  // 2.校验交易是否允许该角色访问，
  // 3.校验交易字段是否齐全和符合需求，
  // 4.交给正式业务方法进行业务操作
  // 返回也是从业务数据层封装好的，这里只要json返回即可
  ret := ml2moxings.FanHui{}
  jwtRet := ml2moxings.FanHui{}
  LuRuRiZhi(canShu)//录入日志，无论如何都录入，也要打在日志文件里
  if jwtRet = JiaoYanJwt(canShu.JWT);jwtRet.ZhuangTai != "00"{
    ret=jwtRet
    return ret
  }
  if ret = JiaoYanQuanXian(canShu,jwtRet.ShuJu["jueSeBianMa"]);ret.ZhuangTai != "00"{
    return ret
  }
  if ret = JiaoYanShuJu(canShu);ret.ZhuangTai != "00"{
    return ret
  }
  ret = FenPeiYeWu(canShu)
  return ret
}
func LuRuRiZhi(canShu ml2moxings.CanShu) ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:"00",
  }
  return ret
}
func JiaoYanJwt(jwt string) ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:"00",
  }
  return ret
}
func JiaoYanQuanXian(canShu ml2moxings.CanShu,jwtJueSe string) ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:"00",
  }
  return ret
}
func JiaoYanShuJu(canShu ml2moxings.CanShu) ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:"00",
  }
  return ret
}
func FenFa(canShu ml2moxings.CanShu) ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:"00",
  }
  //这里根据不同的交易码分配调用不同的业务返回，同样库业务也是能够返回FanHui数据的，这样的数据在每个方法中都能进行解析和显示
  return ret
}
