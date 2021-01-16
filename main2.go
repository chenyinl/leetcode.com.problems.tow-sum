func twoSum(nums []int, target int) []int {
    r:=[]int{0,0};
    for i:=0; i<len(nums); i++ {
        // n2 := ;
        s_nums := nums[i+1:]
        //found := findinslice(s_nums, n2)
        found := findinslice(s_nums, target - nums[i]);
        if(found != -1){
             found = found+i+1;
            r=[]int{i, found};
            return r;
        }
    }
    return r;
}

func findinslice(nums_ []int, target_ int) int {
    l := len(nums_);
    for i:=l-1; i>=0; i--{
        if(nums_[i] == target_){
            return i;
        }
    }
    return -1;
}
