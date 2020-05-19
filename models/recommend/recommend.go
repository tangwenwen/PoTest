package recommendModel

import (
	recommendDao "PoTest/dao/recommend"
	"PoTest/plugins"
	recommendType "PoTest/types/recommend"
)

func PostWeibo(data *recommendType.PostWeiboReq) error {
	recordData := make([]*recommendType.Record, len(data.Req))
	postId := plugins.GetId(plugins.Recommend)
	for index, item := range data.Req {
		tempData := new(recommendType.Record)
		tempData.UserId1 = item.UserId1
		tempData.UserId2 = item.UserId2
		tempData.PostId = postId
		recordData[index] = tempData
	}

	if err := recommendDao.AddRecord(recordData); err != nil {
		return err
	}
	return nil
}

func Suggest(userId int) ([]recommendType.SuggestResp, error) {

	var suggestList []int
	var suggestResp []recommendType.SuggestResp
	type uni struct {
		user int
		post int64
	}

	recordList, err := recommendDao.GetAllPostRecord()
	if err != nil {
		return nil, err
	}
	mmap := make(map[uni][]int, len(recordList))
	for _, item := range recordList {
		recordData, err := recommendDao.GetPostRecordByPost(item.PostId)
		if err != nil {
			return nil, err
		}
		mmap[uni{
			user: item.UserId1,
			post: item.PostId,
		}] = recordData
	}

	for k, v := range mmap {
		if sliceContains(v, userId) {
			suggestList = append(suggestList, v...)
			elseList, err := recommendDao.GetUserPostRecord(k.user, k.post)
			suggestList = append(suggestList, elseList...)
			if err != nil {
				return nil, err
			}
		}
	}
	target := suggestList[:0]
	for _, item := range suggestList {
		if item != userId {
			target = append(target, item)
		}
	}
	for _, item := range target {
		suggestResp = append(suggestResp, recommendType.SuggestResp{UserId: item})
	}

	return suggestResp, nil
}

func sliceContains(slice []int, elem interface{}) bool {
	for _, item := range slice {
		if item == elem {
			return true
		}
	}
	return false
}
