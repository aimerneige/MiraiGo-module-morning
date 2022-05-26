package morning

import (
	"sync"
	"time"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/go-co-op/gocron"
)

var instance *morning
var logger = utils.GetModuleLogger("internal.logging")

type morning struct {
}

func init() {
	instance = &morning{}
	bot.RegisterModule(instance)
}

func (m *morning) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "aimerneige.test.morning",
		Instance: instance,
	}
}

// Init 初始化过程
// 在此处可以进行 Module 的初始化配置
// 如配置读取
func (m *morning) Init() {
}

// PostInit 第二次初始化
// 再次过程中可以进行跨 Module 的动作
// 如通用数据库等等
func (m *morning) PostInit() {
}

// Serve 注册服务函数部分
func (m *morning) Serve(b *bot.Bot) {
	msg := message.NewSendingMessage().Append(message.NewText("早上好啊！"))
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("00:00").Do(func() {
		groupList := b.GroupList
		for _, group := range groupList {
			b.SendGroupMessage(group.Code, msg)
		}
	})
	s.StartAsync()
}

// Start 此函数会新开携程进行调用
// ```go
// 		go exampleModule.Start()
// ```
// 可以利用此部分进行后台操作
// 如 http 服务器等等
func (m *morning) Start(b *bot.Bot) {
}

// Stop 结束部分
// 一般调用此函数时，程序接收到 os.Interrupt 信号
// 即将退出
// 在此处应该释放相应的资源或者对状态进行保存
func (m *morning) Stop(b *bot.Bot, wg *sync.WaitGroup) {
	// 别忘了解锁
	defer wg.Done()
}
