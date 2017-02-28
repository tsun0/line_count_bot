package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func getTotal() (string, error) {

	ddb := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	// GetItem操作の入力を表します。
	params := &dynamodb.GetItemInput{
		TableName: aws.String("count"), // テーブル名

		Key: map[string]*dynamodb.AttributeValue{
			"id": { // キー名
				N: aws.String("1"), // 持ってくるキーの値
			},
		},
		AttributesToGet: []*string{
			aws.String("total"), // 欲しいデータの名前
		},
		ConsistentRead: aws.Bool(true), // 常に最新を取得するかどうか

		//返ってくるデータの種類
		ReturnConsumedCapacity: aws.String("NONE"),
	}

	// GetItem操作は、指定された主キーを持つ項目の属性のセットを返します。
	// respはGetItem操作の出力を表します。
	resp, err := ddb.GetItem(params)

	if err != nil {
		log.Println(err)
		return "", err
	}

	//resp.Item[項目名].型 でデータへのポインタを取得
	log.Println(*resp.Item["total"].S)

	return *resp.Item["total"].S, err
}
