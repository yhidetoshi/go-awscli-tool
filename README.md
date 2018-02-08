![Alt Text](https://github.com/yhidetoshi/Pictures/raw/master/Go_study/gophertraining.png)

# go-awscli-tool
- Go言語でaws-sdk-goを利用してAWSを制御するツールです。(勉強しながら作っているので、どこかできれいに作りなおします...)
- 使用するには awsのapiをコールできる設定が必要です。--> (./awsの設定)

(main以外のソースコード)
(go get) https://github.com/yhidetoshi/clitoolgoaws

- go buildでファイルをクロスコンパイルする
- /usr/local/bin等の場所に保存する。(以下は /usr/local/binに `ore-aws` という名前で配置)

- macOSとLinux(amd64)とWindows用にクロスコンパイルしたバイナリはこちら
  - https://github.com/yhidetoshi/go-awscli-tool/tree/master/bin
  
- 参考
  - `GoDoc`: https://godoc.org/github.com/aws/aws-sdk-go/aws
  - `aws-sdk-go APIリファレンス`: https://docs.aws.amazon.com/sdk-for-go/api/
  
  
 (ex) 
`$ore-aws -resource=ec2 -profile=stg`

# コマンドオプション
- EC2
  - 一覧  
    - `$ ore-aws -resource=ec2 -profile=stg`
  - 起動
    - `$ ore-aws -resource=ec2 -start -instances=<INSTANCENAME> or <INSTANCEID> -profile=stg`
  - 停止
    - `$ ore-aws -resource=ec2 -stop -instances=<INSTANCENAME> or <INSTANCEID> -profile=stg`
  - 削除
    - `$ ore-aws -resource=ec2 -terminate -instances=<INSTANCENAME> or <INSTANCEID> -profile=stg`
  - AMI焼き
    - `$ ore-aws -resource=ec2 -ami -aminame=<AMINAME> -instances=<INSTANCENAME> or <INSTANCEID> -profile=stg`
  - AMI情報の一覧を取得
    - `$ ore-aws -resource=ec2 -amilist -profile=stg`
  - AMIの削除(解除)
    - `$ ore-aws -resource=ec2 -deregister -amiid=<ami-id> -profile=stg`
  - ElasticIPの一覧をを取得
    - `$ ore-aws -resource=ec2 -deleteeip -eiplist`
  - ElasticIPのリリース
    - `$ ore-aws -resource=ec2 -deleteeip -allocationid=<ALLOCATIONID>`
  - SecurityGroupの一覧を取得
    - `$ ore-aws -resource=ec2 -sglist`
  - SecurityGroupのルール確認
    - `$ ore-aws -resource=ec2 -show -sgid=<GROUPID>`
  - インスタンスを 複数同時に操作するときは `,` で区切り複数指定する
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
- S3
  - バケット一覧
    - `$ ore-aws -resource=s3 -profile=stg`
  - バケットのオブジェクト一覧を取得
    - `ore-aws -resource=s3 -show -bucket=<BUCKETNAME> -profile=stg`
  - バケットのサイズ取得
    - `ore-aws -resource=s3 -size -bucket=<BUCKETNAME> -profile=stg` 
  - バケットの削除(条件: bucket　is empty) ./bin配下は未反映
    - `ore-aws -resource=s3 -deletebucket -bucket=<BUCKETNAME> -profile=stg`
  - オブジェクト削除
    - `ore-aws -resource=s3 -deleteobject -bucket=<BUCKETNAME> -object=<FILENAME>`
  - バケット内のオブジェクトを全て削除
    - `ore-aws -resource=s3 -deleteallobject -bucket=<BUCKETNAME>`
  - Tokyoリージョン内の全バケットのACLがPublic or Privateかを取得
    - `ore-aws -resource=s3 -checkacl -profile=stg`
  - Tokyoリージョン内の全バケットサイズを取得
    - `ore-aws -resource=s3 -sizeall -profile=stg`
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

