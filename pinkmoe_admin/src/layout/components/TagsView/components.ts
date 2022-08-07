/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:19:39
 * @FilePath: /pinkmoe_admin/src/layout/components/TagsView/components.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import {
  DownOutlined,
  ReloadOutlined,
  CloseOutlined,
  VerticalRightOutlined,
  VerticalLeftOutlined,
  ColumnWidthOutlined,
  MinusOutlined,
} from '@ant-design/icons-vue';
import { Dropdown, Tabs, Card } from 'ant-design-vue';

export default {
  [Tabs.name]: Tabs,
  [Tabs.TabPane.name]: Tabs.TabPane,
  [Dropdown.name]: Dropdown,
  [Card.name]: Card,
  MinusOutlined,
  DownOutlined,
  ReloadOutlined,
  CloseOutlined,
  VerticalRightOutlined,
  VerticalLeftOutlined,
  ColumnWidthOutlined,
};
