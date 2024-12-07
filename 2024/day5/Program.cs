string input = File.ReadAllText("input.txt");

string[] parts = input.Split("\r\n\r\n", StringSplitOptions.None);

string rules = parts[0];
string updates = parts[1];
int ans = 0;
int ans2 = 0;

foreach(var update in updates.Split("\r\n")) {
    string[] nums = update.Split(",");
    List<string[]> rulesToCheck = [];

    foreach(var rule in rules.Split("\r\n")) {
        string[] ruleNums = rule.Split("|");

        if(nums.Contains(ruleNums[0]) && nums.Contains(ruleNums[1])) {
            rulesToCheck.Add(ruleNums);
        }
    }

    bool ok = true;
    foreach(var rule in rulesToCheck) {
        int before = Array.IndexOf(nums, rule[0]);
        int after = Array.IndexOf(nums, rule[1]);

        if(before >= after) {
            ok = false;
            break;
        }
    }

    if(ok) {
        int middle = nums.Length / 2;
        ans += int.Parse(nums[middle]);
    } else {
        Console.WriteLine("incorrect");
        Console.WriteLine(update);
        for(int i = 0; i < nums.Length; i++) {
            for(int j = 0; j < nums.Length - 1; j++) { // bubble sort it is :DDDD
                string[]? rule = rulesToCheck.Find(r => r[0] == nums[i] && r[1] == nums[j] || r[0] == nums[j] && r[1] == nums[i]);
                bool swap = false;
                if(rule != null) {
                    if(i < j && rule[0] == nums[j] && rule[1] == nums[i]) {
                        swap = true;
                    }
                    if(i > j && rule[0] == nums[i] && rule[1] == nums[j]) {
                        Console.WriteLine("Swap");
                        swap = true;
                    }
                }

                if(swap) {
                    (nums[i], nums[j]) = (nums[j], nums[i]);
                }
            }
        }   
        int middle = nums.Length / 2;
        ans2 += int.Parse(nums[middle]);
    }
}

Console.WriteLine(ans);
Console.WriteLine(ans2);