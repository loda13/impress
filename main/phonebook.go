package main

import (
	"fmt"
	"sort"
)

/**
 * @Author: tang
 * @mail: yuetang2
 * @Date: 2022/5/28 17:26
 * @Desc: 写一个终端管理的内存版本的电话簿，要求： 1）电话簿每个记录包括姓名和号码; 2）电话簿支持有序列表查看(按姓名首字母排序); 3）电话簿支持添加和修改
 */

func Show(phonebook map[string]int) {
	if len(phonebook) == 0 {
		fmt.Println("通讯录为空，请先添加联系人")
	} else {
		fmt.Println("当前联系人:")
		for k, v := range phonebook {
			fmt.Printf("姓名：%v,电话：%v\n", k, v)
		}
	}
}

func Sort(phonebook map[string]int) {
	ks := []string{}
	for k := range phonebook {
		ks = append(ks, k)
	}

	sort.Strings(ks)
	fmt.Println("姓名首字母排序:")
	for i := 0; i < len(phonebook); i++ {
		fmt.Printf("姓名：%v，电话：%v\n", ks[i], phonebook[ks[i]])
	}
}

func Select(phonebook map[string]int) {
	var name string
	index := -1
	fmt.Println("请输入要查询的姓名：")
	fmt.Scan(&name)

	for key, value := range phonebook {
		if key == name {
			index = value
			fmt.Println("联系人姓名", key)
			fmt.Printf("电话:%v\n", phonebook[key])

		}
	}
	if index == -1 {
		fmt.Println("没有查询到该联系人")
		return
	} else {
		return
	}

}

func Delete(phonebook map[string]int) {
	var name string

	fmt.Println("请输入要删除的姓名：")
	fmt.Scan(&name)

	for k, _ := range phonebook {
		if k == name {
			delete(phonebook, k)
			break
		}
	}
	Show(phonebook)
}

func Update(phonebook map[string]int) {
	var name string
	var num int
	var newname string
	var newphone int

	if phonebook != nil {
		for {
			fmt.Println("编辑用户名称请按1，编辑用户电话请按2，退出请按3")
			fmt.Scan(&num)

			switch num {
			case 1:
				fmt.Println("请输入要编辑的用户姓名：")
				fmt.Scan(&name)
				for k, v := range phonebook {
					if k == name {
						fmt.Println("请输入要新的用户姓名：")
						fmt.Scan(&newname)
						phonebook[newname] = v
						delete(phonebook, k)
						fmt.Printf("联系人%s已修改为%s", name, newname)
					}
				}
				Show(phonebook)
			case 2:
				fmt.Println("请输入要编辑的用户姓名：")
				fmt.Scan(&name)
				for k, _ := range phonebook {
					if k == name {
						fmt.Println("请输入新的号码：")
						fmt.Scan(&newphone)
						phonebook[k] = newphone
					}
				}
				Show(phonebook)
			case 3:
				break
			}
			break
		}
	} else {
		fmt.Println("没有查询到要编辑的联系人")
	}
}

func Add(phonebook map[string]int) {
	var name string
	var phone int
	var exit string

	fmt.Println("请输入姓名：")
	fmt.Scan(&name)

	for {
		fmt.Println("请输入电话号码：")
		fmt.Scan(&phone)
		phonebook[name] = phone

		fmt.Println("结束录入输入exit")
		fmt.Scan(&exit)

		if exit == "exit" {
			break
		} else {
			continue
		}
	}
	fmt.Println("添加成功")
	Show(phonebook)
}

func Menu() {
	fmt.Println("-------Menu-------------")
	fmt.Println("输出当前联系人信息，请按1")
	fmt.Println("排序联系人信息，请按2")
	fmt.Println("添加联系人信息，请按3")
	fmt.Println("删除联系人信息，请按4")
	fmt.Println("编辑联系人信息，请按5")
	fmt.Println("查找联系人信息，请按6")
	fmt.Println("退出，请按0")
	fmt.Println("--------End------------")
}
func main() {
	phonebook := make(map[string]int)
	phonebook["Kobe"] = 13323233232
	phonebook["John"] = 15656565566
	phonebook["Bob"] = 17677337733
	phonebook["Obama"] = 16555665561
	phonebook["Alice"] = 17666767338
	Menu()
	var num int
	for {
		fmt.Scan(&num)
		switch num {

		case 1:
			Show(phonebook)
			Menu()
		case 2:
			Sort(phonebook)
			Menu()
		case 3:
			Add(phonebook)
			Menu()
		case 4:
			Delete(phonebook)
			Menu()
		case 5:
			Update(phonebook)
			Menu()
		case 6:
			Select(phonebook)
			Menu()
		case 0:
			return
		}
	}

}

//todo
//异常处理
