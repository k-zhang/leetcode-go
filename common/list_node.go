package common

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}
