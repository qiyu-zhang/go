package main

import "fmt"

type Author struct{
	Name string //名字
	VIP bool	//是否是高贵的带会员
	Icon string	 //头像
	Signature string	//签名
	Focus int	 //关注人数
}
type Video struct {
	AuthorInfo Author	//作者信息
	VideoName string //视频名称
	ThumbUpNum int	//点赞数
	CollectionNum int	//收藏数
	CoinOperatedNum int //投币数
}
//点赞
func(v *Video) ThumbUp()  {
	v.ThumbUpNum++
	fmt.Println("点赞  该视频点赞数：",v.ThumbUpNum)
}
//收藏
func(v *Video) Collection() {
	v.CollectionNum++
	fmt.Printf("收藏  该视频收藏数：%d\n",v.CollectionNum)
}
//投币
func(v *Video) CoinOperated() {
	v.CoinOperatedNum++
	fmt.Printf("投币  该视频投币数：%d \n",v.CoinOperatedNum)
}
//一键三连
func(v *Video) OneKeyThreeLink()  {
	v.ThumbUpNum++
	v.CollectionNum++
	v.CoinOperatedNum++
	fmt.Printf("一键三连  该视频点赞数：%d; 收藏数：%d; 投币数：%d \n",v.ThumbUpNum,v.CollectionNum,v.CoinOperatedNum)
}

//视频发布
func VideoRelease(auName string, vidName string) (reVideo Video){
	fmt.Println("发布作者", auName, "的视频", vidName)
	reVideo = Video{
		AuthorInfo: Author{Name: auName, VIP: false, Icon: "balabala", Signature: "go!go!", Focus: 0},
		VideoName: vidName,
		ThumbUpNum: 0,
		CollectionNum: 0,
		CoinOperatedNum: 0,
	}
	return
}
func main()  {
	var oneVideo Video = VideoRelease("moren", "cheng")
	fmt.Println("视频结构体为：",oneVideo)
	//测试
	fmt.Println("测试")
	//点赞
	oneVideo.ThumbUp()
	//收藏
	oneVideo.Collection()
	//投币
	oneVideo.CoinOperated()
	//一键三连
	oneVideo.OneKeyThreeLink()
}