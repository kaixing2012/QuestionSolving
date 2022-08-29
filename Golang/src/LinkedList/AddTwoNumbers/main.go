package main

import (
	"fmt"
	"strconv"

	"eyr.question.solving/pkg/data_structures"
)

func main() {
	// https://leetcode.com/problems/add-two-numbers/

	listOne := data_structures.LinkedList{}
	listOne.Insert(7)
	listOne.Insert(8)
	listOne.Insert(9)

	listTwo := data_structures.LinkedList{}
	listTwo.Insert(6)
	listTwo.Insert(5)
	listTwo.Insert(4)

	iter := addTwoNumbers(listOne.Head, listTwo.Head)

	for iter != nil {
		fmt.Println(iter.Val)

		iter = iter.Next
	}
}

func addTwoNumbers(l1 *data_structures.LinkNode, l2 *data_structures.LinkNode) *data_structures.LinkNode {
	var carry uint64 = 0
	var result = &data_structures.LinkNode{}
	var pointer = result

	for l1 != nil || l2 != nil || carry > 0 {

		var firstNum uint64 = 0

		if l1 != nil {
			firstNumStr := fmt.Sprint(l1.Val)
			firstNum, _ = strconv.ParseUint(firstNumStr, 0, 10)
		}

		var secondNum uint64 = 0
		if l2 != nil {
			secondNumStr := fmt.Sprint(l2.Val)
			secondNum, _ = strconv.ParseUint(secondNumStr, 0, 10)
		}

		var sum = firstNum + secondNum + carry
		var num = sum % 10
		carry = sum / 10

		pointer.Next = &data_structures.LinkNode{Val: num, Next: nil}
		pointer = pointer.Next

		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}
	}

	return result.Next
}