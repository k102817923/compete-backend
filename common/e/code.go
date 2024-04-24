package e

const (
	SUCCESS                                 = 200
	ERROR                                   = 500
	UNKNOWN                                 = 42000
	USER_AUTH_LIMITED                       = 42001
	RESOURCE_NOT_EXIST                      = 42002
	RESOURCE_EXIST                          = 42003
	RESOURCE_OCCUPY                         = 42004
	DELETE_COMPETE                          = 42100
	DURING_COMPETE_CANNOT_BE_MODIFIED       = 42101
	AFTER_COMPETE_STARTS_CANNOT_BE_MODIFIED = 42102
	AFTER_COMPETE_OVER_CANNOT_BE_MODIFIED   = 42103
	MAXIMUM_NUMBER_OF_SUBMISSIONS           = 42104
	CREATION_FAILED_EVENT_TYPE_INCORRECT    = 42105
	COMEPETE_DOESNT_EXIST                   = 42106
	INCORRECT_COMEPETE_TYPE                 = 42107
	USER_LIMIT                              = 42200
	ENROLMENT_EXIST                         = 42201
	ACCOUNT_DOES_NOT_EXIST                  = 42203
	REGISTRATION_FAILED                     = 42204
	JUDGE_EXIST                             = 42300
	RECOMMENDED_SPACE_FULL                  = 42400
	TEST_PAPER_DOES_NOT_EXIST               = 42500
	UPDATE_FAILED                           = 42999
	INVITATION_FAILED                       = 41307
	ENROLMENT_UPDATE_LIMIE                  = 41308
	COMPETE_NOT_PUBLISHED                   = 41309
	USER_PASSWORD_WRONG                     = 41310
	ANSWER_SHEET_RECORD_SUBMIT_ERROR        = 41311
	DELETE_COMPETE_JUDGE_ERROR              = 41312
	COMPETE_TEAM_UNIQUE                     = 41313
	COMPETE_TEAM_CELLPHONE_UNIQUE           = 41314
	RATE_LIMIT                              = 41315
	RATE_LIMIT_BY_THIRTY_MINS               = 41316
	CELLPHONE_VALIDATE_ERROR                = 41317
	CREATE_USER_ERROR                       = 41318
	TEAMMEMBER_ENROLMENT_EXIST              = 41319
	CELLPHONE_HAS_USEED_BY_LEADER           = 41320
	COMPETE_COPY_FAILED                     = 41321
	NOT_COMPETE_JUDGE                       = 41322
	COMPETE_REVIEW_OVER                     = 41323
)
