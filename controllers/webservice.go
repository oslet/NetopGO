package controllers

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	//	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
//var t f_datetime
type f_datetime time.Time

//var f_datetime = time.Now().Format("2006-01-02T15:04:05")

var _ xml.Name

//var t_time  "%d-%02d-%02dT%02d:%02d:%02d-00:00\n",t.Year(), t.Month(), t.Day(),t.Hour(), t.Minute(), t.Second())

type EUserType string

const (
	EUserType系统用户 EUserType = "系统用户"

	EUserType分店人员 EUserType = "分店人员"

	EUserType客服用户 EUserType = "客服用户"

	EUserType集团用户 EUserType = "集团用户"

	EUserType其它 EUserType = "其它"
)

type AuthenticateUser struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AuthenticateUser"`

	SUserCode string `xml:"sUserCode,omitempty"`

	SPassword string `xml:"sPassword,omitempty"`

	NSystemID int32 `xml:"nSystemID,omitempty"`
}

type AuthenticateUserResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AuthenticateUserResponse"`

	AuthenticateUserResult bool `xml:"AuthenticateUserResult,omitempty"`

	SErrMsg string `xml:"sErrMsg,omitempty"`
}

type AuthenticateUserEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AuthenticateUserEx"`

	SUserCode string `xml:"sUserCode,omitempty"`

	SPassword string `xml:"sPassword,omitempty"`

	NSystemID int32 `xml:"nSystemID,omitempty"`
}

type AuthenticateUserExResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AuthenticateUserExResponse"`

	AuthenticateUserExResult *ResultDataOfCUser `xml:"AuthenticateUserExResult,omitempty"`
}

type AuthenticateKeyCardUser struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AuthenticateKeyCardUser"`

	SUserCode string `xml:"sUserCode,omitempty"`

	SPassword string `xml:"sPassword,omitempty"`

	SDynamicPass string `xml:"sDynamicPass,omitempty"`

	NSystemID int32 `xml:"nSystemID,omitempty"`
}

type AuthenticateKeyCardUserResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AuthenticateKeyCardUserResponse"`

	AuthenticateKeyCardUserResult bool `xml:"AuthenticateKeyCardUserResult,omitempty"`

	SErrMsg string `xml:"sErrMsg,omitempty"`
}

type GetUserByName struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserByName"`

	SName string `xml:"sName,omitempty"`
}

type GetUserByNameResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserByNameResponse"`

	GetUserByNameResult *ArrayOfCUser `xml:"GetUserByNameResult,omitempty"`
}

type GetUserMessageByUserCode struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserMessageByUserCode"`

	SUserCode string `xml:"sUserCode,omitempty"`
}

type GetUserMessageByUserCodeResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserMessageByUserCodeResponse"`

	GetUserMessageByUserCodeResult *CUser `xml:"GetUserMessageByUserCodeResult,omitempty"`
}

type AuthenticateUserEx2 struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AuthenticateUserEx2"`

	SUserCode string `xml:"sUserCode,omitempty"`

	SPassword string `xml:"sPassword,omitempty"`

	NSystemID int32 `xml:"nSystemID,omitempty"`
}

type AuthenticateUserEx2Response struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AuthenticateUserEx2Response"`

	AuthenticateUserEx2Result *ResultDataOfCUser `xml:"AuthenticateUserEx2Result,omitempty"`
}

type AuthenticateUserEx3 struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AuthenticateUserEx3"`

	SUserCode string `xml:"sUserCode,omitempty"`

	SPassword string `xml:"sPassword,omitempty"`

	NSystemID int32 `xml:"nSystemID,omitempty"`
}

type AuthenticateUserEx3Response struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AuthenticateUserEx3Response"`

	AuthenticateUserEx3Result *ResultDataOfCUserEx `xml:"AuthenticateUserEx3Result,omitempty"`
}

type GetUserByUserCode struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserByUserCode"`

	SUserCode string `xml:"sUserCode,omitempty"`
}

type GetUserByUserCodeResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserByUserCodeResponse"`

	GetUserByUserCodeResult *CUser `xml:"GetUserByUserCodeResult,omitempty"`
}

type GetUserByUserCodeEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserByUserCodeEx"`

	SUserCode string `xml:"sUserCode,omitempty"`
}

type GetUserByUserCodeExResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserByUserCodeExResponse"`

	GetUserByUserCodeExResult *CUserEx `xml:"GetUserByUserCodeExResult,omitempty"`
}

type GetUserByUserIDEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserByUserIDEx"`

	NUserID int32 `xml:"nUserID,omitempty"`
}

type GetUserByUserIDExResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserByUserIDExResponse"`

	GetUserByUserIDExResult *CUserEx3 `xml:"GetUserByUserIDExResult,omitempty"`
}

type GetUserByDept struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserByDept"`

	NGotoPage int32 `xml:"nGotoPage,omitempty"`

	NPageSize int32 `xml:"nPageSize,omitempty"`

	NDept string `xml:"nDept,omitempty"`
}

type GetUserByDeptResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserByDeptResponse"`

	GetUserByDeptResult *RowAbilityDateSetOfListOfCUser `xml:"GetUserByDeptResult,omitempty"`
}

type GetUserByType struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserByType"`

	NType int32 `xml:"nType,omitempty"`
}

type GetUserByTypeResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserByTypeResponse"`

	GetUserByTypeResult *ArrayOfCUser `xml:"GetUserByTypeResult,omitempty"`
}

type GetGroupChainUser struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetGroupChainUser"`
}

type GetGroupChainUserResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetGroupChainUserResponse"`

	GetGroupChainUserResult *ArrayOfCUser `xml:"GetGroupChainUserResult,omitempty"`
}

type GetUserByStation struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserByStation"`

	NGotoPage int32 `xml:"nGotoPage,omitempty"`

	NPageSize int32 `xml:"nPageSize,omitempty"`

	NStation int32 `xml:"nStation,omitempty"`
}

type GetUserByStationResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserByStationResponse"`

	GetUserByStationResult *RowAbilityDateSetOfListOfCUser `xml:"GetUserByStationResult,omitempty"`
}

type QueryUser struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ QueryUser"`

	NGotoPage int32 `xml:"nGotoPage,omitempty"`

	NPageSize int32 `xml:"nPageSize,omitempty"`

	SUserName string `xml:"sUserName,omitempty"`

	SMobile string `xml:"sMobile,omitempty"`

	SEmail string `xml:"sEmail,omitempty"`

	SDeptName string `xml:"sDeptName,omitempty"`

	SStationName string `xml:"sStationName,omitempty"`
}

type QueryUserResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ QueryUserResponse"`

	QueryUserResult *RowAbilityDateSetOfListOfCUser `xml:"QueryUserResult,omitempty"`
}

type QueryUserEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ QueryUserEx"`

	NGotoPage int32 `xml:"nGotoPage,omitempty"`

	NPageSize int32 `xml:"nPageSize,omitempty"`

	SQueryString string `xml:"sQueryString,omitempty"`
}

type QueryUserExResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ QueryUserExResponse"`

	QueryUserExResult *RowAbilityDateSetOfListOfCUser `xml:"QueryUserExResult,omitempty"`
}

type QueryUserEx2 struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ QueryUserEx2"`

	NGotoPage int32 `xml:"nGotoPage,omitempty"`

	NPageSize int32 `xml:"nPageSize,omitempty"`

	SQueryString string `xml:"sQueryString,omitempty"`
}

type QueryUserEx2Response struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ QueryUserEx2Response"`

	QueryUserEx2Result *RowAbilityDateSetOfListOfCUserEx2 `xml:"QueryUserEx2Result,omitempty"`
}

type GetManagerMessage struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetManagerMessage"`

	ChainID string `xml:"chainID,omitempty"`
}

type GetManagerMessageResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetManagerMessageResponse"`

	GetManagerMessageResult *ArrayOfCUser `xml:"GetManagerMessageResult,omitempty"`
}

type GetUserInfoByChainIDAndStationID struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserInfoByChainIDAndStationID"`

	ChainID int32 `xml:"chainID,omitempty"`

	StationID int32 `xml:"stationID,omitempty"`
}

type GetUserInfoByChainIDAndStationIDResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserInfoByChainIDAndStationIDResponse"`

	GetUserInfoByChainIDAndStationIDResult *ArrayOfManagerInfo `xml:"GetUserInfoByChainIDAndStationIDResult,omitempty"`
}

type GetUserByUserID struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserByUserID"`

	SUserID int32 `xml:"sUserID,omitempty"`
}

type GetUserByUserIDResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetUserByUserIDResponse"`

	GetUserByUserIDResult *CUser `xml:"GetUserByUserIDResult,omitempty"`
}

type QueryAllDept struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ QueryAllDept"`

	NGotoPage int32 `xml:"nGotoPage,omitempty"`

	NPageSize int32 `xml:"nPageSize,omitempty"`
}

type QueryAllDeptResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ QueryAllDeptResponse"`

	QueryAllDeptResult *RowAbilityDateSetOfListOfDept `xml:"QueryAllDeptResult,omitempty"`
}

type QueryAllDeptEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ QueryAllDeptEx"`

	NGotoPage int32 `xml:"nGotoPage,omitempty"`

	NPageSize int32 `xml:"nPageSize,omitempty"`
}

type QueryAllDeptExResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ QueryAllDeptExResponse"`

	QueryAllDeptExResult *RowAbilityDateSetOfListOfDept `xml:"QueryAllDeptExResult,omitempty"`
}

type GetAllSealedDept struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetAllSealedDept"`
}

type GetAllSealedDeptResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetAllSealedDeptResponse"`

	GetAllSealedDeptResult *ArrayOfDept `xml:"GetAllSealedDeptResult,omitempty"`
}

type GetAllChildDeptByDeptID struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetAllChildDeptByDeptID"`

	NDeptID string `xml:"nDeptID,omitempty"`
}

type GetAllChildDeptByDeptIDResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetAllChildDeptByDeptIDResponse"`

	GetAllChildDeptByDeptIDResult *ArrayOfDept `xml:"GetAllChildDeptByDeptIDResult,omitempty"`
}

type GetAllChildDeptEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetAllChildDeptEx"`

	NDeptID string `xml:"nDeptID,omitempty"`

	SQueryString string `xml:"sQueryString,omitempty"`
}

type GetAllChildDeptExResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetAllChildDeptExResponse"`

	GetAllChildDeptExResult *ArrayOfDept `xml:"GetAllChildDeptExResult,omitempty"`
}

type GetDeptByCode struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetDeptByCode"`

	SDeptID string `xml:"sDeptID,omitempty"`
}

type GetDeptByCodeResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetDeptByCodeResponse"`

	GetDeptByCodeResult *Dept `xml:"GetDeptByCodeResult,omitempty"`
}

type GetAllInn struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetAllInn"`
}

type GetAllInnResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetAllInnResponse"`

	GetAllInnResult *ArrayOfInn `xml:"GetAllInnResult,omitempty"`
}

type GetDeptByQueryString struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetDeptByQueryString"`

	NGotoPage int32 `xml:"nGotoPage,omitempty"`

	NPageSize int32 `xml:"nPageSize,omitempty"`

	SQueryString string `xml:"sQueryString,omitempty"`
}

type GetDeptByQueryStringResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetDeptByQueryStringResponse"`

	GetDeptByQueryStringResult *RowAbilityDateSetOfListOfDept `xml:"GetDeptByQueryStringResult,omitempty"`
}

type GetDeptByExpand struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetDeptByExpand"`

	NGotoPage int32 `xml:"nGotoPage,omitempty"`

	NPageSize int32 `xml:"nPageSize,omitempty"`

	SExpand string `xml:"sExpand,omitempty"`
}

type GetDeptByExpandResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetDeptByExpandResponse"`

	GetDeptByExpandResult *RowAbilityDateSetOfListOfDept `xml:"GetDeptByExpandResult,omitempty"`
}

type GetAllStation struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetAllStation"`
}

type GetAllStationResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetAllStationResponse"`

	GetAllStationResult *ArrayOfStation `xml:"GetAllStationResult,omitempty"`
}

type AddCheckRoles struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AddCheckRoles"`

	List *ArrayOfCheckRoles `xml:"list,omitempty"`
}

type AddCheckRolesResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AddCheckRolesResponse"`

	AddCheckRolesResult bool `xml:"AddCheckRolesResult,omitempty"`
}

type AddCheckGroupRoles struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AddCheckGroupRoles"`

	List *ArrayOfCheckGroupOnRoles `xml:"list,omitempty"`
}

type AddCheckGroupRolesResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AddCheckGroupRolesResponse"`

	AddCheckGroupRolesResult bool `xml:"AddCheckGroupRolesResult,omitempty"`
}

type AddCheckUserRoles struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AddCheckUserRoles"`

	List *ArrayOfCheckUserOnRoles `xml:"list,omitempty"`
}

type AddCheckUserRolesResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AddCheckUserRolesResponse"`

	AddCheckUserRolesResult bool `xml:"AddCheckUserRolesResult,omitempty"`
}

type UpdateUser struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ UpdateUser"`

	ObjUser *CUser `xml:"objUser,omitempty"`
}

type UpdateUserResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ UpdateUserResponse"`
}

type GetPeopleByChainID struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetPeopleByChainID"`

	NChainID int32 `xml:"nChainID,omitempty"`
}

type GetPeopleByChainIDResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetPeopleByChainIDResponse"`

	GetPeopleByChainIDResult *ArrayOfPeople `xml:"GetPeopleByChainIDResult,omitempty"`
}

type GetPeopleByChainIDEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetPeopleByChainIDEx"`

	NChainID int32 `xml:"nChainID,omitempty"`
}

type GetPeopleByChainIDExResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetPeopleByChainIDExResponse"`

	GetPeopleByChainIDExResult *ArrayOfPeopleEx `xml:"GetPeopleByChainIDExResult,omitempty"`
}

type GetChainIDByMebCardCode struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetChainIDByMebCardCode"`

	SMebCardCode string `xml:"sMebCardCode,omitempty"`
}

type GetChainIDByMebCardCodeResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetChainIDByMebCardCodeResponse"`

	GetChainIDByMebCardCodeResult string `xml:"GetChainIDByMebCardCodeResult,omitempty"`
}

type GetManageSalerByMebCardCode struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetManageSalerByMebCardCode"`

	SMebCardCode string `xml:"sMebCardCode,omitempty"`
}

type GetManageSalerByMebCardCodeResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetManageSalerByMebCardCodeResponse"`

	GetManageSalerByMebCardCodeResult *ArrayOfSaleManager `xml:"GetManageSalerByMebCardCodeResult,omitempty"`
}

type UpdateSaleManagerCodeByChainID struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ UpdateSaleManagerCodeByChainID"`

	NChainID int32 `xml:"nChainID,omitempty"`

	NSaleManagerCode int64 `xml:"nSaleManagerCode,omitempty"`

	NType int32 `xml:"nType,omitempty"`
}

type UpdateSaleManagerCodeByChainIDResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ UpdateSaleManagerCodeByChainIDResponse"`

	UpdateSaleManagerCodeByChainIDResult bool `xml:"UpdateSaleManagerCodeByChainIDResult,omitempty"`
}

type GetAllChainUsers struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetAllChainUsers"`

	NChainID int32 `xml:"nChainID,omitempty"`
}

type GetAllChainUsersResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ GetAllChainUsersResponse"`

	GetAllChainUsersResult *ArrayOfCCUser `xml:"GetAllChainUsersResult,omitempty"`
}

type AddUser struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AddUser"`

	ObjUser *AddInnUser `xml:"objUser,omitempty"`

	SOperateMebCard string `xml:"sOperateMebCard,omitempty"`

	SOperateUserName string `xml:"sOperateUserName,omitempty"`
}

type AddUserResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AddUserResponse"`

	AddUserResult int32 `xml:"AddUserResult,omitempty"`
}

type AddUserEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AddUserEx"`

	ObjUser *AddInnUser `xml:"objUser,omitempty"`

	SOperateMebCard string `xml:"sOperateMebCard,omitempty"`

	SOperateUserName string `xml:"sOperateUserName,omitempty"`
}

type AddUserExResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AddUserExResponse"`

	AddUserExResult int32 `xml:"AddUserExResult,omitempty"`

	SErrMsg string `xml:"sErrMsg,omitempty"`
}

type AddUserEx2 struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AddUserEx2"`

	ObjUser *AddInnUser `xml:"objUser,omitempty"`

	SOperateMebCard string `xml:"sOperateMebCard,omitempty"`

	SOperateUserName string `xml:"sOperateUserName,omitempty"`
}

type AddUserEx2Response struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AddUserEx2Response"`

	AddUserEx2Result int32 `xml:"AddUserEx2Result,omitempty"`

	SErrMsg string `xml:"sErrMsg,omitempty"`
}

type ModifyUser struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ModifyUser"`

	NUserID int32 `xml:"nUserID,omitempty"`

	ObjUser *ModifyInnUser `xml:"objUser,omitempty"`

	SOperateMebCard string `xml:"sOperateMebCard,omitempty"`

	SOperateUserName string `xml:"sOperateUserName,omitempty"`
}

type ModifyUserResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ModifyUserResponse"`

	ModifyUserResult bool `xml:"ModifyUserResult,omitempty"`
}

type ModifyUserEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ModifyUserEx"`

	SMebCardCode string `xml:"sMebCardCode,omitempty"`

	ObjUser *ModifyInnUserEx `xml:"objUser,omitempty"`

	SOperateMebCard string `xml:"sOperateMebCard,omitempty"`

	SOperateUserName string `xml:"sOperateUserName,omitempty"`
}

type ModifyUserExResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ModifyUserExResponse"`

	ModifyUserExResult bool `xml:"ModifyUserExResult,omitempty"`
}

type DisableUser struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ DisableUser"`

	NUserID int32 `xml:"nUserID,omitempty"`

	SOperateMebCard string `xml:"sOperateMebCard,omitempty"`

	SOperateUserName string `xml:"sOperateUserName,omitempty"`
}

type DisableUserResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ DisableUserResponse"`

	DisableUserResult bool `xml:"DisableUserResult,omitempty"`
}

type EnableUser struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ EnableUser"`

	NMebCardCode string `xml:"nMebCardCode,omitempty"`

	SOperateMebCard string `xml:"sOperateMebCard,omitempty"`

	SOperateUserName string `xml:"sOperateUserName,omitempty"`
}

type EnableUserResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ EnableUserResponse"`

	EnableUserResult bool `xml:"EnableUserResult,omitempty"`
}

type CheckMebCardIsExist struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ CheckMebCardIsExist"`

	SMebCard string `xml:"sMebCard,omitempty"`
}

type CheckMebCardIsExistResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ CheckMebCardIsExistResponse"`

	CheckMebCardIsExistResult bool `xml:"CheckMebCardIsExistResult,omitempty"`
}

type ModifyPassword struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ModifyPassword"`

	NUserID int32 `xml:"nUserID,omitempty"`

	SNewPassword string `xml:"sNewPassword,omitempty"`

	SOperateMebCard string `xml:"sOperateMebCard,omitempty"`

	SOperateUserName string `xml:"sOperateUserName,omitempty"`
}

type ModifyPasswordResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ModifyPasswordResponse"`

	ModifyPasswordResult bool `xml:"ModifyPasswordResult,omitempty"`
}

type ModifyPasswordEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ModifyPasswordEx"`

	NUserID int32 `xml:"nUserID,omitempty"`

	SOldPassword string `xml:"sOldPassword,omitempty"`

	SNewPassword string `xml:"sNewPassword,omitempty"`

	SOperateMebCard string `xml:"sOperateMebCard,omitempty"`

	SOperateUserName string `xml:"sOperateUserName,omitempty"`
}

type ModifyPasswordExResponse struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ModifyPasswordExResponse"`

	ModifyPasswordExResult bool `xml:"ModifyPasswordExResult,omitempty"`
}

type ModifyPasswordEx2 struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ModifyPasswordEx2"`

	NUserID int32 `xml:"nUserID,omitempty"`

	SOldPassword string `xml:"sOldPassword,omitempty"`

	SNewPassword string `xml:"sNewPassword,omitempty"`

	SOperateMebCard string `xml:"sOperateMebCard,omitempty"`

	SOperateUserName string `xml:"sOperateUserName,omitempty"`
}

type ModifyPasswordEx2Response struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ModifyPasswordEx2Response"`

	ModifyPasswordEx2Result bool `xml:"ModifyPasswordEx2Result,omitempty"`

	SErrMsg string `xml:"sErrMsg,omitempty"`
}

type ModifyPasswordEx3 struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ModifyPasswordEx3"`

	NUserID int32 `xml:"nUserID,omitempty"`

	SNewPassword string `xml:"sNewPassword,omitempty"`

	SOperateMebCard string `xml:"sOperateMebCard,omitempty"`

	SOperateUserName string `xml:"sOperateUserName,omitempty"`
}

type ModifyPasswordEx3Response struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ModifyPasswordEx3Response"`

	ModifyPasswordEx3Result bool `xml:"ModifyPasswordEx3Result,omitempty"`

	SErrMsg string `xml:"sErrMsg,omitempty"`
}

type ResultDataOfCUser struct {
	//XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ResultDataOfCUser"`

	Success bool `xml:"_success,omitempty"`

	Success1 bool `xml:"Success,omitempty"`

	ErrMsg string `xml:"ErrMsg,omitempty"`

	ObjData *CUser `xml:"ObjData,omitempty"`
}

type CUser struct {
	//XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ CUser"`

	UserID int32 `xml:"UserID,omitempty"`

	UserCode string `xml:"UserCode,omitempty"`

	Name string `xml:"Name,omitempty"`

	Title string `xml:"Title,omitempty"`

	Type *EUserType `xml:"Type,omitempty"`

	MebCardCode string `xml:"MebCardCode,omitempty"`

	Flag int32 `xml:"Flag,omitempty"`

	Mobile string `xml:"Mobile,omitempty"`

	Password string `xml:"Password,omitempty"`

	EMail string `xml:"EMail,omitempty"`

	RoleList string `xml:"RoleList,omitempty"`

	SecurityGrade int32 `xml:"SecurityGrade,omitempty"`

	RoleGroupID int32 `xml:"RoleGroupID,omitempty"`

	UserRoleGroupList string `xml:"UserRoleGroupList,omitempty"`

	Passport *Passport `xml:"Passport,omitempty"`

	TelePhone string `xml:"TelePhone,omitempty"`

	IDCardCode string `xml:"IDCardCode,omitempty"`

	Sex string `xml:"Sex,omitempty"`

	Age int32 `xml:"Age,omitempty"`

	Birthday f_datetime `xml:"Birthday,omitempty"`

	Education string `xml:"Education,omitempty"`

	Nation string `xml:"Nation,omitempty"`

	Height string `xml:"Height,omitempty"`

	BirthPlace string `xml:"BirthPlace,omitempty"`

	GraduateSchool string `xml:"GraduateSchool,omitempty"`

	MainOU int32 `xml:"MainOU,omitempty"`

	Address string `xml:"Address,omitempty"`

	DeptList *ArrayOfDeptInfo `xml:"DeptList,omitempty"`

	JoinTime f_datetime `xml:"JoinTime,omitempty"`

	LeaveTime f_datetime `xml:"LeaveTime,omitempty"`

	LeaderUserID int32 `xml:"LeaderUserID,omitempty"`
}

type Passport struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ Passport"`

	MDPassWord string `xml:"MDPassWord,omitempty"`

	PWDTime f_datetime `xml:"PWDTime,omitempty"`

	Code string `xml:"Code,omitempty"`
}

type ArrayOfDeptInfo struct {
	//XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfDeptInfo"`

	DeptInfo []*DeptInfo `xml:"DeptInfo,omitempty"`
}

type DeptInfo struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ DeptInfo"`

	DeptID string `xml:"DeptID,omitempty"`

	DeptName string `xml:"DeptName,omitempty"`

	ParentDeptID string `xml:"ParentDeptID,omitempty"`

	OUID int32 `xml:"OUID,omitempty"`

	PositionID int32 `xml:"PositionID,omitempty"`

	PositionName string `xml:"PositionName,omitempty"`

	StationID int32 `xml:"StationID,omitempty"`

	StationName string `xml:"StationName,omitempty"`

	Remark string `xml:"Remark,omitempty"`

	Flag int32 `xml:"Flag,omitempty"`

	ChainID int32 `xml:"ChainID,omitempty"`

	LeaderUserID string `xml:"LeaderUserID,omitempty"`

	BrandID int32 `xml:"BrandID,omitempty"`
}

type ArrayOfCUser struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfCUser"`

	CUser []*CUser `xml:"CUser,omitempty"`
}

type ResultDataOfCUserEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ResultDataOfCUserEx"`

	Success bool `xml:"_success,omitempty"`

	Success1 bool `xml:"Success,omitempty"`

	ErrMsg string `xml:"ErrMsg,omitempty"`

	ObjData *CUserEx `xml:"ObjData,omitempty"`
}

type CUserEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ CUserEx"`

	UserID int32 `xml:"UserID,omitempty"`

	UserCode string `xml:"UserCode,omitempty"`

	Name string `xml:"Name,omitempty"`

	Title string `xml:"Title,omitempty"`

	Type *EUserType `xml:"Type,omitempty"`

	MebCardCode string `xml:"MebCardCode,omitempty"`

	Flag int32 `xml:"Flag,omitempty"`

	Mobile string `xml:"Mobile,omitempty"`

	Password string `xml:"Password,omitempty"`

	EMail string `xml:"EMail,omitempty"`

	RoleList string `xml:"RoleList,omitempty"`

	SecurityGrade int32 `xml:"SecurityGrade,omitempty"`

	RoleGroupID int32 `xml:"RoleGroupID,omitempty"`

	UserRoleGroupList string `xml:"UserRoleGroupList,omitempty"`

	Passport *Passport `xml:"Passport,omitempty"`

	TelePhone string `xml:"TelePhone,omitempty"`

	IDCardCode string `xml:"IDCardCode,omitempty"`

	Sex string `xml:"Sex,omitempty"`

	Age int32 `xml:"Age,omitempty"`

	Birthday f_datetime `xml:"Birthday,omitempty"`

	Education string `xml:"Education,omitempty"`

	Nation string `xml:"Nation,omitempty"`

	Height string `xml:"Height,omitempty"`

	BirthPlace string `xml:"BirthPlace,omitempty"`

	GraduateSchool string `xml:"GraduateSchool,omitempty"`

	MainOU int32 `xml:"MainOU,omitempty"`

	Address string `xml:"Address,omitempty"`

	DeptListEx *ArrayOfDeptInfoEx `xml:"DeptListEx,omitempty"`

	JoinTime f_datetime `xml:"JoinTime,omitempty"`

	LeaveTime f_datetime `xml:"LeaveTime,omitempty"`

	LeaderUserID int32 `xml:"LeaderUserID,omitempty"`
}

type ArrayOfDeptInfoEx struct {
	//XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfDeptInfoEx"`

	DeptInfoEx []*DeptInfoEx `xml:"DeptInfoEx,omitempty"`
}

type DeptInfoEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ DeptInfoEx"`

	DeptID string `xml:"DeptID,omitempty"`

	DeptName string `xml:"DeptName,omitempty"`

	ParentDeptID string `xml:"ParentDeptID,omitempty"`

	OUID int32 `xml:"OUID,omitempty"`

	PositionID int32 `xml:"PositionID,omitempty"`

	PositionName string `xml:"PositionName,omitempty"`

	StationID int32 `xml:"StationID,omitempty"`

	StationName string `xml:"StationName,omitempty"`

	Remark string `xml:"Remark,omitempty"`

	Flag int32 `xml:"Flag,omitempty"`

	ChainID int32 `xml:"ChainID,omitempty"`

	LeaderUserID string `xml:"LeaderUserID,omitempty"`

	BrandID int32 `xml:"BrandID,omitempty"`

	GroupID int32 `xml:"GroupID,omitempty"`
}

type CUserEx3 struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ CUserEx3"`

	UserID int32 `xml:"UserID,omitempty"`

	UserCode string `xml:"UserCode,omitempty"`

	Name string `xml:"Name,omitempty"`
}

type RowAbilityDateSetOfListOfCUser struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ RowAbilityDateSetOfListOfCUser"`

	DataSet *ArrayOfCUser `xml:"DataSet,omitempty"`

	RowCount int32 `xml:"RowCount,omitempty"`

	PageCount int32 `xml:"PageCount,omitempty"`

	CurrentPage int32 `xml:"CurrentPage,omitempty"`

	PageRowCount int32 `xml:"PageRowCount,omitempty"`
}

type RowAbilityDateSetOfListOfCUserEx2 struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ RowAbilityDateSetOfListOfCUserEx2"`

	DataSet *ArrayOfCUserEx2 `xml:"DataSet,omitempty"`

	RowCount int32 `xml:"RowCount,omitempty"`

	PageCount int32 `xml:"PageCount,omitempty"`

	CurrentPage int32 `xml:"CurrentPage,omitempty"`

	PageRowCount int32 `xml:"PageRowCount,omitempty"`
}

type ArrayOfCUserEx2 struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfCUserEx2"`

	CUserEx2 []*CUserEx2 `xml:"CUserEx2,omitempty"`
}

type CUserEx2 struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ CUserEx2"`

	UserID int32 `xml:"UserID,omitempty"`

	UserCode string `xml:"UserCode,omitempty"`

	Name string `xml:"Name,omitempty"`

	Title string `xml:"Title,omitempty"`

	Type *EUserType `xml:"Type,omitempty"`

	MebCardCode string `xml:"MebCardCode,omitempty"`

	Flag int32 `xml:"Flag,omitempty"`

	Mobile string `xml:"Mobile,omitempty"`

	Password string `xml:"Password,omitempty"`

	EMail string `xml:"EMail,omitempty"`

	RoleList string `xml:"RoleList,omitempty"`

	SecurityGrade int32 `xml:"SecurityGrade,omitempty"`

	RoleGroupID int32 `xml:"RoleGroupID,omitempty"`

	UserRoleGroupList string `xml:"UserRoleGroupList,omitempty"`

	Passport *Passport `xml:"Passport,omitempty"`

	TelePhone string `xml:"TelePhone,omitempty"`

	IDCardCode string `xml:"IDCardCode,omitempty"`

	Sex string `xml:"Sex,omitempty"`

	Age int32 `xml:"Age,omitempty"`

	Birthday f_datetime `xml:"Birthday,omitempty"`

	Education string `xml:"Education,omitempty"`

	Nation string `xml:"Nation,omitempty"`

	Height string `xml:"Height,omitempty"`

	BirthPlace string `xml:"BirthPlace,omitempty"`

	GraduateSchool string `xml:"GraduateSchool,omitempty"`

	MainOU int32 `xml:"MainOU,omitempty"`

	Address string `xml:"Address,omitempty"`

	DeptList *ArrayOfDeptInfoEx2 `xml:"DeptList,omitempty"`

	JoinTime f_datetime `xml:"JoinTime,omitempty"`

	LeaveTime f_datetime `xml:"LeaveTime,omitempty"`

	LeaderUserID int32 `xml:"LeaderUserID,omitempty"`
}

type ArrayOfDeptInfoEx2 struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfDeptInfoEx2"`

	DeptInfoEx2 []*DeptInfoEx2 `xml:"DeptInfoEx2,omitempty"`
}

type DeptInfoEx2 struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ DeptInfoEx2"`

	DeptID string `xml:"DeptID,omitempty"`

	DeptName string `xml:"DeptName,omitempty"`

	Expand string `xml:"Expand,omitempty"`

	ParentDeptID string `xml:"ParentDeptID,omitempty"`

	OUID int32 `xml:"OUID,omitempty"`

	PositionID int32 `xml:"PositionID,omitempty"`

	PositionName string `xml:"PositionName,omitempty"`

	StationID int32 `xml:"StationID,omitempty"`

	StationName string `xml:"StationName,omitempty"`

	Remark string `xml:"Remark,omitempty"`

	Flag int32 `xml:"Flag,omitempty"`

	ChainID int32 `xml:"ChainID,omitempty"`

	LeaderUserID string `xml:"LeaderUserID,omitempty"`

	BrandID int32 `xml:"BrandID,omitempty"`
}

type ArrayOfManagerInfo struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfManagerInfo"`

	ManagerInfo []*ManagerInfo `xml:"ManagerInfo,omitempty"`
}

type ManagerInfo struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ManagerInfo"`

	Name string `xml:"Name,omitempty"`

	MebCardCode string `xml:"MebCardCode,omitempty"`

	EMail string `xml:"EMail,omitempty"`

	Mobile string `xml:"Mobile,omitempty"`

	UserID int32 `xml:"UserID,omitempty"`

	StationID int32 `xml:"StationID,omitempty"`

	StationName string `xml:"StationName,omitempty"`

	LeaderUserID int32 `xml:"LeaderUserID,omitempty"`
}

type RowAbilityDateSetOfListOfDept struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ RowAbilityDateSetOfListOfDept"`

	DataSet *ArrayOfDept `xml:"DataSet,omitempty"`

	RowCount int32 `xml:"RowCount,omitempty"`

	PageCount int32 `xml:"PageCount,omitempty"`

	CurrentPage int32 `xml:"CurrentPage,omitempty"`

	PageRowCount int32 `xml:"PageRowCount,omitempty"`
}

type ArrayOfDept struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfDept"`

	Dept []*Dept `xml:"Dept,omitempty"`
}

type Dept struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ Dept"`

	DeptID string `xml:"DeptID,omitempty"`

	DeptName string `xml:"DeptName,omitempty"`

	ParentDeptID int32 `xml:"ParentDeptID,omitempty"`

	OUID int32 `xml:"OUID,omitempty"`

	Remark string `xml:"Remark,omitempty"`

	Flag int32 `xml:"Flag,omitempty"`

	ChainID int32 `xml:"ChainID,omitempty"`

	LeaderUserID string `xml:"LeaderUserID,omitempty"`

	BrandID int32 `xml:"BrandID,omitempty"`

	Expand string `xml:"Expand,omitempty"`

	SaleManagerCode int64 `xml:"SaleManagerCode,omitempty"`
}

type ArrayOfInn struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfInn"`

	Inn []*Inn `xml:"Inn,omitempty"`
}

type Inn struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ Inn"`

	ProjectID string `xml:"ProjectID,omitempty"`

	CityID int32 `xml:"CityID,omitempty"`

	ProjectName string `xml:"ProjectName,omitempty"`

	AreaName string `xml:"AreaName,omitempty"`
}

type ArrayOfStation struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfStation"`

	Station []*Station `xml:"Station,omitempty"`
}

type Station struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ Station"`

	StationID int32 `xml:"StationID,omitempty"`

	StationName string `xml:"StationName,omitempty"`

	Remark string `xml:"Remark,omitempty"`

	Flag int32 `xml:"Flag,omitempty"`

	LstStationRoles *ArrayOfStationRoles `xml:"lstStationRoles,omitempty"`
}

type ArrayOfStationRoles struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfStationRoles"`

	StationRoles []*StationRoles `xml:"StationRoles,omitempty"`
}

type StationRoles struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ StationRoles"`

	ID int32 `xml:"ID,omitempty"`

	StationID int32 `xml:"StationID,omitempty"`

	UserGroupID int32 `xml:"UserGroupID,omitempty"`

	UserGroupName string `xml:"UserGroupName,omitempty"`

	Flag int32 `xml:"Flag,omitempty"`
}

type ArrayOfCheckRoles struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfCheckRoles"`

	CheckRoles []*CheckRoles `xml:"CheckRoles,omitempty"`
}

type CheckRoles struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ CheckRoles"`

	SysRoleID int32 `xml:"SysRoleID,omitempty"`

	SystemID int32 `xml:"SystemID,omitempty"`

	Name string `xml:"Name,omitempty"`

	Remark string `xml:"Remark,omitempty"`
}

type ArrayOfCheckGroupOnRoles struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfCheckGroupOnRoles"`

	CheckGroupOnRoles []*CheckGroupOnRoles `xml:"CheckGroupOnRoles,omitempty"`
}

type CheckGroupOnRoles struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ CheckGroupOnRoles"`

	GroupID int32 `xml:"GroupID,omitempty"`

	RolesID int32 `xml:"RolesID,omitempty"`

	SystemID int32 `xml:"SystemID,omitempty"`

	IsConflict int32 `xml:"IsConflict,omitempty"`

	GroupName string `xml:"GroupName,omitempty"`

	RolesName string `xml:"RolesName,omitempty"`

	RolesRemark string `xml:"RolesRemark,omitempty"`

	SystemName string `xml:"SystemName,omitempty"`

	ConflictRemark string `xml:"ConflictRemark,omitempty"`

	Remark string `xml:"Remark,omitempty"`
}

type ArrayOfCheckUserOnRoles struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfCheckUserOnRoles"`

	CheckUserOnRoles []*CheckUserOnRoles `xml:"CheckUserOnRoles,omitempty"`
}

type CheckUserOnRoles struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ CheckUserOnRoles"`

	ID int32 `xml:"ID,omitempty"`

	SystemID int32 `xml:"SystemID,omitempty"`

	SystemName string `xml:"SystemName,omitempty"`

	MebCode int64 `xml:"MebCode,omitempty"`

	MebName string `xml:"MebName,omitempty"`

	DeptID string `xml:"DeptID,omitempty"`

	DeptName string `xml:"DeptName,omitempty"`

	GroupID int32 `xml:"GroupID,omitempty"`

	GroupName string `xml:"GroupName,omitempty"`

	RolesID int32 `xml:"RolesID,omitempty"`

	RolesName string `xml:"RolesName,omitempty"`

	RolesRemark string `xml:"RolesRemark,omitempty"`

	IsConflict int32 `xml:"IsConflict,omitempty"`

	ConflictRemark string `xml:"ConflictRemark,omitempty"`
}

type ArrayOfPeople struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfPeople"`

	People []*People `xml:"People,omitempty"`
}

type People struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ People"`

	MebCode int64 `xml:"MebCode,omitempty"`

	UserCode string `xml:"UserCode,omitempty"`

	Name string `xml:"Name,omitempty"`

	Email string `xml:"Email,omitempty"`

	Mobile string `xml:"Mobile,omitempty"`

	TelePhone string `xml:"TelePhone,omitempty"`

	IDCardCode string `xml:"IDCardCode,omitempty"`

	Sex int32 `xml:"Sex,omitempty"`

	Age int32 `xml:"Age,omitempty"`

	Birthday f_datetime `xml:"Birthday,omitempty"`

	Education string `xml:"Education,omitempty"`

	Nation string `xml:"Nation,omitempty"`

	Height string `xml:"Height,omitempty"`

	BirthPlace string `xml:"BirthPlace,omitempty"`

	GraduateSchool string `xml:"GraduateSchool,omitempty"`

	Address string `xml:"Address,omitempty"`

	Type int32 `xml:"Type,omitempty"`

	MainOU int32 `xml:"MainOU,omitempty"`

	Flag int32 `xml:"Flag,omitempty"`

	PeopleDetails *ArrayOfPeopleDetails `xml:"people_Details,omitempty"`

	RoleList *ArrayOfString `xml:"RoleList,omitempty"`

	JoinTime f_datetime `xml:"JoinTime,omitempty"`

	LeaveTime f_datetime `xml:"LeaveTime,omitempty"`

	UserID int32 `xml:"UserID,omitempty"`

	LeaderUserID int32 `xml:"LeaderUserID,omitempty"`
}

type ArrayOfPeopleDetails struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfPeople_Details"`

	PeopleDetails []*PeopleDetails `xml:"People_Details,omitempty"`
}

type PeopleDetails struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ People_Details"`

	DetailID int32 `xml:"DetailID,omitempty"`

	MebCode int64 `xml:"MebCode,omitempty"`

	PositionID int32 `xml:"PositionID,omitempty"`

	PositionName string `xml:"PositionName,omitempty"`

	StationID int32 `xml:"StationID,omitempty"`

	StationName string `xml:"StationName,omitempty"`

	DeptID string `xml:"DeptID,omitempty"`

	DeptName string `xml:"DeptName,omitempty"`

	OUID int32 `xml:"OUID,omitempty"`

	OUName string `xml:"OUName,omitempty"`

	JoinDeptTime f_datetime `xml:"JoinDeptTime,omitempty"`
}

type ArrayOfString struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfString"`

	String []string `xml:"string,omitempty"`
}

type ArrayOfPeopleEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfPeopleEx"`

	PeopleEx []*PeopleEx `xml:"PeopleEx,omitempty"`
}

type PeopleEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ PeopleEx"`

	MebCode int64 `xml:"MebCode,omitempty"`

	UserCode string `xml:"UserCode,omitempty"`

	Name string `xml:"Name,omitempty"`

	Email string `xml:"Email,omitempty"`

	Mobile string `xml:"Mobile,omitempty"`

	TelePhone string `xml:"TelePhone,omitempty"`

	IDCardCode string `xml:"IDCardCode,omitempty"`

	Sex int32 `xml:"Sex,omitempty"`

	Age int32 `xml:"Age,omitempty"`

	Birthday f_datetime `xml:"Birthday,omitempty"`

	Education string `xml:"Education,omitempty"`

	Nation string `xml:"Nation,omitempty"`

	Height string `xml:"Height,omitempty"`

	BirthPlace string `xml:"BirthPlace,omitempty"`

	GraduateSchool string `xml:"GraduateSchool,omitempty"`

	Address string `xml:"Address,omitempty"`

	Type int32 `xml:"Type,omitempty"`

	MainOU int32 `xml:"MainOU,omitempty"`

	Flag int32 `xml:"Flag,omitempty"`

	PeopleDetailsEx *ArrayOfPeopleDetailsEx `xml:"people_DetailsEx,omitempty"`

	RoleList *ArrayOfString `xml:"RoleList,omitempty"`

	JoinTime f_datetime `xml:"JoinTime,omitempty"`

	LeaveTime f_datetime `xml:"LeaveTime,omitempty"`

	UserID int32 `xml:"UserID,omitempty"`

	LeaderUserID int32 `xml:"LeaderUserID,omitempty"`
}

type ArrayOfPeopleDetailsEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfPeople_DetailsEx"`

	PeopleDetailsEx []*PeopleDetailsEx `xml:"People_DetailsEx,omitempty"`
}

type PeopleDetailsEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ People_DetailsEx"`

	DetailID int32 `xml:"DetailID,omitempty"`

	MebCode int64 `xml:"MebCode,omitempty"`

	PositionID int32 `xml:"PositionID,omitempty"`

	PositionName string `xml:"PositionName,omitempty"`

	StationID int32 `xml:"StationID,omitempty"`

	StationName string `xml:"StationName,omitempty"`

	DeptID string `xml:"DeptID,omitempty"`

	DeptName string `xml:"DeptName,omitempty"`

	OUID int32 `xml:"OUID,omitempty"`

	OUName string `xml:"OUName,omitempty"`

	JoinDeptTime f_datetime `xml:"JoinDeptTime,omitempty"`

	GroupID int32 `xml:"GroupID,omitempty"`
}

type ArrayOfSaleManager struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfSaleManager"`

	SaleManager []*SaleManager `xml:"SaleManager,omitempty"`
}

type SaleManager struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ SaleManager"`

	MebCode int64 `xml:"MebCode,omitempty"`

	Name string `xml:"Name,omitempty"`

	Email string `xml:"Email,omitempty"`

	Mobile string `xml:"Mobile,omitempty"`

	JoinTime f_datetime `xml:"JoinTime,omitempty"`

	LeaveTime f_datetime `xml:"LeaveTime,omitempty"`

	UserID string `xml:"UserID,omitempty"`

	DeptID string `xml:"DeptID,omitempty"`

	DeptName string `xml:"DeptName,omitempty"`

	ParentDeptID string `xml:"ParentDeptID,omitempty"`

	Expand string `xml:"Expand,omitempty"`

	ChainID int32 `xml:"ChainID,omitempty"`

	BrandID int32 `xml:"BrandID,omitempty"`

	GroupID int32 `xml:"GroupID,omitempty"`
}

type ArrayOfCCUser struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ArrayOfCCUser"`

	CCUser []*CCUser `xml:"CCUser,omitempty"`
}

type CCUser struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ CCUser"`

	UserID int32 `xml:"UserID,omitempty"`

	UserCode string `xml:"UserCode,omitempty"`

	Name string `xml:"Name,omitempty"`

	Title string `xml:"Title,omitempty"`

	Type int32 `xml:"Type,omitempty"`

	MebCardCode string `xml:"MebCardCode,omitempty"`

	Flag int32 `xml:"Flag,omitempty"`

	Mobile string `xml:"Mobile,omitempty"`

	EMail string `xml:"EMail,omitempty"`

	UserRoleGroupList string `xml:"UserRoleGroupList,omitempty"`

	InnList string `xml:"InnList,omitempty"`
}

type AddInnUser struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ AddInnUser"`

	MebCardCode string `xml:"MebCardCode,omitempty"`

	Name string `xml:"Name,omitempty"`

	Title string `xml:"Title,omitempty"`

	Mobile string `xml:"Mobile,omitempty"`

	Password string `xml:"Password,omitempty"`

	EMail string `xml:"EMail,omitempty"`

	UserRoleGroupList string `xml:"UserRoleGroupList,omitempty"`

	InnList string `xml:"InnList,omitempty"`
}

type ModifyInnUser struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ModifyInnUser"`

	MebCardCode string `xml:"MebCardCode,omitempty"`

	Title string `xml:"Title,omitempty"`

	Mobile string `xml:"Mobile,omitempty"`

	EMail string `xml:"EMail,omitempty"`

	UserRoleGroupList string `xml:"UserRoleGroupList,omitempty"`
}

type ModifyInnUserEx struct {
	XMLName xml.Name `xml:"http://UserManageCenter.7daysinn.cn/ ModifyInnUserEx"`

	MebCardCode string `xml:"MebCardCode,omitempty"`

	Title string `xml:"Title,omitempty"`

	Mobile string `xml:"Mobile,omitempty"`

	EMail string `xml:"EMail,omitempty"`
}

type UserManageCenterServiceSoap struct {
	client *SOAPClient
}

func NewUserManageCenterServiceSoap(url string, tls bool) *UserManageCenterServiceSoap {
	if url == "" {
		url = "http://10.100.113.38:1101/UserManageCenterService.asmx"
	}
	client := NewSOAPClient(url, tls)

	return &UserManageCenterServiceSoap{
		client: client,
	}
}

func (service *UserManageCenterServiceSoap) SetHeader(header interface{}) {
	service.client.SetHeader(header)
}

/* 验证用户的合法性  */
func (service *UserManageCenterServiceSoap) AuthenticateUser(request *AuthenticateUser) (*AuthenticateUserResponse, error) {
	response := new(AuthenticateUserResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/AuthenticateUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 验证用户的合法性  */
func (service *UserManageCenterServiceSoap) AuthenticateUserEx(request *AuthenticateUserEx) (*AuthenticateUserExResponse, error) {
	response := new(AuthenticateUserExResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/AuthenticateUserEx", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 验证用户合法性(含有动态口令卡)  */
func (service *UserManageCenterServiceSoap) AuthenticateKeyCardUser(request *AuthenticateKeyCardUser) (*AuthenticateKeyCardUserResponse, error) {
	response := new(AuthenticateKeyCardUserResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/AuthenticateKeyCardUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 通过用户姓名获取用户  */
func (service *UserManageCenterServiceSoap) GetUserByName(request *GetUserByName) (*GetUserByNameResponse, error) {
	response := new(GetUserByNameResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetUserByName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 通过用户帐号/会员卡号获取有效的用户  */
func (service *UserManageCenterServiceSoap) GetUserMessageByUserCode(request *GetUserMessageByUserCode) (*GetUserMessageByUserCodeResponse, error) {
	response := new(GetUserMessageByUserCodeResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetUserMessageByUserCode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 验证用户的合法性(同时返回用户的信息) */
func (service *UserManageCenterServiceSoap) AuthenticateUserEx2(request *AuthenticateUserEx2) (*AuthenticateUserEx2Response, error) {
	response := new(AuthenticateUserEx2Response)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/AuthenticateUserEx2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 验证用户的合法性(同时返回用户的信息，包含岗位组编码：GroupID) */
func (service *UserManageCenterServiceSoap) AuthenticateUserEx3(request *AuthenticateUserEx3) (*AuthenticateUserEx3Response, error) {
	response := new(AuthenticateUserEx3Response)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/AuthenticateUserEx3", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 通过用户帐号/会员卡号获取可能是无效的或有效的用户  */
func (service *UserManageCenterServiceSoap) GetUserByUserCode(request *GetUserByUserCode) (*GetUserByUserCodeResponse, error) {
	response := new(GetUserByUserCodeResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetUserByUserCode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 通过用户帐号/会员卡号获取可能是无效的或有效的用户  */
func (service *UserManageCenterServiceSoap) GetUserByUserCodeEx(request *GetUserByUserCodeEx) (*GetUserByUserCodeExResponse, error) {
	response := new(GetUserByUserCodeExResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetUserByUserCodeEx", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 通过用户ID获取可能是无效的或有效的用户  */
func (service *UserManageCenterServiceSoap) GetUserByUserIDEx(request *GetUserByUserIDEx) (*GetUserByUserIDExResponse, error) {
	response := new(GetUserByUserIDExResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetUserByUserIDEx", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 通过用户部门查询用户，nGotoPage或nPageSize为0时，查询该部门下的所有用户 */
func (service *UserManageCenterServiceSoap) GetUserByDept(request *GetUserByDept) (*GetUserByDeptResponse, error) {
	response := new(GetUserByDeptResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetUserByDept", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 根据用户类型查询用户(0：程序；1：分店；2：客服；3：集团；4：其它) */
func (service *UserManageCenterServiceSoap) GetUserByType(request *GetUserByType) (*GetUserByTypeResponse, error) {
	response := new(GetUserByTypeResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetUserByType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 取即是分店用户又在总部的用户（申报系统专用）) */
func (service *UserManageCenterServiceSoap) GetGroupChainUser(request *GetGroupChainUser) (*GetGroupChainUserResponse, error) {
	response := new(GetGroupChainUserResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetGroupChainUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 通过岗位查询用户，nGotoPage或nPageSize为0时，查询该岗位下的所有用户 */
func (service *UserManageCenterServiceSoap) GetUserByStation(request *GetUserByStation) (*GetUserByStationResponse, error) {
	response := new(GetUserByStationResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetUserByStation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 查询用户 */
func (service *UserManageCenterServiceSoap) QueryUser(request *QueryUser) (*QueryUserResponse, error) {
	response := new(QueryUserResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/QueryUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 查询用户 */
func (service *UserManageCenterServiceSoap) QueryUserEx(request *QueryUserEx) (*QueryUserExResponse, error) {
	response := new(QueryUserExResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/QueryUserEx", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 查询用户 */
func (service *UserManageCenterServiceSoap) QueryUserEx2(request *QueryUserEx2) (*QueryUserEx2Response, error) {
	response := new(QueryUserEx2Response)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/QueryUserEx2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 根据chainID查询店长的UserCode和Name */
func (service *UserManageCenterServiceSoap) GetManagerMessage(request *GetManagerMessage) (*GetManagerMessageResponse, error) {
	response := new(GetManagerMessageResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetManagerMessage", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 根据分店ID获取店长信息 */
func (service *UserManageCenterServiceSoap) GetUserInfoByChainIDAndStationID(request *GetUserInfoByChainIDAndStationID) (*GetUserInfoByChainIDAndStationIDResponse, error) {
	response := new(GetUserInfoByChainIDAndStationIDResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetUserInfoByChainIDAndStationID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 通过用户ID获取用户 */
func (service *UserManageCenterServiceSoap) GetUserByUserID(request *GetUserByUserID) (*GetUserByUserIDResponse, error) {
	response := new(GetUserByUserIDResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetUserByUserID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 查询所有部门，nGotoPage为0时，查询所有  */
func (service *UserManageCenterServiceSoap) QueryAllDept(request *QueryAllDept) (*QueryAllDeptResponse, error) {
	response := new(QueryAllDeptResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/QueryAllDept", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 查询所有部门，nGotoPage为0时，查询所有  */
func (service *UserManageCenterServiceSoap) QueryAllDeptEx(request *QueryAllDeptEx) (*QueryAllDeptExResponse, error) {
	response := new(QueryAllDeptExResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/QueryAllDeptEx", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 查询所有没有子部门的部门 */
func (service *UserManageCenterServiceSoap) GetAllSealedDept(request *GetAllSealedDept) (*GetAllSealedDeptResponse, error) {
	response := new(GetAllSealedDeptResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetAllSealedDept", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 查询所有子部门 */
func (service *UserManageCenterServiceSoap) GetAllChildDeptByDeptID(request *GetAllChildDeptByDeptID) (*GetAllChildDeptByDeptIDResponse, error) {
	response := new(GetAllChildDeptByDeptIDResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetAllChildDeptByDeptID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 查询所有子部门 */
func (service *UserManageCenterServiceSoap) GetAllChildDeptEx(request *GetAllChildDeptEx) (*GetAllChildDeptExResponse, error) {
	response := new(GetAllChildDeptExResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetAllChildDeptEx", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 按部门编码查询部门 */
func (service *UserManageCenterServiceSoap) GetDeptByCode(request *GetDeptByCode) (*GetDeptByCodeResponse, error) {
	response := new(GetDeptByCodeResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetDeptByCode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 从证照系统中查询所有分店 */
func (service *UserManageCenterServiceSoap) GetAllInn(request *GetAllInn) (*GetAllInnResponse, error) {
	response := new(GetAllInnResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetAllInn", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 查询部门信息 */
func (service *UserManageCenterServiceSoap) GetDeptByQueryString(request *GetDeptByQueryString) (*GetDeptByQueryStringResponse, error) {
	response := new(GetDeptByQueryStringResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetDeptByQueryString", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 根据Expand查询部门信息 */
func (service *UserManageCenterServiceSoap) GetDeptByExpand(request *GetDeptByExpand) (*GetDeptByExpandResponse, error) {
	response := new(GetDeptByExpandResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetDeptByExpand", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 获取岗位数据 */
func (service *UserManageCenterServiceSoap) GetAllStation(request *GetAllStation) (*GetAllStationResponse, error) {
	response := new(GetAllStationResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetAllStation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 添加系统所有权限项，一次只能添加一个系统的权限项，并且会删除该系统的历史权限项 */
func (service *UserManageCenterServiceSoap) AddCheckRoles(request *AddCheckRoles) (*AddCheckRolesResponse, error) {
	response := new(AddCheckRolesResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/AddCheckRoles", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 添加权限组权限，一次只能添加一个系统的权限，并且会删除该系统的历史权限记录 */
func (service *UserManageCenterServiceSoap) AddCheckGroupRoles(request *AddCheckGroupRoles) (*AddCheckGroupRolesResponse, error) {
	response := new(AddCheckGroupRolesResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/AddCheckGroupRoles", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 添加用户权限，一次只能添加一个系统的权限，并且会删除该系统的历史权限记录 */
func (service *UserManageCenterServiceSoap) AddCheckUserRoles(request *AddCheckUserRoles) (*AddCheckUserRolesResponse, error) {
	response := new(AddCheckUserRolesResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/AddCheckUserRoles", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 修改用户信息 */
func (service *UserManageCenterServiceSoap) UpdateUser(request *UpdateUser) (*UpdateUserResponse, error) {
	response := new(UpdateUserResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/UpdateUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 根据分店ID查找有效用户 */
func (service *UserManageCenterServiceSoap) GetPeopleByChainID(request *GetPeopleByChainID) (*GetPeopleByChainIDResponse, error) {
	response := new(GetPeopleByChainIDResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetPeopleByChainID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 根据分店ID查找有效用户（包含GroupID） */
func (service *UserManageCenterServiceSoap) GetPeopleByChainIDEx(request *GetPeopleByChainIDEx) (*GetPeopleByChainIDExResponse, error) {
	response := new(GetPeopleByChainIDExResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetPeopleByChainIDEx", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 根据卡号查找分店编码列表字符串 */
func (service *UserManageCenterServiceSoap) GetChainIDByMebCardCode(request *GetChainIDByMebCardCode) (*GetChainIDByMebCardCodeResponse, error) {
	response := new(GetChainIDByMebCardCodeResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetChainIDByMebCardCode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 根据卡号查找分店编码列表字符串 */
func (service *UserManageCenterServiceSoap) GetManageSalerByMebCardCode(request *GetManageSalerByMebCardCode) (*GetManageSalerByMebCardCodeResponse, error) {
	response := new(GetManageSalerByMebCardCodeResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetManageSalerByMebCardCode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 根据分店编码更新销售经理卡号 */
func (service *UserManageCenterServiceSoap) UpdateSaleManagerCodeByChainID(request *UpdateSaleManagerCodeByChainID) (*UpdateSaleManagerCodeByChainIDResponse, error) {
	response := new(UpdateSaleManagerCodeByChainIDResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/UpdateSaleManagerCodeByChainID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* PMS专用，违者砍头-查询分店的所有用户，包括有效跟无效的 */
func (service *UserManageCenterServiceSoap) GetAllChainUsers(request *GetAllChainUsers) (*GetAllChainUsersResponse, error) {
	response := new(GetAllChainUsersResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/GetAllChainUsers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* PMS专用，违者砍头-新增用户 */
func (service *UserManageCenterServiceSoap) AddUser(request *AddUser) (*AddUserResponse, error) {
	response := new(AddUserResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/AddUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* PMS专用，违者砍头-新增用户 */
func (service *UserManageCenterServiceSoap) AddUserEx(request *AddUserEx) (*AddUserExResponse, error) {
	response := new(AddUserExResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/AddUserEx", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* PMS专用，违者砍头-新增用户 */
func (service *UserManageCenterServiceSoap) AddUserEx2(request *AddUserEx2) (*AddUserEx2Response, error) {
	response := new(AddUserEx2Response)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/AddUserEx2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* PMS专用，违者砍头-修改用户 */
func (service *UserManageCenterServiceSoap) ModifyUser(request *ModifyUser) (*ModifyUserResponse, error) {
	response := new(ModifyUserResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/ModifyUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 修改用户（根据会员卡号） */
func (service *UserManageCenterServiceSoap) ModifyUserEx(request *ModifyUserEx) (*ModifyUserExResponse, error) {
	response := new(ModifyUserExResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/ModifyUserEx", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* PMS专用，违者砍头-禁用用户 */
func (service *UserManageCenterServiceSoap) DisableUser(request *DisableUser) (*DisableUserResponse, error) {
	response := new(DisableUserResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/DisableUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* PMS专用，违者砍头-禁用用户 */
func (service *UserManageCenterServiceSoap) EnableUser(request *EnableUser) (*EnableUserResponse, error) {
	response := new(EnableUserResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/EnableUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* PMS专用，违者砍头-判断是否允许添加该卡号的用户 */
func (service *UserManageCenterServiceSoap) CheckMebCardIsExist(request *CheckMebCardIsExist) (*CheckMebCardIsExistResponse, error) {
	response := new(CheckMebCardIsExistResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/CheckMebCardIsExist", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* PMS专用，违者砍头-修改用户密码 */
func (service *UserManageCenterServiceSoap) ModifyPassword(request *ModifyPassword) (*ModifyPasswordResponse, error) {
	response := new(ModifyPasswordResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/ModifyPassword", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 修改用户密码 */
func (service *UserManageCenterServiceSoap) ModifyPasswordEx(request *ModifyPasswordEx) (*ModifyPasswordExResponse, error) {
	response := new(ModifyPasswordExResponse)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/ModifyPasswordEx", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* 修改用户密码 */
func (service *UserManageCenterServiceSoap) ModifyPasswordEx2(request *ModifyPasswordEx2) (*ModifyPasswordEx2Response, error) {
	response := new(ModifyPasswordEx2Response)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/ModifyPasswordEx2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* PMS专用，违者砍头-修改用户密码 */
func (service *UserManageCenterServiceSoap) ModifyPasswordEx3(request *ModifyPasswordEx3) (*ModifyPasswordEx3Response, error) {
	response := new(ModifyPasswordEx3Response)
	err := service.client.Call("http://UserManageCenter.7daysinn.cn/ModifyPasswordEx3", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *SOAPHeader
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url    string
	tls    bool
	auth   *BasicAuth
	header interface{}
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool) *SOAPClient {
	return &SOAPClient{
		url: url,
		tls: tls,
	}
}

func (s *SOAPClient) SetHeader(header interface{}) {
	s.header = header
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{}

	if s.header != nil {
		envelope.Header = &SOAPHeader{Header: s.header}
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	//log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	//log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
