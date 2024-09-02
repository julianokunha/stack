// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type V2ListWorkflowsResponseCursor struct {
	Data     []V2Workflow `json:"data"`
	HasMore  bool         `json:"hasMore"`
	Next     *string      `json:"next,omitempty"`
	PageSize int64        `json:"pageSize"`
	Previous *string      `json:"previous,omitempty"`
}

func (o *V2ListWorkflowsResponseCursor) GetData() []V2Workflow {
	if o == nil {
		return []V2Workflow{}
	}
	return o.Data
}

func (o *V2ListWorkflowsResponseCursor) GetHasMore() bool {
	if o == nil {
		return false
	}
	return o.HasMore
}

func (o *V2ListWorkflowsResponseCursor) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *V2ListWorkflowsResponseCursor) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *V2ListWorkflowsResponseCursor) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

type V2ListWorkflowsResponse struct {
	Cursor V2ListWorkflowsResponseCursor `json:"cursor"`
}

func (o *V2ListWorkflowsResponse) GetCursor() V2ListWorkflowsResponseCursor {
	if o == nil {
		return V2ListWorkflowsResponseCursor{}
	}
	return o.Cursor
}
