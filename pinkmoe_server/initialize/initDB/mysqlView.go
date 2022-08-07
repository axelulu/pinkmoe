/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:03
 * @FilePath: /pinkmoe_server/initialize/initDB/mysqlView.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package initDB

import (
	"fmt"
	"reflect"
	"server/global"
	"server/model"
	"strings"

	"gorm.io/gorm/schema"

	"github.com/pkg/errors"
)

var ViewAuthorityMenuMysql = new(viewAuthorityMenuMysql)

type viewAuthorityMenuMysql struct{}

func (v *viewAuthorityMenuMysql) TableName() string {
	var entity model.XdMenu
	return entity.TableName()
}

type AuthorityMenus struct {
	AuthorityId string `gorm:"column:xd_authority_authority_id"`
	BaseMenuId  uint   `gorm:"column:xd_base_menu_id"`
}

func (a *AuthorityMenus) TableName() string {
	var entity model.XdAuthority
	types := reflect.TypeOf(entity)
	if s, o := types.FieldByName("XdBaseMenus"); o {
		m1 := schema.ParseTagSetting(s.Tag.Get("gorm"), ";")
		return m1["MANY2MANY"]
	}
	return ""
}

func (v *viewAuthorityMenuMysql) Initialize() error {
	var entity AuthorityMenus
	sql := `
	CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW @table_name AS
	select @menus.id                AS id,
		   @menus.path              AS path,
		   @menus.icon              AS icon,
		   @menus.name              AS name,
		   @menus.sort              AS sort,
		   @menus.title             AS title,
		   @menus.hidden            AS hidden,
		   @menus.component         AS component,
		   @menus.parent_id         AS parent_id,
		   @menus.created_at        AS created_at,
		   @menus.updated_at        AS updated_at,
		   @menus.deleted_at        AS deleted_at,
		   @authorities_menus.xd_base_menu_id      AS menu_id,
		   @authorities_menus.xd_authority_authority_id AS authority_id
	from (@authorities_menus
			 join @menus on ((@authorities_menus.xd_base_menu_id = @menus.id)));
	`
	sql = strings.ReplaceAll(sql, "@table_name", v.TableName())
	sql = strings.ReplaceAll(sql, "@menus", "xd_base_menus")
	sql = strings.ReplaceAll(sql, "@authorities_menus", entity.TableName())
	if err := global.XD_DB.Exec(sql).Error; err != nil {
		return errors.Wrap(err, v.TableName()+"视图创建失败!")
	}
	return nil
}

func (v *viewAuthorityMenuMysql) CheckDataExist() bool {
	err1 := global.XD_DB.Find(&[]model.XdMenu{}).Error
	err2 := errors.New(fmt.Sprintf("Error 1146: Table '%v.%v' doesn't exist", global.XD_CONFIG.MySqlConfig.Dbname, v.TableName()))
	if errors.As(err1, &err2) {
		return false
	}
	return true
}
