package ml7yewus

import (
  "xm1shengcheng/ml2moxings"
)
// 业务发生时，是在界面点击添加弹出添加界面然后录入了表名
func TianJiaBiao(canShu ml2moxings.CanShu) ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:"00",
  }
  // 调用kuyewu的方法，所有错误日志都要记录到数据表，这个表后面用nosql来替代，canShu中必然有一个表名栏位
  // 1. 录入总领表字段，录入一条没有字段的总领表数据，此时字段名和字段计数是默认值
  // 2. 新增一个表数据表
  // 3. 不用录入结构表，因为结构表里全都是字段的说明。
  return ret
}
// 业务发生时，是在界面添加了一个表之后
func TianJiaZiDuan(canShu ml2moxings.CanShu) ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:"00",
  }
  // 调用kuyewu的方法，所有错误日志都要记录到数据表，这个表后面用nosql来替代，canShu必然有表名和字段名栏位，带上表编码和字段编码吧
  // 1. 录入总领表字段，这时候要查出表名所在的所有数据看是否存在现在录入的名称，前端也控制一次。然后根据录入的名称转化为拼音生成一条字段数据录入。
  // 同时把表数据名查出来。
  // 2. 录入结构表，此时结构表可以生成一条数据，生成一个最大的字段名，搭上最大的1查出来的表名录入一条数据
  // 3. 创建字段数据表，根据新生成的表名用最大的表名新生成一张字段数据表
  // 由于生成表这个操作不可控，所以还是用执行sql的方式完成数据表的创建比较好。
  return ret
}
//由于只有一个结构表进行表示，所以查询字段的意思就是查询所有字段，但有一个总领
func ChaXunZiDuan(canShu ml2moxings.CanShu) ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:"00",
  }
  return ret
}
