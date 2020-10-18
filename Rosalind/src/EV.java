import java.util.Scanner;

public class EV {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int[] genotypeCounts = new int[6];

        for (int i = 0; i < 6; i++) {
            genotypeCounts[i] = in.nextInt();
        }
        double[] expectedValues = new double[]{2, 2, 2, 1.5, 1, 0};
        double sumEV = 0;

        for (int i = 0; i < 6; i++) {
            sumEV += genotypeCounts[i] * expectedValues[i];
        }

        System.out.print(sumEV);
    }
}
