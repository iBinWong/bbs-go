package spam

import (
	"bbs-go/model"
	"bbs-go/pkg/common"
	"bbs-go/services"
	"github.com/dchest/captcha"
)

type CaptchaStrategy struct{}

func (CaptchaStrategy) Name() string {
	return "CaptchaStrategy"
}

func (CaptchaStrategy) IsOpen() bool {
	// TODO 这个要做成配置
	return true
}

func (CaptchaStrategy) CheckTopic(user *model.User, form model.CreateTopicForm) error {
	if services.SysConfigService.GetConfig().TopicCaptcha && !captcha.VerifyString(form.CaptchaId, form.CaptchaCode) {
		return common.CaptchaError
	}
	return nil
}

func (CaptchaStrategy) CheckArticle(user *model.User, form model.CreateArticleForm) error {
	return nil
}

func (CaptchaStrategy) CheckComment(user *model.User, form model.CreateCommentForm) error {
	return nil
}
