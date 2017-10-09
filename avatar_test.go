package main

import "testing"

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("値が存在しない場合 AuthAvatar.getAvatarURL は ErrNoAvatarURL を返すべき")
	}

	// 値をセット
	testUrl := "http://url-to-avatar/"
	client.userData = map[string]interface{}{"avatar_url": testUrl}
	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("値が存在する場合 AuthAvatar.getAvatarURL はエラーを返すべきではない")
	} else if url != testUrl {
		t.Error("AuthAvatar.getAvatarURL は正しいURLを返すべき")
	}
}
