package Programmers;

import java.util.*;

class 완주하지_못한_선수 {
    public String solution(String[] participant, String[] completion) {
        String answer = "";
        HashMap<String, Integer> map = new HashMap<String, Integer>(99999);
        
        for (int i = 0; i < completion.length; i++) {
            if (!map.containsKey(completion[i])) {
                map.put(completion[i], 1);
            }
            else {
                map.put(completion[i], map.get(completion[i]) + 1);
            }
        }
        
        for (int i = 0; i < participant.length; i++) {
            if (!map.containsKey(participant[i]) || map.get(participant[i]) == 0) {
                answer = participant[i];
                break;
            }
            
            map.put(participant[i], map.get(participant[i]) - 1);
        }
        
        return answer;
    }
}
