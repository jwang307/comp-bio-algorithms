import java.util.Scanner;

public class FibonacciRabbits {

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int m = in.nextInt();
        long[][] dp = new long[n][3];
        dp[0][0] = 0;
        dp[0][1] = 0;
        dp[0][2] = 1;
        for (int i = 1; i < n; i++) {
            if (i < m) {
                dp[i][0] = 0;
                dp[i][2] = dp[i-1][1];
                dp[i][1] = dp[i-1][1] + dp[i - 1][2];
            } else {
                dp[i][0] = dp[i-m][2];
                dp[i][1] = dp[i - 1][1] + dp[i - 1][2] - dp[i][0];
                dp[i][2] = dp[i - 1][1];
            }
        }
        if (dp[n - 1][1] + dp[n - 1][2] > 0) {
            System.out.print(dp[n - 1][1] + dp[n - 1][2]);
        } else {
            System.out.print(0);
        }
    }
}
