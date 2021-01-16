func twoSum(nums []int, target int) []int {
    var n1, n2, l, i, j int;
    r := []int{0,1};
    l = len(nums);
    if(l==2){
        return r;
    }
    for i=0; i<l; i++ {
        n1=nums[i];
        for j=i+1; j<l; j++ {
            n2=target - n1;
            if(nums[j]==n2){
                return []int{i,j};
            }
        }        
    }
    return r;
}
