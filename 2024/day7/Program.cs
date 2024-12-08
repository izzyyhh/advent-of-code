string[] lines = File.ReadAllLines("input.txt");
var operations = new[]{"+","*"};
double ans = 0;
double ans2 = 0;


foreach(var line in lines) {
    var parts = line.Split(":");
    // Console.WriteLine($"prob{parts[0]}");
    double target = double.Parse(parts[0]);
    List<int> nums = parts[1].Trim().Split(" ").Select(int.Parse).ToList();
    int slots = nums.Count - 1;

    List<string> combs = [];
    HashSet<string> combinatinos = [];
    GetCombinations("", slots, combinatinos);

    // Console.WriteLine($"amk{combinatinos.Count}");
    bool good = false;
    bool good2 = false;

    // Console.WriteLine(parts[1].Trim());
    foreach(var c in combinatinos) {
        // Console.WriteLine($"c{c}");
        double res = 0;
        for(int i = 1; i < nums.Count; i+=1) {
            // Console.WriteLine($"i{i}");
            bool isMulti = c[i - 1] == '*';

            if(i == 1) {
                if(isMulti) res = nums[i] * nums[i-1];
                else res = nums[i] + nums[i-1];
            } else {

                if(isMulti) {
                    res *= nums[i]; 
                } else {
                    res += nums[i];
                }
            }
        }
        // Console.WriteLine(res);
        if(res == target){
            good = true;
            break;
        } 
    }

    if(good) {
        ans += target;
    } else {
            // Console.WriteLine("checking furter");
            // Console.WriteLine(line);
            HashSet<string> newC = [];
            GetCombinations("", slots, newC, true);
            foreach (var nc in newC) {
                // Console.WriteLine(nc);
                // if(nc == "*|*") {
                //     Console.WriteLine("should fix");
                // }
                int[] copyNums = new int[nums.Count];
                nums.CopyTo(copyNums); 
                List<double> toSum = [];
                
                for(int i = 1; i < copyNums.Length; i+=1) {
                    if(i == 1) {
                        switch(nc[i-1]) {
                            case '|': {
                                // copyNums[i] = (int)double.Parse($"{copyNums[i-1]}{copyNums[i]}");
                                toSum.Add(double.Parse($"{copyNums[i-1]}{copyNums[i]}"));
                                break;
                            }
                            case '+': {
                                toSum.Add(copyNums[i-1] + copyNums[i]);
                                break;
                            }
                            case '*': {
                                toSum.Add(copyNums[i-1] * copyNums[i]);
                                break;
                            }
                        }
                    } else {
                        switch(nc[i-1]) {
                            case '|': {
                                var last = toSum[toSum.Count-1];
                                toSum.RemoveAt(toSum.Count -1);
                                toSum.Add((int)double.Parse($"{last}{copyNums[i]}"));
                                break;
                            }
                            case '+': {
                                toSum.Add(copyNums[i]);
                                break;
                            }
                            case '*': {
                                var sum = toSum.Sum();
                                if(sum == 0) {
                                    sum = 1;
                                }
                                toSum = [sum * copyNums[i]];
                                break;
                            }
                        }
                    }
                }

                if(target == toSum.Sum()) {
                    good2 = true;
                    break;
                }
            }
            if(good2) {
                // Console.WriteLine("fixed");
                // Console.WriteLine(line);
                ans2 += target;
            } 
        }
}

Console.WriteLine(ans);
Console.WriteLine(ans2 + ans);

static void GetCombinations(string combination, int slots, HashSet<string> set, bool part2 = false) {
    if(combination.Length == slots) {
        set.Add(combination);
        return;
    }
    GetCombinations(combination + "+", slots, set, part2);
    GetCombinations(combination + "*", slots, set, part2);
    if(part2) {
        GetCombinations(combination + "|", slots, set, part2);
    }
}