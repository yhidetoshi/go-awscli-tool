# go-awscli-tool

Go言語でaws-sdk-goを利用してAWSを制御するツールです。


`$ ore-aws -h`
```
Usage of ore-aws:
  -billing
    	get billing info
  -deregister
    	Deregister Instances to ELB
  -elbname string
    	input elbname
  -instances string
    	 slect Instance ID or Instance Tag:Name or RDSinstanceName
  -profile string
    	slect profile.
  -region string
    	slect Region (default "ap-northeast-1")
  -register
    	Register Instances to ELB
  -resource string
    	select resource
  -show
    	show ELB backendend Instances
  -start
    	Instance start
  -stop
    	Instance stop
  -terminate
    	Instance terminate
  ```
  
 (ex) 
`$ore-aws -resource=ec2 -profile=stg`
```
+------------------------------------+---------------------+--------------+-----------------+----------------+---------------+---------+--------------+-----------------+------------+----------+
|              TAG:NAME              |     INSTANCEID      | INSTANCETYPE |       AZ        |   PRIVATEIP    |   PUBLICIP    | STATUS  |    VPCID     |    SUBNETID     | DEVICETYPE | KEYNAME  |
+------------------------------------+---------------------+--------------+-----------------+----------------+---------------+---------+--------------+-----------------+------------+----------+
| test                           | i-076cbe7cxxxxxx | t2.micro     | ap-northeast-1c | x.x.x.x | 172.31.8.141  | running | vpc-e85a898c | subnet-729b3f2a | ebs        | test-key |
```

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
