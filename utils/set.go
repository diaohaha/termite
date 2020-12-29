package utils

type Empty struct { }

var empty Empty
//set类型
type Set struct {
	m map[string]interface{}
}
func SetFactory() *Set{
	return &Set{
		m:make(map[string]interface{}),
	}
}
//添加元素
func (s *Set) Push(val string,data interface{}){
	s.m[val] = data
}
func (s *Set) RandPop() interface{}{
	for v:= range s.m{
		p := s.m[v]
		s.Remove(v)
		return p
	}
	return nil
}
//删除元素
func (s *Set) Remove(val string) {
	delete(s.m, val)
}

func (s *Set) Check(val string) bool{
	_,isOK := s.m[val]
	return isOK
}
//获取长度
func (s *Set) Len() int {
	return len(s.m)
}

//清空set
func (s *Set) Clear() {
	s.m = make(map[string]interface{})
}

////遍历set
//func (s *Set) Traverse(){
//	for v := range s.m {
//		fmt.Println(v)
//	}
//}
//
////排序输出
//func (s *Set) SortTraverse(){
//	vals := make([]int, 0, s.Len())
//
//	for v := range s.m {
//		vals = append(vals, v)
//	}
//
//	//排序
//	sort.Ints(vals)
//
//	for _, v := range vals {
//		fmt.Println(v)
//	}
//}
