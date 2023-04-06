using System.Diagnostics;
using Taxicab;


var sw = new Stopwatch();
decimal n = Searcher.TA6;

sw.Start();
var solutions = Searcher.ModuloSearch(n, 10);
sw.Stop();
Console.WriteLine($"Test took {sw.Elapsed.TotalSeconds}s");
            
foreach (var s in solutions)
    Console.WriteLine(s.ToString());



namespace Taxicab
{
    using System.Diagnostics;
    using System.Numerics;

    static class Searcher
    {
        public const decimal TA2 = 1729;                            // 2^10,8
        public const decimal TA3 = 87539319;                        // 2^26,4
        public const decimal TA4 = 6963472309248;                   // 2^42,7
        public const decimal TA5 = 48988659276962496;               // 2^55,4
        public const decimal TA6 = 24153319581254312065344M;        // 2^74,4
        public const decimal PTA7 = 24885189317885898975235988544M; // 2^94,3

        public static uint64 FloorLog2(decimal n)
        {
            if (n == 0)
                return 0;

            uint64 counter = 0;

            BigInteger x = new BigInteger(n);
            
            while (x != 1)
            {
                x = x >> 1;
                counter += 1;
            }
            return counter;
        }
        

        public static decimal Pow(decimal n, uint64 power)
        {
            decimal x = 1;

            for (int i = 0; i < power; i++)
            {
                x *= n;
            }
            return x;
        }
        

        public static ulong FloorCubeRoot(decimal n)
        {
            uint64 floorLog8 = FloorLog2(n) / 3;
            uint64 ceilLog8 = floorLog8 + 1;

            ulong lowGuess = (ulong) Pow((decimal)2, floorLog8);
            ulong highGuess = (ulong) Pow((decimal)2, ceilLog8);

            while (highGuess - lowGuess > 1)
            {
                ulong guess = (highGuess + lowGuess) / 2;
                if (Pow((decimal)guess, 3) > n)
                    highGuess = guess;
                else
                    lowGuess = guess;
            }
            return lowGuess;
        }
        

        public static ulong SafeXIncreasePerYDecrease(decimal target, decimal y)
        {
            //slope = y^2 / (target - y^3)^(2/3);
            decimal numerator = y * y;
            decimal denominator = Pow((decimal)FloorCubeRoot(target - y * y * y) + 1, 2); // Round up by adding one
            decimal slope = decimal.Floor(numerator / denominator);
            
            //Console.WriteLine($"Target: {target}, y: {y}, slope: {slope}");
            //return (floor(slope) - 1 / 6) * 6 + 1
            return (ulong)decimal.Truncate((slope - 1) / 6) * 6 + 1;
        }
        
        public static List<(ulong, ulong)> ModuloSearch(decimal target, int threads = 10)
        {
            if (threads < 1 || target < 1 || Math.Round(target) != target)
            {
                Console.WriteLine("Invalid arguments to function, target and thread both have to be positive nonzero integers");
                return null;
            }
            
            ulong upperBound = FloorCubeRoot(target);
            ulong lowerBound = FloorCubeRoot(target / 2);

            //Test how big the numbers will be at most
            decimal biggestNumberInCalculation = (decimal)upperBound * upperBound * upperBound - target;
            bool fitsInLong = Math.Abs(biggestNumberInCalculation) < long.MaxValue;

            //Partition the search range in [thread] chunks    
            ulong diff = upperBound - lowerBound;
            threads = (int) Math.Min((ulong) threads, diff);

            Tuple<ulong, ulong>[] ranges = new Tuple<ulong, ulong>[threads];
            for (int i = 0; i < threads; i++)
            {
                ulong size = (upperBound - lowerBound + 1) / (ulong) (threads - i);
                ranges[i] = new Tuple<ulong, ulong>(upperBound, upperBound - size + 1);
                upperBound -= size;
            }

            //Create a separate task for each chunk
            var t = new Task<List<(ulong, ulong)>>[threads];
            
            for (int i = 0; i < threads; i++)
            {
                int localI = i;
                
                t[i] = Task.Run(() =>
                    {
                        Stopwatch sw = Stopwatch.StartNew();
                        
                        ulong maxY = ranges[localI].Item1;
                        ulong minY = ranges[localI].Item2;

                        var res =
                            fitsInLong ? 
                                TraceLineSegment64(maxY, minY, target) :
                                TraceLineSegment96(maxY, minY, target);
                        
                        
                        sw.Stop();
                        Console.WriteLine($"Thread {localI} completed in {sw.Elapsed.TotalSeconds}s");
                        return res;
                    }
                );
            }
            Console.WriteLine($"Running on {threads} threads using " +
                              (fitsInLong ? "int64 algorithm" : "int96 algorithm"));
            Task.WaitAll(t.ToArray<Task>());

            List<(ulong, ulong)> results =
                (from task in t
                    from partialResults in task.Result
                    select partialResults).ToList();

            return results;
        }

        private static List<(ulong, ulong)> TraceLineSegment64(ulong maxY, ulong minY, decimal target)
        {
            var solutions = new List<(ulong, ulong)>();
            ulong y = maxY;
            ulong x = (ulong) ((target - y) % 6);

            x += FloorCubeRoot(target - (decimal) y * y * y) / 6 * 6;
            
            long relativeSum = (long) ((decimal) y * y * y + (decimal) x * x * x - target);

            while (y >= minY)
            {
                if (relativeSum > 0)
                {
                    // (y-1)^3 - y^3 + (x+1)^3 - x^3    =   3(x^2 - y^2) + 3(x+y) = 3 (x^2 + x + y - y^2)
                    relativeSum += 3 * (long) (x * x + x + y - y * y);
                    y -= 1;
                    x += 1;
                    continue;
                }

                if (relativeSum < 0)
                {
                    // (x+6)^3 - x^3    =   18*x^2 + 108x + 216
                    relativeSum += (long) (18 * x * x + 108 * x + 216);
                    x += 6;
                    continue;
                }

                solutions.Add((x, y));

                // (y-1)^3 - y^3 + (x+7)^3 - x^3    =   3(7x^2 + 49x + y - y^2 + 114)
                relativeSum += 3 * (long) (7 * x * x + 49 * x + y - y * y + 114);
                y -= 1;
                x += 7;
            }

            return solutions;
        }
        
        private static List<(ulong, ulong)> TraceLineSegment96(ulong maxY, ulong minY, decimal target)
        {
            var solutions = new List<(ulong, ulong)>();
            ulong y = maxY;
            ulong x = FloorCubeRoot(target - (decimal) y * y * y) / 6 * 6;

            x += (ulong) ((target - y) % 6);

            decimal relativeSum = (decimal) y * y * y + (decimal) x * x * x - target;
            
            while (y >= minY)
            {
                if (relativeSum > 0)
                {
                    // (y-1)^3 - y^3 + (x+slope)^3 - x^3    =   3(x^2 - y^2) + 3(x+y) = 3 (x^2 + x + y - y^2)
                    relativeSum += 3 * ((decimal)x * x + x + y - (decimal)y * y);
                    y -= 1;
                    x += 1;
                    continue;
                }

                if (relativeSum < 0)
                {
                    // (x+6)^3 - x^3    =   18*x^2 + 108x + 216
                    relativeSum += (18 * (decimal)x * x + 108 * x + 216);
                    x += 6;
                    continue;
                }

                solutions.Add((x, y));

                // (y-1)^3 - y^3 + (x+7)^3 - x^3    =   3(7x^2 + 49x + y - y^2 + 114
                relativeSum += 3 * (7 * (decimal)x * x + 49 * x + y - (decimal)y * y + 114);
                y -= 1;
                x += 7;
            }

            return solutions;
        }
        
    }
}
