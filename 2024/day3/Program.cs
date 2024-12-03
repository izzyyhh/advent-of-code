string input = File.ReadAllText("input.txt");
const string MUL = "mul("; 

List<Factors> multiplications = [];
List<Factors> enabledMultiplications = [];

bool enabled = true;

for(int i = 0; i < input.Length; i++) {
    if(i + "don't()".Length < input.Length && input.Substring(i, "don't()".Length) == "don't()") {
        enabled = false;
    }
    if(i + "don't()".Length < input.Length && input.Substring(i, "do()".Length) == "do()") {
        enabled = true;
    }

    string tmp = "";
    for(int j = i; j < input.Length && j - i < 12; j++) {
        tmp += input[j];

        if(tmp.Length == 4 && !tmp.StartsWith(MUL)) {
            break;
        } else if(tmp.Length > 4 && tmp.EndsWith(')')) {
            i = j;
            Console.WriteLine(tmp);
            multiplications.Add(getFactorsFromOperation(tmp));
            if(enabled) {
                enabledMultiplications.Add(getFactorsFromOperation(tmp));
            }
            break;
        }
    }
}

int sum = 0;

foreach(var m in multiplications){
    sum += m.a * m.b;
}
int sum2 = 0;

foreach(var m in enabledMultiplications){
    sum2 += m.a * m.b;
}

Console.WriteLine(sum);
Console.WriteLine(sum2);

static Factors getFactorsFromOperation(string op) {
    try {

    string csv = op.Substring(MUL.Length, op.Length - MUL.Length -1);
    string[] nums = csv.Split(",");

    return new Factors(int.Parse(nums[0]),
                       int.Parse(nums[1]));
    } catch {
        Console.WriteLine("invalid factors");
        return new Factors(0,0);
    }
}

struct Factors(int a, int b)
{
    public int a = a;
    public int b = b;
}
