package Programmers;

import java.lang.*;
import java.util.*;

class 폰켓몬 {
    public int solution(int[] nums) {
        int answer = 0;
        int getPokemonNums = nums.length / 2;  
        HashMap<Integer, Integer> pokemonType = new HashMap<Integer, Integer>(10000);
        
        for (int i = 0; i < nums.length; i++) {
            pokemonType.put(nums[i], i);
        }
        
        answer = Math.min(getPokemonNums, pokemonType.size());
        
        return answer;
    }
}
