// This file is be designed to define any common error so that we can use it in any service simply.

package errno

var (
	// Success
	Success = NewErrNo(SuccessCode, "Success")

	ServiceError             = NewErrNo(ServiceErrorCode, "service is unable to start successfully")
	ServiceInternalError     = NewErrNo(ServiceErrorCode, "service internal error")
	ParamError               = NewErrNo(ParamErrorCode, "parameter error")
	AuthorizationFailedError = NewErrNo(AuthorizationFailedErrCode, "authorization failed")

	// User
	UserExistedError = NewErrNo(ParamErrorCode, "user existed")

	// Comment
	UnexpectedTypeError     = NewErrNo(UnexpectedTypeErrorCode, "unexpected type")
	NotImplementError       = NewErrNo(NotImplementErrorCode, "not implement")
	SensitiveWordsError     = NewErrNo(SensitiveWordsErrorCode, "existed sensitive words")
	SensitiveWordsHTTPError = NewErrNo(ServiceErrorCode, "sensitive-words api error")

	// Favorite
	LikeNoExistError      = NewErrNo(LikeNoExistErrorCode, "you did not like the video")
	LikeAlreadyExistError = NewErrNo(LikeAlreadyExistErrorCode, "you already like the video")

	// Video
	FileUploadError = NewErrNo(FileUploadErrorCode, "upload meet error")
	NotVideoFile    = NewErrNo(FileUploadError.ErrorCode, "not video file")

	// Follow
	FollowYourselfError = NewErrNo(FollowYourselfErrorCode, "you should not follow yourself")
	NotFollowError      = NewErrNo(NotFollowErrorCode, "you are not following this user")
	AlreadyFollowError  = NewErrNo(AlreadyFollowErrorCode, "you already follow this user")
	UserNotFoundError   = NewErrNo(ParamErrorCode, "user not found")

	// Chat
	CharacterLimitError = NewErrNo(CharacterLimitErrorCode, "character limit error")
)
