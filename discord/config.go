// FIXME: 일단 함수형 옵션 패턴 보류

package discord

import (
	"errors"
)

type Options struct {
	Token *string
}

type SetOption func(options *Options) error

func WithToken(token string) SetOption {
	return func(options *Options) error {
		if token == "" {
			return errors.New("bot Token 이 비었습니다.")
		}
		options.Token = &token
		return nil
	}
}
