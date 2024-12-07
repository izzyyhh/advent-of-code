string input = File.ReadAllText("input.txt");

string[] parts = input.Split("\r\n\r\n", StringSplitOptions.None);

string rules = parts[0];
string updates = parts[1];
int ans = 0;

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
    }
}

Console.WriteLine(ans);