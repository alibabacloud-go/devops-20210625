// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Id           *int64  `json:"id,omitempty" xml:"id,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 标签分类
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateHostGroupResponse) SetBody(v *CreateHostGroupResponseBody) *CreateHostGroupResponse {
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 项目信息
	Project *CreateProjectResponseBodyProject `json:"project,omitempty" xml:"project,omitempty" type:"Struct"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true或者false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// 空间大类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 空间大类id
	CategoryIdentifier *string `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	// 创建人id
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 自定义编号
	CustomCode *string `json:"customCode,omitempty" xml:"customCode,omitempty"`
	// 描述信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 创建时间
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 图标
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 大图标
	IconBig *string `json:"iconBig,omitempty" xml:"iconBig,omitempty"`
	// 图表组，json格式字符串
	IconGroup *string `json:"iconGroup,omitempty" xml:"iconGroup,omitempty"`
	// 小图标
	IconSmall *string `json:"iconSmall,omitempty" xml:"iconSmall,omitempty"`
	// 项目id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 项目唯一标识符
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// 空间路径
	IdentifierPath *string `json:"identifierPath,omitempty" xml:"identifierPath,omitempty"`
	// 项目状态
	LogicalStatus *string `json:"logicalStatus,omitempty" xml:"logicalStatus,omitempty"`
	// 修改人
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 项目名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 企业id
	OrganizationIdentifier *string `json:"organizationIdentifier,omitempty" xml:"organizationIdentifier,omitempty"`
	// 父项目id
	ParentIdentifier *string `json:"parentIdentifier,omitempty" xml:"parentIdentifier,omitempty"`
	// 可见范围
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// 状态id
	StatusIdentifier *string `json:"statusIdentifier,omitempty" xml:"statusIdentifier,omitempty"`
	// 状态阶段
	StatusStageIdentifier *string `json:"statusStageIdentifier,omitempty" xml:"statusStageIdentifier,omitempty"`
	// 空间小类
	SubType *string `json:"subType,omitempty" xml:"subType,omitempty"`
	// 空间小类id
	TypeIdentifier *string `json:"typeIdentifier,omitempty" xml:"typeIdentifier,omitempty"`
}

func (s CreateProjectResponseBodyProject) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyProject) SetCategory(v string) *CreateProjectResponseBodyProject {
	s.Category = &v
	return s
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

func (s *CreateProjectResponseBodyProject) SetIconBig(v string) *CreateProjectResponseBodyProject {
	s.IconBig = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetIconGroup(v string) *CreateProjectResponseBodyProject {
	s.IconGroup = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetIconSmall(v string) *CreateProjectResponseBodyProject {
	s.IconSmall = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetId(v string) *CreateProjectResponseBodyProject {
	s.Id = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetIdentifier(v string) *CreateProjectResponseBodyProject {
	s.Identifier = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetIdentifierPath(v string) *CreateProjectResponseBodyProject {
	s.IdentifierPath = &v
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

func (s *CreateProjectResponseBodyProject) SetParentIdentifier(v string) *CreateProjectResponseBodyProject {
	s.ParentIdentifier = &v
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

func (s *CreateProjectResponseBodyProject) SetSubType(v string) *CreateProjectResponseBodyProject {
	s.SubType = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetTypeIdentifier(v string) *CreateProjectResponseBodyProject {
	s.TypeIdentifier = &v
	return s
}

type CreateProjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type CreateResourceMemberRequest struct {
	// 用户id
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 角色部署组 deployGroup   user  成员，使用权限   admin 管理员，使用编辑权限 流水线 pipeline   admin 查看、运行、编辑权限   member  运行权限   viewer 查看权限
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateResourceMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateResourceMemberResponse) SetBody(v *CreateResourceMemberResponseBody) *CreateResourceMemberResponse {
	s.Body = v
	return s
}

type CreateSprintRequest struct {
	// 结束时间
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 迭代名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 项目id
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// 负责人列表
	StaffIds []*string `json:"staffIds,omitempty" xml:"staffIds,omitempty" type:"Repeated"`
	// 开始时间
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 迭代信息
	Sprint *CreateSprintResponseBodySprint `json:"sprint,omitempty" xml:"sprint,omitempty" type:"Struct"`
	// true或者false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// 创建人id
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 描述信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 结束时间
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 创建时间
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 迭代唯一标识符
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// 修改人
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 迭代名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 可见范围
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// 项目id
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// 开始时间
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSprintResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateSprintResponse) SetBody(v *CreateSprintResponseBody) *CreateSprintResponse {
	s.Body = v
	return s
}

type CreateSshKeyResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 企业公钥
	SshKey *CreateSshKeyResponseBodySshKey `json:"sshKey,omitempty" xml:"sshKey,omitempty" type:"Struct"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// 企业公钥id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 企业公钥
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSshKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateSshKeyResponse) SetBody(v *CreateSshKeyResponseBody) *CreateSshKeyResponse {
	s.Body = v
	return s
}

type CreateVariableGroupRequest struct {
	// 变量组描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 变量组名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 变量信息json字符串 isEncrypted 是否加密 name 变量名称 value 变量值
	Variables *string `json:"variables,omitempty" xml:"variables,omitempty"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 新建的变量组id
	VariableGroupId *int64 `json:"variableGroupId,omitempty" xml:"variableGroupId,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVariableGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateVariableGroupResponse) SetBody(v *CreateVariableGroupResponseBody) *CreateVariableGroupResponse {
	s.Body = v
	return s
}

type CreateWorkspaceRequest struct {
	// 代码来源URL（当前仅支持云效 Codeup 来源）
	CodeUrl *string `json:"codeUrl,omitempty" xml:"codeUrl,omitempty"`
	// 代码版本，支持 commitSHA、分支、标签
	CodeVersion *string `json:"codeVersion,omitempty" xml:"codeVersion,omitempty"`
	// 打开空间默认打开的文件相对路径
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// 工作空间名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 请求来源（用于统计，云产品集成时需要传入）
	RequestFrom *string `json:"requestFrom,omitempty" xml:"requestFrom,omitempty"`
	// 资源标识，提供给非标代码源作为空间复用的唯一标识
	ResourceIdentifier *string `json:"resourceIdentifier,omitempty" xml:"resourceIdentifier,omitempty"`
	// 工作空间复用标识，按照"用户+技术栈+代码地址+版本"进行复用 true - 复用 false - 不复用，每次均为新创建
	Reuse *bool `json:"reuse,omitempty" xml:"reuse,omitempty"`
	// 技术栈
	WorkspaceTemplate *string `json:"workspaceTemplate,omitempty" xml:"workspaceTemplate,omitempty"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求ID
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 请求是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 工作空间信息
	Workspace *CreateWorkspaceResponseBodyWorkspace `json:"workspace,omitempty" xml:"workspace,omitempty" type:"Struct"`
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
	// 创建时间戳
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建者，阿里云PK
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 工作空间唯一标识，字符串形式，可在云效DevStudio访问空间链接中获取
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 工作空间名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 空间状态，枚举：CREATING-创建中, SUCCESS-运行中, FROZEN-冻结中, RECOVERING-恢复中
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 工作空间模板
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateWorkspaceResponse) SetBody(v *CreateWorkspaceResponseBody) *CreateWorkspaceResponse {
	s.Body = v
	return s
}

type DeleteFlowTagResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlowTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteFlowTagResponse) SetBody(v *DeleteFlowTagResponseBody) *DeleteFlowTagResponse {
	s.Body = v
	return s
}

type DeleteFlowTagGroupResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlowTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteFlowTagGroupResponse) SetBody(v *DeleteFlowTagGroupResponseBody) *DeleteFlowTagGroupResponse {
	s.Body = v
	return s
}

type DeleteHostGroupResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteHostGroupResponse) SetBody(v *DeleteHostGroupResponseBody) *DeleteHostGroupResponse {
	s.Body = v
	return s
}

type DeletePipelineResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeletePipelineResponse) SetBody(v *DeletePipelineResponseBody) *DeletePipelineResponse {
	s.Body = v
	return s
}

type DeleteProjectRequest struct {
	// 项目id
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetProjectId(v string) *DeleteProjectRequest {
	s.ProjectId = &v
	return s
}

type DeleteProjectResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true或者false
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// true或者false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
	return s
}

type DeleteResourceMemberResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteResourceMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteResourceMemberResponse) SetBody(v *DeleteResourceMemberResponseBody) *DeleteResourceMemberResponse {
	s.Body = v
	return s
}

type DeleteVariableGroupResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVariableGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteVariableGroupResponse) SetBody(v *DeleteVariableGroupResponseBody) *DeleteVariableGroupResponse {
	s.Body = v
	return s
}

type FrozenWorkspaceResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求ID
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 请求是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FrozenWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *FrozenWorkspaceResponse) SetBody(v *FrozenWorkspaceResponseBody) *FrozenWorkspaceResponse {
	s.Body = v
	return s
}

type GetCustomFieldOptionRequest struct {
	// 项目id
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// 类型
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	// 工作项类型id
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 字段值信息
	Fileds []*GetCustomFieldOptionResponseBodyFileds `json:"fileds,omitempty" xml:"fileds,omitempty" type:"Repeated"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true或者false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// 展示的值
	DisplayValue *string `json:"displayValue,omitempty" xml:"displayValue,omitempty"`
	// 字段唯一标识
	FieldIdentifier *string `json:"fieldIdentifier,omitempty" xml:"fieldIdentifier,omitempty"`
	// 迭代唯一标识符
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// 展示级别，数字范围1~9，数字越大，颜色越浅
	Level *int64 `json:"level,omitempty" xml:"level,omitempty"`
	// 待选值顺序
	Position *int64 `json:"position,omitempty" xml:"position,omitempty"`
	// 字段中文名称
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 字段英文名称
	ValueEn *string `json:"valueEn,omitempty" xml:"valueEn,omitempty"`
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCustomFieldOptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetCustomFieldOptionResponse) SetBody(v *GetCustomFieldOptionResponseBody) *GetCustomFieldOptionResponse {
	s.Body = v
	return s
}

type GetFlowTagGroupResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string                                  `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	FlowTagGroup *GetFlowTagGroupResponseBodyFlowTagGroup `json:"flowTagGroup,omitempty" xml:"flowTagGroup,omitempty" type:"Struct"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFlowTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetFlowTagGroupResponse) SetBody(v *GetFlowTagGroupResponseBody) *GetFlowTagGroupResponse {
	s.Body = v
	return s
}

type GetHostGroupResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string                            `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HostGroup    *GetHostGroupResponseBodyHostGroup `json:"hostGroup,omitempty" xml:"hostGroup,omitempty" type:"Struct"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetHostGroupResponse) SetBody(v *GetHostGroupResponseBody) *GetHostGroupResponse {
	s.Body = v
	return s
}

type GetOrganizationMemberResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 成员
	Member *GetOrganizationMemberResponseBodyMember `json:"member,omitempty" xml:"member,omitempty" type:"Struct"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// 阿里云用户PK
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 生日
	Birthday *int64 `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 部门名称列表
	DeptLists []*string `json:"deptLists,omitempty" xml:"deptLists,omitempty" type:"Repeated"`
	// 邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// 入职时间
	HiredDate *int64 `json:"hiredDate,omitempty" xml:"hiredDate,omitempty"`
	// 第三方信息
	Identities *GetOrganizationMemberResponseBodyMemberIdentities `json:"identities,omitempty" xml:"identities,omitempty" type:"Struct"`
	// 加入云效企业时间
	JoinTime *int64 `json:"joinTime,omitempty" xml:"joinTime,omitempty"`
	// 最近一次访问时间
	LastVisitTime *int64 `json:"lastVisitTime,omitempty" xml:"lastVisitTime,omitempty"`
	// 手机号
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 企业成员名
	OrganizationMemberName *string `json:"organizationMemberName,omitempty" xml:"organizationMemberName,omitempty"`
	// 企业角色Id
	OrganizationRoleId *string `json:"organizationRoleId,omitempty" xml:"organizationRoleId,omitempty"`
	// 企业角色名字
	OrganizationRoleName *string `json:"organizationRoleName,omitempty" xml:"organizationRoleName,omitempty"`
	// 用户状态
	State *string `json:"state,omitempty" xml:"state,omitempty"`
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
	// 第三方系统的用户 id
	ExternUid *string `json:"externUid,omitempty" xml:"externUid,omitempty"`
	// 第三方系统
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOrganizationMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetOrganizationMemberResponse) SetBody(v *GetOrganizationMemberResponseBody) *GetOrganizationMemberResponse {
	s.Body = v
	return s
}

type GetPipelineResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 流水线
	Pipeline *GetPipelineResponseBodyPipeline `json:"pipeline,omitempty" xml:"pipeline,omitempty" type:"Struct"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// 创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建者阿里云账号id
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// 环境id 0 日常环境  1预发环境 2正式环境
	EnvId *int32 `json:"envId,omitempty" xml:"envId,omitempty"`
	// 环境名称
	EnvName *string `json:"envName,omitempty" xml:"envName,omitempty"`
	// 流水线分组id
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 更新人阿里云账号id
	ModifierAccountId *string `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	// 流水线名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 流水线配置
	PipelineConfig *GetPipelineResponseBodyPipelinePipelineConfig `json:"pipelineConfig,omitempty" xml:"pipelineConfig,omitempty" type:"Struct"`
	// 标签
	TagList []*GetPipelineResponseBodyPipelineTagList `json:"tagList,omitempty" xml:"tagList,omitempty" type:"Repeated"`
	// 更新时间
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
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
	// 流水线配置信息
	Flow *string `json:"flow,omitempty" xml:"flow,omitempty"`
	// 流水线环境变量等
	Settings *string `json:"settings,omitempty" xml:"settings,omitempty"`
	// 代码源
	Sources []*GetPipelineResponseBodyPipelinePipelineConfigSources `json:"sources,omitempty" xml:"sources,omitempty" type:"Repeated"`
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
	// 代码数据
	Data *GetPipelineResponseBodyPipelinePipelineConfigSourcesData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 代码源唯一标识
	Sign *string `json:"sign,omitempty" xml:"sign,omitempty"`
	// 代码源类型aliyunGit 阿里云代码库 customGitlab  自建git giteeGit 码云 codeup Codeup git 通用git gitlab gitlab bitbucket bitbucket githubOAuth github
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	// 分支
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// 克隆深度
	CloneDepth *int64 `json:"cloneDepth,omitempty" xml:"cloneDepth,omitempty"`
	// Credential Id
	CredentialId *int64 `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// Credential Label
	CredentialLabel *string `json:"credentialLabel,omitempty" xml:"credentialLabel,omitempty"`
	// Credential Type
	CredentialType *string `json:"credentialType,omitempty" xml:"credentialType,omitempty"`
	// 触发事件
	Events []*string `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	// 是否分支模式
	IsBranchMode *bool `json:"isBranchMode,omitempty" xml:"isBranchMode,omitempty"`
	// 是否设置clone深度
	IsCloneDepth *bool `json:"isCloneDepth,omitempty" xml:"isCloneDepth,omitempty"`
	// 是否子模块
	IsSubmodule *bool `json:"isSubmodule,omitempty" xml:"isSubmodule,omitempty"`
	// 是否提交触发
	IsTrigger *bool `json:"isTrigger,omitempty" xml:"isTrigger,omitempty"`
	// 代码源显示标签
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// github命名空间
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// 代码库地址
	Repo *string `json:"repo,omitempty" xml:"repo,omitempty"`
	// 服务连接Id
	ServiceConnectionId *int64 `json:"serviceConnectionId,omitempty" xml:"serviceConnectionId,omitempty"`
	// 触发过滤条件
	TriggerFilter *string `json:"triggerFilter,omitempty" xml:"triggerFilter,omitempty"`
	// webhhook地址
	Webhook *string `json:"webhook,omitempty" xml:"webhook,omitempty"`
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
	// 标签id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 标签名称
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	FileUrl      *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPipelineArtifactUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	FileUrl      *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPipelineEmasArtifactUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetPipelineEmasArtifactUrlResponse) SetBody(v *GetPipelineEmasArtifactUrlResponseBody) *GetPipelineEmasArtifactUrlResponse {
	s.Body = v
	return s
}

type GetPipelineRunResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 流水线运行实例
	PipelineRun *GetPipelineRunResponseBodyPipelineRun `json:"pipelineRun,omitempty" xml:"pipelineRun,omitempty" type:"Struct"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// 创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建者阿里云账号id
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// 更新人阿里云账号id
	ModifierAccountId *string `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	// 流水线Id
	PipelineId *int64 `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	// 流水线运行实例id
	PipelineRunId *int64 `json:"pipelineRunId,omitempty" xml:"pipelineRunId,omitempty"`
	// 代码源
	Sources []*GetPipelineRunResponseBodyPipelineRunSources `json:"sources,omitempty" xml:"sources,omitempty" type:"Repeated"`
	// 阶段拓扑信息
	StageGroup [][]*string `json:"stageGroup,omitempty" xml:"stageGroup,omitempty" type:"Repeated"`
	// 阶段信息
	Stages []*GetPipelineRunResponseBodyPipelineRunStages `json:"stages,omitempty" xml:"stages,omitempty" type:"Repeated"`
	// 状态 FAIL 运行失败 SUCCESS 运行成功 RUNNING 运行中
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 触发模式 1人工触发 2定时触发 3代码提交触发
	TriggerMode *int32 `json:"triggerMode,omitempty" xml:"triggerMode,omitempty"`
	// 更新时间
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
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
	// 代码源信息
	Data *GetPipelineRunResponseBodyPipelineRunSourcesData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 代码源唯一标识
	Sign *string `json:"sign,omitempty" xml:"sign,omitempty"`
	// 代码库类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	// 分支
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// 提交信息 json数据
	Commint *string `json:"commint,omitempty" xml:"commint,omitempty"`
	// 代码库地址
	Repo *string `json:"repo,omitempty" xml:"repo,omitempty"`
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
	// 阶段名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 阶段详情
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
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 任务
	Jobs []*GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	// 阶段名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
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
	// 后续操作
	Actions []*GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 任务Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 任务名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 触发参数
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
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
	// 是否可用
	Disable *bool `json:"disable,omitempty" xml:"disable,omitempty"`
	// API参数
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// API名称
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	ReportUrl    *string `json:"reportUrl,omitempty" xml:"reportUrl,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPipelineScanReportUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetPipelineScanReportUrlResponse) SetBody(v *GetPipelineScanReportUrlResponseBody) *GetPipelineScanReportUrlResponse {
	s.Body = v
	return s
}

type GetProjectInfoResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 项目信息
	Project   *GetProjectInfoResponseBodyProject `json:"project,omitempty" xml:"project,omitempty" type:"Struct"`
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                              `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProjectInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetProjectInfoResponse) SetBody(v *GetProjectInfoResponseBody) *GetProjectInfoResponse {
	s.Body = v
	return s
}

type GetVMDeployOrderResponseBody struct {
	// 部署单
	DeployOrder *GetVMDeployOrderResponseBodyDeployOrder `json:"deployOrder,omitempty" xml:"deployOrder,omitempty" type:"Struct"`
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// 后续action
	Actions []*GetVMDeployOrderResponseBodyDeployOrderActions `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// 创建时时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建人
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 当前发布批次
	CurrentBatch *int32 `json:"currentBatch,omitempty" xml:"currentBatch,omitempty"`
	// 部署机器信息
	DeployMachineInfo *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo `json:"deployMachineInfo,omitempty" xml:"deployMachineInfo,omitempty" type:"Struct"`
	// 部署单ID
	DeployOrderId *string `json:"deployOrderId,omitempty" xml:"deployOrderId,omitempty"`
	// 错误码
	ExceptionCode *string `json:"exceptionCode,omitempty" xml:"exceptionCode,omitempty"`
	// 发布状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 总发布批次
	TotalBatch *int32 `json:"totalBatch,omitempty" xml:"totalBatch,omitempty"`
	// 修改时间
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
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
	// 是否可用
	Disable *bool `json:"disable,omitempty" xml:"disable,omitempty"`
	// 参数
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// Action
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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

func (s *GetVMDeployOrderResponseBodyDeployOrderActions) SetParams(v map[string]interface{}) *GetVMDeployOrderResponseBodyDeployOrderActions {
	s.Params = v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderActions) SetType(v string) *GetVMDeployOrderResponseBodyDeployOrderActions {
	s.Type = &v
	return s
}

type GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfo struct {
	// 发布批次
	BatchNum *int32 `json:"batchNum,omitempty" xml:"batchNum,omitempty"`
	// 部署机器列表
	DeployMachines []*GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachines `json:"deployMachines,omitempty" xml:"deployMachines,omitempty" type:"Repeated"`
	// 主机组ID
	HostGroupId *int64 `json:"hostGroupId,omitempty" xml:"hostGroupId,omitempty"`
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
	// 后续action
	Actions []*GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// 部署批次
	BatchNum *int32 `json:"batchNum,omitempty" xml:"batchNum,omitempty"`
	// 机器状态
	ClientStatus *string `json:"clientStatus,omitempty" xml:"clientStatus,omitempty"`
	// 开始时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 机器IP
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
	// 机器sn
	MachineSn *string `json:"machineSn,omitempty" xml:"machineSn,omitempty"`
	// 部署状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 修改时间
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
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
	// 是否可用
	Disable *bool `json:"disable,omitempty" xml:"disable,omitempty"`
	// 参数
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// Action
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) SetParams(v map[string]interface{}) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions {
	s.Params = v
	return s
}

func (s *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions) SetType(v string) *GetVMDeployOrderResponseBodyDeployOrderDeployMachineInfoDeployMachinesActions {
	s.Type = &v
	return s
}

type GetVMDeployOrderResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVMDeployOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetVMDeployOrderResponse) SetBody(v *GetVMDeployOrderResponseBody) *GetVMDeployOrderResponse {
	s.Body = v
	return s
}

type GetVariableGroupResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 变量组
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
	// 创建人阿里云账号id
	CcreatorAccountId *string `json:"ccreatorAccountId,omitempty" xml:"ccreatorAccountId,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 变量组描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 变量组id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 更新人阿里云账号id
	ModifierAccountId *string `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	// 变量组名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 关联的流水线
	RelatedPipelines []*GetVariableGroupResponseBodyVariableGroupRelatedPipelines `json:"relatedPipelines,omitempty" xml:"relatedPipelines,omitempty" type:"Repeated"`
	// 更新时间
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// 变量
	Variables []*GetVariableGroupResponseBodyVariableGroupVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
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
	// 关联的流水线Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 关联的流水线名称
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
	// 是否加密
	IsEncrypted *bool `json:"isEncrypted,omitempty" xml:"isEncrypted,omitempty"`
	// 变量名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 变量值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVariableGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetVariableGroupResponse) SetBody(v *GetVariableGroupResponseBody) *GetVariableGroupResponse {
	s.Body = v
	return s
}

type GetWorkItemActivityResponseBody struct {
	// 动态信息
	Activities []*GetWorkItemActivityResponseBodyActivities `json:"activities,omitempty" xml:"activities,omitempty" type:"Repeated"`
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true或者false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// 动作类型
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 事件id
	EventId *int64 `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// 事件时间
	EventTime *int64 `json:"eventTime,omitempty" xml:"eventTime,omitempty"`
	// 事件类型
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// 操作者
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 父事件id
	ParentEventId *int64 `json:"parentEventId,omitempty" xml:"parentEventId,omitempty"`
	// 修改属性
	Property *GetWorkItemActivityResponseBodyActivitiesProperty `json:"property,omitempty" xml:"property,omitempty" type:"Struct"`
	// 操作对象
	ResourceIdentifier *string `json:"resourceIdentifier,omitempty" xml:"resourceIdentifier,omitempty"`
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

type GetWorkItemActivityResponseBodyActivitiesProperty struct {
	// 属性的展示名
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 资源id
	PropertyIdentifier *string `json:"propertyIdentifier,omitempty" xml:"propertyIdentifier,omitempty"`
	// 属性key
	PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty"`
	// 类型
	PropertyType *string `json:"propertyType,omitempty" xml:"propertyType,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWorkItemActivityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetWorkItemActivityResponse) SetBody(v *GetWorkItemActivityResponseBody) *GetWorkItemActivityResponse {
	s.Body = v
	return s
}

type GetWorkItemInfoResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true或者false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 工作项信息
	Workitem *GetWorkItemInfoResponseBodyWorkitem `json:"workitem,omitempty" xml:"workitem,omitempty" type:"Struct"`
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
	// 负责人
	AssignedTo *string `json:"assignedTo,omitempty" xml:"assignedTo,omitempty"`
	// 工作项的类型id
	CategoryIdentifier *string `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	// 创建人
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 工作项内容
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// 创建时间
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 工作项唯一标识
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// 逻辑状态
	LogicalStatus *string `json:"logicalStatus,omitempty" xml:"logicalStatus,omitempty"`
	// 修改人
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 父工作项id
	ParentIdentifier *string `json:"parentIdentifier,omitempty" xml:"parentIdentifier,omitempty"`
	// 编号
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	// 所属项目id
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// 所属项目名称
	SpaceName *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	// 项目类型
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	// 状态名称
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 状态阶段id
	StatusStageIdentifier *string `json:"statusStageIdentifier,omitempty" xml:"statusStageIdentifier,omitempty"`
	// 工作项标题
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 状态更新时间
	UpdateStatusAt *int64 `json:"updateStatusAt,omitempty" xml:"updateStatusAt,omitempty"`
	// 工作项类型id
	WorkitemTypeIdentifier *string `json:"workitemTypeIdentifier,omitempty" xml:"workitemTypeIdentifier,omitempty"`
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

func (s *GetWorkItemInfoResponseBodyWorkitem) SetStatus(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.Status = &v
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

func (s *GetWorkItemInfoResponseBodyWorkitem) SetUpdateStatusAt(v int64) *GetWorkItemInfoResponseBodyWorkitem {
	s.UpdateStatusAt = &v
	return s
}

func (s *GetWorkItemInfoResponseBodyWorkitem) SetWorkitemTypeIdentifier(v string) *GetWorkItemInfoResponseBodyWorkitem {
	s.WorkitemTypeIdentifier = &v
	return s
}

type GetWorkItemInfoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWorkItemInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetWorkItemInfoResponse) SetBody(v *GetWorkItemInfoResponseBody) *GetWorkItemInfoResponse {
	s.Body = v
	return s
}

type GetWorkItemWorkFlowInfoRequest struct {
	// 项目id
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true或者false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 工作项信息
	Workflow *GetWorkItemWorkFlowInfoResponseBodyWorkflow `json:"workflow,omitempty" xml:"workflow,omitempty" type:"Struct"`
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
	// 创建人
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 工作流的默认状态
	DefaultStatusIdentifier *string `json:"defaultStatusIdentifier,omitempty" xml:"defaultStatusIdentifier,omitempty"`
	// 工作流的描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 创建时间
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 工作流唯一标识
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// 修改人
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 工作流名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 工作流所属的团队空间或项目的identifier
	OwnerSpaceIdentifier *string `json:"ownerSpaceIdentifier,omitempty" xml:"ownerSpaceIdentifier,omitempty"`
	// 工作流所属的团队项目类型
	OwnerSpaceType *string `json:"ownerSpaceType,omitempty" xml:"ownerSpaceType,omitempty"`
	// 资源类型
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// 工作流来源
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 工作流的状态顺序
	StatusOrder *string `json:"statusOrder,omitempty" xml:"statusOrder,omitempty"`
	// 状态列表
	Statuses []*GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses `json:"statuses,omitempty" xml:"statuses,omitempty" type:"Repeated"`
	// 工作流的流转步骤
	WorkflowActions []*GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions `json:"workflowActions,omitempty" xml:"workflowActions,omitempty" type:"Repeated"`
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
	// 创建人
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 描述信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 创建时间
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 状态唯一标识
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// 修改人
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 状态名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 资源来源
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// 状态来源
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 阶段信息-阶段的唯一标识
	WorkflowStageIdentifier *string `json:"workflowStageIdentifier,omitempty" xml:"workflowStageIdentifier,omitempty"`
	// 阶段信息-名称
	WorkflowStageName *string `json:"workflowStageName,omitempty" xml:"workflowStageName,omitempty"`
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
	// 流转步骤的id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// action的名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// action对应的下个状态的信息id
	NextWorkflowStatusIdentifier *string `json:"nextWorkflowStatusIdentifier,omitempty" xml:"nextWorkflowStatusIdentifier,omitempty"`
	// action对应的工作流
	WorkflowIdentifier *string `json:"workflowIdentifier,omitempty" xml:"workflowIdentifier,omitempty"`
	// action对应的当前状态id
	WorkflowStatusIdentifier *string `json:"workflowStatusIdentifier,omitempty" xml:"workflowStatusIdentifier,omitempty"`
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWorkItemWorkFlowInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetWorkItemWorkFlowInfoResponse) SetBody(v *GetWorkItemWorkFlowInfoResponseBody) *GetWorkItemWorkFlowInfoResponse {
	s.Body = v
	return s
}

type GetWorkspaceResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求ID
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 请求是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 工作空间信息
	Workspace *GetWorkspaceResponseBodyWorkspace `json:"workspace,omitempty" xml:"workspace,omitempty" type:"Struct"`
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
	// 代码来源URL
	CodeUrl *string `json:"codeUrl,omitempty" xml:"codeUrl,omitempty"`
	// 代码版本，支持 commitSHA、分支、标签
	CodeVersion *string `json:"codeVersion,omitempty" xml:"codeVersion,omitempty"`
	// 创建时间戳
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 工作空间唯一标识，字符串形式，可在云效DevStudio访问空间链接中获取
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 工作空间名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 机器规格
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
	// 空间状态，枚举：CREATING-创建中, SUCCESS-运行中, FROZEN-冻结中, RECOVERING-恢复中
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 工作空间模板
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// 用户阿里云PK
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetWorkspaceResponse) SetBody(v *GetWorkspaceResponseBody) *GetWorkspaceResponse {
	s.Body = v
	return s
}

type ListFlowTagGroupsResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 标签分类
	FlowTagGroups []*ListFlowTagGroupsResponseBodyFlowTagGroups `json:"flowTagGroups,omitempty" xml:"flowTagGroups,omitempty" type:"Repeated"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// 创建人
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// 标签分类id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 修改人
	ModiferAccountId *string `json:"modiferAccountId,omitempty" xml:"modiferAccountId,omitempty"`
	// 标签分类名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowTagGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListFlowTagGroupsResponse) SetBody(v *ListFlowTagGroupsResponseBody) *ListFlowTagGroupsResponse {
	s.Body = v
	return s
}

type ListHostGroupsRequest struct {
	// 主机组结束时间
	CreateEndTime *int64 `json:"createEndTime,omitempty" xml:"createEndTime,omitempty"`
	// 主机组创建时间
	CreateStartTime *int64 `json:"createStartTime,omitempty" xml:"createStartTime,omitempty"`
	// 创建阿里云账号id，多个逗号分割
	CreatorAccountIds *string `json:"creatorAccountIds,omitempty" xml:"creatorAccountIds,omitempty"`
	// 主机组id，多个逗号分割
	Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
	// 结果返回个数
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 主机组名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 分页token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 排序顺序
	PageOrder *string `json:"pageOrder,omitempty" xml:"pageOrder,omitempty"`
	// 排序条件ID
	PageSort *string `json:"pageSort,omitempty" xml:"pageSort,omitempty"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 主机组
	HostGroups []*ListHostGroupsResponseBodyHostGroups `json:"hostGroups,omitempty" xml:"hostGroups,omitempty" type:"Repeated"`
	// 分页token,空表示最后一页
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 总数
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	// 阿里云区域
	AliyunRegion *string `json:"aliyunRegion,omitempty" xml:"aliyunRegion,omitempty"`
	// 主机时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建人阿里云账号id
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// ecs标签Key
	EcsLabelKey *string `json:"ecsLabelKey,omitempty" xml:"ecsLabelKey,omitempty"`
	// Ecs标签值
	EcsLabelValue *string `json:"ecsLabelValue,omitempty" xml:"ecsLabelValue,omitempty"`
	// 主机类型
	EcsType *string `json:"ecsType,omitempty" xml:"ecsType,omitempty"`
	// 主机个数
	HostNum *int64 `json:"hostNum,omitempty" xml:"hostNum,omitempty"`
	// 323232
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 修改人阿里云账号id
	ModifierAccountId *string `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	// 部署组名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 服务连接Id
	ServiceConnectionId *int64 `json:"serviceConnectionId,omitempty" xml:"serviceConnectionId,omitempty"`
	// 类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 更新时间
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHostGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 成员列表
	Members []*ListOrganizationMembersResponseBodyMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 分页Token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 总数
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	// 阿里云用户ID
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 生日
	Birthday *int64 `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 部门名称列表
	DeptLists []*string `json:"deptLists,omitempty" xml:"deptLists,omitempty" type:"Repeated"`
	// 邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// 入职时间
	HiredDate *int64 `json:"hiredDate,omitempty" xml:"hiredDate,omitempty"`
	// 第三方信息
	Identities *ListOrganizationMembersResponseBodyMembersIdentities `json:"identities,omitempty" xml:"identities,omitempty" type:"Struct"`
	// 加入云效企业时间
	JoinTime *int64 `json:"joinTime,omitempty" xml:"joinTime,omitempty"`
	// 最近一次访问时间
	LastVisitTime *int64 `json:"lastVisitTime,omitempty" xml:"lastVisitTime,omitempty"`
	// 手机号
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 企业成员名
	OrganizationMemberName *string `json:"organizationMemberName,omitempty" xml:"organizationMemberName,omitempty"`
	// 企业角色Id
	OrganizationRoleId *string `json:"organizationRoleId,omitempty" xml:"organizationRoleId,omitempty"`
	// 企业角色名字
	OrganizationRoleName *string `json:"organizationRoleName,omitempty" xml:"organizationRoleName,omitempty"`
	// 用户状态
	State *string `json:"state,omitempty" xml:"state,omitempty"`
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
	// 第三方系统的用户Id
	ExternUid *string `json:"externUid,omitempty" xml:"externUid,omitempty"`
	// 第三方系统
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOrganizationMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListOrganizationMembersResponse) SetBody(v *ListOrganizationMembersResponseBody) *ListOrganizationMembersResponse {
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string                                    `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Jobs         []*ListPipelineJobHistorysResponseBodyJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	NextToken    *string                                    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success    *bool  `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPipelineJobHistorysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string                             `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Jobs         []*ListPipelineJobsResponseBodyJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPipelineJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListPipelineJobsResponse) SetBody(v *ListPipelineJobsResponseBody) *ListPipelineJobsResponse {
	s.Body = v
	return s
}

type ListPipelineRunsRequest struct {
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 最大返回数量
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页Token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 状态 状态 FAIL 运行失败 SUCCESS 运行成功 RUNNING 运行中
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 触发模式 1人工触发 2定时触发 3代码提交触发
	TriggerMode *int32 `json:"triggerMode,omitempty" xml:"triggerMode,omitempty"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 下一个分页token，为空时，表示没有下一页
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 流水线运行实例
	PipelineRuns []*ListPipelineRunsResponseBodyPipelineRuns `json:"pipelineRuns,omitempty" xml:"pipelineRuns,omitempty" type:"Repeated"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 总数
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	// 运行人阿里云账号id
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 流水线id
	PipelineId *int64 `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	// 流水线实例id
	PipelineRunId *int64 `json:"pipelineRunId,omitempty" xml:"pipelineRunId,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 运行状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 触发模式
	TriggerMode *int64 `json:"triggerMode,omitempty" xml:"triggerMode,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPipelineRunsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListPipelineRunsResponse) SetBody(v *ListPipelineRunsResponseBody) *ListPipelineRunsResponse {
	s.Body = v
	return s
}

type ListPipelinesRequest struct {
	// 创建结束时间
	CreateEndTime *int64 `json:"createEndTime,omitempty" xml:"createEndTime,omitempty"`
	// 创建开始时间
	CreateStartTime *int64 `json:"createStartTime,omitempty" xml:"createStartTime,omitempty"`
	// 创建人阿里云账号Id
	CreatorAccountIds *string `json:"creatorAccountIds,omitempty" xml:"creatorAccountIds,omitempty"`
	// 执行人阿里云账号id
	ExecuteAccountIds *string `json:"executeAccountIds,omitempty" xml:"executeAccountIds,omitempty"`
	// 执行结束时间
	ExecuteEndTime *int64 `json:"executeEndTime,omitempty" xml:"executeEndTime,omitempty"`
	// 执行开始时间
	ExecuteStartTime *int64 `json:"executeStartTime,omitempty" xml:"executeStartTime,omitempty"`
	// 返回的总数
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页Token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 流水线名称
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// 状态列表，多个逗号分割
	StatusList *string `json:"statusList,omitempty" xml:"statusList,omitempty"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 分页Token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 流水线
	Pipelines []*ListPipelinesResponseBodyPipelines `json:"pipelines,omitempty" xml:"pipelines,omitempty" type:"Repeated"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 总数
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	// 创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建人阿里云账号id
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// 流水线id
	PipelineId *int64 `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	// 流水线名称
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
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

func (s *ListPipelinesResponseBodyPipelines) SetPipelineId(v int64) *ListPipelinesResponseBodyPipelines {
	s.PipelineId = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetPipelineName(v string) *ListPipelinesResponseBodyPipelines {
	s.PipelineName = &v
	return s
}

type ListPipelinesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// member信息
	Members []*ListProjectMembersResponseBodyMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true或者false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// 登陆账号
	Account *string `json:"account,omitempty" xml:"account,omitempty"`
	// 用户头像
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 钉钉id
	DingTalkId *string `json:"dingTalkId,omitempty" xml:"dingTalkId,omitempty"`
	// 展示名
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 展示昵称
	DisplayNickName *string `json:"displayNickName,omitempty" xml:"displayNickName,omitempty"`
	// 展示真名
	DisplayRealName *string `json:"displayRealName,omitempty" xml:"displayRealName,omitempty"`
	// 部门信息
	Division *ListProjectMembersResponseBodyMembersDivision `json:"division,omitempty" xml:"division,omitempty" type:"Struct"`
	// 邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// 性别
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// 用户唯一 标识符
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// 手机号
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 英文名
	NameEn *string `json:"nameEn,omitempty" xml:"nameEn,omitempty"`
	// 昵称
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// 昵称拼音
	NickNamePinyin *string `json:"nickNamePinyin,omitempty" xml:"nickNamePinyin,omitempty"`
	// 企业信息
	OrganizationUserInfo *ListProjectMembersResponseBodyMembersOrganizationUserInfo `json:"organizationUserInfo,omitempty" xml:"organizationUserInfo,omitempty" type:"Struct"`
	// 真名
	RealName *string `json:"realName,omitempty" xml:"realName,omitempty"`
	// 真名拼音
	RealNamePinyin *string `json:"realNamePinyin,omitempty" xml:"realNamePinyin,omitempty"`
	// 用户类型
	Stamp *string `json:"stamp,omitempty" xml:"stamp,omitempty"`
	// 角色id
	TbRoleId *string `json:"tbRoleId,omitempty" xml:"tbRoleId,omitempty"`
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
	// 部门唯一标识
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
	// 企业唯一标识符
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProjectMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListProjectMembersResponse) SetBody(v *ListProjectMembersResponseBody) *ListProjectMembersResponse {
	s.Body = v
	return s
}

type ListResourceMembersResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 成员
	ResourceMembers []*ListResourceMembersResponseBodyResourceMembers `json:"resourceMembers,omitempty" xml:"resourceMembers,omitempty" type:"Repeated"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// 账号id
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 角色
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// 用户名称
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListResourceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListResourceMembersResponse) SetBody(v *ListResourceMembersResponseBody) *ListResourceMembersResponse {
	s.Body = v
	return s
}

type ListServiceConnectionsRequest struct {
	// aliyun_code  阿里云代码 Codeup       Codeup  Gitee        码云 github       Github ack       容器服务Kubernetes(ACK) docker_register_aliyun    容器镜像服务(ACR) ecs          对象存储(OSS) edas          企业级分布式应用(EDAS) emas         移动研发平台(EMAS) fc            阿里云函数计算(FC) kubernetes     自建k8s集群 oss            对象存储(OSS) PACKAGES       制品仓库 ros   资源编排服务(ROS) sae       Serverless应用引擎(SAE)
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 服务连接
	ServiceConnections []*ListServiceConnectionsResponseBodyServiceConnections `json:"serviceConnections,omitempty" xml:"serviceConnections,omitempty" type:"Repeated"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// 创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 服务连接Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 服务连接名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 拥有者阿里云账号id
	OwnerAccountId *int64 `json:"ownerAccountId,omitempty" xml:"ownerAccountId,omitempty"`
	// 服务连接类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListServiceConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListServiceConnectionsResponse) SetBody(v *ListServiceConnectionsResponseBody) *ListServiceConnectionsResponse {
	s.Body = v
	return s
}

type ListVariableGroupsRequest struct {
	// 最大返回数，默认30
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页token，上一次请求的出参nextToken
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 排序顺序
	PageOrder *string `json:"pageOrder,omitempty" xml:"pageOrder,omitempty"`
	// 排序条件
	PageSort *string `json:"pageSort,omitempty" xml:"pageSort,omitempty"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 下一次查询的token，为空表示最后一页
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 变量组总数
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 变量组
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
	// 创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建人阿里云账号id
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// 变量组描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 变量组id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 更新人阿里云账号id
	ModifierAccountId *string `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	// 变量组名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 关联的流水线
	RelatedPipelines []*ListVariableGroupsResponseBodyVariableGroupsRelatedPipelines `json:"relatedPipelines,omitempty" xml:"relatedPipelines,omitempty" type:"Repeated"`
	// 更新时间
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// 变量
	Variables []*ListVariableGroupsResponseBodyVariableGroupsVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
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
	// 关联的流水线Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 关联的流水线名称
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
	// 是否加密
	IsEncrypted *bool `json:"isEncrypted,omitempty" xml:"isEncrypted,omitempty"`
	// 变量名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 变量值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVariableGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListVariableGroupsResponse) SetBody(v *ListVariableGroupsResponseBody) *ListVariableGroupsResponse {
	s.Body = v
	return s
}

type ListWorkspacesRequest struct {
	// 本次读取的最大数据记录数量，默认10，最大100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 用来标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 枚举值：CREATING-创建中, SUCCESS-运行中, FROZEN-冻结中, RECOVERING-恢复中
	StatusList []*string `json:"statusList,omitempty" xml:"statusList,omitempty" type:"Repeated"`
	// 空间模板列表
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
	// 本次读取的最大数据记录数量，默认10，最大100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 用来标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 枚举值：CREATING-创建中, SUCCESS-运行中, FROZEN-冻结中, RECOVERING-恢复中
	StatusListShrink *string `json:"statusList,omitempty" xml:"statusList,omitempty"`
	// 空间模板列表
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// MaxResults本次请求所返回的最大记录条数
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 请求ID
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 请求是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// TotalCount本次请求条件下的数据总量，此参数为可选参数，默认可不返回
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 工作空间列表
	Workspaces []*ListWorkspacesResponseBodyWorkspaces `json:"workspaces,omitempty" xml:"workspaces,omitempty" type:"Repeated"`
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
	// 代码来源URL
	CodeUrl *string `json:"codeUrl,omitempty" xml:"codeUrl,omitempty"`
	// 代码版本，支持 commitSHA、分支、标签
	CodeVersion *string `json:"codeVersion,omitempty" xml:"codeVersion,omitempty"`
	// 创建时间戳
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 工作空间唯一标识，字符串形式，可在云效DevStudio访问空间链接中获取
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 工作空间名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 机器规格
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
	// 空间状态，枚举：CREATING-创建中, SUCCESS-运行中, FROZEN-冻结中, RECOVERING-恢复中
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 工作空间模板
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// 用户阿里云PK
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LogPipelineJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *LogPipelineJobRunResponse) SetBody(v *LogPipelineJobRunResponseBody) *LogPipelineJobRunResponse {
	s.Body = v
	return s
}

type LogVMDeployMachineResponseBody struct {
	// 部署单
	DeployMachineLog *LogVMDeployMachineResponseBodyDeployMachineLog `json:"deployMachineLog,omitempty" xml:"deployMachineLog,omitempty" type:"Struct"`
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// 部署地域
	AliyunRegion *string `json:"aliyunRegion,omitempty" xml:"aliyunRegion,omitempty"`
	// 部署开始时间
	DeployBeginTime *int64 `json:"deployBeginTime,omitempty" xml:"deployBeginTime,omitempty"`
	// 部署结束时间
	DeployEndTime *int64 `json:"deployEndTime,omitempty" xml:"deployEndTime,omitempty"`
	// 部署日志
	DeployLog *string `json:"deployLog,omitempty" xml:"deployLog,omitempty"`
	// 部署日志路径
	DeployLogPath *string `json:"deployLogPath,omitempty" xml:"deployLogPath,omitempty"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LogVMDeployMachineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *LogVMDeployMachineResponse) SetBody(v *LogVMDeployMachineResponseBody) *LogVMDeployMachineResponse {
	s.Body = v
	return s
}

type PassPipelineValidateResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PassPipelineValidateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PassPipelineValidateResponse) SetBody(v *PassPipelineValidateResponseBody) *PassPipelineValidateResponse {
	s.Body = v
	return s
}

type RefusePipelineValidateResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefusePipelineValidateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RefusePipelineValidateResponse) SetBody(v *RefusePipelineValidateResponseBody) *RefusePipelineValidateResponse {
	s.Body = v
	return s
}

type ReleaseWorkspaceResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求ID
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 请求是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ReleaseWorkspaceResponse) SetBody(v *ReleaseWorkspaceResponseBody) *ReleaseWorkspaceResponse {
	s.Body = v
	return s
}

type ResetSshKeyResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 企业公钥
	SshKey *ResetSshKeyResponseBodySshKey `json:"sshKey,omitempty" xml:"sshKey,omitempty" type:"Struct"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// 企业公钥id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 企业公钥
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetSshKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ResetSshKeyResponse) SetBody(v *ResetSshKeyResponseBody) *ResetSshKeyResponse {
	s.Body = v
	return s
}

type ResumeVMDeployOrderResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResumeVMDeployOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RetryPipelineJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RetryPipelineJobRunResponse) SetBody(v *RetryPipelineJobRunResponseBody) *RetryPipelineJobRunResponse {
	s.Body = v
	return s
}

type RetryVMDeployMachineResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RetryVMDeployMachineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SkipPipelineJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SkipPipelineJobRunResponse) SetBody(v *SkipPipelineJobRunResponseBody) *SkipPipelineJobRunResponse {
	s.Body = v
	return s
}

type SkipVMDeployMachineResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SkipVMDeployMachineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SkipVMDeployMachineResponse) SetBody(v *SkipVMDeployMachineResponseBody) *SkipVMDeployMachineResponse {
	s.Body = v
	return s
}

type StartPipelineRunRequest struct {
	// 流水线运行参数,json字符串 branchModeBranchs  分支模式运行的分支 envs  环境变量 runningBranchs 运行分支 runningTags  运行代码tag comment  运行备注
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 流水线运行实例id
	PipelineRunId *int64 `json:"pipelineRunId,omitempty" xml:"pipelineRunId,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StartPipelineRunResponse) SetBody(v *StartPipelineRunResponseBody) *StartPipelineRunResponse {
	s.Body = v
	return s
}

type StopPipelineJobRunResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopPipelineJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StopPipelineJobRunResponse) SetBody(v *StopPipelineJobRunResponseBody) *StopPipelineJobRunResponse {
	s.Body = v
	return s
}

type StopPipelineRunResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StopPipelineRunResponse) SetBody(v *StopPipelineRunResponseBody) *StopPipelineRunResponse {
	s.Body = v
	return s
}

type StopVMDeployOrderResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopVMDeployOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StopVMDeployOrderResponse) SetBody(v *StopVMDeployOrderResponseBody) *StopVMDeployOrderResponse {
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFlowTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFlowTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePipelineBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdatePipelineBaseInfoResponse) SetBody(v *UpdatePipelineBaseInfoResponseBody) *UpdatePipelineBaseInfoResponse {
	s.Body = v
	return s
}

type UpdateProjectMemberRequest struct {
	// 角色id
	RoleIdentifier *string `json:"roleIdentifier,omitempty" xml:"roleIdentifier,omitempty"`
	// 资源id，也就是项目id
	TargetIdentifier *string `json:"targetIdentifier,omitempty" xml:"targetIdentifier,omitempty"`
	// 资源类型
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// 用户id
	UserIdentifier *string `json:"userIdentifier,omitempty" xml:"userIdentifier,omitempty"`
	// 用户类型
	UserType *string `json:"userType,omitempty" xml:"userType,omitempty"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 成员信息
	Member *UpdateProjectMemberResponseBodyMember `json:"member,omitempty" xml:"member,omitempty" type:"Struct"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true或者false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	// 创建时间
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 角色id
	RoleIdentifier *string `json:"roleIdentifier,omitempty" xml:"roleIdentifier,omitempty"`
	// 资源id，也就是项目id
	TargetIdentifier *string `json:"targetIdentifier,omitempty" xml:"targetIdentifier,omitempty"`
	// 资源类型
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// 用户id
	UserIdentifier *string `json:"userIdentifier,omitempty" xml:"userIdentifier,omitempty"`
	// 用户类型
	UserType *string `json:"userType,omitempty" xml:"userType,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateProjectMemberResponse) SetBody(v *UpdateProjectMemberResponseBody) *UpdateProjectMemberResponse {
	s.Body = v
	return s
}

type UpdateResourceMemberRequest struct {
	// 角色部署组 deployGroup   user  成员，使用权限   admin 管理员，使用编辑权限   owner 拥有者，所有权限 流水线 pipeline   owner 拥有者，所有权限   admin 查看、运行、编辑权限   member  运行权限   viewer 查看权限
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateResourceMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateResourceMemberResponse) SetBody(v *UpdateResourceMemberResponseBody) *UpdateResourceMemberResponse {
	s.Body = v
	return s
}

type UpdateVariableGroupRequest struct {
	// 变量组描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 变量组名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 变量信息json字符串 isEncrypted 是否加密 name 变量名称 value 变量值
	Variables *string `json:"variables,omitempty" xml:"variables,omitempty"`
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateVariableGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateVariableGroupResponse) SetBody(v *UpdateVariableGroupResponseBody) *UpdateVariableGroupResponse {
	s.Body = v
	return s
}

type UpdateWorkItemRequest struct {
	// 工作项id
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// 更新的字段名
	PropertyKey *string `json:"propertyKey,omitempty" xml:"propertyKey,omitempty"`
	// 更新后的值
	PropertyValue *string `json:"propertyValue,omitempty" xml:"propertyValue,omitempty"`
}

func (s UpdateWorkItemRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkItemRequest) GoString() string {
	return s.String()
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
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 请求id，每次请求都是唯一值，便于后续排查问题
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true或者false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 工作项信息
	Workitem *UpdateWorkItemResponseBodyWorkitem `json:"workitem,omitempty" xml:"workitem,omitempty" type:"Struct"`
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
	// 负责人
	AssignedTo *string `json:"assignedTo,omitempty" xml:"assignedTo,omitempty"`
	// 工作项的类型id
	CategoryIdentifier *string `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	// 创建人
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 工作项内容
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// 创建时间
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 工作项唯一标识
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// 逻辑状态
	LogicalStatus *string `json:"logicalStatus,omitempty" xml:"logicalStatus,omitempty"`
	// 修改人
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 父工作项id
	ParentIdentifier *string `json:"parentIdentifier,omitempty" xml:"parentIdentifier,omitempty"`
	// 编号
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	// 所属项目id
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// 所属项目名称
	SpaceName *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	// 项目类型
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	// 状态名称
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 状态阶段id
	StatusStageIdentifier *string `json:"statusStageIdentifier,omitempty" xml:"statusStageIdentifier,omitempty"`
	// 工作项标题
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 状态更新时间
	UpdateStatusAt *int64 `json:"updateStatusAt,omitempty" xml:"updateStatusAt,omitempty"`
	// 工作项类型id
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

func (s *UpdateWorkItemResponseBodyWorkitem) SetStatus(v string) *UpdateWorkItemResponseBodyWorkitem {
	s.Status = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateWorkItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateWorkItemResponse) SetBody(v *UpdateWorkItemResponseBody) *UpdateWorkItemResponse {
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/flow/tags"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/flow/tagGroups"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/hostGroups"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/projects/createProject"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	resourceType = openapiutil.GetEncodeParam(resourceType)
	resourceId = openapiutil.GetEncodeParam(resourceId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/" + tea.StringValue(resourceType) + "/" + tea.StringValue(resourceId) + "/members"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/sprints/create"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSshKey"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/sshKey"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/variableGroups"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	id = openapiutil.GetEncodeParam(id)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFlowTag"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/flow/tags/" + tea.StringValue(id)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	id = openapiutil.GetEncodeParam(id)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFlowTagGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/flow/tagGroups/" + tea.StringValue(id)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	id = openapiutil.GetEncodeParam(id)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHostGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/hostGroups/" + tea.StringValue(id)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePipeline"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId)),
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

func (client *Client) DeleteProject(organizationId *string, projectId *string, request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(organizationId, projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectWithOptions(organizationId *string, projectId *string, request *DeleteProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	organizationId = openapiutil.GetEncodeParam(organizationId)
	projectId = openapiutil.GetEncodeParam(projectId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProject"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/projects/" + tea.StringValue(projectId) + "/delete"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	resourceType = openapiutil.GetEncodeParam(resourceType)
	resourceId = openapiutil.GetEncodeParam(resourceId)
	accountId = openapiutil.GetEncodeParam(accountId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceMember"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/" + tea.StringValue(resourceType) + "/" + tea.StringValue(resourceId) + "/members/" + tea.StringValue(accountId)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	id = openapiutil.GetEncodeParam(id)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVariableGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/variableGroups/" + tea.StringValue(id)),
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
	workspaceId = openapiutil.GetEncodeParam(workspaceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("FrozenWorkspace"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/workspaces/" + tea.StringValue(workspaceId) + "/frozen"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	fieldId = openapiutil.GetEncodeParam(fieldId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/fields/" + tea.StringValue(fieldId) + "/getCustomOption"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	id = openapiutil.GetEncodeParam(id)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetFlowTagGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/flow/tagGroups/" + tea.StringValue(id)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	id = openapiutil.GetEncodeParam(id)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetHostGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/hostGroups/" + tea.StringValue(id)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	accountId = openapiutil.GetEncodeParam(accountId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetOrganizationMember"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/members/" + tea.StringValue(accountId)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipeline"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipeline/getArtifactDownloadUrl"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	emasJobInstanceId = openapiutil.GetEncodeParam(emasJobInstanceId)
	md5 = openapiutil.GetEncodeParam(md5)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
	pipelineRunId = openapiutil.GetEncodeParam(pipelineRunId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipeline/" + tea.StringValue(pipelineId) + "/pipelineRun/" + tea.StringValue(pipelineRunId) + "/emas/artifact/" + tea.StringValue(emasJobInstanceId) + "/" + tea.StringValue(md5)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
	pipelineRunId = openapiutil.GetEncodeParam(pipelineRunId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineRun"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId) + "/pipelineRuns/" + tea.StringValue(pipelineRunId)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipeline/getPipelineScanReportUrl"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	projectId = openapiutil.GetEncodeParam(projectId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetProjectInfo"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/project/" + tea.StringValue(projectId)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
	deployOrderId = openapiutil.GetEncodeParam(deployOrderId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetVMDeployOrder"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId) + "/deploy/" + tea.StringValue(deployOrderId)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	id = openapiutil.GetEncodeParam(id)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetVariableGroup"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/variableGroups/" + tea.StringValue(id)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	workitemId = openapiutil.GetEncodeParam(workitemId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkItemActivity"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/workitems/" + tea.StringValue(workitemId) + "/getActivity"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	workitemId = openapiutil.GetEncodeParam(workitemId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkItemInfo"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/workitems/" + tea.StringValue(workitemId)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	workitemId = openapiutil.GetEncodeParam(workitemId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/workitems/" + tea.StringValue(workitemId) + "/getWorkflowInfo"),
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
	workspaceId = openapiutil.GetEncodeParam(workspaceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkspace"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/workspaces/" + tea.StringValue(workspaceId)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlowTagGroups"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/flow/tagGroups"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/hostGroups"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/members"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipeline/" + tea.StringValue(pipelineId) + "/job/historys"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipeline/" + tea.StringValue(pipelineId) + "/jobs"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId) + "/pipelineRuns"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	projectId = openapiutil.GetEncodeParam(projectId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/projects/" + tea.StringValue(projectId) + "/listMembers"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	resourceType = openapiutil.GetEncodeParam(resourceType)
	resourceId = openapiutil.GetEncodeParam(resourceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceMembers"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/" + tea.StringValue(resourceType) + "/" + tea.StringValue(resourceId) + "/members"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/serviceConnections"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/variableGroups"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
	jobId = openapiutil.GetEncodeParam(jobId)
	pipelineRunId = openapiutil.GetEncodeParam(pipelineRunId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("LogPipelineJobRun"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipeline/" + tea.StringValue(pipelineId) + "/pipelineRun/" + tea.StringValue(pipelineRunId) + "/job/" + tea.StringValue(jobId) + "/logs"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
	deployOrderId = openapiutil.GetEncodeParam(deployOrderId)
	machineSn = openapiutil.GetEncodeParam(machineSn)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("LogVMDeployMachine"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId) + "/deploy/" + tea.StringValue(deployOrderId) + "/machine/" + tea.StringValue(machineSn) + "/log"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
	pipelineRunId = openapiutil.GetEncodeParam(pipelineRunId)
	jobId = openapiutil.GetEncodeParam(jobId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("PassPipelineValidate"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId) + "/pipelineRuns/" + tea.StringValue(pipelineRunId) + "/jobs/" + tea.StringValue(jobId) + "/pass"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
	pipelineRunId = openapiutil.GetEncodeParam(pipelineRunId)
	jobId = openapiutil.GetEncodeParam(jobId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RefusePipelineValidate"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId) + "/pipelineRuns/" + tea.StringValue(pipelineRunId) + "/jobs/" + tea.StringValue(jobId) + "/refuse"),
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
	workspaceId = openapiutil.GetEncodeParam(workspaceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseWorkspace"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/workspaces/" + tea.StringValue(workspaceId) + "/release"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ResetSshKey"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/sshKey"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
	deployOrderId = openapiutil.GetEncodeParam(deployOrderId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeVMDeployOrder"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId) + "/deploy/" + tea.StringValue(deployOrderId) + "/resume"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
	pipelineRunId = openapiutil.GetEncodeParam(pipelineRunId)
	jobId = openapiutil.GetEncodeParam(jobId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RetryPipelineJobRun"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId) + "/pipelineRuns/" + tea.StringValue(pipelineRunId) + "/jobs/" + tea.StringValue(jobId)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
	deployOrderId = openapiutil.GetEncodeParam(deployOrderId)
	machineSn = openapiutil.GetEncodeParam(machineSn)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RetryVMDeployMachine"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId) + "/deploy/" + tea.StringValue(deployOrderId) + "/machine/" + tea.StringValue(machineSn) + "/retry"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
	pipelineRunId = openapiutil.GetEncodeParam(pipelineRunId)
	jobId = openapiutil.GetEncodeParam(jobId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("SkipPipelineJobRun"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId) + "/pipelineRuns/" + tea.StringValue(pipelineRunId) + "/jobs/" + tea.StringValue(jobId) + "/skip"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
	deployOrderId = openapiutil.GetEncodeParam(deployOrderId)
	machineSn = openapiutil.GetEncodeParam(machineSn)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("SkipVMDeployMachine"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId) + "/deploy/" + tea.StringValue(deployOrderId) + "/machine/" + tea.StringValue(machineSn) + "/skip"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
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
		Pathname:    tea.String("/organizations/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId) + "/run"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
	pipelineRunId = openapiutil.GetEncodeParam(pipelineRunId)
	jobId = openapiutil.GetEncodeParam(jobId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopPipelineJobRun"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId) + "/pipelineRuns/" + tea.StringValue(pipelineRunId) + "/jobs/" + tea.StringValue(jobId) + "/stop"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
	pipelineRunId = openapiutil.GetEncodeParam(pipelineRunId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopPipelineRun"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId) + "/pipelineRuns/" + tea.StringValue(pipelineRunId) + "/stop"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
	deployOrderId = openapiutil.GetEncodeParam(deployOrderId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopVMDeployOrder"),
		Version:     tea.String("2021-06-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId) + "/deploy/" + tea.StringValue(deployOrderId) + "/stop"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	id = openapiutil.GetEncodeParam(id)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/flow/tags/" + tea.StringValue(id)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	id = openapiutil.GetEncodeParam(id)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/flow/tagGroups/" + tea.StringValue(id)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	id = openapiutil.GetEncodeParam(id)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/hostGroups/" + tea.StringValue(id)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	pipelineId = openapiutil.GetEncodeParam(pipelineId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/pipelines/" + tea.StringValue(pipelineId) + "/baseInfo"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	projectId = openapiutil.GetEncodeParam(projectId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/projects/" + tea.StringValue(projectId) + "/updateMember"),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	resourceType = openapiutil.GetEncodeParam(resourceType)
	resourceId = openapiutil.GetEncodeParam(resourceId)
	accountId = openapiutil.GetEncodeParam(accountId)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/" + tea.StringValue(resourceType) + "/" + tea.StringValue(resourceId) + "/members/" + tea.StringValue(accountId)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	id = openapiutil.GetEncodeParam(id)
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/variableGroups/" + tea.StringValue(id)),
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
	organizationId = openapiutil.GetEncodeParam(organizationId)
	body := map[string]interface{}{}
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
		Pathname:    tea.String("/organization/" + tea.StringValue(organizationId) + "/workitems/update"),
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
