// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddWayToCompositeWay(ctx context.Context, arg AddWayToCompositeWayParams) (CompositeWay, error)
	CountUsers(ctx context.Context, arg CountUsersParams) (int64, error)
	CountWaysByType(ctx context.Context, arg CountWaysByTypeParams) (int64, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (CreateCommentRow, error)
	CreateDayReport(ctx context.Context, arg CreateDayReportParams) (DayReport, error)
	CreateFavoriteUser(ctx context.Context, arg CreateFavoriteUserParams) (FavoriteUser, error)
	CreateFavoriteUserWay(ctx context.Context, arg CreateFavoriteUserWayParams) (FavoriteUsersWay, error)
	CreateFormerMentorsWay(ctx context.Context, arg CreateFormerMentorsWayParams) (FormerMentorsWay, error)
	CreateFromUserMentoringRequest(ctx context.Context, arg CreateFromUserMentoringRequestParams) (FromUserMentoringRequest, error)
	CreateJobDone(ctx context.Context, arg CreateJobDoneParams) (CreateJobDoneRow, error)
	CreateJobDonesJobTag(ctx context.Context, arg CreateJobDonesJobTagParams) (JobDonesJobTag, error)
	CreateJobTag(ctx context.Context, arg CreateJobTagParams) (JobTag, error)
	CreateMentorUserWay(ctx context.Context, arg CreateMentorUserWayParams) (MentorUsersWay, error)
	CreateMetric(ctx context.Context, arg CreateMetricParams) (Metric, error)
	CreatePlan(ctx context.Context, arg CreatePlanParams) (CreatePlanRow, error)
	CreatePlansJobTag(ctx context.Context, arg CreatePlansJobTagParams) (PlansJobTag, error)
	CreateProblem(ctx context.Context, arg CreateProblemParams) (CreateProblemRow, error)
	CreateToUserMentoringRequest(ctx context.Context, arg CreateToUserMentoringRequestParams) (ToUserMentoringRequest, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserTag(ctx context.Context, tagName string) (UserTag, error)
	CreateUsersUserTag(ctx context.Context, arg CreateUsersUserTagParams) (UsersUserTag, error)
	CreateWay(ctx context.Context, arg CreateWayParams) (CreateWayRow, error)
	CreateWayCollection(ctx context.Context, arg CreateWayCollectionParams) (WayCollection, error)
	CreateWayCollectionsWays(ctx context.Context, arg CreateWayCollectionsWaysParams) (WayCollectionsWay, error)
	CreateWayTag(ctx context.Context, name string) (WayTag, error)
	CreateWaysWayTag(ctx context.Context, arg CreateWaysWayTagParams) (WaysWayTag, error)
	DeleteComment(ctx context.Context, commentUuid pgtype.UUID) error
	DeleteFavoriteUserByIds(ctx context.Context, arg DeleteFavoriteUserByIdsParams) error
	DeleteFavoriteUserWayByIds(ctx context.Context, arg DeleteFavoriteUserWayByIdsParams) error
	DeleteFormerMentorWayIfExist(ctx context.Context, arg DeleteFormerMentorWayIfExistParams) error
	DeleteFromUserMentoringRequest(ctx context.Context, arg DeleteFromUserMentoringRequestParams) error
	DeleteJobDone(ctx context.Context, jobDoneUuid pgtype.UUID) error
	DeleteJobDonesJobTagByJobDoneId(ctx context.Context, arg DeleteJobDonesJobTagByJobDoneIdParams) error
	DeleteJobTagById(ctx context.Context, jobTagUuid pgtype.UUID) error
	DeleteMentorUserWayByIds(ctx context.Context, arg DeleteMentorUserWayByIdsParams) error
	DeleteMetric(ctx context.Context, metricsUuid pgtype.UUID) (Metric, error)
	DeletePlan(ctx context.Context, planUuid pgtype.UUID) error
	DeletePlansJobTagByIds(ctx context.Context, arg DeletePlansJobTagByIdsParams) error
	DeleteProblem(ctx context.Context, problemUuid pgtype.UUID) error
	DeleteToUserMentoringRequestByIds(ctx context.Context, arg DeleteToUserMentoringRequestByIdsParams) error
	DeleteUser(ctx context.Context, userUuid pgtype.UUID) error
	DeleteUserTagFromUser(ctx context.Context, arg DeleteUserTagFromUserParams) error
	DeleteWay(ctx context.Context, wayUuid pgtype.UUID) error
	DeleteWayCollection(ctx context.Context, wayCollectionsUuid pgtype.UUID) error
	DeleteWayCollectionsWaysByIds(ctx context.Context, arg DeleteWayCollectionsWaysByIdsParams) error
	DeleteWayFromCompositeWay(ctx context.Context, arg DeleteWayFromCompositeWayParams) error
	DeleteWayTagFromWay(ctx context.Context, arg DeleteWayTagFromWayParams) error
	GetDayReportsByRankRange(ctx context.Context, arg GetDayReportsByRankRangeParams) ([]GetDayReportsByRankRangeRow, error)
	GetDayReportsCountByWayId(ctx context.Context, wayUuid pgtype.UUID) (int64, error)
	GetFavoriteForUserUuidsByWayId(ctx context.Context, wayUuid pgtype.UUID) (int64, error)
	GetFavoriteUserByDonorUserId(ctx context.Context, donorUserUuid pgtype.UUID) ([]User, error)
	GetFavoriteUserUuidsByAcceptorUserId(ctx context.Context, acceptorUserUuid pgtype.UUID) ([]pgtype.UUID, error)
	GetFavoriteWaysByUserId(ctx context.Context, userUuid pgtype.UUID) ([]GetFavoriteWaysByUserIdRow, error)
	GetFormerMentorUsersByWayId(ctx context.Context, wayUuid pgtype.UUID) ([]User, error)
	GetFromUserMentoringRequestWaysByUserId(ctx context.Context, userUuid pgtype.UUID) ([]GetFromUserMentoringRequestWaysByUserIdRow, error)
	GetFromUserMentoringRequestWaysByWayId(ctx context.Context, wayUuid pgtype.UUID) ([]User, error)
	GetIsUserHavingPermissionsForComment(ctx context.Context, arg GetIsUserHavingPermissionsForCommentParams) (GetIsUserHavingPermissionsForCommentRow, error)
	GetIsUserHavingPermissionsForDayReport(ctx context.Context, arg GetIsUserHavingPermissionsForDayReportParams) (GetIsUserHavingPermissionsForDayReportRow, error)
	GetIsUserHavingPermissionsForJobDone(ctx context.Context, arg GetIsUserHavingPermissionsForJobDoneParams) (GetIsUserHavingPermissionsForJobDoneRow, error)
	GetIsUserHavingPermissionsForPlan(ctx context.Context, arg GetIsUserHavingPermissionsForPlanParams) (GetIsUserHavingPermissionsForPlanRow, error)
	GetIsUserHavingPermissionsForProblem(ctx context.Context, arg GetIsUserHavingPermissionsForProblemParams) (GetIsUserHavingPermissionsForProblemRow, error)
	GetJobDonesByDayReportUuids(ctx context.Context, dayReportUuids []pgtype.UUID) ([]GetJobDonesByDayReportUuidsRow, error)
	GetJobTagByUuid(ctx context.Context, jobTagUuid pgtype.UUID) (JobTag, error)
	GetLabelStatistics(ctx context.Context, arg GetLabelStatisticsParams) ([]GetLabelStatisticsRow, error)
	GetLastDayReportDate(ctx context.Context, wayUuids []pgtype.UUID) (GetLastDayReportDateRow, error)
	GetListCommentsByDayReportUuids(ctx context.Context, dayReportUuids []pgtype.UUID) ([]Comment, error)
	GetListDayReportsByWayUuid(ctx context.Context, arg GetListDayReportsByWayUuidParams) ([]DayReport, error)
	GetListJobTagsByWayUuid(ctx context.Context, wayUuid pgtype.UUID) ([]JobTag, error)
	GetListJobTagsByWayUuids(ctx context.Context, wayUuids []pgtype.UUID) ([]JobTag, error)
	GetListJobsDoneByDayReportId(ctx context.Context, dayReportUuid pgtype.UUID) ([]JobDone, error)
	GetListLabelsByLabelUuids(ctx context.Context, jobTagUuids []pgtype.UUID) ([]JobTag, error)
	GetListMetricsByWayUuid(ctx context.Context, wayUuid pgtype.UUID) ([]Metric, error)
	GetListPlansByDayReportId(ctx context.Context, dayReportUuid pgtype.UUID) ([]Plan, error)
	GetListProblemsByDayReportId(ctx context.Context, dayReportUuid pgtype.UUID) ([]Problem, error)
	GetListUserTagsByUserId(ctx context.Context, userUuid pgtype.UUID) ([]UserTag, error)
	GetListWayCollectionsByUserId(ctx context.Context, ownerUuid pgtype.UUID) ([]WayCollection, error)
	GetListWayTagsByWayId(ctx context.Context, wayUuid pgtype.UUID) ([]WayTag, error)
	GetListWayTagsByWayIds(ctx context.Context, wayUuids []pgtype.UUID) ([]GetListWayTagsByWayIdsRow, error)
	GetMentorUsersByWayId(ctx context.Context, wayUuid pgtype.UUID) ([]User, error)
	GetMentorUsersByWayIds(ctx context.Context, wayUuids []pgtype.UUID) ([]GetMentorUsersByWayIdsRow, error)
	GetMentoringWaysByMentorId(ctx context.Context, userUuid pgtype.UUID) ([]GetMentoringWaysByMentorIdRow, error)
	GetMentoringWaysCountByUserId(ctx context.Context, userUuid pgtype.UUID) (int64, error)
	GetOverallInformation(ctx context.Context, arg GetOverallInformationParams) (GetOverallInformationRow, error)
	GetOwnWaysByUserId(ctx context.Context, ownerUuid pgtype.UUID) ([]GetOwnWaysByUserIdRow, error)
	GetOwnWaysCountByUserId(ctx context.Context, userUuid pgtype.UUID) (int64, error)
	GetPlansByDayReportUuids(ctx context.Context, dayReportUuids []pgtype.UUID) ([]GetPlansByDayReportUuidsRow, error)
	GetPricingPlanByUserId(ctx context.Context, userUuid pgtype.UUID) (PricingPlanType, error)
	GetPrivateWaysCountByUserId(ctx context.Context, userUuid pgtype.UUID) (int64, error)
	GetProblemsByDayReportUuids(ctx context.Context, dollar_1 []pgtype.UUID) ([]Problem, error)
	GetTagsCountByUserId(ctx context.Context, userUuid pgtype.UUID) (int64, error)
	GetTimeSpentByDayChart(ctx context.Context, arg GetTimeSpentByDayChartParams) ([]GetTimeSpentByDayChartRow, error)
	GetToMentorUserRequestsByWayId(ctx context.Context, wayUuid pgtype.UUID) ([]pgtype.UUID, error)
	GetUserByEmail(ctx context.Context, userEmail string) (User, error)
	GetUserById(ctx context.Context, userUuid pgtype.UUID) (User, error)
	GetUserByIds(ctx context.Context, dollar_1 []pgtype.UUID) ([]User, error)
	GetUserTagByName(ctx context.Context, tagName string) (UserTag, error)
	GetUsersByIds(ctx context.Context, userUuids []pgtype.UUID) ([]GetUsersByIdsRow, error)
	GetWayById(ctx context.Context, wayUuid pgtype.UUID) (GetWayByIdRow, error)
	GetWayChildren(ctx context.Context, wayUuid pgtype.UUID) ([]pgtype.UUID, error)
	GetWayCollectionJoinWayByUserId(ctx context.Context, ownerUuid pgtype.UUID) ([]GetWayCollectionJoinWayByUserIdRow, error)
	GetWayCollectionsByUserId(ctx context.Context, ownerUuid pgtype.UUID) ([]WayCollection, error)
	GetWayCollectionsCountByUserId(ctx context.Context, userUuid pgtype.UUID) (int64, error)
	GetWayRelatedUsers(ctx context.Context, wayUuids []pgtype.UUID) ([]GetWayRelatedUsersRow, error)
	GetWayTagByName(ctx context.Context, wayTagName string) (WayTag, error)
	GetWaysByCollectionId(ctx context.Context, wayCollectionUuid pgtype.UUID) ([]GetWaysByCollectionIdRow, error)
	IsAllMetricsDone(ctx context.Context, wayUuid pgtype.UUID) (bool, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]ListUsersRow, error)
	ListWays(ctx context.Context, arg ListWaysParams) ([]ListWaysRow, error)
	RegenerateDbData(ctx context.Context) error
	RemoveEverything(ctx context.Context) error
	UpdateComment(ctx context.Context, arg UpdateCommentParams) (UpdateCommentRow, error)
	UpdateJobDone(ctx context.Context, arg UpdateJobDoneParams) (UpdateJobDoneRow, error)
	UpdateJobTag(ctx context.Context, arg UpdateJobTagParams) (JobTag, error)
	UpdateMetric(ctx context.Context, arg UpdateMetricParams) (UpdateMetricRow, error)
	UpdatePlan(ctx context.Context, arg UpdatePlanParams) (UpdatePlanRow, error)
	UpdatePricingPlanByUserId(ctx context.Context, arg UpdatePricingPlanByUserIdParams) (ProfileSetting, error)
	UpdateProblem(ctx context.Context, arg UpdateProblemParams) (UpdateProblemRow, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateWay(ctx context.Context, arg UpdateWayParams) (UpdateWayRow, error)
	UpdateWayCollection(ctx context.Context, arg UpdateWayCollectionParams) (WayCollection, error)
}

var _ Querier = (*Queries)(nil)
