import java.io.File;

public class main 
{
    public static void main(String[] argv)
    {

        if (argv.length == 0) {
            System.out.println("I require a path to be put in as an argument.");
            return;
        }

        String path = argv[0];

        File file = new File(path);

        if (!file.exists()) {
            System.out.println("Sorry, the path you have entered does not exist!");
            return;
        }

        file.renameTo(new File(file.getName().replaceFirst("\\.[^.]*$", "")+".lua"));


    }
}