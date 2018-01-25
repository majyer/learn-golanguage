package main
import b64 "encoding/base64"
import(
"fmt"
"math"
"time"
"errors"
"sync"
"sync/atomic"
"math/rand"
"sort"
"os"
"strings"
"regexp"
"bytes"
"encoding/json"
"strconv"
"net"
"net/url"
"crypto/sha1"
"io/ioutil"
"os/exec"
"syscall"
) 
var P=fmt.Println

const s1 string="const"
func main() {
	fmt.Println("hello world")

	fmt.Println("go"+" "+"lang")

	fmt.Println(true && false)
	fmt.Println(true || false)

	fmt.Println(!true)

	var astr string="hello world"
	fmt.Println(astr)

	var d int = 1
	fmt.Println(d)

	var a=true
	fmt.Println(a)

	var b int
	fmt.Println(b)

	c:="short"
	fmt.Println(c)

	fmt.Println(s1)
	
	const n int=1
	fmt.Println(n)

	const m=6000000000
	fmt.Println(m)

	const e=3e20/m
	fmt.Println(e)

	fmt.Println(math.Sin(m))

	i:=1
	for i<3 {
		fmt.Println(i)
		i++
	}

	for j:=7;j<=9;j++{
		fmt.Println(j)
	}

	for k:=0;k<=5;k++{
		if k%2==0{
			continue
		}
		if k%3==0{
			break
		}
		fmt.Println(n)
	}

	if num:=9;num<0{
		fmt.Println(num,"is negative")
	}else if num<10{
		fmt.Println(num,"has a digit")
	}else{
		fmt.Println(num,"has multiple digit")
	}

	f:=2
	fmt.Println("write",f," as")
	switch f{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("work day")
	}

	whatAmI:=func(i interface{}){
		switch t:=i.(type){
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Println("don't know type %T\n",t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	var A [5]int
	fmt.Println("emp:",A)

	A[4]=100
	fmt.Println("set :",A)
	fmt.Println("get:",A[4])

	fmt.Println("len",len(A))

	B:=[5]int{1,2,3,4,5}
	fmt.Println("dcl:",B)

	var twoD [2][3]int
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			twoD[i][j]=i+j
		}
	}
	fmt.Println("2d:",twoD)

	S:=make([]string,3)
	fmt.Println("emp:",S)

	S[0]="a"
	S[1]="b"
	S[2]="c"

	fmt.Println("set",S)
	fmt.Println("get:",S[2])
	fmt.Println("get s:",S[1:])
	fmt.Println("get S:",S[1:2])
	fmt.Println("len:",len(S))

	S=append(S,"d")
	S=append(S,"e")

	fmt.Println("apd:",S)
	fmt.Println("S new len",len(S))

	C:=make([]string,len(S))
	copy(C,S)
	fmt.Println("C:",C)

	g:=S[2:5]
	fmt.Println("S2-5:",g)

	g=S[2:]
	fmt.Println("S2:",g)

	g=S[:5]
	fmt.Println("S:5:",g)

	h:=[]string{"g","h","i"}
	fmt.Println("dcl:",h)

	twoH:=make([][]int,3)
	for i:=0;i<3;i++{
		innerlen:=i+1
		twoH[i]=make([]int,innerlen)
		for j:=0;j<innerlen;j++{
			twoH[i][j]=i+j
		}
	}
	fmt.Println("2D H:",twoH)

	l:=make(map[string]int)

	l["k1"]=1
	l["k2"]=2
	fmt.Println("map:",l)

	v1:=l["k1"]
	fmt.Println("v1:",v1)
	fmt.Println("len:",len(l))

	delete(l,"k2")
	fmt.Println("map:",l)

	_,prs:=l["k2"]
	fmt.Println("prs:",prs)

	n1:=map[string]int{"foo":1,"bar":2}
	fmt.Println("map:",n1)

	nums:=[]int{2,3,4}
	sum:=0
	for _,num:=range nums{
		sum+=num
	}
	fmt.Println("sum:",sum)

	for i,num:=range nums{
		if num==3{
			fmt.Println("index:",i)
		}
	}

	kvs:=map[string]string{"a":"apple","b":"banana"}
	for k,v:=range kvs{
		fmt.Printf("%s->%s\n",k,v)
	}

	for k:=range kvs{
		fmt.Println("key:",k)
	}

	for _,c:=range "gone"{
		fmt.Printf("%c\n",c)
	}

	res:=plus(1,2)
	fmt.Println("1+2 =:",res)

	res=plusplus(1,2,3)
	fmt.Println("1+2+3 =:",res)

	a1,b1:=vals()
	fmt.Println(a1,b1)

	_,c1:=vals()
	fmt.Println(c1)	

	sum1(1,2)
	sum1(1,3,4)

	nums1:=[]int{1,2,3,4}
	sum1(nums1...)

	nextInt:=inSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts:=inSeq()
	fmt.Println(newInts())
	fmt.Println(nextInt())
	fmt.Println(newInts())

	fmt.Println(fact(7))

	i2:=1
	fmt.Println("initial:",i2)

	zeroval(i2)
	fmt.Println("zeroval:",i2)

	zeroptr(&i2)
	fmt.Println("zeroptr:",i2)

	fmt.Println("pointer",&i2)

	fmt.Println(person{"bob",20})

	fmt.Println(person{name:"alice",age:30})

	fmt.Println(person{name:"Fred"})

	fmt.Println(&person{name:"Ann",age:40})

	s2:=person{name:"sean",age:50}
	fmt.Println(s2.name)

	sp:=&s2
	fmt.Println(sp.age)

	sp.age=53
	fmt.Println(sp.age)

	r:=rect{width:10,height:5}

	fmt.Println("area:",r.area())
	fmt.Println("perim:",r.perim())

	rp:=&r
	fmt.Println("area:",rp.area())
	fmt.Println("perim:",rp.perim())

	r64:=rect64{width:3,height:4}
	c64:=circle{radius:5}

	measure(r64)
	measure(c64)

	for _,i:=range []int{7,42}{
		if r,e:=f1(i);e!=nil{
			fmt.Println("f1 failed:",e)
		}else{
			fmt.Println("f1 worked:",r)
		}
	}
	for _,i:=range []int{7,42}{
		if r,e:=f2(i);e!=nil{
			fmt.Println("f2 failed:",e)
		}else{
			fmt.Println("f2 worked:",r)
		}
	}

	_,e3:=f2(42)
	if ae,ok:=e3.(*argError);ok{
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	f3("direct")

	go f3("goroutine")

	go func(msg string){
		fmt.Println(msg)
	}("going")

	//var input string
	//fmt.Scanln(&input)
	fmt.Println("done")

	//channels
	messages:=make(chan string)

	go func(){
		messages<-"ping"
	}()

	msg:=<-messages
	fmt.Println(msg)

	messages1:=make(chan string,2)

	messages1<-"buffered"
	messages1<-"channel"

	fmt.Println(<-messages1)
	fmt.Println(<-messages1)

	done:=make(chan bool,1)
	go worker(done)
	<-done

	pings:=make(chan string,1)
	pongs:=make(chan string,1)

	ping(pings,"pass message")
	pong(pings,pongs)
	fmt.Println(<-pongs)

	c3:=make(chan string)
	c4:=make(chan string)

	go func(){
		time.Sleep(time.Second*1)
		c3<-"one"
	}()

	go func(){
		time.Sleep(time.Second*1)
		c4<-"two"
	}()

	for i:=0;i<2;i++{
		select{
		case msg1:=<-c3:
			fmt.Println("received",msg1)
		case msg2:=<-c4:
			fmt.Println("received",msg2)
		}
	}

	c5:=make(chan string,1)

	go func(){
		time.Sleep(time.Second*2)
		c5<-"result 1"
	}()

	select{
	case res:=<-c5:
		fmt.Println(res)
	case <-time.After(time.Second*1):
		fmt.Println("timeout 1")
	}

	c6:=make(chan string,1)

	go func(){
		time.Sleep(time.Second*2)
		c6<-"result 2"
	}()

	select{
	case res:=<-c6:
		fmt.Println(res)
	case <-time.After(time.Second*3):
		fmt.Println("timeout 2")
	}

	message3:=make(chan string)
	signals:=make(chan bool)

	select{
	case msg:=<-message3:
		fmt.Println("received message",msg)
	default:
		fmt.Println("no message received")
	}

	msg3:="hi"
	select{
	case message3<-msg3:
		fmt.Println("send message",msg3)
	default:
		fmt.Println("no message sent")
	}

	select{
		case msg:=<-message3:
			fmt.Println("sent message",msg)
		case sig:=<-signals:
			fmt.Println("received signal",sig)
		default:
			fmt.Println("no activity")
	}

	jobs:=make(chan int,5)
	done1:=make(chan bool)

	go func(){
		for{
			j,more:=<-jobs
			if more{
				fmt.Println("received job",j)
			}else{
				fmt.Println("received all jobs")
				done1<-true
				return
			}
		}
	}()

	for j:=1;j<=3;j++{
		jobs <- j;
		fmt.Println("send jobs",j)
	}

	close(jobs)
	fmt.Println("send all jobs")
	<-done1

	queue:=make(chan string,2)
	queue<-"one"
	queue<-"two"
	close(queue)

	for elem:=range queue{
		fmt.Println(elem)
	}

	timer1:=time.NewTimer(time.Second*2)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2:=time.NewTimer(time.Second)
	go func(){
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	//time.Sleep(time.Second)

	stop2:=timer2.Stop()
	if stop2{
		fmt.Println("timer 2 stopped")
	}

	ticker:=time.NewTicker(time.Millisecond*500)
	go func(){
		for t:=range ticker.C{
			fmt.Println("Tick at",t)
		}
	}()

	time.Sleep(time.Millisecond*1600)
	ticker.Stop()
	fmt.Println("ticker stopped")

	jobs1:=make(chan int,100)
	results:=make(chan int,100)

	for i:=1;i<=5;i++{
		go workerpool(i,jobs1,results)
	}

	for j:=1;j<=10;j++{
		jobs1<-j
	}
	close(jobs1)

	for k:=1;k<=5;k++{
		ret:=<-results
		fmt.Println("ret:",ret)
	}
	//time.Sleep(time.Second*1)

	requests:=make(chan int,5)
	for i:=1;i<=5;i++{
		requests<-i
	}
	close(requests)

	limiter:=time.Tick(time.Millisecond*200)

	for req:=range requests{
		<-limiter
		fmt.Println("request-1",req,time.Now())
	}

	burstyLimiter:=make(chan time.Time,3)

	for i:=0;i<3;i++{
		burstyLimiter<-time.Now()
	}

	go func(){
		for t:=range time.Tick(time.Millisecond*200){
			burstyLimiter<-t
		}
	}()

	burstyRequests:=make(chan int,5)
	for i:=1;i<=5;i++{
		burstyRequests<-i
	}
	close(burstyRequests)

	for req:=range burstyRequests{
		<-burstyLimiter
		fmt.Println("request-2",req,time.Now())
	}

	var ops uint64=0

	for i:=0;i<50;i++{
		go func(){
			for{
				atomic.AddUint64(&ops,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)

	opsFinal:=atomic.LoadUint64(&ops)
	fmt.Println("ops:",opsFinal)

	var state=make(map[int]int)
	var mutex=&sync.Mutex{}

	var readops uint64=0
	var writeops uint64=0

	for i:=0;i<100;i++{
		go func(){
			total:=0
			for{
				key:=rand.Intn(5)
				mutex.Lock()
				total+=state[key]
				mutex.Unlock()
				atomic.AddUint64(&readops,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for j:=0;j<10;j++{
		go func(){
			for{
				key:=rand.Intn(5)
				val:=rand.Intn(100)
				mutex.Lock()
				state[key]=val
				mutex.Unlock()
				atomic.AddUint64(&writeops,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)

	readopsFinal:=atomic.LoadUint64(&readops)
	fmt.Println("readops",readopsFinal)
	writeopsFinal:=atomic.LoadUint64(&writeops)
	fmt.Println("writeops",writeopsFinal)

	mutex.Lock()
	fmt.Println("state",state)
	mutex.Unlock()

    type readops1 struct{
    	key int
    	resp chan int
    }
    type writeops1 struct{
    	key int
    	val int
    	resp chan bool
    }
	reads:=make(chan *readops1)
	writes:=make(chan *writeops1)

	go func(){
		var state=make(map[int]int)
		for{
			select{
			case read:=<-reads:
				read.resp<-state[read.key]
			case write:=<-writes:
				state[write.key]=write.val
				write.resp<-true
			}
		}
	}()
	for r:=0;r<100;r++{
		go func(){
			for{
				read:=&readops1{
					key:rand.Intn(5),
					resp:make(chan int)}
				reads<-read
				<-read.resp
				atomic.AddUint64(&readops,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for w:=0;w<10;w++{
		go func(){
			for{
				write:=&writeops1{
					key:rand.Intn(5),
					val:rand.Intn(100),
					resp:make(chan bool)}
				writes<-write
				<-write.resp
				atomic.AddUint64(&writeops,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)

	readOpsFinal1:=atomic.LoadUint64(&readops)
	fmt.Println("readops-1",readOpsFinal1)
	writeopsFinal1:=atomic.LoadUint64(&writeops)
	fmt.Println("writeops-1",writeopsFinal1)

	strs:=[]string{"c","a","b"}
	sort.Strings(strs)
	fmt.Println("strings:",strs)

	ints:=[]int{5,8,2}
	sort.Ints(ints)
	fmt.Println("ints:",ints)

	s:=sort.IntsAreSorted(ints)
	fmt.Println("sorted:",s)

	furits:=[]string{"apple","peach","banana","kiwi"}
	sort.Sort(ByLength(furits))
	fmt.Println(furits)
    
	//panic("a problem")
	// _,err:=os.Create("./tmp/file")
	// if err !=nil{
	// 	panic(err)
	// }

	file1:=createFile("./defer.txt")
	defer closeFile(file1)
	writeFile(file1)

	var strs1=[]string{"peach","apple","pear","plum"}

	fmt.Println(Index(strs1,"pear"))
	fmt.Println(Include(strs1,"apple"))

	fmt.Println(Any(strs1,func(v string)bool{
		return strings.HasPrefix(v,"p")
		}))

	fmt.Println(All(strs1,func(v string)bool{
		return strings.HasSuffix(v,"p")
		}))

	fmt.Println(Filter(strs1,func(v string)bool{
		return strings.Contains(v,"e")
		}))

	fmt.Println(Map(strs1,strings.ToUpper))

	P("Contains:",strings.Contains("test","es"))
	P("Count",strings.Count("test","t"))
	P("HasPrefix",strings.HasPrefix("test","te"))
	P("HasSuffix",strings.HasSuffix("test","st"))
	P("Index",strings.Index("test","e"))
	P("Join",strings.Join([]string{"a","b","c"},"-"))
	P("Repeat",strings.Repeat("a",5))
	P("Replace",strings.Replace("foo","o","0",-1))
	P("Replace",strings.Replace("foo","o","0",1))
	P("Split",strings.Split("a-b-c-d-e","-"))
	P("ToLower",strings.ToLower("TEST"))
	P("ToUpper",strings.ToUpper("test"))
	P()
	P("len",len("hello"))
	P("Char","hello"[1])

	match,_:=regexp.MatchString("p([a-z]+)ch","peach")
	fmt.Println(match)

	r1,_:=regexp.Compile("p([a-z]+)ch")

	fmt.Println(r1.MatchString("peach"))

	fmt.Println(r1.FindString("peach punch"))

	fmt.Println(r1.FindStringIndex("peach punch"))

	fmt.Println(r1.FindStringSubmatch("peach punch"))

	in:=[]byte("a peach")
	out:=r1.ReplaceAllFunc(in,bytes.ToUpper)
	fmt.Println(string(out))

	bolB,_:=json.Marshal(true)
	fmt.Println(string(bolB))

	intB,_:=json.Marshal(1)
	fmt.Println(string(intB))

	fltB,_:=json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB,_:=json.Marshal("goher")
	fmt.Println(string(strB))

	slcD:=[]string{"apple","banana"}
	slcB,_:=json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD:=map[string]int{"apple":1,"banana":2}
	mapB,_:=json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D:=&Response1{
		Page:1,
		Fruits:[]string{"apple","banana","pear"}}
	res1B,_:=json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D:=&Response2{
		Page:1,
		Fruits:[]string{"apple","banana","peach"}}
	res2B,_:=json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt:=[]byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err:=json.Unmarshal(byt,&dat);err!=nil{
		panic(err)
	}
    fmt.Println(dat)

    num:=dat["num"].(float64)
    fmt.Println(num)

    str1:=dat["strs"].([]interface{})
	str2:=str1[0].(string)
	fmt.Println(str2)

	str3:=`{"page":1,"fruits":["apple","peach"]}`
	res1:=Response2{}
	json.Unmarshal([]byte(str3),&res1)
	fmt.Println(res1)
	fmt.Println(res1.Fruits[0])

	enc:=json.NewEncoder(os.Stdout)
	d1:=map[string]int{"apple":5,"lettuce":7}
	enc.Encode(d1)

	now:=time.Now()
	secs:=now.Unix()
	nanos:=now.UnixNano()
	fmt.Println(now)

	millis:=nanos/100000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs,0))
	fmt.Println(time.Unix(0,nanos))

	t1:=time.Now()
	P(t1.Format(time.RFC3339))

	t2,e1:=time.Parse(time.RFC3339,
		"2018-01-24T17:24:33+00:00")
	P(t2)
	P(e1)

	P(t1.Format("3:04AM"))
	P(t1.Format("Mon Jan _2 15:04:05 2017"))
	P(t1.Format("2006-01-01T15:04:05.999999-07:00"))
	form:="3 04 PM"
	t3,e2:=time.Parse(form,"8 41 PM")
	P(t3)
	P(e2)

	f2,_:=strconv.ParseFloat("1.234",64)
	fmt.Println(f2)

	f3,_:=strconv.ParseInt("123",0,64)
	fmt.Println(f3)

	k1,_:=strconv.Atoi("123")
	fmt.Println(k1)

	_,e4:=strconv.Atoi("wat")
	fmt.Println(e4)

	surl := "postgres://user:pass@host.com:5432/path?k=v#f"
	u,err:=url.Parse(surl)
	if err!=nil{
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	pwd,_:=u.User.Password()
	fmt.Println(pwd)

	fmt.Println(u.Host)

	host,port,_:=net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m1,_:=url.ParseQuery(u.RawQuery)
	fmt.Println(m1)
	fmt.Println(m1["k"][0])

	s3:="sha1 this string"

	h3:=sha1.New()

	h3.Write([]byte(s3))
	bs:=h3.Sum(nil)

	fmt.Println(s3)
	fmt.Printf("%x\n",bs)

	data1 := "abc123!?$*&()'-=@~"
	sEnc:=b64.StdEncoding.EncodeToString([]byte(data1))
	fmt.Println(sEnc)

	sDec,_:=b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc:=b64.URLEncoding.EncodeToString([]byte(data1))
	fmt.Println(uEnc)
	uDec,_:=b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

	os.Setenv("Foo","test")
	fmt.Println("FOO:",os.Getenv("FOO"))
	fmt.Println("BAR:",os.Getenv("BAR"))

	fmt.Println()
	for _,e:=range os.Environ(){
		pair:=strings.Split(e,"=")
		fmt.Println(pair[0])
	}

	datecmd:=exec.Command("date")

	dateOut,err:=datecmd.Output()
	if err!=nil{
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepcmd:=exec.Command("grep","hello")

	grepIn,_:=grepcmd.StdinPipe()
	grepOut,_:=grepcmd.StdoutPipe()
	grepcmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes,_:=ioutil.ReadAll(grepOut)
	grepcmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))


	lscmd:=exec.Command("bash","-c","ls  -a -l -h")
	lsout,err:=lscmd.Output()

	if err!=nil{
		panic(err)
	}
	fmt.Println(">ls -a -l -h")
	fmt.Println(string(lsout))

	binary,lookptr:=exec.LookPath("ls")
	if lookptr!=nil{
		panic(lookptr)
	}

	args:=[]string{"ls","-a","-l","-h"}

	env:=os.Environ()

	execErr:=syscall.Exec(binary,args,env)
	if execErr!=nil{
		panic(execErr)
	}
}

type Response1 struct{
	Page int
	Fruits[]string
}
type Response2 struct{
	Page int `json:"page"`
	Fruits[]string `json:"fruits"`
}

func Index(vs []string,t string)int{
	for i,v:=range vs{
		if v==t{
			return i
		}
	}
	return -1
}

func Include(vs []string,t string)bool{
	return Index(vs,t)>=0
}

func Any(vs []string,f func(string)bool)bool{
	for _,v:=range vs{
		if f(v){
			return true
		}
	}
	return false
}
func All(vs []string,f func(string)bool)bool{
	for _,v:=range vs{
		if !f(v){
			return false
		}
	}
	return true
}

func Filter(vs []string,f func(string)bool)[]string{
	vsf:=make([]string,0)
	for _,v:=range vs{
		if f(v){
			vsf=append(vsf,v)
		}
	}
	return vsf
}
func Map(vs []string,f func(string) string) []string{
	vsm:=make([]string,len(vs))
	for i,v:=range vs{
		vsm[i]=f(v)
	}
	return vsm
}

func createFile(p string)*os.File{
	fmt.Println("creating")
	f,err:=os.Create(p)
	if err!=nil{
		panic(err)
	}
	return f
}
func writeFile(f *os.File){
	fmt.Println("writing")
	fmt.Fprintln(f,"data")
}
func closeFile(f *os.File){
	fmt.Println("closing")
	f.Close()
}
type ByLength[]string
func (s ByLength)Len() int{
	return len(s)
}
func (s ByLength)Swap(i,j int){
	s[i],s[j]=s[j],s[i]
}
func (s ByLength)Less(i,j int)bool{
	return len(s[i])<len(s[j])
}
func ping(pings chan<- string,msg string){
	pings<-msg
}
func pong(pings <-chan string,pongs chan<- string){
	msg:=<-pings
	pongs<-msg
}
func workerpool(id int,jobs <-chan int,result chan<- int){
	for j:=range jobs{
		fmt.Println("worker",id,"start job",j)
		time.Sleep(time.Second)
		fmt.Println("worker",id,"stop job",j)
		result<- j*2
	}
}
func worker(done chan bool){
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done<-true
}
func plus(a int,b int)int{
	return a+b
}
func plusplus(a,b,c int)int{
	return a+b+c
}
func vals()(int,int){
	return 3,7
}
func sum1(nums ...int){
	fmt.Println(nums,"")
	total:=0
	for _,num:=range nums{
		total+=num
	}
	fmt.Println(total)
}

func inSeq()func() int{
	i:=0
	return func() int{
		i+=1
		return i
	}
}

func fact(n int) int{
	if n==0{
		return 1
	}
	return n*fact(n-1)
}

func zeroval(ival int){
	ival=0
}

func zeroptr(ival *int){
	*ival =0
}

type person struct{
	name string
	age int
}

type rect struct{
	width,height int
}

func (r *rect) area() int{
	return r.width+r.height
}
func (r rect)perim() int{
	return 2*r.width+2*r.height
}

type geometry interface{
	area64() float64
	perim64() float64
}
type rect64 struct{
	width,height float64
}
type circle struct{
	radius float64
}

func (r rect64) area64() float64{
	return r.width*r.height
}
func (r rect64) perim64() float64{
	return 2*r.width+2*r.height
}
func (c circle) area64() float64{
	return math.Pi*c.radius*c.radius
}

func (c circle) perim64() float64{
	return 2*math.Pi*c.radius
}
func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area64())
	fmt.Println(g.perim64())
}

func f1(arg int)(int ,error){
	if arg==42{
		return -1,errors.New("can't work with 42")
	}
	return arg+3,nil
}

type argError struct{
	arg int
	prob string
}

func(e *argError)Error() string{
	return fmt.Sprintf("%d-%s",e.arg,e.prob)
}
func f2(arg int)(int ,error){
	error1:=argError{arg,"can not work with 42"}
	if arg==42{
		fmt.Println("error:",error1.Error())
	}
	if arg==42{
		return -1,&argError{arg,"can't work with 42"}
	}
	return arg+3,nil
}
func f3(from string){
	for i:=0;i<3;i++{
		fmt.Println(from,":",i)
	}
}