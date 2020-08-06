# 系统设计

## 层次设计

* 工具
* 常量
* 模型
* 数据库操作
* 库业务
* 业务
* 接口交易提供
* main.go调用交易提供服务

# 模型设计

* 总领表:表和各个字段的对应关系，关联主键，用于生成和更新结构表中指定的数据表
* 结构表
* 总领历史表
* 数据表

# js方法全部promise，参数第一个必须是调用者方法名


# 接收参数和返回值

## 接收参数

包含内容：

授权码：根据某个时期或版本进行设置的验证密钥，保证app的授权是本企业开发的。虽然可以模仿但是可以增强一下校验也好。最好是临时获取，但这个接口也肯定是会被公开，所以还是得固化在app中更好。
头：
JWT---保存一些用户基本信息，匿名则会有默认值，肯定都会有一个，这是接下来用户访问的标志，所以无论如果第一次访问链接都会带上这个返回给用户，如果存在这个值且没有过期就不会重新生成。
YEWU---想要访问的业务，这个会根据用户是否有权限进行数据库设置
SHIJIAN---访问的时间戳，精确到毫秒
KEHUDUANLEIXING---客户端类型，比如微信，安卓app，苹果app，电脑浏览器，手机浏览器
KEHUDUANBIANMA---根据每个手机下载的不同升序分配一个编码，作为下载次数的计数
json:{SHOUQUAN:"",TOU:{JWT:"",YEWU:"",SHIJIAN:"",KEHUDUANLEIXING:"",KEHUDUANBIANMA:""},SHUJU:{}}
返回：
json:{BianMa:"000000000000",ZhuangTai:"00",MiaoShu:"正常数据",ShuJu:shuJu}
根据业务请求的不同，比如第一次访问业务没有用户信息也没有jwt的时候就应该把jwt返回回去那么jwt就是shuJu，仍然要在内部用名称指定，也就是说用bianma和zhuangtai来作为返回是否有正常数据的判定，必须是一一对应的，后台一定要注意处理所有可能的错误正常返回给前端。前端也要把所有的错误规整到一起，然后未预料的错误就发一个错误日志给后台保存起来。发错误日志这个交易就必须要保证无意外情况。否则将会产生业务死循环。
接下来的重点应该是根据业务操作需求进行不同的代码生成，然后提供一个自动生成的系统给另一个系统使用，也就是说如果生成的系统有修改最终都会体现在生成系统内，这样就可以保证所有的代码生成都是可控的，而且每次生成都是必然可执行的。另外就是自动生成的代码既要有前端后端的可定制性，但也不能完全为了通用而强耦合。
curl -H "Content-Type:application/json" -X POST --data '{SHOUQUAN:"xyz",TOU:{JWT:"abc.bcd.cde",YEWU:"HFX000000001",SHIJIAN:"2020_04_06_17_09_30_222",KEHUDUANLEIXING:"CURL",KEHUDUANBIANMA:"test"},SHUJU:{YongHuMing:"clz",XingBie:"nan",NianLing:"123",MiMa:"123456"}}' localhost:8080/hanfuxin
curl -H "Content-Type:application/json" -X POST --data '{"SHOUQUAN":"xyz","TOU":{"JWT":"abc.bcd.cde","YEWU":"HFX000000001","SHIJIAN":"2020_04_06_17_09_30_222","KEHUDUANLEIXING":"CURL","KEHUDUANBIANMA":"test"},"SHUJU":{"YongHuMing":"clz","XingBie":"nan","NianLing":"123","MiMa":"123456"}}' localhost:8080/hanfuxin










































