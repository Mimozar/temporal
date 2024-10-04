// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by cmd/tools/genrpcserverinterceptors. DO NOT EDIT.

package logtags

import (
	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common/log/tag"
)

func (wt *WorkflowTags) extractFromMatchingServiceServerRequest(req any) []tag.Tag {
	switch r := req.(type) {
	case *matchingservice.AddActivityTaskRequest:
		return []tag.Tag{
			tag.WorkflowID(r.GetExecution().GetWorkflowId()),
			tag.WorkflowRunID(r.GetExecution().GetRunId()),
		}
	case *matchingservice.AddWorkflowTaskRequest:
		return []tag.Tag{
			tag.WorkflowID(r.GetExecution().GetWorkflowId()),
			tag.WorkflowRunID(r.GetExecution().GetRunId()),
		}
	case *matchingservice.ApplyTaskQueueUserDataReplicationEventRequest:
		return nil
	case *matchingservice.CancelOutstandingPollRequest:
		return nil
	case *matchingservice.CreateNexusEndpointRequest:
		return nil
	case *matchingservice.DeleteNexusEndpointRequest:
		return nil
	case *matchingservice.DescribeTaskQueueRequest:
		return nil
	case *matchingservice.DescribeTaskQueuePartitionRequest:
		return nil
	case *matchingservice.DispatchNexusTaskRequest:
		return nil
	case *matchingservice.ForceLoadTaskQueuePartitionRequest:
		return nil
	case *matchingservice.ForceUnloadTaskQueuePartitionRequest:
		return nil
	case *matchingservice.GetBuildIdTaskQueueMappingRequest:
		return nil
	case *matchingservice.GetTaskQueueUserDataRequest:
		return nil
	case *matchingservice.GetWorkerBuildIdCompatibilityRequest:
		return nil
	case *matchingservice.GetWorkerVersioningRulesRequest:
		return nil
	case *matchingservice.ListNexusEndpointsRequest:
		return nil
	case *matchingservice.ListTaskQueuePartitionsRequest:
		return nil
	case *matchingservice.PollActivityTaskQueueRequest:
		return nil
	case *matchingservice.PollNexusTaskQueueRequest:
		return nil
	case *matchingservice.PollWorkflowTaskQueueRequest:
		return nil
	case *matchingservice.QueryWorkflowRequest:
		return []tag.Tag{
			tag.WorkflowID(r.GetQueryRequest().GetExecution().GetWorkflowId()),
			tag.WorkflowRunID(r.GetQueryRequest().GetExecution().GetRunId()),
		}
	case *matchingservice.ReplicateTaskQueueUserDataRequest:
		return nil
	case *matchingservice.RespondNexusTaskCompletedRequest:
		return nil
	case *matchingservice.RespondNexusTaskFailedRequest:
		return nil
	case *matchingservice.RespondQueryTaskCompletedRequest:
		return nil
	case *matchingservice.UpdateNexusEndpointRequest:
		return nil
	case *matchingservice.UpdateTaskQueueUserDataRequest:
		return nil
	case *matchingservice.UpdateWorkerBuildIdCompatibilityRequest:
		return nil
	case *matchingservice.UpdateWorkerVersioningRulesRequest:
		return nil
	default:
		return nil
	}
}
