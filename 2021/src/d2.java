import java.io.File;  
import java.util.ArrayList;
import java.io.FileNotFoundException;  
import java.util.Scanner;

public class d2 {
    public static void main(String args[]){
        
        int numbers = 0;
        String direction = new String("");
        int depth=0, distance=0,aim=0;

        try 
        {
            File myObj = new File("inp2.txt");
            Scanner myReader = new Scanner(myObj);
            myReader.useDelimiter("\n");

            while (myReader.hasNextLine()) 
            {
                String s = myReader.nextLine();
                String comp[] = s.trim().split("\\s+");
                
                direction = comp[0];
                numbers = Integer.parseInt(comp[1]);
                if(direction.equals("up"))
                    aim-=numbers;
                else if(direction.equals("down"))
                    aim+=numbers;
                else if(direction.equals("forward"))
                {
                    distance+=numbers;
                    depth+=aim*numbers;
                }
            }
            System.out.println(depth);
            System.out.println(distance);
            System.out.println(depth*distance);
            myReader.close();
        } 

        catch (FileNotFoundException e) {
            System.out.println("An error occurred.");
            e.printStackTrace();
        }
    }
}
