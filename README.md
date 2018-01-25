# go-awscli-tool

Go言語でaws-sdk-goを利用してAWSを制御するツールです。

(main以外のソースコード)
(go get) https://github.com/yhidetoshi/clitoolgoaws

- go buildでファイルをクロスコンパイルする
- /usr/local/bin等の場所に保存する。(以下は /usr/local/binに `ore-aws` という名前で配置)

- macOSとLinux(amd64)用にクロスコンパイルしたバイナリはこちら
  - https://github.com/yhidetoshi/go-awscli-tool/tree/master/bin
  
- 参考
  - `GoDoc`: https://godoc.org/github.com/aws/aws-sdk-go/aws
  - `aws-sdk-go APIリファレンス`: https://docs.aws.amazon.com/sdk-for-go/api/
  
  
 (ex) 
`$ore-aws -resource=ec2 -profile=stg`
```
+------------------------------------+---------------------+--------------+-----------------+----------------+---------------+---------+--------------+-----------------+------------+----------+
|              TAG:NAME              |     INSTANCEID      | INSTANCETYPE |       AZ        |   PRIVATEIP    |   PUBLICIP    | STATUS  |    VPCID     |    SUBNETID     | DEVICETYPE | KEYNAME  |
+------------------------------------+---------------------+--------------+-----------------+----------------+---------------+---------+--------------+-----------------+------------+----------+
```

### コマンドオプション
- EC2
  - 一覧  
    - `$ ore-aws -resource=ec2 -profile=stg`
  - 起動
    - `$ ore-aws -resource=ec2 -start -instances=<INSTANCENAME> or <INSTANCEID> -profile=stg`
  - 停止
    - `$ ore-aws -resource=ec2 -stop -instances=<INSTANCENAME> or <INSTANCEID> -profile=stg`
  - 削除
    - `$ ore-aws -resource=ec2 -terminate -instances=<INSTANCENAME> or <INSTANCEID> -profile=stg`
  - AMI焼き(/binのバイナリには未追加)  
    - `$ ore-aws -resource=ec2 -ami -aminame=<AMINAME> -instances=<INSTANCENAME> or <INSTANCEID> -profile=stg`
  - AMI情報を取得(/binのバイナリには未追加)  
    - `$ ore-aws -resource=ec2 -amilist -profile=stg`
  - AMIの削除(解除)(/binのバイナリには未追加)
    - `$ ore-aws -resource=ec2 -deregister -amiid=<ami-id> -profile=stg`
  
- RDS
  - 一覧  
    - `$ ore-aws -resource=rds -profile=stg`
  - 起動
    - `$ ore-aws -resource=rds -start -instances=<INSTANCENAME> or <INSTANCEID> -profile=stg`
  - 停止
    - `$ ore-aws -resource=rds -stop -instances=<INSTANCENAME> or <INSTANCEID> -profile=stg`  
- ELB
  - 一覧
    - `$ ore-aws -resource=elb -profile=stg`
  - ELBのバックエンドインスタンスを取得
    - `$ ore-aws -resource=elb -show -elbname=<ELBNAME> -profile=stg`
  - ELBにバックエンドインスタンスを登録
    - `$ ore-aws -resource=elb -register -elbname=<ELBNAME> -instances=<INSTANCEID> -profile=stg`
  - ELBにバックエンドインスタンスを解除
    - `$ ore-aws -resource=elb -show -elbname=<ELBNAME> -profile=stg`
- Cloudwatch
  - Billing
    - `$ ore-aws -resource=cloudwatch -billing`
  - Alarm
    - `$ ore-aws -resource=cloudwatch`
- IAM
  - ユーザ一覧
    - `$ ore-aws -resource=iam-user -profile=stg`
  - グループ一覧
    - `$ ore-aws -resource=iam-group -profile=stg`  
    
