//Write your code here
class Calculator {
    private:
    public:
    int power(int n, int p) {
        if (n < 0 || p < 0) {
            throw invalid_argument("n and p should be non-negative");
        }
        return pow(n, p);
    }
};