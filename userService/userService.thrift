namespace go userapi

//图片验证码
struct ImgCaptchaRequest{
    1:string uuid
}

struct ImgCaptchaResponse{
    1:binary img
}

//邮箱验证码
struct SmsCaptchaRequest{
    1: string email
    2: string imgCode
    3: string uuid
}

struct SmsCaptchaResponse{
    1: string errno
    2: string errmsg
}

//注册
struct RegisterUserRequest{
    1: string email
    2: string password
    3: string smscode
}

struct RegisterUserResponse{
    1: string errno
    2: string errmsg
}

service userService{
    ImgCaptchaResponse getImgCaptcha(1:ImgCaptchaRequest req)
    SmsCaptchaResponse sendSmsCaptcha(1:SmsCaptchaRequest req)
    RegisterUserResponse registerUser(1:RegisterUserRequest req)
}