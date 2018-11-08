package common

var OkErrorCode int = 0
var OkErrorString string = "success"

var JsonParseErrorCode int = 10
var JsonParseErrorString string = "jsonParseError: "

var DbExecuteErrorCode int = 20
var DbExecuteErrorString string = "dbExecuteError: "

var MqttCommErrorCode int = 30
var MqttCommErrorString string = "mqttCommError: "

var DeleteFileErrorCode int = 40
var DeleteFileErrorString string = "deleteFileError: "

var ParamErrorCode int = 50
var ParamErrorString string = "paramError: "

var UserAlreadyExistErrorCode int = 100
var UserAlreadyExistErrorString string = "userAlreadyExistError: "

var UserUnRegisterErrorCode int = 110
var UserUnRegisterErrorString string = "userUnRegisterError: "

var UserPasswordIsFalseErrorCode int = 120
var UserPasswordIsFalseErrorString string = "userPasswordError: "

var UserCookieIsVaildErrorCode int = 130
var UserCookieIsVaildErrorString string = "userCookieIsVaildError: "

var DownloadPictureFailedErrorCode int = 140
var DownloadPictureFailedErrorString string = "downloadPictureFailedError: "

var AddPictureHeaderNoPathrootErrorCode int = 150
var AddPictureHeaderNoPathrootErrorString string = "addPictureHaderNoPathrootError: "

var DeletePicturePicurlIsnullErrorCode int = 160
var DeletePicturePicurlIsnullErrorString string = "deletePicturePicurlIsnullError: "
