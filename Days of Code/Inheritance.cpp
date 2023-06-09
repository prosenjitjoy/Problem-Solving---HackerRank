class Student :  public Person{
	private:
		vector<int> testScores;  
	public:
        /*	
        *   Class Constructor
        *   
        *   Parameters:
        *   firstName - A string denoting the Person's first name.
        *   lastName - A string denoting the Person's last name.
        *   id - An integer denoting the Person's ID number.
        *   scores - An array of integers denoting the Person's test scores.
        */
        // Write your constructor here
        Student(string firstName, string lastName, int id, vector<int>scores): Person(firstName, lastName, id) {
            this->firstName = firstName;
            this->lastName = lastName;
            this->id = id;
            this->testScores = scores;
        }
        /*	
        *   Function Name: calculate
        *   Return: A character denoting the grade.
        */
        // Write your function here
        char calculate() {
            double total = 0;
            for (int i=0; i<int(this->testScores.size()); i++) {
                total += this->testScores[i];
            }
            int avg = total / int(this->testScores.size());
            
            if (avg >= 90 && avg <= 100) return 'O';
            else if (avg >= 80 && avg < 90) return 'E';
            else if (avg >= 70 && avg < 80) return 'A';
            else if (avg >= 55 && avg < 70) return 'P';
            else if (avg >= 40 && avg < 55) return 'D';
            else return 'T';
        }
};