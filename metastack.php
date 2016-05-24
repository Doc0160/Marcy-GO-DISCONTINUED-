<?php
function capitalize(string $a){
	$a[0]=strtoupper($a[0]);
	return $a;
}
$types = ["string", "int"];
$nils =  ['""',       "0"];
$t = '
package main
import(
	"errors"
	"sync"
)
';
foreach($types as $k=>$type){
	$stacktype = "Stack".capitalize($type);
	$t.=
'
type '.$stacktype.' struct{
	Things []'.$type.'
	sync.Mutex
}

func (t *'.$stacktype.')Push(v '.$type.'){
	//TODO(doc):do better
	t.Lock()
	t.Things=append(t.Things, v)
	t.Unlock()
}

func (t *'.$stacktype.')Pop()('.$type.', error){
	if t.Size()<1 {
		return '.$nils[$k].', errors.New("Vide")
	}
	a:= t.Things[len(t.Things)-1]
	t.Lock()
	x:= t.Size()-1
	t.Things[x]='.$nils[$k].'
	t.Things=t.Things[:x]
	t.Unlock()
	return a, nil
}

func (t '.$stacktype.')Size()int {
	return len(t.Things)
}

func (t '.$stacktype.')Peek()'.$type.'{
	return t.Things[len(t.Things)-1] 
}
';
}
$t=trim($t);
file_put_contents("Stack.go", $t);
print($t);
