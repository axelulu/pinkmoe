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

//
//type Postss struct { //带结构标签，反引号来包围字符串
//	Title     string `json:"title"`
//	Cover     string `json:"cover"`
//	Expert    string `json:"expert"`
//	Tag       string `json:"tags"`
//	Content   string `json:"content"`
//	Download  string `json:"download"`
//	UpdatedAt string `json:"updated_at"`
//}
//
//type pinkPost struct {
//	ID          string `json:"ID"`
//	PostTitle   string `json:"post_title"`
//	PostContent string `json:"post_content"`
//	PostExcerpt string `json:"post_excerpt"`
//	MenuOrder   string `json:"menu_order"`
//}
//
//type resPinkPost struct {
//	ID          string                 `json:"ID"`
//	PostTitle   string                 `json:"post_title"`
//	PostContent string                 `json:"post_content"`
//	PostExcerpt string                 `json:"post_excerpt"`
//	MenuOrder   string                 `json:"menu_order"`
//	Cover       string                 `json:"cover"`
//	Menu        string                 `json:"menu"`
//	Downloads   []model.XdPostDownload `json:"downloads"`
//	Tag         []string               `json:"tag"`
//}
//
//type termRelationships struct {
//	TermTaxonomyId string `json:"term_taxonomy_id"`
//	ObjectId       string `json:"object_id"`
//}
//
//type termTaxonomy struct {
//	TermTaxonomyId string `json:"term_taxonomy_id"`
//	TermId         string `json:"term_id"`
//	Taxonomy       string `json:"taxonomy"`
//}
//
//type terms struct {
//	TermId string `json:"term_id"`
//	Name   string `json:"name"`
//}
//
//type downloads struct {
//	Name   string `json:"name"`
//	Credit string `json:"credit"`
//	Link   string `json:"link"`
//	Pwd    string `json:"pwd"`
//	Pwd2   string `json:"pwd2"`
//}
//
//type download struct {
//	MetaValue string `json:"meta_value"`
//}
//
//func (demo *Demo) ExchangeGet(c *gin.Context) {
//	mysqlConfig := my.Config{
//		DSN:                       "root:pinkmoe" + "@tcp(localhost:3306)/pinkacg?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
//		DefaultStringSize:         191,                                                                                      // string 类型字段的默认长度
//		SkipInitializeWithVersion: false,                                                                                    // 根据版本自动配置
//	}
//	if db, err := gorm.Open(my.New(mysqlConfig), &gorm.Config{}); err != nil {
//		os.Exit(0)
//	} else {
//		numssss := 0
//		numssss2 := 0
//		var PinkPost []pinkPost
//		var ResPinkPost []resPinkPost
//		db.Table("wp_posts").Where("post_type = 'post' OR post_type = 'revision'").Where("ID < ? AND ID > ?", 60000, 39999).Find(&PinkPost)
//		for i, post := range PinkPost {
//			var down download
//			var dss []downloads
//			var downs []model.XdPostDownload
//			db.Table("wp_postmeta").Where("post_id = ?", post.ID).Where("meta_key = 'ghost_download'").First(&down)
//			phpserialize.Unmarshal([]byte(down.MetaValue), &dss)
//			for i := 0; i < len(dss); i++ {
//				downs = append(downs, model.XdPostDownload{
//					Price:      20,
//					PriceType:  "credit",
//					Name:       post.PostTitle,
//					Url:        dss[i].Link,
//					ExtractPwd: dss[i].Pwd,
//					UnpackPwd:  dss[i].Pwd2,
//				})
//			}
//			post.PostContent = strings.Replace(post.PostContent, "https://img.catacg.cn/", "https://pinkmoe.coderzhaolu.com/", -1)
//			var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
//			imgs := imgRE.FindAllStringSubmatch(post.PostContent, -1)
//			if len(imgs) > 0 {
//				var po = resPinkPost{
//					ID:          post.ID,
//					PostTitle:   post.PostTitle,
//					PostContent: post.PostContent,
//					PostExcerpt: post.PostExcerpt,
//					Cover:       imgs[0][1],
//					Downloads:   downs,
//				}
//				var TermRelationships []termRelationships
//				db.Table("wp_term_relationships").Where("object_id = ?", po.ID).Find(&TermRelationships)
//				if len(TermRelationships) == 0 {
//					continue
//				}
//				for _, relationship := range TermRelationships {
//					// 标签
//					var TermTaxonomy []termTaxonomy
//					db.Table("wp_term_taxonomy").Where("term_taxonomy_id = ?", relationship.TermTaxonomyId).Where("taxonomy = 'post_tag'").Find(&TermTaxonomy)
//					for _, taxonomy := range TermTaxonomy {
//						var Terms terms
//						db.Table("wp_terms").Where("term_id = ?", taxonomy.TermId).First(&Terms)
//						po.Tag = append(po.Tag, Terms.Name)
//					}
//					// 分类
//					var TermTaxonomyCat []terms
//					db.Table("wp_term_taxonomy").Where("term_taxonomy_id = ?", relationship.TermTaxonomyId).Where("taxonomy = 'category'").Find(&TermTaxonomyCat)
//					var Termss []terms
//					if len(TermTaxonomyCat) > 0 {
//						db.Table("wp_terms").Where("term_id = ?", TermTaxonomyCat[0].TermId).Find(&Termss)
//						po.Menu = Termss[0].Name
//					}
//				}
//				currentTime := time.Now()
//				rand.Seed(time.Now().Unix() * int64(i))
//				num := rand.Int31n(2) + 1
//				num2 := rand.Int31n(3) + 1
//				oldTime2 := currentTime.AddDate(0, 0, -int(int(num)+i))
//				oldTime := oldTime2.Add(time.Hour * time.Duration(num2))
//				uid, _ := uuid.FromString("7587adf5-faee-4097-9e82-1424ec818047")
//				numssss2++
//				if po.Menu == "番剧" {
//					numssss++
//					var cate = "japanAnime"
//					mysql.CreatePost(request.CreatePostParams{
//						Title:            po.PostTitle,
//						Exerpt:           po.PostExcerpt,
//						Content:          strings.TrimSpace(po.PostContent),
//						Category:         cate,
//						Author:           uid,
//						Cover:            po.Cover,
//						Type:             "post",
//						From:             "",
//						CommentStatus:    "true",
//						Status:           "published",
//						Topic:            po.Tag,
//						UpdatedAt:        global.XdTime{Time: oldTime},
//						DownloadRelation: downs,
//					})
//				}
//				if po.Menu == "OST" || po.Menu == "OP" || po.Menu == "ED" {
//					numssss++
//					var cate = "musicDownload"
//					mysql.CreatePost(request.CreatePostParams{
//						Title:            po.PostTitle,
//						Exerpt:           po.PostExcerpt,
//						Content:          strings.TrimSpace(po.PostContent),
//						Category:         cate,
//						Author:           uid,
//						Cover:            po.Cover,
//						Type:             "post",
//						From:             "",
//						CommentStatus:    "true",
//						Status:           "published",
//						Topic:            po.Tag,
//						UpdatedAt:        global.XdTime{Time: oldTime},
//						DownloadRelation: downs,
//					})
//				}
//				if po.Menu == "美影" {
//					numssss++
//					var cate = "europeMovie"
//					mysql.CreatePost(request.CreatePostParams{
//						Title:            po.PostTitle,
//						Exerpt:           po.PostExcerpt,
//						Content:          strings.TrimSpace(po.PostContent),
//						Category:         cate,
//						Author:           uid,
//						Cover:            po.Cover,
//						Type:             "post",
//						From:             "",
//						CommentStatus:    "true",
//						Status:           "published",
//						Topic:            po.Tag,
//						UpdatedAt:        global.XdTime{Time: oldTime},
//						DownloadRelation: downs,
//					})
//				}
//				if po.Menu == "美剧" {
//					numssss++
//					var cate = "USATv"
//					mysql.CreatePost(request.CreatePostParams{
//						Title:            po.PostTitle,
//						Exerpt:           po.PostExcerpt,
//						Content:          strings.TrimSpace(po.PostContent),
//						Category:         cate,
//						Author:           uid,
//						Cover:            po.Cover,
//						Type:             "post",
//						From:             "",
//						CommentStatus:    "true",
//						Status:           "published",
//						Topic:            po.Tag,
//						UpdatedAt:        global.XdTime{Time: oldTime},
//						DownloadRelation: downs,
//					})
//				}
//				if po.Menu == "日剧" || po.Menu == "韩剧" || po.Menu == "影视" {
//					numssss++
//					var cate = "japanTv"
//					mysql.CreatePost(request.CreatePostParams{
//						Title:            po.PostTitle,
//						Exerpt:           po.PostExcerpt,
//						Content:          strings.TrimSpace(po.PostContent),
//						Category:         cate,
//						Author:           uid,
//						Cover:            po.Cover,
//						Type:             "post",
//						From:             "",
//						CommentStatus:    "true",
//						Status:           "published",
//						Topic:            po.Tag,
//						UpdatedAt:        global.XdTime{Time: oldTime},
//						DownloadRelation: downs,
//					})
//				}
//				if po.Menu == "轻小说" || po.Menu == "其他转载" {
//					numssss++
//					var cate = "lightNovelDownload"
//					mysql.CreatePost(request.CreatePostParams{
//						Title:            po.PostTitle,
//						Exerpt:           po.PostExcerpt,
//						Content:          strings.TrimSpace(po.PostContent),
//						Category:         cate,
//						Author:           uid,
//						Cover:            po.Cover,
//						Type:             "post",
//						From:             "",
//						CommentStatus:    "true",
//						Status:           "published",
//						Topic:            po.Tag,
//						UpdatedAt:        global.XdTime{Time: oldTime},
//						DownloadRelation: downs,
//					})
//				}
//				if po.Menu == "游戏" || po.Menu == "pc" {
//					numssss++
//					var cate = "steamPc"
//					mysql.CreatePost(request.CreatePostParams{
//						Title:            po.PostTitle,
//						Exerpt:           po.PostExcerpt,
//						Content:          strings.TrimSpace(po.PostContent),
//						Category:         cate,
//						Author:           uid,
//						Cover:            po.Cover,
//						Type:             "post",
//						From:             "",
//						CommentStatus:    "true",
//						Status:           "published",
//						Topic:            po.Tag,
//						UpdatedAt:        global.XdTime{Time: oldTime},
//						DownloadRelation: downs,
//					})
//				}
//				ResPinkPost = append(ResPinkPost, po)
//			}
//		}
//		println("-----------------------------------")
//		println(numssss)
//		println("-----------------------------------")
//		println(numssss2)
//		response.OkWithData(ResPinkPost, c)
//	}
//}

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
