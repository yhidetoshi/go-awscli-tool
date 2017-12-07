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
  
`$ore-aws -resource=ec2 -profile=stg`
```
+------------------------------------+---------------------+--------------+-----------------+----------------+---------------+---------+--------------+-----------------+------------+----------+
|              TAG:NAME              |     INSTANCEID      | INSTANCETYPE |       AZ        |   PRIVATEIP    |   PUBLICIP    | STATUS  |    VPCID     |    SUBNETID     | DEVICETYPE | KEYNAME  |
+------------------------------------+---------------------+--------------+-----------------+----------------+---------------+---------+--------------+-----------------+------------+----------+
| test                           | i-076cbe7cxxxxxx | t2.micro     | ap-northeast-1c | x.x.x.x | 172.31.8.141  | running | vpc-e85a898c | subnet-729b3f2a | ebs        | test-key |
```
