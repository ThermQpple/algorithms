package rbTree

import (
	"flag"
	"testing"
	"algorithms/tree"
	"algorithms/tree/binaryTree/genericBinaryTree"
)

var debug = flag.Bool("debug", false, "debug flag")

func TestRBT_Insert(t *testing.T) {
	arr := tree.RandomSlice(0, 20, 10)
	nodeCnt := 0
	rbt := New()
	for i, v := range arr {
		rbt.Insert(uint32(v))
		stop := rbt.PreOrderWalk(rbt.Root(), genericBinaryTree.CheckGBT(t, &nodeCnt, *debug))
		if stop {
			return
		}
		if nodeCnt != i+1 {
			t.Log("node cnt expect to ", i+1, "but get:", nodeCnt)
			t.Fail()
		}
		nodeCnt = 0
		stop = checkRBT(t, rbt)
		if stop {
			return
		}
	}
}

func TestRBT_Delete(t *testing.T) {
	arr := tree.RandomSlice(0, 20, 10)
	deleteSequence := tree.RandomSlice(0, 10, 10)
	nodeCnt := 0
	rbt := New()
	for _, v := range arr {
		rbt.Insert(uint32(v))
	}
	for i, v := range deleteSequence {
		rbt.Delete(uint32(arr[v]))
		stop := rbt.PreOrderWalk(rbt.Root(), genericBinaryTree.CheckGBT(t, &nodeCnt, *debug))
		if stop {
			return
		}
		if nodeCnt != len(deleteSequence)-1-i {
			t.Log("node cnt expect to ", len(deleteSequence)-1-i, "but get:", nodeCnt)
			t.Fail()
		}
		nodeCnt = 0
		if i != len(deleteSequence)-1 {
			stop = checkRBT(t, rbt)
			if stop {
				return
			}
		}
	}
}