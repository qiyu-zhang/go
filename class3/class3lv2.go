package main
/*有一份代码，实现了登录注册、账号密码加密存储的功能。不仅写的稀烂，还有很多地方忘了加锁！这会导致程序很不稳定！你需要帮他加上足够多的锁，来保证程序运行时的线程安全。*/
import (
	"bufio"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)
var lock sync.Mutex //锁 uh userHash 用户信息
var lock1 sync.Mutex //锁 saveUsers  中的fail
const (
	filePath = "./users.data"
	key = "woshifeiwu"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type userHash map[string]string

type Checker struct {
	uh userHash	// 用户信息
	registerUsers []User // 注册了但未保存的用户
}

func (c *Checker) SignIn() {
	defer fix()

	fmt.Println("请输入用户名和密码")
	var username, password string
	fmt.Scan(&username, &password)
	if _, ok := c.uh[username]; !ok {
		fmt.Println("查无此人")
		return
	}
	if c.uh[username] != password {
		fmt.Println("用户名密码错误")
		return
	}

	fmt.Println("登录成功")
}

func (c *Checker) SignUp() {
	defer fix()

	fmt.Println("请输入用户名")
	var username, password string
	fmt.Scan(&username)
	if _, ok := c.uh[username]; ok {
		fmt.Println("用户名已被占用")
		return
	}
	fmt.Println("请输入密码")
	for {
		fmt.Scan(&password)
		if len(password) >= 6 {
			break
		}
		fmt.Println("密码长度应大于六位，请重新输入")
	}


	c.registerUsers = append(c.registerUsers, User{
		Username: username,
		Password: password,
	})
	if len(c.registerUsers) > 10 {
		go c.Save()
	}
	c.uh[username] = password
}

func (c *Checker) Save() {
	defer fix()

	fail := saveUsers(c.registerUsers)
	c.registerUsers = fail
}

func initUsers() (userHash, error){
	defer fix()

	f, err := os.OpenFile(filePath, os.O_CREATE | os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()

	uh := make(userHash)
	var wg sync.WaitGroup // WaitGroup的作用是确保所有协程都执行完毕
	reader := bufio.NewReader(f)
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return nil, err
		}

		wg.Add(1)
		go func(buf []byte) {
			defer wg.Done()

			arr := strings.Split(string(buf), ".")
			sign , err := base64.StdEncoding.DecodeString(arr[1])
			if err != nil {
				fmt.Println(err)
				return
			}

			mac := hmac.New(sha256.New, []byte(key))
			mac.Write([]byte(arr[0]))
			s := mac.Sum(nil)
			if res := bytes.Compare(sign,s); res != 0 {
				fmt.Println("data error")
				return
			}

			u, err := base64.StdEncoding.DecodeString(arr[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			var user User
			err = json.Unmarshal(u, &user)
			if err != nil {
				fmt.Println(err)
				return
			}
			lock.Lock()   // 在进行存储操作之前，先加锁
			uh[user.Username] = user.Password
			lock.Unlock()	// 当存储完毕后，进行解锁
		}(buf)

	}
	wg.Wait()
	return uh, nil
}

func saveUsers(users []User) (fail []User){
	defer fix()

	f, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE | os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var wg sync.WaitGroup // WaitGroup的作用是确保所有协程都执行完毕
	writer := bufio.NewWriter(f)
	for _, user := range users{
		wg.Add(1)
		go func(user User) {
			defer wg.Done()
			buf, err := json.Marshal(user)
			user64 := base64.StdEncoding.EncodeToString(buf)
			if err != nil {
				fmt.Println(err)
				lock1.Lock() // 在进行存储操作之前，先加锁
				fail = append(fail, user)
				lock1.Unlock()// 当存储完毕后，进行解锁
				return
			}

			mac := hmac.New(sha256.New, []byte(key))
			mac.Write([]byte(user64))
			s := mac.Sum(nil)
			signature := base64.StdEncoding.EncodeToString(s)

			n, err := writer.Write(append([]byte(user64 + "." + signature), byte('\n')))
			if err != nil {
				fmt.Println(n, err)
				lock1.Lock() // 在进行存储操作之前，先加锁
				fail = append(fail, user)
				lock1.Unlock()// 当存储完毕后，进行解锁
				return
			}
		} (user)
	}
	wg.Wait()
	writer.Flush()
	return
}

func showList() {
	fmt.Println("请选择操作：")
	fmt.Println("1、登录")
	fmt.Println("2、注册")
	fmt.Println("3、退出")
}

func main() {
	defer fix()
	checker := Checker{}
	var err error
	checker.uh, err = initUsers()
	if err != nil {
		return
	}


	var opt int
	for {
		showList()
		_, err := fmt.Scanln(&opt)
		if err != nil || opt < 1 || opt > 3 {
			fmt.Println("请输入正确的操作序号")
			continue
		}

		switch opt {
		case 1:
			checker.SignIn()
		case 2:
			checker.SignUp()
		case 3:
			checker.Save()
			return
		}
	}
}

func fix() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}