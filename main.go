func twoSum(nums []int, target int) []int {
    // sort.Ints(nums);
    var n2 int; // the second number
    // var m  int; // the second key
    length := len(nums);
    //fmt.Println(length);
    for i:=0; i<length; i++{
        //fmt.Println("needleNumber:");
        //fmt.Println(nums[i]);
        n2 = target - nums[i];
        
        s_nums := nums[i+1:]
        
        //l := len(s_nums);
        //fmt.Println(l);
        found := findinslice(s_nums, n2)
        if(found != -1){
             found = found+i+1;
            //fmt.Println("Found ");
            //fmt.Println(found);
            r:=[]int{i, found};
            return r;
        }
        //fmt.Println(found);
        //copy(s_nums[:],nums[i:len+1])
        //fmt.Println(n2);
        //fmt.Println(s_nums);
    }
    // 
    //fmt.Println("failed");
    r := []int{9,9};
    return r;
}

func findinslice(nums_ []int, target_ int) int {
    l := len(nums_);
    //fmt.Println("The Array");
    //fmt.Println(nums_);
    for i:=l-1; i>=0; i--{
        if(nums_[i] == target_){
            //fmt.Println("found");
            //fmt.Println(i);
            //fmt.Println(nums_[i]);
            //fmt.Println(target_);
            return i;
        }
    }
    return -1;
}
