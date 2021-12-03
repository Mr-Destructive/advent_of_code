import java.io.File;  
import java.io.FileNotFoundException;  
import java.util.Arrays;
import java.util.Scanner;

public class d3 {
    public static void main(String args[]){
        
        String binary= "";
        String gamma= "";
        String epsilon= "";
        int one, zero;
        int[] freq= new int[24];
        try 
        {
            File myObj = new File("inp3.txt");
            Scanner myReader = new Scanner(myObj);

            while (myReader.hasNextLine()) 
            {
                one=0;zero=0;
                binary = myReader.nextLine();
                for(int i=0;i<12;i++)
                {
                    if(Character.compare(binary.charAt(i),'1')==0)
                        freq[i]+=1;
                    else
                        freq[i+12]+=1;
                }
            }
            myReader.close();
        } 

        catch (FileNotFoundException e) {
            System.out.println("An error occurred.");
            e.printStackTrace();
        }
        for(int i=0;i<12;i++)
        {
            if(freq[i]>freq[i+12])
            {
                gamma+="0";
                epsilon+="1";
            }
            else
            {
                gamma+="1";
                epsilon+="0";
            }
        }
        int n=gamma.length(),gam=0,eps=0;

        gam = Integer.parseInt(gamma,2);
        eps = Integer.parseInt(epsilon,2);

        System.out.println(eps*gam);
    }
}
