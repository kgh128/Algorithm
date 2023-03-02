package Baekjoon;

import java.io.*;
import java.util.*;

public class p15686 {
    static int N, M, minChickenDistance;
    static ArrayList<Point> houses;
    static ArrayList<Point> chickens;
    static int[] selectedChickens;
    static int[][] distance; // 행: 집 번호, 열: 치킨집 번호

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 집과 치킨집 정리하기
        StringTokenizer st = new StringTokenizer(br.readLine());
        N = Integer.parseInt(st.nextToken());
        M = Integer.parseInt(st.nextToken());
        houses = new ArrayList<>();
        chickens = new ArrayList<>();
        selectedChickens = new int[M];

        for (int r = 1; r <= N; r++) {
            st = new StringTokenizer(br.readLine());

            for (int c = 1; c <= N; c++) {
                int current = Integer.parseInt(st.nextToken());

                // 집인 경우
                if (current == 1) houses.add(new Point(r, c));

                // 치킨집인 경우
                if (current == 2) chickens.add(new Point(r, c));
            }
        }

        // 2. 전체 집-치킨집 사이의 거리 구하기
        // 집-치킨집 사이의 거리는 계속 사용되므로 한 번에 구해서 배열에 저장해놓고 꺼내쓰기
        distance = new int[houses.size()][chickens.size()];
        getAllDistance();

        // 3. 도시의 최소 치킨 거리 구하기
        minChickenDistance = Integer.MAX_VALUE;
        getChickenDistance(0, 0);

        // 4. 출력하기
        System.out.println(minChickenDistance);

    }

    public static class Point {
        public int r;
        public int c;

        public Point(int r, int c) {
            this.r = r;
            this.c = c;
        }
    }

    public static void getAllDistance() {
        for (int houseNum = 0; houseNum < houses.size(); houseNum++) {
            Point house = houses.get(houseNum);

            for (int chickenNum = 0; chickenNum < chickens.size(); chickenNum++) {
                Point chicken = chickens.get(chickenNum);

                distance[houseNum][chickenNum] = Math.abs(house.r - chicken.r) + Math.abs(house.c - chicken.c);
            }
        }
    }

    public static void getChickenDistance(int current, int count) {
        // 치킨집을 최대 M개 골랐을 경우
        if (count == M) {
            // 도시의 치킨 거리 구하기
            int cityChickenDistance = 0;

            for (int houseNum = 0; houseNum < houses.size(); houseNum++) {
                int houseChickenDistance = Integer.MAX_VALUE;

                // 각 집의 치킨 거리 구하기
                for (int chickenNum: selectedChickens) {
                    houseChickenDistance = Math.min(houseChickenDistance, distance[houseNum][chickenNum]);
                }

                cityChickenDistance += houseChickenDistance;
            }

            // 도시의 최소 치킨 거리 구하기
            minChickenDistance = Math.min(minChickenDistance, cityChickenDistance);
            return;
        }

        // 치킨집 고르기
        for (int chickenNum = current; chickenNum < chickens.size(); chickenNum++) {
            selectedChickens[count] = chickenNum;
            getChickenDistance(chickenNum+1, count+1);
        }
    }
}
