using System;
using System.IO;

namespace cs
{
  internal class Program
  {
    public static void Main(string[] args)
    {
      string[] arguments = Environment.GetCommandLineArgs();
      string PathTo = arguments[1];

      string path = Path.GetFullPath(PathTo);
      if (!File.Exists(path))
      {
        Console.WriteLine("A valid path is required.");
        return;
      } 
      if (Path.GetExtension(path) == "txt")
      {
        Console.WriteLine("Sorry, only txt files are able to be converted to lua files");
        return;
      }

      string newPath = path.Substring(0, path.Length- Path.GetExtension(path).Length ) + ".lua";
      File.Move(path,newPath);
      
    }
  }
}