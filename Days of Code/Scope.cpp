// Add your code here
    Difference(vector<int>elements) {
        this->elements = elements;
    }
    
    void computeDifference() {
        sort(this->elements.begin(), this->elements.end());
        this->maximumDifference = abs(this->elements[0]-this->elements[int(this->elements.size())-1]);
    }