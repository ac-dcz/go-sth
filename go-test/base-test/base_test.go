package base

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	s := "abc-123-dcz-356"
	r := []string{"abc", "123", "dcz", "356"}
	w := Split(s, "-")
	if ok := reflect.DeepEqual(r, w); !ok {
		t.Errorf("expected: %v, result: %v\n", r, w)
	}
}

func TestSplitComplexSeq(t *testing.T) {
	s := "dczqqlhdqqsg"
	r := []string{"dcz", "lhd", "sg"}
	w := Split(s, "qq")
	if ok := reflect.DeepEqual(r, w); !ok {
		t.Errorf("expected: %v, result: %v\n", r, w)
	}
}

func TestMulCaseSplit(t *testing.T) {
	t.Run("case1", TestSplit) //运行子测试
	t.Run("case2", TestSplitComplexSeq)
}

func TestSimpleSplit(t *testing.T) { //表格驱动测试
	items := []struct {
		Name string
		In   []string
		Out  []string
	}{
		{"base-case", []string{"abc-123-dcz-356", "-"}, []string{"abc", "123", "dcz", "356"}},
		{"wrong-case", []string{"abc-123-dcz-356", ","}, []string{"abc-123-dcz-356"}},
		{"complex-case", []string{"dczqqlhdqqsg", "qq"}, []string{"dcz", "lhd", "sg"}},
		{"explode-case", []string{"dcz", ""}, []string{"d", "c", "z"}},
	}

	for _, item := range items {
		t.Run(item.Name, func(t *testing.T) {
			r := Split(item.In[0], item.In[1])
			if !reflect.DeepEqual(r, item.Out) {
				t.Errorf("Expected: %v, result: %v \n", item.Out, r)
			}
		})
	}

}

func TestSimpleSplitByTestify(t *testing.T) { //表格驱动测试
	items := []struct {
		Name string
		In   []string
		Out  []string
	}{
		{"base-case", []string{"abc-123-dcz-356", "-"}, []string{"abc", "123", "dcz", "356"}},
		{"wrong-case", []string{"abc-123-dcz-356", ","}, []string{"abc-123-dcz-356"}},
		{"complex-case", []string{"dczqqlhdqqsg", "qq"}, []string{"dcz", "lhd", "sg"}},
		{"explode-case", []string{"dcz", ""}, []string{"d", "c", "z"}},
	}

	for _, item := range items {
		t.Run(item.Name, func(t *testing.T) {
			r := Split(item.In[0], item.In[1])
			assert.Equal(t, item.Out, r) //使用Testify
		})
	}

}

func TestTimeConsumer(t *testing.T) {
	if testing.Short() { //如果开启了Short flag -short
		t.Skipf("Skip %s \n", t.Name())
	}
	time.Sleep(time.Second * 2)
}

// 并行测试
func TestParallelSplit(t *testing.T) {
	items := []struct {
		Name string
		In   []string
		Out  []string
	}{
		{"base-case", []string{"abc-123-dcz-356", "-"}, []string{"abc", "123", "dcz", "356"}},
		{"wrong-case", []string{"abc-123-dcz-356", ","}, []string{"abc-123-dcz-356"}},
		{"complex-case", []string{"dczqqlhdqqsg", "qq"}, []string{"dcz", "lhd", "sg"}},
		{"explode-case", []string{"dcz", ""}, []string{"d", "c", "z"}},
	}

	for _, item := range items {
		temp := item
		t.Run(temp.Name, func(t *testing.T) {
			t.Parallel()
			r := Split(temp.In[0], temp.In[1])
			if !reflect.DeepEqual(r, temp.Out) {
				t.Errorf("Expected: %v, result: %v \n", temp.Out, r)
			}
		})
	}
}
