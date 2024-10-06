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

package repo

import (
	"github.com/google/wire"
	"github.com/sungorilla2036/crowdsourced/internal/base/data"
	"github.com/sungorilla2036/crowdsourced/internal/repo/activity"
	"github.com/sungorilla2036/crowdsourced/internal/repo/activity_common"
	"github.com/sungorilla2036/crowdsourced/internal/repo/answer"
	"github.com/sungorilla2036/crowdsourced/internal/repo/auth"
	"github.com/sungorilla2036/crowdsourced/internal/repo/badge"
	"github.com/sungorilla2036/crowdsourced/internal/repo/badge_award"
	"github.com/sungorilla2036/crowdsourced/internal/repo/badge_group"
	"github.com/sungorilla2036/crowdsourced/internal/repo/captcha"
	"github.com/sungorilla2036/crowdsourced/internal/repo/collection"
	"github.com/sungorilla2036/crowdsourced/internal/repo/comment"
	"github.com/sungorilla2036/crowdsourced/internal/repo/config"
	"github.com/sungorilla2036/crowdsourced/internal/repo/export"
	"github.com/sungorilla2036/crowdsourced/internal/repo/limit"
	"github.com/sungorilla2036/crowdsourced/internal/repo/meta"
	"github.com/sungorilla2036/crowdsourced/internal/repo/notification"
	"github.com/sungorilla2036/crowdsourced/internal/repo/plugin_config"
	"github.com/sungorilla2036/crowdsourced/internal/repo/question"
	"github.com/sungorilla2036/crowdsourced/internal/repo/rank"
	"github.com/sungorilla2036/crowdsourced/internal/repo/reason"
	"github.com/sungorilla2036/crowdsourced/internal/repo/report"
	"github.com/sungorilla2036/crowdsourced/internal/repo/review"
	"github.com/sungorilla2036/crowdsourced/internal/repo/revision"
	"github.com/sungorilla2036/crowdsourced/internal/repo/role"
	"github.com/sungorilla2036/crowdsourced/internal/repo/search_common"
	"github.com/sungorilla2036/crowdsourced/internal/repo/site_info"
	"github.com/sungorilla2036/crowdsourced/internal/repo/tag"
	"github.com/sungorilla2036/crowdsourced/internal/repo/tag_common"
	"github.com/sungorilla2036/crowdsourced/internal/repo/unique"
	"github.com/sungorilla2036/crowdsourced/internal/repo/user"
	"github.com/sungorilla2036/crowdsourced/internal/repo/user_external_login"
	"github.com/sungorilla2036/crowdsourced/internal/repo/user_notification_config"
)

// ProviderSetRepo is data providers.
var ProviderSetRepo = wire.NewSet(
	data.NewData,
	data.NewDB,
	data.NewCache,
	comment.NewCommentRepo,
	comment.NewCommentCommonRepo,
	captcha.NewCaptchaRepo,
	unique.NewUniqueIDRepo,
	report.NewReportRepo,
	activity_common.NewFollowRepo,
	activity_common.NewVoteRepo,
	config.NewConfigRepo,
	user.NewUserRepo,
	user.NewUserAdminRepo,
	rank.NewUserRankRepo,
	question.NewQuestionRepo,
	answer.NewAnswerRepo,
	activity_common.NewActivityRepo,
	activity.NewVoteRepo,
	activity.NewFollowRepo,
	activity.NewAnswerActivityRepo,
	activity.NewUserActiveActivityRepo,
	activity.NewActivityRepo,
	activity.NewReviewActivityRepo,
	tag.NewTagRepo,
	tag_common.NewTagCommonRepo,
	tag.NewTagRelRepo,
	collection.NewCollectionRepo,
	collection.NewCollectionGroupRepo,
	auth.NewAuthRepo,
	revision.NewRevisionRepo,
	search_common.NewSearchRepo,
	meta.NewMetaRepo,
	export.NewEmailRepo,
	reason.NewReasonRepo,
	site_info.NewSiteInfo,
	notification.NewNotificationRepo,
	role.NewRoleRepo,
	role.NewUserRoleRelRepo,
	role.NewRolePowerRelRepo,
	role.NewPowerRepo,
	user_external_login.NewUserExternalLoginRepo,
	plugin_config.NewPluginConfigRepo,
	user_notification_config.NewUserNotificationConfigRepo,
	limit.NewRateLimitRepo,
	plugin_config.NewPluginUserConfigRepo,
	review.NewReviewRepo,
	badge.NewBadgeRepo,
	badge.NewEventRuleRepo,
	badge_group.NewBadgeGroupRepo,
	badge_award.NewBadgeAwardRepo,
)
