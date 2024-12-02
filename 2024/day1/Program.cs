using System.IO;

var file = File.ReadAllLines(@"input.txt");

List<int> numbersLeft = new List<int>();
List<int> numbersRight = new List<int>();

foreach (string line in file){
    var spaces = line.Split("   ");
    int leftNumber;
    int rightNumber;

    int.TryParse(spaces[0], out leftNumber);
    int.TryParse(spaces[1], out rightNumber);

    numbersLeft.Add(leftNumber);
    numbersRight.Add(rightNumber);
}

numbersLeft.Sort();
numbersRight.Sort();

int sum = 0;

for(int i = 0; i< numbersLeft.Count; i++) {
    int left = numbersLeft[i];
    int right = numbersRight[i];
    int distance = Math.Abs(left - right);

    sum += distance;
}

Console.WriteLine(sum);

List<int> similarities = new List<int>();

for(int i = 0; i< numbersLeft.Count; i++) {
    int left = numbersLeft[i];
    int count = 0;
    foreach(int right in numbersRight) {
        if(left == right) {
            count++;
        }
    }

    similarities.Add(left * count);
}

Console.WriteLine(similarities.Sum());