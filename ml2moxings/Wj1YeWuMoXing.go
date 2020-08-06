package ml2moxings
//最终发现JiShu表不能有数据表，关联表也不能有，否则没完没了了，当然还有Bsj和Zdsj不能有数据表。
//也就是说计数表和关联表以及本身已经是数据表的表不能有数据表对应，那就在Biao1s和ZiDuan2s表中以BuKeYou的值方式存放
//按照这个逻辑的话Biao1s和ZiDuan2s也是不应该存放在Biao1s中的，因为如果按照存放的意思的话，业务数据表和系统结构表就重复存在在Biao1s
//指定的数据表中了，这个业务又不是为了生成我现在这个系统的，如果要从里面拿到关联数据反而把获取数据的方式整复杂了。所以不要用Biao1s存放'Biao1s'的方式，除非以后要专门针对这个设计做一个系统来优化现在的编码本身。
//从设计的方案来说会在某个阶段被创建的表只有Bsj和Zdsj这两种，从一个业务系统的角度来说也只要这两种就够了。最麻烦的是用户做了一个表格，然后这个表格的字段成了其他用户录入的表单数据
//从现在的设计思路看来，shujuku层也就是之前称之为dao层的地方只是一个操作数据库的地方，不应该跟对象有什么组合，只有在yewus层才会有对象概念。

// 这两个对象应该是贯穿整个业务周期的，所有应该返回提醒给其他调用者的内容都由这两个引发，所以这两个对象应该是底层的，放在高层不太好
type CanShu struct{
  JWT string
  SHOUQUAN string
  YEWU string
  SHIJIAN string
  KEHUDUANLEIXING string
  KEHUDUANJISHU string
  SHUJU map[string]string
}

type FanHui struct{
  Jwt string
  BianMa string
  ZhuangTai string
  MiaoShu string
  BiaoZhi string//说明返回类型等，此字段只能是一个完整词汇
  ShuJuLieBiaos []string
  ShuJuDuiXiangs []map[string]string
  ShuJuLieBiaosDuiXiangs []map[string][]string
}

