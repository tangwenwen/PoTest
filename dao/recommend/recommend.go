package recommendDao

import (
	"PoTest/plugins/db"
	recommendType "PoTest/types/recommend"
)

func AddRecord(recordList []*recommendType.Record) error {
	orm, err := db.GetEngine()
	if err != nil {
		return err
	}
	_, err = orm.Insert(recordList)
	if err != nil {
		return err
	}
	return nil
}

func GetAllPostRecord() ([]recommendType.Record, error) {
	orm, err := db.GetEngine()
	if err != nil {
		return nil, err
	}
	var recordList []recommendType.Record
	if err = orm.GroupBy("post_id,user_id_1").Find(&recordList); err != nil {
		return nil, err
	}
	return recordList, nil
}

func GetPostRecordByPost(postId int64) ([]int, error) {
	var result []int
	orm, err := db.GetEngine()
	if err != nil {
		return result, err
	}
	var recordList []recommendType.Record
	if err = orm.Where("Post_id = ?", postId).Find(&recordList); err != nil {
		return result, err
	}
	for _, item := range recordList {
		result = append(result, item.UserId2)
	}
	return result, nil
}

func GetUserPostRecord(userId int,postId int64)([]int,error){
	var result []int
	orm, err := db.GetEngine()
	if err != nil {
		return result, err
	}
	var recordList []recommendType.Record
	if err = orm.Where("Post_id != ?", postId).And("user_id_1 = ?",userId ).Find(&recordList); err != nil {
		return result, err
	}
	for _, item := range recordList {
		result = append(result, item.UserId2)
	}
	return result, nil
}
