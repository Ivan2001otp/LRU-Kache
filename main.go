package main

import (
	"fmt"
)

//linkedlist in golang
type Node struct{
	Value int
	Next *Node
}

const MAX_LEN = 4;

var CURR_LEN = 0;

//head of linked list
type LinkedList struct{
	Head *Node
}

func (list *LinkedList)GetCount(){
	var len int = 0;

	if list.Head!=nil{
		len = 0;
		CURR_LEN = len;
		return;
	}

	startPtr := list.Head;

	for(startPtr!=nil){
		len++;
		startPtr.Next = startPtr;
	}

	CURR_LEN = len;
}

func NewNode(value int) *Node{
	return &Node{Value: value}
}

func (list *LinkedList) AddFirst(value int){
	newNode := NewNode(value);
	newNode.Next = list.Head;
	list.Head = newNode;	
	CURR_LEN++;
}

func (list *LinkedList) AddLast(value int){
	newNode := NewNode(value);

	if (list.Head==nil){
		list.Head = newNode;
		return;
	}

	startPtr := list.Head;

	for startPtr.Next!=nil{
		startPtr = startPtr.Next
	}

	startPtr.Next = newNode;
	CURR_LEN++;
}

func (list *LinkedList) RemoveTarget(value int){
	if (list.Head==nil){
		return ;
	}


	if (list.Head.Value == value){
		list.Head = list.Head.Next;
		return;
	}

	startPtr := list.Head;

	for(startPtr.Next!=nil && startPtr.Next.Value!=value){
		
		startPtr = startPtr.Next;
	}


	if(startPtr.Next!=nil){
		startPtr.Next = startPtr.Next.Next;
		CURR_LEN--;
	}

	fmt.Println("removed ",value);


}

func (list *LinkedList)PrintList(){
	current := list.Head;

	for(current!=nil){
		fmt.Println(current.Value,"->");
		current = current.Next;
	}

}

func RemoveLastCacheItem(list *LinkedList){

	if (list.Head==nil){
		return;
	}

	if(list.Head.Next==nil){
		list.Head = nil;
		return;
	}

	startPtr := list.Head;

	for startPtr.Next.Next!=nil{
		startPtr = startPtr.Next;
	}

	startPtr.Next = nil;
	CURR_LEN--;
	
}

func CheckIfCachedItemIsPresent(list *LinkedList,val int) (bool){
	flag:=false;


	if list.Head==nil{
		return flag;
	}

	start := list.Head;

	for start!=nil{

		if(start.Value == val){
			flag = true;
			break;
		}
		start = start.Next;
	}

	return flag;
}

func (list *LinkedList)AddCacheItem(value int){

	if(CURR_LEN>MAX_LEN){
		RemoveLastCacheItem(list);
	}else{
		if(CheckIfCachedItemIsPresent(list,value)){
			list.RemoveTarget(value);
			list.AddFirst(value);
		}else{
			list.AddFirst(value);
		}
	}
}

func main(){
	fmt.Println("hi");
	list := LinkedList{};

	list.AddCacheItem(1);
	list.AddCacheItem(11);
	list.AddCacheItem(122);
	list.AddCacheItem(123);
	list.AddCacheItem(124);

	// list.RemoveTarget(11);
	RemoveLastCacheItem(&list);
	

	list.PrintList();
}