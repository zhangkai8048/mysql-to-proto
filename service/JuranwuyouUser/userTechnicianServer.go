
package JuranwuyouUser

import (
	"context"
    
    "zhangkai.com/mysql-to-proto/pb/JuranwuyouUser"
    
)

type UserTechnicianServer struct{}

func (server *UserTechnicianServer) GetUserTechnician(context context.Context, request *JuranwuyouUser.UserTechnicianFilter) (response *JuranwuyouUser.UserTechnicianRequest, err error) {
	panic("implement me")
    return 
}

func (server *UserTechnicianServer) CreateUserTechnician(context context.Context, request *JuranwuyouUser.UserTechnicianRequest) (response *JuranwuyouUser.UserTechnicianResponse, err error) {
	panic("implement me")
    return 
}

func (server *UserTechnicianServer) UpdateUserTechnician(context context.Context, request *JuranwuyouUser.UserTechnicianRequest) (response *JuranwuyouUser.UserTechnicianResponse, err error) {
	panic("implement me")
    return 
}

