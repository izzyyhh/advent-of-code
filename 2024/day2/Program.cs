string[] lines = File.ReadAllLines("input.txt");
int count = 0;
int count2 = 0;

foreach(string line in lines) {
    string[] report = line.Split(" ");
    List<int> levels = report.Select(int.Parse).ToList();

    if(isSafeLevel(levels)) {
        count++;
    }

    for(int i = 0; i < levels.Count; i++) {
        List<int> dampenedLevel = levels.Take(i).Concat(levels.Skip(i + 1)).ToList();
        if(isSafeLevel(dampenedLevel)) {
            count2++;
            break;
        } 
    }
}

static bool isSafeDistance(int distance) {
    return Math.Abs(distance) <= 3 && Math.Abs(distance) >= 1;
}

static bool isSafeLevel(List<int> levels) {
    List<int> distances = getDistances(levels);
    return distances.All(isSafeDistance) && (distances.All(d => d > 0) || distances.All(d => d < 0));
}

static List<int> getDistances(List<int> levels) {
    List<int> distances = [];

    for(int i = 0; i < levels.Count - 1; i++) {
        int level = levels[i];
        int nextLevel = levels[i + 1];
        distances.Add(level - nextLevel);
    }

    return distances;
}

Console.WriteLine(count);
Console.WriteLine(count2);

