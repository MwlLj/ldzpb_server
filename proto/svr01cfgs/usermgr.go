package svr01cfgs

// add user / delete user
var User string = join("user")

// login / logout
var User_login string = join("user/login")

// get useriuuid by phone number
var User_useruuid_way_phoneno string = join("user/useruuid/way/phoneno")

// get useruuid by email number
var User_useruuid_way_emailno string = join("user/useruuid/way/emailno")

// get useruuid by phone number or email number
var User_useruuid_way_no string = join("user/useruuid/way/no")

// get isexist by phone number
var User_status_exist_way_phoneno string = join("user/status/exist/way/phoneno")

// get isexist by email number
var User_status_exist_way_emailno string = join("user/status/exist/way/emailno")

// get isexist by phone number or email number
var User_status_exist_way_no string = join("user/status/exist/way/no")

// get password is true
var User_status_true_pwd string = join("user/status/true/pwd")
