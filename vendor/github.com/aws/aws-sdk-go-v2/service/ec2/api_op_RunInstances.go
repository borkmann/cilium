// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RunInstancesInput struct {
	_ struct{} `type:"structure"`

	// Reserved.
	AdditionalInfo *string `locationName:"additionalInfo" type:"string"`

	// The block device mapping entries.
	BlockDeviceMappings []BlockDeviceMapping `locationName:"BlockDeviceMapping" locationNameList:"BlockDeviceMapping" type:"list"`

	// Information about the Capacity Reservation targeting option. If you do not
	// specify this parameter, the instance's Capacity Reservation preference defaults
	// to open, which enables it to run in any open Capacity Reservation that has
	// matching attributes (instance type, platform, Availability Zone).
	CapacityReservationSpecification *CapacityReservationSpecification `type:"structure"`

	// Unique, case-sensitive identifier you provide to ensure the idempotency of
	// the request. For more information, see Ensuring Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	//
	// Constraints: Maximum 64 ASCII characters
	ClientToken *string `locationName:"clientToken" type:"string"`

	// The CPU options for the instance. For more information, see Optimizing CPU
	// Options (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	CpuOptions *CpuOptionsRequest `type:"structure"`

	// The credit option for CPU usage of the burstable performance instance. Valid
	// values are standard and unlimited. To change this attribute after launch,
	// use ModifyInstanceCreditSpecification (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceCreditSpecification.html).
	// For more information, see Burstable Performance Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// Default: standard (T2 instances) or unlimited (T3/T3a instances)
	CreditSpecification *CreditSpecificationRequest `type:"structure"`

	// If you set this parameter to true, you can't terminate the instance using
	// the Amazon EC2 console, CLI, or API; otherwise, you can. To change this attribute
	// after launch, use ModifyInstanceAttribute (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceAttribute.html).
	// Alternatively, if you set InstanceInitiatedShutdownBehavior to terminate,
	// you can terminate the instance by running the shutdown command from the instance.
	//
	// Default: false
	DisableApiTermination *bool `locationName:"disableApiTermination" type:"boolean"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// Indicates whether the instance is optimized for Amazon EBS I/O. This optimization
	// provides dedicated throughput to Amazon EBS and an optimized configuration
	// stack to provide optimal Amazon EBS I/O performance. This optimization isn't
	// available with all instance types. Additional usage charges apply when using
	// an EBS-optimized instance.
	//
	// Default: false
	EbsOptimized *bool `locationName:"ebsOptimized" type:"boolean"`

	// An elastic GPU to associate with the instance. An Elastic GPU is a GPU resource
	// that you can attach to your Windows instance to accelerate the graphics performance
	// of your applications. For more information, see Amazon EC2 Elastic GPUs (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-graphics.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	ElasticGpuSpecification []ElasticGpuSpecification `locationNameList:"item" type:"list"`

	// An elastic inference accelerator to associate with the instance. Elastic
	// inference accelerators are a resource you can attach to your Amazon EC2 instances
	// to accelerate your Deep Learning (DL) inference workloads.
	ElasticInferenceAccelerators []ElasticInferenceAccelerator `locationName:"ElasticInferenceAccelerator" locationNameList:"item" type:"list"`

	// Indicates whether an instance is enabled for hibernation. For more information,
	// see Hibernate Your Instance (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	HibernationOptions *HibernationOptionsRequest `type:"structure"`

	// The IAM instance profile.
	IamInstanceProfile *IamInstanceProfileSpecification `locationName:"iamInstanceProfile" type:"structure"`

	// The ID of the AMI. An AMI ID is required to launch an instance and must be
	// specified here or in a launch template.
	ImageId *string `type:"string"`

	// Indicates whether an instance stops or terminates when you initiate shutdown
	// from the instance (using the operating system command for system shutdown).
	//
	// Default: stop
	InstanceInitiatedShutdownBehavior ShutdownBehavior `locationName:"instanceInitiatedShutdownBehavior" type:"string" enum:"true"`

	// The market (purchasing) option for the instances.
	//
	// For RunInstances, persistent Spot Instance requests are only supported when
	// InstanceInterruptionBehavior is set to either hibernate or stop.
	InstanceMarketOptions *InstanceMarketOptionsRequest `type:"structure"`

	// The instance type. For more information, see Instance Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// Default: m1.small
	InstanceType InstanceType `type:"string" enum:"true"`

	// [EC2-VPC] The number of IPv6 addresses to associate with the primary network
	// interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
	// You cannot specify this option and the option to assign specific IPv6 addresses
	// in the same request. You can specify this option if you've specified a minimum
	// number of instances to launch.
	//
	// You cannot specify this option and the network interfaces option in the same
	// request.
	Ipv6AddressCount *int64 `type:"integer"`

	// [EC2-VPC] The IPv6 addresses from the range of the subnet to associate with
	// the primary network interface. You cannot specify this option and the option
	// to assign a number of IPv6 addresses in the same request. You cannot specify
	// this option if you've specified a minimum number of instances to launch.
	//
	// You cannot specify this option and the network interfaces option in the same
	// request.
	Ipv6Addresses []InstanceIpv6Address `locationName:"Ipv6Address" locationNameList:"item" type:"list"`

	// The ID of the kernel.
	//
	// We recommend that you use PV-GRUB instead of kernels and RAM disks. For more
	// information, see PV-GRUB (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	KernelId *string `type:"string"`

	// The name of the key pair. You can create a key pair using CreateKeyPair (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateKeyPair.html)
	// or ImportKeyPair (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportKeyPair.html).
	//
	// If you do not specify a key pair, you can't connect to the instance unless
	// you choose an AMI that is configured to allow users another way to log in.
	KeyName *string `type:"string"`

	// The launch template to use to launch the instances. Any parameters that you
	// specify in RunInstances override the same parameters in the launch template.
	// You can specify either the name or ID of a launch template, but not both.
	LaunchTemplate *LaunchTemplateSpecification `type:"structure"`

	// The license configurations.
	LicenseSpecifications []LicenseConfigurationRequest `locationName:"LicenseSpecification" locationNameList:"item" type:"list"`

	// The maximum number of instances to launch. If you specify more instances
	// than Amazon EC2 can launch in the target Availability Zone, Amazon EC2 launches
	// the largest possible number of instances above MinCount.
	//
	// Constraints: Between 1 and the maximum number you're allowed for the specified
	// instance type. For more information about the default limits, and how to
	// request an increase, see How many instances can I run in Amazon EC2 (http://aws.amazon.com/ec2/faqs/#How_many_instances_can_I_run_in_Amazon_EC2)
	// in the Amazon EC2 FAQ.
	//
	// MaxCount is a required field
	MaxCount *int64 `type:"integer" required:"true"`

	// The metadata options for the instance. For more information, see Instance
	// Metadata and User Data (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html).
	MetadataOptions *InstanceMetadataOptionsRequest `type:"structure"`

	// The minimum number of instances to launch. If you specify a minimum that
	// is more instances than Amazon EC2 can launch in the target Availability Zone,
	// Amazon EC2 launches no instances.
	//
	// Constraints: Between 1 and the maximum number you're allowed for the specified
	// instance type. For more information about the default limits, and how to
	// request an increase, see How many instances can I run in Amazon EC2 (http://aws.amazon.com/ec2/faqs/#How_many_instances_can_I_run_in_Amazon_EC2)
	// in the Amazon EC2 General FAQ.
	//
	// MinCount is a required field
	MinCount *int64 `type:"integer" required:"true"`

	// Specifies whether detailed monitoring is enabled for the instance.
	Monitoring *RunInstancesMonitoringEnabled `type:"structure"`

	// The network interfaces to associate with the instance. If you specify a network
	// interface, you must specify any security groups and subnets as part of the
	// network interface.
	NetworkInterfaces []InstanceNetworkInterfaceSpecification `locationName:"networkInterface" locationNameList:"item" type:"list"`

	// The placement for the instance.
	Placement *Placement `type:"structure"`

	// [EC2-VPC] The primary IPv4 address. You must specify a value from the IPv4
	// address range of the subnet.
	//
	// Only one private IP address can be designated as primary. You can't specify
	// this option if you've specified the option to designate a private IP address
	// as the primary IP address in a network interface specification. You cannot
	// specify this option if you're launching more than one instance in the request.
	//
	// You cannot specify this option and the network interfaces option in the same
	// request.
	PrivateIpAddress *string `locationName:"privateIpAddress" type:"string"`

	// The ID of the RAM disk to select. Some kernels require additional drivers
	// at launch. Check the kernel requirements for information about whether you
	// need to specify a RAM disk. To find kernel requirements, go to the AWS Resource
	// Center and search for the kernel ID.
	//
	// We recommend that you use PV-GRUB instead of kernels and RAM disks. For more
	// information, see PV-GRUB (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	RamdiskId *string `type:"string"`

	// The IDs of the security groups. You can create a security group using CreateSecurityGroup
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateSecurityGroup.html).
	//
	// If you specify a network interface, you must specify any security groups
	// as part of the network interface.
	SecurityGroupIds []string `locationName:"SecurityGroupId" locationNameList:"SecurityGroupId" type:"list"`

	// [EC2-Classic, default VPC] The names of the security groups. For a nondefault
	// VPC, you must use security group IDs instead.
	//
	// If you specify a network interface, you must specify any security groups
	// as part of the network interface.
	//
	// Default: Amazon EC2 uses the default security group.
	SecurityGroups []string `locationName:"SecurityGroup" locationNameList:"SecurityGroup" type:"list"`

	// [EC2-VPC] The ID of the subnet to launch the instance into.
	//
	// If you specify a network interface, you must specify any subnets as part
	// of the network interface.
	SubnetId *string `type:"string"`

	// The tags to apply to the resources during launch. You can only tag instances
	// and volumes on launch. The specified tags are applied to all instances or
	// volumes that are created during launch. To tag a resource after it has been
	// created, see CreateTags (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTags.html).
	TagSpecifications []TagSpecification `locationName:"TagSpecification" locationNameList:"item" type:"list"`

	// The user data to make available to the instance. For more information, see
	// Running Commands on Your Linux Instance at Launch (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html)
	// (Linux) and Adding User Data (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-instance-metadata.html#instancedata-add-user-data)
	// (Windows). If you are using a command line tool, base64-encoding is performed
	// for you, and you can load the text from a file. Otherwise, you must provide
	// base64-encoded text. User data is limited to 16 KB.
	UserData *string `type:"string"`
}

// String returns the string representation
func (s RunInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RunInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RunInstancesInput"}

	if s.MaxCount == nil {
		invalidParams.Add(aws.NewErrParamRequired("MaxCount"))
	}

	if s.MinCount == nil {
		invalidParams.Add(aws.NewErrParamRequired("MinCount"))
	}
	if s.CreditSpecification != nil {
		if err := s.CreditSpecification.Validate(); err != nil {
			invalidParams.AddNested("CreditSpecification", err.(aws.ErrInvalidParams))
		}
	}
	if s.ElasticGpuSpecification != nil {
		for i, v := range s.ElasticGpuSpecification {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ElasticGpuSpecification", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.ElasticInferenceAccelerators != nil {
		for i, v := range s.ElasticInferenceAccelerators {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ElasticInferenceAccelerators", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Monitoring != nil {
		if err := s.Monitoring.Validate(); err != nil {
			invalidParams.AddNested("Monitoring", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes a reservation.
type RunInstancesOutput struct {
	_ struct{} `type:"structure"`

	// [EC2-Classic only] The security groups.
	Groups []GroupIdentifier `locationName:"groupSet" locationNameList:"item" type:"list"`

	// The instances.
	Instances []Instance `locationName:"instancesSet" locationNameList:"item" type:"list"`

	// The ID of the AWS account that owns the reservation.
	OwnerId *string `locationName:"ownerId" type:"string"`

	// The ID of the requester that launched the instances on your behalf (for example,
	// AWS Management Console or Auto Scaling).
	RequesterId *string `locationName:"requesterId" type:"string"`

	// The ID of the reservation.
	ReservationId *string `locationName:"reservationId" type:"string"`
}

// String returns the string representation
func (s RunInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opRunInstances = "RunInstances"

// RunInstancesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Launches the specified number of instances using an AMI for which you have
// permissions.
//
// You can specify a number of options, or leave the default options. The following
// rules apply:
//
//    * [EC2-VPC] If you don't specify a subnet ID, we choose a default subnet
//    from your default VPC for you. If you don't have a default VPC, you must
//    specify a subnet ID in the request.
//
//    * [EC2-Classic] If don't specify an Availability Zone, we choose one for
//    you.
//
//    * Some instance types must be launched into a VPC. If you do not have
//    a default VPC, or if you do not specify a subnet ID, the request fails.
//    For more information, see Instance Types Available Only in a VPC (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-vpc.html#vpc-only-instance-types).
//
//    * [EC2-VPC] All instances have a network interface with a primary private
//    IPv4 address. If you don't specify this address, we choose one from the
//    IPv4 range of your subnet.
//
//    * Not all instance types support IPv6 addresses. For more information,
//    see Instance Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html).
//
//    * If you don't specify a security group ID, we use the default security
//    group. For more information, see Security Groups (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html).
//
//    * If any of the AMIs have a product code attached for which the user has
//    not subscribed, the request fails.
//
// You can create a launch template (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html),
// which is a resource that contains the parameters to launch an instance. When
// you launch an instance using RunInstances, you can specify the launch template
// instead of specifying the launch parameters.
//
// To ensure faster instance launches, break up large requests into smaller
// batches. For example, create five separate launch requests for 100 instances
// each instead of one launch request for 500 instances.
//
// An instance is ready for you to use when it's in the running state. You can
// check the state of your instance using DescribeInstances. You can tag instances
// and EBS volumes during launch, after launch, or both. For more information,
// see CreateTags and Tagging Your Amazon EC2 Resources (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html).
//
// Linux instances have access to the public key of the key pair at boot. You
// can use this key to provide secure access to the instance. Amazon EC2 public
// images use this feature to provide secure access without passwords. For more
// information, see Key Pairs (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// For troubleshooting, see What To Do If An Instance Immediately Terminates
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_InstanceStraightToTerminated.html),
// and Troubleshooting Connecting to Your Instance (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstancesConnecting.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using RunInstancesRequest.
//    req := client.RunInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RunInstances
func (c *Client) RunInstancesRequest(input *RunInstancesInput) RunInstancesRequest {
	op := &aws.Operation{
		Name:       opRunInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RunInstancesInput{}
	}

	req := c.newRequest(op, input, &RunInstancesOutput{})
	return RunInstancesRequest{Request: req, Input: input, Copy: c.RunInstancesRequest}
}

// RunInstancesRequest is the request type for the
// RunInstances API operation.
type RunInstancesRequest struct {
	*aws.Request
	Input *RunInstancesInput
	Copy  func(*RunInstancesInput) RunInstancesRequest
}

// Send marshals and sends the RunInstances API request.
func (r RunInstancesRequest) Send(ctx context.Context) (*RunInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RunInstancesResponse{
		RunInstancesOutput: r.Request.Data.(*RunInstancesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RunInstancesResponse is the response type for the
// RunInstances API operation.
type RunInstancesResponse struct {
	*RunInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RunInstances request.
func (r *RunInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
