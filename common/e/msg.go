package e

var MsgFlags = map[int]string{
	SUCCESS:                                 "ok",
	ERROR:                                   "fail",
	UNKNOWN:                                 "未知错误",
	USER_AUTH_LIMITED:                       "没有权限",
	RESOURCE_NOT_EXIST:                      "资源不存在",
	RESOURCE_EXIST:                          "资源已存在",
	RESOURCE_OCCUPY:                         "更新失败",
	DELETE_COMPETE:                          "资源被引用，请确认后再进行删除操作",
	DURING_COMPETE_CANNOT_BE_MODIFIED:       "比赛时间内，不能修改",
	AFTER_COMPETE_STARTS_CANNOT_BE_MODIFIED: "比赛开始后，不能修改",
	AFTER_COMPETE_OVER_CANNOT_BE_MODIFIED:   "比赛结束后，不能修改",
	MAXIMUM_NUMBER_OF_SUBMISSIONS:           "比赛时间内，不能进行删除操作",
	CREATION_FAILED_EVENT_TYPE_INCORRECT:    "提交次数上限",
	COMEPETE_DOESNT_EXIST:                   "赛事不存在",
	INCORRECT_COMEPETE_TYPE:                 "赛事类型不正确",
	USER_LIMIT:                              "已报名过该比赛，无需重复报名",
	ENROLMENT_EXIST:                         "评委已存在，不能重复邀请",
	ACCOUNT_DOES_NOT_EXIST:                  "最多配置十个轮播推荐位",
	REGISTRATION_FAILED:                     "导入数据超出5000条，无法导入。",
	JUDGE_EXIST:                             "试卷不存在",
	RECOMMENDED_SPACE_FULL:                  "账号不存在",
	TEST_PAPER_DOES_NOT_EXIST:               "报名失败",
	UPDATE_FAILED:                           "邀请失败",
	INVITATION_FAILED:                       "创建失败，赛事类型不正确",
	ENROLMENT_UPDATE_LIMIE:                  "当前不允许修改报名信息",
	COMPETE_NOT_PUBLISHED:                   "比赛已取消，请联系管理员核实！",
	USER_PASSWORD_WRONG:                     "密码错误",
	ANSWER_SHEET_RECORD_SUBMIT_ERROR:        "试卷提交失败",
	DELETE_COMPETE_JUDGE_ERROR:              "评审时间内不允许删除评委",
	COMPETE_TEAM_UNIQUE:                     "已存在该团队名称",
	COMPETE_TEAM_CELLPHONE_UNIQUE:           "团队成员电话信息不能重复，请重新输入！",
	RATE_LIMIT:                              "操作过于频繁，请稍后再试！",
	RATE_LIMIT_BY_THIRTY_MINS:               "操作过于频繁，请30分钟后再试！",
	CELLPHONE_VALIDATE_ERROR:                "手机号格式校验不通过",
	CREATE_USER_ERROR:                       "创建用户账号失败",
	TEAMMEMBER_ENROLMENT_EXIST:              "队员已报名，不允许重复报名！",
	CELLPHONE_HAS_USEED_BY_LEADER:           "此电话已被队长绑定，请更换手机号！",
	COMPETE_COPY_FAILED:                     "赛事复制失败",
	NOT_COMPETE_JUDGE:                       "评审失败，您不是赛事评委",
	COMPETE_REVIEW_OVER:                     "评审失败，评审时间已结束",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
