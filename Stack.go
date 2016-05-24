package main
import(
	"errors"
	"sync"
)
type StackString struct{
	Things []string
	sync.Mutex
}
func (t *StackString)Push(v string){
	//TODO(doc):do better
	t.Lock()
	t.Things=append(t.Things, v)
	t.Unlock()
}
func (t *StackString)Pop()(string, error){
	if t.Size()<1 {
		return "", errors.New("Vide")
	}
	a:= t.Things[len(t.Things)-1]
	t.Lock()
	x:= t.Size()-1
	t.Things[x]=""
	t.Things=t.Things[:x]
	t.Unlock()
	return a, nil
}
func (t StackString)Size()int {
	return len(t.Things)
}
func (t StackString)Peek() string {
	return t.Things[len(t.Things)-1]
}
type StackInt struct{
	Things []int
	sync.Mutex
}
func (t *StackInt)Push(v int){
	//TODO(doc):do better
	t.Lock()
	t.Things=append(t.Things, v)
	t.Unlock()
}
func (t *StackInt)Pop()(int, error){
	if t.Size()<1 {
		return 0, errors.New("Vide")
	}
	a:= t.Things[len(t.Things)-1]
	t.Lock()
	x:= t.Size()-1
	t.Things[x]=0
	t.Things=t.Things[:x]
	t.Unlock()
	return a, nil
}
func (t StackInt)Size() int {
	return len(t.Things)
}
func (t StackInt)Peek() int {
	return t.Things[len(t.Things)-1] 
}