package database

import (
	"time"

	. "vid/exceptions"
	. "vid/models"
	. "vid/utils"
)

type passDao struct{}

var PassDao = new(passDao)

const (
	col_pass_uid           = "uid"
	col_pass_encryptedPass = "encrypted_pass"
)

// db 内部使用 查询密码项
//
// @return `isExist`
func (p *passDao) queryPassRecord(uid int) (*PassRecord, bool) {
	var pass PassRecord
	nf := DB.Where(col_pass_uid+" = ?", uid).Find(&pass).RecordNotFound()
	if nf {
		return nil, false
	} else {
		return &pass, true
	}
}

// db 注册 插入用户和密码项
//
// @return `*user` `err`
//
// @error `UserExistException` `InsertUserException`
func (p *passDao) InsertUserPassRecord(username string, encryptedPass string) (*User, error) {

	if _, ok := UserDao.QueryUserByUserName(username); ok {
		return nil, UserExistException
	}

	tx := DB.Begin()

	// 先创建用户
	user := &User{
		Username:     username,
		AvatarUrl:    CmnUtil.GetDefaultAvatarUrl(),
		RegisterTime: time.Now(),
	}
	user.ToDB()
	tx.Create(user)

	// 用户是否被创建
	var queryUser User
	nf := tx.Where(col_user_username+" = ?", username).Find(&queryUser).RecordNotFound()
	if nf {
		tx.Rollback()
		return nil, InsertUserException
	}

	// 创建密码项
	tx.Create(&PassRecord{
		Uid:           queryUser.Uid,
		EncryptedPass: encryptedPass,
	})

	// 密码项是否被创建
	nf = tx.Where(col_pass_uid+" = ?", queryUser.Uid).Find(&PassRecord{}).RecordNotFound()
	if nf {
		tx.Rollback()
		return nil, InsertUserException
	}

	tx.Commit()

	queryUser.ToServer()
	return &queryUser, nil
}

// db 登录 查询密码项
//
// @return `*user` `*pass` `isExist`
func (p *passDao) QueryPassRecordByUsername(username string) (*User, *PassRecord, bool) {

	user, ok := UserDao.QueryUserByUserName(username)
	if !ok {
		return nil, nil, false
	}

	var pass PassRecord
	nf := DB.Where(col_pass_uid+" = ?", user.Uid).Find(&pass).RecordNotFound()
	if nf {
		return nil, nil, false
	} else {
		return user, &pass, true
	}
}

// db 登录 修改密码
//
// @return `uid` `err`
//
// @error `UserExistException` `ModifyPassException` `NotUpdateUserException`
func (p *passDao) UpdatePass(pass PassRecord) (int, error) {
	queryBefore, ok := p.queryPassRecord(pass.Uid)
	if !ok {
		return -1, UserExistException
	}
	DB.Model(&pass).Updates(map[string]interface{}{
		col_pass_encryptedPass: pass.EncryptedPass,
	})
	query, ok := p.queryPassRecord(pass.Uid)
	if !ok {
		return -1, ModifyPassException
	} else {
		if queryBefore.EncryptedPass == query.EncryptedPass {
			return -1, NotUpdateUserException
		} else {
			return pass.Uid, nil
		}
	}
}