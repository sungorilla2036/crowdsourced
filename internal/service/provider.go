/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package service

import (
	"github.com/google/wire"
	"github.com/sungorilla2036/crowdsourced/internal/service/action"
	"github.com/sungorilla2036/crowdsourced/internal/service/activity"
	"github.com/sungorilla2036/crowdsourced/internal/service/activity_common"
	"github.com/sungorilla2036/crowdsourced/internal/service/activity_queue"
	answercommon "github.com/sungorilla2036/crowdsourced/internal/service/answer_common"
	"github.com/sungorilla2036/crowdsourced/internal/service/auth"
	"github.com/sungorilla2036/crowdsourced/internal/service/badge"
	"github.com/sungorilla2036/crowdsourced/internal/service/collection"
	collectioncommon "github.com/sungorilla2036/crowdsourced/internal/service/collection_common"
	"github.com/sungorilla2036/crowdsourced/internal/service/comment"
	"github.com/sungorilla2036/crowdsourced/internal/service/comment_common"
	"github.com/sungorilla2036/crowdsourced/internal/service/config"
	"github.com/sungorilla2036/crowdsourced/internal/service/content"
	"github.com/sungorilla2036/crowdsourced/internal/service/dashboard"
	"github.com/sungorilla2036/crowdsourced/internal/service/event_queue"
	"github.com/sungorilla2036/crowdsourced/internal/service/export"
	"github.com/sungorilla2036/crowdsourced/internal/service/follow"
	"github.com/sungorilla2036/crowdsourced/internal/service/meta"
	"github.com/sungorilla2036/crowdsourced/internal/service/meta_common"
	"github.com/sungorilla2036/crowdsourced/internal/service/notice_queue"
	"github.com/sungorilla2036/crowdsourced/internal/service/notification"
	notficationcommon "github.com/sungorilla2036/crowdsourced/internal/service/notification_common"
	"github.com/sungorilla2036/crowdsourced/internal/service/object_info"
	"github.com/sungorilla2036/crowdsourced/internal/service/plugin_common"
	questioncommon "github.com/sungorilla2036/crowdsourced/internal/service/question_common"
	"github.com/sungorilla2036/crowdsourced/internal/service/rank"
	"github.com/sungorilla2036/crowdsourced/internal/service/reason"
	"github.com/sungorilla2036/crowdsourced/internal/service/report"
	"github.com/sungorilla2036/crowdsourced/internal/service/report_handle"
	"github.com/sungorilla2036/crowdsourced/internal/service/review"
	"github.com/sungorilla2036/crowdsourced/internal/service/revision_common"
	"github.com/sungorilla2036/crowdsourced/internal/service/role"
	"github.com/sungorilla2036/crowdsourced/internal/service/search_parser"
	"github.com/sungorilla2036/crowdsourced/internal/service/siteinfo"
	"github.com/sungorilla2036/crowdsourced/internal/service/siteinfo_common"
	"github.com/sungorilla2036/crowdsourced/internal/service/tag"
	tagcommon "github.com/sungorilla2036/crowdsourced/internal/service/tag_common"
	"github.com/sungorilla2036/crowdsourced/internal/service/uploader"
	"github.com/sungorilla2036/crowdsourced/internal/service/user_admin"
	usercommon "github.com/sungorilla2036/crowdsourced/internal/service/user_common"
	"github.com/sungorilla2036/crowdsourced/internal/service/user_external_login"
	"github.com/sungorilla2036/crowdsourced/internal/service/user_notification_config"
)

// ProviderSetService is providers.
var ProviderSetService = wire.NewSet(
	comment.NewCommentService,
	comment_common.NewCommentCommonService,
	report.NewReportService,
	content.NewVoteService,
	tag.NewTagService,
	follow.NewFollowService,
	collection.NewCollectionGroupService,
	collection.NewCollectionService,
	action.NewCaptchaService,
	auth.NewAuthService,
	content.NewUserService,
	content.NewQuestionService,
	content.NewAnswerService,
	export.NewEmailService,
	tagcommon.NewTagCommonService,
	usercommon.NewUserCommon,
	questioncommon.NewQuestionCommon,
	answercommon.NewAnswerCommon,
	uploader.NewUploaderService,
	collectioncommon.NewCollectionCommon,
	revision_common.NewRevisionService,
	content.NewRevisionService,
	rank.NewRankService,
	search_parser.NewSearchParser,
	content.NewSearchService,
	metacommon.NewMetaCommonService,
	object_info.NewObjService,
	report_handle.NewReportHandle,
	user_admin.NewUserAdminService,
	reason.NewReasonService,
	siteinfo_common.NewSiteInfoCommonService,
	siteinfo.NewSiteInfoService,
	notficationcommon.NewNotificationCommon,
	notification.NewNotificationService,
	activity.NewAnswerActivityService,
	dashboard.NewDashboardService,
	activity_common.NewActivityCommon,
	activity.NewActivityService,
	role.NewRoleService,
	role.NewUserRoleRelService,
	role.NewRolePowerRelService,
	user_external_login.NewUserExternalLoginService,
	user_external_login.NewUserCenterLoginService,
	plugin_common.NewPluginCommonService,
	config.NewConfigService,
	notice_queue.NewNotificationQueueService,
	activity_queue.NewActivityQueueService,
	user_notification_config.NewUserNotificationConfigService,
	notification.NewExternalNotificationService,
	notice_queue.NewNewQuestionNotificationQueueService,
	review.NewReviewService,
	meta.NewMetaService,
	event_queue.NewEventQueueService,
	badge.NewBadgeService,
	badge.NewBadgeEventService,
	badge.NewBadgeAwardService,
	badge.NewBadgeGroupService,
)
