package mysql

//import (
//	"encoding/json"
//	"fmt"
//)
//
//// GetList 获取列表
//func GetList(req *dao.PluginsManageSearchReq) (total, page int, list []*dao.CsPluginListRes, err error) {
//	//同步服务端插件商城
//	total, page, list, err = syncFromStore(req)
//	return
//}
//
//// ChangeStatus 修改状态
//func ChangeStatus(ctx context.Context, req *dao.PluginsManageStatusReq) error {
//	_, err := dao.PluginsManage.Ctx(ctx).Where(dao.PluginsManage.Columns.StoreId, req.PluginId).Update(g.Map{
//		dao.PluginsManage.Columns.Status: req.Status,
//	})
//	return err
//}
//
////同步插件商城中的插件
//func syncFromStore(req *dao.PluginsManageSearchReq) (total, page int, csPluginList []*dao.CsPluginListRes, err error) {
//	storeUrl := g.Cfg().GetString("plugin.serverUrl") + "/codeStore/pluginList"
//	res := (*ghttp.ClientResponse)(nil)
//	if req.PageNum == 0 {
//		req.PageNum = 1
//	}
//	page = req.PageNum
//	res, err = g.Client().Ctx(req.Ctx).Get(storeUrl, g.MapStrAny{
//		"pageNum":    req.PageNum,
//		"PageSize":   req.PageSize,
//		"pluginName": req.PTitle,
//	})
//	if err != nil {
//		g.Log().Error(err)
//		err = gerror.New("获取插件数据失败")
//		return
//	}
//	defer res.Close()
//	var data map[string]interface{}
//	b := res.ReadAll()
//	err = json.Unmarshal(b, &data)
//	if err != nil {
//		g.Log().Error(err)
//		err = gerror.New("获取插件数据失败")
//		return
//	}
//	if gconv.Int(data["code"]) == 0 {
//		err = gconv.Structs((data["data"].(g.Map))["list"], &csPluginList)
//		if err != nil {
//			return
//		}
//		csPluginList, err = s.updatePlugins(req.Ctx, csPluginList)
//		if err != nil {
//			return
//		}
//		total = gconv.Int((data["data"].(g.Map))["total"])
//	} else {
//		err = gerror.New(data["msg"].(string))
//		return
//	}
//	return
//}
//
//// 更新插件数据
//func updatePlugins(ctx context.Context, csList []*dao.CsPluginListRes) (newList []*dao.CsPluginListRes, err error) {
//	ids := make([]uint, len(csList))
//	for k, v := range csList {
//		ids[k] = v.PluginId
//	}
//	//查询插件信息
//	var pluginList []*model.PluginsManage
//	err = dao.PluginsManage.Ctx(ctx).Where(dao.PluginsManage.Columns.StoreId+" in(?)", ids).Scan(&pluginList)
//	if err != nil {
//		return
//	}
//	hasIds := garray.NewArraySize(len(pluginList), 100)
//	gmp := gmap.New()
//	for k, v := range pluginList {
//		hasIds.Set(k, v.StoreId)
//		gmp.Set(v.StoreId, v)
//	}
//	for _, v := range csList {
//		pluginId := gconv.Int(v.PluginId)
//		if hasIds.Len() > 0 && hasIds.Contains(pluginId) {
//			plugin := gmp.Get(pluginId).(*model.PluginsManage)
//			//修改
//			version := plugin.Version
//			if plugin.IsInstall == 0 && len(v.PluginInfo) > 0 {
//				version = v.PluginInfo[0].InfoVersion
//			}
//			err = s.Edit(ctx, &dao.PluginsManageEditReq{
//				Id:            plugin.Id,
//				StoreId:       pluginId,
//				PName:         v.CodeName,
//				PTitle:        v.PluginName,
//				PDescription:  v.Description,
//				PAuth:         v.MemName,
//				Status:        plugin.Status,
//				Version:       version,
//				Price:         v.PluginPrice,
//				DownloadTimes: gconv.Uint(v.DownloadTimes),
//				IsInstall:     plugin.IsInstall,
//			})
//			v.Status = plugin.Status
//			v.Version = version
//			v.IsInstall = plugin.IsInstall
//			v.PluginPriceStr = s.Int64ToDecimal(gconv.Int64(v.PluginPrice))
//		} else {
//			//新增
//			version := ""
//			if len(v.PluginInfo) > 0 {
//				version = v.PluginInfo[0].InfoVersion
//			}
//			err = s.Add(ctx, &dao.PluginsManageAddReq{
//				StoreId:       pluginId,
//				PName:         v.CodeName,
//				PTitle:        v.PluginName,
//				PDescription:  v.Description,
//				PAuth:         v.MemName,
//				Status:        0,
//				Version:       version,
//				Price:         v.PluginPrice,
//				DownloadTimes: gconv.Uint(v.DownloadTimes),
//				IsInstall:     0,
//			})
//			v.Status = 0
//			v.Version = version
//			v.IsInstall = 0
//			v.PluginPriceStr = s.Int64ToDecimal(gconv.Int64(v.PluginPrice))
//		}
//		if err != nil {
//			return
//		}
//	}
//	newList = csList
//	return
//}
//
//// DecimalToInt64 元转分
//func DecimalToInt64(decimal string) (i int64) {
//	pos := gstr.PosR(decimal, ".")
//	integer := gconv.Int64(gstr.SubStr(decimal, 0, pos)) * 100
//	dec := int64(0)
//	if pos > -1 {
//		dec = gconv.Int64(gstr.SubStr(decimal, pos+1, 2))
//	}
//	i = integer + dec
//	return
//}
//
//// Int64ToDecimal 分转元
//func Int64ToDecimal(i int64) (decimal string) {
//	b := []byte(gconv.String(i))
//	for {
//		if len(b) >= 2 {
//			break
//		}
//		b = append([]byte{'0'}, b...)
//	}
//	integer := b[:len(b)-2]
//	if len(integer) == 0 {
//		integer = []byte{'0'}
//	}
//	dec := b[len(b)-2:]
//	decimal = fmt.Sprintf("%s.%s", integer, dec)
//	return
//}
//
//// Install 插件安装
//func Install(ctx context.Context, req *dao.PluginsManageInstallReq) (err error) {
//	//生成下载链接
//	storeUrl := g.Cfg().GetString("plugin.serverUrl") + "/codeStoreFrontAdmin/getDownloadInfo"
//	res := (*ghttp.ClientResponse)(nil)
//	res, err = g.Client().Ctx(ctx).Get(fmt.Sprintf("%s?pluginId=%d&version=%s&token=%s", storeUrl, req.PluginId, req.Version,
//		gurl.RawEncode(req.RToken)))
//	if err != nil {
//		return
//	}
//	defer res.Close()
//	var data map[string]interface{}
//	b := res.ReadAll()
//	err = json.Unmarshal(b, &data)
//	if err != nil {
//		return
//	}
//	if gconv.Int(data["code"]) == 0 {
//		url := fmt.Sprintf("%s/%s&token=%s", g.Cfg().GetString("plugin.serverUrl"),
//			(data["data"]).(string), gurl.RawEncode(req.RToken))
//		//下载插件并安装
//		err = s.downloadAndInstall(ctx, url)
//	} else {
//		err = gerror.New(data["msg"].(string))
//	}
//	return
//}
//
//// 下载并安装插件
//func downloadAndInstall(ctx context.Context, url string) error {
//	res, err := g.Client().Ctx(ctx).Get(url)
//	if err != nil {
//		return err
//	}
//	defer res.Close()
//	ct := res.Header.Get("Content-Type")
//	if gstr.ContainsI(ct, "json") {
//		var data map[string]interface{}
//		b := res.ReadAll()
//		err = json.Unmarshal(b, &data)
//		if err != nil {
//			return err
//		}
//		if gconv.Int(data["code"]) != 0 {
//			err = gerror.New(data["msg"].(string))
//		}
//		return err
//	} else {
//		//安装
//		//获取插件名称
//		fileName := res.Header.Get("content-disposition")
//		fileName = gstr.SubStr(fileName, gstr.PosR(fileName, "=")+1,
//			gstr.PosR(fileName, ".")-gstr.PosR(fileName, "=")-1)
//		err = s.InstallFile(ctx, res.ReadAll(), fileName)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//// InstallFile 安装插件文件
//func InstallFile(ctx context.Context, data []byte, fileName string) (err error) {
//	//获取插件下载路径
//	downloadPath := gfile.MainPkgPath() + "/data/installPlugins"
//	if !gfile.IsDir(downloadPath) {
//		err = gfile.Mkdir(downloadPath)
//		if err != nil {
//			return
//		}
//	}
//	g.Log().Debug(downloadPath, fileName)
//	// 删除安装临时文件
//	defer gfile.Remove(downloadPath + "/" + fileName)
//	err = gcompress.UnZipContent(data, downloadPath)
//	if err != nil {
//		return
//	}
//	//获取插件配置信息
//	var installCfg *gjson.Json
//	installCfg, err = gjson.Load(downloadPath + "/" + fileName + "/install.json")
//	if err != nil {
//		return
//	}
//	mustGfastVersion := installCfg.GetString("minGfastVersion")
//	if gstr.CompareVersion(mustGfastVersion, global.Version) > 0 {
//		err = gerror.New(fmt.Sprintf("您的gfast版本过低，此插件要求gfast版本为:%s", mustGfastVersion))
//		return
//	}
//	//获取本项目安装情况
//	var plugin *model.PluginsManage
//	err = dao.PluginsManage.Ctx(ctx).Where(dao.PluginsManage.Columns.PName, fileName).Limit(1).Scan(&plugin)
//	if err != nil {
//		return
//	}
//	if plugin == nil {
//		err = gerror.New("插件信息不存在，请刷新页面后再安装。")
//		return
//	}
//	//复制插件文件到对应目录
//	//1.后端
//	err = gfile.Copy(downloadPath+"/"+fileName+"/go/", gfile.MainPkgPath())
//	if err != nil {
//		return
//	}
//	//2.前端
//	fontRoot := g.Cfg().GetString("gen.frontDir")
//	if !gfile.IsDir(fontRoot) {
//		err = gerror.New("前端路径不存在，请配置gen.frontDir")
//		return
//	}
//	err = gfile.Copy(downloadPath+"/"+fileName+"/vue/", fontRoot+"/src")
//	if err != nil {
//		return
//	}
//	// 安装成功后修改插件安装状态及安装路径
//	_, err = dao.PluginsManage.Ctx(ctx).WherePri(plugin.Id).Update(g.Map{
//		dao.PluginsManage.Columns.IsInstall:   1,
//		dao.PluginsManage.Columns.Status:      1,
//		dao.PluginsManage.Columns.InstallPath: installCfg.GetString("installPath"),
//		dao.PluginsManage.Columns.Version:     installCfg.GetString("version"),
//	})
//	return
//}
