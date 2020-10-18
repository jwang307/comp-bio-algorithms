import java.util.Scanner;

public class Mendels {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String[] input = in.nextLine().split(" ");
        double k = Double.parseDouble(input[0]);
        double m = Double.parseDouble(input[1]);
        double n = Double.parseDouble(input[2]);

        double sum = k + m + n;

        double homorecessive = n/sum*(n-1)/(sum-1);
        double mix = (n / sum) * (m / (sum - 1))/2;
        double mix1 = (m/sum) * (n/(sum-1))/2;
        double hetero = (m / sum)*((m-1)/(sum-1))/4;
        System.out.print(1-homorecessive-mix-mix1-hetero);
    }
}
