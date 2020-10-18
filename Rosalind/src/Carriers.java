import java.util.Scanner;

public class Carriers {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);

        String[] recessive = in.nextLine().split(" ");
        double[] recessiveprop = new double[recessive.length];
        for (int i = 0; i < recessive.length; i++) {
            recessiveprop[i] = Double.parseDouble(recessive[i]);

            double homoRecessive = recessiveprop[i];
            double homoDominant = 1.0 - Math.sqrt(homoRecessive);
            double carrier = 1.0 - Math.pow(homoDominant, 2);

            System.out.print(carrier + " ");
        }
    }
}
