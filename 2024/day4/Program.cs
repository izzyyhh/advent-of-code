string[] grid = File.ReadAllLines("input.txt");
int rows = grid.Length;
int columns = grid[0].Length;
List<Tuple<int, int>> directions = new()
{
    new Tuple<int, int>(-1, 0),  
    new Tuple<int, int>(1, 0),   
    new Tuple<int, int>(0, -1),  
    new Tuple<int, int>(0, 1),   
    new Tuple<int, int>(-1, -1),
    new Tuple<int, int>(-1, 1), 
    new Tuple<int, int>(1, -1), 
    new Tuple<int, int>(1, 1)   
};
int ans = 0;
int ans2 = 0;

for(int r = 0; r < rows; r++) {
    for(int c = 0; c < rows; c++) {
      if(grid[r][c] == 'X') {
        foreach(var d in directions) {
          string mas = "";
          int step = 1;
          while(mas.Length < 3) {
            try {
              mas += grid[r + d.Item1 * step][c + d.Item2 * step];
              step++;
            } catch(IndexOutOfRangeException) {
              break;
            }
          }
          if(mas == "MAS") {
            ans++;
          }
        }
      }

      if(grid[r][c] == 'A') {
            try {
                string corners = "";
                corners += grid[r+1][c-1];
                corners += grid[r+1][c+1];
                corners += grid[r-1][c+1];
                corners += grid[r-1][c-1];

                if(corners == "MMSS" || corners == "MSSM" || corners == "SMMS" || corners == "SSMM") {
                    ans2++;
                }
 
            } catch(IndexOutOfRangeException) {
                continue;
            }
      } 
    }
}

Console.WriteLine(ans);
Console.WriteLine(ans2);