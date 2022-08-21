package model

import (
	_ "github.com/go-sql-driver/mysql"
)

// Users ユーザー情報のテーブル情報
type Users struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Gender    int    `json:"gender"`
	Age       string `json:"age"`
	Height    int    `json:"height"`
	Uuid      string `json:"uuid"`
	Mail      string `json:"mail"`
	Icon      string `json:"icon"`
	CreatedAt string `json:"created_at" sql:"not null;type:date"`
	UpdatedAt string `json:"update_at" sql:"not null;type:date"`
}

//ユーザーの情報をresponseで返すための構造体
type UsersAddStatus struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender int    `json:"gender"`
	Age    string `json:"age"`
	Height int    `json:"height"`
	Uuid   string `json:"uuid"`
	Mail   string `json:"mail"`
	Icon   string `json:"icon"`
	Status bool   `json:"status"`
}

// Coordinates コーディネート情報のテーブル情報
type Likes struct {
	Id              string `json:"id"`
	Coordinate_id   string `json:"coordinate_id"`
	Send_user_id    string `json:"send_user_id"`
	Receive_user_id string `json:"receive_user_id"`
	Lat             string `json:"lat"`
	Lng             string `json:"lng"`
	CreatedAt       string `json:"created_at" sql:"not null;type:date"`
	UpdatedAt       string `json:"update_at" sql:"not null;type:date"`
}

// Coordinates コーディネート情報のテーブル情報
type Coordinates struct {
	Coordinate_id string `json:"coordinate_id"`
	User_id       string `json:"user_id"`
	Put_flag      int    `json:"put_flag"`
	Public        int    `json:"public"`
	Image         string `json:"image"`
	Category      string `json:"category"`
	Brand         string `json:"brand"`
	Price         string `json:"price"`
	Ble           string `json:"ble"`
	CreatedAt     string `json:"created_at" sql:"not null;type:date"`
	UpdatedAt     string `json:"update_at" sql:"not null;type:date"`
}

// ユーザー情報
type Requestuser struct {
	Name   string `json:"name"`
	Gender int    `json:"gender"`
	Age    string `json:"age"`
	Height int    `json:"height"`
	Uuid   string `json:"uuid"`
	Mail   string `json:"mail"`
	Icon   string `json:"icon"`
}

//服のコーディネート情報
type CoordinatesAdd struct {
	User_id string `json:"user_id"`
	Ble     string `json:"ble"`
	Public  bool   `json:"public"`
	Image   string `json:"image"`
	Items   []Item `json:"items"`
}

//服の情報
type Item struct {
	Category string `json:"category"`
	Brand    string `json:"brand"`
	Price    string `json:"price"`
}

type Ble struct {
	Coordinate_id string  `json:"coordinate_id"`
	Image         string  `json:"image"`
	Items         []*Item `json:"items"`
	Users         Users   `json:"users"`
	Status        bool    `json:"status"`
}

type Send_user struct {
	Gender int    `json:"gender"`
	Age    string `json:"age"`
	Height int    `json:"height"`
	Lat    string `json:"lat"`
	Lng    string `json:"lng"`
}

type Map struct {
	Coordinate_id string       `json:"coordinate_id"`
	User_id       string       `json:"user_id"`
	Image         string       `json:"image"`
	Send_users    []*Send_user `json:"send_users"`
}

type Maps struct {
	Maps   []*Map `json:"map"`
	Status bool   `json:"status"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
type UserResponse struct {
	Status bool   `json:"status"`
	Id     string `json:"id"`
}
type TrueResponse struct {
	Status bool `json:"status"`
}
