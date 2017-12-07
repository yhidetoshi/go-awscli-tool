# go-awscli-tool

Go言語でaws-sdk-goを利用してAWSを制御するツールです。

(main以外のソースコード)
(go get) https://github.com/yhidetoshi/clitoolgoaws

- go buildでファイルをクロスコンパイルする
- /usr/local/bin等の場所に保存する。(以下は /usr/local/binに `ore-aws` という名前で配置)
  
  
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
    - `$ ore-aws main.go -resource=ec2 -start -instances=<INSTANCEIDNAME> or <INSTANCEID> -profile=stg`
  - 停止
    - `$ ore-aws main.go -resource=ec2 -stop -instances=<INSTANCEIDNAME> or <INSTANCEID> -profile=stg`
  - 削除
    - `$ ore-aws main.go -resource=ec2 -terminate -instances=<INSTANCEIDNAME> or <INSTANCEID> -profile=stg`    
- RDS
  - 一覧  
    - `$ ore-aws -resource=rds -profile=stg`
  - 起動
    - `$ ore-aws main.go -resource=rds -start -instances=<INSTANCEIDNAME> or <INSTANCEID> -profile=stg`
  - 停止
    - `$ ore-aws main.go -resource=rds -stop -instances=<INSTANCEIDNAME> or <INSTANCEID> -profile=stg`
- ELB
  - 一覧
    - `$ ore-aws -resource=elb -profile=stg`
  - ELBのバックエンドインスタンスを取得
    - `$ ore-aws -resource=elb -show -elbname=beaconnect-lb-1 -profile=bct-stg`
  - ELBにバックエンドインスタンスを登録
    - `$ ore-aws -resource=elb -register -elbname=<ELBNAME> -instances=<INSTANCEID> -profile=stg`
  - ELBにバックエンドインスタンスを解除
    - `$ ore-aws -resource=elb -show -elbname=beaconnect-lb-1 -profile=bct-stg`
- Cloudwatch
  - Billing
    - `$ ore-aws -resource=cloudwatch -billing`
  - Alarm
    - `$ ore-aws -resource=cloudwatch`
    
    
