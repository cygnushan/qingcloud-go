package client

import (
	"fmt"
	"time"

	"github.com/chai2010/qingcloud-go/pkg/logger"
	service "github.com/chai2010/qingcloud-go/pkg/service.pb"
	"github.com/chai2010/qingcloud-go/pkg/utils"
)

// WaitJob wait the job with this jobID finish
func WaitJob(jobService *service.JobService, jobID string, timeout time.Duration, waitInterval time.Duration) error {
	logger.Debug("Waiting for Job [%s] finished", jobID)
	return utils.WaitForSpecificOrError(func() (bool, error) {
		input := &service.DescribeJobsInput{Jobs: []string{jobID}}
		output, err := jobService.DescribeJobs(input)
		if err != nil {
			//network or api error, not considered job fail.
			return false, nil
		}
		if len(output.JobSet) == 0 {
			return false, fmt.Errorf("Can not find job [%s]", jobID)
		}
		j := output.JobSet[0]
		if j.GetStatus() == "" {
			logger.Error("Job [%s] status is nil ", jobID)
			return false, nil
		}
		if j.GetStatus() == "working" || j.GetStatus() == "pending" {
			return false, nil
		}
		if j.GetStatus() == "successful" {
			return true, nil
		}
		if j.GetStatus() == "failed" {
			return false, fmt.Errorf("Job [%s] failed", jobID)
		}
		logger.Error("Unknow status [%s] for job [%s]", j.Status, jobID)
		return false, nil
	}, timeout, waitInterval)
}

// CheckJobStatus get job status
func CheckJobStatus(jobService *service.JobService, jobID string) (string, error) {
	input := &service.DescribeJobsInput{Jobs: []string{jobID}}
	output, err := jobService.DescribeJobs(input)
	if err != nil {
		return JobStatusUnknown, nil
	}
	if len(output.JobSet) == 0 {
		return "", fmt.Errorf("Can not find job [%s]", jobID)
	}
	j := output.JobSet[0]
	if j.GetStatus() == "" {
		logger.Error("Job [%s] status is nil ", jobID)
		return JobStatusUnknown, nil
	}
	return j.GetStatus(), nil
}

func describeInstance(instanceService *service.InstanceService, instanceID string) (*service.Instance, error) {
	input := &service.DescribeInstancesInput{Instances: []string{instanceID}}
	output, err := instanceService.DescribeInstances(input)
	if err != nil {
		return nil, err
	}
	if len(output.InstanceSet) == 0 {
		return nil, fmt.Errorf("Instance with id [%s] not exist", instanceID)
	}
	return output.InstanceSet[0], nil
}

// WaitInstanceStatus wait the instance with this instanceID to expect status
func WaitInstanceStatus(instanceService *service.InstanceService, instanceID string, status string, timeout time.Duration, waitInterval time.Duration) (ins *service.Instance, err error) {
	logger.Debug("Waiting for Instance [%s] status [%s] ", instanceID, status)
	errorTimes := 0
	err = utils.WaitForSpecificOrError(func() (bool, error) {
		i, err := describeInstance(instanceService, instanceID)
		if err != nil {
			logger.Error("DescribeInstance [%s] error : [%s]", instanceID, err.Error())
			errorTimes++
			if errorTimes > 3 {
				return false, err
			}
			return false, nil
		}
		if i.GetStatus() != "" && i.GetStatus() == status {
			if i.GetTransitionStatus() != "" {
				//wait transition to finished
				return false, nil
			}
			logger.Debug("Instance [%s] status is [%s] ", instanceID, i.Status)
			ins = i
			return true, nil
		}
		return false, nil
	}, timeout, waitInterval)
	return
}

// WaitInstanceNetwork wait the instance with this instanceID network become ready
func WaitInstanceNetwork(instanceService *service.InstanceService, instanceID string, timeout time.Duration, waitInterval time.Duration) (ins *service.Instance, err error) {
	logger.Debug("Waiting for IP address to be assigned to Instance [%s]", instanceID)
	err = utils.WaitForSpecificOrError(func() (bool, error) {
		i, err := describeInstance(instanceService, instanceID)
		if err != nil {
			return false, err
		}
		if len(i.Vxnets) == 0 || i.Vxnets[0] == nil || i.Vxnets[0].GetPrivateIp() == "" {
			return false, nil
		}
		ins = i
		logger.Debug("Instance [%s] get IP address [%s]", instanceID, ins.Vxnets[0].PrivateIp)
		return true, nil
	}, timeout, waitInterval)
	return
}

func describeLoadBalancer(lbService *service.LoadBalancerService, loadBalancerID string) (*service.LoadBalancer, error) {
	output, err := lbService.DescribeLoadBalancers(&service.DescribeLoadBalancersInput{
		Loadbalancers: []string{loadBalancerID},
	})
	if err != nil {
		return nil, err
	}
	if len(output.LoadbalancerSet) == 0 {
		return nil, nil
	}

	return output.LoadbalancerSet[0], nil
}

// WaitLoadBalancerStatus wait the loadBalancer with this loadBalancerID to expect status
func WaitLoadBalancerStatus(lbService *service.LoadBalancerService, loadBalancerID string, status string, timeout time.Duration, waitInterval time.Duration) (lb *service.LoadBalancer, err error) {
	logger.Debug("Waiting for LoadBalancer [%s] status [%s] ", loadBalancerID, status)
	errorTimes := 0
	err = utils.WaitForSpecificOrError(func() (bool, error) {
		i, err := describeLoadBalancer(lbService, loadBalancerID)
		if err != nil {
			logger.Error("DescribeLoadBalancer [%s] error : [%s]", loadBalancerID, err.Error())
			errorTimes++
			if errorTimes > 3 {
				return false, err
			}
			return false, nil
		}
		if i.GetStatus() != "" && i.GetStatus() == status {
			if i.GetTransitionStatus() != "" {
				//wait transition to finished
				return false, nil
			}
			lb = i
			logger.Debug("LoadBalancer [%s] status is [%s] ", loadBalancerID, i.Status)
			return true, nil
		}
		return false, nil
	}, timeout, waitInterval)
	return
}
