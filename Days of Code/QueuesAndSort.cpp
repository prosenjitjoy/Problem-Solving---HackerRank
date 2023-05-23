class Solution {
    //Write your code here
    private:
    stack<char>s;
    queue<char>q;
    public:
    void pushCharacter(char c) {
        s.push(c);
    }
    void enqueueCharacter(char c) {
        q.push(c);
    }
    char popCharacter() {
        char top = s.top();
        s.pop();
        return top;
    }
    char dequeueCharacter() {
        char front = q.front();
        q.pop();
        return front;
    }
};