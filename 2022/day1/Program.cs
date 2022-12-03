
/*
example:
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
*/

internal class Program
{
    // private static async Task<string> ReadInput()
    // {
    //     using FileStream fileStream = File.Open("./input.txt", FileMode.Open);
    //     using StreamReader streamReader = new StreamReader(fileStream);
        
    //     return await streamReader.ReadToEndAsync();
    // }

    private static async Task Main(string[] args)
    {
        //pipe file into cin
        var input = await Console.In.ReadToEndAsync();
        
        // Split by elf, then sum each elf, sort by descending
        var elfs = input
            .Trim()
            .Split("\n\n")
            .Select(text => text
                .Split('\n')
                .Select(v => int.Parse(v))
                .Sum())
            .OrderByDescending(v => v)
            .ToArray();

        //First Star
        var topElf = elfs.First();
        Console.WriteLine($"Top Elf: {topElf}");
        
        //Second Star
        var topThreeElfs = elfs.Take(3).ToArray();
        Console.WriteLine($"Top three Elfs: {String.Join(" + ", topThreeElfs)} = {topThreeElfs.Sum()}");
    }
}