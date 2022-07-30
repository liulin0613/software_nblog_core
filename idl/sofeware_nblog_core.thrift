namespace go software_nblog_core

// 公共返回，所有 Resp 应该都要有
struct Error {
    1: required i64 ErrNo
    2: required string ErrTips
}

// hello
struct HelloReq {
    1: required string Name (api.query="name"); // 添加 api 注解为方便进行参数绑定
}

struct HelloResp {
    1: required string RespBody;
    2: optional string LogID

    255: required Error Error
}

// get Log Details
struct QueryLogReq {
    1: required string LogID // logID
    2: required string Day // 日期 形式： 20220102
}

struct QueryLogResp {
    1: required list<string> LogDetail
    2: optional string LogID

    255: required Error Error
}

// 一个 service 对应一个文件
service HelloService {
    //@1:hello
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
}

service QueryService {
    //@2:QueryLogDetails
    QueryLogResp QueryLog(1: QueryLogReq request) (api.get = "/api/admin/log")
}