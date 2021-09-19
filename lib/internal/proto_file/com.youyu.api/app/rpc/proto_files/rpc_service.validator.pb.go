// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc_service.proto

package proto_files

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Null) Validate() error {
	return nil
}
func (this *Article_Response) Validate() error {
	for _, item := range this.Articles {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Articles", err)
			}
		}
	}
	return nil
}
func (this *GetArticleRequest) Validate() error {
	return nil
}
func (this *Article) Validate() error {
	return nil
}
func (this *ArticleOptions) Validate() error {
	return nil
}
func (this *AdvertisementOptions) Validate() error {
	return nil
}
func (this *ArticleStatistics) Validate() error {
	return nil
}
func (this *ArticleLinkTab) Validate() error {
	return nil
}
func (this *BaseData) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AdvertisementRequest) Validate() error {
	return nil
}
func (this *AdvertisementResponse) Validate() error {
	for _, item := range this.AdvertisementList {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AdvertisementList", err)
			}
		}
	}
	return nil
}
func (this *Advertisement) Validate() error {
	return nil
}
func (this *Tag) Validate() error {
	return nil
}
func (this *UserAuth) Validate() error {
	return nil
}
func (this *UserCheckPhone) Validate() error {
	if this.Ua != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ua); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ua", err)
		}
	}
	return nil
}

var _regex_UserCheckEmail_Email = regexp.MustCompile(`^[0-9A-Za-z][\.-_0-9A-Za-z]*@[0-9A-Za-z]+(?:\.[0-9A-Za-z]+)+$`)

func (this *UserCheckEmail) Validate() error {
	if !_regex_UserCheckEmail_Email.MatchString(this.Email) {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must be a string conforming to regex "^[0-9A-Za-z][\\.-_0-9A-Za-z]*@[0-9A-Za-z]+(?:\\.[0-9A-Za-z]+)+$"`, this.Email))
	}
	if this.Ua != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ua); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ua", err)
		}
	}
	return nil
}
func (this *UserCheckWechat) Validate() error {
	if this.Openid == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Openid", fmt.Errorf(`value '%v' must not be an empty string`, this.Openid))
	}
	if this.Code == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Code", fmt.Errorf(`value '%v' must not be an empty string`, this.Code))
	}
	if this.Ua != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ua); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ua", err)
		}
	}
	return nil
}
func (this *UserSign) Validate() error {
	if _, ok := LoginAndSignType_name[int32(this.SignType)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("SignType", fmt.Errorf(`value '%v' must be a valid LoginAndSignType field`, this.SignType))
	}
	if this.WechatData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WechatData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WechatData", err)
		}
	}
	return nil
}
func (this *UserLogin) Validate() error {
	if !(this.Save > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Save", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Save))
	}
	if !(this.Save < 3) {
		return github_com_mwitkow_go_proto_validators.FieldError("Save", fmt.Errorf(`value '%v' must be less than '3'`, this.Save))
	}
	if _, ok := LoginAndSignType_name[int32(this.LoginType)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("LoginType", fmt.Errorf(`value '%v' must be a valid LoginAndSignType field`, this.LoginType))
	}
	if this.WechatData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WechatData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WechatData", err)
		}
	}
	return nil
}
func (this *UserInfoShow) Validate() error {
	return nil
}
func (this *UserInfoOtherShow) Validate() error {
	return nil
}
func (this *UserInfoSet) Validate() error {
	if !(this.Sex > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Sex", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Sex))
	}
	if !(this.Sex < 3) {
		return github_com_mwitkow_go_proto_validators.FieldError("Sex", fmt.Errorf(`value '%v' must be less than '3'`, this.Sex))
	}
	if !(this.Age > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Age", fmt.Errorf(`value '%v' must be greater than '0'`, this.Age))
	}
	if !(this.Age < 160) {
		return github_com_mwitkow_go_proto_validators.FieldError("Age", fmt.Errorf(`value '%v' must be less than '160'`, this.Age))
	}
	if !(len(this.UserNickName) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserNickName", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.UserNickName))
	}
	if !(len(this.UserNickName) < 11) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserNickName", fmt.Errorf(`value '%v' must have a length smaller than '11'`, this.UserNickName))
	}
	if !(len(this.Explain) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Explain", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Explain))
	}
	if !(len(this.Explain) < 500) {
		return github_com_mwitkow_go_proto_validators.FieldError("Explain", fmt.Errorf(`value '%v' must have a length smaller than '500'`, this.Explain))
	}
	return nil
}
func (this *UserHeadPortraitSet) Validate() error {
	return nil
}
func (this *WechatUserinfo) Validate() error {
	return nil
}
func (this *PhoneEmailLoginInfo) Validate() error {
	return nil
}
func (this *CommentMasterShow) Validate() error {
	if _, ok := CommentType_name[int32(this.Type)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("Type", fmt.Errorf(`value '%v' must be a valid CommentType field`, this.Type))
	}
	if this.Text == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Text", fmt.Errorf(`value '%v' must not be an empty string`, this.Text))
	}
	if this.CreateTime == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CreateTime", fmt.Errorf(`value '%v' must not be an empty string`, this.CreateTime))
	}
	for _, item := range this.SlaveComment {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SlaveComment", err)
			}
		}
	}
	return nil
}
func (this *CommentSlave) Validate() error {
	if _, ok := CommentType_name[int32(this.Type)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("Type", fmt.Errorf(`value '%v' must be a valid CommentType field`, this.Type))
	}
	if this.Text == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Text", fmt.Errorf(`value '%v' must not be an empty string`, this.Text))
	}
	return nil
}
func (this *CommentShow) Validate() error {
	for _, item := range this.Master {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Master", err)
			}
		}
	}
	return nil
}
func (this *UpdateCommentOption) Validate() error {
	if _, ok := CommentType_name[int32(this.Type)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("Type", fmt.Errorf(`value '%v' must be a valid CommentType field`, this.Type))
	}
	if _, ok := CommentOptions_name[int32(this.Options)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("Options", fmt.Errorf(`value '%v' must be a valid CommentOptions field`, this.Options))
	}
	return nil
}
