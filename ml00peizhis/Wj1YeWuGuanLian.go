package ml00peizhis

import(
  "xm1shengcheng/ml1changliangs"
  "xm1shengcheng/ml2moxings"
)
//再次进行数据结构的约定，
//JieGou中，BianMa=BiaoMing+"_"+ZiDuanMing，当需要进行某一个表的查询是必然可以通过这个前缀拿到这个表的所有字段
//BsjBiao=Bsj+N+s，当需要进行数据表增加时必须用当前所有数据表的名称拿出之后加一进行新创建，
//ZdsjBiao=Zdsj+N+s同表数据表的创建
//CREATE TABLE `jiegou` (
// 	`BianMa` VARCHAR(1024) NOT NULL DEFAULT 'hfxmoren',
// 	`MingCheng` VARCHAR(1024) NOT NULL DEFAULT 'hfxmoren',
// 	`BsjBiao` VARCHAR(1024) NOT NULL DEFAULT 'hfxmoren',
// 	`ZdsjBiao` VARCHAR(1024) NOT NULL DEFAULT 'hfxmoren',
// 	`ZhengZe` VARCHAR(1024) NOT NULL DEFAULT 'hfxmoren',
// 	`ShiJian` VARCHAR(1024) NOT NULL DEFAULT 'hfxmoren'
// )
// COLLATE='utf8mb4_general_ci'
// ENGINE=InnoDB
// ;

//也就是说计数和字段汇总就在一个表中完成了，当界面上查看所有表时，意思是要截取JieGou表中的所有表前缀进行分组
//并把名称也进行相同格式分组显示，然后当点击时也是用前缀进行匹配完成查询所有的字段，这样001交易就算是完成了。于是又删除一个基础表，
//所有的基础表只有一个。只是说明了结构，这个表没有业务数据，被拿走也没关系，而业务数据都在bsj这种表中，
//也是只能拿一个字段，天生具有一定的分布式安全性

//所有的交易都要来查这个配置，然后根据说明进行数据组装返回
// func SmJY0000000001(canShu ml2moxings.CanShu)ml2moxings.FanHui{
//   ret := ml2moxings.FanHui{
//     ZhuangTai:ml1changliangs.CHENGGONG,
//     BianMa:ml1changliangs.HFX000000000,
//     BiaoZhi:ml1changliangs.LieBiaosDuiXiangs,
//   }
//   ret.ShuJuLieBiaosDuiXiangs=[]map[string][]string{}
//   liuCheng1=map[string][]string{
//     "LeiBie":[]string{"ChaXun"},
//     "LieTou":[]string{"BiaoMing","ZhongWenBiaoMing"},
//     "CanShuZhi":[]string{},
//     "ShuoMing":[]string{"SQL:截取获取所有JieGou中的表名编码，中文表名编码","放到FanHui中的ShuJu"},
//   }
//   ret.ShuJuLieBiaosDuiXiangs=append(ret.ShuJuLieBiaosDuiXiangs,liuCheng1)
//   return ret
// }
  //这里配置所有业务的关联sql和查询方案，回显内容格式等，这是一个非常重要的方法
func YeWuGuanLian(canShu ml2moxings.CanShu)ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:ml1changliangs.CHENGGONG,
    BianMa:ml1changliangs.HFX000000000,
    BiaoZhi:ml1changliangs.LieBiaosDuiXiangs,
  }
  if canShu.YEWU == ml1changliangs.JY0000000001 {
    return GlJY0000000001(canShu)
  }
  if canShu.YEWU == ml1changliangs.JY0000000002 {
    return GlJY0000000002(canShu)
  }
  return ret
}


//比如001交易，查询所有制定好的业务表，这是一个系统生成表，当查询此交易时，一般是管理员在操作，所以全部字段返回是没问题的。
//而根据命名方案来说，001交易对应的sql也就只有一个，而即使只有一个也要对应下来
//就是select BianMa,MingCheng,BsjBiao,ZdsjBiao,ZhengZe,ShiJian, from JieGou而不能用星号所以这里返回的数据
//是设置完值之后返回，这个方法感觉不太喜欢，但是好像append也是这样做的。
//由于这个设计方案的原因，后面查询业务数据的时候很多join的数据，所以会有很多表名应该如何定这些值呢？全都用数组的方式，
//用下划线进行分割？也就是放到map里了，这样还是不太合适，那就在FanHui里再定一个数据对象返回map中的value为字符串列表，
//这样的话如果是对象列表时，界面上可以直接循环展示，列表则可以更加直接的使用，
//而value是一个列表的情况则说明有一对多的关系需要处理，就是这里纯框架处理的功能
//这个设计思路就是，dao层不能自己写sql，而是必须从配置中读取业务所关联的sql在kuyewus层进行组织调用这里的方法得到了sql全量，
//然后dao层执行kuyewus传下来的真实sql，执行之后根据情况把数据组装之后返回给kuyewus层，根本不需要知道业务含义，
//也就是说statement要专注于执行statement的事情，不关心任何业务层面的事情。
//因为最麻烦的就是处理sql和业务之间的关系，这里进行梳理之后会极大解放后续维护问题

//从使用逻辑上来说，JieGou表只有管理员会进行增删改，业务只会来查询并进行组装，所以001交易主要是查询有多少根据下划线拆分的表名

//查询所有表名，只要通过了前面的所有检查，到这里就是直接查询所有的表名回去就行了，编码名称而已。
//有另外一个需求，比如增加一个用户，要增加很多个表的数据，这种情况就会需要很多条sql执行，并且必须事务提交的方式进行，
//于是就会要很多条sql
//查询系统中的所有表名
//应该制定一个sql执行方案，每个交易应该对应很多sql，每个sql执行又分为几个阶段，所以应该是一个交易对应的是sql的个数，根据个数循环挨个找出
//找出sql之后执行得到数据，然后指定执行方案，得到中间结果，等待执行第二个sql，和第二个需要执行的sql数据进行合并，依次类推之后得到想要返回给界面的结果
//所以应该是这样的方法
//交易对应sql的计数，然后又对应着执行方案，和返回方案。
// Sql1-n:根据交易码规定只有一个sql，是根据交易码区分的。一般情况下一个交易应该不会超过10个sql需要执行的情况。
//这样每个返回都只会有一个sql，lie也就不需要split就可以直接用了
//说明交易要我自己设计一个针对交易的说明方案，这里返回一个对象即可
func ShuoMingJiaoYi(canShu ml2moxings.CanShu)ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:ml1changliangs.CHENGGONG,
    BianMa:ml1changliangs.HFX000000000,
    BiaoZhi:ml1changliangs.LieBiaosDuiXiangs,
  }
  return ret
}

func GlJY0000000001(canShu ml2moxings.CanShu)ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:ml1changliangs.CHENGGONG,
    BianMa:ml1changliangs.HFX000000000,
    BiaoZhi:ml1changliangs.LieBiaosDuiXiangs,
  }
  ret.ShuJuLieBiaosDuiXiangs=[]map[string][]string{}
  firstSqlMap := map[string][]string{
    //这里决定需要多少个sql,需要多少个列头来组装map，需要多少个
    "LeiXings":[]string{"ChaXun"},//有多少个leixing就有多少个sql需要执行
    "Lies":[]string{"BiaoMing","ZhongWenBiaoMing"},
    "Biaos":[]string{"(SELECT SUBSTR(BianMa,1,instr(BianMa,'_') - 1) BiaoMing,SUBSTR(MingCheng,1,instr(MingCheng,'_') - 1) ZhongWenBiaoMing FROM jiegou) linShi"},
    "TiaoJians":[]string{""},
    "TiaoJianZhis":[]string{""},
    "JuHes":[]string{"GROUP BY BiaoMing"},
  }
  ret.ShuJuLieBiaosDuiXiangs=append(ret.ShuJuLieBiaosDuiXiangs,firstSqlMap)
  return ret
}

// 查询结构，肯定是选中了某个表
func GlJY0000000002(canShu ml2moxings.CanShu)ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:ml1changliangs.CHENGGONG,
    BianMa:ml1changliangs.HFX000000000,
    BiaoZhi:ml1changliangs.LieBiaosDuiXiangs,
  }
  ret.ShuJuLieBiaosDuiXiangs=[]map[string][]string{}
  firstSqlMap := map[string][]string{
    //这里决定需要多少个sql,需要多少个列头来组装map，需要多少个
    "LeiXings":[]string{"ChaXun"},//有多少个leixing就有多少个sql需要执行
    "Lies":[]string{"BiaoMing","ZhongWenBiaoMing"},
    "Biaos":[]string{"(SELECT SUBSTR(BianMa,1,instr(BianMa,'_') - 1) BiaoMing,SUBSTR(MingCheng,1,instr(MingCheng,'_') - 1) ZhongWenBiaoMing FROM jiegou) linShi"},
    "TiaoJians":[]string{""},
    "TiaoJianZhis":[]string{""},
    "JuHes":[]string{"GROUP BY BiaoMing"},
  }
  ret.ShuJuLieBiaosDuiXiangs=append(ret.ShuJuLieBiaosDuiXiangs,firstSqlMap)
  return ret
}
// 删除字段，不回收数据表名，设计表的时候用心设计

func GlJY0000000001(canShu ml2moxings.CanShu)ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:ml1changliangs.CHENGGONG,
    BianMa:ml1changliangs.HFX000000000,
    BiaoZhi:ml1changliangs.LieBiaosDuiXiangs,
  }
  ret.ShuJuLieBiaosDuiXiangs=[]map[string][]string{}
  firstSqlMap := map[string][]string{
    //这里决定需要多少个sql,需要多少个列头来组装map，需要多少个
    "LeiXings":[]string{"ChaXun"},//有多少个leixing就有多少个sql需要执行
    "Lies":[]string{"BiaoMing","ZhongWenBiaoMing"},
    "Biaos":[]string{"(SELECT SUBSTR(BianMa,1,instr(BianMa,'_') - 1) BiaoMing,SUBSTR(MingCheng,1,instr(MingCheng,'_') - 1) ZhongWenBiaoMing FROM jiegou) linShi"},
    "TiaoJians":[]string{""},
    "TiaoJianZhis":[]string{""},
    "JuHes":[]string{"GROUP BY BiaoMing"},
  }
  ret.ShuJuLieBiaosDuiXiangs=append(ret.ShuJuLieBiaosDuiXiangs,firstSqlMap)
  
  return ret
}
//修改字段，这个交易一般用来改名称或者编码，其他的不改，因为是程序生成控制的
func GlJY0000000004(canShu ml2moxings.CanShu)ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:ml1changliangs.CHENGGONG,
    BianMa:ml1changliangs.HFX000000000,
    BiaoZhi:ml1changliangs.LieBiaosDuiXiangs,
  }
  ret.ShuJuLieBiaosDuiXiangs=[]map[string][]string{
    "sqls":[]string{"select $lies FROM (SELECT SUBSTR(BianMa,1,instr(BianMa,'_') - 1) BiaoMing,SUBSTR(MingCheng,1,instr(MingCheng,'_') - 1) ZhongWenBiaoMing FROM jiegou) linShi GROUP BY BiaoMing"},
    "lies":[]string{"BiaoMing,ZhongWenBiaoMing"},
    "zhis":[]string{""},
  }
  return ret
}
//增加用户：能够到这一步肯定是数据参数已经能够满足交易入库了，
//1.先根据交易看有多少sql顺序查的，查结构表，指向YongHu前缀的结构，找到所有的Zdsj表，然后根据这个表，所以交易应该在每条sql前指定主表和关联表用于执行查询
//2.所以，如果是多个sql，这些sql都最终会形成一个通用化内容。
func GlJY0000000005(canShu ml2moxings.CanShu)ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:ml1changliangs.CHENGGONG,
    BianMa:ml1changliangs.HFX000000000,
    BiaoZhi:ml1changliangs.LieBiaosDuiXiangs,
  }
  ret.ShuJuLieBiaosDuiXiangs=[]map[string][]string{//这个列表也是可以查JieGou表得到的，只要这个sql足够通用就行。知道这个表的业务逻辑
    "sqls":[]string{//sql写的时候要控制好可以直接用列当作查询返回值的方式
      "SELECT lies#1 FROM JieGou WHERE BianMa LIKE 'Zhis#1_%'",
      "insert into Bsj1s(lies#2) values(?)",
      "insert into Zdsj1s(lies#3) values(?,?)",//这两个是根据第一条查出的Bsj个数进行生成的
      "insert into Zdsj2s(lies#4) values(?,?)",//这两个是根据第一条查出的Bsj个数进行生成的
    },
  }
  return ret
}
