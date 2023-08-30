// 1. Stack 타입을 만들어서 사용
type Stack []rune

func (stack *Stack) isEmpty() bool {
  return len(*stack) == 0
}

func (stack *Stack) push(data rune) {
  *stack = append(*stack, data)
}

func (stack *Stack) pop() rune {
  if stack.isEmpty() {
    return 0
  }
  
  top := len(*stack) - 1
  data := (*stack)[top]
  *stack = (*stack)[:top]
  return data
}

func (stack *Stack) peek() rune {
  if stack.isEmpty() {
    return 0
  }

  top := len(*stack) - 1
  return (*stack)[top]
}

func isValid(s string) bool {
    var stack Stack

    for _, data := range s {
      switch data {
        case ')':
          if stack.peek() == '(' {
            stack.pop()
          } else {
            return false
          }
        case '}':
          if stack.peek() == '{' {
            stack.pop()
          } else {
            return false
          }
        case ']':
          if stack.peek() == '[' {
            stack.pop()
          } else {
            return false
          }
        default:
          stack.push(data)
      }
    }

    if stack.isEmpty() {
      return true
    } else {
      return false
    }
}

// 2. slice를 간단하게 stack으로 이용하기
// - 따로 Stack type과 그를 위한 메소드를 만들지 않는다.
// - slice 연산을 이용해서 isEmpty(), push(), pop()의 효과를 낸다.
// - isEmpty(): len(stack) == 0
// - push(): stack = append(stack, data)
// - pop(): stack[top]과 stack = stack[:top]
// - peek(): stack[top]
// - 괄호 쌍을 map의 key-value로 저장하여 괄호의 종류마다 if문을 사용하지 않는다.
func isValid(s string) bool {
    pairs := map[rune]rune {
      '(': ')',
      '{': '}',
      '[': ']',
    }

    stack := []rune{}

    for _, data := range s {
      top := len(stack) - 1

      if _, isOpen := pairs[data]; isOpen {                        // 1) 여는 괄호인 경우 stack에 집어넣는다.(push)
        stack = append(stack, data)
      } else if len(stack) == 0 || pairs[stack[top]] != data {     // 2) 닫는 괄호인데, stack이 비어있거나 stack의 top이 자신과 쌍을 이루는 여는 괄호가 아니면 false를 리턴한다.
        return false
      } else {                                                     // 3) 닫는 괄호이고, stack의 top과 쌍을 이룬다면 stack의 top을 제거한다.(pop)
        stack = stack[:top]
      }
    }

    return len(stack) == 0                                         // 4) stack이 비어있으면 true를, 비어있지 않으면 false를 리턴한다.
}
