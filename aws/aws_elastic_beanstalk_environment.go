// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

func ListElasticBeanstalkEnvironment(client *Client) ([]Resource, error) {
	req := client.elasticbeanstalkconn.DescribeEnvironmentsRequest(&elasticbeanstalk.DescribeEnvironmentsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Environments) > 0 {
		for _, r := range resp.Environments {

			result = append(result, Resource{
				Type:   "aws_elastic_beanstalk_environment",
				ID:     *r.EnvironmentId,
				Region: client.elasticbeanstalkconn.Config.Region,
			})
		}
	}

	return result, nil
}
