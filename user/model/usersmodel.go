package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	usersFieldNames          = builderx.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUsersIdPrefix            = "cache:users:id:"
	cacheUsersEmailPrefix         = "cache:users:email:"
	cacheUsersPhonePrefix         = "cache:users:phone:"
	cacheUsersWeixinOpenidPrefix  = "cache:users:weixinOpenid:"
	cacheUsersWeixinUnionidPrefix = "cache:users:weixinUnionid:"
)

type (
	UsersModel interface {
		Insert(data Users) (sql.Result, error)
		FindOne(id int64) (*Users, error)
		FindOneByEmail(email sql.NullString) (*Users, error)
		FindOneByPhone(phone sql.NullString) (*Users, error)
		FindOneByWeixinOpenid(weixinOpenid sql.NullString) (*Users, error)
		FindOneByWeixinUnionid(weixinUnionid sql.NullString) (*Users, error)
		Update(data Users) error
		Delete(id int64) error
	}

	defaultUsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		Id                int64          `db:"id"`
		Name              string         `db:"name"`
		Phone             sql.NullString `db:"phone"`
		Email             sql.NullString `db:"email"`
		EmailVerifiedAt   sql.NullTime   `db:"email_verified_at"`
		Password          sql.NullString `db:"password"`
		WeixinOpenid      sql.NullString `db:"weixin_openid"`
		WeixinUnionid     sql.NullString `db:"weixin_unionid"`
		RememberToken     sql.NullString `db:"remember_token"`
		CreatedAt         sql.NullTime   `db:"created_at"`
		UpdatedAt         sql.NullTime   `db:"updated_at"`
		Avatar            sql.NullString `db:"avatar"`
		Introduction      sql.NullString `db:"introduction"`
		NotificationCount int64          `db:"notification_count"`
		LastActivedAt     sql.NullTime   `db:"last_actived_at"`
		RegistrationId    sql.NullString `db:"registration_id"`
	}
)

func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf) UsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`users`",
	}
}

func (m *defaultUsersModel) Insert(data Users) (sql.Result, error) {
	usersEmailKey := fmt.Sprintf("%s%v", cacheUsersEmailPrefix, data.Email)
	usersPhoneKey := fmt.Sprintf("%s%v", cacheUsersPhonePrefix, data.Phone)
	usersWeixinOpenidKey := fmt.Sprintf("%s%v", cacheUsersWeixinOpenidPrefix, data.WeixinOpenid)
	usersWeixinUnionidKey := fmt.Sprintf("%s%v", cacheUsersWeixinUnionidPrefix, data.WeixinUnionid)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, usersRowsExpectAutoSet)
		return conn.Exec(query, data.Name, data.Phone, data.Email, data.EmailVerifiedAt, data.Password, data.WeixinOpenid, data.WeixinUnionid, data.RememberToken, data.CreatedAt, data.UpdatedAt, data.Avatar, data.Introduction, data.NotificationCount, data.LastActivedAt, data.RegistrationId)
	}, usersEmailKey, usersPhoneKey, usersWeixinOpenidKey, usersWeixinUnionidKey)
	return ret, err
}

func (m *defaultUsersModel) FindOne(id int64) (*Users, error) {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, id)
	var resp Users
	err := m.QueryRow(&resp, usersIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByEmail(email sql.NullString) (*Users, error) {
	usersEmailKey := fmt.Sprintf("%s%v", cacheUsersEmailPrefix, email)
	var resp Users
	err := m.QueryRowIndex(&resp, usersEmailKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", usersRows, m.table)
		if err := conn.QueryRow(&resp, query, email); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByPhone(phone sql.NullString) (*Users, error) {
	usersPhoneKey := fmt.Sprintf("%s%v", cacheUsersPhonePrefix, phone)
	var resp Users
	err := m.QueryRowIndex(&resp, usersPhoneKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `phone` = ? limit 1", usersRows, m.table)
		if err := conn.QueryRow(&resp, query, phone); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByWeixinOpenid(weixinOpenid sql.NullString) (*Users, error) {
	usersWeixinOpenidKey := fmt.Sprintf("%s%v", cacheUsersWeixinOpenidPrefix, weixinOpenid)
	var resp Users
	err := m.QueryRowIndex(&resp, usersWeixinOpenidKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `weixin_openid` = ? limit 1", usersRows, m.table)
		if err := conn.QueryRow(&resp, query, weixinOpenid); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByWeixinUnionid(weixinUnionid sql.NullString) (*Users, error) {
	usersWeixinUnionidKey := fmt.Sprintf("%s%v", cacheUsersWeixinUnionidPrefix, weixinUnionid)
	var resp Users
	err := m.QueryRowIndex(&resp, usersWeixinUnionidKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `weixin_unionid` = ? limit 1", usersRows, m.table)
		if err := conn.QueryRow(&resp, query, weixinUnionid); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Update(data Users) error {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)
	usersEmailKey := fmt.Sprintf("%s%v", cacheUsersEmailPrefix, data.Email)
	usersPhoneKey := fmt.Sprintf("%s%v", cacheUsersPhonePrefix, data.Phone)
	usersWeixinOpenidKey := fmt.Sprintf("%s%v", cacheUsersWeixinOpenidPrefix, data.WeixinOpenid)
	usersWeixinUnionidKey := fmt.Sprintf("%s%v", cacheUsersWeixinUnionidPrefix, data.WeixinUnionid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.Phone, data.Email, data.EmailVerifiedAt, data.Password, data.WeixinOpenid, data.WeixinUnionid, data.RememberToken, data.CreatedAt, data.UpdatedAt, data.Avatar, data.Introduction, data.NotificationCount, data.LastActivedAt, data.RegistrationId, data.Id)
	}, usersPhoneKey, usersWeixinOpenidKey, usersWeixinUnionidKey, usersIdKey, usersEmailKey)
	return err
}

func (m *defaultUsersModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	usersEmailKey := fmt.Sprintf("%s%v", cacheUsersEmailPrefix, data.Email)
	usersPhoneKey := fmt.Sprintf("%s%v", cacheUsersPhonePrefix, data.Phone)
	usersWeixinOpenidKey := fmt.Sprintf("%s%v", cacheUsersWeixinOpenidPrefix, data.WeixinOpenid)
	usersWeixinUnionidKey := fmt.Sprintf("%s%v", cacheUsersWeixinUnionidPrefix, data.WeixinUnionid)
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, id)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, usersPhoneKey, usersWeixinOpenidKey, usersWeixinUnionidKey, usersIdKey, usersEmailKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUsersIdPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
	return conn.QueryRow(v, query, primary)
}
