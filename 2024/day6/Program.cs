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

int initR = 0;
int initC = 0;


for(int r = 0; r < map.Length; r++) {
    for(int c = 0; c < map[0].Length; c++) {
        if(isTraveler(map[r][c])) {
            var traveler = new Traveler{R=r, C=c};
            initR = r;
            initC = c;
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

int ans2 = 0;
// part 2, brute force every obstacle position, and see if cycles
for(int r = 0; r < map.Length; r++) {
    for(int c = 0; c < map[0].Length; c++) {
        var seen = new HashSet<string>();
        var guard = new Traveler{R=initR,C=initC};
        (int,int) dir = movingDir.GetValueOrDefault(map[guard.R][guard.C]);

        while(true) {
            if(seen.Contains($"{guard.R},{guard.C},{dir}")) {
                ans2++;
                break;
            }
            seen.Add($"{guard.R},{guard.C},{dir}");

            int newR = guard.R + dir.Item1;
            int newC = guard.C + dir.Item2;
            if(newR < 0 || newR >= map.Length || newC < 0 || newC >= map[0].Length) break;

            if(map[newR][newC] == '#' || (newR == r && newC == c)) {
                dir = rotateDir(dir);
                continue;
            } 
            guard.R = newR;
            guard.C = newC;  
        }
    }
}


Console.WriteLine(track.Count);
Console.WriteLine(ans2);

static (int, int) rotateDir((int,int) vec) {
    return(vec.Item2, -vec.Item1);
}

struct Traveler {
    public int R;
    public int C;
}
