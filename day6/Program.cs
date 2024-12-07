string input = File.ReadAllText("input.txt");
string[] map = input.Split("\r\n");
HashSet<string> track = [];
Dictionary<char, (int,int)> movingDir = [];

movingDir.Add('v', (1,0));
movingDir.Add('^', (-1,0));
movingDir.Add('>', (0,1));
movingDir.Add('<', (0,-1));

static bool isTraveler(char c) {
    return c == 'v' || c == '^' || c == '>' || c == '<';
}

for(int r = 0; r < map.Length; r++) {
    for(int c = 0; c < map[0].Length; c++) {
        if(isTraveler(map[r][c])) {
            var traveler = new Traveler{R=r, C=c};
            (int,int) dir = movingDir.GetValueOrDefault(map[r][c]);
            while(traveler.R < map.Length && traveler.C < map[0].Length) {
                track.Add($"{traveler.R},{traveler.C}");
                
                if(traveler.R + dir.Item1 < map.Length && traveler.C + dir.Item2 < map[0].Length && map[traveler.R + dir.Item1][traveler.C + dir.Item2] == '#') {
                    dir = rotateDir(dir);
                }
                traveler.R += dir.Item1;
                traveler.C += dir.Item2;
            }
        }
    }
}

Console.WriteLine(track.Count);

static (int, int) rotateDir((int,int) vec) {
    return(vec.Item2, -vec.Item1);
}

struct Traveler {
    public int R;
    public int C;
}
