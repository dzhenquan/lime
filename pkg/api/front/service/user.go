package service

import (
	"errors"
	"lime/pkg/api/admin/model"
	"lime/pkg/api/front/dao"
	"lime/pkg/api/front/domain/auth"
	"lime/pkg/api/front/dto"
	"lime/pkg/api/utils/log"
	"time"
)

type UserService struct{}

var UserDao = dao.UsersDao{}

func (U UserService) Login(gDto dto.LoginDto) (loginInfo dto.LoginInfoDto, err error) {
	var tokenErr = map[string]string{}

	u, userErr := UserDao.GetUserByUsernameAndPassword(gDto.Username, gDto.Password)
	if userErr != nil {
		return loginInfo, errors.New("no user")
	}
	ts, tErr := auth.NewToken().CreateToken(u.Id)
	if tErr != nil {
		tokenErr["token_error"] = tErr.Error()
		return loginInfo, tErr
	}
	saveErr := auth.CreateAuth(u.Id, ts)
	if saveErr != nil {
		return loginInfo, saveErr
	}
	return dto.LoginInfoDto{
		AccessToken:  ts.AccessToken,
		RefreshToken: ts.RefreshToken,
		AtExpires:    0,
		RtExpires:    0,
	}, nil
}

func (U UserService) Register(gDto dto.RegisterDto) (err error) {
	regUsername := UserDao.GetUserByField(gDto.Username,"username")
	if regUsername.Username != "" {
		return errors.New("已经有存在该账号")
	}
	regEmail := UserDao.GetUserByField(gDto.Email,"email")
	if regEmail.Email != "" {
		return errors.New("已经有存在该Email")
	}
	regMobile := UserDao.GetUserByField(gDto.Mobile,"mobile")
	if regMobile.Mobile != "" {
		return errors.New("已经有存在该手机号码")
	}
	salt, _ := auth.MakeSalt()
	password, _ := auth.HashPassword(gDto.Password, salt)
	Model := model.Users{
		Username:        gDto.Username,
		Mobile:          gDto.Mobile,
		Sex:             0,
		Password:        password,
		Salt:            salt,
		Faceicon:        "",
		//Wechat:          "",
		Email:           gDto.Email,
		Amount:          0,
		Coin:            0,
		Source:          1,
		Is_vip:          0,
		Channel_id:      1,
		Status:          0,
		CreatedAt:       time.Now(),
		Last_login_time: time.Now(),
	}
	c := UserDao.Create(&Model)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return nil
}

// UpdatePassword - update password only
func (U UserService) UpdatePassword(Id int, dto dto.UserEditPasswordDto) int64 {
	salt, _ := auth.MakeSalt()
	pwd, _ := auth.HashPassword(dto.Password, salt)
	u := UserDao.Get(Id)
	c := UserDao.Update(&u, map[string]interface{}{
		"password": pwd,
		"salt":     salt,
	})
	return c.RowsAffected
}

func (U UserService) Info(Id int) dto.InfoDto {
	u := UserDao.Get(Id)
	return dto.InfoDto{
		Id:       u.Id,
		Username: u.Username,
		Faceicon: u.Faceicon,
	}
}