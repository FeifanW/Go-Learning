package monster

import "testing"

// 测试用例，测试Store方法
func TestStore(t *testing.T) {
	// 先创建一个Monster实例
	monster := Monster{
		Name:  "孙悟空",
		Age:   500,
		Skill: "变身",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() 错误，希望为%v 实际为=%v", true, res)
	}
	t.Logf("monster.Store() 测试成功！")
}
