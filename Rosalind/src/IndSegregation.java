import java.math.BigInteger;
import java.util.Scanner;

public class IndSegregation {

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);

        int n = in.nextInt();
        int chromosomes = 2*n;
        double[] probabilities = new double[chromosomes];
        double[] combinations = new double[chromosomes];
        combinations[0] = 1; combinations[chromosomes - 1] = 1;
        for (int i = 1; i < chromosomes/2; i++) {
            combinations[i] = binomial(chromosomes, i).doubleValue() + combinations[i-1];
        }

        for (int i = chromosomes - 2; i >= chromosomes/2; i--) {
            combinations[i] = binomial(chromosomes, i+1).doubleValue() + combinations[i+1];
        }
        probabilities[0] = 0.000;
        double lastprob = 0.0;
        for (int i = 1; i < chromosomes; i++) {
            double combinatorics = combinations[i];
            double probability = combinatorics/Math.pow(2, chromosomes);
            if (i < chromosomes/2) {
                probability = (Math.pow(2, chromosomes) - combinatorics)/Math.pow(2, chromosomes);
            }
            double logprob = Math.log10(probability);

            probabilities[i] = logprob;
        }
        for (double d : probabilities) {
            System.out.print(d + " ");
        }
    }

    static BigInteger binomial(final int N, final int K) {
        BigInteger ret = BigInteger.ONE;
        for (int k = 0; k < K; k++) {
            ret = ret.multiply(BigInteger.valueOf(N-k))
                    .divide(BigInteger.valueOf(k+1));
        }
        return ret;
    }

}
