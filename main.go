package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/yhidetoshi/clitoolgoaws"
)

const (
	BILLING_REGION = "us-east-1"
)

var (
	argResource         = flag.String("resource", "", "select resource")
	argProfile          = flag.String("profile", "", "slect profile.")
	argRegion           = flag.String("region", "ap-northeast-1", "slect Region")
	argInstances        = flag.String("instances", "", " slect Instance ID or Instance Tag:Name or RDSinstanceName ")
	argELBName          = flag.String("elbname", "", "input elbname")
	argAmiName          = flag.String("aminame", "", "input ami name")
	argAmiId            = flag.String("amiid", "", "input ami id")
	argBucket           = flag.String("bucket", "", "input bucket name")
	argObject           = flag.String("object", "", "input object name")
	argASG              = flag.String("asg", "", "input autoscaling group name")
	argLaunchConfigName = flag.String("lcname", "", "input launchconfig name")
	argInstanceProfile  = flag.String("instanceprofile", "", "input instance proifle")
	argImageId          = flag.String("imageid", "", "input image id")
	argImageName        = flag.String("imagename", "", "input image name")
	argKeyName          = flag.String("keyname", "", "input key name")
	argInstanceType     = flag.String("instancetype", "", "input instance type")
	argSecurityGroupId  = flag.String("sgids", "", "input securitygroup id")
	argAllocationId     = flag.String("allocationid", "", "input allocationid")
	argNum              = flag.Int64("num", 0, "input num")
	argSGId             = flag.String("sgid", "", "input securityGroup id")
	argZoneId           = flag.String("zoneid", "", "input zone id")
	argAMI              = flag.Bool("ami", false, "create ami")
	argCreate           = flag.Bool("create", false, "create launch config")
	argAMIList          = flag.Bool("amilist", false, "list ami")
	argEIPList          = flag.Bool("eiplist", false, "eiplist ami")
	argStop             = flag.Bool("stop", false, "Instance stop")
	argStart            = flag.Bool("start", false, "Instance start")
	argShow             = flag.Bool("show", false, "show ELB backendend Instances")
	argBilling          = flag.Bool("billing", false, "get billing info")
	argBucketList       = flag.Bool("bucketlist", false, "get billing info")
	argSecurityGroup    = flag.Bool("sglist", false, "get security group")
	argBucketDelete     = flag.Bool("deletebucket", false, "delete bucket")
	argObjectDelete     = flag.Bool("deleteobject", false, "delete object")
	argObjectsAllDelete = flag.Bool("deleteallobject", false, "delete object")
	argAMItest          = flag.Bool("amitest", false, "delete object")
	argEIPDelete        = flag.Bool("deleteeip", false, "delete eip")
	argSize             = flag.Bool("size", false, "calc bucket size")
	argSizeAll          = flag.Bool("sizeall", false, "calc all bucket size")
	argCheckACL         = flag.Bool("checkacl", false, "calc all bucket size")
	argsTerminate       = flag.Bool("terminate", false, "Instance terminate")
	argRegister         = flag.Bool("register", false, "Register Instances to ELB")
	argDeregister       = flag.Bool("deregister", false, "Deregister Instances to ELB")
	argMax              = flag.Bool("max", false, "input maxsize num")
	argUpdate           = flag.Bool("update", false, "input laucnconfig name")
	argMin              = flag.Bool("min", false, "input minsize num")
	argDesire           = flag.Bool("desire", false, "input desiresize num")
)

func main() {
	flag.Parse()

	ec2Client := clitoolgoaws.AwsEC2Client(*argProfile, *argRegion)
	rdsClient := clitoolgoaws.AwsRDSClient(*argProfile, *argRegion)
	elbClient := clitoolgoaws.AwsELBClient(*argProfile, *argRegion)
	cloudwatchClient := clitoolgoaws.AwsCloudwatchClient(*argProfile, *argRegion)
	kinesisClient := clitoolgoaws.AwsKinesisClient(*argProfile, *argRegion)
	iamClient := clitoolgoaws.AwsIAMClient(*argProfile, *argRegion)
	S3Client := clitoolgoaws.AwsS3Client(*argProfile, *argRegion)
	asClient := clitoolgoaws.AwsASClient(*argProfile, *argRegion)
	route53Client := clitoolgoaws.AwsRoute53Client(*argProfile, *argRegion)

	// EC2のコマンド
	var ec2Instances []*string
	var ec2InstancesAMI *string
	exeFlag := true
	if *argResource == "ec2" {
		if *argAMIList {
			clitoolgoaws.ListAMI(ec2Client, nil)
			exeFlag = false
		} else if *argDeregister {
			clitoolgoaws.DeregisterAMI(ec2Client, argAmiId)
			exeFlag = false
		} else if *argEIPList {
			clitoolgoaws.ShowElasticIP(ec2Client)
			exeFlag = false
		} else if *argEIPDelete {
			clitoolgoaws.DeleteElasticIP(ec2Client, argAllocationId)
			exeFlag = false
		} else if *argSecurityGroup {
			clitoolgoaws.ListSecurityGroup(ec2Client)
			exeFlag = false
		} else if *argAMItest {
			clitoolgoaws.GetAmiId(ec2Client, *argImageName)
			exeFlag = false
		} else if *argShow {
			sliceSGInfo := []*string{
				argSGId,
			}
			clitoolgoaws.ShowSecurityGroup(ec2Client, sliceSGInfo)
			exeFlag = false
		}

		if *argInstances != "" {
			ec2Instances = clitoolgoaws.GetEC2InstanceIds(ec2Client, *argInstances)
			if *argStart {
				clitoolgoaws.ControlEC2Instances(ec2Client, ec2Instances, "start")
			} else if *argStop {
				clitoolgoaws.ControlEC2Instances(ec2Client, ec2Instances, "stop")
			} else if *argsTerminate {
				clitoolgoaws.ControlEC2Instances(ec2Client, ec2Instances, "terminate")
			} else if *argAMI {
				ec2InstancesAMI = clitoolgoaws.GetEC2InstanceIdsAMI(ec2Client, *argInstances)
				clitoolgoaws.RegisterAMI(ec2Client, argAmiName, ec2InstancesAMI)
			} else {
				fmt.Println("`-start` or `-stop` or `-terminate` or `-ami` slect option")
				os.Exit(1)
			}
		} else if exeFlag {
			clitoolgoaws.ListEC2Instances(ec2Client, nil)
		}
	}
	// LaunchConfigのコマンド
	if *argResource == "lc" {
		if *argLaunchConfigName != "" {
			sliceSGId := []*string{
				argSecurityGroupId,
			}
			argImageId = clitoolgoaws.GetAmiId(ec2Client, *argImageName)
			clitoolgoaws.CreateLaunchConfig(asClient, argLaunchConfigName, argInstanceProfile, argImageId, argInstanceType, argKeyName, sliceSGId)
		}
	}

	//AutoScalingのコマンド
	if *argResource == "as" {
		if *argASG != "" {
			// -update
			if *argMax {
				clitoolgoaws.ChangeMaxSizeInstances(asClient, argASG, argNum)
			} else if *argMin {
				clitoolgoaws.ChangeMinSizeInstances(asClient, argASG, argNum)
			} else if *argDesire {
				clitoolgoaws.ChangeDesireSizeInstances(asClient, argASG, argNum)
			}
			if *argUpdate {
				clitoolgoaws.ChangeLaunchConfig(asClient, argASG, argLaunchConfigName)
			}
		} else {
			clitoolgoaws.ShowAutoScaling(asClient)
		}
	}

	// route53のコマンド
	if *argResource == "route53" {
		if *argShow {
			clitoolgoaws.ShowListResourceRecordSets(route53Client, argZoneId)
		} else {
			clitoolgoaws.ShowHostedZone(route53Client)
		}
	}

	// RDSのコマンド
	var rdsInstances *string
	if *argResource == "rds" {
		if *argInstances != "" {
			rdsInstances = clitoolgoaws.GetRDSInstanceName(rdsClient, *argInstances)
			if *argStart {
				clitoolgoaws.ControlRDSInstances(rdsClient, rdsInstances, "start")
			} else if *argStop {
				clitoolgoaws.ControlRDSInstances(rdsClient, rdsInstances, "stop")
			} else {
				fmt.Println("`-start` or `-stop` slect option")
				os.Exit(1)
			}
		} else {
			clitoolgoaws.ListRDSInstances(rdsClient, nil)
		}
	}

	// S3のコマンド
	exeFlagS3 := true
	if *argResource == "s3" {
		if *argSizeAll {
			start := time.Now()                       //  時間計測
			clitoolgoaws.TotalGetBucketSize(S3Client) // バケットの数だけスレッドを設けて計算する
			//clitoolgoaws.TotalGetBucketSizeSingle(S3Client) //シングルスレッドで計算する
			exeFlagS3 = false
			end := time.Now() // 時間計測
			fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
		} else if *argCheckACL {
			// バケットのACL情報のレスポンスを確認するため
			clitoolgoaws.ShowPublicBucket(S3Client)
			exeFlagS3 = false
		}

		if *argBucket != "" {
			if *argShow {
				clitoolgoaws.ShowObjects(S3Client, argBucket)
			} else if *argSize {
				clitoolgoaws.ShowBucketSize(S3Client, argBucket)
			} else if *argBucketDelete {
				clitoolgoaws.DeleteBucket(S3Client, argBucket)
			} else if *argObjectDelete {
				clitoolgoaws.DeleteObject(S3Client, argBucket, argObject)
			} else if *argObjectsAllDelete {
				clitoolgoaws.DeleteAllObjects(S3Client, argBucket)
			}

		} else if exeFlagS3 {
			clitoolgoaws.ShowBuckets(S3Client)
		}
	}

	// ELBのコマンド
	var elasticLoadbalancers []*string
	if *argResource == "elb" {
		if *argELBName != "" {
			elasticLoadbalancers = clitoolgoaws.GetELBInfo(elbClient, *argELBName)
			if *argShow {
				clitoolgoaws.ListELBBackendInstances(elbClient, elasticLoadbalancers, "show")
			} else if *argRegister && *argInstances != "" {
				clitoolgoaws.ControlELB(elbClient, *argELBName, *argInstances, "register")
				clitoolgoaws.ListELBBackendInstances(elbClient, elasticLoadbalancers, "show")
			} else if *argDeregister && *argInstances != "" {
				clitoolgoaws.ControlELB(elbClient, *argELBName, *argInstances, "deregister")
				clitoolgoaws.ListELBBackendInstances(elbClient, elasticLoadbalancers, "show")
			} else {
				fmt.Println("`-show` slect option")
				os.Exit(1)
			}
		} else {
			clitoolgoaws.ListELB(elbClient, nil)
		}
	}

	// Cloudwatchのコマンド
	if *argResource == "cloudwatch" {
		if *argBilling {
			clitoolgoaws.GetBilling(*argProfile, BILLING_REGION)
		} else {
			clitoolgoaws.ListCloudwatch(cloudwatchClient, nil)
		}
	}

	// Kinesisのコマンド
	if *argResource == "kinesis" {
		clitoolgoaws.ListKinesis(kinesisClient, nil)
	}

	// IAM-Userコマンド
	if *argResource == "iam-user" {
		clitoolgoaws.ListIAMUser(iamClient, nil)
	}
	// IAM-Groupコマンド
	if *argResource == "iam-group" {
		clitoolgoaws.ListIAMGroup(iamClient, nil)
	}

}
