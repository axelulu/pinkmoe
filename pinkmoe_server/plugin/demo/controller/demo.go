/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:52:54
 * @FilePath: /pinkmoe_server/plugin/demo/controller/demo.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package demo

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"regexp"
	"server/controller"
	"server/global"
	"server/model/response"
	"strings"
	"time"
)

// 这里每个controller执行init方法都要注册自动路由
func init() {
	controller.Register(&Demo{})
}

type Demo struct{}

type Postss struct { //带结构标签，反引号来包围字符串
	Title   string `json:"title"`
	Cover   string `json:"cover"`
	Postid  string `json:"postid"`
	Popp    string `json:"popp"`
	Content string `json:"content"`
}

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

func (demo *Demo) ExchangeGet(c *gin.Context) {
	var posts []Postss
	global.XD_DB.Table("xd_posts_copy1").Find(&posts)
	for i, p := range posts {
		if i < 2 {
			currentTime := time.Now()
			rand.Seed(time.Now().Unix() * int64(i))
			num := rand.Int31n(2) + 1
			num2 := rand.Int31n(9) + 1
			oldTime2 := currentTime.AddDate(0, 0, -int(int(num)+i))
			oldTime := oldTime2.Add(time.Hour * time.Duration(num2))
			//p.Cover = strings.Replace(p.Cover, `<img style="background:url( `, "", -1)
			//reg2 := regexp.MustCompile(` \) no-repeat(.+?)1920w\">`)
			//p.Cover = reg2.ReplaceAllString(p.Cover, "")
			//p.Cover = strings.Replace(p.Cover, "http://", "https://", -1)

			//reg3 := regexp.MustCompile(`((?:https?:\/\/)?(?:yun|pan|eyun)\.quark\.cn\/(?:s\/\w*(((-)?\w*)*)?|share\/\S*\d\w*))`)
			//link1 := reg3.FindAllString(p.Download, -1)
			//reg5 := regexp.MustCompile(`暗号：([a-z0-9]{4})`)
			//pwd := reg5.FindAllString(p.Download, -1)
			//pwd2 := ""
			//if len(pwd) > 0 {
			//	pwd2 = strings.Replace(pwd[0], "暗号：", "", -1)
			//}
			//reg4 := regexp.MustCompile(`((?:https?:\/\/)?(?:yun|pan|eyun)\.baidu\.com\/(?:s\/\w*(((-)?\w*)*)?|share\/\S*\d\w*))`)
			//link2 := reg4.FindAllString(p.Download, -1)
			////tags := strings.Split(p.Tag, " ")
			////uid, _ := uuid.FromString("7587adf5-faee-4097-9e82-1424ec818047")
			////var cate = "pixivPic"
			//var download []model.XdPostDownload
			//if len(link1) > 0 && len(link2) > 0 {
			//	download = make([]model.XdPostDownload, 2)
			//} else {
			//	download = make([]model.XdPostDownload, 1)
			//}

			//reg6 := regexp.MustCompile(`<div id="download-box" class="download-box">([\s\S]+?)</div></div></div>`)
			//p.Content = reg6.ReplaceAllString(p.Content, "")
			reg8 := regexp.MustCompile(`<div class="download-box" id="download-box" ref="downloadBox">([\s\S]+?)</div>\n                </div>\n            </div>\n`)
			p.Content = reg8.ReplaceAllString(p.Content, "")
			reg9 := regexp.MustCompile(`<p>==========================</p>([\s\S]+?)</a></h3>`)
			p.Content = reg9.ReplaceAllString(p.Content, "")
			reg10 := regexp.MustCompile(`<p>预览视频</p>([\s\S]+?)<p>购买正版</p>`)
			p.Content = reg10.ReplaceAllString(p.Content, "")
			//fmt.Printf("---%s---\n", p.Content)
			tags := strings.Split(p.Popp, "\n")
			fmt.Printf("-------------------tags:---%s-----------------\n", tags)
			opts := append(chromedp.DefaultExecAllocatorOptions[:],
				chromedp.Flag("headless", false),
			)
			allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
			defer cancel()

			// create chrome instance
			ctx, cancel := chromedp.NewContext(
				allocCtx,
				chromedp.WithLogf(log.Printf),
			)
			defer cancel()

			// create a timeout
			ctx, cancel = context.WithTimeout(ctx, 600*time.Second)
			defer cancel()

			// navigate to a page, wait for an element, click
			sel := `//*[@id="username"]`
			chromedp.Run(ctx,
				chromedp.Navigate("https://www.vrmoo.cn/"+p.Postid+".html"),
				chromedp.WaitVisible("#download-box"),
				chromedp.ActionFunc(func(ctx context.Context) error {
					_, exp, err := runtime.Evaluate(`window.localStorage.setItem('userData', '{"token":"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwczpcL1wvd3d3LnZybW9vLmNuIiwiaWF0IjoxNjY2NDE5MTM4LCJuYmYiOjE2NjY0MTkxMzgsImV4cCI6MTY2NzAyMzkzOCwiZGF0YSI6eyJ1c2VyIjp7ImlkIjoiMzgyMTQifX19.Dn_yBa6hEKUp-7x7slnRvZGZRgyciQtZFVp54BKncq8","user_display_name":"zhaoxingyue999","name":"zhaoxingyue999","id":"38214","avatar":"https://www.vrmoo.cn/wp-content/themes/b2/Assets/fontend/images/default-avatar.png","user_link":"https://www.vrmoo.cn/users/38214","is_admin":false,"verify":"","verify_icon":"","nonce":"","user_code":"ezAwmrlkn"}')`).Do(ctx)
					if err != nil {
						return err
					}
					if exp != nil {
						return exp
					}
					return nil
				}),
				chromedp.Reload(),
				chromedp.WaitVisible("document.querySelector(\"#gg-box > div > div > span\")", chromedp.ByJSPath),
				chromedp.Click("document.querySelector(\"#gg-box > div > div > span\")", chromedp.ByJSPath),
				chromedp.WaitVisible("document.querySelector(\"#download-box > div.download-list > div > div.download-info > div.download-button-box > button:nth-child(1)\")", chromedp.ByJSPath),
				chromedp.Click("document.querySelector(\"#download-box > div.download-list > div > div.download-info > div.download-button-box > button:nth-child(1)\")", chromedp.ByJSPath),
				chromedp.WaitVisible("document.querySelector(\"#gg-box > div > div > span\")", chromedp.ByJSPath),
				chromedp.Click("document.querySelector(\"#gg-box > div > div > span\")", chromedp.ByJSPath),
				chromedp.WaitVisible("document.querySelector(\"#download-page > a\")", chromedp.ByJSPath),
				chromedp.Click("document.querySelector(\"#download-page > a\")", chromedp.ByJSPath),
				//缓一缓
				chromedp.Sleep(100*time.Second),

				chromedp.SendKeys(sel, "username", chromedp.BySearch), //匹配xpath

			)
			//if len(link1) > 0 {
			//	download = append(download, model.XdPostDownload{
			//		Price:      10,
			//		PriceType:  "credit",
			//		Name:       p.Title,
			//		Url:        link1[0],
			//		ExtractPwd: "",
			//		UnpackPwd:  "",
			//	})
			//}
			//if len(link2) > 0 {
			//	download = append(download, model.XdPostDownload{
			//		Price:      10,
			//		PriceType:  "credit",
			//		Name:       p.Title,
			//		Url:        link2[0],
			//		ExtractPwd: pwd2,
			//		UnpackPwd:  "",
			//	})
			//}
			//err = mysql.CreatePost(request.CreatePostParams{
			//	Title:            p.Title,
			//	Exerpt:           p.Title,
			//	Content:          strings.TrimSpace(p.Content),
			//	Category:         cate,
			//	Author:           uid,
			//	Cover:            p.Cover,
			//	Type:             "post",
			//	From:             "",
			//	CommentStatus:    "true",
			//	Status:           "published",
			//	Topic:            tags,
			//	UpdatedAt:        global.XdTime{Time: oldTime},
			//	DownloadRelation: download,
			//})
			fmt.Printf("---%s---\n", oldTime)
		}
	}
	response.OkWithData(posts[0], c)
}
