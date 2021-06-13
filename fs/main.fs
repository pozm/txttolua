open System.IO;;

[<EntryPoint>]
let main(args) =
    Directory.Move(args.[0], args.[0].Replace(".txt", ".lua"))
    0