package main

// package 外部可访问
type Monkey struct {
	Name   string `json:"name"`
	age    int
	Gender int `json:"gender"`
}

// package 内部可访问
type gorilla struct {
	*Monkey
	Weight int
}

// package 外部可访问
// monkey 带 * 表示是一个指针
// err 没带 * 表示是值类型
func BuildMonkey(name string, age, gender int) (monkey *Monkey, err error) {
	err = nil
	// & 表示引用指针，别担心，用错了编译不过
	monkey = &Monkey{
		Name:   name,
		age:    age,
		Gender: gender,
	}
	return
}

// monkey.Age()
func (monkey *Monkey) Age() int {
	return monkey.age
}
