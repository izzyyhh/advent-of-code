string[] grid = File.ReadAllLines("input.txt");
int rows = grid.Length;
int columns = grid[0].Length;
List<Tuple<int, int>> directions = new()
{
    new Tuple<int, int>(-1, 0),  // Up
    new Tuple<int, int>(1, 0),   // Down
    new Tuple<int, int>(0, -1),  // Left
    new Tuple<int, int>(0, 1),   // Right
    new Tuple<int, int>(-1, -1), // Up-Left
    new Tuple<int, int>(-1, 1),  // Up-Right
    new Tuple<int, int>(1, -1),  // Down-Left
    new Tuple<int, int>(1, 1)    // Down-Right
};

for(int r = 0; r < rows; r++) {
    for(int c = 0; c < rows; c++) {
      if(grid[r][c] == 'X') {
        foreach(var d in directions) {
          // while loop? going so far bis gefunden 
        }
      }
    }
}
