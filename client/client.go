// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddRepositoryMemberRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	AccessLevel    *int32  `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	AliyunPks      *string `json:"aliyunPks,omitempty" xml:"aliyunPks,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s AddRepositoryMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRepositoryMemberRequest) GoString() string {
	return s.String()
}

func (s *AddRepositoryMemberRequest) SetAccessToken(v string) *AddRepositoryMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *AddRepositoryMemberRequest) SetAccessLevel(v int32) *AddRepositoryMemberRequest {
	s.AccessLevel = &v
	return s
}

func (s *AddRepositoryMemberRequest) SetAliyunPks(v string) *AddRepositoryMemberRequest {
	s.AliyunPks = &v
	return s
}

func (s *AddRepositoryMemberRequest) SetOrganizationId(v string) *AddRepositoryMemberRequest {
	s.OrganizationId = &v
	return s
}

type AddRepositoryMemberResponseBody struct {
	ErrorCode    *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*AddRepositoryMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRepositoryMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRepositoryMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddRepositoryMemberResponseBody) SetErrorCode(v string) *AddRepositoryMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddRepositoryMemberResponseBody) SetErrorMessage(v string) *AddRepositoryMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddRepositoryMemberResponseBody) SetRequestId(v string) *AddRepositoryMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRepositoryMemberResponseBody) SetResult(v []*AddRepositoryMemberResponseBodyResult) *AddRepositoryMemberResponseBody {
	s.Result = v
	return s
}

func (s *AddRepositoryMemberResponseBody) SetSuccess(v bool) *AddRepositoryMemberResponseBody {
	s.Success = &v
	return s
}

type AddRepositoryMemberResponseBodyResult struct {
	AccessLevel  *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s AddRepositoryMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddRepositoryMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddRepositoryMemberResponseBodyResult) SetAccessLevel(v int32) *AddRepositoryMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetAvatarUrl(v string) *AddRepositoryMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetEmail(v string) *AddRepositoryMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetExternUserId(v string) *AddRepositoryMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetId(v int64) *AddRepositoryMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetState(v string) *AddRepositoryMemberResponseBodyResult {
	s.State = &v
	return s
}

type AddRepositoryMemberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddRepositoryMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddRepositoryMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRepositoryMemberResponse) GoString() string {
	return s.String()
}

func (s *AddRepositoryMemberResponse) SetHeaders(v map[string]*string) *AddRepositoryMemberResponse {
	s.Headers = v
	return s
}

func (s *AddRepositoryMemberResponse) SetStatusCode(v int32) *AddRepositoryMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRepositoryMemberResponse) SetBody(v *AddRepositoryMemberResponseBody) *AddRepositoryMemberResponse {
	s.Body = v
	return s
}

type AddWebhookRequest struct {
	AccessToken           *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId        *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Description           *string `json:"description,omitempty" xml:"description,omitempty"`
	EnableSslVerification *bool   `json:"enableSslVerification,omitempty" xml:"enableSslVerification,omitempty"`
	MergeRequestsEvents   *bool   `json:"mergeRequestsEvents,omitempty" xml:"mergeRequestsEvents,omitempty"`
	NoteEvents            *bool   `json:"noteEvents,omitempty" xml:"noteEvents,omitempty"`
	PushEvents            *bool   `json:"pushEvents,omitempty" xml:"pushEvents,omitempty"`
	SecretToken           *string `json:"secretToken,omitempty" xml:"secretToken,omitempty"`
	TagPushEvents         *bool   `json:"tagPushEvents,omitempty" xml:"tagPushEvents,omitempty"`
	Url                   *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s AddWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s AddWebhookRequest) GoString() string {
	return s.String()
}

func (s *AddWebhookRequest) SetAccessToken(v string) *AddWebhookRequest {
	s.AccessToken = &v
	return s
}

func (s *AddWebhookRequest) SetOrganizationId(v string) *AddWebhookRequest {
	s.OrganizationId = &v
	return s
}

func (s *AddWebhookRequest) SetDescription(v string) *AddWebhookRequest {
	s.Description = &v
	return s
}

func (s *AddWebhookRequest) SetEnableSslVerification(v bool) *AddWebhookRequest {
	s.EnableSslVerification = &v
	return s
}

func (s *AddWebhookRequest) SetMergeRequestsEvents(v bool) *AddWebhookRequest {
	s.MergeRequestsEvents = &v
	return s
}

func (s *AddWebhookRequest) SetNoteEvents(v bool) *AddWebhookRequest {
	s.NoteEvents = &v
	return s
}

func (s *AddWebhookRequest) SetPushEvents(v bool) *AddWebhookRequest {
	s.PushEvents = &v
	return s
}

func (s *AddWebhookRequest) SetSecretToken(v string) *AddWebhookRequest {
	s.SecretToken = &v
	return s
}

func (s *AddWebhookRequest) SetTagPushEvents(v bool) *AddWebhookRequest {
	s.TagPushEvents = &v
	return s
}

func (s *AddWebhookRequest) SetUrl(v string) *AddWebhookRequest {
	s.Url = &v
	return s
}

type AddWebhookResponseBody struct {
	ErrorCode    *string                       `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                       `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       *AddWebhookResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success      *bool                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *AddWebhookResponseBody) SetErrorCode(v string) *AddWebhookResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddWebhookResponseBody) SetErrorMessage(v string) *AddWebhookResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddWebhookResponseBody) SetRequestId(v string) *AddWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWebhookResponseBody) SetResult(v *AddWebhookResponseBodyResult) *AddWebhookResponseBody {
	s.Result = v
	return s
}

func (s *AddWebhookResponseBody) SetSuccess(v bool) *AddWebhookResponseBody {
	s.Success = &v
	return s
}

type AddWebhookResponseBodyResult struct {
	CreatedAt             *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description           *string `json:"description,omitempty" xml:"description,omitempty"`
	EnableSslVerification *bool   `json:"enableSslVerification,omitempty" xml:"enableSslVerification,omitempty"`
	Id                    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	LastTestResult        *string `json:"lastTestResult,omitempty" xml:"lastTestResult,omitempty"`
	MergeRequestsEvents   *bool   `json:"mergeRequestsEvents,omitempty" xml:"mergeRequestsEvents,omitempty"`
	NoteEvents            *bool   `json:"noteEvents,omitempty" xml:"noteEvents,omitempty"`
	PushEvents            *bool   `json:"pushEvents,omitempty" xml:"pushEvents,omitempty"`
	RepositoryId          *int64  `json:"repositoryId,omitempty" xml:"repositoryId,omitempty"`
	SecretToken           *string `json:"secretToken,omitempty" xml:"secretToken,omitempty"`
	TagPushEvents         *bool   `json:"tagPushEvents,omitempty" xml:"tagPushEvents,omitempty"`
	Url                   *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s AddWebhookResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddWebhookResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddWebhookResponseBodyResult) SetCreatedAt(v string) *AddWebhookResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetDescription(v string) *AddWebhookResponseBodyResult {
	s.Description = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetEnableSslVerification(v bool) *AddWebhookResponseBodyResult {
	s.EnableSslVerification = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetId(v int64) *AddWebhookResponseBodyResult {
	s.Id = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetLastTestResult(v string) *AddWebhookResponseBodyResult {
	s.LastTestResult = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetMergeRequestsEvents(v bool) *AddWebhookResponseBodyResult {
	s.MergeRequestsEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetNoteEvents(v bool) *AddWebhookResponseBodyResult {
	s.NoteEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetPushEvents(v bool) *AddWebhookResponseBodyResult {
	s.PushEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetRepositoryId(v int64) *AddWebhookResponseBodyResult {
	s.RepositoryId = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetSecretToken(v string) *AddWebhookResponseBodyResult {
	s.SecretToken = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetTagPushEvents(v bool) *AddWebhookResponseBodyResult {
	s.TagPushEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetUrl(v string) *AddWebhookResponseBodyResult {
	s.Url = &v
	return s
}

type AddWebhookResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s AddWebhookResponse) GoString() string {
	return s.String()
}

func (s *AddWebhookResponse) SetHeaders(v map[string]*string) *AddWebhookResponse {
	s.Headers = v
	return s
}

func (s *AddWebhookResponse) SetStatusCode(v int32) *AddWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWebhookResponse) SetBody(v *AddWebhookResponseBody) *AddWebhookResponse {
	s.Body = v
	return s
}

type CreateBranchRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	BranchName     *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	Ref            *string `json:"ref,omitempty" xml:"ref,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s CreateBranchRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBranchRequest) GoString() string {
	return s.String()
}

func (s *CreateBranchRequest) SetAccessToken(v string) *CreateBranchRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateBranchRequest) SetBranchName(v string) *CreateBranchRequest {
	s.BranchName = &v
	return s
}

func (s *CreateBranchRequest) SetRef(v string) *CreateBranchRequest {
	s.Ref = &v
	return s
}

func (s *CreateBranchRequest) SetOrganizationId(v string) *CreateBranchRequest {
	s.OrganizationId = &v
	return s
}

type CreateBranchResponseBody struct {
	ErrorCode    *string                         `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                         `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       *CreateBranchResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success      *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateBranchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBranchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBranchResponseBody) SetErrorCode(v string) *CreateBranchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateBranchResponseBody) SetErrorMessage(v string) *CreateBranchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateBranchResponseBody) SetRequestId(v string) *CreateBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBranchResponseBody) SetResult(v *CreateBranchResponseBodyResult) *CreateBranchResponseBody {
	s.Result = v
	return s
}

func (s *CreateBranchResponseBody) SetSuccess(v bool) *CreateBranchResponseBody {
	s.Success = &v
	return s
}

type CreateBranchResponseBodyResult struct {
	Commit    *CreateBranchResponseBodyResultCommit `json:"commit,omitempty" xml:"commit,omitempty" type:"Struct"`
	Name      *string                               `json:"name,omitempty" xml:"name,omitempty"`
	Protected *bool                                 `json:"protected,omitempty" xml:"protected,omitempty"`
}

func (s CreateBranchResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateBranchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateBranchResponseBodyResult) SetCommit(v *CreateBranchResponseBodyResultCommit) *CreateBranchResponseBodyResult {
	s.Commit = v
	return s
}

func (s *CreateBranchResponseBodyResult) SetName(v string) *CreateBranchResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateBranchResponseBodyResult) SetProtected(v bool) *CreateBranchResponseBodyResult {
	s.Protected = &v
	return s
}

type CreateBranchResponseBodyResultCommit struct {
	AuthorEmail    *string   `json:"authorEmail,omitempty" xml:"authorEmail,omitempty"`
	AuthorName     *string   `json:"authorName,omitempty" xml:"authorName,omitempty"`
	AuthoredDate   *string   `json:"authoredDate,omitempty" xml:"authoredDate,omitempty"`
	CommittedDate  *string   `json:"committedDate,omitempty" xml:"committedDate,omitempty"`
	CommitterEmail *string   `json:"committerEmail,omitempty" xml:"committerEmail,omitempty"`
	CommitterName  *string   `json:"committerName,omitempty" xml:"committerName,omitempty"`
	CreatedAt      *string   `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Id             *string   `json:"id,omitempty" xml:"id,omitempty"`
	Message        *string   `json:"message,omitempty" xml:"message,omitempty"`
	ParentIds      []*string `json:"parentIds,omitempty" xml:"parentIds,omitempty" type:"Repeated"`
	ShortId        *string   `json:"shortId,omitempty" xml:"shortId,omitempty"`
	Title          *string   `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateBranchResponseBodyResultCommit) String() string {
	return tea.Prettify(s)
}

func (s CreateBranchResponseBodyResultCommit) GoString() string {
	return s.String()
}

func (s *CreateBranchResponseBodyResultCommit) SetAuthorEmail(v string) *CreateBranchResponseBodyResultCommit {
	s.AuthorEmail = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetAuthorName(v string) *CreateBranchResponseBodyResultCommit {
	s.AuthorName = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetAuthoredDate(v string) *CreateBranchResponseBodyResultCommit {
	s.AuthoredDate = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetCommittedDate(v string) *CreateBranchResponseBodyResultCommit {
	s.CommittedDate = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetCommitterEmail(v string) *CreateBranchResponseBodyResultCommit {
	s.CommitterEmail = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetCommitterName(v string) *CreateBranchResponseBodyResultCommit {
	s.CommitterName = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetCreatedAt(v string) *CreateBranchResponseBodyResultCommit {
	s.CreatedAt = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetId(v string) *CreateBranchResponseBodyResultCommit {
	s.Id = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetMessage(v string) *CreateBranchResponseBodyResultCommit {
	s.Message = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetParentIds(v []*string) *CreateBranchResponseBodyResultCommit {
	s.ParentIds = v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetShortId(v string) *CreateBranchResponseBodyResultCommit {
	s.ShortId = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetTitle(v string) *CreateBranchResponseBodyResultCommit {
	s.Title = &v
	return s
}

type CreateBranchResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBranchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBranchResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBranchResponse) GoString() string {
	return s.String()
}

func (s *CreateBranchResponse) SetHeaders(v map[string]*string) *CreateBranchResponse {
	s.Headers = v
	return s
}

func (s *CreateBranchResponse) SetStatusCode(v int32) *CreateBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBranchResponse) SetBody(v *CreateBranchResponseBody) *CreateBranchResponse {
	s.Body = v
	return s
}

type CreateFileRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	BranchName     *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	CommitMessage  *string `json:"commitMessage,omitempty" xml:"commitMessage,omitempty"`
	Content        *string `json:"content,omitempty" xml:"content,omitempty"`
	Encoding       *string `json:"encoding,omitempty" xml:"encoding,omitempty"`
	FilePath       *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s CreateFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFileRequest) GoString() string {
	return s.String()
}

func (s *CreateFileRequest) SetAccessToken(v string) *CreateFileRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateFileRequest) SetBranchName(v string) *CreateFileRequest {
	s.BranchName = &v
	return s
}

func (s *CreateFileRequest) SetCommitMessage(v string) *CreateFileRequest {
	s.CommitMessage = &v
	return s
}

func (s *CreateFileRequest) SetContent(v string) *CreateFileRequest {
	s.Content = &v
	return s
}

func (s *CreateFileRequest) SetEncoding(v string) *CreateFileRequest {
	s.Encoding = &v
	return s
}

func (s *CreateFileRequest) SetFilePath(v string) *CreateFileRequest {
	s.FilePath = &v
	return s
}

func (s *CreateFileRequest) SetOrganizationId(v string) *CreateFileRequest {
	s.OrganizationId = &v
	return s
}

type CreateFileResponseBody struct {
	ErrorCode    *string                       `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                       `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       *CreateFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success      *bool                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileResponseBody) SetErrorCode(v string) *CreateFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFileResponseBody) SetErrorMessage(v string) *CreateFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateFileResponseBody) SetRequestId(v string) *CreateFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileResponseBody) SetResult(v *CreateFileResponseBodyResult) *CreateFileResponseBody {
	s.Result = v
	return s
}

func (s *CreateFileResponseBody) SetSuccess(v bool) *CreateFileResponseBody {
	s.Success = &v
	return s
}

type CreateFileResponseBodyResult struct {
	BranchName *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	FilePath   *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
}

func (s CreateFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateFileResponseBodyResult) SetBranchName(v string) *CreateFileResponseBodyResult {
	s.BranchName = &v
	return s
}

func (s *CreateFileResponseBodyResult) SetFilePath(v string) *CreateFileResponseBodyResult {
	s.FilePath = &v
	return s
}

type CreateFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFileResponse) GoString() string {
	return s.String()
}

func (s *CreateFileResponse) SetHeaders(v map[string]*string) *CreateFileResponse {
	s.Headers = v
	return s
}

func (s *CreateFileResponse) SetStatusCode(v int32) *CreateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileResponse) SetBody(v *CreateFileResponseBody) *CreateFileResponse {
	s.Body = v
	return s
}

type CreateFlowTagRequest struct {
	Color          *string `json:"color,omitempty" xml:"color,omitempty"`
	FlowTagGroupId *int64  `json:"flowTagGroupId,omitempty" xml:"flowTagGroupId,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateFlowTagRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowTagRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowTagRequest) SetColor(v string) *CreateFlowTagRequest {
	s.Color = &v
	return s
}

func (s *CreateFlowTagRequest) SetFlowTagGroupId(v int64) *CreateFlowTagRequest {
	s.FlowTagGroupId = &v
	return s
}

func (s *CreateFlowTagRequest) SetName(v string) *CreateFlowTagRequest {
	s.Name = &v
	return s
}

type CreateFlowTagResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Id           *int64  `json:"id,omitempty" xml:"id,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateFlowTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowTagResponseBody) SetErrorCode(v string) *CreateFlowTagResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFlowTagResponseBody) SetErrorMessage(v string) *CreateFlowTagResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateFlowTagResponseBody) SetId(v int64) *CreateFlowTagResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFlowTagResponseBody) SetRequestId(v string) *CreateFlowTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowTagResponseBody) SetSuccess(v bool) *CreateFlowTagResponseBody {
	s.Success = &v
	return s
}

type CreateFlowTagResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFlowTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowTagResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowTagResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowTagResponse) SetHeaders(v map[string]*string) *CreateFlowTagResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowTagResponse) SetStatusCode(v int32) *CreateFlowTagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFlowTagResponse) SetBody(v *CreateFlowTagResponseBody) *CreateFlowTagResponse {
	s.Body = v
	return s
}

type CreateFlowTagGroupRequest struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateFlowTagGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowTagGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowTagGroupRequest) SetName(v string) *CreateFlowTagGroupRequest {
	s.Name = &v
	return s
}

type CreateFlowTagGroupResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Id           *int64  `json:"id,omitempty" xml:"id,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateFlowTagGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowTagGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowTagGroupResponseBody) SetErrorCode(v string) *CreateFlowTagGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFlowTagGroupResponseBody) SetErrorMessage(v string) *CreateFlowTagGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateFlowTagGroupResponseBody) SetId(v int64) *CreateFlowTagGroupResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFlowTagGroupResponseBody) SetRequestId(v string) *CreateFlowTagGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowTagGroupResponseBody) SetSuccess(v bool) *CreateFlowTagGroupResponseBody {
	s.Success = &v
	return s
}

type CreateFlowTagGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFlowTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowTagGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowTagGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowTagGroupResponse) SetHeaders(v map[string]*string) *CreateFlowTagGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowTagGroupResponse) SetStatusCode(v int32) *CreateFlowTagGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFlowTagGroupResponse) SetBody(v *CreateFlowTagGroupResponseBody) *CreateFlowTagGroupResponse {
	s.Body = v
	return s
}

type CreateHostGroupRequest struct {
	AliyunRegion        *string `json:"aliyunRegion,omitempty" xml:"aliyunRegion,omitempty"`
	EcsLabelKey         *string `json:"ecsLabelKey,omitempty" xml:"ecsLabelKey,omitempty"`
	EcsLabelValue       *string `json:"ecsLabelValue,omitempty" xml:"ecsLabelValue,omitempty"`
	EcsType             *string `json:"ecsType,omitempty" xml:"ecsType,omitempty"`
	EnvId               *string `json:"envId,omitempty" xml:"envId,omitempty"`
	MachineInfos        *string `json:"machineInfos,omitempty" xml:"machineInfos,omitempty"`
	Name                *string `json:"name,omitempty" xml:"name,omitempty"`
	ServiceConnectionId *int64  `json:"serviceConnectionId,omitempty" xml:"serviceConnectionId,omitempty"`
	TagIds              *string `json:"tagIds,omitempty" xml:"tagIds,omitempty"`
	Type                *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateHostGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHostGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateHostGroupRequest) SetAliyunRegion(v string) *CreateHostGroupRequest {
	s.AliyunRegion = &v
	return s
}

func (s *CreateHostGroupRequest) SetEcsLabelKey(v string) *CreateHostGroupRequest {
	s.EcsLabelKey = &v
	return s
}

func (s *CreateHostGroupRequest) SetEcsLabelValue(v string) *CreateHostGroupRequest {
	s.EcsLabelValue = &v
	return s
}

func (s *CreateHostGroupRequest) SetEcsType(v string) *CreateHostGroupRequest {
	s.EcsType = &v
	return s
}

func (s *CreateHostGroupRequest) SetEnvId(v string) *CreateHostGroupRequest {
	s.EnvId = &v
	return s
}

func (s *CreateHostGroupRequest) SetMachineInfos(v string) *CreateHostGroupRequest {
	s.MachineInfos = &v
	return s
}

func (s *CreateHostGroupRequest) SetName(v string) *CreateHostGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateHostGroupRequest) SetServiceConnectionId(v int64) *CreateHostGroupRequest {
	s.ServiceConnectionId = &v
	return s
}

func (s *CreateHostGroupRequest) SetTagIds(v string) *CreateHostGroupRequest {
	s.TagIds = &v
	return s
}

func (s *CreateHostGroupRequest) SetType(v string) *CreateHostGroupRequest {
	s.Type = &v
	return s
}

type CreateHostGroupResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HostGroupId  *int64  `json:"hostGroupId,omitempty" xml:"hostGroupId,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateHostGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHostGroupResponseBody) SetErrorCode(v string) *CreateHostGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateHostGroupResponseBody) SetErrorMessage(v string) *CreateHostGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateHostGroupResponseBody) SetHostGroupId(v int64) *CreateHostGroupResponseBody {
	s.HostGroupId = &v
	return s
}

func (s *CreateHostGroupResponseBody) SetRequestId(v string) *CreateHostGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHostGroupResponseBody) SetSuccess(v bool) *CreateHostGroupResponseBody {
	s.Success = &v
	return s
}

type CreateHostGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHostGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHostGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateHostGroupResponse) SetHeaders(v map[string]*string) *CreateHostGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateHostGroupResponse) SetStatusCode(v int32) *CreateHostGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHostGroupResponse) SetBody(v *CreateHostGroupResponseBody) *CreateHostGroupResponse {
	s.Body = v
	return s
}

type CreateOAuthTokenRequest struct {
	ClientId     *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	ClientSecret *string `json:"clientSecret,omitempty" xml:"clientSecret,omitempty"`
	Code         *string `json:"code,omitempty" xml:"code,omitempty"`
	GrantType    *string `json:"grantType,omitempty" xml:"grantType,omitempty"`
	Login        *string `json:"login,omitempty" xml:"login,omitempty"`
	Scope        *string `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s CreateOAuthTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateOAuthTokenRequest) SetClientId(v string) *CreateOAuthTokenRequest {
	s.ClientId = &v
	return s
}

func (s *CreateOAuthTokenRequest) SetClientSecret(v string) *CreateOAuthTokenRequest {
	s.ClientSecret = &v
	return s
}

func (s *CreateOAuthTokenRequest) SetCode(v string) *CreateOAuthTokenRequest {
	s.Code = &v
	return s
}

func (s *CreateOAuthTokenRequest) SetGrantType(v string) *CreateOAuthTokenRequest {
	s.GrantType = &v
	return s
}

func (s *CreateOAuthTokenRequest) SetLogin(v string) *CreateOAuthTokenRequest {
	s.Login = &v
	return s
}

func (s *CreateOAuthTokenRequest) SetScope(v string) *CreateOAuthTokenRequest {
	s.Scope = &v
	return s
}

type CreateOAuthTokenResponseBody struct {
	ErrorCode    *string                             `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                             `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       *CreateOAuthTokenResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success      *string                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateOAuthTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOAuthTokenResponseBody) SetErrorCode(v string) *CreateOAuthTokenResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateOAuthTokenResponseBody) SetErrorMessage(v string) *CreateOAuthTokenResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateOAuthTokenResponseBody) SetRequestId(v string) *CreateOAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOAuthTokenResponseBody) SetResult(v *CreateOAuthTokenResponseBodyResult) *CreateOAuthTokenResponseBody {
	s.Result = v
	return s
}

func (s *CreateOAuthTokenResponseBody) SetSuccess(v string) *CreateOAuthTokenResponseBody {
	s.Success = &v
	return s
}

type CreateOAuthTokenResponseBodyResult struct {
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	Scope       *string `json:"scope,omitempty" xml:"scope,omitempty"`
	TokenType   *string `json:"tokenType,omitempty" xml:"tokenType,omitempty"`
}

func (s CreateOAuthTokenResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateOAuthTokenResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateOAuthTokenResponseBodyResult) SetAccessToken(v string) *CreateOAuthTokenResponseBodyResult {
	s.AccessToken = &v
	return s
}

func (s *CreateOAuthTokenResponseBodyResult) SetId(v string) *CreateOAuthTokenResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateOAuthTokenResponseBodyResult) SetScope(v string) *CreateOAuthTokenResponseBodyResult {
	s.Scope = &v
	return s
}

func (s *CreateOAuthTokenResponseBodyResult) SetTokenType(v string) *CreateOAuthTokenResponseBodyResult {
	s.TokenType = &v
	return s
}

type CreateOAuthTokenResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOAuthTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateOAuthTokenResponse) SetHeaders(v map[string]*string) *CreateOAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateOAuthTokenResponse) SetStatusCode(v int32) *CreateOAuthTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOAuthTokenResponse) SetBody(v *CreateOAuthTokenResponseBody) *CreateOAuthTokenResponse {
	s.Body = v
	return s
}

type CreatePipelineGroupRequest struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreatePipelineGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineGroupRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineGroupRequest) SetName(v string) *CreatePipelineGroupRequest {
	s.Name = &v
	return s
}

type CreatePipelineGroupResponseBody struct {
	ErrorCode     *string                                       `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage  *string                                       `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	PipelineGroup *CreatePipelineGroupResponseBodyPipelineGroup `json:"pipelineGroup,omitempty" xml:"pipelineGroup,omitempty" type:"Struct"`
	RequestId     *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success       *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreatePipelineGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePipelineGroupResponseBody) SetErrorCode(v string) *CreatePipelineGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreatePipelineGroupResponseBody) SetErrorMessage(v string) *CreatePipelineGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreatePipelineGroupResponseBody) SetPipelineGroup(v *CreatePipelineGroupResponseBodyPipelineGroup) *CreatePipelineGroupResponseBody {
	s.PipelineGroup = v
	return s
}

func (s *CreatePipelineGroupResponseBody) SetRequestId(v string) *CreatePipelineGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePipelineGroupResponseBody) SetSuccess(v bool) *CreatePipelineGroupResponseBody {
	s.Success = &v
	return s
}

type CreatePipelineGroupResponseBodyPipelineGroup struct {
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreatePipelineGroupResponseBodyPipelineGroup) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineGroupResponseBodyPipelineGroup) GoString() string {
	return s.String()
}

func (s *CreatePipelineGroupResponseBodyPipelineGroup) SetId(v int64) *CreatePipelineGroupResponseBodyPipelineGroup {
	s.Id = &v
	return s
}

func (s *CreatePipelineGroupResponseBodyPipelineGroup) SetName(v string) *CreatePipelineGroupResponseBodyPipelineGroup {
	s.Name = &v
	return s
}

type CreatePipelineGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePipelineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePipelineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineGroupResponse) GoString() string {
	return s.String()
}

func (s *CreatePipelineGroupResponse) SetHeaders(v map[string]*string) *CreatePipelineGroupResponse {
	s.Headers = v
	return s
}

func (s *CreatePipelineGroupResponse) SetStatusCode(v int32) *CreatePipelineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePipelineGroupResponse) SetBody(v *CreatePipelineGroupResponseBody) *CreatePipelineGroupResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	CustomCode         *string `json:"customCode,omitempty" xml:"customCode,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	Scope              *string `json:"scope,omitempty" xml:"scope,omitempty"`
	TemplateIdentifier *string `json:"templateIdentifier,omitempty" xml:"templateIdentifier,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetCustomCode(v string) *CreateProjectRequest {
	s.CustomCode = &v
	return s
}

func (s *CreateProjectRequest) SetName(v string) *CreateProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectRequest) SetScope(v string) *CreateProjectRequest {
	s.Scope = &v
	return s
}

func (s *CreateProjectRequest) SetTemplateIdentifier(v string) *CreateProjectRequest {
	s.TemplateIdentifier = &v
	return s
}

type CreateProjectResponseBody struct {
	ErrorCode *string                           `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                           `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Project   *CreateProjectResponseBodyProject `json:"project,omitempty" xml:"project,omitempty" type:"Struct"`
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetErrorCode(v string) *CreateProjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateProjectResponseBody) SetErrorMsg(v string) *CreateProjectResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateProjectResponseBody) SetProject(v *CreateProjectResponseBodyProject) *CreateProjectResponseBody {
	s.Project = v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) SetSuccess(v bool) *CreateProjectResponseBody {
	s.Success = &v
	return s
}

type CreateProjectResponseBodyProject struct {
	CategoryIdentifier     *string `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	Creator                *string `json:"creator,omitempty" xml:"creator,omitempty"`
	CustomCode             *string `json:"customCode,omitempty" xml:"customCode,omitempty"`
	Description            *string `json:"description,omitempty" xml:"description,omitempty"`
	GmtCreate              *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified            *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Icon                   *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Identifier             *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	LogicalStatus          *string `json:"logicalStatus,omitempty" xml:"logicalStatus,omitempty"`
	Modifier               *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Name                   *string `json:"name,omitempty" xml:"name,omitempty"`
	OrganizationIdentifier *string `json:"organizationIdentifier,omitempty" xml:"organizationIdentifier,omitempty"`
	Scope                  *string `json:"scope,omitempty" xml:"scope,omitempty"`
	StatusIdentifier       *string `json:"statusIdentifier,omitempty" xml:"statusIdentifier,omitempty"`
	StatusStageIdentifier  *string `json:"statusStageIdentifier,omitempty" xml:"statusStageIdentifier,omitempty"`
	TypeIdentifier         *string `json:"typeIdentifier,omitempty" xml:"typeIdentifier,omitempty"`
}

func (s CreateProjectResponseBodyProject) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyProject) SetCategoryIdentifier(v string) *CreateProjectResponseBodyProject {
	s.CategoryIdentifier = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetCreator(v string) *CreateProjectResponseBodyProject {
	s.Creator = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetCustomCode(v string) *CreateProjectResponseBodyProject {
	s.CustomCode = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetDescription(v string) *CreateProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetGmtCreate(v int64) *CreateProjectResponseBodyProject {
	s.GmtCreate = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetGmtModified(v int64) *CreateProjectResponseBodyProject {
	s.GmtModified = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetIcon(v string) *CreateProjectResponseBodyProject {
	s.Icon = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetIdentifier(v string) *CreateProjectResponseBodyProject {
	s.Identifier = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetLogicalStatus(v string) *CreateProjectResponseBodyProject {
	s.LogicalStatus = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetModifier(v string) *CreateProjectResponseBodyProject {
	s.Modifier = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetName(v string) *CreateProjectResponseBodyProject {
	s.Name = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetOrganizationIdentifier(v string) *CreateProjectResponseBodyProject {
	s.OrganizationIdentifier = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetScope(v string) *CreateProjectResponseBodyProject {
	s.Scope = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetStatusIdentifier(v string) *CreateProjectResponseBodyProject {
	s.StatusIdentifier = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetStatusStageIdentifier(v string) *CreateProjectResponseBodyProject {
	s.StatusStageIdentifier = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetTypeIdentifier(v string) *CreateProjectResponseBodyProject {
	s.TypeIdentifier = &v
	return s
}

type CreateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetStatusCode(v int32) *CreateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type CreateRepositoryRequest struct {
	AccessToken          *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	AvatarUrl            *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Description          *string `json:"description,omitempty" xml:"description,omitempty"`
	GitignoreType        *string `json:"gitignoreType,omitempty" xml:"gitignoreType,omitempty"`
	ImportAccount        *string `json:"importAccount,omitempty" xml:"importAccount,omitempty"`
	ImportDemoProject    *bool   `json:"importDemoProject,omitempty" xml:"importDemoProject,omitempty"`
	ImportRepoType       *string `json:"importRepoType,omitempty" xml:"importRepoType,omitempty"`
	ImportToken          *string `json:"importToken,omitempty" xml:"importToken,omitempty"`
	ImportTokenEncrypted *string `json:"importTokenEncrypted,omitempty" xml:"importTokenEncrypted,omitempty"`
	ImportUrl            *string `json:"importUrl,omitempty" xml:"importUrl,omitempty"`
	InitStandardService  *bool   `json:"initStandardService,omitempty" xml:"initStandardService,omitempty"`
	IsCryptoEnabled      *bool   `json:"isCryptoEnabled,omitempty" xml:"isCryptoEnabled,omitempty"`
	LocalImportUrl       *string `json:"localImportUrl,omitempty" xml:"localImportUrl,omitempty"`
	Name                 *string `json:"name,omitempty" xml:"name,omitempty"`
	NamespaceId          *int64  `json:"namespaceId,omitempty" xml:"namespaceId,omitempty"`
	Path                 *string `json:"path,omitempty" xml:"path,omitempty"`
	ReadmeType           *string `json:"readmeType,omitempty" xml:"readmeType,omitempty"`
	VisibilityLevel      *int32  `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	CreateParentPath     *bool   `json:"createParentPath,omitempty" xml:"createParentPath,omitempty"`
	OrganizationId       *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	Sync                 *bool   `json:"sync,omitempty" xml:"sync,omitempty"`
}

func (s CreateRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryRequest) GoString() string {
	return s.String()
}

func (s *CreateRepositoryRequest) SetAccessToken(v string) *CreateRepositoryRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateRepositoryRequest) SetAvatarUrl(v string) *CreateRepositoryRequest {
	s.AvatarUrl = &v
	return s
}

func (s *CreateRepositoryRequest) SetDescription(v string) *CreateRepositoryRequest {
	s.Description = &v
	return s
}

func (s *CreateRepositoryRequest) SetGitignoreType(v string) *CreateRepositoryRequest {
	s.GitignoreType = &v
	return s
}

func (s *CreateRepositoryRequest) SetImportAccount(v string) *CreateRepositoryRequest {
	s.ImportAccount = &v
	return s
}

func (s *CreateRepositoryRequest) SetImportDemoProject(v bool) *CreateRepositoryRequest {
	s.ImportDemoProject = &v
	return s
}

func (s *CreateRepositoryRequest) SetImportRepoType(v string) *CreateRepositoryRequest {
	s.ImportRepoType = &v
	return s
}

func (s *CreateRepositoryRequest) SetImportToken(v string) *CreateRepositoryRequest {
	s.ImportToken = &v
	return s
}

func (s *CreateRepositoryRequest) SetImportTokenEncrypted(v string) *CreateRepositoryRequest {
	s.ImportTokenEncrypted = &v
	return s
}

func (s *CreateRepositoryRequest) SetImportUrl(v string) *CreateRepositoryRequest {
	s.ImportUrl = &v
	return s
}

func (s *CreateRepositoryRequest) SetInitStandardService(v bool) *CreateRepositoryRequest {
	s.InitStandardService = &v
	return s
}

func (s *CreateRepositoryRequest) SetIsCryptoEnabled(v bool) *CreateRepositoryRequest {
	s.IsCryptoEnabled = &v
	return s
}

func (s *CreateRepositoryRequest) SetLocalImportUrl(v string) *CreateRepositoryRequest {
	s.LocalImportUrl = &v
	return s
}

func (s *CreateRepositoryRequest) SetName(v string) *CreateRepositoryRequest {
	s.Name = &v
	return s
}

func (s *CreateRepositoryRequest) SetNamespaceId(v int64) *CreateRepositoryRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateRepositoryRequest) SetPath(v string) *CreateRepositoryRequest {
	s.Path = &v
	return s
}

func (s *CreateRepositoryRequest) SetReadmeType(v string) *CreateRepositoryRequest {
	s.ReadmeType = &v
	return s
}

func (s *CreateRepositoryRequest) SetVisibilityLevel(v int32) *CreateRepositoryRequest {
	s.VisibilityLevel = &v
	return s
}

func (s *CreateRepositoryRequest) SetCreateParentPath(v bool) *CreateRepositoryRequest {
	s.CreateParentPath = &v
	return s
}

func (s *CreateRepositoryRequest) SetOrganizationId(v string) *CreateRepositoryRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateRepositoryRequest) SetSync(v bool) *CreateRepositoryRequest {
	s.Sync = &v
	return s
}

type CreateRepositoryResponseBody struct {
	ErrorCode    *string                             `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                             `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       *CreateRepositoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success      *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponseBody) SetErrorCode(v string) *CreateRepositoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetErrorMessage(v string) *CreateRepositoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetRequestId(v string) *CreateRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetResult(v *CreateRepositoryResponseBodyResult) *CreateRepositoryResponseBody {
	s.Result = v
	return s
}

func (s *CreateRepositoryResponseBody) SetSuccess(v bool) *CreateRepositoryResponseBody {
	s.Success = &v
	return s
}

type CreateRepositoryResponseBodyResult struct {
	ImportFromSvn     *bool                                        `json:"Import_from_svn,omitempty" xml:"Import_from_svn,omitempty"`
	Archived          *bool                                        `json:"archived,omitempty" xml:"archived,omitempty"`
	AvatarUrl         *string                                      `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
	CreatedAt         *string                                      `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatorId         *int64                                       `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	DefaultBranch     *string                                      `json:"defaultBranch,omitempty" xml:"defaultBranch,omitempty"`
	DemoProject       *bool                                        `json:"demoProject,omitempty" xml:"demoProject,omitempty"`
	Description       *string                                      `json:"description,omitempty" xml:"description,omitempty"`
	HttpUrlToRepo     *string                                      `json:"httpUrlToRepo,omitempty" xml:"httpUrlToRepo,omitempty"`
	Id                *int64                                       `json:"id,omitempty" xml:"id,omitempty"`
	LastActivityAt    *string                                      `json:"lastActivityAt,omitempty" xml:"lastActivityAt,omitempty"`
	Name              *string                                      `json:"name,omitempty" xml:"name,omitempty"`
	NameWithNamespace *string                                      `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	Namespace         *CreateRepositoryResponseBodyResultNamespace `json:"namespace,omitempty" xml:"namespace,omitempty" type:"Struct"`
	Path              *string                                      `json:"path,omitempty" xml:"path,omitempty"`
	PathWithNamespace *string                                      `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	SshUrlToRepo      *string                                      `json:"sshUrlToRepo,omitempty" xml:"sshUrlToRepo,omitempty"`
	VisibilityLevel   *string                                      `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	WebUrl            *string                                      `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s CreateRepositoryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponseBodyResult) SetImportFromSvn(v bool) *CreateRepositoryResponseBodyResult {
	s.ImportFromSvn = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetArchived(v bool) *CreateRepositoryResponseBodyResult {
	s.Archived = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetAvatarUrl(v string) *CreateRepositoryResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetCreatedAt(v string) *CreateRepositoryResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetCreatorId(v int64) *CreateRepositoryResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetDefaultBranch(v string) *CreateRepositoryResponseBodyResult {
	s.DefaultBranch = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetDemoProject(v bool) *CreateRepositoryResponseBodyResult {
	s.DemoProject = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetDescription(v string) *CreateRepositoryResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetHttpUrlToRepo(v string) *CreateRepositoryResponseBodyResult {
	s.HttpUrlToRepo = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetId(v int64) *CreateRepositoryResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetLastActivityAt(v string) *CreateRepositoryResponseBodyResult {
	s.LastActivityAt = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetName(v string) *CreateRepositoryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetNameWithNamespace(v string) *CreateRepositoryResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetNamespace(v *CreateRepositoryResponseBodyResultNamespace) *CreateRepositoryResponseBodyResult {
	s.Namespace = v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetPath(v string) *CreateRepositoryResponseBodyResult {
	s.Path = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetPathWithNamespace(v string) *CreateRepositoryResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetSshUrlToRepo(v string) *CreateRepositoryResponseBodyResult {
	s.SshUrlToRepo = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetVisibilityLevel(v string) *CreateRepositoryResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetWebUrl(v string) *CreateRepositoryResponseBodyResult {
	s.WebUrl = &v
	return s
}

type CreateRepositoryResponseBodyResultNamespace struct {
	Avatar          *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	CreatedAt       *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	Id              *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	OwnerId         *int64  `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	Path            *string `json:"path,omitempty" xml:"path,omitempty"`
	Public          *bool   `json:"public,omitempty" xml:"public,omitempty"`
	UpdatedAt       *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	VisibilityLevel *string `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
}

func (s CreateRepositoryResponseBodyResultNamespace) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryResponseBodyResultNamespace) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetAvatar(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.Avatar = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetCreatedAt(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.CreatedAt = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetDescription(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.Description = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetId(v int64) *CreateRepositoryResponseBodyResultNamespace {
	s.Id = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetName(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.Name = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetOwnerId(v int64) *CreateRepositoryResponseBodyResultNamespace {
	s.OwnerId = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetPath(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.Path = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetPublic(v bool) *CreateRepositoryResponseBodyResultNamespace {
	s.Public = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetUpdatedAt(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.UpdatedAt = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetVisibilityLevel(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.VisibilityLevel = &v
	return s
}

type CreateRepositoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryResponse) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponse) SetHeaders(v map[string]*string) *CreateRepositoryResponse {
	s.Headers = v
	return s
}

func (s *CreateRepositoryResponse) SetStatusCode(v int32) *CreateRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepositoryResponse) SetBody(v *CreateRepositoryResponseBody) *CreateRepositoryResponse {
	s.Body = v
	return s
}

type CreateResourceMemberRequest struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	RoleName  *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s CreateResourceMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceMemberRequest) SetAccountId(v string) *CreateResourceMemberRequest {
	s.AccountId = &v
	return s
}

func (s *CreateResourceMemberRequest) SetRoleName(v string) *CreateResourceMemberRequest {
	s.RoleName = &v
	return s
}

type CreateResourceMemberResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateResourceMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceMemberResponseBody) SetErrorCode(v string) *CreateResourceMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateResourceMemberResponseBody) SetErrorMessage(v string) *CreateResourceMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateResourceMemberResponseBody) SetRequestId(v string) *CreateResourceMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceMemberResponseBody) SetSuccess(v bool) *CreateResourceMemberResponseBody {
	s.Success = &v
	return s
}

type CreateResourceMemberResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateResourceMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateResourceMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceMemberResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceMemberResponse) SetHeaders(v map[string]*string) *CreateResourceMemberResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceMemberResponse) SetStatusCode(v int32) *CreateResourceMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceMemberResponse) SetBody(v *CreateResourceMemberResponseBody) *CreateResourceMemberResponse {
	s.Body = v
	return s
}

type CreateSprintRequest struct {
	EndDate         *string   `json:"endDate,omitempty" xml:"endDate,omitempty"`
	Name            *string   `json:"name,omitempty" xml:"name,omitempty"`
	SpaceIdentifier *string   `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	StaffIds        []*string `json:"staffIds,omitempty" xml:"staffIds,omitempty" type:"Repeated"`
	StartDate       *string   `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s CreateSprintRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSprintRequest) GoString() string {
	return s.String()
}

func (s *CreateSprintRequest) SetEndDate(v string) *CreateSprintRequest {
	s.EndDate = &v
	return s
}

func (s *CreateSprintRequest) SetName(v string) *CreateSprintRequest {
	s.Name = &v
	return s
}

func (s *CreateSprintRequest) SetSpaceIdentifier(v string) *CreateSprintRequest {
	s.SpaceIdentifier = &v
	return s
}

func (s *CreateSprintRequest) SetStaffIds(v []*string) *CreateSprintRequest {
	s.StaffIds = v
	return s
}

func (s *CreateSprintRequest) SetStartDate(v string) *CreateSprintRequest {
	s.StartDate = &v
	return s
}

type CreateSprintResponseBody struct {
	ErrorCode *string                         `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                         `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId *string                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Sprint    *CreateSprintResponseBodySprint `json:"sprint,omitempty" xml:"sprint,omitempty" type:"Struct"`
	Success   *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateSprintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSprintResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSprintResponseBody) SetErrorCode(v string) *CreateSprintResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateSprintResponseBody) SetErrorMsg(v string) *CreateSprintResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateSprintResponseBody) SetRequestId(v string) *CreateSprintResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSprintResponseBody) SetSprint(v *CreateSprintResponseBodySprint) *CreateSprintResponseBody {
	s.Sprint = v
	return s
}

func (s *CreateSprintResponseBody) SetSuccess(v bool) *CreateSprintResponseBody {
	s.Success = &v
	return s
}

type CreateSprintResponseBodySprint struct {
	Creator         *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	EndDate         *int64  `json:"endDate,omitempty" xml:"endDate,omitempty"`
	GmtCreate       *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified     *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Identifier      *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Modifier        *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	Scope           *string `json:"scope,omitempty" xml:"scope,omitempty"`
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	StartDate       *int64  `json:"startDate,omitempty" xml:"startDate,omitempty"`
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateSprintResponseBodySprint) String() string {
	return tea.Prettify(s)
}

func (s CreateSprintResponseBodySprint) GoString() string {
	return s.String()
}

func (s *CreateSprintResponseBodySprint) SetCreator(v string) *CreateSprintResponseBodySprint {
	s.Creator = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetDescription(v string) *CreateSprintResponseBodySprint {
	s.Description = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetEndDate(v int64) *CreateSprintResponseBodySprint {
	s.EndDate = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetGmtCreate(v int64) *CreateSprintResponseBodySprint {
	s.GmtCreate = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetGmtModified(v int64) *CreateSprintResponseBodySprint {
	s.GmtModified = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetIdentifier(v string) *CreateSprintResponseBodySprint {
	s.Identifier = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetModifier(v string) *CreateSprintResponseBodySprint {
	s.Modifier = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetName(v string) *CreateSprintResponseBodySprint {
	s.Name = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetScope(v string) *CreateSprintResponseBodySprint {
	s.Scope = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetSpaceIdentifier(v string) *CreateSprintResponseBodySprint {
	s.SpaceIdentifier = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetStartDate(v int64) *CreateSprintResponseBodySprint {
	s.StartDate = &v
	return s
}

func (s *CreateSprintResponseBodySprint) SetStatus(v string) *CreateSprintResponseBodySprint {
	s.Status = &v
	return s
}

type CreateSprintResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSprintResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSprintResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSprintResponse) GoString() string {
	return s.String()
}

func (s *CreateSprintResponse) SetHeaders(v map[string]*string) *CreateSprintResponse {
	s.Headers = v
	return s
}

func (s *CreateSprintResponse) SetStatusCode(v int32) *CreateSprintResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSprintResponse) SetBody(v *CreateSprintResponseBody) *CreateSprintResponse {
	s.Body = v
	return s
}

type CreateSshKeyResponseBody struct {
	ErrorCode    *string                         `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                         `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SshKey       *CreateSshKeyResponseBodySshKey `json:"sshKey,omitempty" xml:"sshKey,omitempty" type:"Struct"`
	Success      *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateSshKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSshKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSshKeyResponseBody) SetErrorCode(v string) *CreateSshKeyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateSshKeyResponseBody) SetErrorMessage(v string) *CreateSshKeyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSshKeyResponseBody) SetRequestId(v string) *CreateSshKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSshKeyResponseBody) SetSshKey(v *CreateSshKeyResponseBodySshKey) *CreateSshKeyResponseBody {
	s.SshKey = v
	return s
}

func (s *CreateSshKeyResponseBody) SetSuccess(v bool) *CreateSshKeyResponseBody {
	s.Success = &v
	return s
}

type CreateSshKeyResponseBodySshKey struct {
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	PublicKey *string `json:"publicKey,omitempty" xml:"publicKey,omitempty"`
}

func (s CreateSshKeyResponseBodySshKey) String() string {
	return tea.Prettify(s)
}

func (s CreateSshKeyResponseBodySshKey) GoString() string {
	return s.String()
}

func (s *CreateSshKeyResponseBodySshKey) SetId(v int64) *CreateSshKeyResponseBodySshKey {
	s.Id = &v
	return s
}

func (s *CreateSshKeyResponseBodySshKey) SetPublicKey(v string) *CreateSshKeyResponseBodySshKey {
	s.PublicKey = &v
	return s
}

type CreateSshKeyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSshKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSshKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSshKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateSshKeyResponse) SetHeaders(v map[string]*string) *CreateSshKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateSshKeyResponse) SetStatusCode(v int32) *CreateSshKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSshKeyResponse) SetBody(v *CreateSshKeyResponseBody) *CreateSshKeyResponse {
	s.Body = v
	return s
}

type CreateVariableGroupRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Variables   *string `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s CreateVariableGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVariableGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateVariableGroupRequest) SetDescription(v string) *CreateVariableGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateVariableGroupRequest) SetName(v string) *CreateVariableGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateVariableGroupRequest) SetVariables(v string) *CreateVariableGroupRequest {
	s.Variables = &v
	return s
}

type CreateVariableGroupResponseBody struct {
	ErrorCode       *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage    *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId       *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success         *bool   `json:"success,omitempty" xml:"success,omitempty"`
	VariableGroupId *int64  `json:"variableGroupId,omitempty" xml:"variableGroupId,omitempty"`
}

func (s CreateVariableGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVariableGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVariableGroupResponseBody) SetErrorCode(v string) *CreateVariableGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateVariableGroupResponseBody) SetErrorMessage(v string) *CreateVariableGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateVariableGroupResponseBody) SetRequestId(v string) *CreateVariableGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVariableGroupResponseBody) SetSuccess(v bool) *CreateVariableGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVariableGroupResponseBody) SetVariableGroupId(v int64) *CreateVariableGroupResponseBody {
	s.VariableGroupId = &v
	return s
}

type CreateVariableGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVariableGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVariableGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVariableGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateVariableGroupResponse) SetHeaders(v map[string]*string) *CreateVariableGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateVariableGroupResponse) SetStatusCode(v int32) *CreateVariableGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVariableGroupResponse) SetBody(v *CreateVariableGroupResponseBody) *CreateVariableGroupResponse {
	s.Body = v
	return s
}

type CreateWorkitemRequest struct {
	AssignedTo        *string                                `json:"assignedTo,omitempty" xml:"assignedTo,omitempty"`
	Category          *string                                `json:"category,omitempty" xml:"category,omitempty"`
	Description       *string                                `json:"description,omitempty" xml:"description,omitempty"`
	DescriptionFormat *string                                `json:"descriptionFormat,omitempty" xml:"descriptionFormat,omitempty"`
	FieldValueList    []*CreateWorkitemRequestFieldValueList `json:"fieldValueList,omitempty" xml:"fieldValueList,omitempty" type:"Repeated"`
	Parent            *string                                `json:"parent,omitempty" xml:"parent,omitempty"`
	Participant       []*string                              `json:"participant,omitempty" xml:"participant,omitempty" type:"Repeated"`
	Space             *string                                `json:"space,omitempty" xml:"space,omitempty"`
	SpaceIdentifier   *string                                `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	SpaceType         *string                                `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	Sprint            []*string                              `json:"sprint,omitempty" xml:"sprint,omitempty" type:"Repeated"`
	Subject           *string                                `json:"subject,omitempty" xml:"subject,omitempty"`
	Tracker           []*string                              `json:"tracker,omitempty" xml:"tracker,omitempty" type:"Repeated"`
	Verifier          []*string                              `json:"verifier,omitempty" xml:"verifier,omitempty" type:"Repeated"`
	WorkitemType      *string                                `json:"workitemType,omitempty" xml:"workitemType,omitempty"`
}

func (s CreateWorkitemRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkitemRequest) SetAssignedTo(v string) *CreateWorkitemRequest {
	s.AssignedTo = &v
	return s
}

func (s *CreateWorkitemRequest) SetCategory(v string) *CreateWorkitemRequest {
	s.Category = &v
	return s
}

func (s *CreateWorkitemRequest) SetDescription(v string) *CreateWorkitemRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkitemRequest) SetDescriptionFormat(v string) *CreateWorkitemRequest {
	s.DescriptionFormat = &v
	return s
}

func (s *CreateWorkitemRequest) SetFieldValueList(v []*CreateWorkitemRequestFieldValueList) *CreateWorkitemRequest {
	s.FieldValueList = v
	return s
}

func (s *CreateWorkitemRequest) SetParent(v string) *CreateWorkitemRequest {
	s.Parent = &v
	return s
}

func (s *CreateWorkitemRequest) SetParticipant(v []*string) *CreateWorkitemRequest {
	s.Participant = v
	return s
}

func (s *CreateWorkitemRequest) SetSpace(v string) *CreateWorkitemRequest {
	s.Space = &v
	return s
}

func (s *CreateWorkitemRequest) SetSpaceIdentifier(v string) *CreateWorkitemRequest {
	s.SpaceIdentifier = &v
	return s
}

func (s *CreateWorkitemRequest) SetSpaceType(v string) *CreateWorkitemRequest {
	s.SpaceType = &v
	return s
}

func (s *CreateWorkitemRequest) SetSprint(v []*string) *CreateWorkitemRequest {
	s.Sprint = v
	return s
}

func (s *CreateWorkitemRequest) SetSubject(v string) *CreateWorkitemRequest {
	s.Subject = &v
	return s
}

func (s *CreateWorkitemRequest) SetTracker(v []*string) *CreateWorkitemRequest {
	s.Tracker = v
	return s
}

func (s *CreateWorkitemRequest) SetVerifier(v []*string) *CreateWorkitemRequest {
	s.Verifier = v
	return s
}

func (s *CreateWorkitemRequest) SetWorkitemType(v string) *CreateWorkitemRequest {
	s.WorkitemType = &v
	return s
}

type CreateWorkitemRequestFieldValueList struct {
	FieldIdentifier    *string `json:"fieldIdentifier,omitempty" xml:"fieldIdentifier,omitempty"`
	Value              *string `json:"value,omitempty" xml:"value,omitempty"`
	WorkitemIdentifier *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s CreateWorkitemRequestFieldValueList) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemRequestFieldValueList) GoString() string {
	return s.String()
}

func (s *CreateWorkitemRequestFieldValueList) SetFieldIdentifier(v string) *CreateWorkitemRequestFieldValueList {
	s.FieldIdentifier = &v
	return s
}

func (s *CreateWorkitemRequestFieldValueList) SetValue(v string) *CreateWorkitemRequestFieldValueList {
	s.Value = &v
	return s
}

func (s *CreateWorkitemRequestFieldValueList) SetWorkitemIdentifier(v string) *CreateWorkitemRequestFieldValueList {
	s.WorkitemIdentifier = &v
	return s
}

type CreateWorkitemResponseBody struct {
	ErrorCode *string                             `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                             `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                               `json:"success,omitempty" xml:"success,omitempty"`
	Workitem  *CreateWorkitemResponseBodyWorkitem `json:"workitem,omitempty" xml:"workitem,omitempty" type:"Struct"`
}

func (s CreateWorkitemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkitemResponseBody) SetErrorCode(v string) *CreateWorkitemResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateWorkitemResponseBody) SetErrorMsg(v string) *CreateWorkitemResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateWorkitemResponseBody) SetRequestId(v string) *CreateWorkitemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkitemResponseBody) SetSuccess(v bool) *CreateWorkitemResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWorkitemResponseBody) SetWorkitem(v *CreateWorkitemResponseBodyWorkitem) *CreateWorkitemResponseBody {
	s.Workitem = v
	return s
}

type CreateWorkitemResponseBodyWorkitem struct {
	AssignedTo             *string `json:"assignedTo,omitempty" xml:"assignedTo,omitempty"`
	CategoryIdentifier     *string `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	Creator                *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Document               *string `json:"document,omitempty" xml:"document,omitempty"`
	GmtCreate              *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified            *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Identifier             *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	LogicalStatus          *string `json:"logicalStatus,omitempty" xml:"logicalStatus,omitempty"`
	Modifier               *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	ParentIdentifier       *string `json:"parentIdentifier,omitempty" xml:"parentIdentifier,omitempty"`
	SerialNumber           *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	SpaceIdentifier        *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	SpaceName              *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	SpaceType              *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	SprintIdentifier       *string `json:"sprintIdentifier,omitempty" xml:"sprintIdentifier,omitempty"`
	Status                 *string `json:"status,omitempty" xml:"status,omitempty"`
	StatusIdentifier       *string `json:"statusIdentifier,omitempty" xml:"statusIdentifier,omitempty"`
	StatusStageIdentifier  *string `json:"statusStageIdentifier,omitempty" xml:"statusStageIdentifier,omitempty"`
	Subject                *string `json:"subject,omitempty" xml:"subject,omitempty"`
	UpdateStatusAt         *int64  `json:"updateStatusAt,omitempty" xml:"updateStatusAt,omitempty"`
	WorkitemTypeIdentifier *string `json:"workitemTypeIdentifier,omitempty" xml:"workitemTypeIdentifier,omitempty"`
}

func (s CreateWorkitemResponseBodyWorkitem) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemResponseBodyWorkitem) GoString() string {
	return s.String()
}

func (s *CreateWorkitemResponseBodyWorkitem) SetAssignedTo(v string) *CreateWorkitemResponseBodyWorkitem {
	s.AssignedTo = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetCategoryIdentifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.CategoryIdentifier = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetCreator(v string) *CreateWorkitemResponseBodyWorkitem {
	s.Creator = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetDocument(v string) *CreateWorkitemResponseBodyWorkitem {
	s.Document = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetGmtCreate(v int64) *CreateWorkitemResponseBodyWorkitem {
	s.GmtCreate = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetGmtModified(v int64) *CreateWorkitemResponseBodyWorkitem {
	s.GmtModified = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetIdentifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.Identifier = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetLogicalStatus(v string) *CreateWorkitemResponseBodyWorkitem {
	s.LogicalStatus = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetModifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.Modifier = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetParentIdentifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetSerialNumber(v string) *CreateWorkitemResponseBodyWorkitem {
	s.SerialNumber = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetSpaceIdentifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.SpaceIdentifier = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetSpaceName(v string) *CreateWorkitemResponseBodyWorkitem {
	s.SpaceName = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetSpaceType(v string) *CreateWorkitemResponseBodyWorkitem {
	s.SpaceType = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetSprintIdentifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.SprintIdentifier = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetStatus(v string) *CreateWorkitemResponseBodyWorkitem {
	s.Status = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetStatusIdentifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.StatusIdentifier = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetStatusStageIdentifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.StatusStageIdentifier = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetSubject(v string) *CreateWorkitemResponseBodyWorkitem {
	s.Subject = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetUpdateStatusAt(v int64) *CreateWorkitemResponseBodyWorkitem {
	s.UpdateStatusAt = &v
	return s
}

func (s *CreateWorkitemResponseBodyWorkitem) SetWorkitemTypeIdentifier(v string) *CreateWorkitemResponseBodyWorkitem {
	s.WorkitemTypeIdentifier = &v
	return s
}

type CreateWorkitemResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWorkitemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWorkitemResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkitemResponse) SetHeaders(v map[string]*string) *CreateWorkitemResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkitemResponse) SetStatusCode(v int32) *CreateWorkitemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkitemResponse) SetBody(v *CreateWorkitemResponseBody) *CreateWorkitemResponse {
	s.Body = v
	return s
}

type CreateWorkitemCommentRequest struct {
	Content            *string `json:"content,omitempty" xml:"content,omitempty"`
	FormatType         *string `json:"formatType,omitempty" xml:"formatType,omitempty"`
	ParentId           *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	WorkitemIdentifier *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s CreateWorkitemCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemCommentRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkitemCommentRequest) SetContent(v string) *CreateWorkitemCommentRequest {
	s.Content = &v
	return s
}

func (s *CreateWorkitemCommentRequest) SetFormatType(v string) *CreateWorkitemCommentRequest {
	s.FormatType = &v
	return s
}

func (s *CreateWorkitemCommentRequest) SetParentId(v string) *CreateWorkitemCommentRequest {
	s.ParentId = &v
	return s
}

func (s *CreateWorkitemCommentRequest) SetWorkitemIdentifier(v string) *CreateWorkitemCommentRequest {
	s.WorkitemIdentifier = &v
	return s
}

type CreateWorkitemCommentResponseBody struct {
	Comment   *CreateWorkitemCommentResponseBodyComment `json:"Comment,omitempty" xml:"Comment,omitempty" type:"Struct"`
	ErrorCode *string                                   `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                   `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *string                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateWorkitemCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemCommentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkitemCommentResponseBody) SetComment(v *CreateWorkitemCommentResponseBodyComment) *CreateWorkitemCommentResponseBody {
	s.Comment = v
	return s
}

func (s *CreateWorkitemCommentResponseBody) SetErrorCode(v string) *CreateWorkitemCommentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateWorkitemCommentResponseBody) SetErrorMsg(v string) *CreateWorkitemCommentResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateWorkitemCommentResponseBody) SetRequestId(v string) *CreateWorkitemCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkitemCommentResponseBody) SetSuccess(v string) *CreateWorkitemCommentResponseBody {
	s.Success = &v
	return s
}

type CreateWorkitemCommentResponseBodyComment struct {
	Id               *int64                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	Content          *string                                       `json:"content,omitempty" xml:"content,omitempty"`
	CreateTime       *int64                                        `json:"createTime,omitempty" xml:"createTime,omitempty"`
	FormatType       *string                                       `json:"formatType,omitempty" xml:"formatType,omitempty"`
	IsTop            *bool                                         `json:"isTop,omitempty" xml:"isTop,omitempty"`
	ModifiedTime     *int64                                        `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ParentId         *int64                                        `json:"parentId,omitempty" xml:"parentId,omitempty"`
	TargetIdentifier *string                                       `json:"targetIdentifier,omitempty" xml:"targetIdentifier,omitempty"`
	TargetType       *string                                       `json:"targetType,omitempty" xml:"targetType,omitempty"`
	TopTime          *int64                                        `json:"topTime,omitempty" xml:"topTime,omitempty"`
	User             *CreateWorkitemCommentResponseBodyCommentUser `json:"user,omitempty" xml:"user,omitempty" type:"Struct"`
}

func (s CreateWorkitemCommentResponseBodyComment) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemCommentResponseBodyComment) GoString() string {
	return s.String()
}

func (s *CreateWorkitemCommentResponseBodyComment) SetId(v int64) *CreateWorkitemCommentResponseBodyComment {
	s.Id = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetContent(v string) *CreateWorkitemCommentResponseBodyComment {
	s.Content = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetCreateTime(v int64) *CreateWorkitemCommentResponseBodyComment {
	s.CreateTime = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetFormatType(v string) *CreateWorkitemCommentResponseBodyComment {
	s.FormatType = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetIsTop(v bool) *CreateWorkitemCommentResponseBodyComment {
	s.IsTop = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetModifiedTime(v int64) *CreateWorkitemCommentResponseBodyComment {
	s.ModifiedTime = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetParentId(v int64) *CreateWorkitemCommentResponseBodyComment {
	s.ParentId = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetTargetIdentifier(v string) *CreateWorkitemCommentResponseBodyComment {
	s.TargetIdentifier = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetTargetType(v string) *CreateWorkitemCommentResponseBodyComment {
	s.TargetType = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetTopTime(v int64) *CreateWorkitemCommentResponseBodyComment {
	s.TopTime = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetUser(v *CreateWorkitemCommentResponseBodyCommentUser) *CreateWorkitemCommentResponseBodyComment {
	s.User = v
	return s
}

type CreateWorkitemCommentResponseBodyCommentUser struct {
	Account     *string `json:"account,omitempty" xml:"account,omitempty"`
	Avatar      *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Identifier  *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	NickName    *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	RealName    *string `json:"realName,omitempty" xml:"realName,omitempty"`
	TargetType  *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s CreateWorkitemCommentResponseBodyCommentUser) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemCommentResponseBodyCommentUser) GoString() string {
	return s.String()
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) SetAccount(v string) *CreateWorkitemCommentResponseBodyCommentUser {
	s.Account = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) SetAvatar(v string) *CreateWorkitemCommentResponseBodyCommentUser {
	s.Avatar = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) SetDisplayName(v string) *CreateWorkitemCommentResponseBodyCommentUser {
	s.DisplayName = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) SetIdentifier(v string) *CreateWorkitemCommentResponseBodyCommentUser {
	s.Identifier = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) SetNickName(v string) *CreateWorkitemCommentResponseBodyCommentUser {
	s.NickName = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) SetRealName(v string) *CreateWorkitemCommentResponseBodyCommentUser {
	s.RealName = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) SetTargetType(v string) *CreateWorkitemCommentResponseBodyCommentUser {
	s.TargetType = &v
	return s
}

type CreateWorkitemCommentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWorkitemCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWorkitemCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemCommentResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkitemCommentResponse) SetHeaders(v map[string]*string) *CreateWorkitemCommentResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkitemCommentResponse) SetStatusCode(v int32) *CreateWorkitemCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkitemCommentResponse) SetBody(v *CreateWorkitemCommentResponseBody) *CreateWorkitemCommentResponse {
	s.Body = v
	return s
}

type CreateWorkitemEstimateRequest struct {
	Description          *string `json:"description,omitempty" xml:"description,omitempty"`
	RecordUserIdentifier *string `json:"recordUserIdentifier,omitempty" xml:"recordUserIdentifier,omitempty"`
	SpentTime            *string `json:"spentTime,omitempty" xml:"spentTime,omitempty"`
	Type                 *string `json:"type,omitempty" xml:"type,omitempty"`
	WorkitemIdentifier   *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s CreateWorkitemEstimateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemEstimateRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkitemEstimateRequest) SetDescription(v string) *CreateWorkitemEstimateRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkitemEstimateRequest) SetRecordUserIdentifier(v string) *CreateWorkitemEstimateRequest {
	s.RecordUserIdentifier = &v
	return s
}

func (s *CreateWorkitemEstimateRequest) SetSpentTime(v string) *CreateWorkitemEstimateRequest {
	s.SpentTime = &v
	return s
}

func (s *CreateWorkitemEstimateRequest) SetType(v string) *CreateWorkitemEstimateRequest {
	s.Type = &v
	return s
}

func (s *CreateWorkitemEstimateRequest) SetWorkitemIdentifier(v string) *CreateWorkitemEstimateRequest {
	s.WorkitemIdentifier = &v
	return s
}

type CreateWorkitemEstimateResponseBody struct {
	WorkitemTimeEstimate *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate `json:"WorkitemTimeEstimate,omitempty" xml:"WorkitemTimeEstimate,omitempty" type:"Struct"`
	ErrorCode            *string                                                 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg             *string                                                 `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId            *string                                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success              *bool                                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateWorkitemEstimateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemEstimateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkitemEstimateResponseBody) SetWorkitemTimeEstimate(v *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) *CreateWorkitemEstimateResponseBody {
	s.WorkitemTimeEstimate = v
	return s
}

func (s *CreateWorkitemEstimateResponseBody) SetErrorCode(v string) *CreateWorkitemEstimateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBody) SetErrorMsg(v string) *CreateWorkitemEstimateResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBody) SetRequestId(v string) *CreateWorkitemEstimateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBody) SetSuccess(v bool) *CreateWorkitemEstimateResponseBody {
	s.Success = &v
	return s
}

type CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate struct {
	Description        *string                                                           `json:"description,omitempty" xml:"description,omitempty"`
	Identifier         *string                                                           `json:"identifier,omitempty" xml:"identifier,omitempty"`
	RecordUser         *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser `json:"recordUser,omitempty" xml:"recordUser,omitempty" type:"Struct"`
	SpentTime          *int64                                                            `json:"spentTime,omitempty" xml:"spentTime,omitempty"`
	Type               *string                                                           `json:"type,omitempty" xml:"type,omitempty"`
	WorkitemIdentifier *string                                                           `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) GoString() string {
	return s.String()
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetDescription(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.Description = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetIdentifier(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.Identifier = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetRecordUser(v *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.RecordUser = v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetSpentTime(v int64) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.SpentTime = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetType(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.Type = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetWorkitemIdentifier(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.WorkitemIdentifier = &v
	return s
}

type CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser struct {
	Account         *string `json:"account,omitempty" xml:"account,omitempty"`
	Avatar          *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	DingTalkId      *string `json:"dingTalkId,omitempty" xml:"dingTalkId,omitempty"`
	DisplayName     *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	DisplayNickName *string `json:"displayNickName,omitempty" xml:"displayNickName,omitempty"`
	DisplayRealName *string `json:"displayRealName,omitempty" xml:"displayRealName,omitempty"`
	Email           *string `json:"email,omitempty" xml:"email,omitempty"`
	Gender          *string `json:"gender,omitempty" xml:"gender,omitempty"`
	Identifier      *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	IsDisabled      *bool   `json:"isDisabled,omitempty" xml:"isDisabled,omitempty"`
	Mobile          *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	NameEn          *string `json:"nameEn,omitempty" xml:"nameEn,omitempty"`
	NickName        *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	NickNamePinyin  *string `json:"nickNamePinyin,omitempty" xml:"nickNamePinyin,omitempty"`
	RealName        *string `json:"realName,omitempty" xml:"realName,omitempty"`
	RealNamePinyin  *string `json:"realNamePinyin,omitempty" xml:"realNamePinyin,omitempty"`
	Stamp           *string `json:"stamp,omitempty" xml:"stamp,omitempty"`
	TbRoleId        *string `json:"tbRoleId,omitempty" xml:"tbRoleId,omitempty"`
}

func (s CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GoString() string {
	return s.String()
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetAccount(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.Account = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetAvatar(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.Avatar = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetDingTalkId(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.DingTalkId = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetDisplayName(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.DisplayName = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetDisplayNickName(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.DisplayNickName = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetDisplayRealName(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.DisplayRealName = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetEmail(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.Email = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetGender(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.Gender = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetIdentifier(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.Identifier = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetIsDisabled(v bool) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.IsDisabled = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetMobile(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.Mobile = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetNameEn(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.NameEn = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetNickName(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.NickName = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetNickNamePinyin(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.NickNamePinyin = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetRealName(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.RealName = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetRealNamePinyin(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.RealNamePinyin = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetStamp(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.Stamp = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetTbRoleId(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.TbRoleId = &v
	return s
}

type CreateWorkitemEstimateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWorkitemEstimateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWorkitemEstimateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemEstimateResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkitemEstimateResponse) SetHeaders(v map[string]*string) *CreateWorkitemEstimateResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkitemEstimateResponse) SetStatusCode(v int32) *CreateWorkitemEstimateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkitemEstimateResponse) SetBody(v *CreateWorkitemEstimateResponseBody) *CreateWorkitemEstimateResponse {
	s.Body = v
	return s
}

type CreateWorkitemRecordRequest struct {
	ActualTime           *string `json:"actualTime,omitempty" xml:"actualTime,omitempty"`
	Description          *string `json:"description,omitempty" xml:"description,omitempty"`
	GmtEnd               *string `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	GmtStart             *string `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
	RecordUserIdentifier *string `json:"recordUserIdentifier,omitempty" xml:"recordUserIdentifier,omitempty"`
	Type                 *string `json:"type,omitempty" xml:"type,omitempty"`
	WorkitemIdentifier   *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s CreateWorkitemRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkitemRecordRequest) SetActualTime(v string) *CreateWorkitemRecordRequest {
	s.ActualTime = &v
	return s
}

func (s *CreateWorkitemRecordRequest) SetDescription(v string) *CreateWorkitemRecordRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkitemRecordRequest) SetGmtEnd(v string) *CreateWorkitemRecordRequest {
	s.GmtEnd = &v
	return s
}

func (s *CreateWorkitemRecordRequest) SetGmtStart(v string) *CreateWorkitemRecordRequest {
	s.GmtStart = &v
	return s
}

func (s *CreateWorkitemRecordRequest) SetRecordUserIdentifier(v string) *CreateWorkitemRecordRequest {
	s.RecordUserIdentifier = &v
	return s
}

func (s *CreateWorkitemRecordRequest) SetType(v string) *CreateWorkitemRecordRequest {
	s.Type = &v
	return s
}

func (s *CreateWorkitemRecordRequest) SetWorkitemIdentifier(v string) *CreateWorkitemRecordRequest {
	s.WorkitemIdentifier = &v
	return s
}

type CreateWorkitemRecordResponseBody struct {
	WorkitemTime *CreateWorkitemRecordResponseBodyWorkitemTime `json:"WorkitemTime,omitempty" xml:"WorkitemTime,omitempty" type:"Struct"`
	ErrorCode    *string                                       `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg     *string                                       `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId    *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateWorkitemRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkitemRecordResponseBody) SetWorkitemTime(v *CreateWorkitemRecordResponseBodyWorkitemTime) *CreateWorkitemRecordResponseBody {
	s.WorkitemTime = v
	return s
}

func (s *CreateWorkitemRecordResponseBody) SetErrorCode(v string) *CreateWorkitemRecordResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateWorkitemRecordResponseBody) SetErrorMsg(v string) *CreateWorkitemRecordResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateWorkitemRecordResponseBody) SetRequestId(v string) *CreateWorkitemRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkitemRecordResponseBody) SetSuccess(v bool) *CreateWorkitemRecordResponseBody {
	s.Success = &v
	return s
}

type CreateWorkitemRecordResponseBodyWorkitemTime struct {
	ActualTime         *int64                                                  `json:"actualTime,omitempty" xml:"actualTime,omitempty"`
	Description        *string                                                 `json:"description,omitempty" xml:"description,omitempty"`
	GmtEnd             *int64                                                  `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	GmtStart           *int64                                                  `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
	Identifier         *string                                                 `json:"identifier,omitempty" xml:"identifier,omitempty"`
	RecordUser         *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser `json:"recordUser,omitempty" xml:"recordUser,omitempty" type:"Struct"`
	Type               *string                                                 `json:"type,omitempty" xml:"type,omitempty"`
	WorkitemIdentifier *string                                                 `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s CreateWorkitemRecordResponseBodyWorkitemTime) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemRecordResponseBodyWorkitemTime) GoString() string {
	return s.String()
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) SetActualTime(v int64) *CreateWorkitemRecordResponseBodyWorkitemTime {
	s.ActualTime = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) SetDescription(v string) *CreateWorkitemRecordResponseBodyWorkitemTime {
	s.Description = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) SetGmtEnd(v int64) *CreateWorkitemRecordResponseBodyWorkitemTime {
	s.GmtEnd = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) SetGmtStart(v int64) *CreateWorkitemRecordResponseBodyWorkitemTime {
	s.GmtStart = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) SetIdentifier(v string) *CreateWorkitemRecordResponseBodyWorkitemTime {
	s.Identifier = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) SetRecordUser(v *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) *CreateWorkitemRecordResponseBodyWorkitemTime {
	s.RecordUser = v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) SetType(v string) *CreateWorkitemRecordResponseBodyWorkitemTime {
	s.Type = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) SetWorkitemIdentifier(v string) *CreateWorkitemRecordResponseBodyWorkitemTime {
	s.WorkitemIdentifier = &v
	return s
}

type CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser struct {
	Account         *string `json:"account,omitempty" xml:"account,omitempty"`
	Avatar          *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	DingTalkId      *string `json:"dingTalkId,omitempty" xml:"dingTalkId,omitempty"`
	DisplayName     *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	DisplayNickName *string `json:"displayNickName,omitempty" xml:"displayNickName,omitempty"`
	DisplayRealName *string `json:"displayRealName,omitempty" xml:"displayRealName,omitempty"`
	Email           *string `json:"email,omitempty" xml:"email,omitempty"`
	Gender          *string `json:"gender,omitempty" xml:"gender,omitempty"`
	Identifier      *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	IsDisabled      *bool   `json:"isDisabled,omitempty" xml:"isDisabled,omitempty"`
	Mobile          *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	NameEn          *string `json:"nameEn,omitempty" xml:"nameEn,omitempty"`
	NickName        *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	NickNamePinyin  *string `json:"nickNamePinyin,omitempty" xml:"nickNamePinyin,omitempty"`
	RealName        *string `json:"realName,omitempty" xml:"realName,omitempty"`
	RealNamePinyin  *string `json:"realNamePinyin,omitempty" xml:"realNamePinyin,omitempty"`
	Stamp           *string `json:"stamp,omitempty" xml:"stamp,omitempty"`
	TbRoleId        *string `json:"tbRoleId,omitempty" xml:"tbRoleId,omitempty"`
}

func (s CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GoString() string {
	return s.String()
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetAccount(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.Account = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetAvatar(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.Avatar = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetDingTalkId(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.DingTalkId = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetDisplayName(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.DisplayName = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetDisplayNickName(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.DisplayNickName = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetDisplayRealName(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.DisplayRealName = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetEmail(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.Email = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetGender(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.Gender = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetIdentifier(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.Identifier = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetIsDisabled(v bool) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.IsDisabled = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetMobile(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.Mobile = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetNameEn(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.NameEn = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetNickName(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.NickName = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetNickNamePinyin(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.NickNamePinyin = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetRealName(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.RealName = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetRealNamePinyin(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.RealNamePinyin = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetStamp(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.Stamp = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetTbRoleId(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.TbRoleId = &v
	return s
}

type CreateWorkitemRecordResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWorkitemRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWorkitemRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkitemRecordResponse) SetHeaders(v map[string]*string) *CreateWorkitemRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkitemRecordResponse) SetStatusCode(v int32) *CreateWorkitemRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkitemRecordResponse) SetBody(v *CreateWorkitemRecordResponseBody) *CreateWorkitemRecordResponse {
	s.Body = v
	return s
}

type CreateWorkspaceRequest struct {
	CodeUrl            *string `json:"codeUrl,omitempty" xml:"codeUrl,omitempty"`
	CodeVersion        *string `json:"codeVersion,omitempty" xml:"codeVersion,omitempty"`
	FilePath           *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	RequestFrom        *string `json:"requestFrom,omitempty" xml:"requestFrom,omitempty"`
	ResourceIdentifier *string `json:"resourceIdentifier,omitempty" xml:"resourceIdentifier,omitempty"`
	Reuse              *bool   `json:"reuse,omitempty" xml:"reuse,omitempty"`
	WorkspaceTemplate  *string `json:"workspaceTemplate,omitempty" xml:"workspaceTemplate,omitempty"`
}

func (s CreateWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) SetCodeUrl(v string) *CreateWorkspaceRequest {
	s.CodeUrl = &v
	return s
}

func (s *CreateWorkspaceRequest) SetCodeVersion(v string) *CreateWorkspaceRequest {
	s.CodeVersion = &v
	return s
}

func (s *CreateWorkspaceRequest) SetFilePath(v string) *CreateWorkspaceRequest {
	s.FilePath = &v
	return s
}

func (s *CreateWorkspaceRequest) SetName(v string) *CreateWorkspaceRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceRequest) SetRequestFrom(v string) *CreateWorkspaceRequest {
	s.RequestFrom = &v
	return s
}

func (s *CreateWorkspaceRequest) SetResourceIdentifier(v string) *CreateWorkspaceRequest {
	s.ResourceIdentifier = &v
	return s
}

func (s *CreateWorkspaceRequest) SetReuse(v bool) *CreateWorkspaceRequest {
	s.Reuse = &v
	return s
}

func (s *CreateWorkspaceRequest) SetWorkspaceTemplate(v string) *CreateWorkspaceRequest {
	s.WorkspaceTemplate = &v
	return s
}

type CreateWorkspaceResponseBody struct {
	ErrorCode    *string                               `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                               `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
	Workspace    *CreateWorkspaceResponseBodyWorkspace `json:"workspace,omitempty" xml:"workspace,omitempty" type:"Struct"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) SetErrorCode(v string) *CreateWorkspaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetErrorMessage(v string) *CreateWorkspaceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetRequestId(v string) *CreateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetSuccess(v bool) *CreateWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetWorkspace(v *CreateWorkspaceResponseBodyWorkspace) *CreateWorkspaceResponseBody {
	s.Workspace = v
	return s
}

type CreateWorkspaceResponseBodyWorkspace struct {
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator    *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
	Template   *string `json:"template,omitempty" xml:"template,omitempty"`
}

func (s CreateWorkspaceResponseBodyWorkspace) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponseBodyWorkspace) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBodyWorkspace) SetCreateTime(v string) *CreateWorkspaceResponseBodyWorkspace {
	s.CreateTime = &v
	return s
}

func (s *CreateWorkspaceResponseBodyWorkspace) SetCreator(v string) *CreateWorkspaceResponseBodyWorkspace {
	s.Creator = &v
	return s
}

func (s *CreateWorkspaceResponseBodyWorkspace) SetId(v string) *CreateWorkspaceResponseBodyWorkspace {
	s.Id = &v
	return s
}

func (s *CreateWorkspaceResponseBodyWorkspace) SetName(v string) *CreateWorkspaceResponseBodyWorkspace {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceResponseBodyWorkspace) SetStatus(v string) *CreateWorkspaceResponseBodyWorkspace {
	s.Status = &v
	return s
}

func (s *CreateWorkspaceResponseBodyWorkspace) SetTemplate(v string) *CreateWorkspaceResponseBodyWorkspace {
	s.Template = &v
	return s
}

type CreateWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponse) SetHeaders(v map[string]*string) *CreateWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkspaceResponse) SetStatusCode(v int32) *CreateWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkspaceResponse) SetBody(v *CreateWorkspaceResponseBody) *CreateWorkspaceResponse {
	s.Body = v
	return s
}

type DeleteBranchRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	BranchName     *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s DeleteBranchRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBranchRequest) GoString() string {
	return s.String()
}

func (s *DeleteBranchRequest) SetAccessToken(v string) *DeleteBranchRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteBranchRequest) SetBranchName(v string) *DeleteBranchRequest {
	s.BranchName = &v
	return s
}

func (s *DeleteBranchRequest) SetOrganizationId(v string) *DeleteBranchRequest {
	s.OrganizationId = &v
	return s
}

type DeleteBranchResponseBody struct {
	ErrorCode    *string                         `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                         `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       *DeleteBranchResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success      *string                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteBranchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBranchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBranchResponseBody) SetErrorCode(v string) *DeleteBranchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteBranchResponseBody) SetErrorMessage(v string) *DeleteBranchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteBranchResponseBody) SetRequestId(v string) *DeleteBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBranchResponseBody) SetResult(v *DeleteBranchResponseBodyResult) *DeleteBranchResponseBody {
	s.Result = v
	return s
}

func (s *DeleteBranchResponseBody) SetSuccess(v string) *DeleteBranchResponseBody {
	s.Success = &v
	return s
}

type DeleteBranchResponseBodyResult struct {
	BranchName *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
}

func (s DeleteBranchResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteBranchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteBranchResponseBodyResult) SetBranchName(v string) *DeleteBranchResponseBodyResult {
	s.BranchName = &v
	return s
}

type DeleteBranchResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBranchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBranchResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBranchResponse) GoString() string {
	return s.String()
}

func (s *DeleteBranchResponse) SetHeaders(v map[string]*string) *DeleteBranchResponse {
	s.Headers = v
	return s
}

func (s *DeleteBranchResponse) SetStatusCode(v int32) *DeleteBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBranchResponse) SetBody(v *DeleteBranchResponseBody) *DeleteBranchResponse {
	s.Body = v
	return s
}

type DeleteFileRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	BranchName     *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	CommitMessage  *string `json:"commitMessage,omitempty" xml:"commitMessage,omitempty"`
	FilePath       *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s DeleteFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileRequest) SetAccessToken(v string) *DeleteFileRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteFileRequest) SetBranchName(v string) *DeleteFileRequest {
	s.BranchName = &v
	return s
}

func (s *DeleteFileRequest) SetCommitMessage(v string) *DeleteFileRequest {
	s.CommitMessage = &v
	return s
}

func (s *DeleteFileRequest) SetFilePath(v string) *DeleteFileRequest {
	s.FilePath = &v
	return s
}

func (s *DeleteFileRequest) SetOrganizationId(v string) *DeleteFileRequest {
	s.OrganizationId = &v
	return s
}

type DeleteFileResponseBody struct {
	ErrorCode    *string                       `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                       `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       *DeleteFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success      *bool                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBody) SetErrorCode(v string) *DeleteFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteFileResponseBody) SetErrorMessage(v string) *DeleteFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFileResponseBody) SetRequestId(v string) *DeleteFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFileResponseBody) SetResult(v *DeleteFileResponseBodyResult) *DeleteFileResponseBody {
	s.Result = v
	return s
}

func (s *DeleteFileResponseBody) SetSuccess(v bool) *DeleteFileResponseBody {
	s.Success = &v
	return s
}

type DeleteFileResponseBodyResult struct {
	BranchName *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	FilePath   *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
}

func (s DeleteFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBodyResult) SetBranchName(v string) *DeleteFileResponseBodyResult {
	s.BranchName = &v
	return s
}

func (s *DeleteFileResponseBodyResult) SetFilePath(v string) *DeleteFileResponseBodyResult {
	s.FilePath = &v
	return s
}

type DeleteFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileResponse) SetHeaders(v map[string]*string) *DeleteFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileResponse) SetStatusCode(v int32) *DeleteFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFileResponse) SetBody(v *DeleteFileResponseBody) *DeleteFileResponse {
	s.Body = v
	return s
}

type DeleteFlowTagResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteFlowTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowTagResponseBody) SetErrorCode(v string) *DeleteFlowTagResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteFlowTagResponseBody) SetErrorMessage(v string) *DeleteFlowTagResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFlowTagResponseBody) SetRequestId(v string) *DeleteFlowTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowTagResponseBody) SetSuccess(v bool) *DeleteFlowTagResponseBody {
	s.Success = &v
	return s
}

type DeleteFlowTagResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFlowTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowTagResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowTagResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowTagResponse) SetHeaders(v map[string]*string) *DeleteFlowTagResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowTagResponse) SetStatusCode(v int32) *DeleteFlowTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFlowTagResponse) SetBody(v *DeleteFlowTagResponseBody) *DeleteFlowTagResponse {
	s.Body = v
	return s
}

type DeleteFlowTagGroupResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteFlowTagGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowTagGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowTagGroupResponseBody) SetErrorCode(v string) *DeleteFlowTagGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteFlowTagGroupResponseBody) SetErrorMessage(v string) *DeleteFlowTagGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFlowTagGroupResponseBody) SetRequestId(v string) *DeleteFlowTagGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowTagGroupResponseBody) SetSuccess(v bool) *DeleteFlowTagGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteFlowTagGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFlowTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowTagGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowTagGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowTagGroupResponse) SetHeaders(v map[string]*string) *DeleteFlowTagGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowTagGroupResponse) SetStatusCode(v int32) *DeleteFlowTagGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFlowTagGroupResponse) SetBody(v *DeleteFlowTagGroupResponseBody) *DeleteFlowTagGroupResponse {
	s.Body = v
	return s
}

type DeleteHostGroupResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteHostGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHostGroupResponseBody) SetErrorCode(v string) *DeleteHostGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteHostGroupResponseBody) SetErrorMessage(v string) *DeleteHostGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteHostGroupResponseBody) SetRequestId(v string) *DeleteHostGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHostGroupResponseBody) SetSuccess(v bool) *DeleteHostGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteHostGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteHostGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteHostGroupResponse) SetHeaders(v map[string]*string) *DeleteHostGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteHostGroupResponse) SetStatusCode(v int32) *DeleteHostGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHostGroupResponse) SetBody(v *DeleteHostGroupResponseBody) *DeleteHostGroupResponse {
	s.Body = v
	return s
}

type DeletePipelineResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeletePipelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePipelineResponseBody) SetErrorCode(v string) *DeletePipelineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeletePipelineResponseBody) SetErrorMessage(v string) *DeletePipelineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeletePipelineResponseBody) SetRequestId(v string) *DeletePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePipelineResponseBody) SetSuccess(v bool) *DeletePipelineResponseBody {
	s.Success = &v
	return s
}

type DeletePipelineResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineResponse) GoString() string {
	return s.String()
}

func (s *DeletePipelineResponse) SetHeaders(v map[string]*string) *DeletePipelineResponse {
	s.Headers = v
	return s
}

func (s *DeletePipelineResponse) SetStatusCode(v int32) *DeletePipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePipelineResponse) SetBody(v *DeletePipelineResponseBody) *DeletePipelineResponse {
	s.Body = v
	return s
}

type DeletePipelineGroupResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeletePipelineGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePipelineGroupResponseBody) SetErrorCode(v string) *DeletePipelineGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeletePipelineGroupResponseBody) SetErrorMessage(v string) *DeletePipelineGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeletePipelineGroupResponseBody) SetRequestId(v string) *DeletePipelineGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePipelineGroupResponseBody) SetSuccess(v bool) *DeletePipelineGroupResponseBody {
	s.Success = &v
	return s
}

type DeletePipelineGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePipelineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePipelineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineGroupResponse) GoString() string {
	return s.String()
}

func (s *DeletePipelineGroupResponse) SetHeaders(v map[string]*string) *DeletePipelineGroupResponse {
	s.Headers = v
	return s
}

func (s *DeletePipelineGroupResponse) SetStatusCode(v int32) *DeletePipelineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePipelineGroupResponse) SetBody(v *DeletePipelineGroupResponseBody) *DeletePipelineGroupResponse {
	s.Body = v
	return s
}

type DeleteProjectRequest struct {
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetIdentifier(v string) *DeleteProjectRequest {
	s.Identifier = &v
	return s
}

type DeleteProjectResponseBody struct {
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) SetErrorCode(v string) *DeleteProjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteProjectResponseBody) SetErrorMsg(v string) *DeleteProjectResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectResponseBody) SetResult(v bool) *DeleteProjectResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteProjectResponseBody) SetSuccess(v bool) *DeleteProjectResponseBody {
	s.Success = &v
	return s
}

type DeleteProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) SetHeaders(v map[string]*string) *DeleteProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectResponse) SetStatusCode(v int32) *DeleteProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
	return s
}

type DeleteProtectedBranchRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s DeleteProtectedBranchRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtectedBranchRequest) GoString() string {
	return s.String()
}

func (s *DeleteProtectedBranchRequest) SetAccessToken(v string) *DeleteProtectedBranchRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteProtectedBranchRequest) SetOrganizationId(v string) *DeleteProtectedBranchRequest {
	s.OrganizationId = &v
	return s
}

type DeleteProtectedBranchResponseBody struct {
	ErrorCode    *string                                  `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                  `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       *DeleteProtectedBranchResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success      *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteProtectedBranchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtectedBranchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProtectedBranchResponseBody) SetErrorCode(v string) *DeleteProtectedBranchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteProtectedBranchResponseBody) SetErrorMessage(v string) *DeleteProtectedBranchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteProtectedBranchResponseBody) SetRequestId(v string) *DeleteProtectedBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProtectedBranchResponseBody) SetResult(v *DeleteProtectedBranchResponseBodyResult) *DeleteProtectedBranchResponseBody {
	s.Result = v
	return s
}

func (s *DeleteProtectedBranchResponseBody) SetSuccess(v bool) *DeleteProtectedBranchResponseBody {
	s.Success = &v
	return s
}

type DeleteProtectedBranchResponseBodyResult struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteProtectedBranchResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtectedBranchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteProtectedBranchResponseBodyResult) SetResult(v bool) *DeleteProtectedBranchResponseBodyResult {
	s.Result = &v
	return s
}

type DeleteProtectedBranchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProtectedBranchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProtectedBranchResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtectedBranchResponse) GoString() string {
	return s.String()
}

func (s *DeleteProtectedBranchResponse) SetHeaders(v map[string]*string) *DeleteProtectedBranchResponse {
	s.Headers = v
	return s
}

func (s *DeleteProtectedBranchResponse) SetStatusCode(v int32) *DeleteProtectedBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProtectedBranchResponse) SetBody(v *DeleteProtectedBranchResponseBody) *DeleteProtectedBranchResponse {
	s.Body = v
	return s
}

type DeleteRepositoryRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	Reason         *string `json:"reason,omitempty" xml:"reason,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s DeleteRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryRequest) SetAccessToken(v string) *DeleteRepositoryRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteRepositoryRequest) SetReason(v string) *DeleteRepositoryRequest {
	s.Reason = &v
	return s
}

func (s *DeleteRepositoryRequest) SetOrganizationId(v string) *DeleteRepositoryRequest {
	s.OrganizationId = &v
	return s
}

type DeleteRepositoryResponseBody struct {
	ErrorCode    *string                             `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                             `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       *DeleteRepositoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success      *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryResponseBody) SetErrorCode(v string) *DeleteRepositoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryResponseBody) SetErrorMessage(v string) *DeleteRepositoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryResponseBody) SetRequestId(v string) *DeleteRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryResponseBody) SetResult(v *DeleteRepositoryResponseBodyResult) *DeleteRepositoryResponseBody {
	s.Result = v
	return s
}

func (s *DeleteRepositoryResponseBody) SetSuccess(v bool) *DeleteRepositoryResponseBody {
	s.Success = &v
	return s
}

type DeleteRepositoryResponseBodyResult struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteRepositoryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryResponseBodyResult) SetResult(v bool) *DeleteRepositoryResponseBodyResult {
	s.Result = &v
	return s
}

type DeleteRepositoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryResponse) SetHeaders(v map[string]*string) *DeleteRepositoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepositoryResponse) SetStatusCode(v int32) *DeleteRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepositoryResponse) SetBody(v *DeleteRepositoryResponseBody) *DeleteRepositoryResponse {
	s.Body = v
	return s
}

type DeleteResourceMemberResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteResourceMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceMemberResponseBody) SetErrorCode(v string) *DeleteResourceMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteResourceMemberResponseBody) SetErrorMessage(v string) *DeleteResourceMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteResourceMemberResponseBody) SetRequestId(v string) *DeleteResourceMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceMemberResponseBody) SetSuccess(v bool) *DeleteResourceMemberResponseBody {
	s.Success = &v
	return s
}

type DeleteResourceMemberResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteResourceMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteResourceMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceMemberResponse) SetHeaders(v map[string]*string) *DeleteResourceMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceMemberResponse) SetStatusCode(v int32) *DeleteResourceMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceMemberResponse) SetBody(v *DeleteResourceMemberResponseBody) *DeleteResourceMemberResponse {
	s.Body = v
	return s
}

type DeleteVariableGroupResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteVariableGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVariableGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVariableGroupResponseBody) SetErrorCode(v string) *DeleteVariableGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteVariableGroupResponseBody) SetErrorMessage(v string) *DeleteVariableGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteVariableGroupResponseBody) SetRequestId(v string) *DeleteVariableGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVariableGroupResponseBody) SetSuccess(v bool) *DeleteVariableGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteVariableGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVariableGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVariableGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVariableGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteVariableGroupResponse) SetHeaders(v map[string]*string) *DeleteVariableGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteVariableGroupResponse) SetStatusCode(v int32) *DeleteVariableGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVariableGroupResponse) SetBody(v *DeleteVariableGroupResponseBody) *DeleteVariableGroupResponse {
	s.Body = v
	return s
}

type DeleteWorkitemAllCommentRequest struct {
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
}

func (s DeleteWorkitemAllCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkitemAllCommentRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkitemAllCommentRequest) SetIdentifier(v string) *DeleteWorkitemAllCommentRequest {
	s.Identifier = &v
	return s
}

type DeleteWorkitemAllCommentResponseBody struct {
	DeleteFlag *bool   `json:"deleteFlag,omitempty" xml:"deleteFlag,omitempty"`
	ErrorCode  *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg   *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success    *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteWorkitemAllCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkitemAllCommentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkitemAllCommentResponseBody) SetDeleteFlag(v bool) *DeleteWorkitemAllCommentResponseBody {
	s.DeleteFlag = &v
	return s
}

func (s *DeleteWorkitemAllCommentResponseBody) SetErrorCode(v string) *DeleteWorkitemAllCommentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteWorkitemAllCommentResponseBody) SetErrorMsg(v string) *DeleteWorkitemAllCommentResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteWorkitemAllCommentResponseBody) SetRequestId(v string) *DeleteWorkitemAllCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkitemAllCommentResponseBody) SetSuccess(v string) *DeleteWorkitemAllCommentResponseBody {
	s.Success = &v
	return s
}

type DeleteWorkitemAllCommentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteWorkitemAllCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWorkitemAllCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkitemAllCommentResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkitemAllCommentResponse) SetHeaders(v map[string]*string) *DeleteWorkitemAllCommentResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkitemAllCommentResponse) SetStatusCode(v int32) *DeleteWorkitemAllCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkitemAllCommentResponse) SetBody(v *DeleteWorkitemAllCommentResponseBody) *DeleteWorkitemAllCommentResponse {
	s.Body = v
	return s
}

type DeleteWorkitemCommentRequest struct {
	CommentId  *int64  `json:"commentId,omitempty" xml:"commentId,omitempty"`
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
}

func (s DeleteWorkitemCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkitemCommentRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkitemCommentRequest) SetCommentId(v int64) *DeleteWorkitemCommentRequest {
	s.CommentId = &v
	return s
}

func (s *DeleteWorkitemCommentRequest) SetIdentifier(v string) *DeleteWorkitemCommentRequest {
	s.Identifier = &v
	return s
}

type DeleteWorkitemCommentResponseBody struct {
	DeleteFlag *bool   `json:"deleteFlag,omitempty" xml:"deleteFlag,omitempty"`
	ErrorCode  *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg   *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success    *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteWorkitemCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkitemCommentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkitemCommentResponseBody) SetDeleteFlag(v bool) *DeleteWorkitemCommentResponseBody {
	s.DeleteFlag = &v
	return s
}

func (s *DeleteWorkitemCommentResponseBody) SetErrorCode(v string) *DeleteWorkitemCommentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteWorkitemCommentResponseBody) SetErrorMsg(v string) *DeleteWorkitemCommentResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteWorkitemCommentResponseBody) SetRequestId(v string) *DeleteWorkitemCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkitemCommentResponseBody) SetSuccess(v string) *DeleteWorkitemCommentResponseBody {
	s.Success = &v
	return s
}

type DeleteWorkitemCommentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteWorkitemCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWorkitemCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkitemCommentResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkitemCommentResponse) SetHeaders(v map[string]*string) *DeleteWorkitemCommentResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkitemCommentResponse) SetStatusCode(v int32) *DeleteWorkitemCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkitemCommentResponse) SetBody(v *DeleteWorkitemCommentResponseBody) *DeleteWorkitemCommentResponse {
	s.Body = v
	return s
}

type FrozenWorkspaceResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s FrozenWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FrozenWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *FrozenWorkspaceResponseBody) SetErrorCode(v string) *FrozenWorkspaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *FrozenWorkspaceResponseBody) SetErrorMessage(v string) *FrozenWorkspaceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *FrozenWorkspaceResponseBody) SetRequestId(v string) *FrozenWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *FrozenWorkspaceResponseBody) SetSuccess(v bool) *FrozenWorkspaceResponseBody {
	s.Success = &v
	return s
}

type FrozenWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FrozenWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FrozenWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s FrozenWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *FrozenWorkspaceResponse) SetHeaders(v map[string]*string) *FrozenWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *FrozenWorkspaceResponse) SetStatusCode(v int32) *FrozenWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *FrozenWorkspaceResponse) SetBody(v *FrozenWorkspaceResponseBody) *FrozenWorkspaceResponse {
	s.Body = v
	return s
}

type GetBranchInfoRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	BranchName     *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s GetBranchInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoRequest) GoString() string {
	return s.String()
}

func (s *GetBranchInfoRequest) SetAccessToken(v string) *GetBranchInfoRequest {
	s.AccessToken = &v
	return s
}

func (s *GetBranchInfoRequest) SetBranchName(v string) *GetBranchInfoRequest {
	s.BranchName = &v
	return s
}

func (s *GetBranchInfoRequest) SetOrganizationId(v string) *GetBranchInfoRequest {
	s.OrganizationId = &v
	return s
}

type GetBranchInfoResponseBody struct {
	ErrorCode    *string                          `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                          `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       *GetBranchInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success      *bool                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetBranchInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBody) SetErrorCode(v string) *GetBranchInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetBranchInfoResponseBody) SetErrorMessage(v string) *GetBranchInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetBranchInfoResponseBody) SetRequestId(v string) *GetBranchInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBranchInfoResponseBody) SetResult(v *GetBranchInfoResponseBodyResult) *GetBranchInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetBranchInfoResponseBody) SetSuccess(v bool) *GetBranchInfoResponseBody {
	s.Success = &v
	return s
}

type GetBranchInfoResponseBodyResult struct {
	Commit    *GetBranchInfoResponseBodyResultCommit `json:"commit,omitempty" xml:"commit,omitempty" type:"Struct"`
	Name      *string                                `json:"name,omitempty" xml:"name,omitempty"`
	Protected *string                                `json:"protected,omitempty" xml:"protected,omitempty"`
}

func (s GetBranchInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBodyResult) SetCommit(v *GetBranchInfoResponseBodyResultCommit) *GetBranchInfoResponseBodyResult {
	s.Commit = v
	return s
}

func (s *GetBranchInfoResponseBodyResult) SetName(v string) *GetBranchInfoResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetBranchInfoResponseBodyResult) SetProtected(v string) *GetBranchInfoResponseBodyResult {
	s.Protected = &v
	return s
}

type GetBranchInfoResponseBodyResultCommit struct {
	Author         *GetBranchInfoResponseBodyResultCommitAuthor    `json:"author,omitempty" xml:"author,omitempty" type:"Struct"`
	AuthorEmail    *string                                         `json:"authorEmail,omitempty" xml:"authorEmail,omitempty"`
	AuthorName     *string                                         `json:"authorName,omitempty" xml:"authorName,omitempty"`
	AuthoredDate   *string                                         `json:"authoredDate,omitempty" xml:"authoredDate,omitempty"`
	CommentsCount  *int64                                          `json:"commentsCount,omitempty" xml:"commentsCount,omitempty"`
	CommittedDate  *string                                         `json:"committedDate,omitempty" xml:"committedDate,omitempty"`
	Committer      *GetBranchInfoResponseBodyResultCommitCommitter `json:"committer,omitempty" xml:"committer,omitempty" type:"Struct"`
	CommitterEmail *string                                         `json:"committerEmail,omitempty" xml:"committerEmail,omitempty"`
	CommitterName  *string                                         `json:"committerName,omitempty" xml:"committerName,omitempty"`
	CreatedAt      *string                                         `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Id             *string                                         `json:"id,omitempty" xml:"id,omitempty"`
	Message        *string                                         `json:"message,omitempty" xml:"message,omitempty"`
	ParentIds      []*string                                       `json:"parentIds,omitempty" xml:"parentIds,omitempty" type:"Repeated"`
	ShortId        *string                                         `json:"shortId,omitempty" xml:"shortId,omitempty"`
	Signature      *GetBranchInfoResponseBodyResultCommitSignature `json:"signature,omitempty" xml:"signature,omitempty" type:"Struct"`
	Title          *string                                         `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetBranchInfoResponseBodyResultCommit) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoResponseBodyResultCommit) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBodyResultCommit) SetAuthor(v *GetBranchInfoResponseBodyResultCommitAuthor) *GetBranchInfoResponseBodyResultCommit {
	s.Author = v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetAuthorEmail(v string) *GetBranchInfoResponseBodyResultCommit {
	s.AuthorEmail = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetAuthorName(v string) *GetBranchInfoResponseBodyResultCommit {
	s.AuthorName = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetAuthoredDate(v string) *GetBranchInfoResponseBodyResultCommit {
	s.AuthoredDate = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetCommentsCount(v int64) *GetBranchInfoResponseBodyResultCommit {
	s.CommentsCount = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetCommittedDate(v string) *GetBranchInfoResponseBodyResultCommit {
	s.CommittedDate = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetCommitter(v *GetBranchInfoResponseBodyResultCommitCommitter) *GetBranchInfoResponseBodyResultCommit {
	s.Committer = v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetCommitterEmail(v string) *GetBranchInfoResponseBodyResultCommit {
	s.CommitterEmail = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetCommitterName(v string) *GetBranchInfoResponseBodyResultCommit {
	s.CommitterName = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetCreatedAt(v string) *GetBranchInfoResponseBodyResultCommit {
	s.CreatedAt = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetId(v string) *GetBranchInfoResponseBodyResultCommit {
	s.Id = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetMessage(v string) *GetBranchInfoResponseBodyResultCommit {
	s.Message = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetParentIds(v []*string) *GetBranchInfoResponseBodyResultCommit {
	s.ParentIds = v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetShortId(v string) *GetBranchInfoResponseBodyResultCommit {
	s.ShortId = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetSignature(v *GetBranchInfoResponseBodyResultCommitSignature) *GetBranchInfoResponseBodyResultCommit {
	s.Signature = v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetTitle(v string) *GetBranchInfoResponseBodyResultCommit {
	s.Title = &v
	return s
}

type GetBranchInfoResponseBodyResultCommitAuthor struct {
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Email     *string `json:"email,omitempty" xml:"email,omitempty"`
	ExternUid *string `json:"externUid,omitempty" xml:"externUid,omitempty"`
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	State     *string `json:"state,omitempty" xml:"state,omitempty"`
	TbUserId  *string `json:"tbUserId,omitempty" xml:"tbUserId,omitempty"`
	Username  *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetBranchInfoResponseBodyResultCommitAuthor) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoResponseBodyResultCommitAuthor) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) SetAvatarUrl(v string) *GetBranchInfoResponseBodyResultCommitAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) SetEmail(v string) *GetBranchInfoResponseBodyResultCommitAuthor {
	s.Email = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) SetExternUid(v string) *GetBranchInfoResponseBodyResultCommitAuthor {
	s.ExternUid = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) SetId(v int64) *GetBranchInfoResponseBodyResultCommitAuthor {
	s.Id = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) SetName(v string) *GetBranchInfoResponseBodyResultCommitAuthor {
	s.Name = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) SetState(v string) *GetBranchInfoResponseBodyResultCommitAuthor {
	s.State = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) SetTbUserId(v string) *GetBranchInfoResponseBodyResultCommitAuthor {
	s.TbUserId = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) SetUsername(v string) *GetBranchInfoResponseBodyResultCommitAuthor {
	s.Username = &v
	return s
}

type GetBranchInfoResponseBodyResultCommitCommitter struct {
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Email     *string `json:"email,omitempty" xml:"email,omitempty"`
	ExternUid *string `json:"externUid,omitempty" xml:"externUid,omitempty"`
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	State     *string `json:"state,omitempty" xml:"state,omitempty"`
	TbUserId  *string `json:"tbUserId,omitempty" xml:"tbUserId,omitempty"`
	Username  *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetBranchInfoResponseBodyResultCommitCommitter) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoResponseBodyResultCommitCommitter) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) SetAvatarUrl(v string) *GetBranchInfoResponseBodyResultCommitCommitter {
	s.AvatarUrl = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) SetEmail(v string) *GetBranchInfoResponseBodyResultCommitCommitter {
	s.Email = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) SetExternUid(v string) *GetBranchInfoResponseBodyResultCommitCommitter {
	s.ExternUid = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) SetId(v int64) *GetBranchInfoResponseBodyResultCommitCommitter {
	s.Id = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) SetName(v string) *GetBranchInfoResponseBodyResultCommitCommitter {
	s.Name = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) SetState(v string) *GetBranchInfoResponseBodyResultCommitCommitter {
	s.State = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) SetTbUserId(v string) *GetBranchInfoResponseBodyResultCommitCommitter {
	s.TbUserId = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) SetUsername(v string) *GetBranchInfoResponseBodyResultCommitCommitter {
	s.Username = &v
	return s
}

type GetBranchInfoResponseBodyResultCommitSignature struct {
	GpgKeyId           *string `json:"gpgKeyId,omitempty" xml:"gpgKeyId,omitempty"`
	VerificationStatus *string `json:"verificationStatus,omitempty" xml:"verificationStatus,omitempty"`
}

func (s GetBranchInfoResponseBodyResultCommitSignature) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoResponseBodyResultCommitSignature) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBodyResultCommitSignature) SetGpgKeyId(v string) *GetBranchInfoResponseBodyResultCommitSignature {
	s.GpgKeyId = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitSignature) SetVerificationStatus(v string) *GetBranchInfoResponseBodyResultCommitSignature {
	s.VerificationStatus = &v
	return s
}

type GetBranchInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBranchInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBranchInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoResponse) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponse) SetHeaders(v map[string]*string) *GetBranchInfoResponse {
	s.Headers = v
	return s
}

func (s *GetBranchInfoResponse) SetStatusCode(v int32) *GetBranchInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBranchInfoResponse) SetBody(v *GetBranchInfoResponseBody) *GetBranchInfoResponse {
	s.Body = v
	return s
}

type GetCodeupOrganizationRequest struct {
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
}

func (s GetCodeupOrganizationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCodeupOrganizationRequest) GoString() string {
	return s.String()
}

func (s *GetCodeupOrganizationRequest) SetAccessToken(v string) *GetCodeupOrganizationRequest {
	s.AccessToken = &v
	return s
}

type GetCodeupOrganizationResponseBody struct {
	ErrorCode    *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetCodeupOrganizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCodeupOrganizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCodeupOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *GetCodeupOrganizationResponseBody) SetErrorCode(v string) *GetCodeupOrganizationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetCodeupOrganizationResponseBody) SetErrorMessage(v string) *GetCodeupOrganizationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCodeupOrganizationResponseBody) SetRequestId(v string) *GetCodeupOrganizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCodeupOrganizationResponseBody) SetResult(v *GetCodeupOrganizationResponseBodyResult) *GetCodeupOrganizationResponseBody {
	s.Result = v
	return s
}

func (s *GetCodeupOrganizationResponseBody) SetSuccess(v bool) *GetCodeupOrganizationResponseBody {
	s.Success = &v
	return s
}

type GetCodeupOrganizationResponseBodyResult struct {
	CreatedAt      *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	NamespaceId    *int64  `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Path           *string `json:"Path,omitempty" xml:"Path,omitempty"`
	UpdatedAt      *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	UserRole       *string `json:"UserRole,omitempty" xml:"UserRole,omitempty"`
}

func (s GetCodeupOrganizationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetCodeupOrganizationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCodeupOrganizationResponseBodyResult) SetCreatedAt(v string) *GetCodeupOrganizationResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetId(v int64) *GetCodeupOrganizationResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetNamespaceId(v int64) *GetCodeupOrganizationResponseBodyResult {
	s.NamespaceId = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetOrganizationId(v string) *GetCodeupOrganizationResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetPath(v string) *GetCodeupOrganizationResponseBodyResult {
	s.Path = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetUpdatedAt(v string) *GetCodeupOrganizationResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetUserRole(v string) *GetCodeupOrganizationResponseBodyResult {
	s.UserRole = &v
	return s
}

type GetCodeupOrganizationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCodeupOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCodeupOrganizationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCodeupOrganizationResponse) GoString() string {
	return s.String()
}

func (s *GetCodeupOrganizationResponse) SetHeaders(v map[string]*string) *GetCodeupOrganizationResponse {
	s.Headers = v
	return s
}

func (s *GetCodeupOrganizationResponse) SetStatusCode(v int32) *GetCodeupOrganizationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCodeupOrganizationResponse) SetBody(v *GetCodeupOrganizationResponseBody) *GetCodeupOrganizationResponse {
	s.Body = v
	return s
}

type GetCustomFieldOptionRequest struct {
	SpaceIdentifier        *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	SpaceType              *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	WorkitemTypeIdentifier *string `json:"workitemTypeIdentifier,omitempty" xml:"workitemTypeIdentifier,omitempty"`
}

func (s GetCustomFieldOptionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomFieldOptionRequest) GoString() string {
	return s.String()
}

func (s *GetCustomFieldOptionRequest) SetSpaceIdentifier(v string) *GetCustomFieldOptionRequest {
	s.SpaceIdentifier = &v
	return s
}

func (s *GetCustomFieldOptionRequest) SetSpaceType(v string) *GetCustomFieldOptionRequest {
	s.SpaceType = &v
	return s
}

func (s *GetCustomFieldOptionRequest) SetWorkitemTypeIdentifier(v string) *GetCustomFieldOptionRequest {
	s.WorkitemTypeIdentifier = &v
	return s
}

type GetCustomFieldOptionResponseBody struct {
	ErrorCode *string                                   `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                   `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Fileds    []*GetCustomFieldOptionResponseBodyFileds `json:"fileds,omitempty" xml:"fileds,omitempty" type:"Repeated"`
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetCustomFieldOptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomFieldOptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomFieldOptionResponseBody) SetErrorCode(v string) *GetCustomFieldOptionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetCustomFieldOptionResponseBody) SetErrorMsg(v string) *GetCustomFieldOptionResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetCustomFieldOptionResponseBody) SetFileds(v []*GetCustomFieldOptionResponseBodyFileds) *GetCustomFieldOptionResponseBody {
	s.Fileds = v
	return s
}

func (s *GetCustomFieldOptionResponseBody) SetRequestId(v string) *GetCustomFieldOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomFieldOptionResponseBody) SetSuccess(v bool) *GetCustomFieldOptionResponseBody {
	s.Success = &v
	return s
}

type GetCustomFieldOptionResponseBodyFileds struct {
	DisplayValue    *string `json:"displayValue,omitempty" xml:"displayValue,omitempty"`
	FieldIdentifier *string `json:"fieldIdentifier,omitempty" xml:"fieldIdentifier,omitempty"`
	Identifier      *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Level           *int64  `json:"level,omitempty" xml:"level,omitempty"`
	Position        *int64  `json:"position,omitempty" xml:"position,omitempty"`
	Value           *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueEn         *string `json:"valueEn,omitempty" xml:"valueEn,omitempty"`
}

func (s GetCustomFieldOptionResponseBodyFileds) String() string {
	return tea.Prettify(s)
}

func (s GetCustomFieldOptionResponseBodyFileds) GoString() string {
	return s.String()
}

func (s *GetCustomFieldOptionResponseBodyFileds) SetDisplayValue(v string) *GetCustomFieldOptionResponseBodyFileds {
	s.DisplayValue = &v
	return s
}

func (s *GetCustomFieldOptionResponseBodyFileds) SetFieldIdentifier(v string) *GetCustomFieldOptionResponseBodyFileds {
	s.FieldIdentifier = &v
	return s
}

func (s *GetCustomFieldOptionResponseBodyFileds) SetIdentifier(v string) *GetCustomFieldOptionResponseBodyFileds {
	s.Identifier = &v
	return s
}

func (s *GetCustomFieldOptionResponseBodyFileds) SetLevel(v int64) *GetCustomFieldOptionResponseBodyFileds {
	s.Level = &v
	return s
}

func (s *GetCustomFieldOptionResponseBodyFileds) SetPosition(v int64) *GetCustomFieldOptionResponseBodyFileds {
	s.Position = &v
	return s
}

func (s *GetCustomFieldOptionResponseBodyFileds) SetValue(v string) *GetCustomFieldOptionResponseBodyFileds {
	s.Value = &v
	return s
}

func (s *GetCustomFieldOptionResponseBodyFileds) SetValueEn(v string) *GetCustomFieldOptionResponseBodyFileds {
	s.ValueEn = &v
	return s
}

type GetCustomFieldOptionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCustomFieldOptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCustomFieldOptionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomFieldOptionResponse) GoString() string {
	return s.String()
}

func (s *GetCustomFieldOptionResponse) SetHeaders(v map[string]*string) *GetCustomFieldOptionResponse {
	s.Headers = v
	return s
}

func (s *GetCustomFieldOptionResponse) SetStatusCode(v int32) *GetCustomFieldOptionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomFieldOptionResponse) SetBody(v *GetCustomFieldOptionResponseBody) *GetCustomFieldOptionResponse {
	s.Body = v
	return s
}

type GetFileBlobsRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	FilePath       *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	From           *int64  `json:"from,omitempty" xml:"from,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	Ref            *string `json:"ref,omitempty" xml:"ref,omitempty"`
	To             *int64  `json:"to,omitempty" xml:"to,omitempty"`
}

func (s GetFileBlobsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileBlobsRequest) GoString() string {
	return s.String()
}

func (s *GetFileBlobsRequest) SetAccessToken(v string) *GetFileBlobsRequest {
	s.AccessToken = &v
	return s
}

func (s *GetFileBlobsRequest) SetFilePath(v string) *GetFileBlobsRequest {
	s.FilePath = &v
	return s
}

func (s *GetFileBlobsRequest) SetFrom(v int64) *GetFileBlobsRequest {
	s.From = &v
	return s
}

func (s *GetFileBlobsRequest) SetOrganizationId(v string) *GetFileBlobsRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetFileBlobsRequest) SetRef(v string) *GetFileBlobsRequest {
	s.Ref = &v
	return s
}

func (s *GetFileBlobsRequest) SetTo(v int64) *GetFileBlobsRequest {
	s.To = &v
	return s
}

type GetFileBlobsResponseBody struct {
	ErrorCode    *string                         `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                         `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       *GetFileBlobsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success      *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetFileBlobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileBlobsResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileBlobsResponseBody) SetErrorCode(v string) *GetFileBlobsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetFileBlobsResponseBody) SetErrorMessage(v string) *GetFileBlobsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetFileBlobsResponseBody) SetRequestId(v string) *GetFileBlobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileBlobsResponseBody) SetResult(v *GetFileBlobsResponseBodyResult) *GetFileBlobsResponseBody {
	s.Result = v
	return s
}

func (s *GetFileBlobsResponseBody) SetSuccess(v bool) *GetFileBlobsResponseBody {
	s.Success = &v
	return s
}

type GetFileBlobsResponseBodyResult struct {
	Content    *string `json:"content,omitempty" xml:"content,omitempty"`
	TotalLines *int32  `json:"totalLines,omitempty" xml:"totalLines,omitempty"`
}

func (s GetFileBlobsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFileBlobsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFileBlobsResponseBodyResult) SetContent(v string) *GetFileBlobsResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetFileBlobsResponseBodyResult) SetTotalLines(v int32) *GetFileBlobsResponseBodyResult {
	s.TotalLines = &v
	return s
}

type GetFileBlobsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFileBlobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileBlobsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileBlobsResponse) GoString() string {
	return s.String()
}

func (s *GetFileBlobsResponse) SetHeaders(v map[string]*string) *GetFileBlobsResponse {
	s.Headers = v
	return s
}

func (s *GetFileBlobsResponse) SetStatusCode(v int32) *GetFileBlobsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileBlobsResponse) SetBody(v *GetFileBlobsResponseBody) *GetFileBlobsResponse {
	s.Body = v
	return s
}

type GetFileLastCommitRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	Filepath       *string `json:"filepath,omitempty" xml:"filepath,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	Sha            *string `json:"sha,omitempty" xml:"sha,omitempty"`
	ShowSignature  *bool   `json:"showSignature,omitempty" xml:"showSignature,omitempty"`
}

func (s GetFileLastCommitRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileLastCommitRequest) GoString() string {
	return s.String()
}

func (s *GetFileLastCommitRequest) SetAccessToken(v string) *GetFileLastCommitRequest {
	s.AccessToken = &v
	return s
}

func (s *GetFileLastCommitRequest) SetFilepath(v string) *GetFileLastCommitRequest {
	s.Filepath = &v
	return s
}

func (s *GetFileLastCommitRequest) SetOrganizationId(v string) *GetFileLastCommitRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetFileLastCommitRequest) SetSha(v string) *GetFileLastCommitRequest {
	s.Sha = &v
	return s
}

func (s *GetFileLastCommitRequest) SetShowSignature(v bool) *GetFileLastCommitRequest {
	s.ShowSignature = &v
	return s
}

type GetFileLastCommitResponseBody struct {
	ErrorCode    *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetFileLastCommitResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFileLastCommitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileLastCommitResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileLastCommitResponseBody) SetErrorCode(v string) *GetFileLastCommitResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetFileLastCommitResponseBody) SetErrorMessage(v string) *GetFileLastCommitResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetFileLastCommitResponseBody) SetRequestId(v string) *GetFileLastCommitResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileLastCommitResponseBody) SetResult(v *GetFileLastCommitResponseBodyResult) *GetFileLastCommitResponseBody {
	s.Result = v
	return s
}

func (s *GetFileLastCommitResponseBody) SetSuccess(v bool) *GetFileLastCommitResponseBody {
	s.Success = &v
	return s
}

type GetFileLastCommitResponseBodyResult struct {
	AuthorDate     *string                                       `json:"AuthorDate,omitempty" xml:"AuthorDate,omitempty"`
	AuthorEmail    *string                                       `json:"AuthorEmail,omitempty" xml:"AuthorEmail,omitempty"`
	AuthorName     *string                                       `json:"AuthorName,omitempty" xml:"AuthorName,omitempty"`
	CommittedDate  *string                                       `json:"CommittedDate,omitempty" xml:"CommittedDate,omitempty"`
	CommitterEmail *string                                       `json:"CommitterEmail,omitempty" xml:"CommitterEmail,omitempty"`
	CommitterName  *string                                       `json:"CommitterName,omitempty" xml:"CommitterName,omitempty"`
	CreatedAt      *string                                       `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id             *string                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	ParentIds      []*string                                     `json:"ParentIds,omitempty" xml:"ParentIds,omitempty" type:"Repeated"`
	ShortId        *string                                       `json:"ShortId,omitempty" xml:"ShortId,omitempty"`
	Signature      *GetFileLastCommitResponseBodyResultSignature `json:"Signature,omitempty" xml:"Signature,omitempty" type:"Struct"`
	Title          *string                                       `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetFileLastCommitResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFileLastCommitResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFileLastCommitResponseBodyResult) SetAuthorDate(v string) *GetFileLastCommitResponseBodyResult {
	s.AuthorDate = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetAuthorEmail(v string) *GetFileLastCommitResponseBodyResult {
	s.AuthorEmail = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetAuthorName(v string) *GetFileLastCommitResponseBodyResult {
	s.AuthorName = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetCommittedDate(v string) *GetFileLastCommitResponseBodyResult {
	s.CommittedDate = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetCommitterEmail(v string) *GetFileLastCommitResponseBodyResult {
	s.CommitterEmail = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetCommitterName(v string) *GetFileLastCommitResponseBodyResult {
	s.CommitterName = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetCreatedAt(v string) *GetFileLastCommitResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetId(v string) *GetFileLastCommitResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetMessage(v string) *GetFileLastCommitResponseBodyResult {
	s.Message = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetParentIds(v []*string) *GetFileLastCommitResponseBodyResult {
	s.ParentIds = v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetShortId(v string) *GetFileLastCommitResponseBodyResult {
	s.ShortId = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetSignature(v *GetFileLastCommitResponseBodyResultSignature) *GetFileLastCommitResponseBodyResult {
	s.Signature = v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetTitle(v string) *GetFileLastCommitResponseBodyResult {
	s.Title = &v
	return s
}

type GetFileLastCommitResponseBodyResultSignature struct {
	GpgKeyId           *string `json:"GpgKeyId,omitempty" xml:"GpgKeyId,omitempty"`
	VerificationStatus *string `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
}

func (s GetFileLastCommitResponseBodyResultSignature) String() string {
	return tea.Prettify(s)
}

func (s GetFileLastCommitResponseBodyResultSignature) GoString() string {
	return s.String()
}

func (s *GetFileLastCommitResponseBodyResultSignature) SetGpgKeyId(v string) *GetFileLastCommitResponseBodyResultSignature {
	s.GpgKeyId = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResultSignature) SetVerificationStatus(v string) *GetFileLastCommitResponseBodyResultSignature {
	s.VerificationStatus = &v
	return s
}

type GetFileLastCommitResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFileLastCommitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileLastCommitResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileLastCommitResponse) GoString() string {
	return s.String()
}

func (s *GetFileLastCommitResponse) SetHeaders(v map[string]*string) *GetFileLastCommitResponse {
	s.Headers = v
	return s
}

func (s *GetFileLastCommitResponse) SetStatusCode(v int32) *GetFileLastCommitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileLastCommitResponse) SetBody(v *GetFileLastCommitResponseBody) *GetFileLastCommitResponse {
	s.Body = v
	return s
}

type GetFlowTagGroupResponseBody struct {
	ErrorCode    *string                                  `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                  `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	FlowTagGroup *GetFlowTagGroupResponseBodyFlowTagGroup `json:"flowTagGroup,omitempty" xml:"flowTagGroup,omitempty" type:"Struct"`
	RequestId    *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetFlowTagGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFlowTagGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowTagGroupResponseBody) SetErrorCode(v string) *GetFlowTagGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetFlowTagGroupResponseBody) SetErrorMessage(v string) *GetFlowTagGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetFlowTagGroupResponseBody) SetFlowTagGroup(v *GetFlowTagGroupResponseBodyFlowTagGroup) *GetFlowTagGroupResponseBody {
	s.FlowTagGroup = v
	return s
}

func (s *GetFlowTagGroupResponseBody) SetRequestId(v string) *GetFlowTagGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFlowTagGroupResponseBody) SetSuccess(v bool) *GetFlowTagGroupResponseBody {
	s.Success = &v
	return s
}

type GetFlowTagGroupResponseBodyFlowTagGroup struct {
	CreatorAccountId *string                                               `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	FlowTagList      []*GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList `json:"flowTagList,omitempty" xml:"flowTagList,omitempty" type:"Repeated"`
	Id               *int64                                                `json:"id,omitempty" xml:"id,omitempty"`
	ModiferAccountId *string                                               `json:"modiferAccountId,omitempty" xml:"modiferAccountId,omitempty"`
	Name             *string                                               `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetFlowTagGroupResponseBodyFlowTagGroup) String() string {
	return tea.Prettify(s)
}

func (s GetFlowTagGroupResponseBodyFlowTagGroup) GoString() string {
	return s.String()
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroup) SetCreatorAccountId(v string) *GetFlowTagGroupResponseBodyFlowTagGroup {
	s.CreatorAccountId = &v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroup) SetFlowTagList(v []*GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) *GetFlowTagGroupResponseBodyFlowTagGroup {
	s.FlowTagList = v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroup) SetId(v int64) *GetFlowTagGroupResponseBodyFlowTagGroup {
	s.Id = &v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroup) SetModiferAccountId(v string) *GetFlowTagGroupResponseBodyFlowTagGroup {
	s.ModiferAccountId = &v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroup) SetName(v string) *GetFlowTagGroupResponseBodyFlowTagGroup {
	s.Name = &v
	return s
}

type GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList struct {
	Color            *string `json:"color,omitempty" xml:"color,omitempty"`
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	Id               *int64  `json:"id,omitempty" xml:"id,omitempty"`
	ModiferAccountId *string `json:"modiferAccountId,omitempty" xml:"modiferAccountId,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) String() string {
	return tea.Prettify(s)
}

func (s GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) GoString() string {
	return s.String()
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) SetColor(v string) *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList {
	s.Color = &v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) SetCreatorAccountId(v string) *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList {
	s.CreatorAccountId = &v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) SetId(v int64) *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList {
	s.Id = &v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) SetModiferAccountId(v string) *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList {
	s.ModiferAccountId = &v
	return s
}

func (s *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList) SetName(v string) *GetFlowTagGroupResponseBodyFlowTagGroupFlowTagList {
	s.Name = &v
	return s
}

type GetFlowTagGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFlowTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFlowTagGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFlowTagGroupResponse) GoString() string {
	return s.String()
}

func (s *GetFlowTagGroupResponse) SetHeaders(v map[string]*string) *GetFlowTagGroupResponse {
	s.Headers = v
	return s
}

func (s *GetFlowTagGroupResponse) SetStatusCode(v int32) *GetFlowTagGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFlowTagGroupResponse) SetBody(v *GetFlowTagGroupResponseBody) *GetFlowTagGroupResponse {
	s.Body = v
	return s
}

type GetHostGroupResponseBody struct {
	ErrorCode    *string                            `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                            `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HostGroup    *GetHostGroupResponseBodyHostGroup `json:"hostGroup,omitempty" xml:"hostGroup,omitempty" type:"Struct"`
	RequestId    *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetHostGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetHostGroupResponseBody) SetErrorCode(v string) *GetHostGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetHostGroupResponseBody) SetErrorMessage(v string) *GetHostGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetHostGroupResponseBody) SetHostGroup(v *GetHostGroupResponseBodyHostGroup) *GetHostGroupResponseBody {
	s.HostGroup = v
	return s
}

func (s *GetHostGroupResponseBody) SetRequestId(v string) *GetHostGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHostGroupResponseBody) SetSuccess(v bool) *GetHostGroupResponseBody {
	s.Success = &v
	return s
}

type GetHostGroupResponseBodyHostGroup struct {
	AliyunRegion        *string                                       `json:"aliyunRegion,omitempty" xml:"aliyunRegion,omitempty"`
	CreateTime          *int64                                        `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorAccountId    *string                                       `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	Description         *string                                       `json:"description,omitempty" xml:"description,omitempty"`
	EcsLabelKey         *string                                       `json:"ecsLabelKey,omitempty" xml:"ecsLabelKey,omitempty"`
	EcsLabelValue       *string                                       `json:"ecsLabelValue,omitempty" xml:"ecsLabelValue,omitempty"`
	EcsType             *string                                       `json:"ecsType,omitempty" xml:"ecsType,omitempty"`
	HostInfos           []*GetHostGroupResponseBodyHostGroupHostInfos `json:"hostInfos,omitempty" xml:"hostInfos,omitempty" type:"Repeated"`
	HostNum             *int64                                        `json:"hostNum,omitempty" xml:"hostNum,omitempty"`
	Id                  *int64                                        `json:"id,omitempty" xml:"id,omitempty"`
	ModifierAccountId   *string                                       `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	Name                *string                                       `json:"name,omitempty" xml:"name,omitempty"`
	ServiceConnectionId *int64                                        `json:"serviceConnectionId,omitempty" xml:"serviceConnectionId,omitempty"`
	Type                *string                                       `json:"type,omitempty" xml:"type,omitempty"`
	UpateTIme           *int64                                        `json:"upateTIme,omitempty" xml:"upateTIme,omitempty"`
}

func (s GetHostGroupResponseBodyHostGroup) String() string {
	return tea.Prettify(s)
}

func (s GetHostGroupResponseBodyHostGroup) GoString() string {
	return s.String()
}

func (s *GetHostGroupResponseBodyHostGroup) SetAliyunRegion(v string) *GetHostGroupResponseBodyHostGroup {
	s.AliyunRegion = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetCreateTime(v int64) *GetHostGroupResponseBodyHostGroup {
	s.CreateTime = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetCreatorAccountId(v string) *GetHostGroupResponseBodyHostGroup {
	s.CreatorAccountId = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetDescription(v string) *GetHostGroupResponseBodyHostGroup {
	s.Description = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetEcsLabelKey(v string) *GetHostGroupResponseBodyHostGroup {
	s.EcsLabelKey = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetEcsLabelValue(v string) *GetHostGroupResponseBodyHostGroup {
	s.EcsLabelValue = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetEcsType(v string) *GetHostGroupResponseBodyHostGroup {
	s.EcsType = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetHostInfos(v []*GetHostGroupResponseBodyHostGroupHostInfos) *GetHostGroupResponseBodyHostGroup {
	s.HostInfos = v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetHostNum(v int64) *GetHostGroupResponseBodyHostGroup {
	s.HostNum = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetId(v int64) *GetHostGroupResponseBodyHostGroup {
	s.Id = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetModifierAccountId(v string) *GetHostGroupResponseBodyHostGroup {
	s.ModifierAccountId = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetName(v string) *GetHostGroupResponseBodyHostGroup {
	s.Name = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetServiceConnectionId(v int64) *GetHostGroupResponseBodyHostGroup {
	s.ServiceConnectionId = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetType(v string) *GetHostGroupResponseBodyHostGroup {
	s.Type = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetUpateTIme(v int64) *GetHostGroupResponseBodyHostGroup {
	s.UpateTIme = &v
	return s
}

type GetHostGroupResponseBodyHostGroupHostInfos struct {
	AliyunRegionId    *string `json:"aliyunRegionId,omitempty" xml:"aliyunRegionId,omitempty"`
	CreateTime        *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorAccountId  *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	InstanceName      *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	Ip                *string `json:"ip,omitempty" xml:"ip,omitempty"`
	MachineSn         *string `json:"machineSn,omitempty" xml:"machineSn,omitempty"`
	ModifierAccountId *string `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	ObjectType        *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	UpdateTime        *int64  `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetHostGroupResponseBodyHostGroupHostInfos) String() string {
	return tea.Prettify(s)
}

func (s GetHostGroupResponseBodyHostGroupHostInfos) GoString() string {
	return s.String()
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetAliyunRegionId(v string) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.AliyunRegionId = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetCreateTime(v int64) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.CreateTime = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetCreatorAccountId(v string) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.CreatorAccountId = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetInstanceName(v string) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.InstanceName = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetIp(v string) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.Ip = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetMachineSn(v string) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.MachineSn = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetModifierAccountId(v string) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.ModifierAccountId = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetObjectType(v string) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.ObjectType = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroupHostInfos) SetUpdateTime(v int64) *GetHostGroupResponseBodyHostGroupHostInfos {
	s.UpdateTime = &v
	return s
}

type GetHostGroupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHostGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHostGroupResponse) GoString() string {
	return s.String()
}

func (s *GetHostGroupResponse) SetHeaders(v map[string]*string) *GetHostGroupResponse {
	s.Headers = v
	return s
}

func (s *GetHostGroupResponse) SetStatusCode(v int32) *GetHostGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHostGroupResponse) SetBody(v *GetHostGroupResponseBody) *GetHostGroupResponse {
	s.Body = v
	return s
}

type GetOrganizationMemberResponseBody struct {
	ErrorCode    *string                                  `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                  `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Member       *GetOrganizationMemberResponseBodyMember `json:"member,omitempty" xml:"member,omitempty" type:"Struct"`
	RequestId    *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetOrganizationMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationMemberResponseBody) SetErrorCode(v string) *GetOrganizationMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOrganizationMemberResponseBody) SetErrorMessage(v string) *GetOrganizationMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetOrganizationMemberResponseBody) SetMember(v *GetOrganizationMemberResponseBodyMember) *GetOrganizationMemberResponseBody {
	s.Member = v
	return s
}

func (s *GetOrganizationMemberResponseBody) SetRequestId(v string) *GetOrganizationMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrganizationMemberResponseBody) SetSuccess(v bool) *GetOrganizationMemberResponseBody {
	s.Success = &v
	return s
}

type GetOrganizationMemberResponseBodyMember struct {
	AccountId              *string                                            `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Birthday               *int64                                             `json:"birthday,omitempty" xml:"birthday,omitempty"`
	DeptLists              []*string                                          `json:"deptLists,omitempty" xml:"deptLists,omitempty" type:"Repeated"`
	Email                  *string                                            `json:"email,omitempty" xml:"email,omitempty"`
	HiredDate              *int64                                             `json:"hiredDate,omitempty" xml:"hiredDate,omitempty"`
	Identities             *GetOrganizationMemberResponseBodyMemberIdentities `json:"identities,omitempty" xml:"identities,omitempty" type:"Struct"`
	JoinTime               *int64                                             `json:"joinTime,omitempty" xml:"joinTime,omitempty"`
	LastVisitTime          *int64                                             `json:"lastVisitTime,omitempty" xml:"lastVisitTime,omitempty"`
	Mobile                 *string                                            `json:"mobile,omitempty" xml:"mobile,omitempty"`
	OrganizationMemberName *string                                            `json:"organizationMemberName,omitempty" xml:"organizationMemberName,omitempty"`
	OrganizationRoleId     *string                                            `json:"organizationRoleId,omitempty" xml:"organizationRoleId,omitempty"`
	OrganizationRoleName   *string                                            `json:"organizationRoleName,omitempty" xml:"organizationRoleName,omitempty"`
	State                  *string                                            `json:"state,omitempty" xml:"state,omitempty"`
}

func (s GetOrganizationMemberResponseBodyMember) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationMemberResponseBodyMember) GoString() string {
	return s.String()
}

func (s *GetOrganizationMemberResponseBodyMember) SetAccountId(v string) *GetOrganizationMemberResponseBodyMember {
	s.AccountId = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetBirthday(v int64) *GetOrganizationMemberResponseBodyMember {
	s.Birthday = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetDeptLists(v []*string) *GetOrganizationMemberResponseBodyMember {
	s.DeptLists = v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetEmail(v string) *GetOrganizationMemberResponseBodyMember {
	s.Email = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetHiredDate(v int64) *GetOrganizationMemberResponseBodyMember {
	s.HiredDate = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetIdentities(v *GetOrganizationMemberResponseBodyMemberIdentities) *GetOrganizationMemberResponseBodyMember {
	s.Identities = v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetJoinTime(v int64) *GetOrganizationMemberResponseBodyMember {
	s.JoinTime = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetLastVisitTime(v int64) *GetOrganizationMemberResponseBodyMember {
	s.LastVisitTime = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetMobile(v string) *GetOrganizationMemberResponseBodyMember {
	s.Mobile = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetOrganizationMemberName(v string) *GetOrganizationMemberResponseBodyMember {
	s.OrganizationMemberName = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetOrganizationRoleId(v string) *GetOrganizationMemberResponseBodyMember {
	s.OrganizationRoleId = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetOrganizationRoleName(v string) *GetOrganizationMemberResponseBodyMember {
	s.OrganizationRoleName = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetState(v string) *GetOrganizationMemberResponseBodyMember {
	s.State = &v
	return s
}

type GetOrganizationMemberResponseBodyMemberIdentities struct {
	ExternUid *string `json:"externUid,omitempty" xml:"externUid,omitempty"`
	Provider  *string `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s GetOrganizationMemberResponseBodyMemberIdentities) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationMemberResponseBodyMemberIdentities) GoString() string {
	return s.String()
}

func (s *GetOrganizationMemberResponseBodyMemberIdentities) SetExternUid(v string) *GetOrganizationMemberResponseBodyMemberIdentities {
	s.ExternUid = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMemberIdentities) SetProvider(v string) *GetOrganizationMemberResponseBodyMemberIdentities {
	s.Provider = &v
	return s
}

type GetOrganizationMemberResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOrganizationMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrganizationMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationMemberResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationMemberResponse) SetHeaders(v map[string]*string) *GetOrganizationMemberResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationMemberResponse) SetStatusCode(v int32) *GetOrganizationMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrganizationMemberResponse) SetBody(v *GetOrganizationMemberResponseBody) *GetOrganizationMemberResponse {
	s.Body = v
	return s
}

type GetPipelineResponseBody struct {
	ErrorCode    *string                          `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                          `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Pipeline     *GetPipelineResponseBodyPipeline `json:"pipeline,omitempty" xml:"pipeline,omitempty" type:"Struct"`
	RequestId    *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPipelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBody) SetErrorCode(v string) *GetPipelineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineResponseBody) SetErrorMessage(v string) *GetPipelineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineResponseBody) SetPipeline(v *GetPipelineResponseBodyPipeline) *GetPipelineResponseBody {
	s.Pipeline = v
	return s
}

func (s *GetPipelineResponseBody) SetRequestId(v string) *GetPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineResponseBody) SetSuccess(v bool) *GetPipelineResponseBody {
	s.Success = &v
	return s
}

type GetPipelineResponseBodyPipeline struct {
	CreateTime        *int64                                         `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorAccountId  *string                                        `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	EnvId             *int32                                         `json:"envId,omitempty" xml:"envId,omitempty"`
	EnvName           *string                                        `json:"envName,omitempty" xml:"envName,omitempty"`
	GroupId           *int64                                         `json:"groupId,omitempty" xml:"groupId,omitempty"`
	ModifierAccountId *string                                        `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	Name              *string                                        `json:"name,omitempty" xml:"name,omitempty"`
	PipelineConfig    *GetPipelineResponseBodyPipelinePipelineConfig `json:"pipelineConfig,omitempty" xml:"pipelineConfig,omitempty" type:"Struct"`
	TagList           []*GetPipelineResponseBodyPipelineTagList      `json:"tagList,omitempty" xml:"tagList,omitempty" type:"Repeated"`
	UpdateTime        *int64                                         `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetPipelineResponseBodyPipeline) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineResponseBodyPipeline) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodyPipeline) SetCreateTime(v int64) *GetPipelineResponseBodyPipeline {
	s.CreateTime = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetCreatorAccountId(v string) *GetPipelineResponseBodyPipeline {
	s.CreatorAccountId = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetEnvId(v int32) *GetPipelineResponseBodyPipeline {
	s.EnvId = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetEnvName(v string) *GetPipelineResponseBodyPipeline {
	s.EnvName = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetGroupId(v int64) *GetPipelineResponseBodyPipeline {
	s.GroupId = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetModifierAccountId(v string) *GetPipelineResponseBodyPipeline {
	s.ModifierAccountId = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetName(v string) *GetPipelineResponseBodyPipeline {
	s.Name = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetPipelineConfig(v *GetPipelineResponseBodyPipelinePipelineConfig) *GetPipelineResponseBodyPipeline {
	s.PipelineConfig = v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetTagList(v []*GetPipelineResponseBodyPipelineTagList) *GetPipelineResponseBodyPipeline {
	s.TagList = v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetUpdateTime(v int64) *GetPipelineResponseBodyPipeline {
	s.UpdateTime = &v
	return s
}

type GetPipelineResponseBodyPipelinePipelineConfig struct {
	Flow     *string                                                 `json:"flow,omitempty" xml:"flow,omitempty"`
	Settings *string                                                 `json:"settings,omitempty" xml:"settings,omitempty"`
	Sources  []*GetPipelineResponseBodyPipelinePipelineConfigSources `json:"sources,omitempty" xml:"sources,omitempty" type:"Repeated"`
}

func (s GetPipelineResponseBodyPipelinePipelineConfig) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineResponseBodyPipelinePipelineConfig) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodyPipelinePipelineConfig) SetFlow(v string) *GetPipelineResponseBodyPipelinePipelineConfig {
	s.Flow = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfig) SetSettings(v string) *GetPipelineResponseBodyPipelinePipelineConfig {
	s.Settings = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfig) SetSources(v []*GetPipelineResponseBodyPipelinePipelineConfigSources) *GetPipelineResponseBodyPipelinePipelineConfig {
	s.Sources = v
	return s
}

type GetPipelineResponseBodyPipelinePipelineConfigSources struct {
	Data *GetPipelineResponseBodyPipelinePipelineConfigSourcesData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Sign *string                                                   `json:"sign,omitempty" xml:"sign,omitempty"`
	Type *string                                                   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPipelineResponseBodyPipelinePipelineConfigSources) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineResponseBodyPipelinePipelineConfigSources) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSources) SetData(v *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) *GetPipelineResponseBodyPipelinePipelineConfigSources {
	s.Data = v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSources) SetSign(v string) *GetPipelineResponseBodyPipelinePipelineConfigSources {
	s.Sign = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSources) SetType(v string) *GetPipelineResponseBodyPipelinePipelineConfigSources {
	s.Type = &v
	return s
}

type GetPipelineResponseBodyPipelinePipelineConfigSourcesData struct {
	Branch              *string   `json:"branch,omitempty" xml:"branch,omitempty"`
	CloneDepth          *int64    `json:"cloneDepth,omitempty" xml:"cloneDepth,omitempty"`
	CredentialId        *int64    `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	CredentialLabel     *string   `json:"credentialLabel,omitempty" xml:"credentialLabel,omitempty"`
	CredentialType      *string   `json:"credentialType,omitempty" xml:"credentialType,omitempty"`
	Events              []*string `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	IsBranchMode        *bool     `json:"isBranchMode,omitempty" xml:"isBranchMode,omitempty"`
	IsCloneDepth        *bool     `json:"isCloneDepth,omitempty" xml:"isCloneDepth,omitempty"`
	IsSubmodule         *bool     `json:"isSubmodule,omitempty" xml:"isSubmodule,omitempty"`
	IsTrigger           *bool     `json:"isTrigger,omitempty" xml:"isTrigger,omitempty"`
	Label               *string   `json:"label,omitempty" xml:"label,omitempty"`
	Namespace           *string   `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Repo                *string   `json:"repo,omitempty" xml:"repo,omitempty"`
	ServiceConnectionId *int64    `json:"serviceConnectionId,omitempty" xml:"serviceConnectionId,omitempty"`
	TriggerFilter       *string   `json:"triggerFilter,omitempty" xml:"triggerFilter,omitempty"`
	Webhook             *string   `json:"webhook,omitempty" xml:"webhook,omitempty"`
}

func (s GetPipelineResponseBodyPipelinePipelineConfigSourcesData) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineResponseBodyPipelinePipelineConfigSourcesData) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetBranch(v string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.Branch = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetCloneDepth(v int64) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.CloneDepth = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetCredentialId(v int64) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.CredentialId = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetCredentialLabel(v string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.CredentialLabel = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetCredentialType(v string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.CredentialType = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetEvents(v []*string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.Events = v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetIsBranchMode(v bool) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.IsBranchMode = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetIsCloneDepth(v bool) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.IsCloneDepth = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetIsSubmodule(v bool) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.IsSubmodule = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetIsTrigger(v bool) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.IsTrigger = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetLabel(v string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.Label = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetNamespace(v string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.Namespace = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetRepo(v string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.Repo = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetServiceConnectionId(v int64) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.ServiceConnectionId = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetTriggerFilter(v string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.TriggerFilter = &v
	return s
}

func (s *GetPipelineResponseBodyPipelinePipelineConfigSourcesData) SetWebhook(v string) *GetPipelineResponseBodyPipelinePipelineConfigSourcesData {
	s.Webhook = &v
	return s
}

type GetPipelineResponseBodyPipelineTagList struct {
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetPipelineResponseBodyPipelineTagList) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineResponseBodyPipelineTagList) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodyPipelineTagList) SetId(v int64) *GetPipelineResponseBodyPipelineTagList {
	s.Id = &v
	return s
}

func (s *GetPipelineResponseBodyPipelineTagList) SetName(v string) *GetPipelineResponseBodyPipelineTagList {
	s.Name = &v
	return s
}

type GetPipelineResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineResponse) SetHeaders(v map[string]*string) *GetPipelineResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineResponse) SetStatusCode(v int32) *GetPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineResponse) SetBody(v *GetPipelineResponseBody) *GetPipelineResponse {
	s.Body = v
	return s
}

type GetPipelineArtifactUrlRequest struct {
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
}

func (s GetPipelineArtifactUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineArtifactUrlRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineArtifactUrlRequest) SetFileName(v string) *GetPipelineArtifactUrlRequest {
	s.FileName = &v
	return s
}

func (s *GetPipelineArtifactUrlRequest) SetFilePath(v string) *GetPipelineArtifactUrlRequest {
	s.FilePath = &v
	return s
}

type GetPipelineArtifactUrlResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	FileUrl      *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPipelineArtifactUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineArtifactUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineArtifactUrlResponseBody) SetErrorCode(v string) *GetPipelineArtifactUrlResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineArtifactUrlResponseBody) SetErrorMessage(v string) *GetPipelineArtifactUrlResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineArtifactUrlResponseBody) SetFileUrl(v string) *GetPipelineArtifactUrlResponseBody {
	s.FileUrl = &v
	return s
}

func (s *GetPipelineArtifactUrlResponseBody) SetRequestId(v string) *GetPipelineArtifactUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineArtifactUrlResponseBody) SetSuccess(v bool) *GetPipelineArtifactUrlResponseBody {
	s.Success = &v
	return s
}

type GetPipelineArtifactUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPipelineArtifactUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPipelineArtifactUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineArtifactUrlResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineArtifactUrlResponse) SetHeaders(v map[string]*string) *GetPipelineArtifactUrlResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineArtifactUrlResponse) SetStatusCode(v int32) *GetPipelineArtifactUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineArtifactUrlResponse) SetBody(v *GetPipelineArtifactUrlResponseBody) *GetPipelineArtifactUrlResponse {
	s.Body = v
	return s
}

type GetPipelineEmasArtifactUrlRequest struct {
	ServiceConnectionId *int64 `json:"serviceConnectionId,omitempty" xml:"serviceConnectionId,omitempty"`
}

func (s GetPipelineEmasArtifactUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineEmasArtifactUrlRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineEmasArtifactUrlRequest) SetServiceConnectionId(v int64) *GetPipelineEmasArtifactUrlRequest {
	s.ServiceConnectionId = &v
	return s
}

type GetPipelineEmasArtifactUrlResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	FileUrl      *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPipelineEmasArtifactUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineEmasArtifactUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineEmasArtifactUrlResponseBody) SetErrorCode(v string) *GetPipelineEmasArtifactUrlResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineEmasArtifactUrlResponseBody) SetErrorMessage(v string) *GetPipelineEmasArtifactUrlResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineEmasArtifactUrlResponseBody) SetFileUrl(v string) *GetPipelineEmasArtifactUrlResponseBody {
	s.FileUrl = &v
	return s
}

func (s *GetPipelineEmasArtifactUrlResponseBody) SetRequestId(v string) *GetPipelineEmasArtifactUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineEmasArtifactUrlResponseBody) SetSuccess(v bool) *GetPipelineEmasArtifactUrlResponseBody {
	s.Success = &v
	return s
}

type GetPipelineEmasArtifactUrlResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPipelineEmasArtifactUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPipelineEmasArtifactUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineEmasArtifactUrlResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineEmasArtifactUrlResponse) SetHeaders(v map[string]*string) *GetPipelineEmasArtifactUrlResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineEmasArtifactUrlResponse) SetStatusCode(v int32) *GetPipelineEmasArtifactUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineEmasArtifactUrlResponse) SetBody(v *GetPipelineEmasArtifactUrlResponseBody) *GetPipelineEmasArtifactUrlResponse {
	s.Body = v
	return s
}

type GetPipelineGroupResponseBody struct {
	ErrorCode     *string                                    `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage  *string                                    `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	PipelineGroup *GetPipelineGroupResponseBodyPipelineGroup `json:"pipelineGroup,omitempty" xml:"pipelineGroup,omitempty" type:"Struct"`
	RequestId     *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success       *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPipelineGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineGroupResponseBody) SetErrorCode(v string) *GetPipelineGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineGroupResponseBody) SetErrorMessage(v string) *GetPipelineGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineGroupResponseBody) SetPipelineGroup(v *GetPipelineGroupResponseBodyPipelineGroup) *GetPipelineGroupResponseBody {
	s.PipelineGroup = v
	return s
}

func (s *GetPipelineGroupResponseBody) SetRequestId(v string) *GetPipelineGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineGroupResponseBody) SetSuccess(v bool) *GetPipelineGroupResponseBody {
	s.Success = &v
	return s
}

type GetPipelineGroupResponseBodyPipelineGroup struct {
	CreateTime *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Id         *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetPipelineGroupResponseBodyPipelineGroup) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineGroupResponseBodyPipelineGroup) GoString() string {
	return s.String()
}

func (s *GetPipelineGroupResponseBodyPipelineGroup) SetCreateTime(v int64) *GetPipelineGroupResponseBodyPipelineGroup {
	s.CreateTime = &v
	return s
}

func (s *GetPipelineGroupResponseBodyPipelineGroup) SetId(v int64) *GetPipelineGroupResponseBodyPipelineGroup {
	s.Id = &v
	return s
}

func (s *GetPipelineGroupResponseBodyPipelineGroup) SetName(v string) *GetPipelineGroupResponseBodyPipelineGroup {
	s.Name = &v
	return s
}

type GetPipelineGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPipelineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPipelineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineGroupResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineGroupResponse) SetHeaders(v map[string]*string) *GetPipelineGroupResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineGroupResponse) SetStatusCode(v int32) *GetPipelineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineGroupResponse) SetBody(v *GetPipelineGroupResponseBody) *GetPipelineGroupResponse {
	s.Body = v
	return s
}

type GetPipelineRunResponseBody struct {
	ErrorCode    *string                                `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	PipelineRun  *GetPipelineRunResponseBodyPipelineRun `json:"pipelineRun,omitempty" xml:"pipelineRun,omitempty" type:"Struct"`
	RequestId    *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPipelineRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBody) SetErrorCode(v string) *GetPipelineRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetErrorMessage(v string) *GetPipelineRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetPipelineRun(v *GetPipelineRunResponseBodyPipelineRun) *GetPipelineRunResponseBody {
	s.PipelineRun = v
	return s
}

func (s *GetPipelineRunResponseBody) SetRequestId(v string) *GetPipelineRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetSuccess(v bool) *GetPipelineRunResponseBody {
	s.Success = &v
	return s
}

type GetPipelineRunResponseBodyPipelineRun struct {
	CreateTime        *int64                                          `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorAccountId  *string                                         `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	ModifierAccountId *string                                         `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	PipelineId        *int64                                          `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	PipelineRunId     *int64                                          `json:"pipelineRunId,omitempty" xml:"pipelineRunId,omitempty"`
	Sources           []*GetPipelineRunResponseBodyPipelineRunSources `json:"sources,omitempty" xml:"sources,omitempty" type:"Repeated"`
	StageGroup        [][]*string                                     `json:"stageGroup,omitempty" xml:"stageGroup,omitempty" type:"Repeated"`
	Stages            []*GetPipelineRunResponseBodyPipelineRunStages  `json:"stages,omitempty" xml:"stages,omitempty" type:"Repeated"`
	Status            *string                                         `json:"status,omitempty" xml:"status,omitempty"`
	TriggerMode       *int32                                          `json:"triggerMode,omitempty" xml:"triggerMode,omitempty"`
	UpdateTime        *int64                                          `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetPipelineRunResponseBodyPipelineRun) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunResponseBodyPipelineRun) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetCreateTime(v int64) *GetPipelineRunResponseBodyPipelineRun {
	s.CreateTime = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetCreatorAccountId(v string) *GetPipelineRunResponseBodyPipelineRun {
	s.CreatorAccountId = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetModifierAccountId(v string) *GetPipelineRunResponseBodyPipelineRun {
	s.ModifierAccountId = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetPipelineId(v int64) *GetPipelineRunResponseBodyPipelineRun {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetPipelineRunId(v int64) *GetPipelineRunResponseBodyPipelineRun {
	s.PipelineRunId = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetSources(v []*GetPipelineRunResponseBodyPipelineRunSources) *GetPipelineRunResponseBodyPipelineRun {
	s.Sources = v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetStageGroup(v [][]*string) *GetPipelineRunResponseBodyPipelineRun {
	s.StageGroup = v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetStages(v []*GetPipelineRunResponseBodyPipelineRunStages) *GetPipelineRunResponseBodyPipelineRun {
	s.Stages = v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetStatus(v string) *GetPipelineRunResponseBodyPipelineRun {
	s.Status = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetTriggerMode(v int32) *GetPipelineRunResponseBodyPipelineRun {
	s.TriggerMode = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetUpdateTime(v int64) *GetPipelineRunResponseBodyPipelineRun {
	s.UpdateTime = &v
	return s
}

type GetPipelineRunResponseBodyPipelineRunSources struct {
	Data *GetPipelineRunResponseBodyPipelineRunSourcesData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Sign *string                                           `json:"sign,omitempty" xml:"sign,omitempty"`
	Type *string                                           `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPipelineRunResponseBodyPipelineRunSources) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunResponseBodyPipelineRunSources) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBodyPipelineRunSources) SetData(v *GetPipelineRunResponseBodyPipelineRunSourcesData) *GetPipelineRunResponseBodyPipelineRunSources {
	s.Data = v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunSources) SetSign(v string) *GetPipelineRunResponseBodyPipelineRunSources {
	s.Sign = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunSources) SetType(v string) *GetPipelineRunResponseBodyPipelineRunSources {
	s.Type = &v
	return s
}

type GetPipelineRunResponseBodyPipelineRunSourcesData struct {
	Branch  *string `json:"branch,omitempty" xml:"branch,omitempty"`
	Commint *string `json:"commint,omitempty" xml:"commint,omitempty"`
	Repo    *string `json:"repo,omitempty" xml:"repo,omitempty"`
}

func (s GetPipelineRunResponseBodyPipelineRunSourcesData) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunResponseBodyPipelineRunSourcesData) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBodyPipelineRunSourcesData) SetBranch(v string) *GetPipelineRunResponseBodyPipelineRunSourcesData {
	s.Branch = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunSourcesData) SetCommint(v string) *GetPipelineRunResponseBodyPipelineRunSourcesData {
	s.Commint = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunSourcesData) SetRepo(v string) *GetPipelineRunResponseBodyPipelineRunSourcesData {
	s.Repo = &v
	return s
}

type GetPipelineRunResponseBodyPipelineRunStages struct {
	Name      *string                                               `json:"name,omitempty" xml:"name,omitempty"`
	StageInfo *GetPipelineRunResponseBodyPipelineRunStagesStageInfo `json:"stageInfo,omitempty" xml:"stageInfo,omitempty" type:"Struct"`
}

func (s GetPipelineRunResponseBodyPipelineRunStages) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunResponseBodyPipelineRunStages) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBodyPipelineRunStages) SetName(v string) *GetPipelineRunResponseBodyPipelineRunStages {
	s.Name = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStages) SetStageInfo(v *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) *GetPipelineRunResponseBodyPipelineRunStages {
	s.StageInfo = v
	return s
}

type GetPipelineRunResponseBodyPipelineRunStagesStageInfo struct {
	EndTime   *int64                                                      `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Jobs      []*GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	Name      *string                                                     `json:"name,omitempty" xml:"name,omitempty"`
	StartTime *int64                                                      `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status    *string                                                     `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetPipelineRunResponseBodyPipelineRunStagesStageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunResponseBodyPipelineRunStagesStageInfo) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) SetEndTime(v int64) *GetPipelineRunResponseBodyPipelineRunStagesStageInfo {
	s.EndTime = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) SetJobs(v []*GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) *GetPipelineRunResponseBodyPipelineRunStagesStageInfo {
	s.Jobs = v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) SetName(v string) *GetPipelineRunResponseBodyPipelineRunStagesStageInfo {
	s.Name = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) SetStartTime(v int64) *GetPipelineRunResponseBodyPipelineRunStagesStageInfo {
	s.StartTime = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) SetStatus(v string) *GetPipelineRunResponseBodyPipelineRunStagesStageInfo {
	s.Status = &v
	return s
}

type GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs struct {
	Actions   []*GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	EndTime   *int64                                                             `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Id        *int64                                                             `json:"id,omitempty" xml:"id,omitempty"`
	Name      *string                                                            `json:"name,omitempty" xml:"name,omitempty"`
	Params    *string                                                            `json:"params,omitempty" xml:"params,omitempty"`
	StartTime *int64                                                             `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status    *string                                                            `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetActions(v []*GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.Actions = v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetEndTime(v int64) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.EndTime = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetId(v int64) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.Id = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetName(v string) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.Name = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetParams(v string) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.Params = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetStartTime(v int64) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.StartTime = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetStatus(v string) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.Status = &v
	return s
}

type GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions struct {
	Disable *bool                  `json:"disable,omitempty" xml:"disable,omitempty"`
	Params  map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	Type    *string                `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) SetDisable(v bool) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions {
	s.Disable = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) SetParams(v map[string]interface{}) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions {
	s.Params = v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) SetType(v string) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions {
	s.Type = &v
	return s
}

type GetPipelineRunResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPipelineRunResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponse) SetHeaders(v map[string]*string) *GetPipelineRunResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineRunResponse) SetStatusCode(v int32) *GetPipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineRunResponse) SetBody(v *GetPipelineRunResponseBody) *GetPipelineRunResponse {
	s.Body = v
	return s
}

type GetPipelineScanReportUrlRequest struct {
	ReportPath *string `json:"reportPath,omitempty" xml:"reportPath,omitempty"`
}

func (s GetPipelineScanReportUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineScanReportUrlRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineScanReportUrlRequest) SetReportPath(v string) *GetPipelineScanReportUrlRequest {
	s.ReportPath = &v
	return s
}

type GetPipelineScanReportUrlResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	ReportUrl    *string `json:"reportUrl,omitempty" xml:"reportUrl,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPipelineScanReportUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineScanReportUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineScanReportUrlResponseBody) SetErrorCode(v string) *GetPipelineScanReportUrlResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineScanReportUrlResponseBody) SetErrorMessage(v string) *GetPipelineScanReportUrlResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineScanReportUrlResponseBody) SetReportUrl(v string) *GetPipelineScanReportUrlResponseBody {
	s.ReportUrl = &v
	return s
}

func (s *GetPipelineScanReportUrlResponseBody) SetRequestId(v string) *GetPipelineScanReportUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineScanReportUrlResponseBody) SetSuccess(v bool) *GetPipelineScanReportUrlResponseBody {
	s.Success = &v
	return s
}

type GetPipelineScanReportUrlResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPipelineScanReportUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPipelineScanReportUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineScanReportUrlResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineScanReportUrlResponse) SetHeaders(v map[string]*string) *GetPipelineScanReportUrlResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineScanReportUrlResponse) SetStatusCode(v int32) *GetPipelineScanReportUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineScanReportUrlResponse) SetBody(v *GetPipelineScanReportUrlResponseBody) *GetPipelineScanReportUrlResponse {
	s.Body = v
	return s
}

type GetProjectInfoResponseBody struct {
	ErrorCode    *string                            `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                            `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Project      *GetProjectInfoResponseBodyProject `json:"project,omitempty" xml:"project,omitempty" type:"Struct"`
	RequestId    *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetProjectInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectInfoResponseBody) SetErrorCode(v string) *GetProjectInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProjectInfoResponseBody) SetErrorMessage(v string) *GetProjectInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetProjectInfoResponseBody) SetProject(v *GetProjectInfoResponseBodyProject) *GetProjectInfoResponseBody {
	s.Project = v
	return s
}

func (s *GetProjectInfoResponseBody) SetRequestId(v string) *GetProjectInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectInfoResponseBody) SetSuccess(v bool) *GetProjectInfoResponseBody {
	s.Success = &v
	return s
}

type GetProjectInfoResponseBodyProject struct {
	Category               *string `json:"category,omitempty" xml:"category,omitempty"`
	CategoryIdentifier     *string `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	Creator                *string `json:"creator,omitempty" xml:"creator,omitempty"`
	CustomCode             *string `json:"customCode,omitempty" xml:"customCode,omitempty"`
	Description            *string `json:"description,omitempty" xml:"description,omitempty"`
	GmtCreate              *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified            *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Icon                   *string `json:"icon,omitempty" xml:"icon,omitempty"`
	IconBig                *string `json:"iconBig,omitempty" xml:"iconBig,omitempty"`
	IconGroup              *string `json:"iconGroup,omitempty" xml:"iconGroup,omitempty"`
	IconSmall              *string `json:"iconSmall,omitempty" xml:"iconSmall,omitempty"`
	Id                     *string `json:"id,omitempty" xml:"id,omitempty"`
	Identifier             *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	IdentifierPath         *string `json:"identifierPath,omitempty" xml:"identifierPath,omitempty"`
	LogicalStatus          *string `json:"logicalStatus,omitempty" xml:"logicalStatus,omitempty"`
	Modifier               *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Name                   *string `json:"name,omitempty" xml:"name,omitempty"`
	OrganizationIdentifier *string `json:"organizationIdentifier,omitempty" xml:"organizationIdentifier,omitempty"`
	ParentIdentifier       *string `json:"parentIdentifier,omitempty" xml:"parentIdentifier,omitempty"`
	Scope                  *string `json:"scope,omitempty" xml:"scope,omitempty"`
	StatusIdentifier       *string `json:"statusIdentifier,omitempty" xml:"statusIdentifier,omitempty"`
	StatusStageIdentifier  *string `json:"statusStageIdentifier,omitempty" xml:"statusStageIdentifier,omitempty"`
	SubType                *string `json:"subType,omitempty" xml:"subType,omitempty"`
	TypeIdentifier         *string `json:"typeIdentifier,omitempty" xml:"typeIdentifier,omitempty"`
}

func (s GetProjectInfoResponseBodyProject) String() string {
	return tea.Prettify(s)
}

func (s GetProjectInfoResponseBodyProject) GoString() string {
	return s.String()
}

func (s *GetProjectInfoResponseBodyProject) SetCategory(v string) *GetProjectInfoResponseBodyProject {
	s.Category = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetCategoryIdentifier(v string) *GetProjectInfoResponseBodyProject {
	s.CategoryIdentifier = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetCreator(v string) *GetProjectInfoResponseBodyProject {
	s.Creator = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetCustomCode(v string) *GetProjectInfoResponseBodyProject {
	s.CustomCode = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetDescription(v string) *GetProjectInfoResponseBodyProject {
	s.Description = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetGmtCreate(v int64) *GetProjectInfoResponseBodyProject {
	s.GmtCreate = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetGmtModified(v int64) *GetProjectInfoResponseBodyProject {
	s.GmtModified = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetIcon(v string) *GetProjectInfoResponseBodyProject {
	s.Icon = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetIconBig(v string) *GetProjectInfoResponseBodyProject {
	s.IconBig = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetIconGroup(v string) *GetProjectInfoResponseBodyProject {
	s.IconGroup = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetIconSmall(v string) *GetProjectInfoResponseBodyProject {
	s.IconSmall = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetId(v string) *GetProjectInfoResponseBodyProject {
	s.Id = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetIdentifier(v string) *GetProjectInfoResponseBodyProject {
	s.Identifier = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetIdentifierPath(v string) *GetProjectInfoResponseBodyProject {
	s.IdentifierPath = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetLogicalStatus(v string) *GetProjectInfoResponseBodyProject {
	s.LogicalStatus = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetModifier(v string) *GetProjectInfoResponseBodyProject {
	s.Modifier = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetName(v string) *GetProjectInfoResponseBodyProject {
	s.Name = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetOrganizationIdentifier(v string) *GetProjectInfoResponseBodyProject {
	s.OrganizationIdentifier = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetParentIdentifier(v string) *GetProjectInfoResponseBodyProject {
	s.ParentIdentifier = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetScope(v string) *GetProjectInfoResponseBodyProject {
	s.Scope = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetStatusIdentifier(v string) *GetProjectInfoResponseBodyProject {
	s.StatusIdentifier = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetStatusStageIdentifier(v string) *GetProjectInfoResponseBodyProject {
	s.StatusStageIdentifier = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetSubType(v string) *GetProjectInfoResponseBodyProject {
	s.SubType = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetTypeIdentifier(v string) *GetProjectInfoResponseBodyProject {
	s.TypeIdentifier = &v
	return s
}

type GetProjectInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProjectInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectInfoResponse) GoString() string {
	return s.String()
}

func (s *GetProjectInfoResponse) SetHeaders(v map[string]*string) *GetProjectInfoResponse {
	s.Headers = v
	return s
}

func (s *GetProjectInfoResponse) SetStatusCode(v int32) *GetProjectInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectInfoResponse) SetBody(v *GetProjectInfoResponseBody) *GetProjectInfoResponse {
	s.Body = v
	return s
}

type GetProjectMemberRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	RepositoryId   *int64  `json:"repositoryId,omitempty" xml:"repositoryId,omitempty"`
	UserAliyunPk   *string `json:"userAliyunPk,omitempty" xml:"userAliyunPk,omitempty"`
}

func (s GetProjectMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *GetProjectMemberRequest) SetAccessToken(v string) *GetProjectMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *GetProjectMemberRequest) SetOrganizationId(v string) *GetProjectMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetProjectMemberRequest) SetRepositoryId(v int64) *GetProjectMemberRequest {
	s.RepositoryId = &v
	return s
}

func (s *GetProjectMemberRequest) SetUserAliyunPk(v string) *GetProjectMemberRequest {
	s.UserAliyunPk = &v
	return s
}

type GetProjectMemberResponseBody struct {
	ErrorCode    *string                             `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                             `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       *GetProjectMemberResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success      *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetProjectMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponseBody) SetErrorCode(v string) *GetProjectMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProjectMemberResponseBody) SetErrorMessage(v string) *GetProjectMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetProjectMemberResponseBody) SetRequestId(v string) *GetProjectMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectMemberResponseBody) SetResult(v *GetProjectMemberResponseBodyResult) *GetProjectMemberResponseBody {
	s.Result = v
	return s
}

func (s *GetProjectMemberResponseBody) SetSuccess(v bool) *GetProjectMemberResponseBody {
	s.Success = &v
	return s
}

type GetProjectMemberResponseBodyResult struct {
	AccessLevel  *int32  `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	AvatarUrl    *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	ExternUserId *string `json:"externUserId,omitempty" xml:"externUserId,omitempty"`
	Id           *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetProjectMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponseBodyResult) SetAccessLevel(v int32) *GetProjectMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *GetProjectMemberResponseBodyResult) SetAvatarUrl(v string) *GetProjectMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *GetProjectMemberResponseBodyResult) SetExternUserId(v string) *GetProjectMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *GetProjectMemberResponseBodyResult) SetId(v int64) *GetProjectMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetProjectMemberResponseBodyResult) SetName(v string) *GetProjectMemberResponseBodyResult {
	s.Name = &v
	return s
}

type GetProjectMemberResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemberResponse) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponse) SetHeaders(v map[string]*string) *GetProjectMemberResponse {
	s.Headers = v
	return s
}

func (s *GetProjectMemberResponse) SetStatusCode(v int32) *GetProjectMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectMemberResponse) SetBody(v *GetProjectMemberResponseBody) *GetProjectMemberResponse {
	s.Body = v
	return s
}

type GetRepositoryRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	Identity       *string `json:"identity,omitempty" xml:"identity,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s GetRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryRequest) GoString() string {
	return s.String()
}

func (s *GetRepositoryRequest) SetAccessToken(v string) *GetRepositoryRequest {
	s.AccessToken = &v
	return s
}

func (s *GetRepositoryRequest) SetIdentity(v string) *GetRepositoryRequest {
	s.Identity = &v
	return s
}

func (s *GetRepositoryRequest) SetOrganizationId(v string) *GetRepositoryRequest {
	s.OrganizationId = &v
	return s
}

type GetRepositoryResponseBody struct {
	ErrorCode    *string                              `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                              `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Repository   *GetRepositoryResponseBodyRepository `json:"repository,omitempty" xml:"repository,omitempty" type:"Struct"`
	RequestId    *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepositoryResponseBody) SetErrorCode(v string) *GetRepositoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRepositoryResponseBody) SetErrorMessage(v string) *GetRepositoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRepositoryResponseBody) SetRepository(v *GetRepositoryResponseBodyRepository) *GetRepositoryResponseBody {
	s.Repository = v
	return s
}

func (s *GetRepositoryResponseBody) SetRequestId(v string) *GetRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepositoryResponseBody) SetSuccess(v bool) *GetRepositoryResponseBody {
	s.Success = &v
	return s
}

type GetRepositoryResponseBodyRepository struct {
	Archive             *bool                                         `json:"archive,omitempty" xml:"archive,omitempty"`
	AvatarUrl           *string                                       `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	CreatedAt           *string                                       `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatorId           *int64                                        `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	DefaultBranch       *string                                       `json:"defaultBranch,omitempty" xml:"defaultBranch,omitempty"`
	DemoProjectStatus   *bool                                         `json:"demoProjectStatus,omitempty" xml:"demoProjectStatus,omitempty"`
	Description         *string                                       `json:"description,omitempty" xml:"description,omitempty"`
	HttpUrlToRepository *string                                       `json:"httpUrlToRepository,omitempty" xml:"httpUrlToRepository,omitempty"`
	Id                  *int64                                        `json:"id,omitempty" xml:"id,omitempty"`
	LastActivityAt      *string                                       `json:"lastActivityAt,omitempty" xml:"lastActivityAt,omitempty"`
	Name                *string                                       `json:"name,omitempty" xml:"name,omitempty"`
	NameWithNamespace   *string                                       `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	Namespace           *GetRepositoryResponseBodyRepositoryNamespace `json:"namespace,omitempty" xml:"namespace,omitempty" type:"Struct"`
	Path                *string                                       `json:"path,omitempty" xml:"path,omitempty"`
	PathWithNamespace   *string                                       `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	SshUrlToRepository  *string                                       `json:"sshUrlToRepository,omitempty" xml:"sshUrlToRepository,omitempty"`
	VisibilityLevel     *int32                                        `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	WebUrl              *string                                       `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s GetRepositoryResponseBodyRepository) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryResponseBodyRepository) GoString() string {
	return s.String()
}

func (s *GetRepositoryResponseBodyRepository) SetArchive(v bool) *GetRepositoryResponseBodyRepository {
	s.Archive = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetAvatarUrl(v string) *GetRepositoryResponseBodyRepository {
	s.AvatarUrl = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetCreatedAt(v string) *GetRepositoryResponseBodyRepository {
	s.CreatedAt = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetCreatorId(v int64) *GetRepositoryResponseBodyRepository {
	s.CreatorId = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetDefaultBranch(v string) *GetRepositoryResponseBodyRepository {
	s.DefaultBranch = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetDemoProjectStatus(v bool) *GetRepositoryResponseBodyRepository {
	s.DemoProjectStatus = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetDescription(v string) *GetRepositoryResponseBodyRepository {
	s.Description = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetHttpUrlToRepository(v string) *GetRepositoryResponseBodyRepository {
	s.HttpUrlToRepository = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetId(v int64) *GetRepositoryResponseBodyRepository {
	s.Id = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetLastActivityAt(v string) *GetRepositoryResponseBodyRepository {
	s.LastActivityAt = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetName(v string) *GetRepositoryResponseBodyRepository {
	s.Name = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetNameWithNamespace(v string) *GetRepositoryResponseBodyRepository {
	s.NameWithNamespace = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetNamespace(v *GetRepositoryResponseBodyRepositoryNamespace) *GetRepositoryResponseBodyRepository {
	s.Namespace = v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetPath(v string) *GetRepositoryResponseBodyRepository {
	s.Path = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetPathWithNamespace(v string) *GetRepositoryResponseBodyRepository {
	s.PathWithNamespace = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetSshUrlToRepository(v string) *GetRepositoryResponseBodyRepository {
	s.SshUrlToRepository = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetVisibilityLevel(v int32) *GetRepositoryResponseBodyRepository {
	s.VisibilityLevel = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetWebUrl(v string) *GetRepositoryResponseBodyRepository {
	s.WebUrl = &v
	return s
}

type GetRepositoryResponseBodyRepositoryNamespace struct {
	Avatar          *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	CreatedAt       *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	Id              *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	OwnerId         *int64  `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	Path            *string `json:"path,omitempty" xml:"path,omitempty"`
	UpdatedAt       *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	VisibilityLevel *int32  `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
}

func (s GetRepositoryResponseBodyRepositoryNamespace) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryResponseBodyRepositoryNamespace) GoString() string {
	return s.String()
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetAvatar(v string) *GetRepositoryResponseBodyRepositoryNamespace {
	s.Avatar = &v
	return s
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetCreatedAt(v string) *GetRepositoryResponseBodyRepositoryNamespace {
	s.CreatedAt = &v
	return s
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetDescription(v string) *GetRepositoryResponseBodyRepositoryNamespace {
	s.Description = &v
	return s
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetId(v int64) *GetRepositoryResponseBodyRepositoryNamespace {
	s.Id = &v
	return s
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetName(v string) *GetRepositoryResponseBodyRepositoryNamespace {
	s.Name = &v
	return s
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetOwnerId(v int64) *GetRepositoryResponseBodyRepositoryNamespace {
	s.OwnerId = &v
	return s
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetPath(v string) *GetRepositoryResponseBodyRepositoryNamespace {
	s.Path = &v
	return s
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetUpdatedAt(v string) *GetRepositoryResponseBodyRepositoryNamespace {
	s.UpdatedAt = &v
	return s
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetVisibilityLevel(v int32) *GetRepositoryResponseBodyRepositoryNamespace {
	s.VisibilityLevel = &v
	return s
}

type GetRepositoryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryResponse) GoString() string {
	return s.String()
}

func (s *GetRepositoryResponse) SetHeaders(v map[string]*string) *GetRepositoryResponse {
	s.Headers = v
	return s
}

func (s *GetRepositoryResponse) SetStatusCode(v int32) *GetRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepositoryResponse) SetBody(v *GetRepositoryResponseBody) *GetRepositoryResponse {
	s.Body = v
	return s
}

type GetSprintInfoResponseBody struct {
	ErrorCode    *string                          `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                          `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Sprint       *GetSprintInfoResponseBodySprint `json:"sprint,omitempty" xml:"sprint,omitempty" type:"Struct"`
	Success      *bool                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSprintInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSprintInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSprintInfoResponseBody) SetErrorCode(v string) *GetSprintInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSprintInfoResponseBody) SetErrorMessage(v string) *GetSprintInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSprintInfoResponseBody) SetRequestId(v string) *GetSprintInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSprintInfoResponseBody) SetSprint(v *GetSprintInfoResponseBodySprint) *GetSprintInfoResponseBody {
	s.Sprint = v
	return s
}

func (s *GetSprintInfoResponseBody) SetSuccess(v bool) *GetSprintInfoResponseBody {
	s.Success = &v
	return s
}

type GetSprintInfoResponseBodySprint struct {
	Creator         *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	EndDate         *int64  `json:"endDate,omitempty" xml:"endDate,omitempty"`
	GmtCreate       *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified     *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Identifier      *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Modifier        *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	Scope           *string `json:"scope,omitempty" xml:"scope,omitempty"`
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	StartDate       *int64  `json:"startDate,omitempty" xml:"startDate,omitempty"`
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetSprintInfoResponseBodySprint) String() string {
	return tea.Prettify(s)
}

func (s GetSprintInfoResponseBodySprint) GoString() string {
	return s.String()
}

func (s *GetSprintInfoResponseBodySprint) SetCreator(v string) *GetSprintInfoResponseBodySprint {
	s.Creator = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetDescription(v string) *GetSprintInfoResponseBodySprint {
	s.Description = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetEndDate(v int64) *GetSprintInfoResponseBodySprint {
	s.EndDate = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetGmtCreate(v int64) *GetSprintInfoResponseBodySprint {
	s.GmtCreate = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetGmtModified(v int64) *GetSprintInfoResponseBodySprint {
	s.GmtModified = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetIdentifier(v string) *GetSprintInfoResponseBodySprint {
	s.Identifier = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetModifier(v string) *GetSprintInfoResponseBodySprint {
	s.Modifier = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetName(v string) *GetSprintInfoResponseBodySprint {
	s.Name = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetScope(v string) *GetSprintInfoResponseBodySprint {
	s.Scope = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetSpaceIdentifier(v string) *GetSprintInfoResponseBodySprint {
	s.SpaceIdentifier = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetStartDate(v int64) *GetSprintInfoResponseBodySprint {
	s.StartDate = &v
	return s
}

func (s *GetSprintInfoResponseBodySprint) SetStatus(v string) *GetSprintInfoResponseBodySprint {
	s.Status = &v
	return s
}

type GetSprintInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSprintInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSprintInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSprintInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSprintInfoResponse) SetHeaders(v map[string]*string) *GetSprintInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSprintInfoResponse) SetStatusCode(v int32) *GetSprintInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSprintInfoResponse) SetBody(v *GetSprintInfoResponseBody) *GetSprintInfoResponse {
	s.Body = v
	return s
}

type GetVMDeployOrderResponseBody struct {
	DeployOrder  *GetVMDeployOrderResponseBodyDeployOrder `json:"deployOrder,omitempty" xml:"deployOrder,omitempty" type:"Struct"`
	ErrorCode    *string                                  `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                  `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetVMDeployOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVMDeployOrderResponseBody) GoString() string {
	return s.String()
}

func (s *GetVMDeployOrderResponseBody) SetDeployOrder(v *GetVMDeployOrderResponseBodyDeployOrder) *GetVMDeployOrderResponseBody {
	s.DeployOrder = v
	return s
}

func (s *GetVMDeployOrderResponseBody) SetErrorCode(v string) *GetVMDeployOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetVMDeployOrderResponseBody) SetErrorMessage(v string) *GetVMDeployOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetVMDeployOrderResponseBody) SetRequestId(v string) *GetVMDeployOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVMDeployOrderResponseBody) SetSuccess(v bool) *GetVMDeployOrderResponseBody {
	s.Success = &v
	return s
}

type GetVMDeployOrderResponseBodyDeployOrder struct {
	Actions           []*GetVMDeployOrderResponseBodyDeployOrderActions         `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	CreateTime        *int64                                                    `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator           *string                                                   `json:"creator,omitempty" xml:"creator,omitempty"`
	CurrentBatch      *int32                                                    `json:"currentBatch,omitempty" xml:"currentBatch,omitempty"`
	DeployMachineInfo *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo `json:"deployMachineInfo,omitempty" xml:"deployMachineInfo,omitempty" type:"Struct"`
	DeployOrderId     *string                                                   `json:"deployOrderId,omitempty" xml:"deployOrderId,omitempty"`
	ExceptionCode     *string                                                   `json:"exceptionCode,omitempty" xml:"exceptionCode,omitempty"`
	Status            *string                                                   `json:"status,omitempty" xml:"status,omitempty"`
	TotalBatch        *int32                                                    `json:"totalBatch,omitempty" xml:"totalBatch,omitempty"`
	UpdateTime        *int64                                                    `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetVMDeployOrderResponseBodyDeployOrder) String() string {
	return tea.Prettify(s)
}

func (s GetVMDeployOrderResponseBodyDeployOrder) GoString() string {
	return s.String()
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetActions(v []*GetVMDeployOrderResponseBodyDeployOrderActions) *GetVMDeployOrderResponseBodyDeployOrder {
	s.Actions = v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetCreateTime(v int64) *GetVMDeployOrderResponseBodyDeployOrder {
	s.CreateTime = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetCreator(v string) *GetVMDeployOrderResponseBodyDeployOrder {
	s.Creator = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetCurrentBatch(v int32) *GetVMDeployOrderResponseBodyDeployOrder {
	s.CurrentBatch = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetDeployMachineInfo(v *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo) *GetVMDeployOrderResponseBodyDeployOrder {
	s.DeployMachineInfo = v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetDeployOrderId(v string) *GetVMDeployOrderResponseBodyDeployOrder {
	s.DeployOrderId = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetExceptionCode(v string) *GetVMDeployOrderResponseBodyDeployOrder {
	s.ExceptionCode = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetStatus(v string) *GetVMDeployOrderResponseBodyDeployOrder {
	s.Status = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetTotalBatch(v int32) *GetVMDeployOrderResponseBodyDeployOrder {
	s.TotalBatch = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrder) SetUpdateTime(v int64) *GetVMDeployOrderResponseBodyDeployOrder {
	s.UpdateTime = &v
	return s
}

type GetVMDeployOrderResponseBodyDeployOrderActions struct {
	Disable *bool       `json:"disable,omitempty" xml:"disable,omitempty"`
	Params  interface{} `json:"params,omitempty" xml:"params,omitempty"`
	Type    *string     `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetVMDeployOrderResponseBodyDeployOrderActions) String() string {
	return tea.Prettify(s)
}

func (s GetVMDeployOrderResponseBodyDeployOrderActions) GoString() string {
	return s.String()
}

func (s *GetVMDeployOrderResponseBodyDeployOrderActions) SetDisable(v bool) *GetVMDeployOrderResponseBodyDeployOrderActions {
	s.Disable = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderActions) SetParams(v interface{}) *GetVMDeployOrderResponseBodyDeployOrderActions {
	s.Params = v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderActions) SetType(v string) *GetVMDeployOrderResponseBodyDeployOrderActions {
	s.Type = &v
	return s
}

type GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo struct {
	BatchNum       *int32                                                                    `json:"batchNum,omitempty" xml:"batchNum,omitempty"`
	DeployMachines []*GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines `json:"deployMachines,omitempty" xml:"deployMachines,omitempty" type:"Repeated"`
	HostGroupId    *int64                                                                    `json:"hostGroupId,omitempty" xml:"hostGroupId,omitempty"`
}

func (s GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo) String() string {
	return tea.Prettify(s)
}

func (s GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo) GoString() string {
	return s.String()
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo) SetBatchNum(v int32) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo {
	s.BatchNum = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo) SetDeployMachines(v []*GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo {
	s.DeployMachines = v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo) SetHostGroupId(v int64) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo {
	s.HostGroupId = &v
	return s
}

type GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines struct {
	Actions      []*GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	BatchNum     *int32                                                                           `json:"batchNum,omitempty" xml:"batchNum,omitempty"`
	ClientStatus *string                                                                          `json:"clientStatus,omitempty" xml:"clientStatus,omitempty"`
	CreateTime   *int64                                                                           `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Ip           *string                                                                          `json:"ip,omitempty" xml:"ip,omitempty"`
	MachineSn    *string                                                                          `json:"machineSn,omitempty" xml:"machineSn,omitempty"`
	Status       *string                                                                          `json:"status,omitempty" xml:"status,omitempty"`
	UpdateTime   *int64                                                                           `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) String() string {
	return tea.Prettify(s)
}

func (s GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) GoString() string {
	return s.String()
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) SetActions(v []*GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines {
	s.Actions = v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) SetBatchNum(v int32) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines {
	s.BatchNum = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) SetClientStatus(v string) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines {
	s.ClientStatus = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) SetCreateTime(v int64) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines {
	s.CreateTime = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) SetIp(v string) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines {
	s.Ip = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) SetMachineSn(v string) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines {
	s.MachineSn = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) SetStatus(v string) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines {
	s.Status = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines) SetUpdateTime(v int64) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines {
	s.UpdateTime = &v
	return s
}

type GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions struct {
	Disable *bool       `json:"disable,omitempty" xml:"disable,omitempty"`
	Params  interface{} `json:"params,omitempty" xml:"params,omitempty"`
	Type    *string     `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) String() string {
	return tea.Prettify(s)
}

func (s GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) GoString() string {
	return s.String()
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) SetDisable(v bool) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions {
	s.Disable = &v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) SetParams(v interface{}) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions {
	s.Params = v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) SetType(v string) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions {
	s.Type = &v
	return s
}

type GetVMDeployOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVMDeployOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVMDeployOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVMDeployOrderResponse) GoString() string {
	return s.String()
}

func (s *GetVMDeployOrderResponse) SetHeaders(v map[string]*string) *GetVMDeployOrderResponse {
	s.Headers = v
	return s
}

func (s *GetVMDeployOrderResponse) SetStatusCode(v int32) *GetVMDeployOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVMDeployOrderResponse) SetBody(v *GetVMDeployOrderResponseBody) *GetVMDeployOrderResponse {
	s.Body = v
	return s
}

type GetVariableGroupResponseBody struct {
	ErrorCode     *string                                    `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage  *string                                    `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId     *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success       *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
	VariableGroup *GetVariableGroupResponseBodyVariableGroup `json:"variableGroup,omitempty" xml:"variableGroup,omitempty" type:"Struct"`
}

func (s GetVariableGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVariableGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetVariableGroupResponseBody) SetErrorCode(v string) *GetVariableGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetVariableGroupResponseBody) SetErrorMessage(v string) *GetVariableGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetVariableGroupResponseBody) SetRequestId(v string) *GetVariableGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVariableGroupResponseBody) SetSuccess(v bool) *GetVariableGroupResponseBody {
	s.Success = &v
	return s
}

func (s *GetVariableGroupResponseBody) SetVariableGroup(v *GetVariableGroupResponseBodyVariableGroup) *GetVariableGroupResponseBody {
	s.VariableGroup = v
	return s
}

type GetVariableGroupResponseBodyVariableGroup struct {
	CcreatorAccountId *string                                                      `json:"ccreatorAccountId,omitempty" xml:"ccreatorAccountId,omitempty"`
	CreateTime        *int64                                                       `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description       *string                                                      `json:"description,omitempty" xml:"description,omitempty"`
	Id                *int64                                                       `json:"id,omitempty" xml:"id,omitempty"`
	ModifierAccountId *string                                                      `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	Name              *string                                                      `json:"name,omitempty" xml:"name,omitempty"`
	RelatedPipelines  []*GetVariableGroupResponseBodyVariableGroupRelatedPipelines `json:"relatedPipelines,omitempty" xml:"relatedPipelines,omitempty" type:"Repeated"`
	UpdateTime        *int64                                                       `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	Variables         []*GetVariableGroupResponseBodyVariableGroupVariables        `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s GetVariableGroupResponseBodyVariableGroup) String() string {
	return tea.Prettify(s)
}

func (s GetVariableGroupResponseBodyVariableGroup) GoString() string {
	return s.String()
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetCcreatorAccountId(v string) *GetVariableGroupResponseBodyVariableGroup {
	s.CcreatorAccountId = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetCreateTime(v int64) *GetVariableGroupResponseBodyVariableGroup {
	s.CreateTime = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetDescription(v string) *GetVariableGroupResponseBodyVariableGroup {
	s.Description = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetId(v int64) *GetVariableGroupResponseBodyVariableGroup {
	s.Id = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetModifierAccountId(v string) *GetVariableGroupResponseBodyVariableGroup {
	s.ModifierAccountId = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetName(v string) *GetVariableGroupResponseBodyVariableGroup {
	s.Name = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetRelatedPipelines(v []*GetVariableGroupResponseBodyVariableGroupRelatedPipelines) *GetVariableGroupResponseBodyVariableGroup {
	s.RelatedPipelines = v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetUpdateTime(v int64) *GetVariableGroupResponseBodyVariableGroup {
	s.UpdateTime = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroup) SetVariables(v []*GetVariableGroupResponseBodyVariableGroupVariables) *GetVariableGroupResponseBodyVariableGroup {
	s.Variables = v
	return s
}

type GetVariableGroupResponseBodyVariableGroupRelatedPipelines struct {
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetVariableGroupResponseBodyVariableGroupRelatedPipelines) String() string {
	return tea.Prettify(s)
}

func (s GetVariableGroupResponseBodyVariableGroupRelatedPipelines) GoString() string {
	return s.String()
}

func (s *GetVariableGroupResponseBodyVariableGroupRelatedPipelines) SetId(v int64) *GetVariableGroupResponseBodyVariableGroupRelatedPipelines {
	s.Id = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroupRelatedPipelines) SetName(v string) *GetVariableGroupResponseBodyVariableGroupRelatedPipelines {
	s.Name = &v
	return s
}

type GetVariableGroupResponseBodyVariableGroupVariables struct {
	IsEncrypted *bool   `json:"isEncrypted,omitempty" xml:"isEncrypted,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetVariableGroupResponseBodyVariableGroupVariables) String() string {
	return tea.Prettify(s)
}

func (s GetVariableGroupResponseBodyVariableGroupVariables) GoString() string {
	return s.String()
}

func (s *GetVariableGroupResponseBodyVariableGroupVariables) SetIsEncrypted(v bool) *GetVariableGroupResponseBodyVariableGroupVariables {
	s.IsEncrypted = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroupVariables) SetName(v string) *GetVariableGroupResponseBodyVariableGroupVariables {
	s.Name = &v
	return s
}

func (s *GetVariableGroupResponseBodyVariableGroupVariables) SetValue(v string) *GetVariableGroupResponseBodyVariableGroupVariables {
	s.Value = &v
	return s
}

type GetVariableGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVariableGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVariableGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVariableGroupResponse) GoString() string {
	return s.String()
}

func (s *GetVariableGroupResponse) SetHeaders(v map[string]*string) *GetVariableGroupResponse {
	s.Headers = v
	return s
}

func (s *GetVariableGroupResponse) SetStatusCode(v int32) *GetVariableGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVariableGroupResponse) SetBody(v *GetVariableGroupResponseBody) *GetVariableGroupResponse {
	s.Body = v
	return s
}

type GetWorkItemActivityResponseBody struct {
	Activities []*GetWorkItemActivityResponseBodyActivities `json:"activities,omitempty" xml:"activities,omitempty" type:"Repeated"`
	ErrorCode  *string                                      `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg   *string                                      `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId  *string                                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success    *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetWorkItemActivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkItemActivityResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkItemActivityResponseBody) SetActivities(v []*GetWorkItemActivityResponseBodyActivities) *GetWorkItemActivityResponseBody {
	s.Activities = v
	return s
}

func (s *GetWorkItemActivityResponseBody) SetErrorCode(v string) *GetWorkItemActivityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkItemActivityResponseBody) SetErrorMsg(v string) *GetWorkItemActivityResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetWorkItemActivityResponseBody) SetRequestId(v string) *GetWorkItemActivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkItemActivityResponseBody) SetSuccess(v bool) *GetWorkItemActivityResponseBody {
	s.Success = &v
	return s
}

type GetWorkItemActivityResponseBodyActivities struct {
	ActionType         *string                                              `json:"actionType,omitempty" xml:"actionType,omitempty"`
	EventId            *int64                                               `json:"eventId,omitempty" xml:"eventId,omitempty"`
	EventTime          *int64                                               `json:"eventTime,omitempty" xml:"eventTime,omitempty"`
	EventType          *string                                              `json:"eventType,omitempty" xml:"eventType,omitempty"`
	NewValue           []*GetWorkItemActivityResponseBodyActivitiesNewValue `json:"newValue,omitempty" xml:"newValue,omitempty" type:"Repeated"`
	OldValue           []*GetWorkItemActivityResponseBodyActivitiesOldValue `json:"oldValue,omitempty" xml:"oldValue,omitempty" type:"Repeated"`
	Operator           *string                                              `json:"operator,omitempty" xml:"operator,omitempty"`
	ParentEventId      *int64                                               `json:"parentEventId,omitempty" xml:"parentEventId,omitempty"`
	Property           *GetWorkItemActivityResponseBodyActivitiesProperty   `json:"property,omitempty" xml:"property,omitempty" type:"Struct"`
	ResourceIdentifier *string                                              `json:"resourceIdentifier,omitempty" xml:"resourceIdentifier,omitempty"`
}

func (s GetWorkItemActivityResponseBodyActivities) String() string {
	return tea.Prettify(s)
}

func (s GetWorkItemActivityResponseBodyActivities) GoString() string {
	return s.String()
}

func (s *GetWorkItemActivityResponseBodyActivities) SetActionType(v string) *GetWorkItemActivityResponseBodyActivities {
	s.ActionType = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetEventId(v int64) *GetWorkItemActivityResponseBodyActivities {
	s.EventId = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetEventTime(v int64) *GetWorkItemActivityResponseBodyActivities {
	s.EventTime = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetEventType(v string) *GetWorkItemActivityResponseBodyActivities {
	s.EventType = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetNewValue(v []*GetWorkItemActivityResponseBodyActivitiesNewValue) *GetWorkItemActivityResponseBodyActivities {
	s.NewValue = v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetOldValue(v []*GetWorkItemActivityResponseBodyActivitiesOldValue) *GetWorkItemActivityResponseBodyActivities {
	s.OldValue = v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetOperator(v string) *GetWorkItemActivityResponseBodyActivities {
	s.Operator = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetParentEventId(v int64) *GetWorkItemActivityResponseBodyActivities {
	s.ParentEventId = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetProperty(v *GetWorkItemActivityResponseBodyActivitiesProperty) *GetWorkItemActivityResponseBodyActivities {
	s.Property = v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetResourceIdentifier(v string) *GetWorkItemActivityResponseBodyActivities {
	s.ResourceIdentifier = &v
	return s
}

type GetWorkItemActivityResponseBodyActivitiesNewValue struct {
	DisplayValue *string `json:"displayValue,omitempty" xml:"displayValue,omitempty"`
	PlainValue   *string `json:"plainValue,omitempty" xml:"plainValue,omitempty"`
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetWorkItemActivityResponseBodyActivitiesNewValue) String() string {
	return tea.Prettify(s)
}

func (s GetWorkItemActivityResponseBodyActivitiesNewValue) GoString() string {
	return s.String()
}

func (s *GetWorkItemActivityResponseBodyActivitiesNewValue) SetDisplayValue(v string) *GetWorkItemActivityResponseBodyActivitiesNewValue {
	s.DisplayValue = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivitiesNewValue) SetPlainValue(v string) *GetWorkItemActivityResponseBodyActivitiesNewValue {
	s.PlainValue = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivitiesNewValue) SetResourceType(v string) *GetWorkItemActivityResponseBodyActivitiesNewValue {
	s.ResourceType = &v
	return s
}

type GetWorkItemActivityResponseBodyActivitiesOldValue struct {
	DisplayValue *string `json:"displayValue,omitempty" xml:"displayValue,omitempty"`
	PlainValue   *string `json:"plainValue,omitempty" xml:"plainValue,omitempty"`
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetWorkItemActivityResponseBodyActivitiesOldValue) String() string {
	return tea.Prettify(s)
}

func (s GetWorkItemActivityResponseBodyActivitiesOldValue) GoString() string {
	return s.String()
}

func (s *GetWorkItemActivityResponseBodyActivitiesOldValue) SetDisplayValue(v string) *GetWorkItemActivityResponseBodyActivitiesOldValue {
	s.DisplayValue = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivitiesOldValue) SetPlainValue(v string) *GetWorkItemActivityResponseBodyActivitiesOldValue {
	s.PlainValue = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivitiesOldValue) SetResourceType(v string) *GetWorkItemActivityResponseBodyActivitiesOldValue {
	s.ResourceType = &v
	return s
}

type GetWorkItemActivityResponseBodyActivitiesProperty struct {
	DisplayName        *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	PropertyIdentifier *string `json:"propertyIdentifier,omitempty" xml:"propertyIdentifier,omitempty"`
	PropertyName       *string `json:"propertyName,omitempty" xml:"propertyName,omitempty"`
	PropertyType       *string `json:"propertyType,omitempty" xml:"propertyType,omitempty"`
}

func (s GetWorkItemActivityResponseBodyActivitiesProperty) String() string {
	return tea.Prettify(s)
}

func (s GetWorkItemActivityResponseBodyActivitiesProperty) GoString() string {
	return s.String()
}

func (s *GetWorkItemActivityResponseBodyActivitiesProperty) SetDisplayName(v string) *GetWorkItemActivityResponseBodyActivitiesProperty {
	s.DisplayName = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivitiesProperty) SetPropertyIdentifier(v string) *GetWorkItemActivityResponseBodyActivitiesProperty {
	s.PropertyIdentifier = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivitiesProperty) SetPropertyName(v string) *GetWorkItemActivityResponseBodyActivitiesProperty {
	s.PropertyName = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivitiesProperty) SetPropertyType(v string) *GetWorkItemActivityResponseBodyActivitiesProperty {
	s.PropertyType = &v
	return s
}

type GetWorkItemActivityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWorkItemActivityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWorkItemActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkItemActivityResponse) GoString() string {
	return s.String()
}

func (s *GetWorkItemActivityResponse) SetHeaders(v map[string]*string) *GetWorkItemActivityResponse {
	s.Headers = v
	return s
}

func (s *GetWorkItemActivityResponse) SetStatusCode(v int32) *GetWorkItemActivityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkItemActivityResponse) SetBody(v *GetWorkItemActivityResponseBody) *GetWorkItemActivityResponse {
	s.Body = v
	return s
}

type GetWorkItemInfoResponseBody struct {
	ErrorCode    *string                              `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                              `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                `json:"success,omitempty" xml:"success,omitempty"`
	Workitem     *GetWorkItemInfoResponseBodyWorkitem `json:"workitem,omitempty" xml:"workitem,omitempty" type:"Struct"`
}

func (s GetWorkItemInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkItemInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkItemInfoResponseBody) SetErrorCode(v string) *GetWorkItemInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkItemInfoResponseBody) SetErrorMessage(v string) *GetWorkItemInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetWorkItemInfoResponseBody) SetRequestId(v string) *GetWorkItemInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkItemInfoResponseBody) SetSuccess(v bool) *GetWorkItemInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkItemInfoResponseBody) SetWorkitem(v *GetWorkItemInfoResponseBodyWorkitem) *GetWorkItemInfoResponseBody {
	s.Workitem = v
	return s
}

type GetWorkItemInfoResponseBodyWorkitem struct {
	AssignedTo             *string                                            `json:"assignedTo,omitempty" xml:"assignedTo,omitempty"`
	CategoryIdentifier     *string                                            `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	Creator                *string                                            `json:"creator,omitempty" xml:"creator,omitempty"`
	CustomFields           []*GetWorkItemInfoResponseBodyWorkitemCustomFields `json:"customFields,omitempty" xml:"customFields,omitempty" type:"Repeated"`
	Document               *string                                            `json:"document,omitempty" xml:"document,omitempty"`
	GmtCreate              *int64                                             `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified            *int64                                             `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Identifier             *string                                            `json:"identifier,omitempty" xml:"identifier,omitempty"`
	LogicalStatus          *string                                            `json:"logicalStatus,omitempty" xml:"logicalStatus,omitempty"`
	Modifier               *string                                            `json:"modifier,omitempty" xml:"modifier,omitempty"`
	ParentIdentifier       *string                                            `json:"parentIdentifier,omitempty" xml:"parentIdentifier,omitempty"`
	Participant            []*string                                          `json:"participant,omitempty" xml:"participant,omitempty" type:"Repeated"`
	SerialNumber           *string                                            `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	SpaceIdentifier        *string                                            `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	SpaceName              *string                                            `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	SpaceType              *string                                            `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	Sprint                 []*string                                          `json:"sprint,omitempty" xml:"sprint,omitempty" type:"Repeated"`
	Status                 *string                                            `json:"status,omitempty" xml:"status,omitempty"`
	StatusIdentifier       *string                                            `json:"statusIdentifier,omitempty" xml:"statusIdentifier,omitempty"`
	StatusStageIdentifier  *string                                            `json:"statusStageIdentifier,omitempty" xml:"statusStageIdentifier,omitempty"`
	Subject                *string                                            `json:"subject,omitempty" xml:"subject,omitempty"`
	Tag                    []*string                                          `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
	Tracker                []*string                                          `json:"tracker,omitempty" xml:"tracker,omitempty" type:"Repeated"`
	UpdateStatusAt         *int64                                             `json:"updateStatusAt,omitempty" xml:"updateStatusAt,omitempty"`
	Verifier               []*string                                          `json:"verifier,omitempty" xml:"verifier,omitempty" type:"Repeated"`
	WorkitemTypeIdentifier *string                                            `json:"workitemTypeIdentifier,omitempty" xml:"workitemTypeIdentifier,omitempty"`
}

func (s GetWorkItemInfoResponseBodyWorkitem) String() string {
	return tea.Prettify(s)
}

func (s GetWorkItemInfoResponseBodyWorkitem) GoString() string {
	return s.String()
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetAssignedTo(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.AssignedTo = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetCategoryIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.CategoryIdentifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetCreator(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Creator = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetCustomFields(v []*GetWorkItemInfoResponseBodyWorkitemCustomFields) *GetWorkItemInfoResponseBodyWorkitem {
	s.CustomFields = v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetDocument(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Document = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetGmtCreate(v int64) *GetWorkItemInfoResponseBodyWorkitem {
	s.GmtCreate = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetGmtModified(v int64) *GetWorkItemInfoResponseBodyWorkitem {
	s.GmtModified = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Identifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetLogicalStatus(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.LogicalStatus = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetModifier(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Modifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetParentIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.ParentIdentifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetParticipant(v []*string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Participant = v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetSerialNumber(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.SerialNumber = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetSpaceIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.SpaceIdentifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetSpaceName(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.SpaceName = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetSpaceType(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.SpaceType = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetSprint(v []*string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Sprint = v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetStatus(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Status = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetStatusIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.StatusIdentifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetStatusStageIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.StatusStageIdentifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetSubject(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Subject = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetTag(v []*string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Tag = v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetTracker(v []*string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Tracker = v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetUpdateStatusAt(v int64) *GetWorkItemInfoResponseBodyWorkitem {
	s.UpdateStatusAt = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetVerifier(v []*string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Verifier = v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetWorkitemTypeIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.WorkitemTypeIdentifier = &v
	return s
}

type GetWorkItemInfoResponseBodyWorkitemCustomFields struct {
	FieldClassName     *string                                                     `json:"fieldClassName,omitempty" xml:"fieldClassName,omitempty"`
	FieldFormat        *string                                                     `json:"fieldFormat,omitempty" xml:"fieldFormat,omitempty"`
	FieldIdentifier    *string                                                     `json:"fieldIdentifier,omitempty" xml:"fieldIdentifier,omitempty"`
	Level              *int64                                                      `json:"level,omitempty" xml:"level,omitempty"`
	ObjectValue        *string                                                     `json:"objectValue,omitempty" xml:"objectValue,omitempty"`
	Position           *int64                                                      `json:"position,omitempty" xml:"position,omitempty"`
	Value              *string                                                     `json:"value,omitempty" xml:"value,omitempty"`
	ValueList          []*GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList `json:"valueList,omitempty" xml:"valueList,omitempty" type:"Repeated"`
	WorkitemIdentifier *string                                                     `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s GetWorkItemInfoResponseBodyWorkitemCustomFields) String() string {
	return tea.Prettify(s)
}

func (s GetWorkItemInfoResponseBodyWorkitemCustomFields) GoString() string {
	return s.String()
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetFieldClassName(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.FieldClassName = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetFieldFormat(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.FieldFormat = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetFieldIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.FieldIdentifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetLevel(v int64) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.Level = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetObjectValue(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.ObjectValue = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetPosition(v int64) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.Position = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetValue(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.Value = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetValueList(v []*GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.ValueList = v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFields) SetWorkitemIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFields {
	s.WorkitemIdentifier = &v
	return s
}

type GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList struct {
	DisplayValue *string `json:"displayValue,omitempty" xml:"displayValue,omitempty"`
	Identifier   *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Level        *int64  `json:"level,omitempty" xml:"level,omitempty"`
	Value        *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueEn      *string `json:"valueEn,omitempty" xml:"valueEn,omitempty"`
}

func (s GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) String() string {
	return tea.Prettify(s)
}

func (s GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) GoString() string {
	return s.String()
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) SetDisplayValue(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList {
	s.DisplayValue = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) SetIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList {
	s.Identifier = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) SetLevel(v int64) *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList {
	s.Level = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) SetValue(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList {
	s.Value = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList) SetValueEn(v string) *GetWorkItemInfoResponseBodyWorkitemCustomFieldsValueList {
	s.ValueEn = &v
	return s
}

type GetWorkItemInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWorkItemInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWorkItemInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkItemInfoResponse) GoString() string {
	return s.String()
}

func (s *GetWorkItemInfoResponse) SetHeaders(v map[string]*string) *GetWorkItemInfoResponse {
	s.Headers = v
	return s
}

func (s *GetWorkItemInfoResponse) SetStatusCode(v int32) *GetWorkItemInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkItemInfoResponse) SetBody(v *GetWorkItemInfoResponseBody) *GetWorkItemInfoResponse {
	s.Body = v
	return s
}

type GetWorkItemWorkFlowInfoRequest struct {
	ConfigurationId *string `json:"configurationId,omitempty" xml:"configurationId,omitempty"`
}

func (s GetWorkItemWorkFlowInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkItemWorkFlowInfoRequest) GoString() string {
	return s.String()
}

func (s *GetWorkItemWorkFlowInfoRequest) SetConfigurationId(v string) *GetWorkItemWorkFlowInfoRequest {
	s.ConfigurationId = &v
	return s
}

type GetWorkItemWorkFlowInfoResponseBody struct {
	ErrorCode    *string                                      `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                      `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
	Workflow     *GetWorkItemWorkFlowInfoResponseBodyWorkflow `json:"workflow,omitempty" xml:"workflow,omitempty" type:"Struct"`
}

func (s GetWorkItemWorkFlowInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkItemWorkFlowInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkItemWorkFlowInfoResponseBody) SetErrorCode(v string) *GetWorkItemWorkFlowInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBody) SetErrorMessage(v string) *GetWorkItemWorkFlowInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBody) SetRequestId(v string) *GetWorkItemWorkFlowInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBody) SetSuccess(v bool) *GetWorkItemWorkFlowInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBody) SetWorkflow(v *GetWorkItemWorkFlowInfoResponseBodyWorkflow) *GetWorkItemWorkFlowInfoResponseBody {
	s.Workflow = v
	return s
}

type GetWorkItemWorkFlowInfoResponseBodyWorkflow struct {
	Creator                 *string                                                       `json:"creator,omitempty" xml:"creator,omitempty"`
	DefaultStatusIdentifier *string                                                       `json:"defaultStatusIdentifier,omitempty" xml:"defaultStatusIdentifier,omitempty"`
	Description             *string                                                       `json:"description,omitempty" xml:"description,omitempty"`
	GmtCreate               *int64                                                        `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified             *int64                                                        `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Identifier              *string                                                       `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Modifier                *string                                                       `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Name                    *string                                                       `json:"name,omitempty" xml:"name,omitempty"`
	OwnerSpaceIdentifier    *string                                                       `json:"ownerSpaceIdentifier,omitempty" xml:"ownerSpaceIdentifier,omitempty"`
	OwnerSpaceType          *string                                                       `json:"ownerSpaceType,omitempty" xml:"ownerSpaceType,omitempty"`
	ResourceType            *string                                                       `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Source                  *string                                                       `json:"source,omitempty" xml:"source,omitempty"`
	StatusOrder             *string                                                       `json:"statusOrder,omitempty" xml:"statusOrder,omitempty"`
	Statuses                []*GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses        `json:"statuses,omitempty" xml:"statuses,omitempty" type:"Repeated"`
	WorkflowActions         []*GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions `json:"workflowActions,omitempty" xml:"workflowActions,omitempty" type:"Repeated"`
}

func (s GetWorkItemWorkFlowInfoResponseBodyWorkflow) String() string {
	return tea.Prettify(s)
}

func (s GetWorkItemWorkFlowInfoResponseBodyWorkflow) GoString() string {
	return s.String()
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetCreator(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.Creator = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetDefaultStatusIdentifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.DefaultStatusIdentifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetDescription(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.Description = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetGmtCreate(v int64) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.GmtCreate = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetGmtModified(v int64) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.GmtModified = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetIdentifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.Identifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetModifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.Modifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetName(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.Name = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetOwnerSpaceIdentifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.OwnerSpaceIdentifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetOwnerSpaceType(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.OwnerSpaceType = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetResourceType(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.ResourceType = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetSource(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.Source = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetStatusOrder(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.StatusOrder = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetStatuses(v []*GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.Statuses = v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetWorkflowActions(v []*GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.WorkflowActions = v
	return s
}

type GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses struct {
	Creator                 *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Description             *string `json:"description,omitempty" xml:"description,omitempty"`
	GmtCreate               *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified             *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Identifier              *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Modifier                *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Name                    *string `json:"name,omitempty" xml:"name,omitempty"`
	ResourceType            *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Source                  *string `json:"source,omitempty" xml:"source,omitempty"`
	WorkflowStageIdentifier *string `json:"workflowStageIdentifier,omitempty" xml:"workflowStageIdentifier,omitempty"`
	WorkflowStageName       *string `json:"workflowStageName,omitempty" xml:"workflowStageName,omitempty"`
}

func (s GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) String() string {
	return tea.Prettify(s)
}

func (s GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) GoString() string {
	return s.String()
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetCreator(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.Creator = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetDescription(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.Description = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetGmtCreate(v int64) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.GmtCreate = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetGmtModified(v int64) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.GmtModified = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetIdentifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.Identifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetModifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.Modifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetName(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.Name = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetResourceType(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.ResourceType = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetSource(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.Source = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetWorkflowStageIdentifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.WorkflowStageIdentifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetWorkflowStageName(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.WorkflowStageName = &v
	return s
}

type GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions struct {
	Id                           *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name                         *string `json:"name,omitempty" xml:"name,omitempty"`
	NextWorkflowStatusIdentifier *string `json:"nextWorkflowStatusIdentifier,omitempty" xml:"nextWorkflowStatusIdentifier,omitempty"`
	WorkflowIdentifier           *string `json:"workflowIdentifier,omitempty" xml:"workflowIdentifier,omitempty"`
	WorkflowStatusIdentifier     *string `json:"workflowStatusIdentifier,omitempty" xml:"workflowStatusIdentifier,omitempty"`
}

func (s GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) String() string {
	return tea.Prettify(s)
}

func (s GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) GoString() string {
	return s.String()
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) SetId(v int64) *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions {
	s.Id = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) SetName(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions {
	s.Name = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) SetNextWorkflowStatusIdentifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions {
	s.NextWorkflowStatusIdentifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) SetWorkflowIdentifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions {
	s.WorkflowIdentifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) SetWorkflowStatusIdentifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions {
	s.WorkflowStatusIdentifier = &v
	return s
}

type GetWorkItemWorkFlowInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWorkItemWorkFlowInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWorkItemWorkFlowInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkItemWorkFlowInfoResponse) GoString() string {
	return s.String()
}

func (s *GetWorkItemWorkFlowInfoResponse) SetHeaders(v map[string]*string) *GetWorkItemWorkFlowInfoResponse {
	s.Headers = v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponse) SetStatusCode(v int32) *GetWorkItemWorkFlowInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponse) SetBody(v *GetWorkItemWorkFlowInfoResponseBody) *GetWorkItemWorkFlowInfoResponse {
	s.Body = v
	return s
}

type GetWorkitemCommentListResponseBody struct {
	CommentList []*GetWorkitemCommentListResponseBodyCommentList `json:"commentList,omitempty" xml:"commentList,omitempty" type:"Repeated"`
	ErrorCode   *string                                          `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg    *string                                          `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId   *string                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success     *string                                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetWorkitemCommentListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkitemCommentListResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkitemCommentListResponseBody) SetCommentList(v []*GetWorkitemCommentListResponseBodyCommentList) *GetWorkitemCommentListResponseBody {
	s.CommentList = v
	return s
}

func (s *GetWorkitemCommentListResponseBody) SetErrorCode(v string) *GetWorkitemCommentListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkitemCommentListResponseBody) SetErrorMsg(v string) *GetWorkitemCommentListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetWorkitemCommentListResponseBody) SetRequestId(v string) *GetWorkitemCommentListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkitemCommentListResponseBody) SetSuccess(v string) *GetWorkitemCommentListResponseBody {
	s.Success = &v
	return s
}

type GetWorkitemCommentListResponseBodyCommentList struct {
	Content          *string `json:"content,omitempty" xml:"content,omitempty"`
	CreateTime       *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	FormatType       *string `json:"formatType,omitempty" xml:"formatType,omitempty"`
	Id               *int64  `json:"id,omitempty" xml:"id,omitempty"`
	IsTop            *bool   `json:"isTop,omitempty" xml:"isTop,omitempty"`
	ModifiedTime     *int64  `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ParentId         *int64  `json:"parentId,omitempty" xml:"parentId,omitempty"`
	TargetIdentifier *string `json:"targetIdentifier,omitempty" xml:"targetIdentifier,omitempty"`
	TargetType       *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	TopTime          *int64  `json:"topTime,omitempty" xml:"topTime,omitempty"`
}

func (s GetWorkitemCommentListResponseBodyCommentList) String() string {
	return tea.Prettify(s)
}

func (s GetWorkitemCommentListResponseBodyCommentList) GoString() string {
	return s.String()
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetContent(v string) *GetWorkitemCommentListResponseBodyCommentList {
	s.Content = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetCreateTime(v int64) *GetWorkitemCommentListResponseBodyCommentList {
	s.CreateTime = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetFormatType(v string) *GetWorkitemCommentListResponseBodyCommentList {
	s.FormatType = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetId(v int64) *GetWorkitemCommentListResponseBodyCommentList {
	s.Id = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetIsTop(v bool) *GetWorkitemCommentListResponseBodyCommentList {
	s.IsTop = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetModifiedTime(v int64) *GetWorkitemCommentListResponseBodyCommentList {
	s.ModifiedTime = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetParentId(v int64) *GetWorkitemCommentListResponseBodyCommentList {
	s.ParentId = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetTargetIdentifier(v string) *GetWorkitemCommentListResponseBodyCommentList {
	s.TargetIdentifier = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetTargetType(v string) *GetWorkitemCommentListResponseBodyCommentList {
	s.TargetType = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetTopTime(v int64) *GetWorkitemCommentListResponseBodyCommentList {
	s.TopTime = &v
	return s
}

type GetWorkitemCommentListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWorkitemCommentListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWorkitemCommentListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkitemCommentListResponse) GoString() string {
	return s.String()
}

func (s *GetWorkitemCommentListResponse) SetHeaders(v map[string]*string) *GetWorkitemCommentListResponse {
	s.Headers = v
	return s
}

func (s *GetWorkitemCommentListResponse) SetStatusCode(v int32) *GetWorkitemCommentListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkitemCommentListResponse) SetBody(v *GetWorkitemCommentListResponseBody) *GetWorkitemCommentListResponse {
	s.Body = v
	return s
}

type GetWorkitemRelationsRequest struct {
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s GetWorkitemRelationsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkitemRelationsRequest) GoString() string {
	return s.String()
}

func (s *GetWorkitemRelationsRequest) SetRelationType(v string) *GetWorkitemRelationsRequest {
	s.RelationType = &v
	return s
}

type GetWorkitemRelationsResponseBody struct {
	ErrorCode    *string                                         `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg     *string                                         `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RelationList []*GetWorkitemRelationsResponseBodyRelationList `json:"relationList,omitempty" xml:"relationList,omitempty" type:"Repeated"`
	RequestId    *string                                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetWorkitemRelationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkitemRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkitemRelationsResponseBody) SetErrorCode(v string) *GetWorkitemRelationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkitemRelationsResponseBody) SetErrorMsg(v string) *GetWorkitemRelationsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetWorkitemRelationsResponseBody) SetRelationList(v []*GetWorkitemRelationsResponseBodyRelationList) *GetWorkitemRelationsResponseBody {
	s.RelationList = v
	return s
}

func (s *GetWorkitemRelationsResponseBody) SetRequestId(v string) *GetWorkitemRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkitemRelationsResponseBody) SetSuccess(v bool) *GetWorkitemRelationsResponseBody {
	s.Success = &v
	return s
}

type GetWorkitemRelationsResponseBodyRelationList struct {
	AssignedTo             *string `json:"assignedTo,omitempty" xml:"assignedTo,omitempty"`
	CategoryIdentifier     *string `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	GmtCreate              *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified            *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Identifier             *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	SpaceIdentifier        *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	Subject                *string `json:"subject,omitempty" xml:"subject,omitempty"`
	WorkitemTypeIdentifier *string `json:"workitemTypeIdentifier,omitempty" xml:"workitemTypeIdentifier,omitempty"`
}

func (s GetWorkitemRelationsResponseBodyRelationList) String() string {
	return tea.Prettify(s)
}

func (s GetWorkitemRelationsResponseBodyRelationList) GoString() string {
	return s.String()
}

func (s *GetWorkitemRelationsResponseBodyRelationList) SetAssignedTo(v string) *GetWorkitemRelationsResponseBodyRelationList {
	s.AssignedTo = &v
	return s
}

func (s *GetWorkitemRelationsResponseBodyRelationList) SetCategoryIdentifier(v string) *GetWorkitemRelationsResponseBodyRelationList {
	s.CategoryIdentifier = &v
	return s
}

func (s *GetWorkitemRelationsResponseBodyRelationList) SetGmtCreate(v string) *GetWorkitemRelationsResponseBodyRelationList {
	s.GmtCreate = &v
	return s
}

func (s *GetWorkitemRelationsResponseBodyRelationList) SetGmtModified(v string) *GetWorkitemRelationsResponseBodyRelationList {
	s.GmtModified = &v
	return s
}

func (s *GetWorkitemRelationsResponseBodyRelationList) SetIdentifier(v string) *GetWorkitemRelationsResponseBodyRelationList {
	s.Identifier = &v
	return s
}

func (s *GetWorkitemRelationsResponseBodyRelationList) SetSpaceIdentifier(v string) *GetWorkitemRelationsResponseBodyRelationList {
	s.SpaceIdentifier = &v
	return s
}

func (s *GetWorkitemRelationsResponseBodyRelationList) SetSubject(v string) *GetWorkitemRelationsResponseBodyRelationList {
	s.Subject = &v
	return s
}

func (s *GetWorkitemRelationsResponseBodyRelationList) SetWorkitemTypeIdentifier(v string) *GetWorkitemRelationsResponseBodyRelationList {
	s.WorkitemTypeIdentifier = &v
	return s
}

type GetWorkitemRelationsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWorkitemRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWorkitemRelationsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkitemRelationsResponse) GoString() string {
	return s.String()
}

func (s *GetWorkitemRelationsResponse) SetHeaders(v map[string]*string) *GetWorkitemRelationsResponse {
	s.Headers = v
	return s
}

func (s *GetWorkitemRelationsResponse) SetStatusCode(v int32) *GetWorkitemRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkitemRelationsResponse) SetBody(v *GetWorkitemRelationsResponseBody) *GetWorkitemRelationsResponse {
	s.Body = v
	return s
}

type GetWorkitemTimeTypeListResponseBody struct {
	ErrorCode *string                                        `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                        `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId *string                                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *string                                        `json:"success,omitempty" xml:"success,omitempty"`
	TimeType  []*GetWorkitemTimeTypeListResponseBodyTimeType `json:"timeType,omitempty" xml:"timeType,omitempty" type:"Repeated"`
}

func (s GetWorkitemTimeTypeListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkitemTimeTypeListResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkitemTimeTypeListResponseBody) SetErrorCode(v string) *GetWorkitemTimeTypeListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBody) SetErrorMsg(v string) *GetWorkitemTimeTypeListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBody) SetRequestId(v string) *GetWorkitemTimeTypeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBody) SetSuccess(v string) *GetWorkitemTimeTypeListResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBody) SetTimeType(v []*GetWorkitemTimeTypeListResponseBodyTimeType) *GetWorkitemTimeTypeListResponseBody {
	s.TimeType = v
	return s
}

type GetWorkitemTimeTypeListResponseBodyTimeType struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Identifier  *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Position    *int64  `json:"position,omitempty" xml:"position,omitempty"`
}

func (s GetWorkitemTimeTypeListResponseBodyTimeType) String() string {
	return tea.Prettify(s)
}

func (s GetWorkitemTimeTypeListResponseBodyTimeType) GoString() string {
	return s.String()
}

func (s *GetWorkitemTimeTypeListResponseBodyTimeType) SetDescription(v string) *GetWorkitemTimeTypeListResponseBodyTimeType {
	s.Description = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBodyTimeType) SetDisplayName(v string) *GetWorkitemTimeTypeListResponseBodyTimeType {
	s.DisplayName = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBodyTimeType) SetIdentifier(v string) *GetWorkitemTimeTypeListResponseBodyTimeType {
	s.Identifier = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBodyTimeType) SetName(v string) *GetWorkitemTimeTypeListResponseBodyTimeType {
	s.Name = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponseBodyTimeType) SetPosition(v int64) *GetWorkitemTimeTypeListResponseBodyTimeType {
	s.Position = &v
	return s
}

type GetWorkitemTimeTypeListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWorkitemTimeTypeListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWorkitemTimeTypeListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkitemTimeTypeListResponse) GoString() string {
	return s.String()
}

func (s *GetWorkitemTimeTypeListResponse) SetHeaders(v map[string]*string) *GetWorkitemTimeTypeListResponse {
	s.Headers = v
	return s
}

func (s *GetWorkitemTimeTypeListResponse) SetStatusCode(v int32) *GetWorkitemTimeTypeListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkitemTimeTypeListResponse) SetBody(v *GetWorkitemTimeTypeListResponseBody) *GetWorkitemTimeTypeListResponse {
	s.Body = v
	return s
}

type GetWorkspaceResponseBody struct {
	ErrorCode    *string                            `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                            `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                              `json:"success,omitempty" xml:"success,omitempty"`
	Workspace    *GetWorkspaceResponseBodyWorkspace `json:"workspace,omitempty" xml:"workspace,omitempty" type:"Struct"`
}

func (s GetWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) SetErrorCode(v string) *GetWorkspaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetErrorMessage(v string) *GetWorkspaceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetRequestId(v string) *GetWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetSuccess(v bool) *GetWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetWorkspace(v *GetWorkspaceResponseBodyWorkspace) *GetWorkspaceResponseBody {
	s.Workspace = v
	return s
}

type GetWorkspaceResponseBodyWorkspace struct {
	CodeUrl     *string `json:"codeUrl,omitempty" xml:"codeUrl,omitempty"`
	CodeVersion *string `json:"codeVersion,omitempty" xml:"codeVersion,omitempty"`
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Spec        *string `json:"spec,omitempty" xml:"spec,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
	Template    *string `json:"template,omitempty" xml:"template,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetWorkspaceResponseBodyWorkspace) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponseBodyWorkspace) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyWorkspace) SetCodeUrl(v string) *GetWorkspaceResponseBodyWorkspace {
	s.CodeUrl = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetCodeVersion(v string) *GetWorkspaceResponseBodyWorkspace {
	s.CodeVersion = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetCreateTime(v string) *GetWorkspaceResponseBodyWorkspace {
	s.CreateTime = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.Id = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetName(v string) *GetWorkspaceResponseBodyWorkspace {
	s.Name = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetSpec(v string) *GetWorkspaceResponseBodyWorkspace {
	s.Spec = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetStatus(v string) *GetWorkspaceResponseBodyWorkspace {
	s.Status = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetTemplate(v string) *GetWorkspaceResponseBodyWorkspace {
	s.Template = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetUserId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.UserId = &v
	return s
}

type GetWorkspaceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponse) SetHeaders(v map[string]*string) *GetWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspaceResponse) SetStatusCode(v int32) *GetWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkspaceResponse) SetBody(v *GetWorkspaceResponseBody) *GetWorkspaceResponse {
	s.Body = v
	return s
}

type JoinPipelineGroupRequest struct {
	GroupId     *int64  `json:"groupId,omitempty" xml:"groupId,omitempty"`
	PipelineIds *string `json:"pipelineIds,omitempty" xml:"pipelineIds,omitempty"`
}

func (s JoinPipelineGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinPipelineGroupRequest) GoString() string {
	return s.String()
}

func (s *JoinPipelineGroupRequest) SetGroupId(v int64) *JoinPipelineGroupRequest {
	s.GroupId = &v
	return s
}

func (s *JoinPipelineGroupRequest) SetPipelineIds(v string) *JoinPipelineGroupRequest {
	s.PipelineIds = &v
	return s
}

type JoinPipelineGroupResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s JoinPipelineGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinPipelineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *JoinPipelineGroupResponseBody) SetErrorCode(v string) *JoinPipelineGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *JoinPipelineGroupResponseBody) SetErrorMessage(v string) *JoinPipelineGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *JoinPipelineGroupResponseBody) SetRequestId(v string) *JoinPipelineGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinPipelineGroupResponseBody) SetSuccess(v bool) *JoinPipelineGroupResponseBody {
	s.Success = &v
	return s
}

type JoinPipelineGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *JoinPipelineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s JoinPipelineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinPipelineGroupResponse) GoString() string {
	return s.String()
}

func (s *JoinPipelineGroupResponse) SetHeaders(v map[string]*string) *JoinPipelineGroupResponse {
	s.Headers = v
	return s
}

func (s *JoinPipelineGroupResponse) SetStatusCode(v int32) *JoinPipelineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinPipelineGroupResponse) SetBody(v *JoinPipelineGroupResponseBody) *JoinPipelineGroupResponse {
	s.Body = v
	return s
}

type ListFlowTagGroupsResponseBody struct {
	ErrorCode     *string                                       `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage  *string                                       `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	FlowTagGroups []*ListFlowTagGroupsResponseBodyFlowTagGroups `json:"flowTagGroups,omitempty" xml:"flowTagGroups,omitempty" type:"Repeated"`
	RequestId     *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success       *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListFlowTagGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowTagGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowTagGroupsResponseBody) SetErrorCode(v string) *ListFlowTagGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListFlowTagGroupsResponseBody) SetErrorMessage(v string) *ListFlowTagGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListFlowTagGroupsResponseBody) SetFlowTagGroups(v []*ListFlowTagGroupsResponseBodyFlowTagGroups) *ListFlowTagGroupsResponseBody {
	s.FlowTagGroups = v
	return s
}

func (s *ListFlowTagGroupsResponseBody) SetRequestId(v string) *ListFlowTagGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowTagGroupsResponseBody) SetSuccess(v bool) *ListFlowTagGroupsResponseBody {
	s.Success = &v
	return s
}

type ListFlowTagGroupsResponseBodyFlowTagGroups struct {
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	Id               *int64  `json:"id,omitempty" xml:"id,omitempty"`
	ModiferAccountId *string `json:"modiferAccountId,omitempty" xml:"modiferAccountId,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListFlowTagGroupsResponseBodyFlowTagGroups) String() string {
	return tea.Prettify(s)
}

func (s ListFlowTagGroupsResponseBodyFlowTagGroups) GoString() string {
	return s.String()
}

func (s *ListFlowTagGroupsResponseBodyFlowTagGroups) SetCreatorAccountId(v string) *ListFlowTagGroupsResponseBodyFlowTagGroups {
	s.CreatorAccountId = &v
	return s
}

func (s *ListFlowTagGroupsResponseBodyFlowTagGroups) SetId(v int64) *ListFlowTagGroupsResponseBodyFlowTagGroups {
	s.Id = &v
	return s
}

func (s *ListFlowTagGroupsResponseBodyFlowTagGroups) SetModiferAccountId(v string) *ListFlowTagGroupsResponseBodyFlowTagGroups {
	s.ModiferAccountId = &v
	return s
}

func (s *ListFlowTagGroupsResponseBodyFlowTagGroups) SetName(v string) *ListFlowTagGroupsResponseBodyFlowTagGroups {
	s.Name = &v
	return s
}

type ListFlowTagGroupsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFlowTagGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowTagGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowTagGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowTagGroupsResponse) SetHeaders(v map[string]*string) *ListFlowTagGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListFlowTagGroupsResponse) SetStatusCode(v int32) *ListFlowTagGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlowTagGroupsResponse) SetBody(v *ListFlowTagGroupsResponseBody) *ListFlowTagGroupsResponse {
	s.Body = v
	return s
}

type ListHostGroupsRequest struct {
	CreateEndTime     *int64  `json:"createEndTime,omitempty" xml:"createEndTime,omitempty"`
	CreateStartTime   *int64  `json:"createStartTime,omitempty" xml:"createStartTime,omitempty"`
	CreatorAccountIds *string `json:"creatorAccountIds,omitempty" xml:"creatorAccountIds,omitempty"`
	Ids               *string `json:"ids,omitempty" xml:"ids,omitempty"`
	MaxResults        *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	NextToken         *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PageOrder         *string `json:"pageOrder,omitempty" xml:"pageOrder,omitempty"`
	PageSort          *string `json:"pageSort,omitempty" xml:"pageSort,omitempty"`
}

func (s ListHostGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListHostGroupsRequest) SetCreateEndTime(v int64) *ListHostGroupsRequest {
	s.CreateEndTime = &v
	return s
}

func (s *ListHostGroupsRequest) SetCreateStartTime(v int64) *ListHostGroupsRequest {
	s.CreateStartTime = &v
	return s
}

func (s *ListHostGroupsRequest) SetCreatorAccountIds(v string) *ListHostGroupsRequest {
	s.CreatorAccountIds = &v
	return s
}

func (s *ListHostGroupsRequest) SetIds(v string) *ListHostGroupsRequest {
	s.Ids = &v
	return s
}

func (s *ListHostGroupsRequest) SetMaxResults(v int64) *ListHostGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHostGroupsRequest) SetName(v string) *ListHostGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListHostGroupsRequest) SetNextToken(v string) *ListHostGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListHostGroupsRequest) SetPageOrder(v string) *ListHostGroupsRequest {
	s.PageOrder = &v
	return s
}

func (s *ListHostGroupsRequest) SetPageSort(v string) *ListHostGroupsRequest {
	s.PageSort = &v
	return s
}

type ListHostGroupsResponseBody struct {
	ErrorCode    *string                                 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                 `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HostGroups   []*ListHostGroupsResponseBodyHostGroups `json:"hostGroups,omitempty" xml:"hostGroups,omitempty" type:"Repeated"`
	NextToken    *string                                 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId    *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount   *int64                                  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListHostGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostGroupsResponseBody) SetErrorCode(v string) *ListHostGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListHostGroupsResponseBody) SetErrorMessage(v string) *ListHostGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListHostGroupsResponseBody) SetHostGroups(v []*ListHostGroupsResponseBodyHostGroups) *ListHostGroupsResponseBody {
	s.HostGroups = v
	return s
}

func (s *ListHostGroupsResponseBody) SetNextToken(v string) *ListHostGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListHostGroupsResponseBody) SetRequestId(v string) *ListHostGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostGroupsResponseBody) SetSuccess(v bool) *ListHostGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListHostGroupsResponseBody) SetTotalCount(v int64) *ListHostGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListHostGroupsResponseBodyHostGroups struct {
	AliyunRegion        *string `json:"aliyunRegion,omitempty" xml:"aliyunRegion,omitempty"`
	CreateTime          *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorAccountId    *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	Description         *string `json:"description,omitempty" xml:"description,omitempty"`
	EcsLabelKey         *string `json:"ecsLabelKey,omitempty" xml:"ecsLabelKey,omitempty"`
	EcsLabelValue       *string `json:"ecsLabelValue,omitempty" xml:"ecsLabelValue,omitempty"`
	EcsType             *string `json:"ecsType,omitempty" xml:"ecsType,omitempty"`
	HostNum             *int64  `json:"hostNum,omitempty" xml:"hostNum,omitempty"`
	Id                  *int64  `json:"id,omitempty" xml:"id,omitempty"`
	ModifierAccountId   *string `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	Name                *string `json:"name,omitempty" xml:"name,omitempty"`
	ServiceConnectionId *int64  `json:"serviceConnectionId,omitempty" xml:"serviceConnectionId,omitempty"`
	Type                *string `json:"type,omitempty" xml:"type,omitempty"`
	UpdateTime          *int64  `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListHostGroupsResponseBodyHostGroups) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsResponseBodyHostGroups) GoString() string {
	return s.String()
}

func (s *ListHostGroupsResponseBodyHostGroups) SetAliyunRegion(v string) *ListHostGroupsResponseBodyHostGroups {
	s.AliyunRegion = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetCreateTime(v int64) *ListHostGroupsResponseBodyHostGroups {
	s.CreateTime = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetCreatorAccountId(v string) *ListHostGroupsResponseBodyHostGroups {
	s.CreatorAccountId = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetDescription(v string) *ListHostGroupsResponseBodyHostGroups {
	s.Description = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetEcsLabelKey(v string) *ListHostGroupsResponseBodyHostGroups {
	s.EcsLabelKey = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetEcsLabelValue(v string) *ListHostGroupsResponseBodyHostGroups {
	s.EcsLabelValue = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetEcsType(v string) *ListHostGroupsResponseBodyHostGroups {
	s.EcsType = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetHostNum(v int64) *ListHostGroupsResponseBodyHostGroups {
	s.HostNum = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetId(v int64) *ListHostGroupsResponseBodyHostGroups {
	s.Id = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetModifierAccountId(v string) *ListHostGroupsResponseBodyHostGroups {
	s.ModifierAccountId = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetName(v string) *ListHostGroupsResponseBodyHostGroups {
	s.Name = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetServiceConnectionId(v int64) *ListHostGroupsResponseBodyHostGroups {
	s.ServiceConnectionId = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetType(v string) *ListHostGroupsResponseBodyHostGroups {
	s.Type = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetUpdateTime(v int64) *ListHostGroupsResponseBodyHostGroups {
	s.UpdateTime = &v
	return s
}

type ListHostGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHostGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHostGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListHostGroupsResponse) SetHeaders(v map[string]*string) *ListHostGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListHostGroupsResponse) SetStatusCode(v int32) *ListHostGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostGroupsResponse) SetBody(v *ListHostGroupsResponseBody) *ListHostGroupsResponse {
	s.Body = v
	return s
}

type ListOrganizationMembersRequest struct {
	ExternUid              *string `json:"externUid,omitempty" xml:"externUid,omitempty"`
	JoinTimeFrom           *int64  `json:"joinTimeFrom,omitempty" xml:"joinTimeFrom,omitempty"`
	JoinTimeTo             *int64  `json:"joinTimeTo,omitempty" xml:"joinTimeTo,omitempty"`
	MaxResults             *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken              *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OrganizationMemberName *string `json:"organizationMemberName,omitempty" xml:"organizationMemberName,omitempty"`
	Provider               *string `json:"provider,omitempty" xml:"provider,omitempty"`
	State                  *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListOrganizationMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationMembersRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersRequest) SetExternUid(v string) *ListOrganizationMembersRequest {
	s.ExternUid = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetJoinTimeFrom(v int64) *ListOrganizationMembersRequest {
	s.JoinTimeFrom = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetJoinTimeTo(v int64) *ListOrganizationMembersRequest {
	s.JoinTimeTo = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetMaxResults(v int64) *ListOrganizationMembersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetNextToken(v string) *ListOrganizationMembersRequest {
	s.NextToken = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetOrganizationMemberName(v string) *ListOrganizationMembersRequest {
	s.OrganizationMemberName = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetProvider(v string) *ListOrganizationMembersRequest {
	s.Provider = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetState(v string) *ListOrganizationMembersRequest {
	s.State = &v
	return s
}

type ListOrganizationMembersResponseBody struct {
	ErrorCode    *string                                       `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                       `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Members      []*ListOrganizationMembersResponseBodyMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	NextToken    *string                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId    *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount   *int64                                        `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListOrganizationMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersResponseBody) SetErrorCode(v string) *ListOrganizationMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetErrorMessage(v string) *ListOrganizationMembersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetMembers(v []*ListOrganizationMembersResponseBodyMembers) *ListOrganizationMembersResponseBody {
	s.Members = v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetNextToken(v string) *ListOrganizationMembersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetRequestId(v string) *ListOrganizationMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetSuccess(v bool) *ListOrganizationMembersResponseBody {
	s.Success = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetTotalCount(v int64) *ListOrganizationMembersResponseBody {
	s.TotalCount = &v
	return s
}

type ListOrganizationMembersResponseBodyMembers struct {
	AccountId              *string                                               `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Birthday               *int64                                                `json:"birthday,omitempty" xml:"birthday,omitempty"`
	DeptLists              []*string                                             `json:"deptLists,omitempty" xml:"deptLists,omitempty" type:"Repeated"`
	Email                  *string                                               `json:"email,omitempty" xml:"email,omitempty"`
	HiredDate              *int64                                                `json:"hiredDate,omitempty" xml:"hiredDate,omitempty"`
	Identities             *ListOrganizationMembersResponseBodyMembersIdentities `json:"identities,omitempty" xml:"identities,omitempty" type:"Struct"`
	JoinTime               *int64                                                `json:"joinTime,omitempty" xml:"joinTime,omitempty"`
	LastVisitTime          *int64                                                `json:"lastVisitTime,omitempty" xml:"lastVisitTime,omitempty"`
	Mobile                 *string                                               `json:"mobile,omitempty" xml:"mobile,omitempty"`
	OrganizationMemberName *string                                               `json:"organizationMemberName,omitempty" xml:"organizationMemberName,omitempty"`
	OrganizationRoleId     *string                                               `json:"organizationRoleId,omitempty" xml:"organizationRoleId,omitempty"`
	OrganizationRoleName   *string                                               `json:"organizationRoleName,omitempty" xml:"organizationRoleName,omitempty"`
	State                  *string                                               `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListOrganizationMembersResponseBodyMembers) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersResponseBodyMembers) SetAccountId(v string) *ListOrganizationMembersResponseBodyMembers {
	s.AccountId = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetBirthday(v int64) *ListOrganizationMembersResponseBodyMembers {
	s.Birthday = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetDeptLists(v []*string) *ListOrganizationMembersResponseBodyMembers {
	s.DeptLists = v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetEmail(v string) *ListOrganizationMembersResponseBodyMembers {
	s.Email = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetHiredDate(v int64) *ListOrganizationMembersResponseBodyMembers {
	s.HiredDate = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetIdentities(v *ListOrganizationMembersResponseBodyMembersIdentities) *ListOrganizationMembersResponseBodyMembers {
	s.Identities = v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetJoinTime(v int64) *ListOrganizationMembersResponseBodyMembers {
	s.JoinTime = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetLastVisitTime(v int64) *ListOrganizationMembersResponseBodyMembers {
	s.LastVisitTime = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetMobile(v string) *ListOrganizationMembersResponseBodyMembers {
	s.Mobile = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetOrganizationMemberName(v string) *ListOrganizationMembersResponseBodyMembers {
	s.OrganizationMemberName = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetOrganizationRoleId(v string) *ListOrganizationMembersResponseBodyMembers {
	s.OrganizationRoleId = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetOrganizationRoleName(v string) *ListOrganizationMembersResponseBodyMembers {
	s.OrganizationRoleName = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetState(v string) *ListOrganizationMembersResponseBodyMembers {
	s.State = &v
	return s
}

type ListOrganizationMembersResponseBodyMembersIdentities struct {
	ExternUid *string `json:"externUid,omitempty" xml:"externUid,omitempty"`
	Provider  *string `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s ListOrganizationMembersResponseBodyMembersIdentities) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationMembersResponseBodyMembersIdentities) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersResponseBodyMembersIdentities) SetExternUid(v string) *ListOrganizationMembersResponseBodyMembersIdentities {
	s.ExternUid = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembersIdentities) SetProvider(v string) *ListOrganizationMembersResponseBodyMembersIdentities {
	s.Provider = &v
	return s
}

type ListOrganizationMembersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOrganizationMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrganizationMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationMembersResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersResponse) SetHeaders(v map[string]*string) *ListOrganizationMembersResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationMembersResponse) SetStatusCode(v int32) *ListOrganizationMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationMembersResponse) SetBody(v *ListOrganizationMembersResponseBody) *ListOrganizationMembersResponse {
	s.Body = v
	return s
}

type ListPipelineGroupPipelinesRequest struct {
	CreateEndTime    *int64  `json:"createEndTime,omitempty" xml:"createEndTime,omitempty"`
	CreateStartTime  *int64  `json:"createStartTime,omitempty" xml:"createStartTime,omitempty"`
	ExecuteEndTime   *int64  `json:"executeEndTime,omitempty" xml:"executeEndTime,omitempty"`
	ExecuteStartTime *int64  `json:"executeStartTime,omitempty" xml:"executeStartTime,omitempty"`
	MaxResults       *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken        *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PipelineName     *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	ResultStatusList *string `json:"resultStatusList,omitempty" xml:"resultStatusList,omitempty"`
}

func (s ListPipelineGroupPipelinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineGroupPipelinesRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineGroupPipelinesRequest) SetCreateEndTime(v int64) *ListPipelineGroupPipelinesRequest {
	s.CreateEndTime = &v
	return s
}

func (s *ListPipelineGroupPipelinesRequest) SetCreateStartTime(v int64) *ListPipelineGroupPipelinesRequest {
	s.CreateStartTime = &v
	return s
}

func (s *ListPipelineGroupPipelinesRequest) SetExecuteEndTime(v int64) *ListPipelineGroupPipelinesRequest {
	s.ExecuteEndTime = &v
	return s
}

func (s *ListPipelineGroupPipelinesRequest) SetExecuteStartTime(v int64) *ListPipelineGroupPipelinesRequest {
	s.ExecuteStartTime = &v
	return s
}

func (s *ListPipelineGroupPipelinesRequest) SetMaxResults(v int64) *ListPipelineGroupPipelinesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPipelineGroupPipelinesRequest) SetNextToken(v string) *ListPipelineGroupPipelinesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPipelineGroupPipelinesRequest) SetPipelineName(v string) *ListPipelineGroupPipelinesRequest {
	s.PipelineName = &v
	return s
}

func (s *ListPipelineGroupPipelinesRequest) SetResultStatusList(v string) *ListPipelineGroupPipelinesRequest {
	s.ResultStatusList = &v
	return s
}

type ListPipelineGroupPipelinesResponseBody struct {
	ErrorCode    *string                                            `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                            `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	NextToken    *string                                            `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Pipelines    []*ListPipelineGroupPipelinesResponseBodyPipelines `json:"pipelines,omitempty" xml:"pipelines,omitempty" type:"Repeated"`
	RequestId    *string                                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                              `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount   *int32                                             `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPipelineGroupPipelinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineGroupPipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineGroupPipelinesResponseBody) SetErrorCode(v string) *ListPipelineGroupPipelinesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBody) SetErrorMessage(v string) *ListPipelineGroupPipelinesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBody) SetNextToken(v string) *ListPipelineGroupPipelinesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBody) SetPipelines(v []*ListPipelineGroupPipelinesResponseBodyPipelines) *ListPipelineGroupPipelinesResponseBody {
	s.Pipelines = v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBody) SetRequestId(v string) *ListPipelineGroupPipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBody) SetSuccess(v bool) *ListPipelineGroupPipelinesResponseBody {
	s.Success = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBody) SetTotalCount(v int32) *ListPipelineGroupPipelinesResponseBody {
	s.TotalCount = &v
	return s
}

type ListPipelineGroupPipelinesResponseBodyPipelines struct {
	CreateTime   *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	PipelineId   *int64  `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
}

func (s ListPipelineGroupPipelinesResponseBodyPipelines) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineGroupPipelinesResponseBodyPipelines) GoString() string {
	return s.String()
}

func (s *ListPipelineGroupPipelinesResponseBodyPipelines) SetCreateTime(v int64) *ListPipelineGroupPipelinesResponseBodyPipelines {
	s.CreateTime = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBodyPipelines) SetPipelineId(v int64) *ListPipelineGroupPipelinesResponseBodyPipelines {
	s.PipelineId = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBodyPipelines) SetPipelineName(v string) *ListPipelineGroupPipelinesResponseBodyPipelines {
	s.PipelineName = &v
	return s
}

type ListPipelineGroupPipelinesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPipelineGroupPipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPipelineGroupPipelinesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineGroupPipelinesResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineGroupPipelinesResponse) SetHeaders(v map[string]*string) *ListPipelineGroupPipelinesResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineGroupPipelinesResponse) SetStatusCode(v int32) *ListPipelineGroupPipelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponse) SetBody(v *ListPipelineGroupPipelinesResponseBody) *ListPipelineGroupPipelinesResponse {
	s.Body = v
	return s
}

type ListPipelineGroupsRequest struct {
	MaxResults *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListPipelineGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineGroupsRequest) SetMaxResults(v int64) *ListPipelineGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPipelineGroupsRequest) SetNextToken(v string) *ListPipelineGroupsRequest {
	s.NextToken = &v
	return s
}

type ListPipelineGroupsResponseBody struct {
	ErrorCode      *string                                         `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage   *string                                         `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	NextToken      *string                                         `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PipelineGroups []*ListPipelineGroupsResponseBodyPipelineGroups `json:"pipelineGroups,omitempty" xml:"pipelineGroups,omitempty" type:"Repeated"`
	RequestId      *string                                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                                           `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount     *int32                                          `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPipelineGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineGroupsResponseBody) SetErrorCode(v string) *ListPipelineGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPipelineGroupsResponseBody) SetErrorMessage(v string) *ListPipelineGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListPipelineGroupsResponseBody) SetNextToken(v string) *ListPipelineGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPipelineGroupsResponseBody) SetPipelineGroups(v []*ListPipelineGroupsResponseBodyPipelineGroups) *ListPipelineGroupsResponseBody {
	s.PipelineGroups = v
	return s
}

func (s *ListPipelineGroupsResponseBody) SetRequestId(v string) *ListPipelineGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineGroupsResponseBody) SetSuccess(v bool) *ListPipelineGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListPipelineGroupsResponseBody) SetTotalCount(v int32) *ListPipelineGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListPipelineGroupsResponseBodyPipelineGroups struct {
	CreateTime *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Id         *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListPipelineGroupsResponseBodyPipelineGroups) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineGroupsResponseBodyPipelineGroups) GoString() string {
	return s.String()
}

func (s *ListPipelineGroupsResponseBodyPipelineGroups) SetCreateTime(v int64) *ListPipelineGroupsResponseBodyPipelineGroups {
	s.CreateTime = &v
	return s
}

func (s *ListPipelineGroupsResponseBodyPipelineGroups) SetId(v int64) *ListPipelineGroupsResponseBodyPipelineGroups {
	s.Id = &v
	return s
}

func (s *ListPipelineGroupsResponseBodyPipelineGroups) SetName(v string) *ListPipelineGroupsResponseBodyPipelineGroups {
	s.Name = &v
	return s
}

type ListPipelineGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPipelineGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPipelineGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineGroupsResponse) SetHeaders(v map[string]*string) *ListPipelineGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineGroupsResponse) SetStatusCode(v int32) *ListPipelineGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineGroupsResponse) SetBody(v *ListPipelineGroupsResponseBody) *ListPipelineGroupsResponse {
	s.Body = v
	return s
}

type ListPipelineJobHistorysRequest struct {
	Category   *string `json:"category,omitempty" xml:"category,omitempty"`
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	MaxResults *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListPipelineJobHistorysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineJobHistorysRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineJobHistorysRequest) SetCategory(v string) *ListPipelineJobHistorysRequest {
	s.Category = &v
	return s
}

func (s *ListPipelineJobHistorysRequest) SetIdentifier(v string) *ListPipelineJobHistorysRequest {
	s.Identifier = &v
	return s
}

func (s *ListPipelineJobHistorysRequest) SetMaxResults(v int64) *ListPipelineJobHistorysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPipelineJobHistorysRequest) SetNextToken(v string) *ListPipelineJobHistorysRequest {
	s.NextToken = &v
	return s
}

type ListPipelineJobHistorysResponseBody struct {
	ErrorCode    *string                                    `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                    `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Jobs         []*ListPipelineJobHistorysResponseBodyJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	NextToken    *string                                    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId    *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount   *int32                                     `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPipelineJobHistorysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineJobHistorysResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineJobHistorysResponseBody) SetErrorCode(v string) *ListPipelineJobHistorysResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBody) SetErrorMessage(v string) *ListPipelineJobHistorysResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBody) SetJobs(v []*ListPipelineJobHistorysResponseBodyJobs) *ListPipelineJobHistorysResponseBody {
	s.Jobs = v
	return s
}

func (s *ListPipelineJobHistorysResponseBody) SetNextToken(v string) *ListPipelineJobHistorysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBody) SetRequestId(v string) *ListPipelineJobHistorysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBody) SetSuccess(v bool) *ListPipelineJobHistorysResponseBody {
	s.Success = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBody) SetTotalCount(v int32) *ListPipelineJobHistorysResponseBody {
	s.TotalCount = &v
	return s
}

type ListPipelineJobHistorysResponseBodyJobs struct {
	ExecuteNumber     *int32  `json:"executeNumber,omitempty" xml:"executeNumber,omitempty"`
	Identifier        *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	JobId             *int64  `json:"jobId,omitempty" xml:"jobId,omitempty"`
	JobName           *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	OperatorAccountId *string `json:"operatorAccountId,omitempty" xml:"operatorAccountId,omitempty"`
	PipelineId        *int64  `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	PipelineRunId     *int64  `json:"pipelineRunId,omitempty" xml:"pipelineRunId,omitempty"`
	Sources           *string `json:"sources,omitempty" xml:"sources,omitempty"`
	Status            *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListPipelineJobHistorysResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineJobHistorysResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetExecuteNumber(v int32) *ListPipelineJobHistorysResponseBodyJobs {
	s.ExecuteNumber = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetIdentifier(v string) *ListPipelineJobHistorysResponseBodyJobs {
	s.Identifier = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetJobId(v int64) *ListPipelineJobHistorysResponseBodyJobs {
	s.JobId = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetJobName(v string) *ListPipelineJobHistorysResponseBodyJobs {
	s.JobName = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetOperatorAccountId(v string) *ListPipelineJobHistorysResponseBodyJobs {
	s.OperatorAccountId = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetPipelineId(v int64) *ListPipelineJobHistorysResponseBodyJobs {
	s.PipelineId = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetPipelineRunId(v int64) *ListPipelineJobHistorysResponseBodyJobs {
	s.PipelineRunId = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetSources(v string) *ListPipelineJobHistorysResponseBodyJobs {
	s.Sources = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetStatus(v string) *ListPipelineJobHistorysResponseBodyJobs {
	s.Status = &v
	return s
}

type ListPipelineJobHistorysResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPipelineJobHistorysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPipelineJobHistorysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineJobHistorysResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineJobHistorysResponse) SetHeaders(v map[string]*string) *ListPipelineJobHistorysResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineJobHistorysResponse) SetStatusCode(v int32) *ListPipelineJobHistorysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineJobHistorysResponse) SetBody(v *ListPipelineJobHistorysResponseBody) *ListPipelineJobHistorysResponse {
	s.Body = v
	return s
}

type ListPipelineJobsRequest struct {
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
}

func (s ListPipelineJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineJobsRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineJobsRequest) SetCategory(v string) *ListPipelineJobsRequest {
	s.Category = &v
	return s
}

type ListPipelineJobsResponseBody struct {
	ErrorCode    *string                             `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                             `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Jobs         []*ListPipelineJobsResponseBodyJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	RequestId    *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListPipelineJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineJobsResponseBody) SetErrorCode(v string) *ListPipelineJobsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPipelineJobsResponseBody) SetErrorMessage(v string) *ListPipelineJobsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListPipelineJobsResponseBody) SetJobs(v []*ListPipelineJobsResponseBodyJobs) *ListPipelineJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListPipelineJobsResponseBody) SetRequestId(v string) *ListPipelineJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineJobsResponseBody) SetSuccess(v bool) *ListPipelineJobsResponseBody {
	s.Success = &v
	return s
}

type ListPipelineJobsResponseBodyJobs struct {
	Identifier    *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	JobName       *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	LastJobId     *int64  `json:"lastJobId,omitempty" xml:"lastJobId,omitempty"`
	LastJobParams *string `json:"lastJobParams,omitempty" xml:"lastJobParams,omitempty"`
}

func (s ListPipelineJobsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListPipelineJobsResponseBodyJobs) SetIdentifier(v string) *ListPipelineJobsResponseBodyJobs {
	s.Identifier = &v
	return s
}

func (s *ListPipelineJobsResponseBodyJobs) SetJobName(v string) *ListPipelineJobsResponseBodyJobs {
	s.JobName = &v
	return s
}

func (s *ListPipelineJobsResponseBodyJobs) SetLastJobId(v int64) *ListPipelineJobsResponseBodyJobs {
	s.LastJobId = &v
	return s
}

func (s *ListPipelineJobsResponseBodyJobs) SetLastJobParams(v string) *ListPipelineJobsResponseBodyJobs {
	s.LastJobParams = &v
	return s
}

type ListPipelineJobsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPipelineJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPipelineJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineJobsResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineJobsResponse) SetHeaders(v map[string]*string) *ListPipelineJobsResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineJobsResponse) SetStatusCode(v int32) *ListPipelineJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineJobsResponse) SetBody(v *ListPipelineJobsResponseBody) *ListPipelineJobsResponse {
	s.Body = v
	return s
}

type ListPipelineRunsRequest struct {
	EndTime     *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	MaxResults  *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken   *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	StartTime   *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
	TriggerMode *int32  `json:"triggerMode,omitempty" xml:"triggerMode,omitempty"`
}

func (s ListPipelineRunsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunsRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsRequest) SetEndTime(v int64) *ListPipelineRunsRequest {
	s.EndTime = &v
	return s
}

func (s *ListPipelineRunsRequest) SetMaxResults(v int64) *ListPipelineRunsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPipelineRunsRequest) SetNextToken(v string) *ListPipelineRunsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPipelineRunsRequest) SetStartTime(v int64) *ListPipelineRunsRequest {
	s.StartTime = &v
	return s
}

func (s *ListPipelineRunsRequest) SetStatus(v string) *ListPipelineRunsRequest {
	s.Status = &v
	return s
}

func (s *ListPipelineRunsRequest) SetTriggerMode(v int32) *ListPipelineRunsRequest {
	s.TriggerMode = &v
	return s
}

type ListPipelineRunsResponseBody struct {
	ErrorCode    *string                                     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                     `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	NextToken    *string                                     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PipelineRuns []*ListPipelineRunsResponseBodyPipelineRuns `json:"pipelineRuns,omitempty" xml:"pipelineRuns,omitempty" type:"Repeated"`
	RequestId    *string                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount   *int64                                      `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPipelineRunsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsResponseBody) SetErrorCode(v string) *ListPipelineRunsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPipelineRunsResponseBody) SetErrorMessage(v string) *ListPipelineRunsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListPipelineRunsResponseBody) SetNextToken(v string) *ListPipelineRunsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPipelineRunsResponseBody) SetPipelineRuns(v []*ListPipelineRunsResponseBodyPipelineRuns) *ListPipelineRunsResponseBody {
	s.PipelineRuns = v
	return s
}

func (s *ListPipelineRunsResponseBody) SetRequestId(v string) *ListPipelineRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineRunsResponseBody) SetSuccess(v bool) *ListPipelineRunsResponseBody {
	s.Success = &v
	return s
}

func (s *ListPipelineRunsResponseBody) SetTotalCount(v int64) *ListPipelineRunsResponseBody {
	s.TotalCount = &v
	return s
}

type ListPipelineRunsResponseBodyPipelineRuns struct {
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	EndTime          *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	PipelineId       *int64  `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	PipelineRunId    *int64  `json:"pipelineRunId,omitempty" xml:"pipelineRunId,omitempty"`
	StartTime        *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status           *string `json:"status,omitempty" xml:"status,omitempty"`
	TriggerMode      *int64  `json:"triggerMode,omitempty" xml:"triggerMode,omitempty"`
}

func (s ListPipelineRunsResponseBodyPipelineRuns) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunsResponseBodyPipelineRuns) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetCreatorAccountId(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.CreatorAccountId = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetEndTime(v int64) *ListPipelineRunsResponseBodyPipelineRuns {
	s.EndTime = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetPipelineId(v int64) *ListPipelineRunsResponseBodyPipelineRuns {
	s.PipelineId = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetPipelineRunId(v int64) *ListPipelineRunsResponseBodyPipelineRuns {
	s.PipelineRunId = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetStartTime(v int64) *ListPipelineRunsResponseBodyPipelineRuns {
	s.StartTime = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetStatus(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.Status = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetTriggerMode(v int64) *ListPipelineRunsResponseBodyPipelineRuns {
	s.TriggerMode = &v
	return s
}

type ListPipelineRunsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPipelineRunsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPipelineRunsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunsResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsResponse) SetHeaders(v map[string]*string) *ListPipelineRunsResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineRunsResponse) SetStatusCode(v int32) *ListPipelineRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineRunsResponse) SetBody(v *ListPipelineRunsResponseBody) *ListPipelineRunsResponse {
	s.Body = v
	return s
}

type ListPipelinesRequest struct {
	CreateEndTime     *int64  `json:"createEndTime,omitempty" xml:"createEndTime,omitempty"`
	CreateStartTime   *int64  `json:"createStartTime,omitempty" xml:"createStartTime,omitempty"`
	CreatorAccountIds *string `json:"creatorAccountIds,omitempty" xml:"creatorAccountIds,omitempty"`
	ExecuteAccountIds *string `json:"executeAccountIds,omitempty" xml:"executeAccountIds,omitempty"`
	ExecuteEndTime    *int64  `json:"executeEndTime,omitempty" xml:"executeEndTime,omitempty"`
	ExecuteStartTime  *int64  `json:"executeStartTime,omitempty" xml:"executeStartTime,omitempty"`
	MaxResults        *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken         *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PipelineName      *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	StatusList        *string `json:"statusList,omitempty" xml:"statusList,omitempty"`
}

func (s ListPipelinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelinesRequest) GoString() string {
	return s.String()
}

func (s *ListPipelinesRequest) SetCreateEndTime(v int64) *ListPipelinesRequest {
	s.CreateEndTime = &v
	return s
}

func (s *ListPipelinesRequest) SetCreateStartTime(v int64) *ListPipelinesRequest {
	s.CreateStartTime = &v
	return s
}

func (s *ListPipelinesRequest) SetCreatorAccountIds(v string) *ListPipelinesRequest {
	s.CreatorAccountIds = &v
	return s
}

func (s *ListPipelinesRequest) SetExecuteAccountIds(v string) *ListPipelinesRequest {
	s.ExecuteAccountIds = &v
	return s
}

func (s *ListPipelinesRequest) SetExecuteEndTime(v int64) *ListPipelinesRequest {
	s.ExecuteEndTime = &v
	return s
}

func (s *ListPipelinesRequest) SetExecuteStartTime(v int64) *ListPipelinesRequest {
	s.ExecuteStartTime = &v
	return s
}

func (s *ListPipelinesRequest) SetMaxResults(v int64) *ListPipelinesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPipelinesRequest) SetNextToken(v string) *ListPipelinesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPipelinesRequest) SetPipelineName(v string) *ListPipelinesRequest {
	s.PipelineName = &v
	return s
}

func (s *ListPipelinesRequest) SetStatusList(v string) *ListPipelinesRequest {
	s.StatusList = &v
	return s
}

type ListPipelinesResponseBody struct {
	ErrorCode    *string                               `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                               `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	NextToken    *string                               `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Pipelines    []*ListPipelinesResponseBodyPipelines `json:"pipelines,omitempty" xml:"pipelines,omitempty" type:"Repeated"`
	RequestId    *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount   *int64                                `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPipelinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBody) SetErrorCode(v string) *ListPipelinesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPipelinesResponseBody) SetErrorMessage(v string) *ListPipelinesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListPipelinesResponseBody) SetNextToken(v string) *ListPipelinesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPipelinesResponseBody) SetPipelines(v []*ListPipelinesResponseBodyPipelines) *ListPipelinesResponseBody {
	s.Pipelines = v
	return s
}

func (s *ListPipelinesResponseBody) SetRequestId(v string) *ListPipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelinesResponseBody) SetSuccess(v bool) *ListPipelinesResponseBody {
	s.Success = &v
	return s
}

func (s *ListPipelinesResponseBody) SetTotalCount(v int64) *ListPipelinesResponseBody {
	s.TotalCount = &v
	return s
}

type ListPipelinesResponseBodyPipelines struct {
	CreateTime       *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	GroupId          *int64  `json:"groupId,omitempty" xml:"groupId,omitempty"`
	PipelineId       *int64  `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	PipelineName     *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
}

func (s ListPipelinesResponseBodyPipelines) String() string {
	return tea.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelines) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelines) SetCreateTime(v int64) *ListPipelinesResponseBodyPipelines {
	s.CreateTime = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetCreatorAccountId(v string) *ListPipelinesResponseBodyPipelines {
	s.CreatorAccountId = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetGroupId(v int64) *ListPipelinesResponseBodyPipelines {
	s.GroupId = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetPipelineId(v int64) *ListPipelinesResponseBodyPipelines {
	s.PipelineId = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetPipelineName(v string) *ListPipelinesResponseBodyPipelines {
	s.PipelineName = &v
	return s
}

type ListPipelinesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPipelinesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelinesResponse) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponse) SetHeaders(v map[string]*string) *ListPipelinesResponse {
	s.Headers = v
	return s
}

func (s *ListPipelinesResponse) SetStatusCode(v int32) *ListPipelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelinesResponse) SetBody(v *ListPipelinesResponseBody) *ListPipelinesResponse {
	s.Body = v
	return s
}

type ListProjectMembersRequest struct {
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s ListProjectMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectMembersRequest) GoString() string {
	return s.String()
}

func (s *ListProjectMembersRequest) SetTargetType(v string) *ListProjectMembersRequest {
	s.TargetType = &v
	return s
}

type ListProjectMembersResponseBody struct {
	ErrorCode *string                                  `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                  `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Members   []*ListProjectMembersResponseBodyMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListProjectMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBody) SetErrorCode(v string) *ListProjectMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProjectMembersResponseBody) SetErrorMsg(v string) *ListProjectMembersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListProjectMembersResponseBody) SetMembers(v []*ListProjectMembersResponseBodyMembers) *ListProjectMembersResponseBody {
	s.Members = v
	return s
}

func (s *ListProjectMembersResponseBody) SetRequestId(v string) *ListProjectMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectMembersResponseBody) SetSuccess(v bool) *ListProjectMembersResponseBody {
	s.Success = &v
	return s
}

type ListProjectMembersResponseBodyMembers struct {
	Account              *string                                                    `json:"account,omitempty" xml:"account,omitempty"`
	Avatar               *string                                                    `json:"avatar,omitempty" xml:"avatar,omitempty"`
	DingTalkId           *string                                                    `json:"dingTalkId,omitempty" xml:"dingTalkId,omitempty"`
	DisplayName          *string                                                    `json:"displayName,omitempty" xml:"displayName,omitempty"`
	DisplayNickName      *string                                                    `json:"displayNickName,omitempty" xml:"displayNickName,omitempty"`
	DisplayRealName      *string                                                    `json:"displayRealName,omitempty" xml:"displayRealName,omitempty"`
	Division             *ListProjectMembersResponseBodyMembersDivision             `json:"division,omitempty" xml:"division,omitempty" type:"Struct"`
	Email                *string                                                    `json:"email,omitempty" xml:"email,omitempty"`
	Gender               *string                                                    `json:"gender,omitempty" xml:"gender,omitempty"`
	Identifier           *string                                                    `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Mobile               *string                                                    `json:"mobile,omitempty" xml:"mobile,omitempty"`
	NameEn               *string                                                    `json:"nameEn,omitempty" xml:"nameEn,omitempty"`
	NickName             *string                                                    `json:"nickName,omitempty" xml:"nickName,omitempty"`
	NickNamePinyin       *string                                                    `json:"nickNamePinyin,omitempty" xml:"nickNamePinyin,omitempty"`
	OrganizationUserInfo *ListProjectMembersResponseBodyMembersOrganizationUserInfo `json:"organizationUserInfo,omitempty" xml:"organizationUserInfo,omitempty" type:"Struct"`
	RealName             *string                                                    `json:"realName,omitempty" xml:"realName,omitempty"`
	RealNamePinyin       *string                                                    `json:"realNamePinyin,omitempty" xml:"realNamePinyin,omitempty"`
	Stamp                *string                                                    `json:"stamp,omitempty" xml:"stamp,omitempty"`
	TbRoleId             *string                                                    `json:"tbRoleId,omitempty" xml:"tbRoleId,omitempty"`
}

func (s ListProjectMembersResponseBodyMembers) String() string {
	return tea.Prettify(s)
}

func (s ListProjectMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBodyMembers) SetAccount(v string) *ListProjectMembersResponseBodyMembers {
	s.Account = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetAvatar(v string) *ListProjectMembersResponseBodyMembers {
	s.Avatar = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetDingTalkId(v string) *ListProjectMembersResponseBodyMembers {
	s.DingTalkId = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetDisplayName(v string) *ListProjectMembersResponseBodyMembers {
	s.DisplayName = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetDisplayNickName(v string) *ListProjectMembersResponseBodyMembers {
	s.DisplayNickName = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetDisplayRealName(v string) *ListProjectMembersResponseBodyMembers {
	s.DisplayRealName = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetDivision(v *ListProjectMembersResponseBodyMembersDivision) *ListProjectMembersResponseBodyMembers {
	s.Division = v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetEmail(v string) *ListProjectMembersResponseBodyMembers {
	s.Email = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetGender(v string) *ListProjectMembersResponseBodyMembers {
	s.Gender = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetIdentifier(v string) *ListProjectMembersResponseBodyMembers {
	s.Identifier = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetMobile(v string) *ListProjectMembersResponseBodyMembers {
	s.Mobile = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetNameEn(v string) *ListProjectMembersResponseBodyMembers {
	s.NameEn = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetNickName(v string) *ListProjectMembersResponseBodyMembers {
	s.NickName = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetNickNamePinyin(v string) *ListProjectMembersResponseBodyMembers {
	s.NickNamePinyin = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetOrganizationUserInfo(v *ListProjectMembersResponseBodyMembersOrganizationUserInfo) *ListProjectMembersResponseBodyMembers {
	s.OrganizationUserInfo = v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetRealName(v string) *ListProjectMembersResponseBodyMembers {
	s.RealName = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetRealNamePinyin(v string) *ListProjectMembersResponseBodyMembers {
	s.RealNamePinyin = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetStamp(v string) *ListProjectMembersResponseBodyMembers {
	s.Stamp = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetTbRoleId(v string) *ListProjectMembersResponseBodyMembers {
	s.TbRoleId = &v
	return s
}

type ListProjectMembersResponseBodyMembersDivision struct {
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
}

func (s ListProjectMembersResponseBodyMembersDivision) String() string {
	return tea.Prettify(s)
}

func (s ListProjectMembersResponseBodyMembersDivision) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBodyMembersDivision) SetIdentifier(v string) *ListProjectMembersResponseBodyMembersDivision {
	s.Identifier = &v
	return s
}

type ListProjectMembersResponseBodyMembersOrganizationUserInfo struct {
	OrganizationIdentifier *string `json:"organizationIdentifier,omitempty" xml:"organizationIdentifier,omitempty"`
}

func (s ListProjectMembersResponseBodyMembersOrganizationUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListProjectMembersResponseBodyMembersOrganizationUserInfo) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBodyMembersOrganizationUserInfo) SetOrganizationIdentifier(v string) *ListProjectMembersResponseBodyMembersOrganizationUserInfo {
	s.OrganizationIdentifier = &v
	return s
}

type ListProjectMembersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProjectMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectMembersResponse) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponse) SetHeaders(v map[string]*string) *ListProjectMembersResponse {
	s.Headers = v
	return s
}

func (s *ListProjectMembersResponse) SetStatusCode(v int32) *ListProjectMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectMembersResponse) SetBody(v *ListProjectMembersResponseBody) *ListProjectMembersResponse {
	s.Body = v
	return s
}

type ListProjectTemplatesRequest struct {
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
}

func (s ListProjectTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListProjectTemplatesRequest) SetCategory(v string) *ListProjectTemplatesRequest {
	s.Category = &v
	return s
}

type ListProjectTemplatesResponseBody struct {
	ErrorCode *string                                      `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                      `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId *string                                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
	Templates []*ListProjectTemplatesResponseBodyTemplates `json:"templates,omitempty" xml:"templates,omitempty" type:"Repeated"`
}

func (s ListProjectTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectTemplatesResponseBody) SetErrorCode(v string) *ListProjectTemplatesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProjectTemplatesResponseBody) SetErrorMsg(v string) *ListProjectTemplatesResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListProjectTemplatesResponseBody) SetRequestId(v string) *ListProjectTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectTemplatesResponseBody) SetSuccess(v bool) *ListProjectTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectTemplatesResponseBody) SetTemplates(v []*ListProjectTemplatesResponseBodyTemplates) *ListProjectTemplatesResponseBody {
	s.Templates = v
	return s
}

type ListProjectTemplatesResponseBodyTemplates struct {
	CopyFrom         *string `json:"copyFrom,omitempty" xml:"copyFrom,omitempty"`
	Creator          *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Description      *string `json:"description,omitempty" xml:"description,omitempty"`
	GmtCreate        *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified      *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Icon             *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Identifier       *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Modifier         *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	NameEn           *string `json:"nameEn,omitempty" xml:"nameEn,omitempty"`
	ResourceCategory *string `json:"resourceCategory,omitempty" xml:"resourceCategory,omitempty"`
	ResourceType     *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	SpaceIdentifier  *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	SpaceType        *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	Type             *int64  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListProjectTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListProjectTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetCopyFrom(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.CopyFrom = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetCreator(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.Creator = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetDescription(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetGmtCreate(v int64) *ListProjectTemplatesResponseBodyTemplates {
	s.GmtCreate = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetGmtModified(v int64) *ListProjectTemplatesResponseBodyTemplates {
	s.GmtModified = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetIcon(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.Icon = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetIdentifier(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.Identifier = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetModifier(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.Modifier = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetName(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetNameEn(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.NameEn = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetResourceCategory(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.ResourceCategory = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetResourceType(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.ResourceType = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetSpaceIdentifier(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.SpaceIdentifier = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetSpaceType(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.SpaceType = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetType(v int64) *ListProjectTemplatesResponseBodyTemplates {
	s.Type = &v
	return s
}

type ListProjectTemplatesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProjectTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListProjectTemplatesResponse) SetHeaders(v map[string]*string) *ListProjectTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListProjectTemplatesResponse) SetStatusCode(v int32) *ListProjectTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectTemplatesResponse) SetBody(v *ListProjectTemplatesResponseBody) *ListProjectTemplatesResponse {
	s.Body = v
	return s
}

type ListProjectWorkitemTypesRequest struct {
	Category  *string `json:"category,omitempty" xml:"category,omitempty"`
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s ListProjectWorkitemTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectWorkitemTypesRequest) GoString() string {
	return s.String()
}

func (s *ListProjectWorkitemTypesRequest) SetCategory(v string) *ListProjectWorkitemTypesRequest {
	s.Category = &v
	return s
}

func (s *ListProjectWorkitemTypesRequest) SetSpaceType(v string) *ListProjectWorkitemTypesRequest {
	s.SpaceType = &v
	return s
}

type ListProjectWorkitemTypesResponseBody struct {
	ErrorCode     *string                                              `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage  *string                                              `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId     *string                                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success       *bool                                                `json:"success,omitempty" xml:"success,omitempty"`
	WorkitemTypes []*ListProjectWorkitemTypesResponseBodyWorkitemTypes `json:"workitemTypes,omitempty" xml:"workitemTypes,omitempty" type:"Repeated"`
}

func (s ListProjectWorkitemTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectWorkitemTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectWorkitemTypesResponseBody) SetErrorCode(v string) *ListProjectWorkitemTypesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBody) SetErrorMessage(v string) *ListProjectWorkitemTypesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBody) SetRequestId(v string) *ListProjectWorkitemTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBody) SetSuccess(v bool) *ListProjectWorkitemTypesResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBody) SetWorkitemTypes(v []*ListProjectWorkitemTypesResponseBodyWorkitemTypes) *ListProjectWorkitemTypesResponseBody {
	s.WorkitemTypes = v
	return s
}

type ListProjectWorkitemTypesResponseBodyWorkitemTypes struct {
	AddUser            *string `json:"addUser,omitempty" xml:"addUser,omitempty"`
	CategoryIdentifier *string `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	Creator            *string `json:"creator,omitempty" xml:"creator,omitempty"`
	DefaultType        *bool   `json:"defaultType,omitempty" xml:"defaultType,omitempty"`
	Description        *string `json:"description,omitempty" xml:"description,omitempty"`
	Enable             *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
	GmtAdd             *int64  `json:"gmtAdd,omitempty" xml:"gmtAdd,omitempty"`
	GmtCreate          *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	Identifier         *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	NameEn             *string `json:"nameEn,omitempty" xml:"nameEn,omitempty"`
	SystemDefault      *bool   `json:"systemDefault,omitempty" xml:"systemDefault,omitempty"`
}

func (s ListProjectWorkitemTypesResponseBodyWorkitemTypes) String() string {
	return tea.Prettify(s)
}

func (s ListProjectWorkitemTypesResponseBodyWorkitemTypes) GoString() string {
	return s.String()
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetAddUser(v string) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.AddUser = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetCategoryIdentifier(v string) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.CategoryIdentifier = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetCreator(v string) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.Creator = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetDefaultType(v bool) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.DefaultType = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetDescription(v string) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.Description = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetEnable(v bool) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.Enable = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetGmtAdd(v int64) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.GmtAdd = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetGmtCreate(v int64) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.GmtCreate = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetIdentifier(v string) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.Identifier = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetName(v string) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.Name = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetNameEn(v string) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.NameEn = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetSystemDefault(v bool) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.SystemDefault = &v
	return s
}

type ListProjectWorkitemTypesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProjectWorkitemTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectWorkitemTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectWorkitemTypesResponse) GoString() string {
	return s.String()
}

func (s *ListProjectWorkitemTypesResponse) SetHeaders(v map[string]*string) *ListProjectWorkitemTypesResponse {
	s.Headers = v
	return s
}

func (s *ListProjectWorkitemTypesResponse) SetStatusCode(v int32) *ListProjectWorkitemTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectWorkitemTypesResponse) SetBody(v *ListProjectWorkitemTypesResponseBody) *ListProjectWorkitemTypesResponse {
	s.Body = v
	return s
}

type ListProjectsRequest struct {
	Category        *string `json:"category,omitempty" xml:"category,omitempty"`
	Conditions      *string `json:"conditions,omitempty" xml:"conditions,omitempty"`
	ExtraConditions *string `json:"extraConditions,omitempty" xml:"extraConditions,omitempty"`
	MaxResults      *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken       *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Scope           *string `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) SetCategory(v string) *ListProjectsRequest {
	s.Category = &v
	return s
}

func (s *ListProjectsRequest) SetConditions(v string) *ListProjectsRequest {
	s.Conditions = &v
	return s
}

func (s *ListProjectsRequest) SetExtraConditions(v string) *ListProjectsRequest {
	s.ExtraConditions = &v
	return s
}

func (s *ListProjectsRequest) SetMaxResults(v int64) *ListProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsRequest) SetNextToken(v string) *ListProjectsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProjectsRequest) SetScope(v string) *ListProjectsRequest {
	s.Scope = &v
	return s
}

type ListProjectsResponseBody struct {
	ErrorCode  *string                             `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg   *string                             `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	MaxResults *int64                              `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Projects   []*ListProjectsResponseBodyProjects `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
	RequestId  *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success    *bool                               `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount *int64                              `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetErrorCode(v string) *ListProjectsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProjectsResponseBody) SetErrorMsg(v string) *ListProjectsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListProjectsResponseBody) SetMaxResults(v int64) *ListProjectsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsResponseBody) SetNextToken(v string) *ListProjectsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProjectsResponseBody) SetProjects(v []*ListProjectsResponseBodyProjects) *ListProjectsResponseBody {
	s.Projects = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) SetSuccess(v bool) *ListProjectsResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectsResponseBody) SetTotalCount(v int64) *ListProjectsResponseBody {
	s.TotalCount = &v
	return s
}

type ListProjectsResponseBodyProjects struct {
	CategoryIdentifier    *string `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	Creator               *string `json:"creator,omitempty" xml:"creator,omitempty"`
	CustomCode            *string `json:"customCode,omitempty" xml:"customCode,omitempty"`
	DeleteTime            *int64  `json:"deleteTime,omitempty" xml:"deleteTime,omitempty"`
	Description           *string `json:"description,omitempty" xml:"description,omitempty"`
	GmtCreate             *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	Icon                  *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Identifier            *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	LogicalStatus         *string `json:"logicalStatus,omitempty" xml:"logicalStatus,omitempty"`
	Name                  *string `json:"name,omitempty" xml:"name,omitempty"`
	Scope                 *string `json:"scope,omitempty" xml:"scope,omitempty"`
	StatusStageIdentifier *string `json:"statusStageIdentifier,omitempty" xml:"statusStageIdentifier,omitempty"`
	TypeIdentifier        *string `json:"typeIdentifier,omitempty" xml:"typeIdentifier,omitempty"`
}

func (s ListProjectsResponseBodyProjects) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyProjects) SetCategoryIdentifier(v string) *ListProjectsResponseBodyProjects {
	s.CategoryIdentifier = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetCreator(v string) *ListProjectsResponseBodyProjects {
	s.Creator = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetCustomCode(v string) *ListProjectsResponseBodyProjects {
	s.CustomCode = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDeleteTime(v int64) *ListProjectsResponseBodyProjects {
	s.DeleteTime = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDescription(v string) *ListProjectsResponseBodyProjects {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetGmtCreate(v int64) *ListProjectsResponseBodyProjects {
	s.GmtCreate = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetIcon(v string) *ListProjectsResponseBodyProjects {
	s.Icon = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetIdentifier(v string) *ListProjectsResponseBodyProjects {
	s.Identifier = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetLogicalStatus(v string) *ListProjectsResponseBodyProjects {
	s.LogicalStatus = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetName(v string) *ListProjectsResponseBodyProjects {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetScope(v string) *ListProjectsResponseBodyProjects {
	s.Scope = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetStatusStageIdentifier(v string) *ListProjectsResponseBodyProjects {
	s.StatusStageIdentifier = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetTypeIdentifier(v string) *ListProjectsResponseBodyProjects {
	s.TypeIdentifier = &v
	return s
}

type ListProjectsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsResponse) SetHeaders(v map[string]*string) *ListProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsResponse) SetStatusCode(v int32) *ListProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

type ListProtectedBranchesRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s ListProtectedBranchesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProtectedBranchesRequest) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesRequest) SetAccessToken(v string) *ListProtectedBranchesRequest {
	s.AccessToken = &v
	return s
}

func (s *ListProtectedBranchesRequest) SetOrganizationId(v string) *ListProtectedBranchesRequest {
	s.OrganizationId = &v
	return s
}

type ListProtectedBranchesResponseBody struct {
	ErrorCode    *string                                    `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                    `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       []*ListProtectedBranchesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success      *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListProtectedBranchesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProtectedBranchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBody) SetErrorCode(v string) *ListProtectedBranchesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProtectedBranchesResponseBody) SetErrorMessage(v string) *ListProtectedBranchesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListProtectedBranchesResponseBody) SetRequestId(v string) *ListProtectedBranchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProtectedBranchesResponseBody) SetResult(v []*ListProtectedBranchesResponseBodyResult) *ListProtectedBranchesResponseBody {
	s.Result = v
	return s
}

func (s *ListProtectedBranchesResponseBody) SetSuccess(v bool) *ListProtectedBranchesResponseBody {
	s.Success = &v
	return s
}

type ListProtectedBranchesResponseBodyResult struct {
	AllowMergeRoles     []*int32                                                    `json:"allowMergeRoles,omitempty" xml:"allowMergeRoles,omitempty" type:"Repeated"`
	AllowMergeUserIds   []*int64                                                    `json:"allowMergeUserIds,omitempty" xml:"allowMergeUserIds,omitempty" type:"Repeated"`
	AllowMergeUsers     []*ListProtectedBranchesResponseBodyResultAllowMergeUsers   `json:"allowMergeUsers,omitempty" xml:"allowMergeUsers,omitempty" type:"Repeated"`
	AllowPushRoles      []*int32                                                    `json:"allowPushRoles,omitempty" xml:"allowPushRoles,omitempty" type:"Repeated"`
	AllowPushUserIds    []*int64                                                    `json:"allowPushUserIds,omitempty" xml:"allowPushUserIds,omitempty" type:"Repeated"`
	AllowPushUsers      []*ListProtectedBranchesResponseBodyResultAllowPushUsers    `json:"allowPushUsers,omitempty" xml:"allowPushUsers,omitempty" type:"Repeated"`
	Branch              *string                                                     `json:"branch,omitempty" xml:"branch,omitempty"`
	CreatedAt           *string                                                     `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Id                  *int64                                                      `json:"id,omitempty" xml:"id,omitempty"`
	Matches             []*string                                                   `json:"matches,omitempty" xml:"matches,omitempty" type:"Repeated"`
	MergeRequestSetting *ListProtectedBranchesResponseBodyResultMergeRequestSetting `json:"mergeRequestSetting,omitempty" xml:"mergeRequestSetting,omitempty" type:"Struct"`
	TestSettingDTO      *ListProtectedBranchesResponseBodyResultTestSettingDTO      `json:"testSettingDTO,omitempty" xml:"testSettingDTO,omitempty" type:"Struct"`
	UpdatedAt           *string                                                     `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s ListProtectedBranchesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResult) SetAllowMergeRoles(v []*int32) *ListProtectedBranchesResponseBodyResult {
	s.AllowMergeRoles = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetAllowMergeUserIds(v []*int64) *ListProtectedBranchesResponseBodyResult {
	s.AllowMergeUserIds = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetAllowMergeUsers(v []*ListProtectedBranchesResponseBodyResultAllowMergeUsers) *ListProtectedBranchesResponseBodyResult {
	s.AllowMergeUsers = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetAllowPushRoles(v []*int32) *ListProtectedBranchesResponseBodyResult {
	s.AllowPushRoles = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetAllowPushUserIds(v []*int64) *ListProtectedBranchesResponseBodyResult {
	s.AllowPushUserIds = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetAllowPushUsers(v []*ListProtectedBranchesResponseBodyResultAllowPushUsers) *ListProtectedBranchesResponseBodyResult {
	s.AllowPushUsers = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetBranch(v string) *ListProtectedBranchesResponseBodyResult {
	s.Branch = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetCreatedAt(v string) *ListProtectedBranchesResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetId(v int64) *ListProtectedBranchesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetMatches(v []*string) *ListProtectedBranchesResponseBodyResult {
	s.Matches = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetMergeRequestSetting(v *ListProtectedBranchesResponseBodyResultMergeRequestSetting) *ListProtectedBranchesResponseBodyResult {
	s.MergeRequestSetting = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetTestSettingDTO(v *ListProtectedBranchesResponseBodyResultTestSettingDTO) *ListProtectedBranchesResponseBodyResult {
	s.TestSettingDTO = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetUpdatedAt(v string) *ListProtectedBranchesResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

type ListProtectedBranchesResponseBodyResultAllowMergeUsers struct {
	Avatar   *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	Email    *string `json:"email,omitempty" xml:"email,omitempty"`
	Id       *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	TbUserId *string `json:"tbUserId,omitempty" xml:"tbUserId,omitempty"`
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListProtectedBranchesResponseBodyResultAllowMergeUsers) String() string {
	return tea.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultAllowMergeUsers) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultAllowMergeUsers) SetAvatar(v string) *ListProtectedBranchesResponseBodyResultAllowMergeUsers {
	s.Avatar = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowMergeUsers) SetEmail(v string) *ListProtectedBranchesResponseBodyResultAllowMergeUsers {
	s.Email = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowMergeUsers) SetId(v int64) *ListProtectedBranchesResponseBodyResultAllowMergeUsers {
	s.Id = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowMergeUsers) SetName(v string) *ListProtectedBranchesResponseBodyResultAllowMergeUsers {
	s.Name = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowMergeUsers) SetTbUserId(v string) *ListProtectedBranchesResponseBodyResultAllowMergeUsers {
	s.TbUserId = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowMergeUsers) SetUsername(v string) *ListProtectedBranchesResponseBodyResultAllowMergeUsers {
	s.Username = &v
	return s
}

type ListProtectedBranchesResponseBodyResultAllowPushUsers struct {
	Avatar   *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	Email    *string `json:"email,omitempty" xml:"email,omitempty"`
	Id       *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	TbUserId *string `json:"tbUserId,omitempty" xml:"tbUserId,omitempty"`
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListProtectedBranchesResponseBodyResultAllowPushUsers) String() string {
	return tea.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultAllowPushUsers) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultAllowPushUsers) SetAvatar(v string) *ListProtectedBranchesResponseBodyResultAllowPushUsers {
	s.Avatar = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowPushUsers) SetEmail(v string) *ListProtectedBranchesResponseBodyResultAllowPushUsers {
	s.Email = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowPushUsers) SetId(v int64) *ListProtectedBranchesResponseBodyResultAllowPushUsers {
	s.Id = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowPushUsers) SetName(v string) *ListProtectedBranchesResponseBodyResultAllowPushUsers {
	s.Name = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowPushUsers) SetTbUserId(v string) *ListProtectedBranchesResponseBodyResultAllowPushUsers {
	s.TbUserId = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowPushUsers) SetUsername(v string) *ListProtectedBranchesResponseBodyResultAllowPushUsers {
	s.Username = &v
	return s
}

type ListProtectedBranchesResponseBodyResultMergeRequestSetting struct {
	AllowMergeRequestRoles       []*int32                                                                      `json:"allowMergeRequestRoles,omitempty" xml:"allowMergeRequestRoles,omitempty" type:"Repeated"`
	DefaultAssignees             []*ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees `json:"defaultAssignees,omitempty" xml:"defaultAssignees,omitempty" type:"Repeated"`
	IsAllowSelfApproval          *bool                                                                         `json:"isAllowSelfApproval,omitempty" xml:"isAllowSelfApproval,omitempty"`
	IsRequireDiscussionProcessed *bool                                                                         `json:"isRequireDiscussionProcessed,omitempty" xml:"isRequireDiscussionProcessed,omitempty"`
	IsRequired                   *bool                                                                         `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	IsResetApprovalWhenNewPush   *bool                                                                         `json:"isResetApprovalWhenNewPush,omitempty" xml:"isResetApprovalWhenNewPush,omitempty"`
	MinimumApproval              *int32                                                                        `json:"minimumApproval,omitempty" xml:"minimumApproval,omitempty"`
	MrMode                       *string                                                                       `json:"mrMode,omitempty" xml:"mrMode,omitempty"`
	WhiteList                    *string                                                                       `json:"whiteList,omitempty" xml:"whiteList,omitempty"`
}

func (s ListProtectedBranchesResponseBodyResultMergeRequestSetting) String() string {
	return tea.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultMergeRequestSetting) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetAllowMergeRequestRoles(v []*int32) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.AllowMergeRequestRoles = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetDefaultAssignees(v []*ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.DefaultAssignees = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetIsAllowSelfApproval(v bool) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.IsAllowSelfApproval = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetIsRequireDiscussionProcessed(v bool) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.IsRequireDiscussionProcessed = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetIsRequired(v bool) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.IsRequired = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetIsResetApprovalWhenNewPush(v bool) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.IsResetApprovalWhenNewPush = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetMinimumApproval(v int32) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.MinimumApproval = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetMrMode(v string) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.MrMode = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetWhiteList(v string) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.WhiteList = &v
	return s
}

type ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees struct {
	Avatar   *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	Email    *string `json:"email,omitempty" xml:"email,omitempty"`
	Id       *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	TbUserId *string `json:"tbUserId,omitempty" xml:"tbUserId,omitempty"`
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) String() string {
	return tea.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) SetAvatar(v string) *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees {
	s.Avatar = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) SetEmail(v string) *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees {
	s.Email = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) SetId(v int64) *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees {
	s.Id = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) SetName(v string) *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees {
	s.Name = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) SetTbUserId(v string) *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees {
	s.TbUserId = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) SetUsername(v string) *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees {
	s.Username = &v
	return s
}

type ListProtectedBranchesResponseBodyResultTestSettingDTO struct {
	CheckConfig             *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig             `json:"checkConfig,omitempty" xml:"checkConfig,omitempty" type:"Struct"`
	CheckTaskQualityConfig  *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig  `json:"checkTaskQualityConfig,omitempty" xml:"checkTaskQualityConfig,omitempty" type:"Struct"`
	CodeGuidelinesDetection *ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection `json:"codeGuidelinesDetection,omitempty" xml:"codeGuidelinesDetection,omitempty" type:"Struct"`
	IsRequired              *bool                                                                         `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	SensitiveInfoDetection  *ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection  `json:"sensitiveInfoDetection,omitempty" xml:"sensitiveInfoDetection,omitempty" type:"Struct"`
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTO) String() string {
	return tea.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTO) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTO) SetCheckConfig(v *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) *ListProtectedBranchesResponseBodyResultTestSettingDTO {
	s.CheckConfig = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTO) SetCheckTaskQualityConfig(v *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) *ListProtectedBranchesResponseBodyResultTestSettingDTO {
	s.CheckTaskQualityConfig = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTO) SetCodeGuidelinesDetection(v *ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) *ListProtectedBranchesResponseBodyResultTestSettingDTO {
	s.CodeGuidelinesDetection = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTO) SetIsRequired(v bool) *ListProtectedBranchesResponseBodyResultTestSettingDTO {
	s.IsRequired = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTO) SetSensitiveInfoDetection(v *ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) *ListProtectedBranchesResponseBodyResultTestSettingDTO {
	s.SensitiveInfoDetection = v
	return s
}

type ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig struct {
	CheckItems []*ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems `json:"checkItems,omitempty" xml:"checkItems,omitempty" type:"Repeated"`
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) String() string {
	return tea.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) SetCheckItems(v []*ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig {
	s.CheckItems = v
	return s
}

type ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems struct {
	IsRequired *bool   `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) String() string {
	return tea.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) SetIsRequired(v bool) *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems {
	s.IsRequired = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) SetName(v string) *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems {
	s.Name = &v
	return s
}

type ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig struct {
	BizNo    *string `json:"bizNo,omitempty" xml:"bizNo,omitempty"`
	Enabled  *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Message  *string `json:"message,omitempty" xml:"message,omitempty"`
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) String() string {
	return tea.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetBizNo(v string) *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.BizNo = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetEnabled(v bool) *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.Enabled = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetMessage(v string) *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.Message = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetTaskName(v string) *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.TaskName = &v
	return s
}

type ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection struct {
	Enabled *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) String() string {
	return tea.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) SetEnabled(v bool) *ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection {
	s.Enabled = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) SetMessage(v string) *ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection {
	s.Message = &v
	return s
}

type ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection struct {
	Enabled *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) String() string {
	return tea.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) SetEnabled(v bool) *ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection {
	s.Enabled = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) SetMessage(v string) *ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection {
	s.Message = &v
	return s
}

type ListProtectedBranchesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProtectedBranchesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProtectedBranchesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProtectedBranchesResponse) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponse) SetHeaders(v map[string]*string) *ListProtectedBranchesResponse {
	s.Headers = v
	return s
}

func (s *ListProtectedBranchesResponse) SetStatusCode(v int32) *ListProtectedBranchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProtectedBranchesResponse) SetBody(v *ListProtectedBranchesResponseBody) *ListProtectedBranchesResponse {
	s.Body = v
	return s
}

type ListRepositoriesRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	Archived       *bool   `json:"archived,omitempty" xml:"archived,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	Page           *int64  `json:"page,omitempty" xml:"page,omitempty"`
	PerPage        *int64  `json:"perPage,omitempty" xml:"perPage,omitempty"`
	Search         *string `json:"search,omitempty" xml:"search,omitempty"`
	Sort           *string `json:"sort,omitempty" xml:"sort,omitempty"`
}

func (s ListRepositoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoriesRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoriesRequest) SetAccessToken(v string) *ListRepositoriesRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoriesRequest) SetArchived(v bool) *ListRepositoriesRequest {
	s.Archived = &v
	return s
}

func (s *ListRepositoriesRequest) SetOrderBy(v string) *ListRepositoriesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListRepositoriesRequest) SetOrganizationId(v string) *ListRepositoriesRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoriesRequest) SetPage(v int64) *ListRepositoriesRequest {
	s.Page = &v
	return s
}

func (s *ListRepositoriesRequest) SetPerPage(v int64) *ListRepositoriesRequest {
	s.PerPage = &v
	return s
}

func (s *ListRepositoriesRequest) SetSearch(v string) *ListRepositoriesRequest {
	s.Search = &v
	return s
}

func (s *ListRepositoriesRequest) SetSort(v string) *ListRepositoriesRequest {
	s.Sort = &v
	return s
}

type ListRepositoriesResponseBody struct {
	ErrorCode    *int32                                `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                               `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       []*ListRepositoriesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success      *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
	Total        *int64                                `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListRepositoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoriesResponseBody) SetErrorCode(v int32) *ListRepositoriesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetErrorMessage(v string) *ListRepositoriesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetRequestId(v string) *ListRepositoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetResult(v []*ListRepositoriesResponseBodyResult) *ListRepositoriesResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoriesResponseBody) SetSuccess(v bool) *ListRepositoriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetTotal(v int64) *ListRepositoriesResponseBody {
	s.Total = &v
	return s
}

type ListRepositoriesResponseBodyResult struct {
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	AccessLevel       *int32  `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	Archive           *bool   `json:"archive,omitempty" xml:"archive,omitempty"`
	AvatarUrl         *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	CreatedAt         *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description       *string `json:"description,omitempty" xml:"description,omitempty"`
	ImportStatus      *string `json:"importStatus,omitempty" xml:"importStatus,omitempty"`
	LastActivityAt    *string `json:"lastActivityAt,omitempty" xml:"lastActivityAt,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	NameWithNamespace *string `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	NamespaceId       *int64  `json:"namespaceId,omitempty" xml:"namespaceId,omitempty"`
	Path              *string `json:"path,omitempty" xml:"path,omitempty"`
	PathWithNamespace *string `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	Star              *bool   `json:"star,omitempty" xml:"star,omitempty"`
	StarCount         *int64  `json:"starCount,omitempty" xml:"starCount,omitempty"`
	UpdatedAt         *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	VisibilityLevel   *string `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	WebUrl            *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s ListRepositoriesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoriesResponseBodyResult) SetId(v int64) *ListRepositoriesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetAccessLevel(v int32) *ListRepositoriesResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetArchive(v bool) *ListRepositoriesResponseBodyResult {
	s.Archive = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetAvatarUrl(v string) *ListRepositoriesResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetCreatedAt(v string) *ListRepositoriesResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetDescription(v string) *ListRepositoriesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetImportStatus(v string) *ListRepositoriesResponseBodyResult {
	s.ImportStatus = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetLastActivityAt(v string) *ListRepositoriesResponseBodyResult {
	s.LastActivityAt = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetName(v string) *ListRepositoriesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetNameWithNamespace(v string) *ListRepositoriesResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetNamespaceId(v int64) *ListRepositoriesResponseBodyResult {
	s.NamespaceId = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetPath(v string) *ListRepositoriesResponseBodyResult {
	s.Path = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetPathWithNamespace(v string) *ListRepositoriesResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetStar(v bool) *ListRepositoriesResponseBodyResult {
	s.Star = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetStarCount(v int64) *ListRepositoriesResponseBodyResult {
	s.StarCount = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetUpdatedAt(v string) *ListRepositoriesResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetVisibilityLevel(v string) *ListRepositoriesResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetWebUrl(v string) *ListRepositoriesResponseBodyResult {
	s.WebUrl = &v
	return s
}

type ListRepositoriesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoriesResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoriesResponse) SetHeaders(v map[string]*string) *ListRepositoriesResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoriesResponse) SetStatusCode(v int32) *ListRepositoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoriesResponse) SetBody(v *ListRepositoriesResponseBody) *ListRepositoriesResponse {
	s.Body = v
	return s
}

type ListRepositoryBranchesRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	Page           *int64  `json:"page,omitempty" xml:"page,omitempty"`
	PageSize       *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Search         *string `json:"search,omitempty" xml:"search,omitempty"`
	Sort           *string `json:"sort,omitempty" xml:"sort,omitempty"`
}

func (s ListRepositoryBranchesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryBranchesRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesRequest) SetAccessToken(v string) *ListRepositoryBranchesRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryBranchesRequest) SetOrganizationId(v string) *ListRepositoryBranchesRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryBranchesRequest) SetPage(v int64) *ListRepositoryBranchesRequest {
	s.Page = &v
	return s
}

func (s *ListRepositoryBranchesRequest) SetPageSize(v int64) *ListRepositoryBranchesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepositoryBranchesRequest) SetSearch(v string) *ListRepositoryBranchesRequest {
	s.Search = &v
	return s
}

func (s *ListRepositoryBranchesRequest) SetSort(v string) *ListRepositoryBranchesRequest {
	s.Sort = &v
	return s
}

type ListRepositoryBranchesResponseBody struct {
	ErrorCode    *string                                     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                     `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       []*ListRepositoryBranchesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success      *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
	Total        *string                                     `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListRepositoryBranchesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryBranchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesResponseBody) SetErrorCode(v string) *ListRepositoryBranchesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetErrorMessage(v string) *ListRepositoryBranchesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetRequestId(v string) *ListRepositoryBranchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetResult(v []*ListRepositoryBranchesResponseBodyResult) *ListRepositoryBranchesResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetSuccess(v bool) *ListRepositoryBranchesResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetTotal(v string) *ListRepositoryBranchesResponseBody {
	s.Total = &v
	return s
}

type ListRepositoryBranchesResponseBodyResult struct {
	Commit    *ListRepositoryBranchesResponseBodyResultCommit `json:"commit,omitempty" xml:"commit,omitempty" type:"Struct"`
	Name      *string                                         `json:"name,omitempty" xml:"name,omitempty"`
	Protected *string                                         `json:"protected,omitempty" xml:"protected,omitempty"`
}

func (s ListRepositoryBranchesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryBranchesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesResponseBodyResult) SetCommit(v *ListRepositoryBranchesResponseBodyResultCommit) *ListRepositoryBranchesResponseBodyResult {
	s.Commit = v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResult) SetName(v string) *ListRepositoryBranchesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResult) SetProtected(v string) *ListRepositoryBranchesResponseBodyResult {
	s.Protected = &v
	return s
}

type ListRepositoryBranchesResponseBodyResultCommit struct {
	AuthorEmail    *string   `json:"authorEmail,omitempty" xml:"authorEmail,omitempty"`
	AuthorName     *string   `json:"authorName,omitempty" xml:"authorName,omitempty"`
	AuthoredDate   *string   `json:"authoredDate,omitempty" xml:"authoredDate,omitempty"`
	CommittedDate  *string   `json:"committedDate,omitempty" xml:"committedDate,omitempty"`
	CommitterEmail *string   `json:"committerEmail,omitempty" xml:"committerEmail,omitempty"`
	CommitterName  *string   `json:"committerName,omitempty" xml:"committerName,omitempty"`
	CreatedAt      *string   `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Id             *string   `json:"id,omitempty" xml:"id,omitempty"`
	Message        *string   `json:"message,omitempty" xml:"message,omitempty"`
	ParentIds      []*string `json:"parentIds,omitempty" xml:"parentIds,omitempty" type:"Repeated"`
	ShortId        *string   `json:"shortId,omitempty" xml:"shortId,omitempty"`
	Title          *string   `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListRepositoryBranchesResponseBodyResultCommit) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryBranchesResponseBodyResultCommit) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetAuthorEmail(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.AuthorEmail = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetAuthorName(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.AuthorName = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetAuthoredDate(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.AuthoredDate = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetCommittedDate(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.CommittedDate = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetCommitterEmail(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.CommitterEmail = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetCommitterName(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.CommitterName = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetCreatedAt(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.CreatedAt = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetId(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.Id = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetMessage(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.Message = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetParentIds(v []*string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.ParentIds = v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetShortId(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.ShortId = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetTitle(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.Title = &v
	return s
}

type ListRepositoryBranchesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRepositoryBranchesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoryBranchesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryBranchesResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesResponse) SetHeaders(v map[string]*string) *ListRepositoryBranchesResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryBranchesResponse) SetStatusCode(v int32) *ListRepositoryBranchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryBranchesResponse) SetBody(v *ListRepositoryBranchesResponseBody) *ListRepositoryBranchesResponse {
	s.Body = v
	return s
}

type ListRepositoryCommitDiffRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	ContextLine    *int32  `json:"contextLine,omitempty" xml:"contextLine,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s ListRepositoryCommitDiffRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCommitDiffRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitDiffRequest) SetAccessToken(v string) *ListRepositoryCommitDiffRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryCommitDiffRequest) SetContextLine(v int32) *ListRepositoryCommitDiffRequest {
	s.ContextLine = &v
	return s
}

func (s *ListRepositoryCommitDiffRequest) SetOrganizationId(v string) *ListRepositoryCommitDiffRequest {
	s.OrganizationId = &v
	return s
}

type ListRepositoryCommitDiffResponseBody struct {
	ErrorCode    *string                                       `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                       `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       []*ListRepositoryCommitDiffResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success      *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListRepositoryCommitDiffResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCommitDiffResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitDiffResponseBody) SetErrorCode(v string) *ListRepositoryCommitDiffResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBody) SetErrorMessage(v string) *ListRepositoryCommitDiffResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBody) SetRequestId(v string) *ListRepositoryCommitDiffResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBody) SetResult(v []*ListRepositoryCommitDiffResponseBodyResult) *ListRepositoryCommitDiffResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryCommitDiffResponseBody) SetSuccess(v bool) *ListRepositoryCommitDiffResponseBody {
	s.Success = &v
	return s
}

type ListRepositoryCommitDiffResponseBodyResult struct {
	AMode       *string `json:"aMode,omitempty" xml:"aMode,omitempty"`
	BMode       *string `json:"bMode,omitempty" xml:"bMode,omitempty"`
	DeletedFile *bool   `json:"deletedFile,omitempty" xml:"deletedFile,omitempty"`
	Diff        *string `json:"diff,omitempty" xml:"diff,omitempty"`
	IsBinary    *bool   `json:"isBinary,omitempty" xml:"isBinary,omitempty"`
	IsNewLfs    *bool   `json:"isNewLfs,omitempty" xml:"isNewLfs,omitempty"`
	IsOldLfs    *bool   `json:"isOldLfs,omitempty" xml:"isOldLfs,omitempty"`
	NewFile     *bool   `json:"newFile,omitempty" xml:"newFile,omitempty"`
	NewId       *string `json:"newId,omitempty" xml:"newId,omitempty"`
	NewPath     *string `json:"newPath,omitempty" xml:"newPath,omitempty"`
	OldId       *string `json:"oldId,omitempty" xml:"oldId,omitempty"`
	OldPath     *string `json:"oldPath,omitempty" xml:"oldPath,omitempty"`
	RenamedFile *bool   `json:"renamedFile,omitempty" xml:"renamedFile,omitempty"`
}

func (s ListRepositoryCommitDiffResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCommitDiffResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetAMode(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.AMode = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetBMode(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.BMode = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetDeletedFile(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.DeletedFile = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetDiff(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.Diff = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetIsBinary(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.IsBinary = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetIsNewLfs(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.IsNewLfs = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetIsOldLfs(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.IsOldLfs = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetNewFile(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.NewFile = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetNewId(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.NewId = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetNewPath(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.NewPath = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetOldId(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.OldId = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetOldPath(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.OldPath = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetRenamedFile(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.RenamedFile = &v
	return s
}

type ListRepositoryCommitDiffResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRepositoryCommitDiffResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoryCommitDiffResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCommitDiffResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitDiffResponse) SetHeaders(v map[string]*string) *ListRepositoryCommitDiffResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryCommitDiffResponse) SetStatusCode(v int32) *ListRepositoryCommitDiffResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryCommitDiffResponse) SetBody(v *ListRepositoryCommitDiffResponseBody) *ListRepositoryCommitDiffResponse {
	s.Body = v
	return s
}

type ListRepositoryMemberWithInheritedRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s ListRepositoryMemberWithInheritedRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryMemberWithInheritedRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberWithInheritedRequest) SetAccessToken(v string) *ListRepositoryMemberWithInheritedRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedRequest) SetOrganizationId(v string) *ListRepositoryMemberWithInheritedRequest {
	s.OrganizationId = &v
	return s
}

type ListRepositoryMemberWithInheritedResponseBody struct {
	ErrorCode    *string                                                `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                                `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       []*ListRepositoryMemberWithInheritedResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success      *bool                                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListRepositoryMemberWithInheritedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryMemberWithInheritedResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberWithInheritedResponseBody) SetErrorCode(v string) *ListRepositoryMemberWithInheritedResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBody) SetErrorMessage(v string) *ListRepositoryMemberWithInheritedResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBody) SetRequestId(v string) *ListRepositoryMemberWithInheritedResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBody) SetResult(v []*ListRepositoryMemberWithInheritedResponseBodyResult) *ListRepositoryMemberWithInheritedResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBody) SetSuccess(v bool) *ListRepositoryMemberWithInheritedResponseBody {
	s.Success = &v
	return s
}

type ListRepositoryMemberWithInheritedResponseBodyResult struct {
	AccessLevel  *int32                                                        `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	AvatarUrl    *string                                                       `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Email        *string                                                       `json:"email,omitempty" xml:"email,omitempty"`
	ExternUserId *string                                                       `json:"externUserId,omitempty" xml:"externUserId,omitempty"`
	Id           *int64                                                        `json:"id,omitempty" xml:"id,omitempty"`
	Inherited    *ListRepositoryMemberWithInheritedResponseBodyResultInherited `json:"inherited,omitempty" xml:"inherited,omitempty" type:"Struct"`
	Name         *string                                                       `json:"name,omitempty" xml:"name,omitempty"`
	State        *string                                                       `json:"state,omitempty" xml:"state,omitempty"`
	Username     *string                                                       `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListRepositoryMemberWithInheritedResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryMemberWithInheritedResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetAccessLevel(v int32) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetAvatarUrl(v string) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetEmail(v string) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.Email = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetExternUserId(v string) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetId(v int64) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetInherited(v *ListRepositoryMemberWithInheritedResponseBodyResultInherited) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.Inherited = v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetName(v string) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetState(v string) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetUsername(v string) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.Username = &v
	return s
}

type ListRepositoryMemberWithInheritedResponseBodyResultInherited struct {
	Id                *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	NameWithNamespace *string `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	Path              *string `json:"path,omitempty" xml:"path,omitempty"`
	PathWithNamespace *string `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	Type              *string `json:"type,omitempty" xml:"type,omitempty"`
	VisibilityLevel   *string `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
}

func (s ListRepositoryMemberWithInheritedResponseBodyResultInherited) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryMemberWithInheritedResponseBodyResultInherited) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetId(v int64) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.Id = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetName(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.Name = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetNameWithNamespace(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.NameWithNamespace = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetPath(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.Path = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetPathWithNamespace(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.PathWithNamespace = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetType(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.Type = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetVisibilityLevel(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.VisibilityLevel = &v
	return s
}

type ListRepositoryMemberWithInheritedResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRepositoryMemberWithInheritedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoryMemberWithInheritedResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryMemberWithInheritedResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberWithInheritedResponse) SetHeaders(v map[string]*string) *ListRepositoryMemberWithInheritedResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponse) SetStatusCode(v int32) *ListRepositoryMemberWithInheritedResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponse) SetBody(v *ListRepositoryMemberWithInheritedResponseBody) *ListRepositoryMemberWithInheritedResponse {
	s.Body = v
	return s
}

type ListRepositoryTreeRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	Path           *string `json:"path,omitempty" xml:"path,omitempty"`
	RefName        *string `json:"refName,omitempty" xml:"refName,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListRepositoryTreeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTreeRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryTreeRequest) SetAccessToken(v string) *ListRepositoryTreeRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryTreeRequest) SetOrganizationId(v string) *ListRepositoryTreeRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryTreeRequest) SetPath(v string) *ListRepositoryTreeRequest {
	s.Path = &v
	return s
}

func (s *ListRepositoryTreeRequest) SetRefName(v string) *ListRepositoryTreeRequest {
	s.RefName = &v
	return s
}

func (s *ListRepositoryTreeRequest) SetType(v string) *ListRepositoryTreeRequest {
	s.Type = &v
	return s
}

type ListRepositoryTreeResponseBody struct {
	ErrorCode    *string                                 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                 `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       []*ListRepositoryTreeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success      *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListRepositoryTreeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTreeResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryTreeResponseBody) SetErrorCode(v string) *ListRepositoryTreeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryTreeResponseBody) SetErrorMessage(v string) *ListRepositoryTreeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryTreeResponseBody) SetRequestId(v string) *ListRepositoryTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryTreeResponseBody) SetResult(v []*ListRepositoryTreeResponseBodyResult) *ListRepositoryTreeResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryTreeResponseBody) SetSuccess(v bool) *ListRepositoryTreeResponseBody {
	s.Success = &v
	return s
}

type ListRepositoryTreeResponseBodyResult struct {
	Id    *string `json:"id,omitempty" xml:"id,omitempty"`
	IsLFS *bool   `json:"isLFS,omitempty" xml:"isLFS,omitempty"`
	Mode  *string `json:"mode,omitempty" xml:"mode,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Path  *string `json:"path,omitempty" xml:"path,omitempty"`
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListRepositoryTreeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTreeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryTreeResponseBodyResult) SetId(v string) *ListRepositoryTreeResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRepositoryTreeResponseBodyResult) SetIsLFS(v bool) *ListRepositoryTreeResponseBodyResult {
	s.IsLFS = &v
	return s
}

func (s *ListRepositoryTreeResponseBodyResult) SetMode(v string) *ListRepositoryTreeResponseBodyResult {
	s.Mode = &v
	return s
}

func (s *ListRepositoryTreeResponseBodyResult) SetName(v string) *ListRepositoryTreeResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListRepositoryTreeResponseBodyResult) SetPath(v string) *ListRepositoryTreeResponseBodyResult {
	s.Path = &v
	return s
}

func (s *ListRepositoryTreeResponseBodyResult) SetType(v string) *ListRepositoryTreeResponseBodyResult {
	s.Type = &v
	return s
}

type ListRepositoryTreeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRepositoryTreeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoryTreeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTreeResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryTreeResponse) SetHeaders(v map[string]*string) *ListRepositoryTreeResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryTreeResponse) SetStatusCode(v int32) *ListRepositoryTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryTreeResponse) SetBody(v *ListRepositoryTreeResponseBody) *ListRepositoryTreeResponse {
	s.Body = v
	return s
}

type ListRepositoryWebhookRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	Page           *int64  `json:"page,omitempty" xml:"page,omitempty"`
	PageSize       *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListRepositoryWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryWebhookRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryWebhookRequest) SetAccessToken(v string) *ListRepositoryWebhookRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryWebhookRequest) SetOrganizationId(v string) *ListRepositoryWebhookRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryWebhookRequest) SetPage(v int64) *ListRepositoryWebhookRequest {
	s.Page = &v
	return s
}

func (s *ListRepositoryWebhookRequest) SetPageSize(v int64) *ListRepositoryWebhookRequest {
	s.PageSize = &v
	return s
}

type ListRepositoryWebhookResponseBody struct {
	ErrorCode    *string                                    `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                    `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       []*ListRepositoryWebhookResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success      *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
	Total        *int64                                     `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListRepositoryWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryWebhookResponseBody) SetErrorCode(v string) *ListRepositoryWebhookResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetErrorMessage(v string) *ListRepositoryWebhookResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetRequestId(v string) *ListRepositoryWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetResult(v []*ListRepositoryWebhookResponseBodyResult) *ListRepositoryWebhookResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetSuccess(v bool) *ListRepositoryWebhookResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetTotal(v int64) *ListRepositoryWebhookResponseBody {
	s.Total = &v
	return s
}

type ListRepositoryWebhookResponseBodyResult struct {
	CreatedAt             *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description           *string `json:"description,omitempty" xml:"description,omitempty"`
	EnableSslVerification *bool   `json:"enableSslVerification,omitempty" xml:"enableSslVerification,omitempty"`
	Id                    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	LastTestResult        *string `json:"lastTestResult,omitempty" xml:"lastTestResult,omitempty"`
	MergeRequestsEvents   *bool   `json:"mergeRequestsEvents,omitempty" xml:"mergeRequestsEvents,omitempty"`
	NoteEvents            *bool   `json:"noteEvents,omitempty" xml:"noteEvents,omitempty"`
	ProjectId             *int64  `json:"projectId,omitempty" xml:"projectId,omitempty"`
	PushEvents            *bool   `json:"pushEvents,omitempty" xml:"pushEvents,omitempty"`
	SecretToken           *string `json:"secretToken,omitempty" xml:"secretToken,omitempty"`
	TagPushEvents         *bool   `json:"tagPushEvents,omitempty" xml:"tagPushEvents,omitempty"`
	Url                   *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListRepositoryWebhookResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryWebhookResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryWebhookResponseBodyResult) SetCreatedAt(v string) *ListRepositoryWebhookResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetDescription(v string) *ListRepositoryWebhookResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetEnableSslVerification(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.EnableSslVerification = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetId(v int64) *ListRepositoryWebhookResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetLastTestResult(v string) *ListRepositoryWebhookResponseBodyResult {
	s.LastTestResult = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetMergeRequestsEvents(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.MergeRequestsEvents = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetNoteEvents(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.NoteEvents = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetProjectId(v int64) *ListRepositoryWebhookResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetPushEvents(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.PushEvents = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetSecretToken(v string) *ListRepositoryWebhookResponseBodyResult {
	s.SecretToken = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetTagPushEvents(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.TagPushEvents = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetUrl(v string) *ListRepositoryWebhookResponseBodyResult {
	s.Url = &v
	return s
}

type ListRepositoryWebhookResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRepositoryWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoryWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryWebhookResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryWebhookResponse) SetHeaders(v map[string]*string) *ListRepositoryWebhookResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryWebhookResponse) SetStatusCode(v int32) *ListRepositoryWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryWebhookResponse) SetBody(v *ListRepositoryWebhookResponseBody) *ListRepositoryWebhookResponse {
	s.Body = v
	return s
}

type ListResourceMembersResponseBody struct {
	ErrorCode       *string                                           `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage    *string                                           `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId       *string                                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResourceMembers []*ListResourceMembersResponseBodyResourceMembers `json:"resourceMembers,omitempty" xml:"resourceMembers,omitempty" type:"Repeated"`
	Success         *bool                                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListResourceMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceMembersResponseBody) SetErrorCode(v string) *ListResourceMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListResourceMembersResponseBody) SetErrorMessage(v string) *ListResourceMembersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListResourceMembersResponseBody) SetRequestId(v string) *ListResourceMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceMembersResponseBody) SetResourceMembers(v []*ListResourceMembersResponseBodyResourceMembers) *ListResourceMembersResponseBody {
	s.ResourceMembers = v
	return s
}

func (s *ListResourceMembersResponseBody) SetSuccess(v bool) *ListResourceMembersResponseBody {
	s.Success = &v
	return s
}

type ListResourceMembersResponseBodyResourceMembers struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	RoleName  *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	Username  *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListResourceMembersResponseBodyResourceMembers) String() string {
	return tea.Prettify(s)
}

func (s ListResourceMembersResponseBodyResourceMembers) GoString() string {
	return s.String()
}

func (s *ListResourceMembersResponseBodyResourceMembers) SetAccountId(v string) *ListResourceMembersResponseBodyResourceMembers {
	s.AccountId = &v
	return s
}

func (s *ListResourceMembersResponseBodyResourceMembers) SetRoleName(v string) *ListResourceMembersResponseBodyResourceMembers {
	s.RoleName = &v
	return s
}

func (s *ListResourceMembersResponseBodyResourceMembers) SetUsername(v string) *ListResourceMembersResponseBodyResourceMembers {
	s.Username = &v
	return s
}

type ListResourceMembersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceMembersResponse) GoString() string {
	return s.String()
}

func (s *ListResourceMembersResponse) SetHeaders(v map[string]*string) *ListResourceMembersResponse {
	s.Headers = v
	return s
}

func (s *ListResourceMembersResponse) SetStatusCode(v int32) *ListResourceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceMembersResponse) SetBody(v *ListResourceMembersResponseBody) *ListResourceMembersResponse {
	s.Body = v
	return s
}

type ListServiceConnectionsRequest struct {
	SericeConnectionType *string `json:"sericeConnectionType,omitempty" xml:"sericeConnectionType,omitempty"`
}

func (s ListServiceConnectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceConnectionsRequest) SetSericeConnectionType(v string) *ListServiceConnectionsRequest {
	s.SericeConnectionType = &v
	return s
}

type ListServiceConnectionsResponseBody struct {
	ErrorCode          *string                                                 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage       *string                                                 `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId          *string                                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ServiceConnections []*ListServiceConnectionsResponseBodyServiceConnections `json:"serviceConnections,omitempty" xml:"serviceConnections,omitempty" type:"Repeated"`
	Success            *bool                                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListServiceConnectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceConnectionsResponseBody) SetErrorCode(v string) *ListServiceConnectionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListServiceConnectionsResponseBody) SetErrorMessage(v string) *ListServiceConnectionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListServiceConnectionsResponseBody) SetRequestId(v string) *ListServiceConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceConnectionsResponseBody) SetServiceConnections(v []*ListServiceConnectionsResponseBodyServiceConnections) *ListServiceConnectionsResponseBody {
	s.ServiceConnections = v
	return s
}

func (s *ListServiceConnectionsResponseBody) SetSuccess(v bool) *ListServiceConnectionsResponseBody {
	s.Success = &v
	return s
}

type ListServiceConnectionsResponseBodyServiceConnections struct {
	CreateTime     *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OwnerAccountId *int64  `json:"ownerAccountId,omitempty" xml:"ownerAccountId,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListServiceConnectionsResponseBodyServiceConnections) String() string {
	return tea.Prettify(s)
}

func (s ListServiceConnectionsResponseBodyServiceConnections) GoString() string {
	return s.String()
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) SetCreateTime(v int64) *ListServiceConnectionsResponseBodyServiceConnections {
	s.CreateTime = &v
	return s
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) SetId(v int64) *ListServiceConnectionsResponseBodyServiceConnections {
	s.Id = &v
	return s
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) SetName(v string) *ListServiceConnectionsResponseBodyServiceConnections {
	s.Name = &v
	return s
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) SetOwnerAccountId(v int64) *ListServiceConnectionsResponseBodyServiceConnections {
	s.OwnerAccountId = &v
	return s
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) SetType(v string) *ListServiceConnectionsResponseBodyServiceConnections {
	s.Type = &v
	return s
}

type ListServiceConnectionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServiceConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServiceConnectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceConnectionsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceConnectionsResponse) SetHeaders(v map[string]*string) *ListServiceConnectionsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceConnectionsResponse) SetStatusCode(v int32) *ListServiceConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceConnectionsResponse) SetBody(v *ListServiceConnectionsResponseBody) *ListServiceConnectionsResponse {
	s.Body = v
	return s
}

type ListSprintsRequest struct {
	MaxResults      *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken       *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	SpaceType       *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s ListSprintsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSprintsRequest) GoString() string {
	return s.String()
}

func (s *ListSprintsRequest) SetMaxResults(v int64) *ListSprintsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSprintsRequest) SetNextToken(v string) *ListSprintsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSprintsRequest) SetSpaceIdentifier(v string) *ListSprintsRequest {
	s.SpaceIdentifier = &v
	return s
}

func (s *ListSprintsRequest) SetSpaceType(v string) *ListSprintsRequest {
	s.SpaceType = &v
	return s
}

type ListSprintsResponseBody struct {
	ErrorCode  *string                           `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg   *string                           `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	MaxResults *int64                            `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                           `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId  *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Sprints    []*ListSprintsResponseBodySprints `json:"sprints,omitempty" xml:"sprints,omitempty" type:"Repeated"`
	Success    *bool                             `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount *int64                            `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSprintsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSprintsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSprintsResponseBody) SetErrorCode(v string) *ListSprintsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSprintsResponseBody) SetErrorMsg(v string) *ListSprintsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListSprintsResponseBody) SetMaxResults(v int64) *ListSprintsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSprintsResponseBody) SetNextToken(v string) *ListSprintsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSprintsResponseBody) SetRequestId(v string) *ListSprintsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSprintsResponseBody) SetSprints(v []*ListSprintsResponseBodySprints) *ListSprintsResponseBody {
	s.Sprints = v
	return s
}

func (s *ListSprintsResponseBody) SetSuccess(v bool) *ListSprintsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSprintsResponseBody) SetTotalCount(v int64) *ListSprintsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSprintsResponseBodySprints struct {
	Creator         *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	EndDate         *int64  `json:"endDate,omitempty" xml:"endDate,omitempty"`
	GmtCreate       *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified     *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Identifier      *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Modifier        *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	Scope           *string `json:"scope,omitempty" xml:"scope,omitempty"`
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	StartDate       *int64  `json:"startDate,omitempty" xml:"startDate,omitempty"`
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListSprintsResponseBodySprints) String() string {
	return tea.Prettify(s)
}

func (s ListSprintsResponseBodySprints) GoString() string {
	return s.String()
}

func (s *ListSprintsResponseBodySprints) SetCreator(v string) *ListSprintsResponseBodySprints {
	s.Creator = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetDescription(v string) *ListSprintsResponseBodySprints {
	s.Description = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetEndDate(v int64) *ListSprintsResponseBodySprints {
	s.EndDate = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetGmtCreate(v int64) *ListSprintsResponseBodySprints {
	s.GmtCreate = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetGmtModified(v int64) *ListSprintsResponseBodySprints {
	s.GmtModified = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetIdentifier(v string) *ListSprintsResponseBodySprints {
	s.Identifier = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetModifier(v string) *ListSprintsResponseBodySprints {
	s.Modifier = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetName(v string) *ListSprintsResponseBodySprints {
	s.Name = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetScope(v string) *ListSprintsResponseBodySprints {
	s.Scope = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetSpaceIdentifier(v string) *ListSprintsResponseBodySprints {
	s.SpaceIdentifier = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetStartDate(v int64) *ListSprintsResponseBodySprints {
	s.StartDate = &v
	return s
}

func (s *ListSprintsResponseBodySprints) SetStatus(v string) *ListSprintsResponseBodySprints {
	s.Status = &v
	return s
}

type ListSprintsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSprintsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSprintsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSprintsResponse) GoString() string {
	return s.String()
}

func (s *ListSprintsResponse) SetHeaders(v map[string]*string) *ListSprintsResponse {
	s.Headers = v
	return s
}

func (s *ListSprintsResponse) SetStatusCode(v int32) *ListSprintsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSprintsResponse) SetBody(v *ListSprintsResponseBody) *ListSprintsResponse {
	s.Body = v
	return s
}

type ListVariableGroupsRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PageOrder  *string `json:"pageOrder,omitempty" xml:"pageOrder,omitempty"`
	PageSort   *string `json:"pageSort,omitempty" xml:"pageSort,omitempty"`
}

func (s ListVariableGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVariableGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListVariableGroupsRequest) SetMaxResults(v int32) *ListVariableGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVariableGroupsRequest) SetNextToken(v string) *ListVariableGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVariableGroupsRequest) SetPageOrder(v string) *ListVariableGroupsRequest {
	s.PageOrder = &v
	return s
}

func (s *ListVariableGroupsRequest) SetPageSort(v string) *ListVariableGroupsRequest {
	s.PageSort = &v
	return s
}

type ListVariableGroupsResponseBody struct {
	ErrorCode      *string                                         `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage   *string                                         `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	NextToken      *string                                         `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId      *string                                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                                           `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount     *int64                                          `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	VariableGroups []*ListVariableGroupsResponseBodyVariableGroups `json:"variableGroups,omitempty" xml:"variableGroups,omitempty" type:"Repeated"`
}

func (s ListVariableGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVariableGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVariableGroupsResponseBody) SetErrorCode(v string) *ListVariableGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListVariableGroupsResponseBody) SetErrorMessage(v string) *ListVariableGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListVariableGroupsResponseBody) SetNextToken(v string) *ListVariableGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVariableGroupsResponseBody) SetRequestId(v string) *ListVariableGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVariableGroupsResponseBody) SetSuccess(v bool) *ListVariableGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListVariableGroupsResponseBody) SetTotalCount(v int64) *ListVariableGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVariableGroupsResponseBody) SetVariableGroups(v []*ListVariableGroupsResponseBodyVariableGroups) *ListVariableGroupsResponseBody {
	s.VariableGroups = v
	return s
}

type ListVariableGroupsResponseBodyVariableGroups struct {
	CreateTime        *int64                                                          `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorAccountId  *string                                                         `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	Description       *string                                                         `json:"description,omitempty" xml:"description,omitempty"`
	Id                *int64                                                          `json:"id,omitempty" xml:"id,omitempty"`
	ModifierAccountId *string                                                         `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	Name              *string                                                         `json:"name,omitempty" xml:"name,omitempty"`
	RelatedPipelines  []*ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines `json:"relatedPipelines,omitempty" xml:"relatedPipelines,omitempty" type:"Repeated"`
	UpdateTime        *int64                                                          `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	Variables         []*ListVariableGroupsResponseBodyVariableGroupsVariables        `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s ListVariableGroupsResponseBodyVariableGroups) String() string {
	return tea.Prettify(s)
}

func (s ListVariableGroupsResponseBodyVariableGroups) GoString() string {
	return s.String()
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetCreateTime(v int64) *ListVariableGroupsResponseBodyVariableGroups {
	s.CreateTime = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetCreatorAccountId(v string) *ListVariableGroupsResponseBodyVariableGroups {
	s.CreatorAccountId = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetDescription(v string) *ListVariableGroupsResponseBodyVariableGroups {
	s.Description = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetId(v int64) *ListVariableGroupsResponseBodyVariableGroups {
	s.Id = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetModifierAccountId(v string) *ListVariableGroupsResponseBodyVariableGroups {
	s.ModifierAccountId = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetName(v string) *ListVariableGroupsResponseBodyVariableGroups {
	s.Name = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetRelatedPipelines(v []*ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines) *ListVariableGroupsResponseBodyVariableGroups {
	s.RelatedPipelines = v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetUpdateTime(v int64) *ListVariableGroupsResponseBodyVariableGroups {
	s.UpdateTime = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroups) SetVariables(v []*ListVariableGroupsResponseBodyVariableGroupsVariables) *ListVariableGroupsResponseBodyVariableGroups {
	s.Variables = v
	return s
}

type ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines struct {
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines) String() string {
	return tea.Prettify(s)
}

func (s ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines) GoString() string {
	return s.String()
}

func (s *ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines) SetId(v int64) *ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines {
	s.Id = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines) SetName(v string) *ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines {
	s.Name = &v
	return s
}

type ListVariableGroupsResponseBodyVariableGroupsVariables struct {
	IsEncrypted *bool   `json:"isEncrypted,omitempty" xml:"isEncrypted,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListVariableGroupsResponseBodyVariableGroupsVariables) String() string {
	return tea.Prettify(s)
}

func (s ListVariableGroupsResponseBodyVariableGroupsVariables) GoString() string {
	return s.String()
}

func (s *ListVariableGroupsResponseBodyVariableGroupsVariables) SetIsEncrypted(v bool) *ListVariableGroupsResponseBodyVariableGroupsVariables {
	s.IsEncrypted = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroupsVariables) SetName(v string) *ListVariableGroupsResponseBodyVariableGroupsVariables {
	s.Name = &v
	return s
}

func (s *ListVariableGroupsResponseBodyVariableGroupsVariables) SetValue(v string) *ListVariableGroupsResponseBodyVariableGroupsVariables {
	s.Value = &v
	return s
}

type ListVariableGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVariableGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVariableGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVariableGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListVariableGroupsResponse) SetHeaders(v map[string]*string) *ListVariableGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListVariableGroupsResponse) SetStatusCode(v int32) *ListVariableGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVariableGroupsResponse) SetBody(v *ListVariableGroupsResponseBody) *ListVariableGroupsResponse {
	s.Body = v
	return s
}

type ListWorkItemAllFieldsRequest struct {
	SpaceIdentifier        *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	SpaceType              *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	WorkitemTypeIdentifier *string `json:"workitemTypeIdentifier,omitempty" xml:"workitemTypeIdentifier,omitempty"`
}

func (s ListWorkItemAllFieldsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkItemAllFieldsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkItemAllFieldsRequest) SetSpaceIdentifier(v string) *ListWorkItemAllFieldsRequest {
	s.SpaceIdentifier = &v
	return s
}

func (s *ListWorkItemAllFieldsRequest) SetSpaceType(v string) *ListWorkItemAllFieldsRequest {
	s.SpaceType = &v
	return s
}

func (s *ListWorkItemAllFieldsRequest) SetWorkitemTypeIdentifier(v string) *ListWorkItemAllFieldsRequest {
	s.WorkitemTypeIdentifier = &v
	return s
}

type ListWorkItemAllFieldsResponseBody struct {
	ErrorCode *string                                    `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                    `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Fields    []*ListWorkItemAllFieldsResponseBodyFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListWorkItemAllFieldsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkItemAllFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkItemAllFieldsResponseBody) SetErrorCode(v string) *ListWorkItemAllFieldsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBody) SetErrorMsg(v string) *ListWorkItemAllFieldsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBody) SetFields(v []*ListWorkItemAllFieldsResponseBodyFields) *ListWorkItemAllFieldsResponseBody {
	s.Fields = v
	return s
}

func (s *ListWorkItemAllFieldsResponseBody) SetRequestId(v string) *ListWorkItemAllFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBody) SetSuccess(v bool) *ListWorkItemAllFieldsResponseBody {
	s.Success = &v
	return s
}

type ListWorkItemAllFieldsResponseBodyFields struct {
	ClassName        *string                                           `json:"className,omitempty" xml:"className,omitempty"`
	Creator          *string                                           `json:"creator,omitempty" xml:"creator,omitempty"`
	DefaultValue     *string                                           `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	Description      *string                                           `json:"description,omitempty" xml:"description,omitempty"`
	Format           *string                                           `json:"format,omitempty" xml:"format,omitempty"`
	GmtCreate        *int64                                            `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified      *int64                                            `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Identifier       *string                                           `json:"identifier,omitempty" xml:"identifier,omitempty"`
	IsRequired       *bool                                             `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	IsShowWhenCreate *bool                                             `json:"isShowWhenCreate,omitempty" xml:"isShowWhenCreate,omitempty"`
	IsSystemRequired *bool                                             `json:"isSystemRequired,omitempty" xml:"isSystemRequired,omitempty"`
	LinkWithService  *string                                           `json:"linkWithService,omitempty" xml:"linkWithService,omitempty"`
	Modifier         *string                                           `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Name             *string                                           `json:"name,omitempty" xml:"name,omitempty"`
	Options          []*ListWorkItemAllFieldsResponseBodyFieldsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	ResourceType     *string                                           `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Type             *string                                           `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListWorkItemAllFieldsResponseBodyFields) String() string {
	return tea.Prettify(s)
}

func (s ListWorkItemAllFieldsResponseBodyFields) GoString() string {
	return s.String()
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetClassName(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.ClassName = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetCreator(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.Creator = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetDefaultValue(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.DefaultValue = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetDescription(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.Description = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetFormat(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.Format = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetGmtCreate(v int64) *ListWorkItemAllFieldsResponseBodyFields {
	s.GmtCreate = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetGmtModified(v int64) *ListWorkItemAllFieldsResponseBodyFields {
	s.GmtModified = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetIdentifier(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.Identifier = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetIsRequired(v bool) *ListWorkItemAllFieldsResponseBodyFields {
	s.IsRequired = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetIsShowWhenCreate(v bool) *ListWorkItemAllFieldsResponseBodyFields {
	s.IsShowWhenCreate = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetIsSystemRequired(v bool) *ListWorkItemAllFieldsResponseBodyFields {
	s.IsSystemRequired = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetLinkWithService(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.LinkWithService = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetModifier(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.Modifier = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetName(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.Name = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetOptions(v []*ListWorkItemAllFieldsResponseBodyFieldsOptions) *ListWorkItemAllFieldsResponseBodyFields {
	s.Options = v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetResourceType(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.ResourceType = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetType(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.Type = &v
	return s
}

type ListWorkItemAllFieldsResponseBodyFieldsOptions struct {
	DisplayValue    *string `json:"displayValue,omitempty" xml:"displayValue,omitempty"`
	FieldIdentifier *string `json:"fieldIdentifier,omitempty" xml:"fieldIdentifier,omitempty"`
	Identifier      *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Level           *int64  `json:"level,omitempty" xml:"level,omitempty"`
	Position        *int64  `json:"position,omitempty" xml:"position,omitempty"`
	Value           *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueEn         *string `json:"valueEn,omitempty" xml:"valueEn,omitempty"`
}

func (s ListWorkItemAllFieldsResponseBodyFieldsOptions) String() string {
	return tea.Prettify(s)
}

func (s ListWorkItemAllFieldsResponseBodyFieldsOptions) GoString() string {
	return s.String()
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) SetDisplayValue(v string) *ListWorkItemAllFieldsResponseBodyFieldsOptions {
	s.DisplayValue = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) SetFieldIdentifier(v string) *ListWorkItemAllFieldsResponseBodyFieldsOptions {
	s.FieldIdentifier = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) SetIdentifier(v string) *ListWorkItemAllFieldsResponseBodyFieldsOptions {
	s.Identifier = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) SetLevel(v int64) *ListWorkItemAllFieldsResponseBodyFieldsOptions {
	s.Level = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) SetPosition(v int64) *ListWorkItemAllFieldsResponseBodyFieldsOptions {
	s.Position = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) SetValue(v string) *ListWorkItemAllFieldsResponseBodyFieldsOptions {
	s.Value = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) SetValueEn(v string) *ListWorkItemAllFieldsResponseBodyFieldsOptions {
	s.ValueEn = &v
	return s
}

type ListWorkItemAllFieldsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListWorkItemAllFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWorkItemAllFieldsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkItemAllFieldsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkItemAllFieldsResponse) SetHeaders(v map[string]*string) *ListWorkItemAllFieldsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkItemAllFieldsResponse) SetStatusCode(v int32) *ListWorkItemAllFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkItemAllFieldsResponse) SetBody(v *ListWorkItemAllFieldsResponseBody) *ListWorkItemAllFieldsResponse {
	s.Body = v
	return s
}

type ListWorkItemWorkFlowStatusRequest struct {
	SpaceIdentifier            *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	SpaceType                  *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	WorkitemCategoryIdentifier *string `json:"workitemCategoryIdentifier,omitempty" xml:"workitemCategoryIdentifier,omitempty"`
	WorkitemTypeIdentifier     *string `json:"workitemTypeIdentifier,omitempty" xml:"workitemTypeIdentifier,omitempty"`
}

func (s ListWorkItemWorkFlowStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkItemWorkFlowStatusRequest) GoString() string {
	return s.String()
}

func (s *ListWorkItemWorkFlowStatusRequest) SetSpaceIdentifier(v string) *ListWorkItemWorkFlowStatusRequest {
	s.SpaceIdentifier = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusRequest) SetSpaceType(v string) *ListWorkItemWorkFlowStatusRequest {
	s.SpaceType = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusRequest) SetWorkitemCategoryIdentifier(v string) *ListWorkItemWorkFlowStatusRequest {
	s.WorkitemCategoryIdentifier = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusRequest) SetWorkitemTypeIdentifier(v string) *ListWorkItemWorkFlowStatusRequest {
	s.WorkitemTypeIdentifier = &v
	return s
}

type ListWorkItemWorkFlowStatusResponseBody struct {
	ErrorCode    *string                                           `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                           `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Statuses     []*ListWorkItemWorkFlowStatusResponseBodyStatuses `json:"statuses,omitempty" xml:"statuses,omitempty" type:"Repeated"`
	Success      *bool                                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListWorkItemWorkFlowStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkItemWorkFlowStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkItemWorkFlowStatusResponseBody) SetErrorCode(v string) *ListWorkItemWorkFlowStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBody) SetErrorMessage(v string) *ListWorkItemWorkFlowStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBody) SetRequestId(v string) *ListWorkItemWorkFlowStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBody) SetStatuses(v []*ListWorkItemWorkFlowStatusResponseBodyStatuses) *ListWorkItemWorkFlowStatusResponseBody {
	s.Statuses = v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBody) SetSuccess(v bool) *ListWorkItemWorkFlowStatusResponseBody {
	s.Success = &v
	return s
}

type ListWorkItemWorkFlowStatusResponseBodyStatuses struct {
	Creator                 *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Description             *string `json:"description,omitempty" xml:"description,omitempty"`
	GmtCreate               *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified             *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Identifier              *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Modifier                *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Name                    *string `json:"name,omitempty" xml:"name,omitempty"`
	ResourceType            *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Source                  *string `json:"source,omitempty" xml:"source,omitempty"`
	WorkflowStageIdentifier *string `json:"workflowStageIdentifier,omitempty" xml:"workflowStageIdentifier,omitempty"`
	WorkflowStageName       *string `json:"workflowStageName,omitempty" xml:"workflowStageName,omitempty"`
}

func (s ListWorkItemWorkFlowStatusResponseBodyStatuses) String() string {
	return tea.Prettify(s)
}

func (s ListWorkItemWorkFlowStatusResponseBodyStatuses) GoString() string {
	return s.String()
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetCreator(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.Creator = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetDescription(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.Description = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetGmtCreate(v int64) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.GmtCreate = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetGmtModified(v int64) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.GmtModified = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetIdentifier(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.Identifier = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetModifier(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.Modifier = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetName(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.Name = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetResourceType(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.ResourceType = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetSource(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.Source = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetWorkflowStageIdentifier(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.WorkflowStageIdentifier = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponseBodyStatuses) SetWorkflowStageName(v string) *ListWorkItemWorkFlowStatusResponseBodyStatuses {
	s.WorkflowStageName = &v
	return s
}

type ListWorkItemWorkFlowStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListWorkItemWorkFlowStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWorkItemWorkFlowStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkItemWorkFlowStatusResponse) GoString() string {
	return s.String()
}

func (s *ListWorkItemWorkFlowStatusResponse) SetHeaders(v map[string]*string) *ListWorkItemWorkFlowStatusResponse {
	s.Headers = v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponse) SetStatusCode(v int32) *ListWorkItemWorkFlowStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusResponse) SetBody(v *ListWorkItemWorkFlowStatusResponseBody) *ListWorkItemWorkFlowStatusResponse {
	s.Body = v
	return s
}

type ListWorkitemTimeResponseBody struct {
	Code         *int64                                      `json:"code,omitempty" xml:"code,omitempty"`
	ErrorCode    *string                                     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg     *string                                     `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId    *string                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
	WorkitemTime []*ListWorkitemTimeResponseBodyWorkitemTime `json:"workitemTime,omitempty" xml:"workitemTime,omitempty" type:"Repeated"`
}

func (s ListWorkitemTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkitemTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkitemTimeResponseBody) SetCode(v int64) *ListWorkitemTimeResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkitemTimeResponseBody) SetErrorCode(v string) *ListWorkitemTimeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListWorkitemTimeResponseBody) SetErrorMsg(v string) *ListWorkitemTimeResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListWorkitemTimeResponseBody) SetRequestId(v string) *ListWorkitemTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkitemTimeResponseBody) SetSuccess(v bool) *ListWorkitemTimeResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkitemTimeResponseBody) SetWorkitemTime(v []*ListWorkitemTimeResponseBodyWorkitemTime) *ListWorkitemTimeResponseBody {
	s.WorkitemTime = v
	return s
}

type ListWorkitemTimeResponseBodyWorkitemTime struct {
	ActualTime         *int64  `json:"actualTime,omitempty" xml:"actualTime,omitempty"`
	Description        *string `json:"description,omitempty" xml:"description,omitempty"`
	GmtCreate          *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtEnd             *int64  `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	GmtModified        *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	GmtStart           *int64  `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
	Identifier         *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	RecordUser         *string `json:"recordUser,omitempty" xml:"recordUser,omitempty"`
	Type               *string `json:"type,omitempty" xml:"type,omitempty"`
	WorkitemIdentifier *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s ListWorkitemTimeResponseBodyWorkitemTime) String() string {
	return tea.Prettify(s)
}

func (s ListWorkitemTimeResponseBodyWorkitemTime) GoString() string {
	return s.String()
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetActualTime(v int64) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.ActualTime = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetDescription(v string) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.Description = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetGmtCreate(v int64) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.GmtCreate = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetGmtEnd(v int64) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.GmtEnd = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetGmtModified(v int64) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.GmtModified = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetGmtStart(v int64) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.GmtStart = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetIdentifier(v string) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.Identifier = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetRecordUser(v string) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.RecordUser = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetType(v string) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.Type = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetWorkitemIdentifier(v string) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.WorkitemIdentifier = &v
	return s
}

type ListWorkitemTimeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListWorkitemTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWorkitemTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkitemTimeResponse) GoString() string {
	return s.String()
}

func (s *ListWorkitemTimeResponse) SetHeaders(v map[string]*string) *ListWorkitemTimeResponse {
	s.Headers = v
	return s
}

func (s *ListWorkitemTimeResponse) SetStatusCode(v int32) *ListWorkitemTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkitemTimeResponse) SetBody(v *ListWorkitemTimeResponseBody) *ListWorkitemTimeResponse {
	s.Body = v
	return s
}

type ListWorkitemsRequest struct {
	Category        *string `json:"category,omitempty" xml:"category,omitempty"`
	Conditions      *string `json:"conditions,omitempty" xml:"conditions,omitempty"`
	ExtraConditions *string `json:"extraConditions,omitempty" xml:"extraConditions,omitempty"`
	GroupCondition  *string `json:"groupCondition,omitempty" xml:"groupCondition,omitempty"`
	MaxResults      *string `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken       *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OrderBy         *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	SearchType      *string `json:"searchType,omitempty" xml:"searchType,omitempty"`
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	SpaceType       *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s ListWorkitemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkitemsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkitemsRequest) SetCategory(v string) *ListWorkitemsRequest {
	s.Category = &v
	return s
}

func (s *ListWorkitemsRequest) SetConditions(v string) *ListWorkitemsRequest {
	s.Conditions = &v
	return s
}

func (s *ListWorkitemsRequest) SetExtraConditions(v string) *ListWorkitemsRequest {
	s.ExtraConditions = &v
	return s
}

func (s *ListWorkitemsRequest) SetGroupCondition(v string) *ListWorkitemsRequest {
	s.GroupCondition = &v
	return s
}

func (s *ListWorkitemsRequest) SetMaxResults(v string) *ListWorkitemsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkitemsRequest) SetNextToken(v string) *ListWorkitemsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkitemsRequest) SetOrderBy(v string) *ListWorkitemsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListWorkitemsRequest) SetSearchType(v string) *ListWorkitemsRequest {
	s.SearchType = &v
	return s
}

func (s *ListWorkitemsRequest) SetSpaceIdentifier(v string) *ListWorkitemsRequest {
	s.SpaceIdentifier = &v
	return s
}

func (s *ListWorkitemsRequest) SetSpaceType(v string) *ListWorkitemsRequest {
	s.SpaceType = &v
	return s
}

type ListWorkitemsResponseBody struct {
	ErrorCode  *string                               `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg   *string                               `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	MaxResults *int64                                `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                               `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId  *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success    *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount *int64                                `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	Workitems  []*ListWorkitemsResponseBodyWorkitems `json:"workitems,omitempty" xml:"workitems,omitempty" type:"Repeated"`
}

func (s ListWorkitemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkitemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkitemsResponseBody) SetErrorCode(v string) *ListWorkitemsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListWorkitemsResponseBody) SetErrorMsg(v string) *ListWorkitemsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListWorkitemsResponseBody) SetMaxResults(v int64) *ListWorkitemsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkitemsResponseBody) SetNextToken(v string) *ListWorkitemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkitemsResponseBody) SetRequestId(v string) *ListWorkitemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkitemsResponseBody) SetSuccess(v bool) *ListWorkitemsResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkitemsResponseBody) SetTotalCount(v int64) *ListWorkitemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkitemsResponseBody) SetWorkitems(v []*ListWorkitemsResponseBodyWorkitems) *ListWorkitemsResponseBody {
	s.Workitems = v
	return s
}

type ListWorkitemsResponseBodyWorkitems struct {
	AssignedTo             *string `json:"assignedTo,omitempty" xml:"assignedTo,omitempty"`
	CategoryIdentifier     *string `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	Creator                *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Document               *string `json:"document,omitempty" xml:"document,omitempty"`
	GmtCreate              *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified            *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Identifier             *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	LogicalStatus          *string `json:"logicalStatus,omitempty" xml:"logicalStatus,omitempty"`
	Modifier               *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	ParentIdentifier       *string `json:"parentIdentifier,omitempty" xml:"parentIdentifier,omitempty"`
	SerialNumber           *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	SpaceIdentifier        *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	SpaceName              *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	SpaceType              *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	SprintIdentifier       *string `json:"sprintIdentifier,omitempty" xml:"sprintIdentifier,omitempty"`
	Status                 *string `json:"status,omitempty" xml:"status,omitempty"`
	StatusIdentifier       *string `json:"statusIdentifier,omitempty" xml:"statusIdentifier,omitempty"`
	StatusStageIdentifier  *string `json:"statusStageIdentifier,omitempty" xml:"statusStageIdentifier,omitempty"`
	Subject                *string `json:"subject,omitempty" xml:"subject,omitempty"`
	UpdateStatusAt         *int64  `json:"updateStatusAt,omitempty" xml:"updateStatusAt,omitempty"`
	WorkitemTypeIdentifier *string `json:"workitemTypeIdentifier,omitempty" xml:"workitemTypeIdentifier,omitempty"`
}

func (s ListWorkitemsResponseBodyWorkitems) String() string {
	return tea.Prettify(s)
}

func (s ListWorkitemsResponseBodyWorkitems) GoString() string {
	return s.String()
}

func (s *ListWorkitemsResponseBodyWorkitems) SetAssignedTo(v string) *ListWorkitemsResponseBodyWorkitems {
	s.AssignedTo = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetCategoryIdentifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.CategoryIdentifier = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetCreator(v string) *ListWorkitemsResponseBodyWorkitems {
	s.Creator = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetDocument(v string) *ListWorkitemsResponseBodyWorkitems {
	s.Document = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetGmtCreate(v int64) *ListWorkitemsResponseBodyWorkitems {
	s.GmtCreate = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetGmtModified(v int64) *ListWorkitemsResponseBodyWorkitems {
	s.GmtModified = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetIdentifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.Identifier = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetLogicalStatus(v string) *ListWorkitemsResponseBodyWorkitems {
	s.LogicalStatus = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetModifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.Modifier = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetParentIdentifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.ParentIdentifier = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetSerialNumber(v string) *ListWorkitemsResponseBodyWorkitems {
	s.SerialNumber = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetSpaceIdentifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.SpaceIdentifier = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetSpaceName(v string) *ListWorkitemsResponseBodyWorkitems {
	s.SpaceName = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetSpaceType(v string) *ListWorkitemsResponseBodyWorkitems {
	s.SpaceType = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetSprintIdentifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.SprintIdentifier = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetStatus(v string) *ListWorkitemsResponseBodyWorkitems {
	s.Status = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetStatusIdentifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.StatusIdentifier = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetStatusStageIdentifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.StatusStageIdentifier = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetSubject(v string) *ListWorkitemsResponseBodyWorkitems {
	s.Subject = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetUpdateStatusAt(v int64) *ListWorkitemsResponseBodyWorkitems {
	s.UpdateStatusAt = &v
	return s
}

func (s *ListWorkitemsResponseBodyWorkitems) SetWorkitemTypeIdentifier(v string) *ListWorkitemsResponseBodyWorkitems {
	s.WorkitemTypeIdentifier = &v
	return s
}

type ListWorkitemsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListWorkitemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWorkitemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkitemsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkitemsResponse) SetHeaders(v map[string]*string) *ListWorkitemsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkitemsResponse) SetStatusCode(v int32) *ListWorkitemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkitemsResponse) SetBody(v *ListWorkitemsResponseBody) *ListWorkitemsResponse {
	s.Body = v
	return s
}

type ListWorkspacesRequest struct {
	MaxResults            *int32    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken             *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	StatusList            []*string `json:"statusList,omitempty" xml:"statusList,omitempty" type:"Repeated"`
	WorkspaceTemplateList []*string `json:"workspaceTemplateList,omitempty" xml:"workspaceTemplateList,omitempty" type:"Repeated"`
}

func (s ListWorkspacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) SetMaxResults(v int32) *ListWorkspacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesRequest) SetNextToken(v string) *ListWorkspacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesRequest) SetStatusList(v []*string) *ListWorkspacesRequest {
	s.StatusList = v
	return s
}

func (s *ListWorkspacesRequest) SetWorkspaceTemplateList(v []*string) *ListWorkspacesRequest {
	s.WorkspaceTemplateList = v
	return s
}

type ListWorkspacesShrinkRequest struct {
	MaxResults                  *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken                   *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	StatusListShrink            *string `json:"statusList,omitempty" xml:"statusList,omitempty"`
	WorkspaceTemplateListShrink *string `json:"workspaceTemplateList,omitempty" xml:"workspaceTemplateList,omitempty"`
}

func (s ListWorkspacesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesShrinkRequest) SetMaxResults(v int32) *ListWorkspacesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetNextToken(v string) *ListWorkspacesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetStatusListShrink(v string) *ListWorkspacesShrinkRequest {
	s.StatusListShrink = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetWorkspaceTemplateListShrink(v string) *ListWorkspacesShrinkRequest {
	s.WorkspaceTemplateListShrink = &v
	return s
}

type ListWorkspacesResponseBody struct {
	ErrorCode    *string                                 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                 `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	MaxResults   *int32                                  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken    *string                                 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId    *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount   *int32                                  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	Workspaces   []*ListWorkspacesResponseBodyWorkspaces `json:"workspaces,omitempty" xml:"workspaces,omitempty" type:"Repeated"`
}

func (s ListWorkspacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) SetErrorCode(v string) *ListWorkspacesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetErrorMessage(v string) *ListWorkspacesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetMaxResults(v int32) *ListWorkspacesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetNextToken(v string) *ListWorkspacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetRequestId(v string) *ListWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetSuccess(v bool) *ListWorkspacesResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetTotalCount(v int32) *ListWorkspacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetWorkspaces(v []*ListWorkspacesResponseBodyWorkspaces) *ListWorkspacesResponseBody {
	s.Workspaces = v
	return s
}

type ListWorkspacesResponseBodyWorkspaces struct {
	CodeUrl     *string `json:"codeUrl,omitempty" xml:"codeUrl,omitempty"`
	CodeVersion *string `json:"codeVersion,omitempty" xml:"codeVersion,omitempty"`
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Spec        *string `json:"spec,omitempty" xml:"spec,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
	Template    *string `json:"template,omitempty" xml:"template,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspaces) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspaces) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCodeUrl(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.CodeUrl = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCodeVersion(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.CodeVersion = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCreateTime(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Id = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetName(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Name = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetSpec(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Spec = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetStatus(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Status = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetTemplate(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Template = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetUserId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.UserId = &v
	return s
}

type ListWorkspacesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWorkspacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponse) SetHeaders(v map[string]*string) *ListWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspacesResponse) SetStatusCode(v int32) *ListWorkspacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspacesResponse) SetBody(v *ListWorkspacesResponseBody) *ListWorkspacesResponse {
	s.Body = v
	return s
}

type LogPipelineJobRunResponseBody struct {
	ErrorCode    *string                           `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                           `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Log          *LogPipelineJobRunResponseBodyLog `json:"log,omitempty" xml:"log,omitempty" type:"Struct"`
	RequestId    *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s LogPipelineJobRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LogPipelineJobRunResponseBody) GoString() string {
	return s.String()
}

func (s *LogPipelineJobRunResponseBody) SetErrorCode(v string) *LogPipelineJobRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *LogPipelineJobRunResponseBody) SetErrorMessage(v string) *LogPipelineJobRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *LogPipelineJobRunResponseBody) SetLog(v *LogPipelineJobRunResponseBodyLog) *LogPipelineJobRunResponseBody {
	s.Log = v
	return s
}

func (s *LogPipelineJobRunResponseBody) SetRequestId(v string) *LogPipelineJobRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *LogPipelineJobRunResponseBody) SetSuccess(v bool) *LogPipelineJobRunResponseBody {
	s.Success = &v
	return s
}

type LogPipelineJobRunResponseBodyLog struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	More    *bool   `json:"more,omitempty" xml:"more,omitempty"`
}

func (s LogPipelineJobRunResponseBodyLog) String() string {
	return tea.Prettify(s)
}

func (s LogPipelineJobRunResponseBodyLog) GoString() string {
	return s.String()
}

func (s *LogPipelineJobRunResponseBodyLog) SetContent(v string) *LogPipelineJobRunResponseBodyLog {
	s.Content = &v
	return s
}

func (s *LogPipelineJobRunResponseBodyLog) SetMore(v bool) *LogPipelineJobRunResponseBodyLog {
	s.More = &v
	return s
}

type LogPipelineJobRunResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LogPipelineJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LogPipelineJobRunResponse) String() string {
	return tea.Prettify(s)
}

func (s LogPipelineJobRunResponse) GoString() string {
	return s.String()
}

func (s *LogPipelineJobRunResponse) SetHeaders(v map[string]*string) *LogPipelineJobRunResponse {
	s.Headers = v
	return s
}

func (s *LogPipelineJobRunResponse) SetStatusCode(v int32) *LogPipelineJobRunResponse {
	s.StatusCode = &v
	return s
}

func (s *LogPipelineJobRunResponse) SetBody(v *LogPipelineJobRunResponseBody) *LogPipelineJobRunResponse {
	s.Body = v
	return s
}

type LogVMDeployMachineResponseBody struct {
	DeployMachineLog *LogVMDeployMachineResponseBodyDeployMachineLog `json:"deployMachineLog,omitempty" xml:"deployMachineLog,omitempty" type:"Struct"`
	ErrorCode        *string                                         `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage     *string                                         `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId        *string                                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success          *bool                                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s LogVMDeployMachineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LogVMDeployMachineResponseBody) GoString() string {
	return s.String()
}

func (s *LogVMDeployMachineResponseBody) SetDeployMachineLog(v *LogVMDeployMachineResponseBodyDeployMachineLog) *LogVMDeployMachineResponseBody {
	s.DeployMachineLog = v
	return s
}

func (s *LogVMDeployMachineResponseBody) SetErrorCode(v string) *LogVMDeployMachineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *LogVMDeployMachineResponseBody) SetErrorMessage(v string) *LogVMDeployMachineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *LogVMDeployMachineResponseBody) SetRequestId(v string) *LogVMDeployMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *LogVMDeployMachineResponseBody) SetSuccess(v bool) *LogVMDeployMachineResponseBody {
	s.Success = &v
	return s
}

type LogVMDeployMachineResponseBodyDeployMachineLog struct {
	AliyunRegion    *string `json:"aliyunRegion,omitempty" xml:"aliyunRegion,omitempty"`
	DeployBeginTime *int64  `json:"deployBeginTime,omitempty" xml:"deployBeginTime,omitempty"`
	DeployEndTime   *int64  `json:"deployEndTime,omitempty" xml:"deployEndTime,omitempty"`
	DeployLog       *string `json:"deployLog,omitempty" xml:"deployLog,omitempty"`
	DeployLogPath   *string `json:"deployLogPath,omitempty" xml:"deployLogPath,omitempty"`
}

func (s LogVMDeployMachineResponseBodyDeployMachineLog) String() string {
	return tea.Prettify(s)
}

func (s LogVMDeployMachineResponseBodyDeployMachineLog) GoString() string {
	return s.String()
}

func (s *LogVMDeployMachineResponseBodyDeployMachineLog) SetAliyunRegion(v string) *LogVMDeployMachineResponseBodyDeployMachineLog {
	s.AliyunRegion = &v
	return s
}

func (s *LogVMDeployMachineResponseBodyDeployMachineLog) SetDeployBeginTime(v int64) *LogVMDeployMachineResponseBodyDeployMachineLog {
	s.DeployBeginTime = &v
	return s
}

func (s *LogVMDeployMachineResponseBodyDeployMachineLog) SetDeployEndTime(v int64) *LogVMDeployMachineResponseBodyDeployMachineLog {
	s.DeployEndTime = &v
	return s
}

func (s *LogVMDeployMachineResponseBodyDeployMachineLog) SetDeployLog(v string) *LogVMDeployMachineResponseBodyDeployMachineLog {
	s.DeployLog = &v
	return s
}

func (s *LogVMDeployMachineResponseBodyDeployMachineLog) SetDeployLogPath(v string) *LogVMDeployMachineResponseBodyDeployMachineLog {
	s.DeployLogPath = &v
	return s
}

type LogVMDeployMachineResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LogVMDeployMachineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LogVMDeployMachineResponse) String() string {
	return tea.Prettify(s)
}

func (s LogVMDeployMachineResponse) GoString() string {
	return s.String()
}

func (s *LogVMDeployMachineResponse) SetHeaders(v map[string]*string) *LogVMDeployMachineResponse {
	s.Headers = v
	return s
}

func (s *LogVMDeployMachineResponse) SetStatusCode(v int32) *LogVMDeployMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *LogVMDeployMachineResponse) SetBody(v *LogVMDeployMachineResponseBody) *LogVMDeployMachineResponse {
	s.Body = v
	return s
}

type PassPipelineValidateResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PassPipelineValidateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PassPipelineValidateResponseBody) GoString() string {
	return s.String()
}

func (s *PassPipelineValidateResponseBody) SetErrorCode(v string) *PassPipelineValidateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PassPipelineValidateResponseBody) SetErrorMessage(v string) *PassPipelineValidateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *PassPipelineValidateResponseBody) SetRequestId(v string) *PassPipelineValidateResponseBody {
	s.RequestId = &v
	return s
}

func (s *PassPipelineValidateResponseBody) SetSuccess(v bool) *PassPipelineValidateResponseBody {
	s.Success = &v
	return s
}

type PassPipelineValidateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PassPipelineValidateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PassPipelineValidateResponse) String() string {
	return tea.Prettify(s)
}

func (s PassPipelineValidateResponse) GoString() string {
	return s.String()
}

func (s *PassPipelineValidateResponse) SetHeaders(v map[string]*string) *PassPipelineValidateResponse {
	s.Headers = v
	return s
}

func (s *PassPipelineValidateResponse) SetStatusCode(v int32) *PassPipelineValidateResponse {
	s.StatusCode = &v
	return s
}

func (s *PassPipelineValidateResponse) SetBody(v *PassPipelineValidateResponseBody) *PassPipelineValidateResponse {
	s.Body = v
	return s
}

type RefusePipelineValidateResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RefusePipelineValidateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefusePipelineValidateResponseBody) GoString() string {
	return s.String()
}

func (s *RefusePipelineValidateResponseBody) SetErrorCode(v string) *RefusePipelineValidateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RefusePipelineValidateResponseBody) SetErrorMessage(v string) *RefusePipelineValidateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RefusePipelineValidateResponseBody) SetRequestId(v string) *RefusePipelineValidateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefusePipelineValidateResponseBody) SetSuccess(v bool) *RefusePipelineValidateResponseBody {
	s.Success = &v
	return s
}

type RefusePipelineValidateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefusePipelineValidateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefusePipelineValidateResponse) String() string {
	return tea.Prettify(s)
}

func (s RefusePipelineValidateResponse) GoString() string {
	return s.String()
}

func (s *RefusePipelineValidateResponse) SetHeaders(v map[string]*string) *RefusePipelineValidateResponse {
	s.Headers = v
	return s
}

func (s *RefusePipelineValidateResponse) SetStatusCode(v int32) *RefusePipelineValidateResponse {
	s.StatusCode = &v
	return s
}

func (s *RefusePipelineValidateResponse) SetBody(v *RefusePipelineValidateResponseBody) *RefusePipelineValidateResponse {
	s.Body = v
	return s
}

type ReleaseWorkspaceResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ReleaseWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseWorkspaceResponseBody) SetErrorCode(v string) *ReleaseWorkspaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReleaseWorkspaceResponseBody) SetErrorMessage(v string) *ReleaseWorkspaceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ReleaseWorkspaceResponseBody) SetRequestId(v string) *ReleaseWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseWorkspaceResponseBody) SetSuccess(v bool) *ReleaseWorkspaceResponseBody {
	s.Success = &v
	return s
}

type ReleaseWorkspaceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleaseWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseWorkspaceResponse) SetHeaders(v map[string]*string) *ReleaseWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseWorkspaceResponse) SetStatusCode(v int32) *ReleaseWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseWorkspaceResponse) SetBody(v *ReleaseWorkspaceResponseBody) *ReleaseWorkspaceResponse {
	s.Body = v
	return s
}

type ResetSshKeyResponseBody struct {
	ErrorCode    *string                        `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                        `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SshKey       *ResetSshKeyResponseBodySshKey `json:"sshKey,omitempty" xml:"sshKey,omitempty" type:"Struct"`
	Success      *bool                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ResetSshKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetSshKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ResetSshKeyResponseBody) SetErrorCode(v string) *ResetSshKeyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ResetSshKeyResponseBody) SetErrorMessage(v string) *ResetSshKeyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ResetSshKeyResponseBody) SetRequestId(v string) *ResetSshKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetSshKeyResponseBody) SetSshKey(v *ResetSshKeyResponseBodySshKey) *ResetSshKeyResponseBody {
	s.SshKey = v
	return s
}

func (s *ResetSshKeyResponseBody) SetSuccess(v bool) *ResetSshKeyResponseBody {
	s.Success = &v
	return s
}

type ResetSshKeyResponseBodySshKey struct {
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	PublicKey *string `json:"publicKey,omitempty" xml:"publicKey,omitempty"`
}

func (s ResetSshKeyResponseBodySshKey) String() string {
	return tea.Prettify(s)
}

func (s ResetSshKeyResponseBodySshKey) GoString() string {
	return s.String()
}

func (s *ResetSshKeyResponseBodySshKey) SetId(v int64) *ResetSshKeyResponseBodySshKey {
	s.Id = &v
	return s
}

func (s *ResetSshKeyResponseBodySshKey) SetPublicKey(v string) *ResetSshKeyResponseBodySshKey {
	s.PublicKey = &v
	return s
}

type ResetSshKeyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetSshKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetSshKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetSshKeyResponse) GoString() string {
	return s.String()
}

func (s *ResetSshKeyResponse) SetHeaders(v map[string]*string) *ResetSshKeyResponse {
	s.Headers = v
	return s
}

func (s *ResetSshKeyResponse) SetStatusCode(v int32) *ResetSshKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetSshKeyResponse) SetBody(v *ResetSshKeyResponseBody) *ResetSshKeyResponse {
	s.Body = v
	return s
}

type ResumeVMDeployOrderResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ResumeVMDeployOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeVMDeployOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeVMDeployOrderResponseBody) SetErrorCode(v string) *ResumeVMDeployOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ResumeVMDeployOrderResponseBody) SetErrorMessage(v string) *ResumeVMDeployOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ResumeVMDeployOrderResponseBody) SetRequestId(v string) *ResumeVMDeployOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeVMDeployOrderResponseBody) SetSuccess(v bool) *ResumeVMDeployOrderResponseBody {
	s.Success = &v
	return s
}

type ResumeVMDeployOrderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResumeVMDeployOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeVMDeployOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeVMDeployOrderResponse) GoString() string {
	return s.String()
}

func (s *ResumeVMDeployOrderResponse) SetHeaders(v map[string]*string) *ResumeVMDeployOrderResponse {
	s.Headers = v
	return s
}

func (s *ResumeVMDeployOrderResponse) SetStatusCode(v int32) *ResumeVMDeployOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeVMDeployOrderResponse) SetBody(v *ResumeVMDeployOrderResponseBody) *ResumeVMDeployOrderResponse {
	s.Body = v
	return s
}

type RetryPipelineJobRunResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RetryPipelineJobRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetryPipelineJobRunResponseBody) GoString() string {
	return s.String()
}

func (s *RetryPipelineJobRunResponseBody) SetErrorCode(v string) *RetryPipelineJobRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RetryPipelineJobRunResponseBody) SetErrorMessage(v string) *RetryPipelineJobRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RetryPipelineJobRunResponseBody) SetRequestId(v string) *RetryPipelineJobRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryPipelineJobRunResponseBody) SetSuccess(v bool) *RetryPipelineJobRunResponseBody {
	s.Success = &v
	return s
}

type RetryPipelineJobRunResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RetryPipelineJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RetryPipelineJobRunResponse) String() string {
	return tea.Prettify(s)
}

func (s RetryPipelineJobRunResponse) GoString() string {
	return s.String()
}

func (s *RetryPipelineJobRunResponse) SetHeaders(v map[string]*string) *RetryPipelineJobRunResponse {
	s.Headers = v
	return s
}

func (s *RetryPipelineJobRunResponse) SetStatusCode(v int32) *RetryPipelineJobRunResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryPipelineJobRunResponse) SetBody(v *RetryPipelineJobRunResponseBody) *RetryPipelineJobRunResponse {
	s.Body = v
	return s
}

type RetryVMDeployMachineResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RetryVMDeployMachineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetryVMDeployMachineResponseBody) GoString() string {
	return s.String()
}

func (s *RetryVMDeployMachineResponseBody) SetErrorCode(v string) *RetryVMDeployMachineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RetryVMDeployMachineResponseBody) SetErrorMessage(v string) *RetryVMDeployMachineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RetryVMDeployMachineResponseBody) SetRequestId(v string) *RetryVMDeployMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryVMDeployMachineResponseBody) SetSuccess(v bool) *RetryVMDeployMachineResponseBody {
	s.Success = &v
	return s
}

type RetryVMDeployMachineResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RetryVMDeployMachineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RetryVMDeployMachineResponse) String() string {
	return tea.Prettify(s)
}

func (s RetryVMDeployMachineResponse) GoString() string {
	return s.String()
}

func (s *RetryVMDeployMachineResponse) SetHeaders(v map[string]*string) *RetryVMDeployMachineResponse {
	s.Headers = v
	return s
}

func (s *RetryVMDeployMachineResponse) SetStatusCode(v int32) *RetryVMDeployMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryVMDeployMachineResponse) SetBody(v *RetryVMDeployMachineResponseBody) *RetryVMDeployMachineResponse {
	s.Body = v
	return s
}

type SkipPipelineJobRunResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SkipPipelineJobRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SkipPipelineJobRunResponseBody) GoString() string {
	return s.String()
}

func (s *SkipPipelineJobRunResponseBody) SetErrorCode(v string) *SkipPipelineJobRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SkipPipelineJobRunResponseBody) SetErrorMessage(v string) *SkipPipelineJobRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SkipPipelineJobRunResponseBody) SetRequestId(v string) *SkipPipelineJobRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *SkipPipelineJobRunResponseBody) SetSuccess(v bool) *SkipPipelineJobRunResponseBody {
	s.Success = &v
	return s
}

type SkipPipelineJobRunResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SkipPipelineJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SkipPipelineJobRunResponse) String() string {
	return tea.Prettify(s)
}

func (s SkipPipelineJobRunResponse) GoString() string {
	return s.String()
}

func (s *SkipPipelineJobRunResponse) SetHeaders(v map[string]*string) *SkipPipelineJobRunResponse {
	s.Headers = v
	return s
}

func (s *SkipPipelineJobRunResponse) SetStatusCode(v int32) *SkipPipelineJobRunResponse {
	s.StatusCode = &v
	return s
}

func (s *SkipPipelineJobRunResponse) SetBody(v *SkipPipelineJobRunResponseBody) *SkipPipelineJobRunResponse {
	s.Body = v
	return s
}

type SkipVMDeployMachineResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SkipVMDeployMachineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SkipVMDeployMachineResponseBody) GoString() string {
	return s.String()
}

func (s *SkipVMDeployMachineResponseBody) SetErrorCode(v string) *SkipVMDeployMachineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SkipVMDeployMachineResponseBody) SetErrorMessage(v string) *SkipVMDeployMachineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SkipVMDeployMachineResponseBody) SetRequestId(v string) *SkipVMDeployMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *SkipVMDeployMachineResponseBody) SetSuccess(v bool) *SkipVMDeployMachineResponseBody {
	s.Success = &v
	return s
}

type SkipVMDeployMachineResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SkipVMDeployMachineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SkipVMDeployMachineResponse) String() string {
	return tea.Prettify(s)
}

func (s SkipVMDeployMachineResponse) GoString() string {
	return s.String()
}

func (s *SkipVMDeployMachineResponse) SetHeaders(v map[string]*string) *SkipVMDeployMachineResponse {
	s.Headers = v
	return s
}

func (s *SkipVMDeployMachineResponse) SetStatusCode(v int32) *SkipVMDeployMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *SkipVMDeployMachineResponse) SetBody(v *SkipVMDeployMachineResponseBody) *SkipVMDeployMachineResponse {
	s.Body = v
	return s
}

type StartPipelineRunRequest struct {
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
}

func (s StartPipelineRunRequest) String() string {
	return tea.Prettify(s)
}

func (s StartPipelineRunRequest) GoString() string {
	return s.String()
}

func (s *StartPipelineRunRequest) SetParams(v string) *StartPipelineRunRequest {
	s.Params = &v
	return s
}

type StartPipelineRunResponseBody struct {
	ErrorCode     *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage  *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	PipelineRunId *int64  `json:"pipelineRunId,omitempty" xml:"pipelineRunId,omitempty"`
	RequestId     *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success       *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StartPipelineRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartPipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *StartPipelineRunResponseBody) SetErrorCode(v string) *StartPipelineRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartPipelineRunResponseBody) SetErrorMessage(v string) *StartPipelineRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StartPipelineRunResponseBody) SetPipelineRunId(v int64) *StartPipelineRunResponseBody {
	s.PipelineRunId = &v
	return s
}

func (s *StartPipelineRunResponseBody) SetRequestId(v string) *StartPipelineRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartPipelineRunResponseBody) SetSuccess(v bool) *StartPipelineRunResponseBody {
	s.Success = &v
	return s
}

type StartPipelineRunResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartPipelineRunResponse) String() string {
	return tea.Prettify(s)
}

func (s StartPipelineRunResponse) GoString() string {
	return s.String()
}

func (s *StartPipelineRunResponse) SetHeaders(v map[string]*string) *StartPipelineRunResponse {
	s.Headers = v
	return s
}

func (s *StartPipelineRunResponse) SetStatusCode(v int32) *StartPipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *StartPipelineRunResponse) SetBody(v *StartPipelineRunResponseBody) *StartPipelineRunResponse {
	s.Body = v
	return s
}

type StopPipelineJobRunResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StopPipelineJobRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopPipelineJobRunResponseBody) GoString() string {
	return s.String()
}

func (s *StopPipelineJobRunResponseBody) SetErrorCode(v string) *StopPipelineJobRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopPipelineJobRunResponseBody) SetErrorMessage(v string) *StopPipelineJobRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopPipelineJobRunResponseBody) SetRequestId(v string) *StopPipelineJobRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopPipelineJobRunResponseBody) SetSuccess(v bool) *StopPipelineJobRunResponseBody {
	s.Success = &v
	return s
}

type StopPipelineJobRunResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopPipelineJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopPipelineJobRunResponse) String() string {
	return tea.Prettify(s)
}

func (s StopPipelineJobRunResponse) GoString() string {
	return s.String()
}

func (s *StopPipelineJobRunResponse) SetHeaders(v map[string]*string) *StopPipelineJobRunResponse {
	s.Headers = v
	return s
}

func (s *StopPipelineJobRunResponse) SetStatusCode(v int32) *StopPipelineJobRunResponse {
	s.StatusCode = &v
	return s
}

func (s *StopPipelineJobRunResponse) SetBody(v *StopPipelineJobRunResponseBody) *StopPipelineJobRunResponse {
	s.Body = v
	return s
}

type StopPipelineRunResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StopPipelineRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopPipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *StopPipelineRunResponseBody) SetErrorCode(v string) *StopPipelineRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopPipelineRunResponseBody) SetErrorMessage(v string) *StopPipelineRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopPipelineRunResponseBody) SetRequestId(v string) *StopPipelineRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopPipelineRunResponseBody) SetSuccess(v bool) *StopPipelineRunResponseBody {
	s.Success = &v
	return s
}

type StopPipelineRunResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopPipelineRunResponse) String() string {
	return tea.Prettify(s)
}

func (s StopPipelineRunResponse) GoString() string {
	return s.String()
}

func (s *StopPipelineRunResponse) SetHeaders(v map[string]*string) *StopPipelineRunResponse {
	s.Headers = v
	return s
}

func (s *StopPipelineRunResponse) SetStatusCode(v int32) *StopPipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *StopPipelineRunResponse) SetBody(v *StopPipelineRunResponseBody) *StopPipelineRunResponse {
	s.Body = v
	return s
}

type StopVMDeployOrderResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StopVMDeployOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopVMDeployOrderResponseBody) GoString() string {
	return s.String()
}

func (s *StopVMDeployOrderResponseBody) SetErrorCode(v string) *StopVMDeployOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopVMDeployOrderResponseBody) SetErrorMessage(v string) *StopVMDeployOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopVMDeployOrderResponseBody) SetRequestId(v string) *StopVMDeployOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopVMDeployOrderResponseBody) SetSuccess(v bool) *StopVMDeployOrderResponseBody {
	s.Success = &v
	return s
}

type StopVMDeployOrderResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopVMDeployOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopVMDeployOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s StopVMDeployOrderResponse) GoString() string {
	return s.String()
}

func (s *StopVMDeployOrderResponse) SetHeaders(v map[string]*string) *StopVMDeployOrderResponse {
	s.Headers = v
	return s
}

func (s *StopVMDeployOrderResponse) SetStatusCode(v int32) *StopVMDeployOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *StopVMDeployOrderResponse) SetBody(v *StopVMDeployOrderResponseBody) *StopVMDeployOrderResponse {
	s.Body = v
	return s
}

type TriggerRepositoryMirrorSyncRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	Account        *string `json:"account,omitempty" xml:"account,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	Token          *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s TriggerRepositoryMirrorSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s TriggerRepositoryMirrorSyncRequest) GoString() string {
	return s.String()
}

func (s *TriggerRepositoryMirrorSyncRequest) SetAccessToken(v string) *TriggerRepositoryMirrorSyncRequest {
	s.AccessToken = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncRequest) SetAccount(v string) *TriggerRepositoryMirrorSyncRequest {
	s.Account = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncRequest) SetOrganizationId(v string) *TriggerRepositoryMirrorSyncRequest {
	s.OrganizationId = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncRequest) SetToken(v string) *TriggerRepositoryMirrorSyncRequest {
	s.Token = &v
	return s
}

type TriggerRepositoryMirrorSyncResponseBody struct {
	ErrorCode    *string                                        `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                        `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       *TriggerRepositoryMirrorSyncResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success      *bool                                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TriggerRepositoryMirrorSyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerRepositoryMirrorSyncResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerRepositoryMirrorSyncResponseBody) SetErrorCode(v string) *TriggerRepositoryMirrorSyncResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponseBody) SetErrorMessage(v string) *TriggerRepositoryMirrorSyncResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponseBody) SetRequestId(v string) *TriggerRepositoryMirrorSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponseBody) SetResult(v *TriggerRepositoryMirrorSyncResponseBodyResult) *TriggerRepositoryMirrorSyncResponseBody {
	s.Result = v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponseBody) SetSuccess(v bool) *TriggerRepositoryMirrorSyncResponseBody {
	s.Success = &v
	return s
}

type TriggerRepositoryMirrorSyncResponseBodyResult struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s TriggerRepositoryMirrorSyncResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s TriggerRepositoryMirrorSyncResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TriggerRepositoryMirrorSyncResponseBodyResult) SetResult(v bool) *TriggerRepositoryMirrorSyncResponseBodyResult {
	s.Result = &v
	return s
}

type TriggerRepositoryMirrorSyncResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TriggerRepositoryMirrorSyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TriggerRepositoryMirrorSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s TriggerRepositoryMirrorSyncResponse) GoString() string {
	return s.String()
}

func (s *TriggerRepositoryMirrorSyncResponse) SetHeaders(v map[string]*string) *TriggerRepositoryMirrorSyncResponse {
	s.Headers = v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponse) SetStatusCode(v int32) *TriggerRepositoryMirrorSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponse) SetBody(v *TriggerRepositoryMirrorSyncResponseBody) *TriggerRepositoryMirrorSyncResponse {
	s.Body = v
	return s
}

type UpdateFileRequest struct {
	AccessToken    *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	BranchName     *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	CommitMessage  *string `json:"commitMessage,omitempty" xml:"commitMessage,omitempty"`
	Content        *string `json:"content,omitempty" xml:"content,omitempty"`
	Encoding       *string `json:"encoding,omitempty" xml:"encoding,omitempty"`
	NewPath        *string `json:"newPath,omitempty" xml:"newPath,omitempty"`
	OldPath        *string `json:"oldPath,omitempty" xml:"oldPath,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s UpdateFileRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileRequest) SetAccessToken(v string) *UpdateFileRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateFileRequest) SetBranchName(v string) *UpdateFileRequest {
	s.BranchName = &v
	return s
}

func (s *UpdateFileRequest) SetCommitMessage(v string) *UpdateFileRequest {
	s.CommitMessage = &v
	return s
}

func (s *UpdateFileRequest) SetContent(v string) *UpdateFileRequest {
	s.Content = &v
	return s
}

func (s *UpdateFileRequest) SetEncoding(v string) *UpdateFileRequest {
	s.Encoding = &v
	return s
}

func (s *UpdateFileRequest) SetNewPath(v string) *UpdateFileRequest {
	s.NewPath = &v
	return s
}

func (s *UpdateFileRequest) SetOldPath(v string) *UpdateFileRequest {
	s.OldPath = &v
	return s
}

func (s *UpdateFileRequest) SetOrganizationId(v string) *UpdateFileRequest {
	s.OrganizationId = &v
	return s
}

type UpdateFileResponseBody struct {
	ErrorCode    *string                       `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                       `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       *UpdateFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success      *bool                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileResponseBody) SetErrorCode(v string) *UpdateFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateFileResponseBody) SetErrorMessage(v string) *UpdateFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateFileResponseBody) SetRequestId(v string) *UpdateFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileResponseBody) SetResult(v *UpdateFileResponseBodyResult) *UpdateFileResponseBody {
	s.Result = v
	return s
}

func (s *UpdateFileResponseBody) SetSuccess(v bool) *UpdateFileResponseBody {
	s.Success = &v
	return s
}

type UpdateFileResponseBodyResult struct {
	BranchName *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	FilePath   *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
}

func (s UpdateFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateFileResponseBodyResult) SetBranchName(v string) *UpdateFileResponseBodyResult {
	s.BranchName = &v
	return s
}

func (s *UpdateFileResponseBodyResult) SetFilePath(v string) *UpdateFileResponseBodyResult {
	s.FilePath = &v
	return s
}

type UpdateFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFileResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileResponse) SetHeaders(v map[string]*string) *UpdateFileResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileResponse) SetStatusCode(v int32) *UpdateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileResponse) SetBody(v *UpdateFileResponseBody) *UpdateFileResponse {
	s.Body = v
	return s
}

type UpdateFlowTagRequest struct {
	Color          *string `json:"color,omitempty" xml:"color,omitempty"`
	FlowTagGroupId *int64  `json:"flowTagGroupId,omitempty" xml:"flowTagGroupId,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateFlowTagRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowTagRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowTagRequest) SetColor(v string) *UpdateFlowTagRequest {
	s.Color = &v
	return s
}

func (s *UpdateFlowTagRequest) SetFlowTagGroupId(v int64) *UpdateFlowTagRequest {
	s.FlowTagGroupId = &v
	return s
}

func (s *UpdateFlowTagRequest) SetName(v string) *UpdateFlowTagRequest {
	s.Name = &v
	return s
}

type UpdateFlowTagResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateFlowTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowTagResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlowTagResponseBody) SetErrorCode(v string) *UpdateFlowTagResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateFlowTagResponseBody) SetErrorMessage(v string) *UpdateFlowTagResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateFlowTagResponseBody) SetRequestId(v string) *UpdateFlowTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlowTagResponseBody) SetSuccess(v bool) *UpdateFlowTagResponseBody {
	s.Success = &v
	return s
}

type UpdateFlowTagResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateFlowTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFlowTagResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowTagResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlowTagResponse) SetHeaders(v map[string]*string) *UpdateFlowTagResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlowTagResponse) SetStatusCode(v int32) *UpdateFlowTagResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFlowTagResponse) SetBody(v *UpdateFlowTagResponseBody) *UpdateFlowTagResponse {
	s.Body = v
	return s
}

type UpdateFlowTagGroupRequest struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateFlowTagGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowTagGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowTagGroupRequest) SetName(v string) *UpdateFlowTagGroupRequest {
	s.Name = &v
	return s
}

type UpdateFlowTagGroupResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateFlowTagGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowTagGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlowTagGroupResponseBody) SetErrorCode(v string) *UpdateFlowTagGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateFlowTagGroupResponseBody) SetErrorMessage(v string) *UpdateFlowTagGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateFlowTagGroupResponseBody) SetRequestId(v string) *UpdateFlowTagGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlowTagGroupResponseBody) SetSuccess(v bool) *UpdateFlowTagGroupResponseBody {
	s.Success = &v
	return s
}

type UpdateFlowTagGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateFlowTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFlowTagGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowTagGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlowTagGroupResponse) SetHeaders(v map[string]*string) *UpdateFlowTagGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlowTagGroupResponse) SetStatusCode(v int32) *UpdateFlowTagGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFlowTagGroupResponse) SetBody(v *UpdateFlowTagGroupResponseBody) *UpdateFlowTagGroupResponse {
	s.Body = v
	return s
}

type UpdateHostGroupRequest struct {
	AliyunRegion        *string `json:"aliyunRegion,omitempty" xml:"aliyunRegion,omitempty"`
	EcsLabelKey         *string `json:"ecsLabelKey,omitempty" xml:"ecsLabelKey,omitempty"`
	EcsLabelValue       *string `json:"ecsLabelValue,omitempty" xml:"ecsLabelValue,omitempty"`
	EcsType             *string `json:"ecsType,omitempty" xml:"ecsType,omitempty"`
	EnvId               *string `json:"envId,omitempty" xml:"envId,omitempty"`
	MachineInfos        *string `json:"machineInfos,omitempty" xml:"machineInfos,omitempty"`
	Name                *string `json:"name,omitempty" xml:"name,omitempty"`
	ServiceConnectionId *int64  `json:"serviceConnectionId,omitempty" xml:"serviceConnectionId,omitempty"`
	TagIds              *string `json:"tagIds,omitempty" xml:"tagIds,omitempty"`
	Type                *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateHostGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHostGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateHostGroupRequest) SetAliyunRegion(v string) *UpdateHostGroupRequest {
	s.AliyunRegion = &v
	return s
}

func (s *UpdateHostGroupRequest) SetEcsLabelKey(v string) *UpdateHostGroupRequest {
	s.EcsLabelKey = &v
	return s
}

func (s *UpdateHostGroupRequest) SetEcsLabelValue(v string) *UpdateHostGroupRequest {
	s.EcsLabelValue = &v
	return s
}

func (s *UpdateHostGroupRequest) SetEcsType(v string) *UpdateHostGroupRequest {
	s.EcsType = &v
	return s
}

func (s *UpdateHostGroupRequest) SetEnvId(v string) *UpdateHostGroupRequest {
	s.EnvId = &v
	return s
}

func (s *UpdateHostGroupRequest) SetMachineInfos(v string) *UpdateHostGroupRequest {
	s.MachineInfos = &v
	return s
}

func (s *UpdateHostGroupRequest) SetName(v string) *UpdateHostGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateHostGroupRequest) SetServiceConnectionId(v int64) *UpdateHostGroupRequest {
	s.ServiceConnectionId = &v
	return s
}

func (s *UpdateHostGroupRequest) SetTagIds(v string) *UpdateHostGroupRequest {
	s.TagIds = &v
	return s
}

func (s *UpdateHostGroupRequest) SetType(v string) *UpdateHostGroupRequest {
	s.Type = &v
	return s
}

type UpdateHostGroupResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateHostGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHostGroupResponseBody) SetErrorCode(v string) *UpdateHostGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateHostGroupResponseBody) SetErrorMessage(v string) *UpdateHostGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateHostGroupResponseBody) SetRequestId(v string) *UpdateHostGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHostGroupResponseBody) SetSuccess(v bool) *UpdateHostGroupResponseBody {
	s.Success = &v
	return s
}

type UpdateHostGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateHostGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHostGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateHostGroupResponse) SetHeaders(v map[string]*string) *UpdateHostGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateHostGroupResponse) SetStatusCode(v int32) *UpdateHostGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHostGroupResponse) SetBody(v *UpdateHostGroupResponseBody) *UpdateHostGroupResponse {
	s.Body = v
	return s
}

type UpdatePipelineBaseInfoRequest struct {
	EnvId        *int64  `json:"envId,omitempty" xml:"envId,omitempty"`
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	TagList      *string `json:"tagList,omitempty" xml:"tagList,omitempty"`
}

func (s UpdatePipelineBaseInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineBaseInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineBaseInfoRequest) SetEnvId(v int64) *UpdatePipelineBaseInfoRequest {
	s.EnvId = &v
	return s
}

func (s *UpdatePipelineBaseInfoRequest) SetPipelineName(v string) *UpdatePipelineBaseInfoRequest {
	s.PipelineName = &v
	return s
}

func (s *UpdatePipelineBaseInfoRequest) SetTagList(v string) *UpdatePipelineBaseInfoRequest {
	s.TagList = &v
	return s
}

type UpdatePipelineBaseInfoResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdatePipelineBaseInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelineBaseInfoResponseBody) SetErrorCode(v string) *UpdatePipelineBaseInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdatePipelineBaseInfoResponseBody) SetErrorMessage(v string) *UpdatePipelineBaseInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdatePipelineBaseInfoResponseBody) SetRequestId(v string) *UpdatePipelineBaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePipelineBaseInfoResponseBody) SetSuccess(v bool) *UpdatePipelineBaseInfoResponseBody {
	s.Success = &v
	return s
}

type UpdatePipelineBaseInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePipelineBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePipelineBaseInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdatePipelineBaseInfoResponse) SetHeaders(v map[string]*string) *UpdatePipelineBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdatePipelineBaseInfoResponse) SetStatusCode(v int32) *UpdatePipelineBaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePipelineBaseInfoResponse) SetBody(v *UpdatePipelineBaseInfoResponseBody) *UpdatePipelineBaseInfoResponse {
	s.Body = v
	return s
}

type UpdatePipelineGroupRequest struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdatePipelineGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineGroupRequest) SetName(v string) *UpdatePipelineGroupRequest {
	s.Name = &v
	return s
}

type UpdatePipelineGroupResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdatePipelineGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelineGroupResponseBody) SetErrorCode(v string) *UpdatePipelineGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdatePipelineGroupResponseBody) SetErrorMessage(v string) *UpdatePipelineGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdatePipelineGroupResponseBody) SetRequestId(v string) *UpdatePipelineGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePipelineGroupResponseBody) SetSuccess(v bool) *UpdatePipelineGroupResponseBody {
	s.Success = &v
	return s
}

type UpdatePipelineGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePipelineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePipelineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdatePipelineGroupResponse) SetHeaders(v map[string]*string) *UpdatePipelineGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdatePipelineGroupResponse) SetStatusCode(v int32) *UpdatePipelineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePipelineGroupResponse) SetBody(v *UpdatePipelineGroupResponseBody) *UpdatePipelineGroupResponse {
	s.Body = v
	return s
}

type UpdateProjectMemberRequest struct {
	RoleIdentifier   *string `json:"roleIdentifier,omitempty" xml:"roleIdentifier,omitempty"`
	TargetIdentifier *string `json:"targetIdentifier,omitempty" xml:"targetIdentifier,omitempty"`
	TargetType       *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	UserIdentifier   *string `json:"userIdentifier,omitempty" xml:"userIdentifier,omitempty"`
	UserType         *string `json:"userType,omitempty" xml:"userType,omitempty"`
}

func (s UpdateProjectMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberRequest) SetRoleIdentifier(v string) *UpdateProjectMemberRequest {
	s.RoleIdentifier = &v
	return s
}

func (s *UpdateProjectMemberRequest) SetTargetIdentifier(v string) *UpdateProjectMemberRequest {
	s.TargetIdentifier = &v
	return s
}

func (s *UpdateProjectMemberRequest) SetTargetType(v string) *UpdateProjectMemberRequest {
	s.TargetType = &v
	return s
}

func (s *UpdateProjectMemberRequest) SetUserIdentifier(v string) *UpdateProjectMemberRequest {
	s.UserIdentifier = &v
	return s
}

func (s *UpdateProjectMemberRequest) SetUserType(v string) *UpdateProjectMemberRequest {
	s.UserType = &v
	return s
}

type UpdateProjectMemberResponseBody struct {
	ErrorCode *string                                `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Member    *UpdateProjectMemberResponseBodyMember `json:"member,omitempty" xml:"member,omitempty" type:"Struct"`
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateProjectMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberResponseBody) SetErrorCode(v string) *UpdateProjectMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateProjectMemberResponseBody) SetErrorMsg(v string) *UpdateProjectMemberResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateProjectMemberResponseBody) SetMember(v *UpdateProjectMemberResponseBodyMember) *UpdateProjectMemberResponseBody {
	s.Member = v
	return s
}

func (s *UpdateProjectMemberResponseBody) SetRequestId(v string) *UpdateProjectMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectMemberResponseBody) SetSuccess(v bool) *UpdateProjectMemberResponseBody {
	s.Success = &v
	return s
}

type UpdateProjectMemberResponseBodyMember struct {
	GmtCreate        *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified      *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id               *string `json:"id,omitempty" xml:"id,omitempty"`
	RoleIdentifier   *string `json:"roleIdentifier,omitempty" xml:"roleIdentifier,omitempty"`
	TargetIdentifier *string `json:"targetIdentifier,omitempty" xml:"targetIdentifier,omitempty"`
	TargetType       *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	UserIdentifier   *string `json:"userIdentifier,omitempty" xml:"userIdentifier,omitempty"`
	UserType         *string `json:"userType,omitempty" xml:"userType,omitempty"`
}

func (s UpdateProjectMemberResponseBodyMember) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectMemberResponseBodyMember) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberResponseBodyMember) SetGmtCreate(v int64) *UpdateProjectMemberResponseBodyMember {
	s.GmtCreate = &v
	return s
}

func (s *UpdateProjectMemberResponseBodyMember) SetGmtModified(v int64) *UpdateProjectMemberResponseBodyMember {
	s.GmtModified = &v
	return s
}

func (s *UpdateProjectMemberResponseBodyMember) SetId(v string) *UpdateProjectMemberResponseBodyMember {
	s.Id = &v
	return s
}

func (s *UpdateProjectMemberResponseBodyMember) SetRoleIdentifier(v string) *UpdateProjectMemberResponseBodyMember {
	s.RoleIdentifier = &v
	return s
}

func (s *UpdateProjectMemberResponseBodyMember) SetTargetIdentifier(v string) *UpdateProjectMemberResponseBodyMember {
	s.TargetIdentifier = &v
	return s
}

func (s *UpdateProjectMemberResponseBodyMember) SetTargetType(v string) *UpdateProjectMemberResponseBodyMember {
	s.TargetType = &v
	return s
}

func (s *UpdateProjectMemberResponseBodyMember) SetUserIdentifier(v string) *UpdateProjectMemberResponseBodyMember {
	s.UserIdentifier = &v
	return s
}

func (s *UpdateProjectMemberResponseBodyMember) SetUserType(v string) *UpdateProjectMemberResponseBodyMember {
	s.UserType = &v
	return s
}

type UpdateProjectMemberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProjectMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberResponse) SetHeaders(v map[string]*string) *UpdateProjectMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectMemberResponse) SetStatusCode(v int32) *UpdateProjectMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectMemberResponse) SetBody(v *UpdateProjectMemberResponseBody) *UpdateProjectMemberResponse {
	s.Body = v
	return s
}

type UpdateProtectedBranchesRequest struct {
	AccessToken         *string                                            `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	AllowMergeRoles     []*int32                                           `json:"allowMergeRoles,omitempty" xml:"allowMergeRoles,omitempty" type:"Repeated"`
	AllowMergeUserIds   []*int64                                           `json:"allowMergeUserIds,omitempty" xml:"allowMergeUserIds,omitempty" type:"Repeated"`
	AllowPushRoles      []*int32                                           `json:"allowPushRoles,omitempty" xml:"allowPushRoles,omitempty" type:"Repeated"`
	AllowPushUserIds    []*int64                                           `json:"allowPushUserIds,omitempty" xml:"allowPushUserIds,omitempty" type:"Repeated"`
	Branch              *string                                            `json:"branch,omitempty" xml:"branch,omitempty"`
	Id                  *int64                                             `json:"id,omitempty" xml:"id,omitempty"`
	MergeRequestSetting *UpdateProtectedBranchesRequestMergeRequestSetting `json:"mergeRequestSetting,omitempty" xml:"mergeRequestSetting,omitempty" type:"Struct"`
	TestSettingDTO      *UpdateProtectedBranchesRequestTestSettingDTO      `json:"testSettingDTO,omitempty" xml:"testSettingDTO,omitempty" type:"Struct"`
	OrganizationId      *string                                            `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s UpdateProtectedBranchesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesRequest) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesRequest) SetAccessToken(v string) *UpdateProtectedBranchesRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetAllowMergeRoles(v []*int32) *UpdateProtectedBranchesRequest {
	s.AllowMergeRoles = v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetAllowMergeUserIds(v []*int64) *UpdateProtectedBranchesRequest {
	s.AllowMergeUserIds = v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetAllowPushRoles(v []*int32) *UpdateProtectedBranchesRequest {
	s.AllowPushRoles = v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetAllowPushUserIds(v []*int64) *UpdateProtectedBranchesRequest {
	s.AllowPushUserIds = v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetBranch(v string) *UpdateProtectedBranchesRequest {
	s.Branch = &v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetId(v int64) *UpdateProtectedBranchesRequest {
	s.Id = &v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetMergeRequestSetting(v *UpdateProtectedBranchesRequestMergeRequestSetting) *UpdateProtectedBranchesRequest {
	s.MergeRequestSetting = v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetTestSettingDTO(v *UpdateProtectedBranchesRequestTestSettingDTO) *UpdateProtectedBranchesRequest {
	s.TestSettingDTO = v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetOrganizationId(v string) *UpdateProtectedBranchesRequest {
	s.OrganizationId = &v
	return s
}

type UpdateProtectedBranchesRequestMergeRequestSetting struct {
	AllowMergeRequestRoles       []*int32 `json:"allowMergeRequestRoles,omitempty" xml:"allowMergeRequestRoles,omitempty" type:"Repeated"`
	DefaultAssignees             []*int32 `json:"defaultAssignees,omitempty" xml:"defaultAssignees,omitempty" type:"Repeated"`
	IsAllowSelfApproval          *bool    `json:"isAllowSelfApproval,omitempty" xml:"isAllowSelfApproval,omitempty"`
	IsRequireDiscussionProcessed *bool    `json:"isRequireDiscussionProcessed,omitempty" xml:"isRequireDiscussionProcessed,omitempty"`
	IsRequired                   *bool    `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	IsResetApprovalWhenNewPush   *bool    `json:"isResetApprovalWhenNewPush,omitempty" xml:"isResetApprovalWhenNewPush,omitempty"`
	MinimumApproval              *int32   `json:"minimumApproval,omitempty" xml:"minimumApproval,omitempty"`
	MrMode                       *string  `json:"mrMode,omitempty" xml:"mrMode,omitempty"`
	WhiteList                    *string  `json:"whiteList,omitempty" xml:"whiteList,omitempty"`
}

func (s UpdateProtectedBranchesRequestMergeRequestSetting) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesRequestMergeRequestSetting) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetAllowMergeRequestRoles(v []*int32) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.AllowMergeRequestRoles = v
	return s
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetDefaultAssignees(v []*int32) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.DefaultAssignees = v
	return s
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetIsAllowSelfApproval(v bool) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.IsAllowSelfApproval = &v
	return s
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetIsRequireDiscussionProcessed(v bool) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.IsRequireDiscussionProcessed = &v
	return s
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetIsRequired(v bool) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.IsRequired = &v
	return s
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetIsResetApprovalWhenNewPush(v bool) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.IsResetApprovalWhenNewPush = &v
	return s
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetMinimumApproval(v int32) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.MinimumApproval = &v
	return s
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetMrMode(v string) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.MrMode = &v
	return s
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetWhiteList(v string) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.WhiteList = &v
	return s
}

type UpdateProtectedBranchesRequestTestSettingDTO struct {
	CheckConfig             *UpdateProtectedBranchesRequestTestSettingDTOCheckConfig             `json:"checkConfig,omitempty" xml:"checkConfig,omitempty" type:"Struct"`
	CheckTaskQualityConfig  *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig  `json:"checkTaskQualityConfig,omitempty" xml:"checkTaskQualityConfig,omitempty" type:"Struct"`
	CodeGuidelinesDetection *UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection `json:"codeGuidelinesDetection,omitempty" xml:"codeGuidelinesDetection,omitempty" type:"Struct"`
	IsRequired              *bool                                                                `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	SensitiveInfoDetection  *UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection  `json:"sensitiveInfoDetection,omitempty" xml:"sensitiveInfoDetection,omitempty" type:"Struct"`
}

func (s UpdateProtectedBranchesRequestTestSettingDTO) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesRequestTestSettingDTO) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesRequestTestSettingDTO) SetCheckConfig(v *UpdateProtectedBranchesRequestTestSettingDTOCheckConfig) *UpdateProtectedBranchesRequestTestSettingDTO {
	s.CheckConfig = v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTO) SetCheckTaskQualityConfig(v *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) *UpdateProtectedBranchesRequestTestSettingDTO {
	s.CheckTaskQualityConfig = v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTO) SetCodeGuidelinesDetection(v *UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection) *UpdateProtectedBranchesRequestTestSettingDTO {
	s.CodeGuidelinesDetection = v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTO) SetIsRequired(v bool) *UpdateProtectedBranchesRequestTestSettingDTO {
	s.IsRequired = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTO) SetSensitiveInfoDetection(v *UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection) *UpdateProtectedBranchesRequestTestSettingDTO {
	s.SensitiveInfoDetection = v
	return s
}

type UpdateProtectedBranchesRequestTestSettingDTOCheckConfig struct {
	CheckItems []*UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems `json:"checkItems,omitempty" xml:"checkItems,omitempty" type:"Repeated"`
}

func (s UpdateProtectedBranchesRequestTestSettingDTOCheckConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesRequestTestSettingDTOCheckConfig) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckConfig) SetCheckItems(v []*UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems) *UpdateProtectedBranchesRequestTestSettingDTOCheckConfig {
	s.CheckItems = v
	return s
}

type UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems struct {
	IsRequired *bool   `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems) SetIsRequired(v bool) *UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems {
	s.IsRequired = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems) SetName(v string) *UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems {
	s.Name = &v
	return s
}

type UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig struct {
	BizNo    *string `json:"bizNo,omitempty" xml:"bizNo,omitempty"`
	Enabled  *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Message  *string `json:"message,omitempty" xml:"message,omitempty"`
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) SetBizNo(v string) *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig {
	s.BizNo = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) SetEnabled(v bool) *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig {
	s.Enabled = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) SetMessage(v string) *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig {
	s.Message = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) SetTaskName(v string) *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig {
	s.TaskName = &v
	return s
}

type UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection struct {
	Enabled *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection) SetEnabled(v bool) *UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection {
	s.Enabled = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection) SetMessage(v string) *UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection {
	s.Message = &v
	return s
}

type UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection struct {
	Enabled *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection) SetEnabled(v bool) *UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection {
	s.Enabled = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection) SetMessage(v string) *UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection {
	s.Message = &v
	return s
}

type UpdateProtectedBranchesResponseBody struct {
	ErrorCode    *string                                    `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                    `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       *UpdateProtectedBranchesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success      *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateProtectedBranchesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBody) SetErrorCode(v string) *UpdateProtectedBranchesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBody) SetErrorMessage(v string) *UpdateProtectedBranchesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBody) SetRequestId(v string) *UpdateProtectedBranchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBody) SetResult(v *UpdateProtectedBranchesResponseBodyResult) *UpdateProtectedBranchesResponseBody {
	s.Result = v
	return s
}

func (s *UpdateProtectedBranchesResponseBody) SetSuccess(v bool) *UpdateProtectedBranchesResponseBody {
	s.Success = &v
	return s
}

type UpdateProtectedBranchesResponseBodyResult struct {
	AllowMergeRoles     []*int32                                                      `json:"allowMergeRoles,omitempty" xml:"allowMergeRoles,omitempty" type:"Repeated"`
	AllowMergeUserIds   []*int64                                                      `json:"allowMergeUserIds,omitempty" xml:"allowMergeUserIds,omitempty" type:"Repeated"`
	AllowPushRoles      []*int32                                                      `json:"allowPushRoles,omitempty" xml:"allowPushRoles,omitempty" type:"Repeated"`
	AllowPushUserIds    []*int64                                                      `json:"allowPushUserIds,omitempty" xml:"allowPushUserIds,omitempty" type:"Repeated"`
	Branch              *string                                                       `json:"branch,omitempty" xml:"branch,omitempty"`
	Id                  *int64                                                        `json:"id,omitempty" xml:"id,omitempty"`
	MergeRequestSetting *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting `json:"mergeRequestSetting,omitempty" xml:"mergeRequestSetting,omitempty" type:"Struct"`
	TestSettingDTO      *UpdateProtectedBranchesResponseBodyResultTestSettingDTO      `json:"testSettingDTO,omitempty" xml:"testSettingDTO,omitempty" type:"Struct"`
}

func (s UpdateProtectedBranchesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBodyResult) SetAllowMergeRoles(v []*int32) *UpdateProtectedBranchesResponseBodyResult {
	s.AllowMergeRoles = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResult) SetAllowMergeUserIds(v []*int64) *UpdateProtectedBranchesResponseBodyResult {
	s.AllowMergeUserIds = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResult) SetAllowPushRoles(v []*int32) *UpdateProtectedBranchesResponseBodyResult {
	s.AllowPushRoles = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResult) SetAllowPushUserIds(v []*int64) *UpdateProtectedBranchesResponseBodyResult {
	s.AllowPushUserIds = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResult) SetBranch(v string) *UpdateProtectedBranchesResponseBodyResult {
	s.Branch = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResult) SetId(v int64) *UpdateProtectedBranchesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResult) SetMergeRequestSetting(v *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) *UpdateProtectedBranchesResponseBodyResult {
	s.MergeRequestSetting = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResult) SetTestSettingDTO(v *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) *UpdateProtectedBranchesResponseBodyResult {
	s.TestSettingDTO = v
	return s
}

type UpdateProtectedBranchesResponseBodyResultMergeRequestSetting struct {
	AllowMergeRequestRoles       []*int32 `json:"allowMergeRequestRoles,omitempty" xml:"allowMergeRequestRoles,omitempty" type:"Repeated"`
	DefaultAssignees             []*int32 `json:"defaultAssignees,omitempty" xml:"defaultAssignees,omitempty" type:"Repeated"`
	IsAllowSelfApproval          *bool    `json:"isAllowSelfApproval,omitempty" xml:"isAllowSelfApproval,omitempty"`
	IsRequireDiscussionProcessed *bool    `json:"isRequireDiscussionProcessed,omitempty" xml:"isRequireDiscussionProcessed,omitempty"`
	IsRequired                   *bool    `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	IsResetApprovalWhenNewPush   *bool    `json:"isResetApprovalWhenNewPush,omitempty" xml:"isResetApprovalWhenNewPush,omitempty"`
	MinimumApproval              *int32   `json:"minimumApproval,omitempty" xml:"minimumApproval,omitempty"`
	MrMode                       *string  `json:"mrMode,omitempty" xml:"mrMode,omitempty"`
	WhiteList                    *string  `json:"whiteList,omitempty" xml:"whiteList,omitempty"`
}

func (s UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetAllowMergeRequestRoles(v []*int32) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.AllowMergeRequestRoles = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetDefaultAssignees(v []*int32) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.DefaultAssignees = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetIsAllowSelfApproval(v bool) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.IsAllowSelfApproval = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetIsRequireDiscussionProcessed(v bool) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.IsRequireDiscussionProcessed = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetIsRequired(v bool) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.IsRequired = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetIsResetApprovalWhenNewPush(v bool) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.IsResetApprovalWhenNewPush = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetMinimumApproval(v int32) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.MinimumApproval = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetMrMode(v string) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.MrMode = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetWhiteList(v string) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.WhiteList = &v
	return s
}

type UpdateProtectedBranchesResponseBodyResultTestSettingDTO struct {
	CheckConfig             *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig             `json:"checkConfig,omitempty" xml:"checkConfig,omitempty" type:"Struct"`
	CheckTaskQualityConfig  *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig  `json:"checkTaskQualityConfig,omitempty" xml:"checkTaskQualityConfig,omitempty" type:"Struct"`
	CodeGuidelinesDetection *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection `json:"codeGuidelinesDetection,omitempty" xml:"codeGuidelinesDetection,omitempty" type:"Struct"`
	IsRequired              *bool                                                                           `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	SensitiveInfoDetection  *UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection  `json:"sensitiveInfoDetection,omitempty" xml:"sensitiveInfoDetection,omitempty" type:"Struct"`
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTO) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTO) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) SetCheckConfig(v *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) *UpdateProtectedBranchesResponseBodyResultTestSettingDTO {
	s.CheckConfig = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) SetCheckTaskQualityConfig(v *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) *UpdateProtectedBranchesResponseBodyResultTestSettingDTO {
	s.CheckTaskQualityConfig = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) SetCodeGuidelinesDetection(v *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) *UpdateProtectedBranchesResponseBodyResultTestSettingDTO {
	s.CodeGuidelinesDetection = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) SetIsRequired(v bool) *UpdateProtectedBranchesResponseBodyResultTestSettingDTO {
	s.IsRequired = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) SetSensitiveInfoDetection(v *UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) *UpdateProtectedBranchesResponseBodyResultTestSettingDTO {
	s.SensitiveInfoDetection = v
	return s
}

type UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig struct {
	CheckItems []*UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems `json:"checkItems,omitempty" xml:"checkItems,omitempty" type:"Repeated"`
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) SetCheckItems(v []*UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig {
	s.CheckItems = v
	return s
}

type UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems struct {
	IsRequired *bool   `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) SetIsRequired(v bool) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems {
	s.IsRequired = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) SetName(v string) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems {
	s.Name = &v
	return s
}

type UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig struct {
	BizNo    *string `json:"bizNo,omitempty" xml:"bizNo,omitempty"`
	Enabled  *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Message  *string `json:"message,omitempty" xml:"message,omitempty"`
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetBizNo(v string) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.BizNo = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetEnabled(v bool) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.Enabled = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetMessage(v string) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.Message = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetTaskName(v string) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.TaskName = &v
	return s
}

type UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection struct {
	Enabled *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) SetEnabled(v bool) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection {
	s.Enabled = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) SetMessage(v string) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection {
	s.Message = &v
	return s
}

type UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection struct {
	Enabled *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) SetEnabled(v bool) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection {
	s.Enabled = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) SetMessage(v string) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection {
	s.Message = &v
	return s
}

type UpdateProtectedBranchesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateProtectedBranchesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProtectedBranchesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProtectedBranchesResponse) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponse) SetHeaders(v map[string]*string) *UpdateProtectedBranchesResponse {
	s.Headers = v
	return s
}

func (s *UpdateProtectedBranchesResponse) SetStatusCode(v int32) *UpdateProtectedBranchesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProtectedBranchesResponse) SetBody(v *UpdateProtectedBranchesResponseBody) *UpdateProtectedBranchesResponse {
	s.Body = v
	return s
}

type UpdateRepositoryRequest struct {
	AccessToken                    *string                                                  `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	AdminSettingLanguage           *string                                                  `json:"adminSettingLanguage,omitempty" xml:"adminSettingLanguage,omitempty"`
	Avatar                         *string                                                  `json:"avatar,omitempty" xml:"avatar,omitempty"`
	BuildsEnabled                  *bool                                                    `json:"buildsEnabled,omitempty" xml:"buildsEnabled,omitempty"`
	CheckEmail                     *bool                                                    `json:"checkEmail,omitempty" xml:"checkEmail,omitempty"`
	DefaultBranch                  *string                                                  `json:"defaultBranch,omitempty" xml:"defaultBranch,omitempty"`
	Description                    *string                                                  `json:"description,omitempty" xml:"description,omitempty"`
	Id                             *int64                                                   `json:"id,omitempty" xml:"id,omitempty"`
	IssuesEnabled                  *bool                                                    `json:"issuesEnabled,omitempty" xml:"issuesEnabled,omitempty"`
	MergeRequestsEnabled           *bool                                                    `json:"mergeRequestsEnabled,omitempty" xml:"mergeRequestsEnabled,omitempty"`
	Name                           *string                                                  `json:"name,omitempty" xml:"name,omitempty"`
	OpenCloneDownloadControl       *bool                                                    `json:"openCloneDownloadControl,omitempty" xml:"openCloneDownloadControl,omitempty"`
	Path                           *string                                                  `json:"path,omitempty" xml:"path,omitempty"`
	ProjectCloneDownloadMethodList []*UpdateRepositoryRequestProjectCloneDownloadMethodList `json:"projectCloneDownloadMethodList,omitempty" xml:"projectCloneDownloadMethodList,omitempty" type:"Repeated"`
	ProjectCloneDownloadRoleList   []*UpdateRepositoryRequestProjectCloneDownloadRoleList   `json:"projectCloneDownloadRoleList,omitempty" xml:"projectCloneDownloadRoleList,omitempty" type:"Repeated"`
	SnippetsEnabled                *bool                                                    `json:"snippetsEnabled,omitempty" xml:"snippetsEnabled,omitempty"`
	VisibilityLevel                *int32                                                   `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	WikiEnabled                    *bool                                                    `json:"wikiEnabled,omitempty" xml:"wikiEnabled,omitempty"`
	OrganizationId                 *string                                                  `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s UpdateRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryRequest) SetAccessToken(v string) *UpdateRepositoryRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateRepositoryRequest) SetAdminSettingLanguage(v string) *UpdateRepositoryRequest {
	s.AdminSettingLanguage = &v
	return s
}

func (s *UpdateRepositoryRequest) SetAvatar(v string) *UpdateRepositoryRequest {
	s.Avatar = &v
	return s
}

func (s *UpdateRepositoryRequest) SetBuildsEnabled(v bool) *UpdateRepositoryRequest {
	s.BuildsEnabled = &v
	return s
}

func (s *UpdateRepositoryRequest) SetCheckEmail(v bool) *UpdateRepositoryRequest {
	s.CheckEmail = &v
	return s
}

func (s *UpdateRepositoryRequest) SetDefaultBranch(v string) *UpdateRepositoryRequest {
	s.DefaultBranch = &v
	return s
}

func (s *UpdateRepositoryRequest) SetDescription(v string) *UpdateRepositoryRequest {
	s.Description = &v
	return s
}

func (s *UpdateRepositoryRequest) SetId(v int64) *UpdateRepositoryRequest {
	s.Id = &v
	return s
}

func (s *UpdateRepositoryRequest) SetIssuesEnabled(v bool) *UpdateRepositoryRequest {
	s.IssuesEnabled = &v
	return s
}

func (s *UpdateRepositoryRequest) SetMergeRequestsEnabled(v bool) *UpdateRepositoryRequest {
	s.MergeRequestsEnabled = &v
	return s
}

func (s *UpdateRepositoryRequest) SetName(v string) *UpdateRepositoryRequest {
	s.Name = &v
	return s
}

func (s *UpdateRepositoryRequest) SetOpenCloneDownloadControl(v bool) *UpdateRepositoryRequest {
	s.OpenCloneDownloadControl = &v
	return s
}

func (s *UpdateRepositoryRequest) SetPath(v string) *UpdateRepositoryRequest {
	s.Path = &v
	return s
}

func (s *UpdateRepositoryRequest) SetProjectCloneDownloadMethodList(v []*UpdateRepositoryRequestProjectCloneDownloadMethodList) *UpdateRepositoryRequest {
	s.ProjectCloneDownloadMethodList = v
	return s
}

func (s *UpdateRepositoryRequest) SetProjectCloneDownloadRoleList(v []*UpdateRepositoryRequestProjectCloneDownloadRoleList) *UpdateRepositoryRequest {
	s.ProjectCloneDownloadRoleList = v
	return s
}

func (s *UpdateRepositoryRequest) SetSnippetsEnabled(v bool) *UpdateRepositoryRequest {
	s.SnippetsEnabled = &v
	return s
}

func (s *UpdateRepositoryRequest) SetVisibilityLevel(v int32) *UpdateRepositoryRequest {
	s.VisibilityLevel = &v
	return s
}

func (s *UpdateRepositoryRequest) SetWikiEnabled(v bool) *UpdateRepositoryRequest {
	s.WikiEnabled = &v
	return s
}

func (s *UpdateRepositoryRequest) SetOrganizationId(v string) *UpdateRepositoryRequest {
	s.OrganizationId = &v
	return s
}

type UpdateRepositoryRequestProjectCloneDownloadMethodList struct {
	Allowed        *bool   `json:"allowed,omitempty" xml:"allowed,omitempty"`
	PermissionCode *string `json:"permissionCode,omitempty" xml:"permissionCode,omitempty"`
}

func (s UpdateRepositoryRequestProjectCloneDownloadMethodList) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryRequestProjectCloneDownloadMethodList) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryRequestProjectCloneDownloadMethodList) SetAllowed(v bool) *UpdateRepositoryRequestProjectCloneDownloadMethodList {
	s.Allowed = &v
	return s
}

func (s *UpdateRepositoryRequestProjectCloneDownloadMethodList) SetPermissionCode(v string) *UpdateRepositoryRequestProjectCloneDownloadMethodList {
	s.PermissionCode = &v
	return s
}

type UpdateRepositoryRequestProjectCloneDownloadRoleList struct {
	Allowed  *bool  `json:"allowed,omitempty" xml:"allowed,omitempty"`
	RoleCode *int32 `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
}

func (s UpdateRepositoryRequestProjectCloneDownloadRoleList) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryRequestProjectCloneDownloadRoleList) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryRequestProjectCloneDownloadRoleList) SetAllowed(v bool) *UpdateRepositoryRequestProjectCloneDownloadRoleList {
	s.Allowed = &v
	return s
}

func (s *UpdateRepositoryRequestProjectCloneDownloadRoleList) SetRoleCode(v int32) *UpdateRepositoryRequestProjectCloneDownloadRoleList {
	s.RoleCode = &v
	return s
}

type UpdateRepositoryResponseBody struct {
	ErrorCode    *string                             `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                             `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       *UpdateRepositoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success      *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryResponseBody) SetErrorCode(v string) *UpdateRepositoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateRepositoryResponseBody) SetErrorMessage(v string) *UpdateRepositoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateRepositoryResponseBody) SetRequestId(v string) *UpdateRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRepositoryResponseBody) SetResult(v *UpdateRepositoryResponseBodyResult) *UpdateRepositoryResponseBody {
	s.Result = v
	return s
}

func (s *UpdateRepositoryResponseBody) SetSuccess(v bool) *UpdateRepositoryResponseBody {
	s.Success = &v
	return s
}

type UpdateRepositoryResponseBodyResult struct {
	Archived             *bool                                        `json:"archived,omitempty" xml:"archived,omitempty"`
	AvatarUrl            *string                                      `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	BuildsEnabled        *bool                                        `json:"buildsEnabled,omitempty" xml:"buildsEnabled,omitempty"`
	CreatedAt            *string                                      `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatorId            *int64                                       `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	DefaultBranch        *string                                      `json:"defaultBranch,omitempty" xml:"defaultBranch,omitempty"`
	Description          *string                                      `json:"description,omitempty" xml:"description,omitempty"`
	HttpUrlToRepo        *string                                      `json:"httpUrlToRepo,omitempty" xml:"httpUrlToRepo,omitempty"`
	Id                   *int64                                       `json:"id,omitempty" xml:"id,omitempty"`
	IssuesEnabled        *bool                                        `json:"issuesEnabled,omitempty" xml:"issuesEnabled,omitempty"`
	LastActivityAt       *string                                      `json:"lastActivityAt,omitempty" xml:"lastActivityAt,omitempty"`
	MergeRequestsEnabled *bool                                        `json:"mergeRequestsEnabled,omitempty" xml:"mergeRequestsEnabled,omitempty"`
	Name                 *string                                      `json:"name,omitempty" xml:"name,omitempty"`
	NameWithNamespace    *string                                      `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	Namespace            *UpdateRepositoryResponseBodyResultNamespace `json:"namespace,omitempty" xml:"namespace,omitempty" type:"Struct"`
	Path                 *string                                      `json:"path,omitempty" xml:"path,omitempty"`
	PathWithNamespace    *string                                      `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	SnippetsEnabled      *bool                                        `json:"snippetsEnabled,omitempty" xml:"snippetsEnabled,omitempty"`
	SshUrlToRepo         *string                                      `json:"sshUrlToRepo,omitempty" xml:"sshUrlToRepo,omitempty"`
	VisibilityLevel      *int32                                       `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	WebUrl               *string                                      `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
	WikiEnabled          *bool                                        `json:"wikiEnabled,omitempty" xml:"wikiEnabled,omitempty"`
}

func (s UpdateRepositoryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryResponseBodyResult) SetArchived(v bool) *UpdateRepositoryResponseBodyResult {
	s.Archived = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetAvatarUrl(v string) *UpdateRepositoryResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetBuildsEnabled(v bool) *UpdateRepositoryResponseBodyResult {
	s.BuildsEnabled = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetCreatedAt(v string) *UpdateRepositoryResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetCreatorId(v int64) *UpdateRepositoryResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetDefaultBranch(v string) *UpdateRepositoryResponseBodyResult {
	s.DefaultBranch = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetDescription(v string) *UpdateRepositoryResponseBodyResult {
	s.Description = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetHttpUrlToRepo(v string) *UpdateRepositoryResponseBodyResult {
	s.HttpUrlToRepo = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetId(v int64) *UpdateRepositoryResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetIssuesEnabled(v bool) *UpdateRepositoryResponseBodyResult {
	s.IssuesEnabled = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetLastActivityAt(v string) *UpdateRepositoryResponseBodyResult {
	s.LastActivityAt = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetMergeRequestsEnabled(v bool) *UpdateRepositoryResponseBodyResult {
	s.MergeRequestsEnabled = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetName(v string) *UpdateRepositoryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetNameWithNamespace(v string) *UpdateRepositoryResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetNamespace(v *UpdateRepositoryResponseBodyResultNamespace) *UpdateRepositoryResponseBodyResult {
	s.Namespace = v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetPath(v string) *UpdateRepositoryResponseBodyResult {
	s.Path = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetPathWithNamespace(v string) *UpdateRepositoryResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetSnippetsEnabled(v bool) *UpdateRepositoryResponseBodyResult {
	s.SnippetsEnabled = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetSshUrlToRepo(v string) *UpdateRepositoryResponseBodyResult {
	s.SshUrlToRepo = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetVisibilityLevel(v int32) *UpdateRepositoryResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetWebUrl(v string) *UpdateRepositoryResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetWikiEnabled(v bool) *UpdateRepositoryResponseBodyResult {
	s.WikiEnabled = &v
	return s
}

type UpdateRepositoryResponseBodyResultNamespace struct {
	Avatar          *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	CreatedAt       *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	Id              *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	OwnerId         *int64  `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	Path            *string `json:"path,omitempty" xml:"path,omitempty"`
	UpdatedAt       *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	VisibilityLevel *int32  `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
}

func (s UpdateRepositoryResponseBodyResultNamespace) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryResponseBodyResultNamespace) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetAvatar(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.Avatar = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetCreatedAt(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.CreatedAt = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetDescription(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.Description = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetId(v int64) *UpdateRepositoryResponseBodyResultNamespace {
	s.Id = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetName(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.Name = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetOwnerId(v int64) *UpdateRepositoryResponseBodyResultNamespace {
	s.OwnerId = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetPath(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.Path = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetUpdatedAt(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetVisibilityLevel(v int32) *UpdateRepositoryResponseBodyResultNamespace {
	s.VisibilityLevel = &v
	return s
}

type UpdateRepositoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryResponse) SetHeaders(v map[string]*string) *UpdateRepositoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateRepositoryResponse) SetStatusCode(v int32) *UpdateRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRepositoryResponse) SetBody(v *UpdateRepositoryResponseBody) *UpdateRepositoryResponse {
	s.Body = v
	return s
}

type UpdateRepositoryMemberRequest struct {
	AccessToken    *string                                      `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	AccessLevel    *int32                                       `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	ExpireAt       *string                                      `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
	MemberType     *string                                      `json:"memberType,omitempty" xml:"memberType,omitempty"`
	RelatedId      *string                                      `json:"relatedId,omitempty" xml:"relatedId,omitempty"`
	RelatedInfos   []*UpdateRepositoryMemberRequestRelatedInfos `json:"relatedInfos,omitempty" xml:"relatedInfos,omitempty" type:"Repeated"`
	OrganizationId *string                                      `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s UpdateRepositoryMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryMemberRequest) SetAccessToken(v string) *UpdateRepositoryMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateRepositoryMemberRequest) SetAccessLevel(v int32) *UpdateRepositoryMemberRequest {
	s.AccessLevel = &v
	return s
}

func (s *UpdateRepositoryMemberRequest) SetExpireAt(v string) *UpdateRepositoryMemberRequest {
	s.ExpireAt = &v
	return s
}

func (s *UpdateRepositoryMemberRequest) SetMemberType(v string) *UpdateRepositoryMemberRequest {
	s.MemberType = &v
	return s
}

func (s *UpdateRepositoryMemberRequest) SetRelatedId(v string) *UpdateRepositoryMemberRequest {
	s.RelatedId = &v
	return s
}

func (s *UpdateRepositoryMemberRequest) SetRelatedInfos(v []*UpdateRepositoryMemberRequestRelatedInfos) *UpdateRepositoryMemberRequest {
	s.RelatedInfos = v
	return s
}

func (s *UpdateRepositoryMemberRequest) SetOrganizationId(v string) *UpdateRepositoryMemberRequest {
	s.OrganizationId = &v
	return s
}

type UpdateRepositoryMemberRequestRelatedInfos struct {
	RelatedId  *string `json:"relatedId,omitempty" xml:"relatedId,omitempty"`
	SourceId   *int64  `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s UpdateRepositoryMemberRequestRelatedInfos) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryMemberRequestRelatedInfos) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryMemberRequestRelatedInfos) SetRelatedId(v string) *UpdateRepositoryMemberRequestRelatedInfos {
	s.RelatedId = &v
	return s
}

func (s *UpdateRepositoryMemberRequestRelatedInfos) SetSourceId(v int64) *UpdateRepositoryMemberRequestRelatedInfos {
	s.SourceId = &v
	return s
}

func (s *UpdateRepositoryMemberRequestRelatedInfos) SetSourceType(v string) *UpdateRepositoryMemberRequestRelatedInfos {
	s.SourceType = &v
	return s
}

type UpdateRepositoryMemberResponseBody struct {
	ErrorCode    *string                                   `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                   `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       *UpdateRepositoryMemberResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success      *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateRepositoryMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryMemberResponseBody) SetErrorCode(v string) *UpdateRepositoryMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBody) SetErrorMessage(v string) *UpdateRepositoryMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBody) SetRequestId(v string) *UpdateRepositoryMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBody) SetResult(v *UpdateRepositoryMemberResponseBodyResult) *UpdateRepositoryMemberResponseBody {
	s.Result = v
	return s
}

func (s *UpdateRepositoryMemberResponseBody) SetSuccess(v bool) *UpdateRepositoryMemberResponseBody {
	s.Success = &v
	return s
}

type UpdateRepositoryMemberResponseBodyResult struct {
	AccessLevel *int32  `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	AvatarUrl   *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Email       *string `json:"email,omitempty" xml:"email,omitempty"`
	ExpireAt    *string `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
	ExternUid   *string `json:"externUid,omitempty" xml:"externUid,omitempty"`
	Id          *int64  `json:"id,omitempty" xml:"id,omitempty"`
	MemberName  *string `json:"memberName,omitempty" xml:"memberName,omitempty"`
	MemberType  *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceId    *int64  `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	SourceType  *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	State       *string `json:"state,omitempty" xml:"state,omitempty"`
	TbUserId    *string `json:"tbUserId,omitempty" xml:"tbUserId,omitempty"`
	Username    *string `json:"username,omitempty" xml:"username,omitempty"`
	WebUrl      *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s UpdateRepositoryMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetAccessLevel(v int32) *UpdateRepositoryMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetAvatarUrl(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetEmail(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetExpireAt(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.ExpireAt = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetExternUid(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.ExternUid = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetId(v int64) *UpdateRepositoryMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetMemberName(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.MemberName = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetMemberType(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.MemberType = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetName(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetSourceId(v int64) *UpdateRepositoryMemberResponseBodyResult {
	s.SourceId = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetSourceType(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetState(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.State = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetTbUserId(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.TbUserId = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetUsername(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.Username = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetWebUrl(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.WebUrl = &v
	return s
}

type UpdateRepositoryMemberResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRepositoryMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRepositoryMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryMemberResponse) SetHeaders(v map[string]*string) *UpdateRepositoryMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateRepositoryMemberResponse) SetStatusCode(v int32) *UpdateRepositoryMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRepositoryMemberResponse) SetBody(v *UpdateRepositoryMemberResponseBody) *UpdateRepositoryMemberResponse {
	s.Body = v
	return s
}

type UpdateResourceMemberRequest struct {
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s UpdateResourceMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceMemberRequest) SetRoleName(v string) *UpdateResourceMemberRequest {
	s.RoleName = &v
	return s
}

type UpdateResourceMemberResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateResourceMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceMemberResponseBody) SetErrorCode(v string) *UpdateResourceMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateResourceMemberResponseBody) SetErrorMessage(v string) *UpdateResourceMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateResourceMemberResponseBody) SetRequestId(v string) *UpdateResourceMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceMemberResponseBody) SetSuccess(v bool) *UpdateResourceMemberResponseBody {
	s.Success = &v
	return s
}

type UpdateResourceMemberResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateResourceMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateResourceMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceMemberResponse) SetHeaders(v map[string]*string) *UpdateResourceMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceMemberResponse) SetStatusCode(v int32) *UpdateResourceMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceMemberResponse) SetBody(v *UpdateResourceMemberResponseBody) *UpdateResourceMemberResponse {
	s.Body = v
	return s
}

type UpdateVariableGroupRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Variables   *string `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s UpdateVariableGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVariableGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateVariableGroupRequest) SetDescription(v string) *UpdateVariableGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateVariableGroupRequest) SetName(v string) *UpdateVariableGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateVariableGroupRequest) SetVariables(v string) *UpdateVariableGroupRequest {
	s.Variables = &v
	return s
}

type UpdateVariableGroupResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateVariableGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVariableGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVariableGroupResponseBody) SetErrorCode(v string) *UpdateVariableGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateVariableGroupResponseBody) SetErrorMessage(v string) *UpdateVariableGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateVariableGroupResponseBody) SetRequestId(v string) *UpdateVariableGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVariableGroupResponseBody) SetSuccess(v bool) *UpdateVariableGroupResponseBody {
	s.Success = &v
	return s
}

type UpdateVariableGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateVariableGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVariableGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVariableGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateVariableGroupResponse) SetHeaders(v map[string]*string) *UpdateVariableGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateVariableGroupResponse) SetStatusCode(v int32) *UpdateVariableGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVariableGroupResponse) SetBody(v *UpdateVariableGroupResponseBody) *UpdateVariableGroupResponse {
	s.Body = v
	return s
}

type UpdateWorkItemRequest struct {
	FieldType     *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	Identifier    *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	PropertyKey   *string `json:"propertyKey,omitempty" xml:"propertyKey,omitempty"`
	PropertyValue *string `json:"propertyValue,omitempty" xml:"propertyValue,omitempty"`
}

func (s UpdateWorkItemRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkItemRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkItemRequest) SetFieldType(v string) *UpdateWorkItemRequest {
	s.FieldType = &v
	return s
}

func (s *UpdateWorkItemRequest) SetIdentifier(v string) *UpdateWorkItemRequest {
	s.Identifier = &v
	return s
}

func (s *UpdateWorkItemRequest) SetPropertyKey(v string) *UpdateWorkItemRequest {
	s.PropertyKey = &v
	return s
}

func (s *UpdateWorkItemRequest) SetPropertyValue(v string) *UpdateWorkItemRequest {
	s.PropertyValue = &v
	return s
}

type UpdateWorkItemResponseBody struct {
	ErrorCode    *string                             `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                             `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool                               `json:"success,omitempty" xml:"success,omitempty"`
	Workitem     *UpdateWorkItemResponseBodyWorkitem `json:"workitem,omitempty" xml:"workitem,omitempty" type:"Struct"`
}

func (s UpdateWorkItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkItemResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkItemResponseBody) SetErrorCode(v string) *UpdateWorkItemResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetErrorMessage(v string) *UpdateWorkItemResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetRequestId(v string) *UpdateWorkItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetSuccess(v bool) *UpdateWorkItemResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetWorkitem(v *UpdateWorkItemResponseBodyWorkitem) *UpdateWorkItemResponseBody {
	s.Workitem = v
	return s
}

type UpdateWorkItemResponseBodyWorkitem struct {
	AssignedTo             *string `json:"assignedTo,omitempty" xml:"assignedTo,omitempty"`
	CategoryIdentifier     *string `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	Creator                *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Document               *string `json:"document,omitempty" xml:"document,omitempty"`
	GmtCreate              *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified            *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Identifier             *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	LogicalStatus          *string `json:"logicalStatus,omitempty" xml:"logicalStatus,omitempty"`
	Modifier               *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	ParentIdentifier       *string `json:"parentIdentifier,omitempty" xml:"parentIdentifier,omitempty"`
	SerialNumber           *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	SpaceIdentifier        *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	SpaceName              *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	SpaceType              *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	SprintIdentifier       *string `json:"sprintIdentifier,omitempty" xml:"sprintIdentifier,omitempty"`
	Status                 *string `json:"status,omitempty" xml:"status,omitempty"`
	StatusIdentifier       *string `json:"statusIdentifier,omitempty" xml:"statusIdentifier,omitempty"`
	StatusStageIdentifier  *string `json:"statusStageIdentifier,omitempty" xml:"statusStageIdentifier,omitempty"`
	Subject                *string `json:"subject,omitempty" xml:"subject,omitempty"`
	UpdateStatusAt         *int64  `json:"updateStatusAt,omitempty" xml:"updateStatusAt,omitempty"`
	WorkitemTypeIdentifier *string `json:"workitemTypeIdentifier,omitempty" xml:"workitemTypeIdentifier,omitempty"`
}

func (s UpdateWorkItemResponseBodyWorkitem) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkItemResponseBodyWorkitem) GoString() string {
	return s.String()
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetAssignedTo(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.AssignedTo = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetCategoryIdentifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.CategoryIdentifier = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetCreator(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.Creator = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetDocument(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.Document = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetGmtCreate(v int64) *UpdateWorkItemResponseBodyWorkitem {
	s.GmtCreate = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetGmtModified(v int64) *UpdateWorkItemResponseBodyWorkitem {
	s.GmtModified = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetIdentifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.Identifier = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetLogicalStatus(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.LogicalStatus = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetModifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.Modifier = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetParentIdentifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.ParentIdentifier = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetSerialNumber(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.SerialNumber = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetSpaceIdentifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.SpaceIdentifier = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetSpaceName(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.SpaceName = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetSpaceType(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.SpaceType = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetSprintIdentifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.SprintIdentifier = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetStatus(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.Status = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetStatusIdentifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.StatusIdentifier = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetStatusStageIdentifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.StatusStageIdentifier = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetSubject(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.Subject = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetUpdateStatusAt(v int64) *UpdateWorkItemResponseBodyWorkitem {
	s.UpdateStatusAt = &v
	return s
}

func (s *UpdateWorkItemResponseBodyWorkitem) SetWorkitemTypeIdentifier(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.WorkitemTypeIdentifier = &v
	return s
}

type UpdateWorkItemResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateWorkItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWorkItemResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkItemResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkItemResponse) SetHeaders(v map[string]*string) *UpdateWorkItemResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkItemResponse) SetStatusCode(v int32) *UpdateWorkItemResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkItemResponse) SetBody(v *UpdateWorkItemResponseBody) *UpdateWorkItemResponse {
	s.Body = v
	return s
}

type UpdateWorkitemCommentRequest struct {
	CommentId          *int64  `json:"commentId,omitempty" xml:"commentId,omitempty"`
	Content            *string `json:"content,omitempty" xml:"content,omitempty"`
	FormatType         *string `json:"formatType,omitempty" xml:"formatType,omitempty"`
	WorkitemIdentifier *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s UpdateWorkitemCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkitemCommentRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemCommentRequest) SetCommentId(v int64) *UpdateWorkitemCommentRequest {
	s.CommentId = &v
	return s
}

func (s *UpdateWorkitemCommentRequest) SetContent(v string) *UpdateWorkitemCommentRequest {
	s.Content = &v
	return s
}

func (s *UpdateWorkitemCommentRequest) SetFormatType(v string) *UpdateWorkitemCommentRequest {
	s.FormatType = &v
	return s
}

func (s *UpdateWorkitemCommentRequest) SetWorkitemIdentifier(v string) *UpdateWorkitemCommentRequest {
	s.WorkitemIdentifier = &v
	return s
}

type UpdateWorkitemCommentResponseBody struct {
	Comment   *UpdateWorkitemCommentResponseBodyComment `json:"comment,omitempty" xml:"comment,omitempty" type:"Struct"`
	ErrorCode *string                                   `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                   `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *string                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateWorkitemCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkitemCommentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemCommentResponseBody) SetComment(v *UpdateWorkitemCommentResponseBodyComment) *UpdateWorkitemCommentResponseBody {
	s.Comment = v
	return s
}

func (s *UpdateWorkitemCommentResponseBody) SetErrorCode(v string) *UpdateWorkitemCommentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBody) SetErrorMsg(v string) *UpdateWorkitemCommentResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBody) SetRequestId(v string) *UpdateWorkitemCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBody) SetSuccess(v string) *UpdateWorkitemCommentResponseBody {
	s.Success = &v
	return s
}

type UpdateWorkitemCommentResponseBodyComment struct {
	Content          *string                                       `json:"content,omitempty" xml:"content,omitempty"`
	CreateTime       *int64                                        `json:"createTime,omitempty" xml:"createTime,omitempty"`
	FormatType       *string                                       `json:"formatType,omitempty" xml:"formatType,omitempty"`
	Id               *string                                       `json:"id,omitempty" xml:"id,omitempty"`
	IsTop            *bool                                         `json:"isTop,omitempty" xml:"isTop,omitempty"`
	ModifiedTime     *int64                                        `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ParentId         *int64                                        `json:"parentId,omitempty" xml:"parentId,omitempty"`
	TargetIdentifier *string                                       `json:"targetIdentifier,omitempty" xml:"targetIdentifier,omitempty"`
	TargetType       *string                                       `json:"targetType,omitempty" xml:"targetType,omitempty"`
	TopTime          *int64                                        `json:"topTime,omitempty" xml:"topTime,omitempty"`
	User             *UpdateWorkitemCommentResponseBodyCommentUser `json:"user,omitempty" xml:"user,omitempty" type:"Struct"`
}

func (s UpdateWorkitemCommentResponseBodyComment) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkitemCommentResponseBodyComment) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetContent(v string) *UpdateWorkitemCommentResponseBodyComment {
	s.Content = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetCreateTime(v int64) *UpdateWorkitemCommentResponseBodyComment {
	s.CreateTime = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetFormatType(v string) *UpdateWorkitemCommentResponseBodyComment {
	s.FormatType = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetId(v string) *UpdateWorkitemCommentResponseBodyComment {
	s.Id = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetIsTop(v bool) *UpdateWorkitemCommentResponseBodyComment {
	s.IsTop = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetModifiedTime(v int64) *UpdateWorkitemCommentResponseBodyComment {
	s.ModifiedTime = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetParentId(v int64) *UpdateWorkitemCommentResponseBodyComment {
	s.ParentId = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetTargetIdentifier(v string) *UpdateWorkitemCommentResponseBodyComment {
	s.TargetIdentifier = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetTargetType(v string) *UpdateWorkitemCommentResponseBodyComment {
	s.TargetType = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetTopTime(v int64) *UpdateWorkitemCommentResponseBodyComment {
	s.TopTime = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetUser(v *UpdateWorkitemCommentResponseBodyCommentUser) *UpdateWorkitemCommentResponseBodyComment {
	s.User = v
	return s
}

type UpdateWorkitemCommentResponseBodyCommentUser struct {
	Account     *string `json:"account,omitempty" xml:"account,omitempty"`
	Avatar      *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Identifier  *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	NickName    *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	RealName    *string `json:"realName,omitempty" xml:"realName,omitempty"`
}

func (s UpdateWorkitemCommentResponseBodyCommentUser) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkitemCommentResponseBodyCommentUser) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) SetAccount(v string) *UpdateWorkitemCommentResponseBodyCommentUser {
	s.Account = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) SetAvatar(v string) *UpdateWorkitemCommentResponseBodyCommentUser {
	s.Avatar = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) SetDisplayName(v string) *UpdateWorkitemCommentResponseBodyCommentUser {
	s.DisplayName = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) SetIdentifier(v string) *UpdateWorkitemCommentResponseBodyCommentUser {
	s.Identifier = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) SetNickName(v string) *UpdateWorkitemCommentResponseBodyCommentUser {
	s.NickName = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) SetRealName(v string) *UpdateWorkitemCommentResponseBodyCommentUser {
	s.RealName = &v
	return s
}

type UpdateWorkitemCommentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateWorkitemCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWorkitemCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkitemCommentResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemCommentResponse) SetHeaders(v map[string]*string) *UpdateWorkitemCommentResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkitemCommentResponse) SetStatusCode(v int32) *UpdateWorkitemCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkitemCommentResponse) SetBody(v *UpdateWorkitemCommentResponseBody) *UpdateWorkitemCommentResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("devops"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddRepositoryMember(repositoryId *string, request *AddRepositoryMemberRequest) (_result *AddRepositoryMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddRepositoryMemberResponse{}
	_body, _err := client.AddRepositoryMemberWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddRepositoryMemberWithOptions(repositoryId *string, request *AddRepositoryMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddRepositoryMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessLevel)) {
		body["accessLevel"] = request.AccessLevel
	}

	if !tea.BoolValue(util.IsUnset(request.AliyunPks)) {
		body["aliyunPks"] = request.AliyunPks
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddRepositoryMember"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/members"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRepositoryMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddWebhook(repositoryId *string, request *AddWebhookRequest) (_result *AddWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddWebhookResponse{}
	_body, _err := client.AddWebhookWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddWebhookWithOptions(repositoryId *string, request *AddWebhookRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSslVerification)) {
		body["enableSslVerification"] = request.EnableSslVerification
	}

	if !tea.BoolValue(util.IsUnset(request.MergeRequestsEvents)) {
		body["mergeRequestsEvents"] = request.MergeRequestsEvents
	}

	if !tea.BoolValue(util.IsUnset(request.NoteEvents)) {
		body["noteEvents"] = request.NoteEvents
	}

	if !tea.BoolValue(util.IsUnset(request.PushEvents)) {
		body["pushEvents"] = request.PushEvents
	}

	if !tea.BoolValue(util.IsUnset(request.SecretToken)) {
		body["secretToken"] = request.SecretToken
	}

	if !tea.BoolValue(util.IsUnset(request.TagPushEvents)) {
		body["tagPushEvents"] = request.TagPushEvents
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddWebhook"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/webhooks/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBranch(repositoryId *string, request *CreateBranchRequest) (_result *CreateBranchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateBranchResponse{}
	_body, _err := client.CreateBranchWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBranchWithOptions(repositoryId *string, request *CreateBranchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateBranchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BranchName)) {
		body["branchName"] = request.BranchName
	}

	if !tea.BoolValue(util.IsUnset(request.Ref)) {
		body["ref"] = request.Ref
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBranch"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/branches"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBranchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFile(repositoryId *string, request *CreateFileRequest) (_result *CreateFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFileResponse{}
	_body, _err := client.CreateFileWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFileWithOptions(repositoryId *string, request *CreateFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BranchName)) {
		body["branchName"] = request.BranchName
	}

	if !tea.BoolValue(util.IsUnset(request.CommitMessage)) {
		body["commitMessage"] = request.CommitMessage
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Encoding)) {
		body["encoding"] = request.Encoding
	}

	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		body["filePath"] = request.FilePath
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFile"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/files"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlowTag(organizationId *string, request *CreateFlowTagRequest) (_result *CreateFlowTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFlowTagResponse{}
	_body, _err := client.CreateFlowTagWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowTagWithOptions(organizationId *string, request *CreateFlowTagRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFlowTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Color)) {
		query["color"] = request.Color
	}

	if !tea.BoolValue(util.IsUnset(request.FlowTagGroupId)) {
		query["flowTagGroupId"] = request.FlowTagGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFlowTag"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/flow/tags"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFlowTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlowTagGroup(organizationId *string, request *CreateFlowTagGroupRequest) (_result *CreateFlowTagGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFlowTagGroupResponse{}
	_body, _err := client.CreateFlowTagGroupWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowTagGroupWithOptions(organizationId *string, request *CreateFlowTagGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFlowTagGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFlowTagGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/flow/tagGroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFlowTagGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHostGroup(organizationId *string, request *CreateHostGroupRequest) (_result *CreateHostGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateHostGroupResponse{}
	_body, _err := client.CreateHostGroupWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHostGroupWithOptions(organizationId *string, request *CreateHostGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateHostGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunRegion)) {
		body["aliyunRegion"] = request.AliyunRegion
	}

	if !tea.BoolValue(util.IsUnset(request.EcsLabelKey)) {
		body["ecsLabelKey"] = request.EcsLabelKey
	}

	if !tea.BoolValue(util.IsUnset(request.EcsLabelValue)) {
		body["ecsLabelValue"] = request.EcsLabelValue
	}

	if !tea.BoolValue(util.IsUnset(request.EcsType)) {
		body["ecsType"] = request.EcsType
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		body["envId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.MachineInfos)) {
		body["machineInfos"] = request.MachineInfos
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceConnectionId)) {
		body["serviceConnectionId"] = request.ServiceConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.TagIds)) {
		body["tagIds"] = request.TagIds
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHostGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/hostGroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOAuthToken(request *CreateOAuthTokenRequest) (_result *CreateOAuthTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateOAuthTokenResponse{}
	_body, _err := client.CreateOAuthTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOAuthTokenWithOptions(request *CreateOAuthTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateOAuthTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["clientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSecret)) {
		body["clientSecret"] = request.ClientSecret
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.GrantType)) {
		body["grantType"] = request.GrantType
	}

	if !tea.BoolValue(util.IsUnset(request.Login)) {
		body["login"] = request.Login
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOAuthToken"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/login/oauth/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOAuthTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePipelineGroup(organizationId *string, request *CreatePipelineGroupRequest) (_result *CreatePipelineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePipelineGroupResponse{}
	_body, _err := client.CreatePipelineGroupWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePipelineGroupWithOptions(organizationId *string, request *CreatePipelineGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePipelineGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePipelineGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelineGroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePipelineGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProject(organizationId *string, request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProjectWithOptions(organizationId *string, request *CreateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomCode)) {
		body["customCode"] = request.CustomCode
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateIdentifier)) {
		body["templateIdentifier"] = request.TemplateIdentifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProject"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/projects/createProject"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRepository(request *CreateRepositoryRequest) (_result *CreateRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRepositoryResponse{}
	_body, _err := client.CreateRepositoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRepositoryWithOptions(request *CreateRepositoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.CreateParentPath)) {
		query["createParentPath"] = request.CreateParentPath
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Sync)) {
		query["sync"] = request.Sync
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AvatarUrl)) {
		body["avatarUrl"] = request.AvatarUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GitignoreType)) {
		body["gitignoreType"] = request.GitignoreType
	}

	if !tea.BoolValue(util.IsUnset(request.ImportAccount)) {
		body["importAccount"] = request.ImportAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ImportDemoProject)) {
		body["importDemoProject"] = request.ImportDemoProject
	}

	if !tea.BoolValue(util.IsUnset(request.ImportRepoType)) {
		body["importRepoType"] = request.ImportRepoType
	}

	if !tea.BoolValue(util.IsUnset(request.ImportToken)) {
		body["importToken"] = request.ImportToken
	}

	if !tea.BoolValue(util.IsUnset(request.ImportTokenEncrypted)) {
		body["importTokenEncrypted"] = request.ImportTokenEncrypted
	}

	if !tea.BoolValue(util.IsUnset(request.ImportUrl)) {
		body["importUrl"] = request.ImportUrl
	}

	if !tea.BoolValue(util.IsUnset(request.InitStandardService)) {
		body["initStandardService"] = request.InitStandardService
	}

	if !tea.BoolValue(util.IsUnset(request.IsCryptoEnabled)) {
		body["isCryptoEnabled"] = request.IsCryptoEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.LocalImportUrl)) {
		body["localImportUrl"] = request.LocalImportUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		body["namespaceId"] = request.NamespaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ReadmeType)) {
		body["readmeType"] = request.ReadmeType
	}

	if !tea.BoolValue(util.IsUnset(request.VisibilityLevel)) {
		body["visibilityLevel"] = request.VisibilityLevel
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRepository"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateResourceMember(organizationId *string, resourceType *string, resourceId *string, request *CreateResourceMemberRequest) (_result *CreateResourceMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceMemberResponse{}
	_body, _err := client.CreateResourceMemberWithOptions(organizationId, resourceType, resourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateResourceMemberWithOptions(organizationId *string, resourceType *string, resourceId *string, request *CreateResourceMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateResourceMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		body["roleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResourceMember"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(resourceType)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(resourceId)) + "/members"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSprint(organizationId *string, request *CreateSprintRequest) (_result *CreateSprintResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSprintResponse{}
	_body, _err := client.CreateSprintWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSprintWithOptions(organizationId *string, request *CreateSprintRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSprintResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceIdentifier)) {
		body["spaceIdentifier"] = request.SpaceIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.StaffIds)) {
		body["staffIds"] = request.StaffIds
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["startDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSprint"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/sprints/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSprintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSshKey(organizationId *string) (_result *CreateSshKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSshKeyResponse{}
	_body, _err := client.CreateSshKeyWithOptions(organizationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSshKeyWithOptions(organizationId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSshKeyResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSshKey"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/sshKey"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSshKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVariableGroup(organizationId *string, request *CreateVariableGroupRequest) (_result *CreateVariableGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateVariableGroupResponse{}
	_body, _err := client.CreateVariableGroupWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVariableGroupWithOptions(organizationId *string, request *CreateVariableGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateVariableGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Variables)) {
		body["variables"] = request.Variables
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVariableGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/variableGroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVariableGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWorkitem(organizationId *string, request *CreateWorkitemRequest) (_result *CreateWorkitemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWorkitemResponse{}
	_body, _err := client.CreateWorkitemWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWorkitemWithOptions(organizationId *string, request *CreateWorkitemRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateWorkitemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssignedTo)) {
		body["assignedTo"] = request.AssignedTo
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DescriptionFormat)) {
		body["descriptionFormat"] = request.DescriptionFormat
	}

	if !tea.BoolValue(util.IsUnset(request.FieldValueList)) {
		body["fieldValueList"] = request.FieldValueList
	}

	if !tea.BoolValue(util.IsUnset(request.Parent)) {
		body["parent"] = request.Parent
	}

	if !tea.BoolValue(util.IsUnset(request.Participant)) {
		body["participant"] = request.Participant
	}

	if !tea.BoolValue(util.IsUnset(request.Space)) {
		body["space"] = request.Space
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceIdentifier)) {
		body["spaceIdentifier"] = request.SpaceIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceType)) {
		body["spaceType"] = request.SpaceType
	}

	if !tea.BoolValue(util.IsUnset(request.Sprint)) {
		body["sprint"] = request.Sprint
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		body["subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.Tracker)) {
		body["tracker"] = request.Tracker
	}

	if !tea.BoolValue(util.IsUnset(request.Verifier)) {
		body["verifier"] = request.Verifier
	}

	if !tea.BoolValue(util.IsUnset(request.WorkitemType)) {
		body["workitemType"] = request.WorkitemType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWorkitem"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/workitems/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkitemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWorkitemComment(organizationId *string, request *CreateWorkitemCommentRequest) (_result *CreateWorkitemCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWorkitemCommentResponse{}
	_body, _err := client.CreateWorkitemCommentWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWorkitemCommentWithOptions(organizationId *string, request *CreateWorkitemCommentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateWorkitemCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.FormatType)) {
		body["formatType"] = request.FormatType
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		body["parentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkitemIdentifier)) {
		body["workitemIdentifier"] = request.WorkitemIdentifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWorkitemComment"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/workitems/comment"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkitemCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWorkitemEstimate(organizationId *string, request *CreateWorkitemEstimateRequest) (_result *CreateWorkitemEstimateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWorkitemEstimateResponse{}
	_body, _err := client.CreateWorkitemEstimateWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWorkitemEstimateWithOptions(organizationId *string, request *CreateWorkitemEstimateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateWorkitemEstimateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.RecordUserIdentifier)) {
		body["recordUserIdentifier"] = request.RecordUserIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.SpentTime)) {
		body["spentTime"] = request.SpentTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.WorkitemIdentifier)) {
		body["workitemIdentifier"] = request.WorkitemIdentifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWorkitemEstimate"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/workitems/estimate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkitemEstimateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWorkitemRecord(organizationId *string, request *CreateWorkitemRecordRequest) (_result *CreateWorkitemRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWorkitemRecordResponse{}
	_body, _err := client.CreateWorkitemRecordWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWorkitemRecordWithOptions(organizationId *string, request *CreateWorkitemRecordRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateWorkitemRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActualTime)) {
		body["actualTime"] = request.ActualTime
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GmtEnd)) {
		body["gmtEnd"] = request.GmtEnd
	}

	if !tea.BoolValue(util.IsUnset(request.GmtStart)) {
		body["gmtStart"] = request.GmtStart
	}

	if !tea.BoolValue(util.IsUnset(request.RecordUserIdentifier)) {
		body["recordUserIdentifier"] = request.RecordUserIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.WorkitemIdentifier)) {
		body["workitemIdentifier"] = request.WorkitemIdentifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWorkitemRecord"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/workitems/record"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkitemRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWorkspace(request *CreateWorkspaceRequest) (_result *CreateWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CreateWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWorkspaceWithOptions(request *CreateWorkspaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CodeUrl)) {
		body["codeUrl"] = request.CodeUrl
	}

	if !tea.BoolValue(util.IsUnset(request.CodeVersion)) {
		body["codeVersion"] = request.CodeVersion
	}

	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		body["filePath"] = request.FilePath
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RequestFrom)) {
		body["requestFrom"] = request.RequestFrom
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIdentifier)) {
		body["resourceIdentifier"] = request.ResourceIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.Reuse)) {
		body["reuse"] = request.Reuse
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceTemplate)) {
		body["workspaceTemplate"] = request.WorkspaceTemplate
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWorkspace"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/workspaces"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBranch(repositoryId *string, request *DeleteBranchRequest) (_result *DeleteBranchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBranchResponse{}
	_body, _err := client.DeleteBranchWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBranchWithOptions(repositoryId *string, request *DeleteBranchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteBranchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.BranchName)) {
		query["branchName"] = request.BranchName
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBranch"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/branches/delete"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBranchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFile(repositoryId *string, request *DeleteFileRequest) (_result *DeleteFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFileResponse{}
	_body, _err := client.DeleteFileWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFileWithOptions(repositoryId *string, request *DeleteFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.BranchName)) {
		query["branchName"] = request.BranchName
	}

	if !tea.BoolValue(util.IsUnset(request.CommitMessage)) {
		query["commitMessage"] = request.CommitMessage
	}

	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		query["filePath"] = request.FilePath
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFile"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/files/delete"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlowTag(organizationId *string, id *string) (_result *DeleteFlowTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFlowTagResponse{}
	_body, _err := client.DeleteFlowTagWithOptions(organizationId, id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowTagWithOptions(organizationId *string, id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFlowTagResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFlowTag"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/flow/tags/" + tea.StringValue(openapiutil.GetEncodeParam(id))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFlowTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlowTagGroup(organizationId *string, id *string) (_result *DeleteFlowTagGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFlowTagGroupResponse{}
	_body, _err := client.DeleteFlowTagGroupWithOptions(organizationId, id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowTagGroupWithOptions(organizationId *string, id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFlowTagGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFlowTagGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/flow/tagGroups/" + tea.StringValue(openapiutil.GetEncodeParam(id))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFlowTagGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHostGroup(organizationId *string, id *string) (_result *DeleteHostGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteHostGroupResponse{}
	_body, _err := client.DeleteHostGroupWithOptions(organizationId, id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteHostGroupWithOptions(organizationId *string, id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteHostGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHostGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/hostGroups/" + tea.StringValue(openapiutil.GetEncodeParam(id))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePipeline(organizationId *string, pipelineId *string) (_result *DeletePipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePipelineResponse{}
	_body, _err := client.DeletePipelineWithOptions(organizationId, pipelineId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePipelineWithOptions(organizationId *string, pipelineId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePipelineResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePipeline"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePipelineGroup(organizationId *string, groupId *string) (_result *DeletePipelineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePipelineGroupResponse{}
	_body, _err := client.DeletePipelineGroupWithOptions(organizationId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePipelineGroupWithOptions(organizationId *string, groupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePipelineGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePipelineGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelineGroups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePipelineGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProject(organizationId *string, request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectWithOptions(organizationId *string, request *DeleteProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		query["identifier"] = request.Identifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProject"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/projects/delete"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProtectedBranch(repositoryId *string, protectedBranchId *string, request *DeleteProtectedBranchRequest) (_result *DeleteProtectedBranchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProtectedBranchResponse{}
	_body, _err := client.DeleteProtectedBranchWithOptions(repositoryId, protectedBranchId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProtectedBranchWithOptions(repositoryId *string, protectedBranchId *string, request *DeleteProtectedBranchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProtectedBranchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProtectedBranch"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/protect_branches/" + tea.StringValue(openapiutil.GetEncodeParam(protectedBranchId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProtectedBranchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRepository(repositoryId *string, request *DeleteRepositoryRequest) (_result *DeleteRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRepositoryResponse{}
	_body, _err := client.DeleteRepositoryWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRepositoryWithOptions(repositoryId *string, request *DeleteRepositoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		body["reason"] = request.Reason
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRepository"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteResourceMember(organizationId *string, resourceType *string, resourceId *string, accountId *string) (_result *DeleteResourceMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceMemberResponse{}
	_body, _err := client.DeleteResourceMemberWithOptions(organizationId, resourceType, resourceId, accountId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteResourceMemberWithOptions(organizationId *string, resourceType *string, resourceId *string, accountId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteResourceMemberResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceMember"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(resourceType)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(resourceId)) + "/members/" + tea.StringValue(openapiutil.GetEncodeParam(accountId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVariableGroup(organizationId *string, id *string) (_result *DeleteVariableGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteVariableGroupResponse{}
	_body, _err := client.DeleteVariableGroupWithOptions(organizationId, id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVariableGroupWithOptions(organizationId *string, id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteVariableGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVariableGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/variableGroups/" + tea.StringValue(openapiutil.GetEncodeParam(id))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVariableGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWorkitemAllComment(organizationId *string, request *DeleteWorkitemAllCommentRequest) (_result *DeleteWorkitemAllCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteWorkitemAllCommentResponse{}
	_body, _err := client.DeleteWorkitemAllCommentWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWorkitemAllCommentWithOptions(organizationId *string, request *DeleteWorkitemAllCommentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteWorkitemAllCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		query["identifier"] = request.Identifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWorkitemAllComment"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/workitems/deleteAllComment"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWorkitemAllCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWorkitemComment(organizationId *string, request *DeleteWorkitemCommentRequest) (_result *DeleteWorkitemCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteWorkitemCommentResponse{}
	_body, _err := client.DeleteWorkitemCommentWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWorkitemCommentWithOptions(organizationId *string, request *DeleteWorkitemCommentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteWorkitemCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommentId)) {
		body["commentId"] = request.CommentId
	}

	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		body["identifier"] = request.Identifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWorkitemComment"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/workitems/deleteComent"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWorkitemCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FrozenWorkspace(workspaceId *string) (_result *FrozenWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FrozenWorkspaceResponse{}
	_body, _err := client.FrozenWorkspaceWithOptions(workspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FrozenWorkspaceWithOptions(workspaceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FrozenWorkspaceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("FrozenWorkspace"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/frozen"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FrozenWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBranchInfo(repositoryId *string, request *GetBranchInfoRequest) (_result *GetBranchInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBranchInfoResponse{}
	_body, _err := client.GetBranchInfoWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBranchInfoWithOptions(repositoryId *string, request *GetBranchInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBranchInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.BranchName)) {
		query["branchName"] = request.BranchName
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBranchInfo"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/branches/detail"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBranchInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCodeupOrganization(identity *string, request *GetCodeupOrganizationRequest) (_result *GetCodeupOrganizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCodeupOrganizationResponse{}
	_body, _err := client.GetCodeupOrganizationWithOptions(identity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCodeupOrganizationWithOptions(identity *string, request *GetCodeupOrganizationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCodeupOrganizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCodeupOrganization"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/organization/" + tea.StringValue(openapiutil.GetEncodeParam(identity))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCodeupOrganizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCustomFieldOption(organizationId *string, fieldId *string, request *GetCustomFieldOptionRequest) (_result *GetCustomFieldOptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCustomFieldOptionResponse{}
	_body, _err := client.GetCustomFieldOptionWithOptions(organizationId, fieldId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCustomFieldOptionWithOptions(organizationId *string, fieldId *string, request *GetCustomFieldOptionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCustomFieldOptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceIdentifier)) {
		query["spaceIdentifier"] = request.SpaceIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceType)) {
		query["spaceType"] = request.SpaceType
	}

	if !tea.BoolValue(util.IsUnset(request.WorkitemTypeIdentifier)) {
		query["workitemTypeIdentifier"] = request.WorkitemTypeIdentifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCustomFieldOption"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/fields/" + tea.StringValue(openapiutil.GetEncodeParam(fieldId)) + "/getCustomOption"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCustomFieldOptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFileBlobs(repositoryId *string, request *GetFileBlobsRequest) (_result *GetFileBlobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFileBlobsResponse{}
	_body, _err := client.GetFileBlobsWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileBlobsWithOptions(repositoryId *string, request *GetFileBlobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFileBlobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		query["filePath"] = request.FilePath
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["from"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Ref)) {
		query["ref"] = request.Ref
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["to"] = request.To
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFileBlobs"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/files/blobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileBlobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFileLastCommit(repositoryId *string, request *GetFileLastCommitRequest) (_result *GetFileLastCommitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFileLastCommitResponse{}
	_body, _err := client.GetFileLastCommitWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileLastCommitWithOptions(repositoryId *string, request *GetFileLastCommitRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFileLastCommitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.Filepath)) {
		query["filepath"] = request.Filepath
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Sha)) {
		query["sha"] = request.Sha
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSignature)) {
		query["showSignature"] = request.ShowSignature
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFileLastCommit"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/files/lastCommit"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileLastCommitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFlowTagGroup(organizationId *string, id *string) (_result *GetFlowTagGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFlowTagGroupResponse{}
	_body, _err := client.GetFlowTagGroupWithOptions(organizationId, id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFlowTagGroupWithOptions(organizationId *string, id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFlowTagGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetFlowTagGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/flow/tagGroups/" + tea.StringValue(openapiutil.GetEncodeParam(id))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFlowTagGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHostGroup(organizationId *string, id *string) (_result *GetHostGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHostGroupResponse{}
	_body, _err := client.GetHostGroupWithOptions(organizationId, id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHostGroupWithOptions(organizationId *string, id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHostGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetHostGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/hostGroups/" + tea.StringValue(openapiutil.GetEncodeParam(id))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrganizationMember(organizationId *string, accountId *string) (_result *GetOrganizationMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOrganizationMemberResponse{}
	_body, _err := client.GetOrganizationMemberWithOptions(organizationId, accountId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrganizationMemberWithOptions(organizationId *string, accountId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetOrganizationMemberResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetOrganizationMember"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/members/" + tea.StringValue(openapiutil.GetEncodeParam(accountId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrganizationMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPipeline(organizationId *string, pipelineId *string) (_result *GetPipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPipelineResponse{}
	_body, _err := client.GetPipelineWithOptions(organizationId, pipelineId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipelineWithOptions(organizationId *string, pipelineId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPipelineResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipeline"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPipelineArtifactUrl(organizationId *string, request *GetPipelineArtifactUrlRequest) (_result *GetPipelineArtifactUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPipelineArtifactUrlResponse{}
	_body, _err := client.GetPipelineArtifactUrlWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipelineArtifactUrlWithOptions(organizationId *string, request *GetPipelineArtifactUrlRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPipelineArtifactUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		query["filePath"] = request.FilePath
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineArtifactUrl"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipeline/getArtifactDownloadUrl"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineArtifactUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPipelineEmasArtifactUrl(organizationId *string, emasJobInstanceId *string, md5 *string, pipelineId *string, pipelineRunId *string, request *GetPipelineEmasArtifactUrlRequest) (_result *GetPipelineEmasArtifactUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPipelineEmasArtifactUrlResponse{}
	_body, _err := client.GetPipelineEmasArtifactUrlWithOptions(organizationId, emasJobInstanceId, md5, pipelineId, pipelineRunId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipelineEmasArtifactUrlWithOptions(organizationId *string, emasJobInstanceId *string, md5 *string, pipelineId *string, pipelineRunId *string, request *GetPipelineEmasArtifactUrlRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPipelineEmasArtifactUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceConnectionId)) {
		query["serviceConnectionId"] = request.ServiceConnectionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineEmasArtifactUrl"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipeline/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/pipelineRun/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineRunId)) + "/emas/artifact/" + tea.StringValue(openapiutil.GetEncodeParam(emasJobInstanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(md5))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineEmasArtifactUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPipelineGroup(organizationId *string, groupId *string) (_result *GetPipelineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPipelineGroupResponse{}
	_body, _err := client.GetPipelineGroupWithOptions(organizationId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipelineGroupWithOptions(organizationId *string, groupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPipelineGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelineGroups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPipelineRun(organizationId *string, pipelineId *string, pipelineRunId *string) (_result *GetPipelineRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPipelineRunResponse{}
	_body, _err := client.GetPipelineRunWithOptions(organizationId, pipelineId, pipelineRunId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipelineRunWithOptions(organizationId *string, pipelineId *string, pipelineRunId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPipelineRunResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineRun"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/pipelineRuns/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineRunId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPipelineScanReportUrl(organizationId *string, request *GetPipelineScanReportUrlRequest) (_result *GetPipelineScanReportUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPipelineScanReportUrlResponse{}
	_body, _err := client.GetPipelineScanReportUrlWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipelineScanReportUrlWithOptions(organizationId *string, request *GetPipelineScanReportUrlRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPipelineScanReportUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReportPath)) {
		body["reportPath"] = request.ReportPath
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineScanReportUrl"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipeline/getPipelineScanReportUrl"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineScanReportUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProjectInfo(organizationId *string, projectId *string) (_result *GetProjectInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectInfoResponse{}
	_body, _err := client.GetProjectInfoWithOptions(organizationId, projectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectInfoWithOptions(organizationId *string, projectId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProjectInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetProjectInfo"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/project/" + tea.StringValue(openapiutil.GetEncodeParam(projectId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProjectMember(request *GetProjectMemberRequest) (_result *GetProjectMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectMemberResponse{}
	_body, _err := client.GetProjectMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectMemberWithOptions(request *GetProjectMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProjectMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.RepositoryId)) {
		query["repositoryId"] = request.RepositoryId
	}

	if !tea.BoolValue(util.IsUnset(request.UserAliyunPk)) {
		query["userAliyunPk"] = request.UserAliyunPk
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProjectMember"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/member/get"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRepository(request *GetRepositoryRequest) (_result *GetRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRepositoryResponse{}
	_body, _err := client.GetRepositoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRepositoryWithOptions(request *GetRepositoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.Identity)) {
		query["identity"] = request.Identity
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRepository"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/get"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSprintInfo(organizationId *string, sprintId *string) (_result *GetSprintInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSprintInfoResponse{}
	_body, _err := client.GetSprintInfoWithOptions(organizationId, sprintId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSprintInfoWithOptions(organizationId *string, sprintId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSprintInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSprintInfo"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/sprints/" + tea.StringValue(openapiutil.GetEncodeParam(sprintId)) + "/getSprintinfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSprintInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVMDeployOrder(organizationId *string, pipelineId *string, deployOrderId *string) (_result *GetVMDeployOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVMDeployOrderResponse{}
	_body, _err := client.GetVMDeployOrderWithOptions(organizationId, pipelineId, deployOrderId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVMDeployOrderWithOptions(organizationId *string, pipelineId *string, deployOrderId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetVMDeployOrderResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetVMDeployOrder"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/deploy/" + tea.StringValue(openapiutil.GetEncodeParam(deployOrderId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVMDeployOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVariableGroup(organizationId *string, id *string) (_result *GetVariableGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVariableGroupResponse{}
	_body, _err := client.GetVariableGroupWithOptions(organizationId, id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVariableGroupWithOptions(organizationId *string, id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetVariableGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetVariableGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/variableGroups/" + tea.StringValue(openapiutil.GetEncodeParam(id))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVariableGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkItemActivity(organizationId *string, workitemId *string) (_result *GetWorkItemActivityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkItemActivityResponse{}
	_body, _err := client.GetWorkItemActivityWithOptions(organizationId, workitemId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkItemActivityWithOptions(organizationId *string, workitemId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetWorkItemActivityResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkItemActivity"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/workitems/" + tea.StringValue(openapiutil.GetEncodeParam(workitemId)) + "/getActivity"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkItemActivityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkItemInfo(organizationId *string, workitemId *string) (_result *GetWorkItemInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkItemInfoResponse{}
	_body, _err := client.GetWorkItemInfoWithOptions(organizationId, workitemId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkItemInfoWithOptions(organizationId *string, workitemId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetWorkItemInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkItemInfo"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/workitems/" + tea.StringValue(openapiutil.GetEncodeParam(workitemId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkItemInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkItemWorkFlowInfo(organizationId *string, workitemId *string, request *GetWorkItemWorkFlowInfoRequest) (_result *GetWorkItemWorkFlowInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkItemWorkFlowInfoResponse{}
	_body, _err := client.GetWorkItemWorkFlowInfoWithOptions(organizationId, workitemId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkItemWorkFlowInfoWithOptions(organizationId *string, workitemId *string, request *GetWorkItemWorkFlowInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetWorkItemWorkFlowInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigurationId)) {
		query["configurationId"] = request.ConfigurationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkItemWorkFlowInfo"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/workitems/" + tea.StringValue(openapiutil.GetEncodeParam(workitemId)) + "/getWorkflowInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkItemWorkFlowInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkitemCommentList(organizationId *string, workitemId *string) (_result *GetWorkitemCommentListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkitemCommentListResponse{}
	_body, _err := client.GetWorkitemCommentListWithOptions(organizationId, workitemId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkitemCommentListWithOptions(organizationId *string, workitemId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetWorkitemCommentListResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkitemCommentList"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/workitems/" + tea.StringValue(openapiutil.GetEncodeParam(workitemId)) + "/commentList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkitemCommentListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkitemRelations(organizationId *string, workitemId *string, request *GetWorkitemRelationsRequest) (_result *GetWorkitemRelationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkitemRelationsResponse{}
	_body, _err := client.GetWorkitemRelationsWithOptions(organizationId, workitemId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkitemRelationsWithOptions(organizationId *string, workitemId *string, request *GetWorkitemRelationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetWorkitemRelationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		query["relationType"] = request.RelationType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkitemRelations"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/workitems/" + tea.StringValue(openapiutil.GetEncodeParam(workitemId)) + "/getRelations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkitemRelationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkitemTimeTypeList(organizationId *string) (_result *GetWorkitemTimeTypeListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkitemTimeTypeListResponse{}
	_body, _err := client.GetWorkitemTimeTypeListWithOptions(organizationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkitemTimeTypeListWithOptions(organizationId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetWorkitemTimeTypeListResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkitemTimeTypeList"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/workitems/type/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkitemTimeTypeListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkspace(workspaceId *string) (_result *GetWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(workspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkspaceWithOptions(workspaceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkspace"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinPipelineGroup(organizationId *string, request *JoinPipelineGroupRequest) (_result *JoinPipelineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &JoinPipelineGroupResponse{}
	_body, _err := client.JoinPipelineGroupWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinPipelineGroupWithOptions(organizationId *string, request *JoinPipelineGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *JoinPipelineGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["groupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineIds)) {
		query["pipelineIds"] = request.PipelineIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("JoinPipelineGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelineGroups/join"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &JoinPipelineGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowTagGroups(organizationId *string) (_result *ListFlowTagGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFlowTagGroupsResponse{}
	_body, _err := client.ListFlowTagGroupsWithOptions(organizationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowTagGroupsWithOptions(organizationId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFlowTagGroupsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlowTagGroups"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/flow/tagGroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowTagGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHostGroups(organizationId *string, request *ListHostGroupsRequest) (_result *ListHostGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListHostGroupsResponse{}
	_body, _err := client.ListHostGroupsWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHostGroupsWithOptions(organizationId *string, request *ListHostGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListHostGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateEndTime)) {
		query["createEndTime"] = request.CreateEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.CreateStartTime)) {
		query["createStartTime"] = request.CreateStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorAccountIds)) {
		query["creatorAccountIds"] = request.CreatorAccountIds
	}

	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageOrder)) {
		query["pageOrder"] = request.PageOrder
	}

	if !tea.BoolValue(util.IsUnset(request.PageSort)) {
		query["pageSort"] = request.PageSort
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHostGroups"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/hostGroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHostGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrganizationMembers(organizationId *string, request *ListOrganizationMembersRequest) (_result *ListOrganizationMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOrganizationMembersResponse{}
	_body, _err := client.ListOrganizationMembersWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrganizationMembersWithOptions(organizationId *string, request *ListOrganizationMembersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListOrganizationMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExternUid)) {
		query["externUid"] = request.ExternUid
	}

	if !tea.BoolValue(util.IsUnset(request.JoinTimeFrom)) {
		query["joinTimeFrom"] = request.JoinTimeFrom
	}

	if !tea.BoolValue(util.IsUnset(request.JoinTimeTo)) {
		query["joinTimeTo"] = request.JoinTimeTo
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationMemberName)) {
		query["organizationMemberName"] = request.OrganizationMemberName
	}

	if !tea.BoolValue(util.IsUnset(request.Provider)) {
		query["provider"] = request.Provider
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["state"] = request.State
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOrganizationMembers"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOrganizationMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPipelineGroupPipelines(organizationId *string, groupId *string, request *ListPipelineGroupPipelinesRequest) (_result *ListPipelineGroupPipelinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineGroupPipelinesResponse{}
	_body, _err := client.ListPipelineGroupPipelinesWithOptions(organizationId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPipelineGroupPipelinesWithOptions(organizationId *string, groupId *string, request *ListPipelineGroupPipelinesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelineGroupPipelinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateEndTime)) {
		query["createEndTime"] = request.CreateEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.CreateStartTime)) {
		query["createStartTime"] = request.CreateStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteEndTime)) {
		query["executeEndTime"] = request.ExecuteEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteStartTime)) {
		query["executeStartTime"] = request.ExecuteStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineName)) {
		query["pipelineName"] = request.PipelineName
	}

	if !tea.BoolValue(util.IsUnset(request.ResultStatusList)) {
		query["resultStatusList"] = request.ResultStatusList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelineGroupPipelines"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelineGroups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId)) + "/pipelines"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelineGroupPipelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPipelineGroups(organizationId *string, request *ListPipelineGroupsRequest) (_result *ListPipelineGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineGroupsResponse{}
	_body, _err := client.ListPipelineGroupsWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPipelineGroupsWithOptions(organizationId *string, request *ListPipelineGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelineGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelineGroups"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelineGroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelineGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPipelineJobHistorys(organizationId *string, pipelineId *string, request *ListPipelineJobHistorysRequest) (_result *ListPipelineJobHistorysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineJobHistorysResponse{}
	_body, _err := client.ListPipelineJobHistorysWithOptions(organizationId, pipelineId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPipelineJobHistorysWithOptions(organizationId *string, pipelineId *string, request *ListPipelineJobHistorysRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelineJobHistorysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		query["identifier"] = request.Identifier
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelineJobHistorys"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipeline/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/job/historys"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelineJobHistorysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPipelineJobs(organizationId *string, pipelineId *string, request *ListPipelineJobsRequest) (_result *ListPipelineJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineJobsResponse{}
	_body, _err := client.ListPipelineJobsWithOptions(organizationId, pipelineId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPipelineJobsWithOptions(organizationId *string, pipelineId *string, request *ListPipelineJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelineJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["category"] = request.Category
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelineJobs"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipeline/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelineJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPipelineRuns(organizationId *string, pipelineId *string, request *ListPipelineRunsRequest) (_result *ListPipelineRunsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineRunsResponse{}
	_body, _err := client.ListPipelineRunsWithOptions(organizationId, pipelineId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPipelineRunsWithOptions(organizationId *string, pipelineId *string, request *ListPipelineRunsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelineRunsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerMode)) {
		query["triggerMode"] = request.TriggerMode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelineRuns"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/pipelineRuns"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelineRunsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPipelines(organizationId *string, request *ListPipelinesRequest) (_result *ListPipelinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelinesResponse{}
	_body, _err := client.ListPipelinesWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPipelinesWithOptions(organizationId *string, request *ListPipelinesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateEndTime)) {
		query["createEndTime"] = request.CreateEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.CreateStartTime)) {
		query["createStartTime"] = request.CreateStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorAccountIds)) {
		query["creatorAccountIds"] = request.CreatorAccountIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteAccountIds)) {
		query["executeAccountIds"] = request.ExecuteAccountIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteEndTime)) {
		query["executeEndTime"] = request.ExecuteEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteStartTime)) {
		query["executeStartTime"] = request.ExecuteStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineName)) {
		query["pipelineName"] = request.PipelineName
	}

	if !tea.BoolValue(util.IsUnset(request.StatusList)) {
		query["statusList"] = request.StatusList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelines"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjectMembers(organizationId *string, projectId *string, request *ListProjectMembersRequest) (_result *ListProjectMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectMembersResponse{}
	_body, _err := client.ListProjectMembersWithOptions(organizationId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectMembersWithOptions(organizationId *string, projectId *string, request *ListProjectMembersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["targetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjectMembers"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/listMembers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjectTemplates(organizationId *string, request *ListProjectTemplatesRequest) (_result *ListProjectTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectTemplatesResponse{}
	_body, _err := client.ListProjectTemplatesWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectTemplatesWithOptions(organizationId *string, request *ListProjectTemplatesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["category"] = request.Category
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjectTemplates"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/projects/listTemplates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjectWorkitemTypes(organizationId *string, projectId *string, request *ListProjectWorkitemTypesRequest) (_result *ListProjectWorkitemTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectWorkitemTypesResponse{}
	_body, _err := client.ListProjectWorkitemTypesWithOptions(organizationId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectWorkitemTypesWithOptions(organizationId *string, projectId *string, request *ListProjectWorkitemTypesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectWorkitemTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceType)) {
		query["spaceType"] = request.SpaceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjectWorkitemTypes"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/getWorkitemType"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectWorkitemTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjects(organizationId *string, request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectsWithOptions(organizationId *string, request *ListProjectsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.Conditions)) {
		query["conditions"] = request.Conditions
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraConditions)) {
		query["extraConditions"] = request.ExtraConditions
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["scope"] = request.Scope
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjects"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/listProjects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProtectedBranches(repositoryId *string, request *ListProtectedBranchesRequest) (_result *ListProtectedBranchesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProtectedBranchesResponse{}
	_body, _err := client.ListProtectedBranchesWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProtectedBranchesWithOptions(repositoryId *string, request *ListProtectedBranchesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProtectedBranchesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProtectedBranches"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/protect_branches"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProtectedBranchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositories(request *ListRepositoriesRequest) (_result *ListRepositoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoriesResponse{}
	_body, _err := client.ListRepositoriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoriesWithOptions(request *ListRepositoriesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.Archived)) {
		query["archived"] = request.Archived
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["orderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PerPage)) {
		query["perPage"] = request.PerPage
	}

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["search"] = request.Search
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["sort"] = request.Sort
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepositories"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositoryBranches(repositoryId *string, request *ListRepositoryBranchesRequest) (_result *ListRepositoryBranchesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoryBranchesResponse{}
	_body, _err := client.ListRepositoryBranchesWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoryBranchesWithOptions(repositoryId *string, request *ListRepositoryBranchesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoryBranchesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["search"] = request.Search
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["sort"] = request.Sort
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepositoryBranches"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/branches"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoryBranchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositoryCommitDiff(repositoryId *string, sha *string, request *ListRepositoryCommitDiffRequest) (_result *ListRepositoryCommitDiffResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoryCommitDiffResponse{}
	_body, _err := client.ListRepositoryCommitDiffWithOptions(repositoryId, sha, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoryCommitDiffWithOptions(repositoryId *string, sha *string, request *ListRepositoryCommitDiffRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoryCommitDiffResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.ContextLine)) {
		query["contextLine"] = request.ContextLine
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepositoryCommitDiff"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/commits/" + tea.StringValue(openapiutil.GetEncodeParam(sha)) + "/diff"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoryCommitDiffResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositoryMemberWithInherited(repositoryId *string, request *ListRepositoryMemberWithInheritedRequest) (_result *ListRepositoryMemberWithInheritedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoryMemberWithInheritedResponse{}
	_body, _err := client.ListRepositoryMemberWithInheritedWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoryMemberWithInheritedWithOptions(repositoryId *string, request *ListRepositoryMemberWithInheritedRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoryMemberWithInheritedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepositoryMemberWithInherited"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/members/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoryMemberWithInheritedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositoryTree(repositoryId *string, request *ListRepositoryTreeRequest) (_result *ListRepositoryTreeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoryTreeResponse{}
	_body, _err := client.ListRepositoryTreeWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoryTreeWithOptions(repositoryId *string, request *ListRepositoryTreeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoryTreeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.RefName)) {
		query["refName"] = request.RefName
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepositoryTree"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/files/tree"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoryTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositoryWebhook(repositoryId *string, request *ListRepositoryWebhookRequest) (_result *ListRepositoryWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoryWebhookResponse{}
	_body, _err := client.ListRepositoryWebhookWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoryWebhookWithOptions(repositoryId *string, request *ListRepositoryWebhookRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoryWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepositoryWebhook"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/webhooks/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoryWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceMembers(organizationId *string, resourceType *string, resourceId *string) (_result *ListResourceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceMembersResponse{}
	_body, _err := client.ListResourceMembersWithOptions(organizationId, resourceType, resourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceMembersWithOptions(organizationId *string, resourceType *string, resourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourceMembersResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceMembers"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(resourceType)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(resourceId)) + "/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceConnections(organizationId *string, request *ListServiceConnectionsRequest) (_result *ListServiceConnectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServiceConnectionsResponse{}
	_body, _err := client.ListServiceConnectionsWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServiceConnectionsWithOptions(organizationId *string, request *ListServiceConnectionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListServiceConnectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SericeConnectionType)) {
		query["sericeConnectionType"] = request.SericeConnectionType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceConnections"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/serviceConnections"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceConnectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSprints(organizationId *string, request *ListSprintsRequest) (_result *ListSprintsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSprintsResponse{}
	_body, _err := client.ListSprintsWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSprintsWithOptions(organizationId *string, request *ListSprintsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSprintsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceIdentifier)) {
		query["spaceIdentifier"] = request.SpaceIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceType)) {
		query["spaceType"] = request.SpaceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSprints"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/sprints/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSprintsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVariableGroups(organizationId *string, request *ListVariableGroupsRequest) (_result *ListVariableGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListVariableGroupsResponse{}
	_body, _err := client.ListVariableGroupsWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVariableGroupsWithOptions(organizationId *string, request *ListVariableGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListVariableGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageOrder)) {
		query["pageOrder"] = request.PageOrder
	}

	if !tea.BoolValue(util.IsUnset(request.PageSort)) {
		query["pageSort"] = request.PageSort
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVariableGroups"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/variableGroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVariableGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWorkItemAllFields(organizationId *string, request *ListWorkItemAllFieldsRequest) (_result *ListWorkItemAllFieldsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkItemAllFieldsResponse{}
	_body, _err := client.ListWorkItemAllFieldsWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWorkItemAllFieldsWithOptions(organizationId *string, request *ListWorkItemAllFieldsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListWorkItemAllFieldsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceIdentifier)) {
		query["spaceIdentifier"] = request.SpaceIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceType)) {
		query["spaceType"] = request.SpaceType
	}

	if !tea.BoolValue(util.IsUnset(request.WorkitemTypeIdentifier)) {
		query["workitemTypeIdentifier"] = request.WorkitemTypeIdentifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkItemAllFields"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/workitems/fields/listAll"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkItemAllFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWorkItemWorkFlowStatus(organizationId *string, request *ListWorkItemWorkFlowStatusRequest) (_result *ListWorkItemWorkFlowStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkItemWorkFlowStatusResponse{}
	_body, _err := client.ListWorkItemWorkFlowStatusWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWorkItemWorkFlowStatusWithOptions(organizationId *string, request *ListWorkItemWorkFlowStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListWorkItemWorkFlowStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceIdentifier)) {
		query["spaceIdentifier"] = request.SpaceIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceType)) {
		query["spaceType"] = request.SpaceType
	}

	if !tea.BoolValue(util.IsUnset(request.WorkitemCategoryIdentifier)) {
		query["workitemCategoryIdentifier"] = request.WorkitemCategoryIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.WorkitemTypeIdentifier)) {
		query["workitemTypeIdentifier"] = request.WorkitemTypeIdentifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkItemWorkFlowStatus"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/workitems/workflow/listWorkflowStatus"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkItemWorkFlowStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWorkitemTime(organizationId *string, workitemId *string) (_result *ListWorkitemTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkitemTimeResponse{}
	_body, _err := client.ListWorkitemTimeWithOptions(organizationId, workitemId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWorkitemTimeWithOptions(organizationId *string, workitemId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListWorkitemTimeResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkitemTime"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/workitems/" + tea.StringValue(openapiutil.GetEncodeParam(workitemId)) + "/time/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkitemTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWorkitems(organizationId *string, request *ListWorkitemsRequest) (_result *ListWorkitemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkitemsResponse{}
	_body, _err := client.ListWorkitemsWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWorkitemsWithOptions(organizationId *string, request *ListWorkitemsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListWorkitemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.Conditions)) {
		query["conditions"] = request.Conditions
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraConditions)) {
		query["extraConditions"] = request.ExtraConditions
	}

	if !tea.BoolValue(util.IsUnset(request.GroupCondition)) {
		query["groupCondition"] = request.GroupCondition
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["orderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.SearchType)) {
		query["searchType"] = request.SearchType
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceIdentifier)) {
		query["spaceIdentifier"] = request.SpaceIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceType)) {
		query["spaceType"] = request.SpaceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkitems"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/listWorkitems"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkitemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (_result *ListWorkspacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkspacesResponse{}
	_body, _err := client.ListWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWorkspacesWithOptions(tmpReq *ListWorkspacesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListWorkspacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.StatusList)) {
		request.StatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StatusList, tea.String("statusList"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.WorkspaceTemplateList)) {
		request.WorkspaceTemplateListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WorkspaceTemplateList, tea.String("workspaceTemplateList"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.StatusListShrink)) {
		query["statusList"] = request.StatusListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceTemplateListShrink)) {
		query["workspaceTemplateList"] = request.WorkspaceTemplateListShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkspaces"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/workspaces"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LogPipelineJobRun(organizationId *string, pipelineId *string, jobId *string, pipelineRunId *string) (_result *LogPipelineJobRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &LogPipelineJobRunResponse{}
	_body, _err := client.LogPipelineJobRunWithOptions(organizationId, pipelineId, jobId, pipelineRunId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LogPipelineJobRunWithOptions(organizationId *string, pipelineId *string, jobId *string, pipelineRunId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *LogPipelineJobRunResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("LogPipelineJobRun"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipeline/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/pipelineRun/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineRunId)) + "/job/" + tea.StringValue(openapiutil.GetEncodeParam(jobId)) + "/logs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &LogPipelineJobRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LogVMDeployMachine(organizationId *string, pipelineId *string, deployOrderId *string, machineSn *string) (_result *LogVMDeployMachineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &LogVMDeployMachineResponse{}
	_body, _err := client.LogVMDeployMachineWithOptions(organizationId, pipelineId, deployOrderId, machineSn, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LogVMDeployMachineWithOptions(organizationId *string, pipelineId *string, deployOrderId *string, machineSn *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *LogVMDeployMachineResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("LogVMDeployMachine"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/deploy/" + tea.StringValue(openapiutil.GetEncodeParam(deployOrderId)) + "/machine/" + tea.StringValue(openapiutil.GetEncodeParam(machineSn)) + "/log"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &LogVMDeployMachineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PassPipelineValidate(organizationId *string, pipelineId *string, pipelineRunId *string, jobId *string) (_result *PassPipelineValidateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PassPipelineValidateResponse{}
	_body, _err := client.PassPipelineValidateWithOptions(organizationId, pipelineId, pipelineRunId, jobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PassPipelineValidateWithOptions(organizationId *string, pipelineId *string, pipelineRunId *string, jobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PassPipelineValidateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("PassPipelineValidate"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/pipelineRuns/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineRunId)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId)) + "/pass"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PassPipelineValidateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefusePipelineValidate(organizationId *string, pipelineId *string, pipelineRunId *string, jobId *string) (_result *RefusePipelineValidateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RefusePipelineValidateResponse{}
	_body, _err := client.RefusePipelineValidateWithOptions(organizationId, pipelineId, pipelineRunId, jobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefusePipelineValidateWithOptions(organizationId *string, pipelineId *string, pipelineRunId *string, jobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RefusePipelineValidateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RefusePipelineValidate"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/pipelineRuns/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineRunId)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId)) + "/refuse"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RefusePipelineValidateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseWorkspace(workspaceId *string) (_result *ReleaseWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReleaseWorkspaceResponse{}
	_body, _err := client.ReleaseWorkspaceWithOptions(workspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseWorkspaceWithOptions(workspaceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReleaseWorkspaceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseWorkspace"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/release"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetSshKey(organizationId *string) (_result *ResetSshKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResetSshKeyResponse{}
	_body, _err := client.ResetSshKeyWithOptions(organizationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetSshKeyWithOptions(organizationId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ResetSshKeyResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ResetSshKey"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/sshKey"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetSshKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeVMDeployOrder(organizationId *string, pipelineId *string, deployOrderId *string) (_result *ResumeVMDeployOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeVMDeployOrderResponse{}
	_body, _err := client.ResumeVMDeployOrderWithOptions(organizationId, pipelineId, deployOrderId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeVMDeployOrderWithOptions(organizationId *string, pipelineId *string, deployOrderId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ResumeVMDeployOrderResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeVMDeployOrder"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/deploy/" + tea.StringValue(openapiutil.GetEncodeParam(deployOrderId)) + "/resume"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeVMDeployOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RetryPipelineJobRun(organizationId *string, pipelineId *string, pipelineRunId *string, jobId *string) (_result *RetryPipelineJobRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RetryPipelineJobRunResponse{}
	_body, _err := client.RetryPipelineJobRunWithOptions(organizationId, pipelineId, pipelineRunId, jobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RetryPipelineJobRunWithOptions(organizationId *string, pipelineId *string, pipelineRunId *string, jobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RetryPipelineJobRunResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RetryPipelineJobRun"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/pipelineRuns/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineRunId)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RetryPipelineJobRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RetryVMDeployMachine(organizationId *string, pipelineId *string, deployOrderId *string, machineSn *string) (_result *RetryVMDeployMachineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RetryVMDeployMachineResponse{}
	_body, _err := client.RetryVMDeployMachineWithOptions(organizationId, pipelineId, deployOrderId, machineSn, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RetryVMDeployMachineWithOptions(organizationId *string, pipelineId *string, deployOrderId *string, machineSn *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RetryVMDeployMachineResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RetryVMDeployMachine"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/deploy/" + tea.StringValue(openapiutil.GetEncodeParam(deployOrderId)) + "/machine/" + tea.StringValue(openapiutil.GetEncodeParam(machineSn)) + "/retry"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RetryVMDeployMachineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SkipPipelineJobRun(organizationId *string, pipelineId *string, pipelineRunId *string, jobId *string) (_result *SkipPipelineJobRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SkipPipelineJobRunResponse{}
	_body, _err := client.SkipPipelineJobRunWithOptions(organizationId, pipelineId, pipelineRunId, jobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SkipPipelineJobRunWithOptions(organizationId *string, pipelineId *string, pipelineRunId *string, jobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SkipPipelineJobRunResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("SkipPipelineJobRun"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/pipelineRuns/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineRunId)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId)) + "/skip"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SkipPipelineJobRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SkipVMDeployMachine(organizationId *string, pipelineId *string, deployOrderId *string, machineSn *string) (_result *SkipVMDeployMachineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SkipVMDeployMachineResponse{}
	_body, _err := client.SkipVMDeployMachineWithOptions(organizationId, pipelineId, deployOrderId, machineSn, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SkipVMDeployMachineWithOptions(organizationId *string, pipelineId *string, deployOrderId *string, machineSn *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SkipVMDeployMachineResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("SkipVMDeployMachine"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/deploy/" + tea.StringValue(openapiutil.GetEncodeParam(deployOrderId)) + "/machine/" + tea.StringValue(openapiutil.GetEncodeParam(machineSn)) + "/skip"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SkipVMDeployMachineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartPipelineRun(organizationId *string, pipelineId *string, request *StartPipelineRunRequest) (_result *StartPipelineRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartPipelineRunResponse{}
	_body, _err := client.StartPipelineRunWithOptions(organizationId, pipelineId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartPipelineRunWithOptions(organizationId *string, pipelineId *string, request *StartPipelineRunRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartPipelineRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartPipelineRun"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organizations/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/run"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartPipelineRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopPipelineJobRun(organizationId *string, pipelineId *string, pipelineRunId *string, jobId *string) (_result *StopPipelineJobRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopPipelineJobRunResponse{}
	_body, _err := client.StopPipelineJobRunWithOptions(organizationId, pipelineId, pipelineRunId, jobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopPipelineJobRunWithOptions(organizationId *string, pipelineId *string, pipelineRunId *string, jobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopPipelineJobRunResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopPipelineJobRun"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/pipelineRuns/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineRunId)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId)) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopPipelineJobRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopPipelineRun(organizationId *string, pipelineId *string, pipelineRunId *string) (_result *StopPipelineRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopPipelineRunResponse{}
	_body, _err := client.StopPipelineRunWithOptions(organizationId, pipelineId, pipelineRunId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopPipelineRunWithOptions(organizationId *string, pipelineId *string, pipelineRunId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopPipelineRunResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopPipelineRun"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/pipelineRuns/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineRunId)) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopPipelineRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopVMDeployOrder(organizationId *string, pipelineId *string, deployOrderId *string) (_result *StopVMDeployOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopVMDeployOrderResponse{}
	_body, _err := client.StopVMDeployOrderWithOptions(organizationId, pipelineId, deployOrderId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopVMDeployOrderWithOptions(organizationId *string, pipelineId *string, deployOrderId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopVMDeployOrderResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopVMDeployOrder"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/deploy/" + tea.StringValue(openapiutil.GetEncodeParam(deployOrderId)) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopVMDeployOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TriggerRepositoryMirrorSync(repositoryId *string, request *TriggerRepositoryMirrorSyncRequest) (_result *TriggerRepositoryMirrorSyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TriggerRepositoryMirrorSyncResponse{}
	_body, _err := client.TriggerRepositoryMirrorSyncWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TriggerRepositoryMirrorSyncWithOptions(repositoryId *string, request *TriggerRepositoryMirrorSyncRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TriggerRepositoryMirrorSyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TriggerRepositoryMirrorSync"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/mirror"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TriggerRepositoryMirrorSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFile(repositoryId *string, request *UpdateFileRequest) (_result *UpdateFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFileResponse{}
	_body, _err := client.UpdateFileWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFileWithOptions(repositoryId *string, request *UpdateFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BranchName)) {
		body["branchName"] = request.BranchName
	}

	if !tea.BoolValue(util.IsUnset(request.CommitMessage)) {
		body["commitMessage"] = request.CommitMessage
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Encoding)) {
		body["encoding"] = request.Encoding
	}

	if !tea.BoolValue(util.IsUnset(request.NewPath)) {
		body["newPath"] = request.NewPath
	}

	if !tea.BoolValue(util.IsUnset(request.OldPath)) {
		body["oldPath"] = request.OldPath
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFile"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/files/update"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFlowTag(organizationId *string, id *string, request *UpdateFlowTagRequest) (_result *UpdateFlowTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFlowTagResponse{}
	_body, _err := client.UpdateFlowTagWithOptions(organizationId, id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFlowTagWithOptions(organizationId *string, id *string, request *UpdateFlowTagRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFlowTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Color)) {
		query["color"] = request.Color
	}

	if !tea.BoolValue(util.IsUnset(request.FlowTagGroupId)) {
		query["flowTagGroupId"] = request.FlowTagGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFlowTag"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/flow/tags/" + tea.StringValue(openapiutil.GetEncodeParam(id))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFlowTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFlowTagGroup(organizationId *string, id *string, request *UpdateFlowTagGroupRequest) (_result *UpdateFlowTagGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFlowTagGroupResponse{}
	_body, _err := client.UpdateFlowTagGroupWithOptions(organizationId, id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFlowTagGroupWithOptions(organizationId *string, id *string, request *UpdateFlowTagGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFlowTagGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFlowTagGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/flow/tagGroups/" + tea.StringValue(openapiutil.GetEncodeParam(id))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFlowTagGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateHostGroup(organizationId *string, id *string, request *UpdateHostGroupRequest) (_result *UpdateHostGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateHostGroupResponse{}
	_body, _err := client.UpdateHostGroupWithOptions(organizationId, id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateHostGroupWithOptions(organizationId *string, id *string, request *UpdateHostGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateHostGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunRegion)) {
		body["aliyunRegion"] = request.AliyunRegion
	}

	if !tea.BoolValue(util.IsUnset(request.EcsLabelKey)) {
		body["ecsLabelKey"] = request.EcsLabelKey
	}

	if !tea.BoolValue(util.IsUnset(request.EcsLabelValue)) {
		body["ecsLabelValue"] = request.EcsLabelValue
	}

	if !tea.BoolValue(util.IsUnset(request.EcsType)) {
		body["ecsType"] = request.EcsType
	}

	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		body["envId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.MachineInfos)) {
		body["machineInfos"] = request.MachineInfos
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceConnectionId)) {
		body["serviceConnectionId"] = request.ServiceConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.TagIds)) {
		body["tagIds"] = request.TagIds
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateHostGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/hostGroups/" + tea.StringValue(openapiutil.GetEncodeParam(id))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePipelineBaseInfo(organizationId *string, pipelineId *string, request *UpdatePipelineBaseInfoRequest) (_result *UpdatePipelineBaseInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePipelineBaseInfoResponse{}
	_body, _err := client.UpdatePipelineBaseInfoWithOptions(organizationId, pipelineId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePipelineBaseInfoWithOptions(organizationId *string, pipelineId *string, request *UpdatePipelineBaseInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdatePipelineBaseInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvId)) {
		query["envId"] = request.EnvId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineName)) {
		query["pipelineName"] = request.PipelineName
	}

	if !tea.BoolValue(util.IsUnset(request.TagList)) {
		query["tagList"] = request.TagList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePipelineBaseInfo"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(pipelineId)) + "/baseInfo"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePipelineBaseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePipelineGroup(organizationId *string, groupId *string, request *UpdatePipelineGroupRequest) (_result *UpdatePipelineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePipelineGroupResponse{}
	_body, _err := client.UpdatePipelineGroupWithOptions(organizationId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePipelineGroupWithOptions(organizationId *string, groupId *string, request *UpdatePipelineGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdatePipelineGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePipelineGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/pipelineGroups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePipelineGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProjectMember(organizationId *string, projectId *string, request *UpdateProjectMemberRequest) (_result *UpdateProjectMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProjectMemberResponse{}
	_body, _err := client.UpdateProjectMemberWithOptions(organizationId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProjectMemberWithOptions(organizationId *string, projectId *string, request *UpdateProjectMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProjectMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleIdentifier)) {
		body["roleIdentifier"] = request.RoleIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.TargetIdentifier)) {
		body["targetIdentifier"] = request.TargetIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		body["targetType"] = request.TargetType
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdentifier)) {
		body["userIdentifier"] = request.UserIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		body["userType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProjectMember"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/updateMember"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProjectMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProtectedBranches(repositoryId *string, id *string, request *UpdateProtectedBranchesRequest) (_result *UpdateProtectedBranchesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProtectedBranchesResponse{}
	_body, _err := client.UpdateProtectedBranchesWithOptions(repositoryId, id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProtectedBranchesWithOptions(repositoryId *string, id *string, request *UpdateProtectedBranchesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProtectedBranchesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowMergeRoles)) {
		body["allowMergeRoles"] = request.AllowMergeRoles
	}

	if !tea.BoolValue(util.IsUnset(request.AllowMergeUserIds)) {
		body["allowMergeUserIds"] = request.AllowMergeUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.AllowPushRoles)) {
		body["allowPushRoles"] = request.AllowPushRoles
	}

	if !tea.BoolValue(util.IsUnset(request.AllowPushUserIds)) {
		body["allowPushUserIds"] = request.AllowPushUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.Branch)) {
		body["branch"] = request.Branch
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.MergeRequestSetting))) {
		body["mergeRequestSetting"] = request.MergeRequestSetting
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TestSettingDTO))) {
		body["testSettingDTO"] = request.TestSettingDTO
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProtectedBranches"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/protect_branches/" + tea.StringValue(openapiutil.GetEncodeParam(id))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProtectedBranchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRepository(repositoryId *string, request *UpdateRepositoryRequest) (_result *UpdateRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRepositoryResponse{}
	_body, _err := client.UpdateRepositoryWithOptions(repositoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRepositoryWithOptions(repositoryId *string, request *UpdateRepositoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdminSettingLanguage)) {
		body["adminSettingLanguage"] = request.AdminSettingLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.Avatar)) {
		body["avatar"] = request.Avatar
	}

	if !tea.BoolValue(util.IsUnset(request.BuildsEnabled)) {
		body["buildsEnabled"] = request.BuildsEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.CheckEmail)) {
		body["checkEmail"] = request.CheckEmail
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultBranch)) {
		body["defaultBranch"] = request.DefaultBranch
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.IssuesEnabled)) {
		body["issuesEnabled"] = request.IssuesEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.MergeRequestsEnabled)) {
		body["mergeRequestsEnabled"] = request.MergeRequestsEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OpenCloneDownloadControl)) {
		body["openCloneDownloadControl"] = request.OpenCloneDownloadControl
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectCloneDownloadMethodList)) {
		body["projectCloneDownloadMethodList"] = request.ProjectCloneDownloadMethodList
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectCloneDownloadRoleList)) {
		body["projectCloneDownloadRoleList"] = request.ProjectCloneDownloadRoleList
	}

	if !tea.BoolValue(util.IsUnset(request.SnippetsEnabled)) {
		body["snippetsEnabled"] = request.SnippetsEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.VisibilityLevel)) {
		body["visibilityLevel"] = request.VisibilityLevel
	}

	if !tea.BoolValue(util.IsUnset(request.WikiEnabled)) {
		body["wikiEnabled"] = request.WikiEnabled
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRepository"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRepositoryMember(repositoryId *string, userId *string, request *UpdateRepositoryMemberRequest) (_result *UpdateRepositoryMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRepositoryMemberResponse{}
	_body, _err := client.UpdateRepositoryMemberWithOptions(repositoryId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRepositoryMemberWithOptions(repositoryId *string, userId *string, request *UpdateRepositoryMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateRepositoryMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["accessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessLevel)) {
		body["accessLevel"] = request.AccessLevel
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireAt)) {
		body["expireAt"] = request.ExpireAt
	}

	if !tea.BoolValue(util.IsUnset(request.MemberType)) {
		body["memberType"] = request.MemberType
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedId)) {
		body["relatedId"] = request.RelatedId
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedInfos)) {
		body["relatedInfos"] = request.RelatedInfos
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRepositoryMember"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/repository/" + tea.StringValue(openapiutil.GetEncodeParam(repositoryId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(userId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRepositoryMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResourceMember(organizationId *string, resourceType *string, resourceId *string, accountId *string, request *UpdateResourceMemberRequest) (_result *UpdateResourceMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceMemberResponse{}
	_body, _err := client.UpdateResourceMemberWithOptions(organizationId, resourceType, resourceId, accountId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResourceMemberWithOptions(organizationId *string, resourceType *string, resourceId *string, accountId *string, request *UpdateResourceMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateResourceMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		body["roleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResourceMember"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(resourceType)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(resourceId)) + "/members/" + tea.StringValue(openapiutil.GetEncodeParam(accountId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVariableGroup(organizationId *string, id *string, request *UpdateVariableGroupRequest) (_result *UpdateVariableGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateVariableGroupResponse{}
	_body, _err := client.UpdateVariableGroupWithOptions(organizationId, id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVariableGroupWithOptions(organizationId *string, id *string, request *UpdateVariableGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateVariableGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Variables)) {
		body["variables"] = request.Variables
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateVariableGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/variableGroups/" + tea.StringValue(openapiutil.GetEncodeParam(id))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVariableGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWorkItem(organizationId *string, request *UpdateWorkItemRequest) (_result *UpdateWorkItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateWorkItemResponse{}
	_body, _err := client.UpdateWorkItemWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWorkItemWithOptions(organizationId *string, request *UpdateWorkItemRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateWorkItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FieldType)) {
		body["fieldType"] = request.FieldType
	}

	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		body["identifier"] = request.Identifier
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyKey)) {
		body["propertyKey"] = request.PropertyKey
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyValue)) {
		body["propertyValue"] = request.PropertyValue
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkItem"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/workitems/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWorkitemComment(organizationId *string, request *UpdateWorkitemCommentRequest) (_result *UpdateWorkitemCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateWorkitemCommentResponse{}
	_body, _err := client.UpdateWorkitemCommentWithOptions(organizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWorkitemCommentWithOptions(organizationId *string, request *UpdateWorkitemCommentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateWorkitemCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommentId)) {
		body["commentId"] = request.CommentId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.FormatType)) {
		body["formatType"] = request.FormatType
	}

	if !tea.BoolValue(util.IsUnset(request.WorkitemIdentifier)) {
		body["workitemIdentifier"] = request.WorkitemIdentifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkitemComment"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(openapiutil.GetEncodeParam(organizationId)) + "/workitems/commentUpdate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkitemCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
