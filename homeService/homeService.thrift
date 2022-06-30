namespace go homeapi

struct AreaRequest{

}

struct AreaResponse{
    1:binary areas
}

service homeService{
    AreaResponse getAreas(1:AreaRequest req)
}