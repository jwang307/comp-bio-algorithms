import java.util.Scanner;

public class PerfectMatching {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);

        String rna = in.nextLine();
        long[] counter = new long[2];
        for (int i = 0; i < rna.length(); i++) {
            String base = rna.substring(i, i+1);

            if (base.equals("A")) {
                counter[0]++;
            } else if (base.equals("C")) {
                counter[1]++;
            }
        }

        System.out.print(factorial(counter[1])*factorial(counter[0]));
    }

    static long factorial(long n)
    {
        if (n == 0)
            return 1;

        return n*factorial(n-1);
    }
}
