import java.io.File;  
import java.util.ArrayList;
import java.io.FileNotFoundException;  
import java.util.Scanner;

public class d1 {
    public static void main(String args[]){
        
        ArrayList<Integer> numbers = new ArrayList<Integer>();
        try 
        {
            File myObj = new File("../inp/inp1.txt");
            Scanner myReader = new Scanner(myObj);

            while (myReader.hasNextLine()) 
            {
                numbers.add(Integer.parseInt(myReader.nextLine()));
            }
            myReader.close();
        } 

        catch (FileNotFoundException e) {
            System.out.println("An error occurred.");
            e.printStackTrace();
        }
        
        int n=numbers.size(),a=0,b=0,c=0;
        int k=numbers.get(0)+numbers.get(1)+numbers.get(2);
        for(int i=1;i<n;i++)
        {
            a=0;
            if(i+2<n)
            {
            for(int j=i;j<=i+2;j++)
                a+=numbers.get(j);

            if(a > k)
                c++;
            k=a;
            }
        }
        System.out.println(c);
    }
}
