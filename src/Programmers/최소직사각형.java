package Programmers;

public class 최소직사각형 {
    public int solution(int[][] sizes) {
        int maxWidth = 0;
        int maxHeight = 0;

        for (int[] size: sizes) {
            if (size[0] < size[1]) {
                int tmp = size[0];
                size[0] = size[1];
                size[1] = tmp;
            }

            maxWidth = Math.max(maxWidth, size[0]);
            maxHeight = Math.max(maxHeight, size[1]);
        }

        return maxWidth * maxHeight;
    }
}
