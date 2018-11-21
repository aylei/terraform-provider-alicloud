package drds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyFullTableScan invokes the drds.ModifyFullTableScan API synchronously
// api document: https://help.aliyun.com/api/drds/modifyfulltablescan.html
func (client *Client) ModifyFullTableScan(request *ModifyFullTableScanRequest) (response *ModifyFullTableScanResponse, err error) {
	response = CreateModifyFullTableScanResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyFullTableScanWithChan invokes the drds.ModifyFullTableScan API asynchronously
// api document: https://help.aliyun.com/api/drds/modifyfulltablescan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyFullTableScanWithChan(request *ModifyFullTableScanRequest) (<-chan *ModifyFullTableScanResponse, <-chan error) {
	responseChan := make(chan *ModifyFullTableScanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyFullTableScan(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyFullTableScanWithCallback invokes the drds.ModifyFullTableScan API asynchronously
// api document: https://help.aliyun.com/api/drds/modifyfulltablescan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyFullTableScanWithCallback(request *ModifyFullTableScanRequest, callback func(response *ModifyFullTableScanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyFullTableScanResponse
		var err error
		defer close(result)
		response, err = client.ModifyFullTableScan(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyFullTableScanRequest is the request struct for api ModifyFullTableScan
type ModifyFullTableScanRequest struct {
	*requests.RpcRequest
	DbName         string           `position:"Query" name:"DbName"`
	TableNames     string           `position:"Query" name:"TableNames"`
	DrdsInstanceId string           `position:"Query" name:"DrdsInstanceId"`
	FullTableScan  requests.Boolean `position:"Query" name:"FullTableScan"`
}

// ModifyFullTableScanResponse is the response struct for api ModifyFullTableScan
type ModifyFullTableScanResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateModifyFullTableScanRequest creates a request to invoke ModifyFullTableScan API
func CreateModifyFullTableScanRequest() (request *ModifyFullTableScanRequest) {
	request = &ModifyFullTableScanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2017-10-16", "ModifyFullTableScan", "", "")
	return
}

// CreateModifyFullTableScanResponse creates a response to parse from ModifyFullTableScan response
func CreateModifyFullTableScanResponse() (response *ModifyFullTableScanResponse) {
	response = &ModifyFullTableScanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}