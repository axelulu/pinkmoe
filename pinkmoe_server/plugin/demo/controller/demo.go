/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:52:54
 * @FilePath: /pinkmoe_server/plugin/demo/controller/demo.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package demo

import (
	"server/controller"
)

//这里每个controller执行init方法都要注册自动路由
func init() {
	controller.Register(&Demo{})
}

type Demo struct{}

type Postss struct { //带结构标签，反引号来包围字符串
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	Expert    string `json:"expert"`
	Tag       string `json:"tags"`
	Content   string `json:"content"`
	Download  string `json:"download"`
	UpdatedAt string `json:"updated_at"`
}

//func (demo *Demo) ExchangeGet(c *gin.Context) {
//global.XD_DB.Where("price_type = ?", "").Where("name = ?", "").Where("url = ?", "").Delete(&model.XdPostDownload{})
//jsonFile, err := os.Open("./p.json")
//if err != nil {
//	fmt.Println("error opening json file")
//	return
//}
//defer jsonFile.Close()
//
//jsonData, err := ioutil.ReadAll(jsonFile)
//if err != nil {
//	fmt.Println("error reading json file")
//	return
//}
//var post []Postss
//json.Unmarshal(jsonData, &post)
//for i, p := range post {
//	if i > 113 {
//		currentTime := time.Now()
//		rand.Seed(time.Now().Unix() * int64(i))
//		num := rand.Int31n(2) + 1
//		num2 := rand.Int31n(9) + 1
//		oldTime2 := currentTime.AddDate(0, 0, -int(int(num)+i))
//		oldTime := oldTime2.Add(time.Hour * time.Duration(num2))
//		p.Cover = strings.Replace(p.Cover, `<img style="background:url( `, "", -1)
//		reg2 := regexp.MustCompile(` \) no-repeat(.+?)1920w\">`)
//		p.Cover = reg2.ReplaceAllString(p.Cover, "")
//		p.Tag = strings.Replace(p.Tag, "标签:", "", -1)
//		p.Cover = strings.Replace(p.Cover, "http://", "https://", -1)
//
//		reg3 := regexp.MustCompile(`((?:https?:\/\/)?(?:yun|pan|eyun)\.quark\.cn\/(?:s\/\w*(((-)?\w*)*)?|share\/\S*\d\w*))`)
//		link1 := reg3.FindAllString(p.Download, -1)
//		reg5 := regexp.MustCompile(`暗号：([a-z0-9]{4})`)
//		pwd := reg5.FindAllString(p.Download, -1)
//		pwd2 := ""
//		if len(pwd) > 0 {
//			pwd2 = strings.Replace(pwd[0], "暗号：", "", -1)
//		}
//		reg4 := regexp.MustCompile(`((?:https?:\/\/)?(?:yun|pan|eyun)\.baidu\.com\/(?:s\/\w*(((-)?\w*)*)?|share\/\S*\d\w*))`)
//		link2 := reg4.FindAllString(p.Download, -1)
//		tags := strings.Split(p.Tag, " ")
//		uid, _ := uuid.FromString("7587adf5-faee-4097-9e82-1424ec818047")
//		var cate = "pixivPic"
//		var download []model.XdPostDownload
//		if len(link1) > 0 && len(link2) > 0 {
//			download = make([]model.XdPostDownload, 2)
//		} else {
//			download = make([]model.XdPostDownload, 1)
//		}
//
//		reg6 := regexp.MustCompile(`<div style="border:1px dashed #999999([\s\S]+?)<p>&nbsp;</p>`)
//		p.Content = reg6.ReplaceAllString(p.Content, "")
//		reg8 := regexp.MustCompile(`<p>• <strong>资源下载([\s\S]+?)<p>&nbsp;</p>`)
//		p.Content = reg8.ReplaceAllString(p.Content, "")
//		reg9 := regexp.MustCompile(`<p>点击下载([\s\S]+?)<p>&nbsp;</p>`)
//		p.Content = reg9.ReplaceAllString(p.Content, "")
//		reg7 := regexp.MustCompile(`<p><img class="aligncenter"([\s\S]+?)</noscript></p>`)
//		p.Content = reg7.ReplaceAllString(p.Content, "")
//		p.Content = strings.Replace(p.Content, "<p><img class=\"aligncenter\" src=\"https://ae01.alicdn.com/kf/H4d81ff54c6214eb5a05d441a192f537bI.jpg\"></p>", "", -1)
//		//fmt.Printf("%s", p.Content)
//		//println(cate)
//		//println(uid.String())
//		//println(tags)
//		//println(oldTime.String())
//
//		if len(link1) > 0 {
//			download = append(download, model.XdPostDownload{
//				Price:      10,
//				PriceType:  "credit",
//				Name:       p.Title,
//				Url:        link1[0],
//				ExtractPwd: "",
//				UnpackPwd:  "",
//			})
//		}
//		if len(link2) > 0 {
//			download = append(download, model.XdPostDownload{
//				Price:      10,
//				PriceType:  "credit",
//				Name:       p.Title,
//				Url:        link2[0],
//				ExtractPwd: pwd2,
//				UnpackPwd:  "",
//			})
//		}
//		err = mysql.CreatePost(request.CreatePostParams{
//			Title:            p.Title,
//			Exerpt:           p.Expert,
//			Content:          strings.TrimSpace(p.Content),
//			Category:         cate,
//			Author:           uid,
//			Cover:            p.Cover,
//			Type:             "post",
//			From:             "",
//			CommentStatus:    "true",
//			Status:           "published",
//			Topic:            tags,
//			UpdatedAt:        global.XdTime{Time: oldTime},
//			DownloadRelation: download,
//		})
//		//fmt.Printf("---%s---\n", oldTime)
//		if err != nil {
//			return
//		}
//	}
//}
//	response.OkWithData("ok", c)
//}
