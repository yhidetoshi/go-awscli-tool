![Alt Text](https://github.com/yhidetoshi/Pictures/raw/master/Go_study/gophertraining.png)

# go-awscli-tool
- Go言語でaws-sdk-goを利用してAWSを制御するツールです。(勉強しながら作っています)

(main以外のソースコード)
(go get) https://github.com/yhidetoshi/clitoolgoaws

- go buildでファイルをクロスコンパイルする
- /usr/local/bin等の場所に保存する。(以下は /usr/local/binに `ore-aws` という名前で配置)
- macOSとLinux(amd64)とWindows用にクロスコンパイルしたバイナリはこちら
  - https://github.com/yhidetoshi/go-awscli-tool/tree/master/bin
    
  
 (ex) 
`$ore-aws -resource=ec2 -profile=stg`

# 使い方(コマンドオプション)
### EC2
```
  ■ 一覧  
    > $ ore-aws -resource=ec2 -profile=<PROFILE>
  
  ■ 起動
    > $ ore-aws -resource=ec2 -start -instances=<INSTANCENAME> or <INSTANCEID> -profile=<PROFILE>
  ■ 停止
    > $ ore-aws -resource=ec2 -stop -instances=<INSTANCENAME> or <INSTANCEID> -profile=<PROFILE>
  
  ■ 削除
    > $ ore-aws -resource=ec2 -terminate -instances=<INSTANCENAME> or <INSTANCEID> -profile=<PROFILE>
  ※ インスタンスを 複数同時に操作するときは `,` で区切り複数指定する
  
  ■ AMI焼き
    > $ ore-aws -resource=ec2 -ami -aminame=<AMINAME> -instances=<INSTANCENAME> or <INSTANCEID> -profile=<PROFILE>
  
  ■ AMI情報の一覧を取得
    > $ ore-aws -resource=ec2 -amilist -profile=<PROFILE>
  
  ■ AMIの削除(解除)
    > $ ore-aws -resource=ec2 -deregister -amiid=<ami-id> -profile=<PROFILE>
  
  ■ ElasticIPの一覧をを取得
    > $ ore-aws -resource=ec2 -eiplist -profile=<PROFILE>
  
  ■ ElasticIPのリリース
    > $ ore-aws -resource=ec2 -deleteeip -allocationid=<ALLOCATIONID> -profile=<PROFILE>
  
  ■ SecurityGroupの一覧を取得
    > $ ore-aws -resource=ec2 -sglist -profile=<PROFILE>
  
  ■ SecurityGroupのルール確認(output:json)
    > $ ore-aws -resource=ec2 -show -sgid=<GROUPID> -profile=<PROFILE>
  
  ■ AutoScalingグループ一覧の情報を出力(コンソール画面相当)
    > $ ore-aws -resource=as -profile=<PROFILE>
  
  ■ AutoScalingグループのインスタンスの`最大`数を変更
    > $ ore-aws -resource=as -asg=<GROUP名> -max -num=<NUM> -profile=<PROFILE>
  
  ■ AutoScalingグループのインスタンスの`最小`数を変更
    > $ ore-aws -resource=as -asg=<GROUP名> -min -num=<NUM> -profile=<PROFILE>
  
  ■ AutoScalingグループのインスタンスの`希望`数を変更
    > $ ore-aws -resource=as -asg=<GROUP名> -desire -num=<NUM> -profile=<PROFILE>
```

### RDS
```
  ■ 一覧  
    > $ ore-aws -resource=rds -profile=<PROFILE>
  
  ■ 起動
    > $ ore-aws -resource=rds -start -instances=<INSTANCENAME> or <INSTANCEID> -profile=<PROFILE>
  
  ■ 停止
    > $ ore-aws -resource=rds -stop -instances=<INSTANCENAME> or <INSTANCEID> -profile=<PROFILE>  
```

### ELB
```
  ■ 一覧
    > $ ore-aws -resource=elb -profile=<PROFILE>
  
  ■ ELBのバックエンドインスタンスを取得
    > $ ore-aws -resource=elb -show -elbname=<ELBNAME> -profile=<PROFILE>
  
  ■ ELBにバックエンドインスタンスを登録
    > $ ore-aws -resource=elb -register -elbname=<ELBNAME> -instances=<INSTANCEID> -profile=<PROFILE>
  
  ■ ELBにバックエンドインスタンスを解除
    > $ ore-aws -resource=elb -deregister -elbname=<ELBNAME> -instances=<INSTANCEID> -profile=<PROFILE>
```

### S3
```
  ■ バケット一覧
    > $ ore-aws -resource=s3 -profile=<PROFILE>
  
  ■ バケットのオブジェクト一覧を取得
    > $ ore-aws -resource=s3 -show -bucket=<BUCKETNAME> -profile=<PROFILE>
  
  ■ バケットのサイズ取得
    > $ ore-aws -resource=s3 -size -bucket=<BUCKETNAME> -profile=<PROFILE> 
  
  ■ バケットの削除(条件: bucket　is empty) 
    > $ ore-aws -resource=s3 -deletebucket -bucket=<BUCKETNAME> -profile=<PROFILE>
  
  ■ オブジェクト削除
    > $ ore-aws -resource=s3 -deleteobject -bucket=<BUCKETNAME> -object=<FILENAME>
  
  ■ バケット内のオブジェクトを全て削除
    > $ ore-aws -resource=s3 -deleteallobject -bucket=<BUCKETNAME> -profile=<PROFILE>

  ■ Tokyoリージョン内の全バケットのACLがPublic or Privateかを取得
    > $ ore-aws -resource=s3 -checkacl -profile=<PROFILE>

  ■ Tokyoリージョン内の全バケットサイズを取得
    > $ ore-aws -resource=s3 -sizeall -profile=<PROFILE>
```

### Route53
```
  ■ Zone一覧
    > $ ore-aws -resource=route53 -profile=<PROFILE>
  ■ Zoneid指定のレコード情報取得
    > $ ore-aws -resource=route53 -show -zoneid=<ZONEID> -profile=<PROFILE>
```

### Cloudwatch
```
  ■ Billing
    > $ ore-aws -resource=cloudwatch -billing -profile=<PROFILE>
  ■ Alarm
    > $ ore-aws -resource=cloudwatch -profile=<PROFILE>
```

### IAM
```
  ■ ユーザ一覧
    > $ ore-aws -resource=iam-user -profile=<PROFILE>
  ■ グループ一覧
    > $ ore-aws -resource=iam-group -profile=<PROFILE>
```

### 1リージョンの全バケットの合計サイズを計算する処理をGoroutineで処理するように変更

- 一部抜粋
```
wg := new(sync.WaitGroup)
ch := make(chan int64, len(allBuckets))

wg.Add(len(allBuckets))
	for i := 0; i < len(allBuckets); i++ {
		buffBucket = &allBuckets[i]
		go CalcThreadBucketSize(S3Client, buffBucket, ch, wg)
	}
```

■ マルチスレッド実行
```
➜  aws-cli-go git:(master) ✗ ore-aws -resource=s3 -sizeall -profile=infra-stg
+-----------------+
| TOTAL SIZE(KIB) |
+-----------------+
|      7897247147 |
+-----------------+
1.711132秒
```

■ シングルスレッド実行
```
➜ aws-cli-go git:(master) ✗ ore-aws -resource=s3 -sizeall -profile=infra-stg
+-----------------+
| TOTAL SIZE(KIB) |
+-----------------+
|      7897247147 |
+-----------------+
3.277896秒
```

### AWS-SDK-Goのドキュメントを読んでいてのメモ

- APIリクエストの必須な引数について
  - `type ReleaseAddressInput`

```
type ReleaseAddressInput struct {

    // [EC2-VPC] The allocation ID. Required for EC2-VPC.
    AllocationId *string `type:"string"`

    // Checks whether you have the required permissions for the action, without
    // actually making the request, and provides an error response. If you have
    // the required permissions, the error response is DryRunOperation. Otherwise,
    // it is UnauthorizedOperation.
    DryRun *bool `locationName:"dryRun" type:"boolean"`

    // [EC2-Classic] The Elastic IP address. Required for EC2-Classic.
    PublicIp *string `type:"string"`
    // contains filtered or unexported fields
}
```

- VPCがデフォルトになってからは
  - `// [EC2-VPC] The allocation ID. Required for EC2-VPC.` に従う
- Clasicの時
  - `// [EC2-Classic] The Elastic IP address. Required for EC2-Classic` に従う
  
- 参考
  - `GoDoc`: https://godoc.org/github.com/aws/aws-sdk-go/aws
  - `aws-sdk-go APIリファレンス`: https://docs.aws.amazon.com/sdk-for-go/api/

