/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:03
 * @FilePath: /pinkmoe_server/initialize/initDB/mysqlView.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
