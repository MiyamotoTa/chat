package main

import "errors"

var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません")

type Avatar interface {
	// GetAvatarURLは指定されたクライアントのアバターのURLを返します
	// 問題が発生した場合にはエラーを返します
	// とくに、URLを取得できなかった場合にはErrNoAvatarURLを返します
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct {
}

var UseAuthAvatar AuthAvatar

func (_ AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}
