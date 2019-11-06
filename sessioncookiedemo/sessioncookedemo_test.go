package sessioncookiedemo

import (
	"container/list"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"testing"
	"time"
)

type Manager struct {
	cookieName string //private cookiename
	lock sync.Mutex //protects session
	provider Provider
	maxlifetime int64
}

type Session interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

var globalSessions *Manager
//初始化一个记录Provider的map
var provides =  make(map[string]Provider)

func NewManager(provideName, cookieName string,maxlifetime int64)(*Manager,error)  {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider:provider,cookieName:cookieName,maxlifetime:maxlifetime},nil
}

func Register(name string,provider Provider)  {
	if provider == nil {
		panic("session: Register provide is nil")
	}
	if _, ok := provides[name]; ok {
		panic("session: Register called twice for provide " + name)
	}
	provides[name]=provider
}

func (manager *Manager)sessionId() string  {
	b := make([]byte,32)
	if _,err := io.ReadFull(rand.Reader,b); err != nil {
		return ""
	}
	sid := base64.URLEncoding.EncodeToString(b)
	println("sid:",sid)
	return sid
}

func (manager *Manager)SessionStart(w http.ResponseWriter, r *http.Request)(session Session)  {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie,err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w,&cookie)
	}else{
		sid,_ := url.QueryUnescape(cookie.Value)
		session,_ = manager.provider.SessionRead(sid)
	}
	return
}

func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}

func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxlifetime)
	time.AfterFunc(time.Duration(manager.maxlifetime), func() { manager.GC() })
}

//初始化一个空链表
var pder = &MProvider{list: list.New()}

type MProvider struct {
	lock sync.Mutex
	sessions map[string]*list.Element
	list *list.List
}

type SessionStore struct {
	sid string
	timeAccessed time.Time
	value  map[interface{}]interface{}
}

func (st *SessionStore)Set(key,value interface{}) error{
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore)Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	}else {
		return nil
	}
	return nil
}

func (st *SessionStore)Delete(key interface{})error  {
	delete(st.value,key)
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore)SessionID()string {
	return st.sid
}

func (pder *MProvider)SessionInit(sid string) (Session,error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{},0)
	newess := &SessionStore{sid:sid,timeAccessed:time.Now(),value:v}
	element := pder.list.PushBack(newess)
	pder.sessions[sid] = element
	return newess, nil
}

func (pder *MProvider) SessionRead(sid string) (Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := pder.SessionInit(sid)
		return sess, err
	}
	return nil, nil
}

func (pder *MProvider) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
		return nil
	}
	return nil
}

func (pder *MProvider)SessionGC(maxlifetime int64)  {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	for{
		element := pder.list.Back()
		if element == nil{
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.sessions,element.Value.(*SessionStore).sid)
		}else {
			break
		}
	}
}


func (pder *MProvider)SessionUpdate(sid string) error{
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}

func count(w http.ResponseWriter,r *http.Request)  {
	println(globalSessions.provider,989)
	sess := globalSessions.SessionStart(w,r)
	createtime := sess.Get("createtime")
	if createtime == nil{
		sess.Set("createtime",time.Now().Unix())
	}else if (createtime.(int64) + 360) < (time.Now().Unix()) {
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil{
		sess.Set("countnum", 1)
	}else{
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}

func login(w http.ResponseWriter,r *http.Request)  {
	session := globalSessions.SessionStart(w,r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, session.Get("username"))
	} else {
		session.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}

func init() {
	println("pder初始化")
	pder.sessions = make(map[string]*list.Element, 0)
	println("pder注册到map里，以memory为key")
	Register("memory",pder)
	println("初始化Provide的名称，Cookie名称，最大存活时间")
	globalSessions, _ = NewManager("memory","gosessionid",3600)
	println("异步并发清理session")
	go globalSessions.GC()
}

func TestSession(t *testing.T)  {
	println("启动服务...")
	http.HandleFunc("/count",count)
	http.HandleFunc("/login",login)
	println("打开浏览器http://localhost:9090")
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("Listen and server:", err)
	}
}